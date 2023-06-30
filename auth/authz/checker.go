package authz

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/go-redis/redis/v8"
	"github.com/interline-io/transitland-lib/log"
	"github.com/interline-io/transitland-server/auth/authn"
	"github.com/interline-io/transitland-server/internal/dbutil"
	"github.com/interline-io/transitland-server/internal/generated/azpb"
)

func init() {
	// Ensure Checker implements CheckerServer
	var _ azpb.CheckerServer = &Checker{}
}

type Checker struct {
	userClient   UserProvider
	fgaClient    FGAProvider
	db           sqlx.Ext
	globalAdmins []string
	azpb.UnsafeCheckerServer
}

func NewCheckerFromConfig(cfg AuthzConfig, db sqlx.Ext, redisClient *redis.Client) (*Checker, error) {
	var userClient UserProvider
	userClient = NewMockUserProvider()
	var fgaClient FGAProvider
	fgaClient = NewMockFGAClient()

	// Use Auth0 if configured
	if cfg.Auth0Domain != "" {
		auth0Client, err := NewAuth0Client(cfg.Auth0Domain, cfg.Auth0ClientID, cfg.Auth0ClientSecret)
		if err != nil {
			return nil, err
		}
		userClient = auth0Client
	}

	// Use FGA if configured
	if cfg.FGAEndpoint != "" {
		fgac, err := NewFGAClient(cfg.FGAEndpoint, cfg.FGAStoreID, cfg.FGAModelID)
		if err != nil {
			return nil, err
		}
		fgaClient = fgac
		// Create test FGA environment
		if cfg.FGALoadModelFile != "" {
			if cfg.FGAStoreID == "" {
				if _, err := fgac.CreateStore(context.Background(), "test"); err != nil {
					return nil, err
				}
			}
			if _, err := fgac.CreateModel(context.Background(), cfg.FGALoadModelFile); err != nil {
				return nil, err
			}
		}
		// Add test data
		for _, tk := range cfg.FGALoadTestData {
			ltk, found, err := ekLookup(db, tk)
			if !found {
				log.Info().Msgf("warning, tuple entities not found in database: %s", tk.String())
			}
			if err != nil {
				return nil, err
			}
			if err := fgaClient.WriteTuple(context.Background(), ltk); err != nil {
				return nil, err
			}
		}
	}

	checker := NewChecker(userClient, fgaClient, db, redisClient)
	if cfg.GlobalAdmin != "" {
		checker.globalAdmins = append(checker.globalAdmins, cfg.GlobalAdmin)
	}
	return checker, nil
}

func NewChecker(n UserProvider, p FGAProvider, db sqlx.Ext, redisClient *redis.Client) *Checker {
	return &Checker{
		userClient: n,
		fgaClient:  p,
		db:         db,
	}
}

// ///////////////////
// USERS
// ///////////////////

func (c *Checker) UserList(ctx context.Context, req *azpb.UserListRequest) (*azpb.UserListResponse, error) {
	// TODO: filter users
	users, err := c.userClient.Users(ctx, req.GetQ())
	if err != nil {
		return nil, err
	}
	return &azpb.UserListResponse{Users: users}, nil
}

func (c *Checker) User(ctx context.Context, req *azpb.UserRequest) (*azpb.UserResponse, error) {
	// Special case "*"
	if req.Id == "*" {
		user := &azpb.User{Id: "*", Name: "All users"}
		return &azpb.UserResponse{User: user}, nil
	}
	// TODO: filter users
	ret, err := c.userClient.UserByID(ctx, req.GetId())
	if ret == nil || err != nil {
		return nil, ErrUnauthorized
	}
	return &azpb.UserResponse{User: ret}, err
}

func (c *Checker) CheckGlobalAdmin(ctx context.Context) (bool, error) {
	return c.checkGlobalAdmin(authn.ForContext(ctx)), nil
}

func (c *Checker) hydrateUsers(ctx context.Context, users []*azpb.User) ([]*azpb.User, error) {
	var ret []*azpb.User
	for _, u := range users {
		uu, err := c.userClient.UserByID(ctx, u.Id)
		if err == nil && uu != nil {
			ret = append(ret, uu)
		}
	}
	return ret, nil
}

