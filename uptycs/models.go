package uptycs

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AlertRule struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Code        types.String `tfsdk:"code"`
	Type        types.String `tfsdk:"type"`
	Rule        types.String `tfsdk:"rule"`
	Grouping    types.String `tfsdk:"grouping"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	GroupingL2  types.String `tfsdk:"grouping_l2"`
	GroupingL3  types.String `tfsdk:"grouping_l3"`
	SQLConfig   SQLConfig    `tfsdk:"sql_config"`
}

type SQLConfig struct {
	IntervalSeconds int `tfsdk:"interval_seconds"`
}

type EventRule struct {
	ID          types.String   `tfsdk:"id"`
	Name        types.String   `tfsdk:"name"`
	Description types.String   `tfsdk:"description"`
	Code        types.String   `tfsdk:"code"`
	Type        types.String   `tfsdk:"type"`
	Rule        types.String   `tfsdk:"rule"`
	Grouping    types.String   `tfsdk:"grouping"`
	Enabled     types.Bool     `tfsdk:"enabled"`
	GroupingL2  types.String   `tfsdk:"grouping_l2"`
	GroupingL3  types.String   `tfsdk:"grouping_l3"`
	EventTags   types.ListType `tfsdk:"event_tags"`
}

//BuilderConfig: uptycs.BuilderConfig{
//	TableName:     "process_open_sockets",
//	Added:         true,
//	MatchesFilter: true,
//	Filters: uptycs.BuilderConfigFilter{
//		And: []uptycs.BuilderConfigFilter{
//			{
//				Name:     "remote_address",
//				Operator: "MATCHES_REGEX",
//				Value:    uptycs.ArrayOrString{"^172.(1[6-9]|2[0-9]|3[01])|^10.|^192.168."},
//				Not:      true,
//			},
//		},
//	},
//	Severity:   "low",
//	Key:        "Test",
//	ValueField: "pid",
//},
