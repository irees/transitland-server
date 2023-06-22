package rest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestAgencyRequest(t *testing.T) {
	srv, te := testRestConfig(t)
	fv := "e535eb2b3b9ac3ef15d82c56575e914575e732e0"
	testcases := []testRest{
		{
			name:         "basic",
			h:            AgencyRequest{},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"caltrain-ca-us", "BART", ""},
		},
		// this used to be caltrain but now bart is imported first.
		{
			name:         "feed_version_sha1",
			h:            AgencyRequest{FeedVersionSHA1: fv},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"BART"},
		},
		{
			name:         "feed_onestop_id",
			h:            AgencyRequest{FeedOnestopID: "BA"},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"BART"},
		},
		{
			name:         "feed_onestop_id,agency_id",
			h:            AgencyRequest{FeedOnestopID: "BA", AgencyID: "BART"},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"BART"},
		},
		{
			name:         "agency_id",
			h:            AgencyRequest{AgencyID: "BART"},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"BART"},
		},
		{
			name:         "agency_name",
			h:            AgencyRequest{AgencyName: "Bay Area Rapid Transit"},
			selector:     "agencies.#.agency_name",
			expectSelect: []string{"Bay Area Rapid Transit"},
		},
		{
			name:         "onestop_id",
			h:            AgencyRequest{OnestopID: "o-9q9-bayarearapidtransit"},
			selector:     "agencies.#.onestop_id",
			expectSelect: []string{"o-9q9-bayarearapidtransit"},
		},
		{
			name:         "onestop_id,feed_version_sha1",
			h:            AgencyRequest{OnestopID: "o-9q9-bayarearapidtransit", FeedVersionSHA1: fv},
			selector:     "agencies.#.feed_version.sha1",
			expectSelect: []string{fv},
		},
		{
			name:         "agency_key onestop_id",
			h:            AgencyRequest{AgencyKey: "o-9q9-bayarearapidtransit"},
			selector:     "agencies.#.onestop_id",
			expectSelect: []string{"o-9q9-bayarearapidtransit"},
		},
		{
			name:         "lat,lon,radius 10m",
			h:            AgencyRequest{Lon: -122.407974, Lat: 37.784471, Radius: 10},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"BART"},
		},
		{
			name:         "lat,lon,radius 2000m",
			h:            AgencyRequest{Lon: -122.407974, Lat: 37.784471, Radius: 2000},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"caltrain-ca-us", "BART"},
		},
		{
			name:         "search",
			h:            AgencyRequest{Search: "caltrain"},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"caltrain-ca-us"},
		},
		{
			name:         "adm0name",
			h:            AgencyRequest{Adm0Name: "united states of america"},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"caltrain-ca-us", "BART", ""},
		},
		{
			name:         "adm1name",
			h:            AgencyRequest{Adm1Name: "california"},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"caltrain-ca-us", "BART"},
		},
		{
			name:         "adm0iso",
			h:            AgencyRequest{Adm0Iso: "us"},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"caltrain-ca-us", "BART", ""},
		},
		{
			name:         "adm1iso:us-ca",
			h:            AgencyRequest{Adm1Iso: "us-ca"},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"caltrain-ca-us", "BART"},
		},
		{
			name:         "adm1iso:us-ny",
			h:            AgencyRequest{Adm1Iso: "us-ny"},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{},
		},
		{
			name:         "city_name:san jose",
			h:            AgencyRequest{CityName: "san jose"},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"caltrain-ca-us"},
		},
		{
			name:         "city_name:oakland",
			h:            AgencyRequest{CityName: "berkeley"},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"BART"},
		},
		{
			name:         "city_name:new york city",
			h:            AgencyRequest{CityName: "new york city"},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{},
		},
		{
			name:         "feed:agency_id",
			h:            AgencyRequest{AgencyKey: "CT:caltrain-ca-us"},
			selector:     "agencies.#.agency_id",
			expectSelect: []string{"caltrain-ca-us"},
		},
		{
			name: "include_alerts:true",
			h:    AgencyRequest{AgencyKey: "BA:BART", IncludeAlerts: true},
			f: func(t *testing.T, jj string) {
				a := gjson.Get(jj, "agencies.0.alerts").Array()
				assert.Equal(t, 2, len(a), "alert count")
			},
		},
		{
			name: "include_alerts:false",
			h:    AgencyRequest{AgencyKey: "BA:BART", IncludeAlerts: false},
			f: func(t *testing.T, jj string) {
				a := gjson.Get(jj, "agencies.0.alerts").Array()
				assert.Equal(t, 0, len(a), "alert count")
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			testquery(t, srv, te, tc)
		})
	}
}