func (c *Checker) hydrateEntityRels(ctx context.Context, ers []*azpb.EntityRelation) ([]*azpb.EntityRelation, error) {
	// This is awful :( :(
	for i, v := range ers {
		if v.Type == TenantType {
			if t, _ := c.getTenants(ctx, []int64{v.Int64()}); len(t) > 0 && t[0] != nil {
				ers[i].Name = t[0].Name
			}
		} else if v.Type == GroupType {
			if t, _ := c.getGroups(ctx, []int64{v.Int64()}); len(t) > 0 && t[0] != nil {
				ers[i].Name = t[0].Name
			}
		} else if v.Type == UserType {
			if t, err := c.User(ctx, &azpb.UserRequest{Id: v.Id}); err == nil && t != nil && t.User != nil {
				ers[i].Name = t.User.Name
			}
		}
	}
	return ers, nil
}

// ///////////////////
// TENANTS
// ///////////////////

func (c *Checker) getTenants(ctx context.Context, ids []int64) ([]*azpb.Tenant, error) {
	return getEntities[*azpb.Tenant](ctx, c.db, ids, "tl_tenants", "id", "coalesce(tenant_name,'') as name")
}

func (c *Checker) TenantList(ctx context.Context, req *azpb.TenantListRequest) (*azpb.TenantListResponse, error) {
	ids, err := c.listCtxObjects(ctx, TenantType, CanView)
	if err != nil {
		return nil, err
	}
	t, err := c.getTenants(ctx, ids)
	return &azpb.TenantListResponse{Tenants: t}, err
}

func (c *Checker) Tenant(ctx context.Context, req *azpb.TenantRequest) (*azpb.TenantResponse, error) {
	tenantId := req.GetId()
	if err := c.checkActionOrError(ctx, CanView, NewEntityID(TenantType, tenantId)); err != nil {
		return nil, err
	}
	t, err := c.getTenants(ctx, []int64{tenantId})
	return &azpb.TenantResponse{Tenant: first(t)}, err
}

func (c *Checker) TenantPermissions(ctx context.Context, req *azpb.TenantRequest) (*azpb.TenantPermissionsResponse, error) {
	ent, err := c.Tenant(ctx, req)
	if err != nil {
		return nil, err
	}
	ret := &azpb.TenantPermissionsResponse{
		Tenant:  ent.Tenant,
		Actions: &azpb.TenantPermissionsResponse_Actions{},
		Users:   &azpb.TenantPermissionsResponse_Users{},
	}

	// Actions
	entKey := NewEntityID(TenantType, req.GetId())
	groupIds, _ := c.listSubjectRelations(ctx, entKey, GroupType, ParentRelation)
	ret.Groups, _ = c.getGroups(ctx, groupIds)
	ret.Actions.CanView, _ = c.checkAction(ctx, CanView, entKey)
	ret.Actions.CanEditMembers, _ = c.checkAction(ctx, CanEditMembers, entKey)
	ret.Actions.CanEdit, _ = c.checkAction(ctx, CanEdit, entKey)
	ret.Actions.CanCreateOrg, _ = c.checkAction(ctx, CanCreateOrg, entKey)
	ret.Actions.CanDeleteOrg, _ = c.checkAction(ctx, CanDeleteOrg, entKey)

	// Get tenant metadata
	tps, err := c.getObjectTuples(ctx, entKey)
	if err != nil {
		return nil, err
	}
	for _, tk := range tps {
		if tk.Relation == AdminRelation {
			ret.Users.Admins = append(ret.Users.Admins, azpb.NewEntityRelation(tk.Subject, tk.Relation))
		}
		if tk.Relation == MemberRelation {
			ret.Users.Members = append(ret.Users.Members, azpb.NewEntityRelation(tk.Subject, tk.Relation))
		}
	}
	ret.Users.Admins, _ = c.hydrateEntityRels(ctx, ret.Users.Admins)
	ret.Users.Members, _ = c.hydrateEntityRels(ctx, ret.Users.Members)
	return ret, nil
}

