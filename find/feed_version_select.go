package find

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/interline-io/transitland-server/model"
)

func FeedVersionSelect(limit *int, after *model.Cursor, ids []int, where *model.FeedVersionFilter) sq.SelectBuilder {
	q := sq.StatementBuilder.
		Select("t.*, tl_feed_version_geometries.geometry").
		From("feed_versions t").
		Join("current_feeds cf on cf.id = t.feed_id").Where(sq.Eq{"cf.deleted_at": nil}).
		JoinClause("left join tl_feed_version_geometries on tl_feed_version_geometries.feed_version_id = t.id").
		Limit(checkLimit(limit)).
		OrderBy("t.fetched_at desc, t.id desc")
	if len(ids) > 0 {
		q = q.Where(sq.Eq{"t.id": ids})
	}
	if after != nil && after.Valid && after.ID > 0 {
		q = q.Where(sq.Expr("(t.fetched_at,t.id) < (select fetched_at from feed_versions where id = ?", after.ID))
	}
	if where != nil {
		if where.Sha1 != nil {
			q = q.Where(sq.Eq{"sha1": *where.Sha1})
		}
		if len(where.FeedIds) > 0 {
			q = q.Where(sq.Eq{"feed_id": where.FeedIds})
		}
		if where.FeedOnestopID != nil {
			q = q.Where(sq.Eq{"cf.onestop_id": *where.FeedOnestopID})
		}
	}
	return q
}

func FeedVersionServiceLevelSelect(limit *int, after *model.Cursor, ids []int, where *model.FeedVersionServiceLevelFilter) sq.SelectBuilder {
	q := quickSelectOrder("feed_version_service_levels", limit, after, nil, "")
	if where == nil {
		where = &model.FeedVersionServiceLevelFilter{}
	}
	q = q.Where(sq.Eq{"route_id": nil})
	if where.StartDate != nil {
		q = q.Where(sq.GtOrEq{"start_date": where.StartDate})
	}
	if where.EndDate != nil {
		q = q.Where(sq.LtOrEq{"end_date": where.EndDate})
	}
	return q
}
