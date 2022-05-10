// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/interline-io/transitland-lib/tl"
)

type AgencyFilter struct {
	OnestopID       *string `json:"onestop_id"`
	FeedVersionSha1 *string `json:"feed_version_sha1"`
	FeedOnestopID   *string `json:"feed_onestop_id"`
	AgencyID        *string `json:"agency_id"`
	// Search for records with this GTFS agency_name
	AgencyName *string     `json:"agency_name"`
	Within     *tl.Polygon `json:"within"`
	// Search for agencies within a radius
	Near *PointRadius `json:"near"`
	// Full text search
	Search *string `json:"search"`
	// Search by city name (provided by Natural Earth)
	CityName *string `json:"city_name"`
	// Search by country name (provided by Natural Earth)
	Adm0Name *string `json:"adm0_name"`
	// Search by state/province/division name (provided by Natural Earth)
	Adm1Name *string `json:"adm1_name"`
	// Search by country 2 letter ISO 3166 code (provided by Natural Earth)
	Adm0Iso *string `json:"adm0_iso"`
	// Search by state/province/division ISO 3166-2 code (provided by Natural Earth)
	Adm1Iso *string `json:"adm1_iso"`
}

type AgencyPlaceFilter struct {
	MinRank *float64 `json:"min_rank"`
}

// [Alert](https://gtfs.org/reference/realtime/v2/#message-alert) message, also called a service alert, provided by a source GTFS Realtime feed.
type Alert struct {
	ActivePeriod       []*RTTimeRange   `json:"active_period"`
	Cause              *string          `json:"cause"`
	Effect             *string          `json:"effect"`
	HeaderText         []*RTTranslation `json:"header_text"`
	DescriptionText    []*RTTranslation `json:"description_text"`
	TtsHeaderText      []*RTTranslation `json:"tts_header_text"`
	TtsDescriptionText []*RTTranslation `json:"tts_description_text"`
	URL                []*RTTranslation `json:"url"`
	SeverityLevel      *string          `json:"severity_level"`
}

type CalendarDateFilter struct {
	Date          *tl.Date `json:"date"`
	ExceptionType *int     `json:"exception_type"`
}

type DirectionRequest struct {
	To       *WaypointInput `json:"to"`
	From     *WaypointInput `json:"from"`
	Mode     StepMode       `json:"mode"`
	DepartAt *time.Time     `json:"depart_at"`
}

type Directions struct {
	Success     bool         `json:"success"`
	Exception   *string      `json:"exception"`
	DataSource  *string      `json:"data_source"`
	Origin      *Waypoint    `json:"origin"`
	Destination *Waypoint    `json:"destination"`
	Duration    *Duration    `json:"duration"`
	Distance    *Distance    `json:"distance"`
	StartTime   *time.Time   `json:"start_time"`
	EndTime     *time.Time   `json:"end_time"`
	Itineraries []*Itinerary `json:"itineraries"`
}

type Distance struct {
	Distance float64      `json:"distance"`
	Units    DistanceUnit `json:"units"`
}

type Duration struct {
	Duration float64      `json:"duration"`
	Units    DurationUnit `json:"units"`
}

type FeedFetchFilter struct {
	Success *bool `json:"success"`
}

type FeedFilter struct {
	// Search for feed with a specific Onestop ID
	OnestopID *string `json:"onestop_id"`
	// Search for feeds of certain data types
	Spec []FeedSpecTypes `json:"spec"`
	// Search for feeds with or without a fetch error
	FetchError *bool `json:"fetch_error"`
	// Search for feeds by their import status
	ImportStatus *ImportStatus `json:"import_status"`
	// Full text search
	Search *string `json:"search"`
	// Search for feeds with a tag
	Tags *tl.Tags `json:"tags"`
	// Search for feeds by their source URLs
	SourceURL *FeedSourceURL `json:"source_url"`
}

type FeedSourceURL struct {
	URL           *string             `json:"url"`
	Type          *FeedSourceURLTypes `json:"type"`
	CaseSensitive *bool               `json:"case_sensitive"`
}

type FeedVersionDeleteResult struct {
	Success bool `json:"success"`
}

type FeedVersionFilter struct {
	FeedOnestopID *string `json:"feed_onestop_id"`
	Sha1          *string `json:"sha1"`
	FeedIds       []int   `json:"feed_ids"`
}

type FeedVersionServiceLevelFilter struct {
	StartDate  *tl.Date `json:"start_date"`
	EndDate    *tl.Date `json:"end_date"`
	AllRoutes  *bool    `json:"all_routes"`
	DistinctOn *string  `json:"distinct_on"`
	RouteIds   []string `json:"route_ids"`
}

type FeedVersionSetInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type FeedVersionUnimportResult struct {
	Success bool `json:"success"`
}