func (c *Checker) TenantSave(ctx context.Context, req *azpb.TenantSaveRequest) (*azpb.TenantSaveResponse, error) {
	t := req.GetTenant()
	tenantId := t.GetId()
	if check, err := c.TenantPermissions(ctx, &azpb.TenantRequest{Id: tenantId}); err != nil {
		return nil, err
	} else if !check.Actions.CanEdit {
		return nil, ErrUnauthorized
	}
	newName := t.GetName()
	log.Trace().Str("tenantName", newName).Int64("id", tenantId).Msg("TenantSave")
	_, err := sq.StatementBuilder.
		RunWith(c.db).
		PlaceholderFormat(sq.Dollar).
		Update("tl_tenants").
		SetMap(map[string]any{
			"tenant_name": newName,
		}).
		Where("id = ?", tenantId).Exec()
	return &azpb.TenantSaveResponse{}, err
}

func (c *Checker) TenantAddPermission(ctx context.Context, req *azpb.TenantModifyPermissionRequest) (*azpb.TenantSaveResponse, error) {
	tenantId := req.GetId()
	if check, err := c.TenantPermissions(ctx, &azpb.TenantRequest{Id: tenantId}); err != nil {
		return nil, err
	} else if !check.Actions.CanEditMembers {
		return nil, ErrUnauthorized
	}
	tk := req.GetEntityRelation().WithObject(NewEntityID(TenantType, tenantId))
	log.Trace().Str("tk", tk.String()).Int64("id", tenantId).Msg("TenantAddPermission")
	return &azpb.TenantSaveResponse{}, c.fgaClient.SetExclusiveSubjectRelation(ctx, tk, MemberRelation, AdminRelation)
}

func (c *Checker) TenantRemovePermission(ctx context.Context, req *azpb.TenantModifyPermissionRequest) (*azpb.TenantSaveResponse, error) {
	tenantId := req.GetId()
	if check, err := c.TenantPermissions(ctx, &azpb.TenantRequest{Id: tenantId}); err != nil {
		return nil, err
	} else if !check.Actions.CanEditMembers {
		return nil, ErrUnauthorized
	}
	tk := req.GetEntityRelation().WithObject(NewEntityID(TenantType, tenantId))
	log.Trace().Str("tk", tk.String()).Int64("id", tenantId).Msg("TenantRemovePermission")
	return &azpb.TenantSaveResponse{}, c.fgaClient.DeleteTuple(ctx, tk)
}

func (c *Checker) TenantCreate(ctx context.Context, req *azpb.TenantCreateRequest) (*azpb.TenantSaveResponse, error) {
	return &azpb.TenantSaveResponse{}, nil
}

func (c *Checker) TenantCreateGroup(ctx context.Context, req *azpb.TenantCreateGroupRequest) (*azpb.GroupSaveResponse, error) {
	tenantId := req.GetId()
	groupName := req.GetGroup().GetName()
	if check, err := c.TenantPermissions(ctx, &azpb.TenantRequest{Id: tenantId}); err != nil {
		return nil, err
	} else if !check.Actions.CanCreateOrg {
		return nil, ErrUnauthorized
	}
	log.Trace().Str("groupName", groupName).Int64("id", tenantId).Msg("TenantCreateGroup")
	groupId := int64(0)
	err := sq.StatementBuilder.
		RunWith(c.db).
		PlaceholderFormat(sq.Dollar).
		Insert("tl_groups").
		Columns("group_name").
		Values(groupName).
		Suffix(`RETURNING "id"`).
		QueryRow().
		Scan(&groupId)
	if err != nil {
		return nil, err
	}
	addTk := NewTupleKey().WithSubjectID(TenantType, tenantId).WithObjectID(GroupType, groupId).WithRelation(ParentRelation)
	if err := c.fgaClient.WriteTuple(ctx, addTk); err != nil {
		return nil, err
	}
	return &azpb.GroupSaveResponse{Group: &azpb.Group{Id: groupId}}, err
}

// ///////////////////
// GROUPS
// ///////////////////

func (c *Checker) getGroups(ctx context.Context, ids []int64) ([]*azpb.Group, error) {
	return getEntities[*azpb.Group](ctx, c.db, ids, "tl_groups", "id", "coalesce(group_name,'') as name")
}

func (c *Checker) GroupList(ctx context.Context, req *azpb.GroupListRequest) (*azpb.GroupListResponse, error) {
	ids, err := c.listCtxObjects(ctx, GroupType, CanView)
	if err != nil {
		return nil, err
	}
	t, err := c.getGroups(ctx, ids)
	return &azpb.GroupListResponse{Groups: t}, err
}