func TestAgencyRequest_Pagination(t *testing.T) {
	srv, te := testRestConfig(t)
	allEnts, err := te.Finder.FindAgencies(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	allIds := []string{}
	for _, ent := range allEnts {
		allIds = append(allIds, ent.AgencyID)
	}
	testcases := []testRest{
		{
			name:         "limit:1",
			h:            AgencyRequest{Limit: 1},
			selector:     "agencies.#.agency_id",
			expectSelect: nil,
			expectLength: 1,
		},
		{
			name:         "pagination exists",
			h:            AgencyRequest{},
			selector:     "meta.after",
			expectSelect: nil,
			expectLength: 1,
		}, // just check presence
		{
			name:         "pagination limit 1",
			h:            AgencyRequest{Limit: 1},
			selector:     "agencies.#.agency_id",
			expectSelect: allIds[:1],
		},
		{
			name:         "pagination after 1",
			h:            AgencyRequest{Limit: 1, After: allEnts[0].ID},
			selector:     "agencies.#.agency_id",
			expectSelect: allIds[1:2],
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			testquery(t, srv, te, tc)
		})
	}
}

func TestAgencyRequest_License(t *testing.T) {
	testcases := []testRest{
		{
			name: "license:share_alike_optional yes",
			h:    AgencyRequest{Limit: 10_000, LicenseFilter: LicenseFilter{LicenseShareAlikeOptional: "yes"}}, selector: "agencies.#.agency_id",
			expectSelect: []string{""},
		},
		{
			name: "license:share_alike_optional no",
			h:    AgencyRequest{Limit: 10_000, LicenseFilter: LicenseFilter{LicenseShareAlikeOptional: "no"}}, selector: "agencies.#.agency_id",
			expectSelect: []string{"BART"},
		},
		{
			name: "license:share_alike_optional exclude_no",
			h:    AgencyRequest{Limit: 10_000, LicenseFilter: LicenseFilter{LicenseShareAlikeOptional: "exclude_no"}}, selector: "agencies.#.agency_id",
			expectSelect: []string{"caltrain-ca-us", ""},
		},
		{
			name: "license:commercial_use_allowed yes",
			h:    AgencyRequest{Limit: 10_000, LicenseFilter: LicenseFilter{LicenseCommercialUseAllowed: "yes"}}, selector: "agencies.#.agency_id",
			expectSelect: []string{""},
		},
		{
			name: "license:commercial_use_allowed no",
			h:    AgencyRequest{Limit: 10_000, LicenseFilter: LicenseFilter{LicenseCommercialUseAllowed: "no"}}, selector: "agencies.#.agency_id",
			expectSelect: []string{"BART"},
		},
		{
			name: "license:commercial_use_allowed exclude_no",
			h:    AgencyRequest{Limit: 10_000, LicenseFilter: LicenseFilter{LicenseCommercialUseAllowed: "exclude_no"}}, selector: "agencies.#.agency_id",
			expectSelect: []string{"caltrain-ca-us", ""},
		},
		{
			name: "license:create_derived_product yes",
			h:    AgencyRequest{Limit: 10_000, LicenseFilter: LicenseFilter{LicenseCreateDerivedProduct: "yes"}}, selector: "agencies.#.agency_id",
			expectSelect: []string{""},
		},
		{
			name: "license:create_derived_product no",
			h:    AgencyRequest{Limit: 10_000, LicenseFilter: LicenseFilter{LicenseCreateDerivedProduct: "no"}}, selector: "agencies.#.agency_id",
			expectSelect: []string{"BART"},
		},
		{
			name: "license:create_derived_product exclude_no",
			h:    AgencyRequest{Limit: 10_000, LicenseFilter: LicenseFilter{LicenseCreateDerivedProduct: "exclude_no"}}, selector: "agencies.#.agency_id",
			expectSelect: []string{"caltrain-ca-us", ""},
		},
	}
	srv, te := testRestConfig(t)
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			testquery(t, srv, te, tc)
		})
	}
}
