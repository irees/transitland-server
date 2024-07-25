package gql

import (
	"context"
	"sort"
	"strconv"
	"time"

	"github.com/interline-io/transitland-lib/tlxy"
	"github.com/interline-io/transitland-server/internal/directions"
	"github.com/interline-io/transitland-server/model"
)

// STOP

type stopResolver struct {
	*Resolver
}

func (r *stopResolver) Cursor(ctx context.Context, obj *model.Stop) (*model.Cursor, error) {
	c := model.NewCursor(obj.FeedVersionID, obj.ID)
	return &c, nil
}

func (r *stopResolver) FeedVersion(ctx context.Context, obj *model.Stop) (*model.FeedVersion, error) {
	return For(ctx).FeedVersionsByID.Load(ctx, obj.FeedVersionID)()
}

func (r *stopResolver) Level(ctx context.Context, obj *model.Stop) (*model.Level, error) {
	if !obj.LevelID.Valid {
		return nil, nil
	}
	return For(ctx).LevelsByID.Load(ctx, obj.LevelID.Int())()
}

func (r *stopResolver) ChildLevels(ctx context.Context, obj *model.Stop, limit *int) ([]*model.Level, error) {
	return For(ctx).LevelsByParentStationID.Load(ctx, model.LevelParam{ParentStationID: obj.ID, Limit: limit})()
}

func (r *stopResolver) Parent(ctx context.Context, obj *model.Stop) (*model.Stop, error) {
	if !obj.ParentStation.Valid {
		return nil, nil
	}
	return For(ctx).StopsByID.Load(ctx, obj.ParentStation.Int())()
}

func (r *stopResolver) Children(ctx context.Context, obj *model.Stop, limit *int) ([]*model.Stop, error) {
	return For(ctx).StopsByParentStopID.Load(ctx, model.StopParam{ParentStopID: obj.ID, Limit: checkLimit(limit)})()
}

func (r *stopResolver) Place(ctx context.Context, obj *model.Stop) (*model.StopPlace, error) {
	pt := tlxy.Point{Lon: obj.Geometry.X(), Lat: obj.Geometry.Y()}
	return For(ctx).StopPlacesByStopID.Load(ctx, model.StopPlaceParam{ID: obj.ID, Point: pt})()
}

func (r *stopResolver) RouteStops(ctx context.Context, obj *model.Stop, limit *int) ([]*model.RouteStop, error) {
	return For(ctx).RouteStopsByStopID.Load(ctx, model.RouteStopParam{StopID: obj.ID, Limit: checkLimit(limit)})()
}

func (r *stopResolver) PathwaysFromStop(ctx context.Context, obj *model.Stop, limit *int) ([]*model.Pathway, error) {
	return For(ctx).PathwaysByFromStopID.Load(ctx, model.PathwayParam{FromStopID: obj.ID, Limit: checkLimit(limit)})()
}

func (r *stopResolver) PathwaysToStop(ctx context.Context, obj *model.Stop, limit *int) ([]*model.Pathway, error) {
	return For(ctx).PathwaysByToStopID.Load(ctx, model.PathwayParam{ToStopID: obj.ID, Limit: checkLimit(limit)})()
}

func (r *stopResolver) ExternalReference(ctx context.Context, obj *model.Stop) (*model.StopExternalReference, error) {
	return For(ctx).StopExternalReferencesByStopID.Load(ctx, obj.ID)()
}

func (r *stopResolver) Observations(ctx context.Context, obj *model.Stop, limit *int, where *model.StopObservationFilter) ([]*model.StopObservation, error) {
	return For(ctx).StopObservationsByStopID.Load(ctx, model.StopObservationParam{StopID: obj.ID, Where: where, Limit: checkLimit(limit)})()
}

func (r *stopResolver) Departures(ctx context.Context, obj *model.Stop, limit *int, where *model.StopTimeFilter) ([]*model.StopTime, error) {
	if where == nil {
		where = &model.StopTimeFilter{}
	}
	t := true
	where.ExcludeLast = &t
	return r.getStopTimes(ctx, obj, limit, where)
}

func (r *stopResolver) Arrivals(ctx context.Context, obj *model.Stop, limit *int, where *model.StopTimeFilter) ([]*model.StopTime, error) {
	if where == nil {
		where = &model.StopTimeFilter{}
	}
	t := true
	where.ExcludeFirst = &t
	return r.getStopTimes(ctx, obj, limit, where)
}