func (c *Checker) Group(ctx context.Context, req *azpb.GroupRequest) (*azpb.GroupResponse, error) {
	groupId := req.GetId()
	if err := c.checkActionOrError(ctx, CanView, NewEntityID(GroupType, groupId)); err != nil {
		return nil, err
	}
	t, err := c.getGroups(ctx, []int64{groupId})
	return &azpb.GroupResponse{Group: first(t)}, err
}

func (c *Checker) GroupPermissions(ctx context.Context, req *azpb.GroupRequest) (*azpb.GroupPermissionsResponse, error) {
	groupId := req.GetId()
	ent, err := c.Group(ctx, req)
	if err != nil {
		return nil, err
	}
	ret := &azpb.GroupPermissionsResponse{
		Group:   ent.Group,
		Users:   &azpb.GroupPermissionsResponse_Users{},
		Actions: &azpb.GroupPermissionsResponse_Actions{},
	}

	// Actions
	entKey := NewEntityID(GroupType, groupId)
	ret.Actions.CanView, _ = c.checkAction(ctx, CanView, entKey)
	ret.Actions.CanEditMembers, _ = c.checkAction(ctx, CanEditMembers, entKey)
	ret.Actions.CanEdit, _ = c.checkAction(ctx, CanEdit, entKey)
	ret.Actions.CanCreateFeed, _ = c.checkAction(ctx, CanCreateFeed, entKey)
	ret.Actions.CanDeleteFeed, _ = c.checkAction(ctx, CanDeleteFeed, entKey)
	ret.Actions.CanSetTenant = c.ctxIsGlobalAdmin(ctx)

	// Get feeds
	feedIds, _ := c.listSubjectRelations(ctx, entKey, FeedType, ParentRelation)
	ret.Feeds, _ = c.getFeeds(ctx, feedIds)

	// Get group metadata
	tps, err := c.getObjectTuples(ctx, entKey)
	if err != nil {
		return nil, err
	}
	for _, tk := range tps {
		if tk.Relation == ParentRelation {
			ct, _ := c.Tenant(ctx, &azpb.TenantRequest{Id: tk.Subject.ID()})
			ret.Tenant = ct.Tenant
		}
		if tk.Relation == ManagerRelation {
			ret.Users.Managers = append(ret.Users.Managers, azpb.NewEntityRelation(tk.Subject, tk.Relation))
		}
		if tk.Relation == EditorRelation {
			ret.Users.Editors = append(ret.Users.Editors, azpb.NewEntityRelation(tk.Subject, tk.Relation))
		}
		if tk.Relation == ViewerRelation {
			ret.Users.Viewers = append(ret.Users.Viewers, azpb.NewEntityRelation(tk.Subject, tk.Relation))
		}
	}
	ret.Users.Managers, _ = c.hydrateEntityRels(ctx, ret.Users.Managers)
	ret.Users.Editors, _ = c.hydrateEntityRels(ctx, ret.Users.Editors)
	ret.Users.Viewers, _ = c.hydrateEntityRels(ctx, ret.Users.Viewers)
	return ret, nil
}

func (c *Checker) GroupSave(ctx context.Context, req *azpb.GroupSaveRequest) (*azpb.GroupSaveResponse, error) {
	group := req.GetGroup()
	groupId := group.GetId()
	newName := group.GetName()
	if check, err := c.GroupPermissions(ctx, &azpb.GroupRequest{Id: groupId}); err != nil {
		return nil, err
	} else if !check.Actions.CanEdit {
		return nil, ErrUnauthorized
	}
	log.Trace().Str("groupName", newName).Int64("id", groupId).Msg("GroupSave")
	_, err := sq.StatementBuilder.
		RunWith(c.db).
		PlaceholderFormat(sq.Dollar).
		Update("tl_groups").
		SetMap(map[string]any{
			"group_name": newName,
		}).
		Where("id = ?", groupId).Exec()
	return &azpb.GroupSaveResponse{}, err
}

func (c *Checker) GroupAddPermission(ctx context.Context, req *azpb.GroupModifyPermissionRequest) (*azpb.GroupSaveResponse, error) {
	groupId := req.GetId()
	if check, err := c.GroupPermissions(ctx, &azpb.GroupRequest{Id: groupId}); err != nil {
		return nil, err
	} else if !check.Actions.CanEditMembers {
		return nil, ErrUnauthorized
	}
	tk := req.GetEntityRelation().WithObject(NewEntityID(GroupType, groupId))
	log.Trace().Str("tk", tk.String()).Int64("id", groupId).Msg("GroupAddPermission")
	return &azpb.GroupSaveResponse{}, c.fgaClient.SetExclusiveSubjectRelation(ctx, tk, ViewerRelation, EditorRelation, ManagerRelation)
}