type Itinerary struct {
	Duration  *Duration `json:"duration"`
	Distance  *Distance `json:"distance"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	From      *Waypoint `json:"from"`
	To        *Waypoint `json:"to"`
	Legs      []*Leg    `json:"legs"`
}

type Leg struct {
	Duration  *Duration     `json:"duration"`
	Distance  *Distance     `json:"distance"`
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
	From      *Waypoint     `json:"from"`
	To        *Waypoint     `json:"to"`
	Steps     []*Step       `json:"steps"`
	Geometry  tl.LineString `json:"geometry"`
}

type OperatorFilter struct {
	Merged        *bool    `json:"merged"`
	OnestopID     *string  `json:"onestop_id"`
	FeedOnestopID *string  `json:"feed_onestop_id"`
	AgencyID      *string  `json:"agency_id"`
	Search        *string  `json:"search"`
	Tags          *tl.Tags `json:"tags"`
	CityName      *string  `json:"city_name"`
	Adm0Name      *string  `json:"adm0_name"`
	Adm1Name      *string  `json:"adm1_name"`
	Adm0Iso       *string  `json:"adm0_iso"`
	Adm1Iso       *string  `json:"adm1_iso"`
}

type PathwayFilter struct {
	PathwayMode *int `json:"pathway_mode"`
}

type PointRadius struct {
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Radius float64 `json:"radius"`
}

// See https://gtfs.org/reference/realtime/v2/#message-timerange
type RTTimeRange struct {
	Start *int `json:"start"`
	End   *int `json:"end"`
}

// See https://gtfs.org/reference/realtime/v2/#message-translatedstring
type RTTranslation struct {
	Text     string  `json:"text"`
	Language *string `json:"language"`
}

// See https://gtfs.org/reference/realtime/v2/#message-tripdescriptor
type RTTripDescriptor struct {
	TripID               *string      `json:"trip_id"`
	RouteID              *string      `json:"route_id"`
	DirectionID          *int         `json:"direction_id"`
	StartTime            *tl.WideTime `json:"start_time"`
	StartDate            *tl.Date     `json:"start_date"`
	ScheduleRelationship *string      `json:"schedule_relationship"`
}

// See https://gtfs.org/reference/realtime/v2/#message-vehicledescriptor
type RTVehicleDescriptor struct {
	ID           *string `json:"id"`
	Label        *string `json:"label"`
	LicensePlate *string `json:"license_plate"`
}

type RouteFilter struct {
	OnestopID               *string      `json:"onestop_id"`
	OnestopIds              []string     `json:"onestop_ids"`
	AllowPreviousOnestopIds *bool        `json:"allow_previous_onestop_ids"`
	FeedVersionSha1         *string      `json:"feed_version_sha1"`
	FeedOnestopID           *string      `json:"feed_onestop_id"`
	RouteID                 *string      `json:"route_id"`
	RouteType               *int         `json:"route_type"`
	Within                  *tl.Polygon  `json:"within"`
	Near                    *PointRadius `json:"near"`
	Search                  *string      `json:"search"`
	OperatorOnestopID       *string      `json:"operator_onestop_id"`
	AgencyIds               []int        `json:"agency_ids"`
}

type Step struct {
	Duration       *Duration `json:"duration"`
	Distance       *Distance `json:"distance"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	To             *Waypoint `json:"to"`
	Mode           StepMode  `json:"mode"`
	Instruction    string    `json:"instruction"`
	GeometryOffset int       `json:"geometry_offset"`
}

type StopFilter struct {
	OnestopID               *string      `json:"onestop_id"`
	OnestopIds              []string     `json:"onestop_ids"`
	AllowPreviousOnestopIds *bool        `json:"allow_previous_onestop_ids"`
	FeedVersionSha1         *string      `json:"feed_version_sha1"`
	FeedOnestopID           *string      `json:"feed_onestop_id"`
	StopID                  *string      `json:"stop_id"`
	StopCode                *string      `json:"stop_code"`
	Within                  *tl.Polygon  `json:"within"`
	Near                    *PointRadius `json:"near"`
	Search                  *string      `json:"search"`
	ServedByOnestopIds      []string     `json:"served_by_onestop_ids"`
	AgencyIds               []int        `json:"agency_ids"`
}

type StopTimeFilter struct {
	ServiceDate                  *tl.Date `json:"service_date"`
	UseServiceWindow             *bool    `json:"use_service_window"`
	StartTime                    *int     `json:"start_time"`
	EndTime                      *int     `json:"end_time"`
	Next                         *int     `json:"next"`
	RouteOnestopIds              []string `json:"route_onestop_ids"`
	AllowPreviousRouteOnestopIds *bool    `json:"allow_previous_route_onestop_ids"`
}

type TripFilter struct {
	ServiceDate     *tl.Date `json:"service_date"`
	TripID          *string  `json:"trip_id"`
	RouteIds        []int    `json:"route_ids"`
	RouteOnestopIds []string `json:"route_onestop_ids"`
	FeedVersionSha1 *string  `json:"feed_version_sha1"`
	FeedOnestopID   *string  `json:"feed_onestop_id"`
}

// [Vehicle Position](https://gtfs.org/reference/realtime/v2/#message-vehicleposition) message provided by a source GTFS Realtime feed.
type VehiclePosition struct {
	Vehicle             *RTVehicleDescriptor `json:"vehicle"`
	Position            *tl.Point            `json:"position"`
	CurrentStopSequence *int                 `json:"current_stop_sequence"`
	StopID              *Stop                `json:"stop_id"`
	CurrentStatus       *string              `json:"current_status"`
	Timestamp           *time.Time           `json:"timestamp"`
	CongestionLevel     *string              `json:"congestion_level"`
}

type Waypoint struct {
	Lon  float64 `json:"lon"`
	Lat  float64 `json:"lat"`
	Name *string `json:"name"`
}

type WaypointInput struct {
	Lon  float64 `json:"lon"`
	Lat  float64 `json:"lat"`
	Name *string `json:"name"`
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

type Role string

const (
	RoleAnon  Role = "ANON"
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

var AllRole = []Role{
	RoleAnon,
	RoleAdmin,
	RoleUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAnon, RoleAdmin, RoleUser:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
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
