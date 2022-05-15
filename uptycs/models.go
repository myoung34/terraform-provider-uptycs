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
	EventTags   types.List   `tfsdk:"event_tags"`
	BuilderConfig BuilderConfig `tfsdk:"builder_config"`
}

type BuilderConfig struct {
	ID              types.String        `tfsdk:"id"`
	CustomerID      types.String        `tfsdk:"customer_id"`
	TableName       types.String        `tfsdk:"table_name"`
	Added           types.Bool          `tfsdk:"added"`
	MatchesFilter   types.Bool          `tfsdk:"matches_filter"`
	//Filters         BuilderConfigFilter `tfsdk:"filters"`
	Severity        types.String        `tfsdk:"severity"`
	Key             types.String        `tfsdk:"key"`
	ValueField      types.String        `tfsdk:"value_field"`
	AutoAlertConfig AutoAlertConfig     `tfsdk:"auto_alert_config"`
}

type AutoAlertConfig struct {
	RaiseAlert   types.Bool `tfsdk:"raise_alert"`
	DisableAlert types.Bool `tfsdk:"disable_alert"`
}

type ArrayOrString []types.String

type BuilderConfigFilter struct {
	//And             []BuilderConfigFilter `tfsdk:"and"`
	//Or              []BuilderConfigFilter `tfsdk:"or"`
	Not             types.Bool            `tfsdk:"not"`
	Name            types.String          `tfsdk:"name"`
	//Value           ArrayOrString         `tfsdk:"value"`
	Operator        types.String          `tfsdk:"operator"`
	IsDate          types.Bool            `tfsdk:"is_date"`
	IsVersion       types.Bool            `tfsdk:"is_version"`
	IsWordMatch     types.Bool            `tfsdk:"is_wordMatch"`
	CaseInsensitive types.Bool            `tfsdk:"case_insensitive"`
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