func (c *Checker) GroupRemovePermission(ctx context.Context, req *azpb.GroupModifyPermissionRequest) (*azpb.GroupSaveResponse, error) {
	groupId := req.GetId()
	if check, err := c.GroupPermissions(ctx, &azpb.GroupRequest{Id: groupId}); err != nil {
		return nil, err
	} else if !check.Actions.CanEditMembers {
		return nil, ErrUnauthorized
	}
	tk := req.GetEntityRelation().WithObject(NewEntityID(GroupType, groupId))
	log.Trace().Str("tk", tk.String()).Int64("id", groupId).Msg("GroupRemovePermission")
	return &azpb.GroupSaveResponse{}, c.fgaClient.DeleteTuple(ctx, tk)
}

func (c *Checker) GroupSetTenant(ctx context.Context, req *azpb.GroupSetTenantRequest) (*azpb.GroupSetTenantResponse, error) {
	groupId := req.GetId()
	newTenantId := req.GetTenantId()
	if check, err := c.GroupPermissions(ctx, &azpb.GroupRequest{Id: groupId}); err != nil {
		return nil, err
	} else if !check.Actions.CanSetTenant {
		return nil, ErrUnauthorized
	}
	tk := NewTupleKey().WithSubjectID(TenantType, newTenantId).WithObjectID(GroupType, groupId).WithRelation(ParentRelation)
	log.Trace().Str("tk", tk.String()).Int64("id", groupId).Msg("GroupSetTenant")
	return &azpb.GroupSetTenantResponse{}, c.fgaClient.SetExclusiveRelation(ctx, tk)
}

// ///////////////////
// FEEDS
// ///////////////////

func (c *Checker) getFeeds(ctx context.Context, ids []int64) ([]*azpb.Feed, error) {
	return getEntities[*azpb.Feed](ctx, c.db, ids, "current_feeds", "id", "onestop_id", "coalesce(name,'') as name")
}

func (c *Checker) FeedList(ctx context.Context, req *azpb.FeedListRequest) (*azpb.FeedListResponse, error) {
	feedIds, err := c.listCtxObjects(ctx, FeedType, CanView)
	if err != nil {
		return nil, err
	}
	t, err := c.getFeeds(ctx, feedIds)
	return &azpb.FeedListResponse{Feeds: t}, err
}

func (c *Checker) Feed(ctx context.Context, req *azpb.FeedRequest) (*azpb.FeedResponse, error) {
	feedId := req.GetId()
	if err := c.checkActionOrError(ctx, CanView, NewEntityID(FeedType, feedId)); err != nil {
		return nil, err
	}
	t, err := c.getFeeds(ctx, []int64{feedId})
	return &azpb.FeedResponse{Feed: first(t)}, err
}

func (c *Checker) FeedPermissions(ctx context.Context, req *azpb.FeedRequest) (*azpb.FeedPermissionsResponse, error) {
	ent, err := c.Feed(ctx, req)
	if err != nil {
		return nil, err
	}
	ret := &azpb.FeedPermissionsResponse{
		Feed:    ent.Feed,
		Actions: &azpb.FeedPermissionsResponse_Actions{},
	}

	// Actions
	entKey := NewEntityID(FeedType, req.GetId())
	ret.Actions.CanView, _ = c.checkAction(ctx, CanView, entKey)
	ret.Actions.CanEdit, _ = c.checkAction(ctx, CanEdit, entKey)
	ret.Actions.CanSetGroup, _ = c.checkAction(ctx, CanSetGroup, entKey)
	ret.Actions.CanCreateFeedVersion, _ = c.checkAction(ctx, CanCreateFeedVersion, entKey)
	ret.Actions.CanDeleteFeedVersion, _ = c.checkAction(ctx, CanDeleteFeedVersion, entKey)

	// Get feed metadata
	tps, err := c.getObjectTuples(ctx, entKey)
	if err != nil {
		return nil, err
	}
	for _, tk := range tps {
		if tk.Relation == ParentRelation {
			ct, _ := c.Group(ctx, &azpb.GroupRequest{Id: tk.Subject.ID()})
			ret.Group = ct.Group
		}
	}
	return ret, nil
}

