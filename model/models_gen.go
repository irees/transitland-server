// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/interline-io/transitland-lib/tl/tt"
)

type AgencyFilter struct {
	OnestopID       *string `json:"onestop_id,omitempty"`
	FeedVersionSha1 *string `json:"feed_version_sha1,omitempty"`
	FeedOnestopID   *string `json:"feed_onestop_id,omitempty"`
	AgencyID        *string `json:"agency_id,omitempty"`
	// Search for records with this GTFS agency_name
	AgencyName *string      `json:"agency_name,omitempty"`
	Bbox       *BoundingBox `json:"bbox,omitempty"`
	Within     *tt.Polygon  `json:"within,omitempty"`
	// Search for agencies within a radius
	Near *PointRadius `json:"near,omitempty"`
	// Full text search
	Search *string `json:"search,omitempty"`
	// Search by city name (provided by Natural Earth)
	CityName *string `json:"city_name,omitempty"`
	// Search by country name (provided by Natural Earth)
	Adm0Name *string `json:"adm0_name,omitempty"`
	// Search by state/province/division name (provided by Natural Earth)
	Adm1Name *string `json:"adm1_name,omitempty"`
	// Search by country 2 letter ISO 3166 code (provided by Natural Earth)
	Adm0Iso *string `json:"adm0_iso,omitempty"`
	// Search by state/province/division ISO 3166-2 code (provided by Natural Earth)
	Adm1Iso *string        `json:"adm1_iso,omitempty"`
	License *LicenseFilter `json:"license,omitempty"`
}

type AgencyPlace struct {
	CityName *string  `json:"city_name,omitempty"`
	Adm0Name *string  `json:"adm0_name,omitempty"`
	Adm1Name *string  `json:"adm1_name,omitempty"`
	Adm0Iso  *string  `json:"adm0_iso,omitempty"`
	Adm1Iso  *string  `json:"adm1_iso,omitempty"`
	Rank     *float64 `json:"rank,omitempty"`
	AgencyID int      `json:"-"`
}

type AgencyPlaceFilter struct {
	MinRank *float64 `json:"min_rank,omitempty"`
}

// [Alert](https://gtfs.org/reference/realtime/v2/#message-alert) message, also called a service alert, provided by a source GTFS Realtime feed.
type Alert struct {
	ActivePeriod       []*RTTimeRange   `json:"active_period,omitempty"`
	Cause              *string          `json:"cause,omitempty"`
	Effect             *string          `json:"effect,omitempty"`
	HeaderText         []*RTTranslation `json:"header_text"`
	DescriptionText    []*RTTranslation `json:"description_text"`
	TtsHeaderText      []*RTTranslation `json:"tts_header_text,omitempty"`
	TtsDescriptionText []*RTTranslation `json:"tts_description_text,omitempty"`
	URL                []*RTTranslation `json:"url,omitempty"`
	SeverityLevel      *string          `json:"severity_level,omitempty"`
}

type BoundingBox struct {
	MinLon float64 `json:"min_lon"`
	MinLat float64 `json:"min_lat"`
	MaxLon float64 `json:"max_lon"`
	MaxLat float64 `json:"max_lat"`
}

type CalendarDateFilter struct {
	Date          *tt.Date `json:"date,omitempty"`
	ExceptionType *int     `json:"exception_type,omitempty"`
}

type CensusGeography struct {
	ID            int            `json:"id"`
	LayerName     string         `json:"layer_name"`
	Geoid         *string        `json:"geoid,omitempty"`
	Name          *string        `json:"name,omitempty"`
	Aland         *float64       `json:"aland,omitempty"`
	Awater        *float64       `json:"awater,omitempty"`
	Geometry      *tt.Polygon    `json:"geometry,omitempty"`
	Values        []*CensusValue `json:"values"`
	MatchEntityID int            `json:"-"`
}

type CensusTable struct {
	ID         int    `json:"id"`
	TableName  string `json:"table_name"`
	TableTitle string `json:"table_title"`
	TableGroup string `json:"table_group"`
}

type CensusValue struct {
	Table       *CensusTable `json:"table"`
	Values      tt.Map       `json:"values"`
	GeographyID int          `json:"-"`
	TableID     int          `json:"-"`
}

type DirectionRequest struct {
	To       *WaypointInput `json:"to"`
	From     *WaypointInput `json:"from"`
	Mode     StepMode       `json:"mode"`
	DepartAt *time.Time     `json:"depart_at,omitempty"`
}

type Directions struct {
	Success     bool         `json:"success"`
	Exception   *string      `json:"exception,omitempty"`
	DataSource  *string      `json:"data_source,omitempty"`
	Origin      *Waypoint    `json:"origin,omitempty"`
	Destination *Waypoint    `json:"destination,omitempty"`
	Duration    *Duration    `json:"duration,omitempty"`
	Distance    *Distance    `json:"distance,omitempty"`
	StartTime   *time.Time   `json:"start_time,omitempty"`
	EndTime     *time.Time   `json:"end_time,omitempty"`
	Itineraries []*Itinerary `json:"itineraries,omitempty"`
}

