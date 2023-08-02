package azcheck

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/interline-io/transitland-lib/log"
	"github.com/interline-io/transitland-server/auth/authz"
	"github.com/interline-io/transitland-server/internal/util"
	"github.com/interline-io/transitland-server/model"
)

func NewServer(checker model.Checker) (http.Handler, error) {
	router := chi.NewRouter()

	/////////////////
	// USERS
	/////////////////

	router.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		ret, err := checker.UserList(r.Context(), &authz.UserListRequest{Q: r.URL.Query().Get("q")})
		handleJson(w, ret, err)
	})
	router.Get("/users/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		ret, err := checker.User(r.Context(), &authz.UserRequest{Id: chi.URLParam(r, "user_id")})
		handleJson(w, ret, err)
	})
	router.Get("/me", func(w http.ResponseWriter, r *http.Request) {
		ret, err := checker.Me(r.Context(), &authz.MeRequest{})
		handleJson(w, ret, err)
	})

	/////////////////
	// TENANTS
	/////////////////

	router.Get("/tenants", func(w http.ResponseWriter, r *http.Request) {
		ret, err := checker.TenantList(r.Context(), &authz.TenantListRequest{})
		handleJson(w, ret, err)
	})
	router.Get("/tenants/{tenant_id}", func(w http.ResponseWriter, r *http.Request) {
		ret, err := checker.TenantPermissions(r.Context(), &authz.TenantRequest{Id: checkId(r, "tenant_id")})
		handleJson(w, ret, err)
	})
	router.Post("/tenants/{tenant_id}", func(w http.ResponseWriter, r *http.Request) {
		check := authz.Tenant{}
		if err := parseJson(r.Body, &check); err != nil {
			handleJson(w, nil, err)
			return
		}
		check.Id = checkId(r, "tenant_id")
		_, err := checker.TenantSave(r.Context(), &authz.TenantSaveRequest{Tenant: &check})
		handleJson(w, nil, err)
	})
	router.Post("/tenants/{tenant_id}/groups", func(w http.ResponseWriter, r *http.Request) {
		check := authz.Group{}
		if err := parseJson(r.Body, &check); err != nil {
			handleJson(w, nil, err)
			return
		}
		_, err := checker.TenantCreateGroup(r.Context(), &authz.TenantCreateGroupRequest{Id: checkId(r, "tenant_id"), Group: &check})
		handleJson(w, nil, err)
	})
	router.Post("/tenants/{tenant_id}/permissions", func(w http.ResponseWriter, r *http.Request) {
		check := &authz.EntityRelation{}
		if err := parseJson(r.Body, check); err != nil {
			handleJson(w, nil, err)
		}
		entId := checkId(r, "tenant_id")
		_, err := checker.TenantAddPermission(r.Context(), &authz.TenantModifyPermissionRequest{Id: entId, EntityRelation: check})
		handleJson(w, nil, err)
	})
	router.Delete("/tenants/{tenant_id}/permissions", func(w http.ResponseWriter, r *http.Request) {
		check := &authz.EntityRelation{}
		if err := parseJson(r.Body, check); err != nil {
			handleJson(w, nil, err)
		}
		entId := checkId(r, "tenant_id")
		_, err := checker.TenantRemovePermission(r.Context(), &authz.TenantModifyPermissionRequest{Id: entId, EntityRelation: check})
		handleJson(w, nil, err)
	})

	/////////////////
	// GROUPS
	/////////////////

	router.Get("/groups", func(w http.ResponseWriter, r *http.Request) {
		ret, err := checker.GroupList(r.Context(), &authz.GroupListRequest{})
		handleJson(w, ret, err)
	})
	router.Post("/groups/{group_id}", func(w http.ResponseWriter, r *http.Request) {
		check := authz.Group{}
		if err := parseJson(r.Body, &check); err != nil {
			handleJson(w, nil, err)
			return
		}
		check.Id = checkId(r, "group_id")
		_, err := checker.GroupSave(r.Context(), &authz.GroupSaveRequest{Group: &check})
		handleJson(w, nil, err)
	})
	router.Get("/groups/{group_id}", func(w http.ResponseWriter, r *http.Request) {
		ret, err := checker.GroupPermissions(r.Context(), &authz.GroupRequest{Id: checkId(r, "group_id")})
		handleJson(w, ret, err)
	})
	router.Post("/groups/{group_id}/permissions", func(w http.ResponseWriter, r *http.Request) {
		check := &authz.EntityRelation{}
		if err := parseJson(r.Body, check); err != nil {
			handleJson(w, nil, err)
		}
		entId := checkId(r, "group_id")
		_, err := checker.GroupAddPermission(r.Context(), &authz.GroupModifyPermissionRequest{Id: entId, EntityRelation: check})
		handleJson(w, nil, err)
	})
	router.Delete("/groups/{group_id}/permissions", func(w http.ResponseWriter, r *http.Request) {
		check := &authz.EntityRelation{}
		if err := parseJson(r.Body, check); err != nil {
			handleJson(w, nil, err)
		}
		entId := checkId(r, "group_id")
		_, err := checker.GroupRemovePermission(r.Context(), &authz.GroupModifyPermissionRequest{Id: entId, EntityRelation: check})
		handleJson(w, nil, err)
	})
	router.Post("/groups/{group_id}/tenant", func(w http.ResponseWriter, r *http.Request) {
		check := authz.GroupSetTenantRequest{}
		if err := parseJson(r.Body, &check); err != nil {
			handleJson(w, nil, err)
			return
		}
		check.Id = checkId(r, "group_id")
		_, err := checker.GroupSetTenant(r.Context(), &check)
		handleJson(w, nil, err)
	})

	/////////////////
	// FEEDS
	/////////////////

	router.Get("/feeds", func(w http.ResponseWriter, r *http.Request) {
		ret, err := checker.FeedList(r.Context(), &authz.FeedListRequest{})
		handleJson(w, ret, err)
	})
	router.Get("/feeds/{feed_id}", func(w http.ResponseWriter, r *http.Request) {
		ret, err := checker.FeedPermissions(r.Context(), &authz.FeedRequest{Id: checkId(r, "feed_id")})
		handleJson(w, ret, err)
	})
	router.Post("/feeds/{feed_id}/group", func(w http.ResponseWriter, r *http.Request) {
		check := authz.FeedSetGroupRequest{}
		if err := parseJson(r.Body, &check); err != nil {
			handleJson(w, nil, err)
			return
		}
		check.Id = checkId(r, "feed_id")
		_, err := checker.FeedSetGroup(r.Context(), &check)
		handleJson(w, nil, err)
	})

	/////////////////
	// FEED VERSIONS
	/////////////////

	router.Get("/feed_versions", func(w http.ResponseWriter, r *http.Request) {
		ret, err := checker.FeedVersionList(r.Context(), &authz.FeedVersionListRequest{})
		handleJson(w, ret, err)
	})
	router.Get("/feed_versions/{feed_version_id}", func(w http.ResponseWriter, r *http.Request) {
		ret, err := checker.FeedVersionPermissions(r.Context(), &authz.FeedVersionRequest{Id: checkId(r, "feed_version_id")})
		handleJson(w, ret, err)
	})
	router.Post("/feed_versions/{feed_version_id}/permissions", func(w http.ResponseWriter, r *http.Request) {
		check := &authz.EntityRelation{}
		if err := parseJson(r.Body, check); err != nil {
			handleJson(w, nil, err)
		}
		entId := checkId(r, "feed_version_id")
		_, err := checker.FeedVersionAddPermission(r.Context(), &authz.FeedVersionModifyPermissionRequest{Id: entId, EntityRelation: check})
		handleJson(w, nil, err)
	})
	router.Delete("/feed_versions/{feed_version_id}/permissions", func(w http.ResponseWriter, r *http.Request) {
		check := &authz.EntityRelation{}
		if err := parseJson(r.Body, check); err != nil {
			handleJson(w, nil, err)
		}
		entId := checkId(r, "feed_version_id")
		_, err := checker.FeedVersionRemovePermission(r.Context(), &authz.FeedVersionModifyPermissionRequest{Id: entId, EntityRelation: check})
		handleJson(w, nil, err)
	})

	return router, nil
}

func handleJson(w http.ResponseWriter, ret any, err error) {
	if err == ErrUnauthorized {
		log.Error().Err(err).Msg("unauthorized")
		http.Error(w, util.MakeJsonError(http.StatusText(http.StatusUnauthorized)), http.StatusUnauthorized)
		return
	} else if err != nil {
		log.Error().Err(err).Msg("admin api error")
		http.Error(w, util.MakeJsonError(http.StatusText(http.StatusInternalServerError)), http.StatusInternalServerError)
		return
	}
	if ret == nil {
		ret = map[string]bool{"success": true}
	}
	jj, _ := json.Marshal(ret)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jj)
}

func checkId(r *http.Request, key string) int64 {
	v, _ := strconv.Atoi(chi.URLParam(r, key))
	return int64(v)
}

func parseJson(r io.Reader, v any) error {
	data, err := ioutil.ReadAll(io.LimitReader(r, 1_000_000))
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