func (c *Checker) FeedSetGroup(ctx context.Context, req *azpb.FeedSetGroupRequest) (*azpb.FeedSaveResponse, error) {
	feedId := req.GetId()
	newGroup := req.GetGroupId()
	if check, err := c.FeedPermissions(ctx, &azpb.FeedRequest{Id: feedId}); err != nil {
		return nil, err
	} else if !check.Actions.CanSetGroup {
		return nil, ErrUnauthorized
	}
	tk := NewTupleKey().WithSubjectID(GroupType, newGroup).WithObjectID(FeedType, feedId).WithRelation(ParentRelation)
	log.Trace().Str("tk", tk.String()).Int64("id", feedId).Msg("FeedSetGroup")
	return &azpb.FeedSaveResponse{}, c.fgaClient.SetExclusiveRelation(ctx, tk)
}

/////////////////////
// FEED VERSIONS
/////////////////////

func (c *Checker) getFeedVersions(ctx context.Context, ids []int64) ([]*azpb.FeedVersion, error) {
	return getEntities[*azpb.FeedVersion](ctx, c.db, ids, "feed_versions", "id", "feed_id", "sha1", "coalesce(name,'') as name")
}

func (c *Checker) FeedVersionList(ctx context.Context, req *azpb.FeedVersionListRequest) (*azpb.FeedVersionListResponse, error) {
	fvids, err := c.listCtxObjects(ctx, FeedVersionType, CanView)
	if err != nil {
		return nil, err
	}
	t, err := c.getFeedVersions(ctx, fvids)
	return &azpb.FeedVersionListResponse{FeedVersions: t}, err
}

func (c *Checker) FeedVersion(ctx context.Context, req *azpb.FeedVersionRequest) (*azpb.FeedVersionResponse, error) {
	fvid := req.GetId()
	feedId := int64(0)
	// We need to get feed id before any other checks
	// If there is a "not found" error here, save it for after the global admin check
	// This is for consistency with other permission checks
	t, fvErr := c.getFeedVersions(ctx, []int64{fvid})
	fv := first(t)
	if fv != nil {
		feedId = fv.FeedId
	}
	ctxTk := NewTupleKey().WithObjectID(FeedVersionType, fvid).WithSubjectID(FeedType, feedId).WithRelation(ParentRelation)
	if err := c.checkActionOrError(ctx, CanView, NewEntityID(FeedVersionType, fvid), ctxTk); err != nil {
		return nil, err
	}
	// Now return deferred fvErr
	if fvErr != nil {
		return nil, fvErr
	}
	return &azpb.FeedVersionResponse{FeedVersion: fv}, nil
}

func (c *Checker) FeedVersionPermissions(ctx context.Context, req *azpb.FeedVersionRequest) (*azpb.FeedVersionPermissionsResponse, error) {
	ent, err := c.FeedVersion(ctx, req)
	if err != nil {
		return nil, err
	}
	ret := &azpb.FeedVersionPermissionsResponse{
		FeedVersion: ent.FeedVersion,
		Users:       &azpb.FeedVersionPermissionsResponse_Users{},
		Actions:     &azpb.FeedVersionPermissionsResponse_Actions{},
	}

	// Actions
	ctxTk := NewTupleKey().WithObjectID(FeedVersionType, ent.FeedVersion.Id).WithSubjectID(FeedType, ent.FeedVersion.FeedId).WithRelation(ParentRelation)
	entKey := NewEntityID(FeedVersionType, req.GetId())
	ret.Actions.CanView, _ = c.checkAction(ctx, CanView, entKey, ctxTk)
	ret.Actions.CanEditMembers, _ = c.checkAction(ctx, CanEditMembers, entKey, ctxTk)
	ret.Actions.CanEdit, _ = c.checkAction(ctx, CanEdit, entKey, ctxTk)

	// Get fv metadata
	tps, err := c.getObjectTuples(ctx, entKey, ctxTk)
	if err != nil {
		return nil, err
	}
	for _, tk := range tps {
		if tk.Relation == EditorRelation {
			ret.Users.Editors = append(ret.Users.Editors, azpb.NewEntityRelation(tk.Subject, tk.Relation))
		}
		if tk.Relation == ViewerRelation {
			ret.Users.Viewers = append(ret.Users.Viewers, azpb.NewEntityRelation(tk.Subject, tk.Relation))
		}
	}
	ret.Users.Editors, _ = c.hydrateEntityRels(ctx, ret.Users.Editors)
	ret.Users.Viewers, _ = c.hydrateEntityRels(ctx, ret.Users.Viewers)
	return ret, nil
}