type Distance struct {
	Distance float64      `json:"distance"`
	Units    DistanceUnit `json:"units"`
}

type Duration struct {
	Duration float64      `json:"duration"`
	Units    DurationUnit `json:"units"`
}

type EntityDeleteResult struct {
	ID int `json:"id"`
}

type FeedFetchFilter struct {
	Success *bool `json:"success,omitempty"`
}

type FeedFilter struct {
	// Search for feed with a specific Onestop ID
	OnestopID *string `json:"onestop_id,omitempty"`
	// Search for feeds of certain data types
	Spec []FeedSpecTypes `json:"spec,omitempty"`
	// Search for feeds with or without a fetch error
	FetchError *bool `json:"fetch_error,omitempty"`
	// Search for feeds by their import status
	ImportStatus *ImportStatus `json:"import_status,omitempty"`
	// Full text search
	Search *string `json:"search,omitempty"`
	// Search for feeds with a tag
	Tags *tt.Tags `json:"tags,omitempty"`
	// Search for feeds by their source URLs
	SourceURL *FeedSourceURL `json:"source_url,omitempty"`
	License   *LicenseFilter `json:"license,omitempty"`
	Bbox      *BoundingBox   `json:"bbox,omitempty"`
	Within    *tt.Polygon    `json:"within,omitempty"`
	Near      *PointRadius   `json:"near,omitempty"`
}

type FeedSourceURL struct {
	URL           *string             `json:"url,omitempty"`
	Type          *FeedSourceURLTypes `json:"type,omitempty"`
	CaseSensitive *bool               `json:"case_sensitive,omitempty"`
}

type FeedVersionDeleteResult struct {
	Success bool `json:"success"`
}

type FeedVersionFetchResult struct {
	FeedVersion  *FeedVersion `json:"feed_version,omitempty"`
	FetchError   *string      `json:"fetch_error,omitempty"`
	FoundSha1    bool         `json:"found_sha1"`
	FoundDirSha1 bool         `json:"found_dir_sha1"`
}

type FeedVersionFilter struct {
	ImportStatus  *ImportStatus        `json:"import_status,omitempty"`
	FeedOnestopID *string              `json:"feed_onestop_id,omitempty"`
	Sha1          *string              `json:"sha1,omitempty"`
	File          *string              `json:"file,omitempty"`
	FeedIds       []int                `json:"feed_ids,omitempty"`
	Covers        *ServiceCoversFilter `json:"covers,omitempty"`
	Bbox          *BoundingBox         `json:"bbox,omitempty"`
	Within        *tt.Polygon          `json:"within,omitempty"`
	Near          *PointRadius         `json:"near,omitempty"`
}

type FeedVersionImportResult struct {
	Success bool `json:"success"`
}

type FeedVersionInput struct {
	ID *int `json:"id,omitempty"`
}

type FeedVersionServiceLevelFilter struct {
	StartDate *tt.Date `json:"start_date,omitempty"`
	EndDate   *tt.Date `json:"end_date,omitempty"`
}

