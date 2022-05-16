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
	ID            types.String  `tfsdk:"id"`
	Name          types.String  `tfsdk:"name"`
	Description   types.String  `tfsdk:"description"`
	Code          types.String  `tfsdk:"code"`
	Type          types.String  `tfsdk:"type"`
	Rule          types.String  `tfsdk:"rule"`
	Grouping      types.String  `tfsdk:"grouping"`
	GroupingL2    types.String  `tfsdk:"grouping_l2"`
	GroupingL3    types.String  `tfsdk:"grouping_l3"`
	Enabled       types.Bool    `tfsdk:"enabled"`
	EventTags     types.List    `tfsdk:"event_tags"`
	BuilderConfig BuilderConfig `tfsdk:"builder_config"`
}

type BuilderConfig struct {
	ID              types.String        `tfsdk:"id"`
	CustomerID      types.String        `tfsdk:"customer_id"`
	TableName       types.String        `tfsdk:"table_name"`
	Added           types.Bool          `tfsdk:"added"`
	MatchesFilter   types.Bool          `tfsdk:"matches_filter"`
	Filters         BuilderConfigFilter `tfsdk:"filters"`
	Severity        types.String        `tfsdk:"severity"`
	Key             types.String        `tfsdk:"key"`
	ValueField      types.String        `tfsdk:"value_field"`
	//AutoAlertConfig AutoAlertConfig     `tfsdk:"auto_alert_config"`
}

type AutoAlertConfig struct {
	RaiseAlert   types.Bool `tfsdk:"raise_alert"`
	DisableAlert types.Bool `tfsdk:"disable_alert"`
}

type BuilderConfigFilter struct {
	//And             []BuilderConfigFilter `tfsdk:"and"`
	//Or              []BuilderConfigFilter `tfsdk:"or"`
	Not             types.Bool   `tfsdk:"not"`
	Name            types.String `tfsdk:"name"`
	Value           types.String `tfsdk:"value"`
	Operator        types.String `tfsdk:"operator"`
	IsDate          types.Bool   `tfsdk:"is_date"`
	IsVersion       types.Bool   `tfsdk:"is_version"`
	IsWordMatch     types.Bool   `tfsdk:"is_word_match"`
	CaseInsensitive types.Bool   `tfsdk:"case_insensitive"`
}
