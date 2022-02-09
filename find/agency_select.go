package find

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/interline-io/transitland-server/model"
)

func AgencySelect(limit *int, after *int, ids []int, active bool, where *model.AgencyFilter) sq.SelectBuilder {
	distinct := false
	qView := sq.StatementBuilder.
		Select(
			"gtfs_agencies.*",
			"tl_agency_geometries.geometry",
			"feed_versions.sha1 AS feed_version_sha1",
			"current_feeds.onestop_id AS feed_onestop_id",
			"coalesce (coif.resolved_onestop_id, '') as onestop_id",
			"coif.id as coif_id",
		).
		From("gtfs_agencies").
		Join("feed_versions ON feed_versions.id = gtfs_agencies.feed_version_id").
		Join("current_feeds ON current_feeds.id = feed_versions.feed_id").
		JoinClause("left join tl_agency_geometries ON tl_agency_geometries.agency_id = gtfs_agencies.id").
		JoinClause("left join current_operators_in_feed coif ON coif.feed_id = current_feeds.id AND coif.resolved_gtfs_agency_id = gtfs_agencies.agency_id").
		Where(sq.Eq{"current_feeds.deleted_at": nil}).
		OrderBy("gtfs_agencies.id")

	if where != nil {
		if where.FeedVersionSha1 != nil {
			qView = qView.Where(sq.Eq{"feed_versions.sha1": *where.FeedVersionSha1})
		}
		if where.FeedOnestopID != nil {
			qView = qView.Where(sq.Eq{"current_feeds.onestop_id": *where.FeedOnestopID})
		}
		if where.AgencyID != nil {
			qView = qView.Where(sq.Eq{"gtfs_agencies.agency_id": *where.AgencyID})
		}
		if where.AgencyName != nil {
			qView = qView.Where(sq.Eq{"gtfs_agencies.agency_name": *where.AgencyName})
		}
		if where.OnestopID != nil {
			qView = qView.Where(sq.Eq{"coif.resolved_onestop_id": *where.OnestopID})
		}
		if where.Within != nil && where.Within.Valid {
			qView = qView.Where("ST_Intersects(tl_agency_geometries.geometry, ?)", where.Within)
		}
		if where.Near != nil {
			radius := checkFloat(&where.Near.Radius, 0, 10_000)
			qView = qView.Where("ST_DWithin(tl_agency_geometries.geometry, ST_MakePoint(?,?), ?)", where.Near.Lon, where.Near.Lat, radius)
		}
		if where.Adm0Iso != nil || where.Adm1Iso != nil || where.Adm0Name != nil || where.Adm1Name != nil || where.CityName != nil {
			distinct = true
			qView = qView.JoinClause(`
			JOIN (
				SELECT
					tlap.agency_id,
					tlap.name,
					tlap.adm1name,
					tlap.adm0name,
					ne_admin.iso_3166_2 as adm1iso,
					ne_admin.iso_a2 as adm0iso
				FROM tl_agency_places tlap
				JOIN ne_10m_admin_1_states_provinces ne_admin on ne_admin.name = tlap.adm1name and ne_admin.admin = tlap.adm0name
			) tlap on tlap.agency_id = gtfs_agencies.id`)
			if where.Adm0Iso != nil {
				qView = qView.Where(sq.ILike{"tlap.adm0iso": *where.Adm0Iso})
			}
			if where.Adm1Iso != nil {
				qView = qView.Where(sq.ILike{"tlap.adm1iso": *where.Adm1Iso})
			}
			if where.Adm0Name != nil {
				qView = qView.Where(sq.ILike{"tlap.adm0name": *where.Adm0Name})
			}
			if where.Adm1Name != nil {
				qView = qView.Where(sq.ILike{"tlap.adm1name": *where.Adm1Name})
			}
			if where.CityName != nil {
				qView = qView.Where(sq.ILike{"tlap.name": *where.CityName})
			}
		}
	}
	if distinct {
		qView = qView.Distinct().Options("on (gtfs_agencies.id)")
	}
	if active {
		qView = qView.Join("feed_states on feed_states.feed_version_id = gtfs_agencies.feed_version_id")
	}
	if len(ids) > 0 {
		qView = qView.Where(sq.Eq{"gtfs_agencies.id": ids})
	}
	if after != nil {
		qView = qView.Where(sq.Gt{"gtfs_agencies.id": *after})
	}
	q := sq.StatementBuilder.Select("t.*").FromSelect(qView, "t").Limit(checkLimit(limit))

	if where != nil {
		if where.Search != nil && len(*where.Search) > 1 {
			rank, wc := tsQuery(*where.Search)
			q = q.Column(rank).Where(wc)
		}
	}
	return q
}

func AgencyPlaceSelect(limit *int, after *int, ids []int, where *model.AgencyPlaceFilter) sq.SelectBuilder {
	q := quickSelectOrder("tl_agency_places", limit, after, ids, "rank desc")
	if where != nil {
		if where.MinRank != nil {
			q = q.Where(sq.GtOrEq{"rank": where.MinRank})
		}
	}
	return q
}
