package resolvers

import (
	"context"

	"github.com/interline-io/transitland-server/model"
)

// FEED

type feedResolver struct{ *Resolver }

func (r *feedResolver) FeedState(ctx context.Context, obj *model.Feed) (*model.FeedState, error) {
	return For(ctx).FeedStatesByFeedID.Load(obj.ID)
}

func (r *feedResolver) FeedVersions(ctx context.Context, obj *model.Feed, limit *int, where *model.FeedVersionFilter) ([]*model.FeedVersion, error) {
	return For(ctx).FeedVersionsByFeedID.Load(model.FeedVersionParam{
		FeedID: obj.ID,
		Limit:  limit,
		Where:  where,
	})
}

func (r *feedResolver) License(ctx context.Context, obj *model.Feed) (*model.FeedLicense, error) {
	return &model.FeedLicense{FeedLicense: obj.License}, nil
}

func (r *feedResolver) Languages(ctx context.Context, obj *model.Feed) ([]string, error) {
	return obj.Languages, nil
}

func (r *feedResolver) Urls(ctx context.Context, obj *model.Feed) (*model.FeedUrls, error) {
	return &model.FeedUrls{FeedUrls: obj.URLs}, nil
}

func (r *feedResolver) AssociatedOperators(ctx context.Context, obj *model.Feed) ([]*model.Operator, error) {
	return For(ctx).OperatorsByFeedID.Load(model.OperatorParam{FeedID: obj.ID})
}

func (r *feedResolver) Authorization(ctx context.Context, obj *model.Feed) (*model.FeedAuthorization, error) {
	return &model.FeedAuthorization{FeedAuthorization: obj.Authorization}, nil
}

func (r *feedResolver) FeedFetches(ctx context.Context, obj *model.Feed, limit *int) ([]*model.FeedFetch, error) {
	return For(ctx).FeedFetchesByFeedID.Load(model.FeedFetchParam{FeedID: obj.ID, Limit: limit})
}

// FEED STATE

type feedStateResolver struct{ *Resolver }