func (c *Checker) FeedVersionAddPermission(ctx context.Context, req *azpb.FeedVersionModifyPermissionRequest) (*azpb.FeedVersionSaveResponse, error) {
	fvid := req.GetId()
	if check, err := c.FeedVersionPermissions(ctx, &azpb.FeedVersionRequest{Id: fvid}); err != nil {
		return nil, err
	} else if !check.Actions.CanEditMembers {
		return nil, ErrUnauthorized
	}
	tk := req.GetEntityRelation().WithObject(NewEntityID(FeedVersionType, fvid))
	log.Trace().Str("tk", tk.String()).Int64("id", fvid).Msg("FeedVersionAddPermission")
	return &azpb.FeedVersionSaveResponse{}, c.fgaClient.SetExclusiveSubjectRelation(ctx, tk, ViewerRelation, EditorRelation, ManagerRelation)
}

func (c *Checker) FeedVersionRemovePermission(ctx context.Context, req *azpb.FeedVersionModifyPermissionRequest) (*azpb.FeedVersionSaveResponse, error) {
	fvid := req.GetId()
	if check, err := c.FeedVersionPermissions(ctx, &azpb.FeedVersionRequest{Id: fvid}); err != nil {
		return nil, err
	} else if !check.Actions.CanEditMembers {
		return nil, ErrUnauthorized
	}
	tk := req.GetEntityRelation().WithObject(NewEntityID(FeedVersionType, fvid))
	log.Trace().Str("tk", tk.String()).Int64("id", fvid).Msg("FeedVersionRemovePermission")
	return &azpb.FeedVersionSaveResponse{}, c.fgaClient.DeleteTuple(ctx, tk)
}

// ///////////////////
// internal
// ///////////////////

func (c *Checker) listCtxObjects(ctx context.Context, objectType ObjectType, action Action) ([]int64, error) {
	checkUser := authn.ForContext(ctx)
	if checkUser == nil {
		return nil, nil
	}
	tk := NewTupleKey().WithUser(checkUser.Name()).WithObject(objectType, "").WithAction(action)
	objTks, err := c.fgaClient.ListObjects(ctx, tk)
	if err != nil {
		return nil, err
	}
	var ret []int64
	for _, tk := range objTks {
		ret = append(ret, tk.Object.ID())
	}
	return ret, nil
}

func (c *Checker) listSubjectRelations(ctx context.Context, sub EntityKey, objectType ObjectType, relation Relation) ([]int64, error) {
	tk := NewTupleKey().WithSubject(sub.Type, sub.Name).WithObject(objectType, "").WithRelation(relation)
	rels, err := c.fgaClient.ListObjects(ctx, tk)
	if err != nil {
		return nil, err
	}
	var ret []int64
	for _, v := range rels {
		ret = append(ret, v.Object.ID())
	}
	return ret, nil
}

func (c *Checker) getObjectTuples(ctx context.Context, obj EntityKey, ctxtk ...TupleKey) ([]TupleKey, error) {
	return c.fgaClient.GetObjectTuples(ctx, NewTupleKey().WithObject(obj.Type, obj.Name))
}

func (c *Checker) checkActionOrError(ctx context.Context, checkAction Action, obj EntityKey, ctxtk ...TupleKey) error {
	ok, err := c.checkAction(ctx, checkAction, obj, ctxtk...)
	if err != nil {
		return err
	}
	if !ok {
		return ErrUnauthorized
	}
	return nil
}