type FeedVersionSetInput struct {
	ID          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type FeedVersionUnimportResult struct {
	Success bool `json:"success"`
}

type GbfsBikeRequest struct {
	Near *PointRadius `json:"near,omitempty"`
}

type GbfsDockRequest struct {
	Near *PointRadius `json:"near,omitempty"`
}

type Itinerary struct {
	Duration  *Duration `json:"duration"`
	Distance  *Distance `json:"distance"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	From      *Waypoint `json:"from"`
	To        *Waypoint `json:"to"`
	Legs      []*Leg    `json:"legs,omitempty"`
}

type Leg struct {
	Duration  *Duration     `json:"duration"`
	Distance  *Distance     `json:"distance"`
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
	From      *Waypoint     `json:"from,omitempty"`
	To        *Waypoint     `json:"to,omitempty"`
	Steps     []*Step       `json:"steps,omitempty"`
	Geometry  tt.LineString `json:"geometry"`
}

type LevelSetInput struct {
	ID          *int              `json:"id,omitempty"`
	FeedVersion *FeedVersionInput `json:"feed_version,omitempty"`
	LevelID     *string           `json:"level_id,omitempty"`
	LevelName   *string           `json:"level_name,omitempty"`
	LevelIndex  *float64          `json:"level_index,omitempty"`
	Geometry    *tt.Polygon       `json:"geometry,omitempty"`
	Parent      *StopSetInput     `json:"parent,omitempty"`
}

type LicenseFilter struct {
	ShareAlikeOptional    *LicenseValue `json:"share_alike_optional,omitempty"`
	CreateDerivedProduct  *LicenseValue `json:"create_derived_product,omitempty"`
	CommercialUseAllowed  *LicenseValue `json:"commercial_use_allowed,omitempty"`
	UseWithoutAttribution *LicenseValue `json:"use_without_attribution,omitempty"`
	RedistributionAllowed *LicenseValue `json:"redistribution_allowed,omitempty"`
}

type Me struct {
	ID           string   `json:"id"`
	Name         *string  `json:"name,omitempty"`
	Email        *string  `json:"email,omitempty"`
	Roles        []string `json:"roles,omitempty"`
	ExternalData tt.Map   `json:"external_data"`
}

type Mutation struct {
}

type OperatorFilter struct {
	Merged        *bool          `json:"merged,omitempty"`
	OnestopID     *string        `json:"onestop_id,omitempty"`
	FeedOnestopID *string        `json:"feed_onestop_id,omitempty"`
	AgencyID      *string        `json:"agency_id,omitempty"`
	Search        *string        `json:"search,omitempty"`
	Tags          *tt.Tags       `json:"tags,omitempty"`
	CityName      *string        `json:"city_name,omitempty"`
	Adm0Name      *string        `json:"adm0_name,omitempty"`
	Adm1Name      *string        `json:"adm1_name,omitempty"`
	Adm0Iso       *string        `json:"adm0_iso,omitempty"`
	Adm1Iso       *string        `json:"adm1_iso,omitempty"`
	License       *LicenseFilter `json:"license,omitempty"`
	Bbox          *BoundingBox   `json:"bbox,omitempty"`
	Within        *tt.Polygon    `json:"within,omitempty"`
	Near          *PointRadius   `json:"near,omitempty"`
}

type PathwayFilter struct {
	PathwayMode *int `json:"pathway_mode,omitempty"`
}

type PathwaySetInput struct {
	ID                  *int              `json:"id,omitempty"`
	FeedVersion         *FeedVersionInput `json:"feed_version,omitempty"`
	PathwayID           *string           `json:"pathway_id,omitempty"`
	PathwayMode         *int              `json:"pathway_mode,omitempty"`
	IsBidirectional     *int              `json:"is_bidirectional,omitempty"`
	Length              *float64          `json:"length,omitempty"`
	TraversalTime       *int              `json:"traversal_time,omitempty"`
	StairCount          *int              `json:"stair_count,omitempty"`
	MaxSlope            *float64          `json:"max_slope,omitempty"`
	MinWidth            *float64          `json:"min_width,omitempty"`
	SignpostedAs        *string           `json:"signposted_as,omitempty"`
	ReverseSignpostedAs *string           `json:"reverse_signposted_as,omitempty"`
	FromStop            *StopSetInput     `json:"from_stop,omitempty"`
	ToStop              *StopSetInput     `json:"to_stop,omitempty"`
}

type Place struct {
	Adm0Name  *string     `json:"adm0_name,omitempty"`
	Adm1Name  *string     `json:"adm1_name,omitempty"`
	CityName  *string     `json:"city_name,omitempty"`
	Count     int         `json:"count"`
	Operators []*Operator `json:"operators,omitempty"`
	AgencyIDs tt.Ints     `db:"agency_ids"`
}

type PlaceFilter struct {
	MinRank  *float64 `json:"min_rank,omitempty"`
	Adm0Name *string  `json:"adm0_name,omitempty"`
	Adm1Name *string  `json:"adm1_name,omitempty"`
	CityName *string  `json:"city_name,omitempty"`
}

type PointRadius struct {
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Radius float64 `json:"radius"`
}

type Query struct {
}

// See https://gtfs.org/reference/realtime/v2/#message-timerange
type RTTimeRange struct {
	Start *int `json:"start,omitempty"`
	End   *int `json:"end,omitempty"`
}

// See https://gtfs.org/reference/realtime/v2/#message-translatedstring
type RTTranslation struct {
	Text     string  `json:"text"`
	Language *string `json:"language,omitempty"`
}

// See https://gtfs.org/reference/realtime/v2/#message-tripdescriptor
type RTTripDescriptor struct {
	TripID               *string      `json:"trip_id,omitempty"`
	RouteID              *string      `json:"route_id,omitempty"`
	DirectionID          *int         `json:"direction_id,omitempty"`
	StartTime            *tt.WideTime `json:"start_time,omitempty"`
	StartDate            *tt.Date     `json:"start_date,omitempty"`
	ScheduleRelationship *string      `json:"schedule_relationship,omitempty"`
}

// See https://gtfs.org/reference/realtime/v2/#message-vehicledescriptor
type RTVehicleDescriptor struct {
	ID           *string `json:"id,omitempty"`
	Label        *string `json:"label,omitempty"`
	LicensePlate *string `json:"license_plate,omitempty"`
}

// MTC GTFS+ Extension: route_attributes.txt
type RouteAttribute struct {
	Category    *int `json:"category,omitempty"`
	Subcategory *int `json:"subcategory,omitempty"`
	RunningWay  *int `json:"running_way,omitempty"`
	RouteID     int  `json:"-"`
}

type RouteFilter struct {
	OnestopID               *string        `json:"onestop_id,omitempty"`
	OnestopIds              []string       `json:"onestop_ids,omitempty"`
	AllowPreviousOnestopIds *bool          `json:"allow_previous_onestop_ids,omitempty"`
	FeedVersionSha1         *string        `json:"feed_version_sha1,omitempty"`
	FeedOnestopID           *string        `json:"feed_onestop_id,omitempty"`
	RouteID                 *string        `json:"route_id,omitempty"`
	RouteType               *int           `json:"route_type,omitempty"`
	Serviced                *bool          `json:"serviced,omitempty"`
	Bbox                    *BoundingBox   `json:"bbox,omitempty"`
	Within                  *tt.Polygon    `json:"within,omitempty"`
	Near                    *PointRadius   `json:"near,omitempty"`
	Search                  *string        `json:"search,omitempty"`
	OperatorOnestopID       *string        `json:"operator_onestop_id,omitempty"`
	License                 *LicenseFilter `json:"license,omitempty"`
	AgencyIds               []int          `json:"agency_ids,omitempty"`
}

type RouteGeometry struct {
	// If true, the source GTFS feed provides no shapes. This route geometry is based on straight lines between stop points.
	Generated             bool           `json:"generated"`
	Geometry              *tt.LineString `json:"geometry,omitempty"`
	CombinedGeometry      *tt.Geometry   `json:"combined_geometry,omitempty"`
	Length                *float64       `json:"length,omitempty"`
	MaxSegmentLength      *float64       `json:"max_segment_length,omitempty"`
	FirstPointMaxDistance *float64       `json:"first_point_max_distance,omitempty"`
	RouteID               int            `json:"-"`
}

type RouteHeadway struct {
	Stop             *Stop          `json:"stop"`
	DowCategory      *int           `json:"dow_category,omitempty"`
	DirectionID      *int           `json:"direction_id,omitempty"`
	HeadwaySecs      *int           `json:"headway_secs,omitempty"`
	ServiceDate      *tt.Date       `json:"service_date,omitempty"`
	StopTripCount    *int           `json:"stop_trip_count,omitempty"`
	DeparturesUnused []*tt.WideTime `json:"departures,omitempty"`
	DepartureInts    tt.Ints        `db:"departures"`
	RouteID          int            `json:"-"`
	SelectedStopID   int            `json:"-"`
}

type RouteStop struct {
	ID       int     `json:"id"`
	StopID   int     `json:"stop_id"`
	RouteID  int     `json:"route_id"`
	AgencyID int     `json:"agency_id"`
	Route    *Route  `json:"route"`
	Stop     *Stop   `json:"stop"`
	Agency   *Agency `json:"agency"`
}

type RouteStopBuffer struct {
	StopPoints     *tt.Geometry `json:"stop_points,omitempty"`
	StopBuffer     *tt.Geometry `json:"stop_buffer,omitempty"`
	StopConvexhull *tt.Polygon  `json:"stop_convexhull,omitempty"`
}

type RouteStopPattern struct {
	StopPatternID int     `json:"stop_pattern_id"`
	DirectionID   int     `json:"direction_id"`
	Count         int     `json:"count"`
	Trips         []*Trip `json:"trips,omitempty"`
	RouteID       int     `json:"-"`
}

type Segment struct {
	ID              int               `json:"id"`
	WayID           int               `json:"way_id"`
	Geometry        tt.LineString     `json:"geometry"`
	SegmentPatterns []*SegmentPattern `json:"segment_patterns,omitempty"`
}

type SegmentPattern struct {
	ID            int      `json:"id"`
	StopPatternID int      `json:"stop_pattern_id"`
	Segment       *Segment `json:"segment"`
	RouteID       int      `json:"-"`
	SegmentID     int      `json:"-"`
}

type ServiceCoversFilter struct {
	FetchedAfter  *time.Time `json:"fetched_after,omitempty"`
	FetchedBefore *time.Time `json:"fetched_before,omitempty"`
	// Search using only feed_info.txt values
	FeedStartDate *tt.Date `json:"feed_start_date,omitempty"`
	// Search using only feed_info.txt values
	FeedEndDate *tt.Date `json:"feed_end_date,omitempty"`
	// Search using feed_info.txt values or calendar maximum service extent
	StartDate *tt.Date `json:"start_date,omitempty"`
	// Search using feed_info.txt values or calendar maximum service extent
	EndDate *tt.Date `json:"end_date,omitempty"`
	// Search using calendar maximum service extent
	EarliestCalendarDate *tt.Date `json:"earliest_calendar_date,omitempty"`
	// Search using calendar maximum service extent
	LatestCalendarDate *tt.Date `json:"latest_calendar_date,omitempty"`
}

type Step struct {
	Duration       *Duration `json:"duration"`
	Distance       *Distance `json:"distance"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	To             *Waypoint `json:"to,omitempty"`
	Mode           StepMode  `json:"mode"`
	Instruction    string    `json:"instruction"`
	GeometryOffset int       `json:"geometry_offset"`
}

type StopExternalReference struct {
	ID                  int     `json:"id"`
	TargetFeedOnestopID *string `json:"target_feed_onestop_id,omitempty"`
	TargetStopID        *string `json:"target_stop_id,omitempty"`
	Inactive            *bool   `json:"inactive,omitempty"`
	TargetActiveStop    *Stop   `json:"target_active_stop,omitempty"`
}

type StopFilter struct {
	OnestopID               *string        `json:"onestop_id,omitempty"`
	OnestopIds              []string       `json:"onestop_ids,omitempty"`
	AllowPreviousOnestopIds *bool          `json:"allow_previous_onestop_ids,omitempty"`
	FeedVersionSha1         *string        `json:"feed_version_sha1,omitempty"`
	FeedOnestopID           *string        `json:"feed_onestop_id,omitempty"`
	StopID                  *string        `json:"stop_id,omitempty"`
	StopCode                *string        `json:"stop_code,omitempty"`
	LocationType            *int           `json:"location_type,omitempty"`
	Serviced                *bool          `json:"serviced,omitempty"`
	Bbox                    *BoundingBox   `json:"bbox,omitempty"`
	Within                  *tt.Polygon    `json:"within,omitempty"`
	Near                    *PointRadius   `json:"near,omitempty"`
	Search                  *string        `json:"search,omitempty"`
	License                 *LicenseFilter `json:"license,omitempty"`
	ServedByOnestopIds      []string       `json:"served_by_onestop_ids,omitempty"`
	ServedByRouteType       *int           `json:"served_by_route_type,omitempty"`
	AgencyIds               []int          `json:"agency_ids,omitempty"`
}

type StopObservation struct {
	ScheduleRelationship   *string      `json:"schedule_relationship,omitempty"`
	TripStartDate          *tt.Date     `json:"trip_start_date,omitempty"`
	TripStartTime          *tt.WideTime `json:"trip_start_time,omitempty"`
	FromStopID             *string      `json:"from_stop_id,omitempty"`
	ToStopID               *string      `json:"to_stop_id,omitempty"`
	AgencyID               *string      `json:"agency_id,omitempty"`
	RouteID                *string      `json:"route_id,omitempty"`
	TripID                 *string      `json:"trip_id,omitempty"`
	StopSequence           *int         `json:"stop_sequence,omitempty"`
	Source                 *string      `json:"source,omitempty"`
	ScheduledArrivalTime   *tt.WideTime `json:"scheduled_arrival_time,omitempty"`
	ScheduledDepartureTime *tt.WideTime `json:"scheduled_departure_time,omitempty"`
	ObservedArrivalTime    *tt.WideTime `json:"observed_arrival_time,omitempty"`
	ObservedDepartureTime  *tt.WideTime `json:"observed_departure_time,omitempty"`
}

type StopObservationFilter struct {
	Source        string  `json:"source"`
	FeedVersionID int     `json:"feed_version_id"`
	TripStartDate tt.Date `json:"trip_start_date"`
}

type StopPlace struct {
	Adm1Name *string `json:"adm1_name,omitempty"`
	Adm0Name *string `json:"adm0_name,omitempty"`
	Adm0Iso  *string `json:"adm0_iso,omitempty"`
	Adm1Iso  *string `json:"adm1_iso,omitempty"`
}

type StopSetInput struct {
	ID                 *int              `json:"id,omitempty"`
	FeedVersion        *FeedVersionInput `json:"feed_version,omitempty"`
	LocationType       *int              `json:"location_type,omitempty"`
	StopCode           *string           `json:"stop_code,omitempty"`
	StopDesc           *string           `json:"stop_desc,omitempty"`
	StopID             *string           `json:"stop_id,omitempty"`
	StopName           *string           `json:"stop_name,omitempty"`
	StopTimezone       *string           `json:"stop_timezone,omitempty"`
	StopURL            *string           `json:"stop_url,omitempty"`
	WheelchairBoarding *int              `json:"wheelchair_boarding,omitempty"`
	ZoneID             *string           `json:"zone_id,omitempty"`
	PlatformCode       *string           `json:"platform_code,omitempty"`
	TtsStopName        *string           `json:"tts_stop_name,omitempty"`
	Geometry           *tt.Point         `json:"geometry,omitempty"`
	Parent             *StopSetInput     `json:"parent,omitempty"`
	Level              *LevelSetInput    `json:"level,omitempty"`
}

type StopTimeEvent struct {
	StopTimezone string       `json:"stop_timezone"`
	Scheduled    *tt.WideTime `json:"scheduled,omitempty"`
	Estimated    *tt.WideTime `json:"estimated,omitempty"`
	EstimatedUtc *time.Time   `json:"estimated_utc,omitempty"`
	Delay        *int         `json:"delay,omitempty"`
	Uncertainty  *int         `json:"uncertainty,omitempty"`
}

type StopTimeFilter struct {
	ServiceDate                  *tt.Date     `json:"service_date,omitempty"`
	UseServiceWindow             *bool        `json:"use_service_window,omitempty"`
	StartTime                    *int         `json:"start_time,omitempty"`
	EndTime                      *int         `json:"end_time,omitempty"`
	Start                        *tt.WideTime `json:"start,omitempty"`
	End                          *tt.WideTime `json:"end,omitempty"`
	Next                         *int         `json:"next,omitempty"`
	RouteOnestopIds              []string     `json:"route_onestop_ids,omitempty"`
	AllowPreviousRouteOnestopIds *bool        `json:"allow_previous_route_onestop_ids,omitempty"`
	ExcludeFirst                 *bool        `json:"exclude_first,omitempty"`
	ExcludeLast                  *bool        `json:"exclude_last,omitempty"`
}

type TripFilter struct {
	ServiceDate     *tt.Date       `json:"service_date,omitempty"`
	TripID          *string        `json:"trip_id,omitempty"`
	StopPatternID   *int           `json:"stop_pattern_id,omitempty"`
	License         *LicenseFilter `json:"license,omitempty"`
	RouteIds        []int          `json:"route_ids,omitempty"`
	RouteOnestopIds []string       `json:"route_onestop_ids,omitempty"`
	FeedVersionSha1 *string        `json:"feed_version_sha1,omitempty"`
	FeedOnestopID   *string        `json:"feed_onestop_id,omitempty"`
}

type TripStopTimeFilter struct {
	Start *tt.WideTime `json:"start,omitempty"`
	End   *tt.WideTime `json:"end,omitempty"`
}

type ValidationRealtimeResult struct {
	URL  string `json:"url"`
	JSON tt.Map `json:"json"`
}

type ValidationReport struct {
	ID                      int                           `json:"id"`
	ReportedAt              *time.Time                    `json:"reported_at,omitempty"`
	ReportedAtLocal         *time.Time                    `json:"reported_at_local,omitempty"`
	ReportedAtLocalTimezone *string                       `json:"reported_at_local_timezone,omitempty"`
	Success                 bool                          `json:"success"`
	FailureReason           *string                       `json:"failure_reason,omitempty"`
	IncludesStatic          *bool                         `json:"includes_static,omitempty"`
	IncludesRt              *bool                         `json:"includes_rt,omitempty"`
	Validator               *string                       `json:"validator,omitempty"`
	ValidatorVersion        *string                       `json:"validator_version,omitempty"`
	Errors                  []*ValidationReportErrorGroup `json:"errors"`
	Warnings                []*ValidationReportErrorGroup `json:"warnings"`
	Details                 *ValidationReportDetails      `json:"details,omitempty"`
	FeedVersionID           int                           `json:"-"`
}

type ValidationReportDetails struct {
	Sha1                 string                      `json:"sha1"`
	EarliestCalendarDate *tt.Date                    `json:"earliest_calendar_date,omitempty"`
	LatestCalendarDate   *tt.Date                    `json:"latest_calendar_date,omitempty"`
	Files                []*FeedVersionFileInfo      `json:"files"`
	ServiceLevels        []*FeedVersionServiceLevel  `json:"service_levels"`
	Agencies             []*Agency                   `json:"agencies"`
	Routes               []*Route                    `json:"routes"`
	Stops                []*Stop                     `json:"stops"`
	FeedInfos            []*FeedInfo                 `json:"feed_infos"`
	Realtime             []*ValidationRealtimeResult `json:"realtime,omitempty"`
}

type ValidationReportError struct {
	Filename                     string       `json:"filename"`
	ErrorType                    string       `json:"error_type"`
	ErrorCode                    string       `json:"error_code"`
	GroupKey                     string       `json:"group_key"`
	EntityID                     string       `json:"entity_id"`
	Field                        string       `json:"field"`
	Line                         int          `json:"line"`
	Value                        string       `json:"value"`
	Message                      string       `json:"message"`
	Geometry                     *tt.Geometry `json:"geometry,omitempty"`
	EntityJSON                   tt.Map       `json:"entity_json"`
	ID                           int          `json:"-"`
	ValidationReportErrorGroupID int          `json:"-"`
}

type ValidationReportErrorGroup struct {
	Filename           string                   `json:"filename"`
	ErrorType          string                   `json:"error_type"`
	ErrorCode          string                   `json:"error_code"`
	GroupKey           string                   `json:"group_key"`
	Field              string                   `json:"field"`
	Count              int                      `json:"count"`
	Errors             []*ValidationReportError `json:"errors"`
	ID                 int                      `json:"-"`
	ValidationReportID int                      `json:"-"`
}

type ValidationReportFilter struct {
	ReportIds        []int   `json:"report_ids,omitempty"`
	Success          *bool   `json:"success,omitempty"`
	Validator        *string `json:"validator,omitempty"`
	ValidatorVersion *string `json:"validator_version,omitempty"`
	IncludesRt       *bool   `json:"includes_rt,omitempty"`
	IncludesStatic   *bool   `json:"includes_static,omitempty"`
}

// [Vehicle Position](https://gtfs.org/reference/realtime/v2/#message-vehicleposition) message provided by a source GTFS Realtime feed.
type VehiclePosition struct {
	Vehicle             *RTVehicleDescriptor `json:"vehicle,omitempty"`
	Position            *tt.Point            `json:"position,omitempty"`
	CurrentStopSequence *int                 `json:"current_stop_sequence,omitempty"`
	StopID              *Stop                `json:"stop_id,omitempty"`
	CurrentStatus       *string              `json:"current_status,omitempty"`
	Timestamp           *time.Time           `json:"timestamp,omitempty"`
	CongestionLevel     *string              `json:"congestion_level,omitempty"`
}

type Waypoint struct {
	Lon  float64 `json:"lon"`
	Lat  float64 `json:"lat"`
	Name *string `json:"name,omitempty"`
}

type WaypointInput struct {
	Lon  float64 `json:"lon"`
	Lat  float64 `json:"lat"`
	Name *string `json:"name,omitempty"`
}

type DistanceUnit string

const (
	DistanceUnitKilometers DistanceUnit = "KILOMETERS"
	DistanceUnitMiles      DistanceUnit = "MILES"
)

var AllDistanceUnit = []DistanceUnit{
	DistanceUnitKilometers,
	DistanceUnitMiles,
}

func (e DistanceUnit) IsValid() bool {
	switch e {
	case DistanceUnitKilometers, DistanceUnitMiles:
		return true
	}
	return false
}

func (e DistanceUnit) String() string {
	return string(e)
}

func (e *DistanceUnit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DistanceUnit(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DistanceUnit", str)
	}
	return nil
}

func (e DistanceUnit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DurationUnit string

const (
	DurationUnitSeconds DurationUnit = "SECONDS"
)

var AllDurationUnit = []DurationUnit{
	DurationUnitSeconds,
}

func (e DurationUnit) IsValid() bool {
	switch e {
	case DurationUnitSeconds:
		return true
	}
	return false
}

func (e DurationUnit) String() string {
	return string(e)
}

func (e *DurationUnit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DurationUnit(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DurationUnit", str)
	}
	return nil
}

func (e DurationUnit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FeedSourceURLTypes string

const (
	FeedSourceURLTypesStaticCurrent            FeedSourceURLTypes = "static_current"
	FeedSourceURLTypesStaticHistoric           FeedSourceURLTypes = "static_historic"
	FeedSourceURLTypesStaticPlanned            FeedSourceURLTypes = "static_planned"
	FeedSourceURLTypesStaticHypothetical       FeedSourceURLTypes = "static_hypothetical"
	FeedSourceURLTypesRealtimeVehiclePositions FeedSourceURLTypes = "realtime_vehicle_positions"
	FeedSourceURLTypesRealtimeTripUpdates      FeedSourceURLTypes = "realtime_trip_updates"
	FeedSourceURLTypesRealtimeAlerts           FeedSourceURLTypes = "realtime_alerts"
	FeedSourceURLTypesGbfsAutoDiscovery        FeedSourceURLTypes = "gbfs_auto_discovery"
	FeedSourceURLTypesMdsProvider              FeedSourceURLTypes = "mds_provider"
)

var AllFeedSourceURLTypes = []FeedSourceURLTypes{
	FeedSourceURLTypesStaticCurrent,
	FeedSourceURLTypesStaticHistoric,
	FeedSourceURLTypesStaticPlanned,
	FeedSourceURLTypesStaticHypothetical,
	FeedSourceURLTypesRealtimeVehiclePositions,
	FeedSourceURLTypesRealtimeTripUpdates,
	FeedSourceURLTypesRealtimeAlerts,
	FeedSourceURLTypesGbfsAutoDiscovery,
	FeedSourceURLTypesMdsProvider,
}

func (e FeedSourceURLTypes) IsValid() bool {
	switch e {
	case FeedSourceURLTypesStaticCurrent, FeedSourceURLTypesStaticHistoric, FeedSourceURLTypesStaticPlanned, FeedSourceURLTypesStaticHypothetical, FeedSourceURLTypesRealtimeVehiclePositions, FeedSourceURLTypesRealtimeTripUpdates, FeedSourceURLTypesRealtimeAlerts, FeedSourceURLTypesGbfsAutoDiscovery, FeedSourceURLTypesMdsProvider:
		return true
	}
	return false
}

func (e FeedSourceURLTypes) String() string {
	return string(e)
}

func (e *FeedSourceURLTypes) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FeedSourceURLTypes(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FeedSourceUrlTypes", str)
	}
	return nil
}

func (e FeedSourceURLTypes) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Type of data contained in a source feed
type FeedSpecTypes string

const (
	FeedSpecTypesGtfs   FeedSpecTypes = "GTFS"
	FeedSpecTypesGtfsRt FeedSpecTypes = "GTFS_RT"
	FeedSpecTypesGbfs   FeedSpecTypes = "GBFS"
	FeedSpecTypesMds    FeedSpecTypes = "MDS"
)

var AllFeedSpecTypes = []FeedSpecTypes{
	FeedSpecTypesGtfs,
	FeedSpecTypesGtfsRt,
	FeedSpecTypesGbfs,
	FeedSpecTypesMds,
}

func (e FeedSpecTypes) IsValid() bool {
	switch e {
	case FeedSpecTypesGtfs, FeedSpecTypesGtfsRt, FeedSpecTypesGbfs, FeedSpecTypesMds:
		return true
	}
	return false
}

func (e FeedSpecTypes) String() string {
	return string(e)
}

func (e *FeedSpecTypes) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FeedSpecTypes(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FeedSpecTypes", str)
	}
	return nil
}

func (e FeedSpecTypes) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ImportStatus string

const (
	ImportStatusSuccess    ImportStatus = "SUCCESS"
	ImportStatusError      ImportStatus = "ERROR"
	ImportStatusInProgress ImportStatus = "IN_PROGRESS"
)

var AllImportStatus = []ImportStatus{
	ImportStatusSuccess,
	ImportStatusError,
	ImportStatusInProgress,
}

func (e ImportStatus) IsValid() bool {
	switch e {
	case ImportStatusSuccess, ImportStatusError, ImportStatusInProgress:
		return true
	}
	return false
}

func (e ImportStatus) String() string {
	return string(e)
}

func (e *ImportStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ImportStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ImportStatus", str)
	}
	return nil
}