func (r *stopResolver) StopTimes(ctx context.Context, obj *model.Stop, limit *int, where *model.StopTimeFilter) ([]*model.StopTime, error) {
	return r.getStopTimes(ctx, obj, limit, where)
}

func (r *stopResolver) getStopTimes(ctx context.Context, obj *model.Stop, limit *int, where *model.StopTimeFilter) ([]*model.StopTime, error) {
	sts, err := (For(ctx).StopTimesByStopID.Load(ctx, model.StopTimeParam{
		StopID:        obj.ID,
		FeedVersionID: obj.FeedVersionID,
		Limit:         checkLimit(limit),
		Where:         where,
	})())
	if err != nil {
		return nil, err
	}

	// Merge scheduled stop times with rt stop times
	// TODO: handle StopTimeFilter in RT
	// Handle scheduled trips; these can be matched on trip_id or (route_id,direction_id,...)
	for _, st := range sts {
		ft := model.Trip{}
		ft.FeedVersionID = obj.FeedVersionID
		ft.TripID, _ = model.ForContext(ctx).RTFinder.GetGtfsTripID(atoi(st.TripID)) // TODO!
		if ste, ok := model.ForContext(ctx).RTFinder.FindStopTimeUpdate(&ft, st); ok {
			st.RTStopTimeUpdate = ste
		}
	}
	// Handle added trips; these must specify stop_id in StopTimeUpdates
	for _, rtTrip := range model.ForContext(ctx).RTFinder.GetAddedTripsForStop(obj) {
		for _, stu := range rtTrip.StopTimeUpdate {
			if stu.GetStopId() != obj.StopID {
				continue
			}
			// create a new StopTime
			rtst := &model.StopTime{}
			rtst.RTTripID = rtTrip.Trip.GetTripId()
			rtst.RTStopTimeUpdate = stu
			rtst.FeedVersionID = obj.FeedVersionID
			rtst.TripID = "0"
			rtst.StopID = strconv.Itoa(obj.ID)
			rtst.StopSequence = int(stu.GetStopSequence())
			sts = append(sts, rtst)
		}
	}
	// Sort by scheduled departure time.
	// TODO: Sort by rt departure time? Requires full StopTime Resolver for timezones, processing, etc.
	sort.Slice(sts, func(i, j int) bool {
		sta := sts[i]
		stb := sts[j]
		a := int(sta.ServiceDate.Val.Unix()) + sta.DepartureTime.Seconds
		b := int(stb.ServiceDate.Val.Unix()) + stb.DepartureTime.Seconds
		return a < b
	})
	return sts, nil
}

func (r *stopResolver) Alerts(ctx context.Context, obj *model.Stop, active *bool, limit *int) ([]*model.Alert, error) {
	rtAlerts := model.ForContext(ctx).RTFinder.FindAlertsForStop(obj, checkLimit(limit), active)
	return rtAlerts, nil
}

func (r *stopResolver) Directions(ctx context.Context, obj *model.Stop, from *model.WaypointInput, to *model.WaypointInput, mode *model.StepMode, departAt *time.Time) (*model.Directions, error) {
	oc := obj.Coordinates()
	swp := &model.WaypointInput{
		Lon:  oc[0],
		Lat:  oc[1],
		Name: &obj.StopName,
	}
	p := model.DirectionRequest{}
	if from != nil {
		p.From = from
		p.To = swp
	} else if to != nil {
		p.From = swp
		p.To = to
	}
	if mode != nil {
		p.Mode = *mode
	}
	return directions.HandleRequest("", p)
}

func (r *stopResolver) NearbyStops(ctx context.Context, obj *model.Stop, limit *int, radius *float64) ([]*model.Stop, error) {
	c := obj.Coordinates()
	nearbyStops, err := model.ForContext(ctx).Finder.FindStops(ctx, checkLimit(limit), nil, nil, &model.StopFilter{Near: &model.PointRadius{Lon: c[0], Lat: c[1], Radius: checkFloat(radius, 0, MAX_RADIUS)}})
	return nearbyStops, err
}

//////////

type stopExternalReferenceResolver struct {
	*Resolver
}

func (r *stopExternalReferenceResolver) TargetActiveStop(ctx context.Context, obj *model.StopExternalReference) (*model.Stop, error) {
	return For(ctx).TargetStopsByStopID.Load(ctx, obj.ID)()
}