func (c *Checker) checkAction(ctx context.Context, checkAction Action, obj EntityKey, ctxtk ...TupleKey) (bool, error) {
	checkUser := authn.ForContext(ctx)
	if checkUser == nil {
		return false, nil
	}
	userName := checkUser.Name()
	if c.checkGlobalAdmin(checkUser) {
		log.Debug().Str("check_user", userName).Str("obj", obj.String()).Str("check_action", checkAction.String()).Msg("global admin action")
		return true, nil
	}
	checkTk := NewTupleKey().WithUser(userName).WithObject(obj.Type, obj.Name).WithAction(checkAction)
	ret, err := c.fgaClient.Check(ctx, checkTk, ctxtk...)
	log.Trace().Str("tk", checkTk.String()).Bool("result", ret).Err(err).Msg("checkAction")
	return ret, err
}

func (c *Checker) ctxIsGlobalAdmin(ctx context.Context) bool {
	checkUser := authn.ForContext(ctx)
	if checkUser == nil {
		return false
	}
	return c.checkGlobalAdmin(checkUser)
}

func (c *Checker) checkGlobalAdmin(checkUser authn.User) bool {
	if c == nil {
		return false
	}
	if checkUser == nil {
		return false
	}
	userName := checkUser.Name()
	for _, v := range c.globalAdmins {
		if v == userName {
			return true
		}
	}
	return false
}

// Helpers

type hasId interface {
	GetId() int64
}

func checkIds[T hasId](ents []T, ids []int64) error {
	if len(ents) != len(ids) {
		return errors.New("not found")
	}
	check := map[int64]bool{}
	for _, ent := range ents {
		check[ent.GetId()] = true
	}
	for _, id := range ids {
		if _, ok := check[id]; !ok {
			return errors.New("not found")
		}
	}
	return nil
}

func getEntities[T hasId](ctx context.Context, db sqlx.Ext, ids []int64, table string, cols ...string) ([]T, error) {
	var t []T
	q := sq.StatementBuilder.Select(cols...).From(table).Where(sq.Eq{"id": ids})
	if err := dbutil.Select(ctx, db, q, &t); err != nil {
		log.Trace().Err(err)
		return nil, err
	}
	if err := checkIds(t, ids); err != nil {
		return nil, err
	}
	return t, nil
}

func first[T any](v []T) T {
	var xt T
	if len(v) > 0 {
		return v[0]
	}
	return xt
}

// todo: rename to dbTestTupleLookup and make arg TestTuple
func dbTupleLookup(t testing.TB, dbx sqlx.Ext, tk TupleKey) TupleKey {
	var err error
	var found bool
	tk.Subject, found, err = dbNameToEntityKey(dbx, tk.Subject)
	if !found && t != nil {
		t.Logf("lookup warning: %s not found", tk.Subject.String())
	}
	if err != nil {
		t.Log(err)
	}
	tk.Object, found, err = dbNameToEntityKey(dbx, tk.Object)
	if !found && t != nil {
		t.Logf("lookup warning: %s not found", tk.Object.String())
	}
	if err != nil {
		t.Log(err)
	}
	return tk
}

func ekLookup(dbx sqlx.Ext, tk TupleKey) (TupleKey, bool, error) {
	var err error
	var found1 bool
	var found2 bool
	tk.Subject, found1, err = dbNameToEntityKey(dbx, tk.Subject)
	if err != nil {
		return tk, false, err
	}
	tk.Object, found2, err = dbNameToEntityKey(dbx, tk.Object)
	if err != nil {
		return tk, false, err
	}
	return tk, found1 && found2, nil
}

func dbNameToEntityKey(dbx sqlx.Ext, ek EntityKey) (EntityKey, bool, error) {
	if ek.Name == "" {
		return ek, false, nil
	}
	nsplit := strings.Split(ek.Name, "#")
	oname := nsplit[0]
	nname := ek.Name
	var err error
	switch ek.Type {
	case TenantType:
		err = sqlx.Get(dbx, &nname, "select id from tl_tenants where tenant_name = $1", oname)
	case GroupType:
		err = sqlx.Get(dbx, &nname, "select id from tl_groups where group_name = $1", oname)
	case FeedType:
		err = sqlx.Get(dbx, &nname, "select id from current_feeds where onestop_id = $1", oname)
	case FeedVersionType:
		err = sqlx.Get(dbx, &nname, "select id from feed_versions where sha1 = $1", oname)
	case UserType:
	}
	found := false
	if err == sql.ErrNoRows {
		err = nil
	} else {
		found = true
	}
	if err != nil {
		return ek, found, err
	}
	nsplit[0] = nname
	ek.Name = strings.Join(nsplit, "#")
	return ek, found, nil
}