func (e ImportStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LicenseValue string

const (
	LicenseValueYes       LicenseValue = "YES"
	LicenseValueNo        LicenseValue = "NO"
	LicenseValueExcludeNo LicenseValue = "EXCLUDE_NO"
	LicenseValueUnknown   LicenseValue = "UNKNOWN"
)

var AllLicenseValue = []LicenseValue{
	LicenseValueYes,
	LicenseValueNo,
	LicenseValueExcludeNo,
	LicenseValueUnknown,
}

func (e LicenseValue) IsValid() bool {
	switch e {
	case LicenseValueYes, LicenseValueNo, LicenseValueExcludeNo, LicenseValueUnknown:
		return true
	}
	return false
}

func (e LicenseValue) String() string {
	return string(e)
}

func (e *LicenseValue) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LicenseValue(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LicenseValue", str)
	}
	return nil
}

func (e LicenseValue) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PlaceAggregationLevel string

const (
	PlaceAggregationLevelAdm0         PlaceAggregationLevel = "ADM0"
	PlaceAggregationLevelAdm0Adm1     PlaceAggregationLevel = "ADM0_ADM1"
	PlaceAggregationLevelAdm0Adm1City PlaceAggregationLevel = "ADM0_ADM1_CITY"
	PlaceAggregationLevelAdm0City     PlaceAggregationLevel = "ADM0_CITY"
	PlaceAggregationLevelAdm1City     PlaceAggregationLevel = "ADM1_CITY"
	PlaceAggregationLevelCity         PlaceAggregationLevel = "CITY"
)

var AllPlaceAggregationLevel = []PlaceAggregationLevel{
	PlaceAggregationLevelAdm0,
	PlaceAggregationLevelAdm0Adm1,
	PlaceAggregationLevelAdm0Adm1City,
	PlaceAggregationLevelAdm0City,
	PlaceAggregationLevelAdm1City,
	PlaceAggregationLevelCity,
}

func (e PlaceAggregationLevel) IsValid() bool {
	switch e {
	case PlaceAggregationLevelAdm0, PlaceAggregationLevelAdm0Adm1, PlaceAggregationLevelAdm0Adm1City, PlaceAggregationLevelAdm0City, PlaceAggregationLevelAdm1City, PlaceAggregationLevelCity:
		return true
	}
	return false
}

func (e PlaceAggregationLevel) String() string {
	return string(e)
}

func (e *PlaceAggregationLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PlaceAggregationLevel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PlaceAggregationLevel", str)
	}
	return nil
}

func (e PlaceAggregationLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ScheduleRelationship string

const (
	ScheduleRelationshipScheduled   ScheduleRelationship = "SCHEDULED"
	ScheduleRelationshipAdded       ScheduleRelationship = "ADDED"
	ScheduleRelationshipUnscheduled ScheduleRelationship = "UNSCHEDULED"
	ScheduleRelationshipCanceled    ScheduleRelationship = "CANCELED"
)

var AllScheduleRelationship = []ScheduleRelationship{
	ScheduleRelationshipScheduled,
	ScheduleRelationshipAdded,
	ScheduleRelationshipUnscheduled,
	ScheduleRelationshipCanceled,
}

func (e ScheduleRelationship) IsValid() bool {
	switch e {
	case ScheduleRelationshipScheduled, ScheduleRelationshipAdded, ScheduleRelationshipUnscheduled, ScheduleRelationshipCanceled:
		return true
	}
	return false
}

func (e ScheduleRelationship) String() string {
	return string(e)
}

func (e *ScheduleRelationship) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ScheduleRelationship(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ScheduleRelationship", str)
	}
	return nil
}

func (e ScheduleRelationship) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StepMode string

const (
	StepModeWalk    StepMode = "WALK"
	StepModeAuto    StepMode = "AUTO"
	StepModeBicycle StepMode = "BICYCLE"
	StepModeTransit StepMode = "TRANSIT"
	StepModeLine    StepMode = "LINE"
)

var AllStepMode = []StepMode{
	StepModeWalk,
	StepModeAuto,
	StepModeBicycle,
	StepModeTransit,
	StepModeLine,
}

func (e StepMode) IsValid() bool {
	switch e {
	case StepModeWalk, StepModeAuto, StepModeBicycle, StepModeTransit, StepModeLine:
		return true
	}
	return false
}

func (e StepMode) String() string {
	return string(e)
}

func (e *StepMode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StepMode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StepMode", str)
	}
	return nil
}

func (e StepMode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
