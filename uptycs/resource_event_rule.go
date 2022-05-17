package uptycs

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/myoung34/uptycs-client-go/uptycs"
)

type resourceEventRuleType struct{}

// Alert Rule Resource schema
func (r resourceEventRuleType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
			},
			"description": {
				Type:     types.StringType,
				Required: true,
			},
			"code": {
				Type:     types.StringType,
				Required: true,
				Computed: false,
			},
			"type": {
				Type:     types.StringType,
				Required: true,
				Computed: false,
			},
			"rule": {
				Type:     types.StringType,
				Required: true,
			},
			"grouping": {
				Type:     types.StringType,
				Required: true,
			},
			"grouping_l2": {
				Type:     types.StringType,
				Optional: true,
			},
			"grouping_l3": {
				Type:     types.StringType,
				Optional: true,
			},
			"enabled": {
				Type:     types.BoolType,
				Optional: true,
			},
			"event_tags": {
				Type:     types.ListType{ElemType: types.StringType},
				Required: true,
			},

			"builder_config": {
				Required: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						Type:     types.StringType,
						Computed: true,
					},
					"customer_id": {
						Type:     types.StringType,
						Computed: true,
					},
					"table_name": {
						Type:     types.StringType,
						Optional: true,
					},
					"added": {
						Type:     types.BoolType,
						Optional: true,
					},
					"matches_filter": {
						Type:     types.BoolType,
						Optional: true,
					},
					"severity": {
						Type:     types.StringType,
						Optional: true,
					},
					"key": {
						Type:     types.StringType,
						Optional: true,
					},
					"value_field": {
						Type:     types.StringType,
						Optional: true,
					},
					"filters": {
						Required: true,
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"value": {
								Type:     types.StringType,
								Optional: true,
							},
							"not": {
								Type:     types.BoolType,
								Optional: true,
							},
							"name": {
								Type:     types.StringType,
								Optional: true,
							},
							"operator": {
								Type:     types.StringType,
								Optional: true,
							},
							"is_date": {
								Type:     types.BoolType,
								Optional: true,
							},
							"is_version": {
								Type:     types.BoolType,
								Optional: true,
							},
							"is_word_match": {
								Type:     types.BoolType,
								Optional: true,
							},
							"case_insensitive": {
								Type:     types.BoolType,
								Optional: true,
							},
						}),
					},
					//"auto_alert_config": {
					//	Optional: true,
					//	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					//		"raise_alert": {
					//			Type:     types.BoolType,
					//			Optional: true,
					//		},
					//		"disable_alert": {
					//			Type:     types.BoolType,
					//			Optional: true,
					//		},
					//	}),
					//},
				}),
			},
		},
	}, nil
}

// New resource instance
func (r resourceEventRuleType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceEventRule{
		p: *(p.(*provider)),
	}, nil
}

type resourceEventRule struct {
	p provider
}

// Create a new resource
func (r resourceEventRule) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	// Retrieve values from plan
	var plan EventRule
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var tags []string
	plan.EventTags.ElementsAs(ctx, &tags, false)

	eventRuleResp, err := r.p.client.CreateEventRule(uptycs.EventRule{
		Name:        plan.Name.Value,
		Code:        plan.Code.Value,
		Description: plan.Description.Value,
		Rule:        plan.Rule.Value,
		Type:        plan.Type.Value,
		Enabled:     plan.Enabled.Value,
		Custom:      true,
		Grouping:    plan.Grouping.Value,
		GroupingL2:  plan.GroupingL2.Value,
		GroupingL3:  plan.GroupingL3.Value,
		EventTags:   tags,
		BuilderConfig: uptycs.BuilderConfig{
			Filters: uptycs.BuilderConfigFilter{
				And: []uptycs.BuilderConfigFilter{
					uptycs.BuilderConfigFilter{
						Not:             plan.BuilderConfig.Filters.Not.Value,
						Name:            plan.BuilderConfig.Filters.Name.Value,
						Operator:        plan.BuilderConfig.Filters.Operator.Value,
						Value:           uptycs.ArrayOrString{plan.BuilderConfig.Filters.Value.Value},
						IsDate:          plan.BuilderConfig.Filters.IsDate.Value,
						IsVersion:       plan.BuilderConfig.Filters.IsVersion.Value,
						IsWordMatch:     plan.BuilderConfig.Filters.IsWordMatch.Value,
						CaseInsensitive: plan.BuilderConfig.Filters.CaseInsensitive.Value,
					},
				},
			},
			TableName:     plan.BuilderConfig.TableName.Value,
			Added:         plan.BuilderConfig.Added.Value,
			MatchesFilter: plan.BuilderConfig.MatchesFilter.Value,
			Severity:      plan.BuilderConfig.Severity.Value,
			Key:           plan.BuilderConfig.Key.Value,
			ValueField:    plan.BuilderConfig.ValueField.Value,
			//AutoAlertConfig: uptycs.AutoAlertConfig{
			//	DisableAlert: plan.BuilderConfig.AutoAlertConfig.DisableAlert.Value,
			//	RaiseAlert:   plan.BuilderConfig.AutoAlertConfig.RaiseAlert.Value,
			//},
		},
	})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating eventRule",
			"Could not create eventRule, unexpected error: "+err.Error(),
		)
		return
	}

	var result = EventRule{
		ID:          types.String{Value: eventRuleResp.ID},
		Enabled:     types.Bool{Value: eventRuleResp.Enabled},
		Name:        types.String{Value: eventRuleResp.Name},
		Description: types.String{Value: eventRuleResp.Description},
		Code:        types.String{Value: eventRuleResp.Code},
		Type:        types.String{Value: eventRuleResp.Type},
		Rule:        types.String{Value: eventRuleResp.Rule},
		Grouping:    types.String{Value: eventRuleResp.Grouping},
		GroupingL2:  types.String{Value: eventRuleResp.GroupingL2},
		GroupingL3:  types.String{Value: eventRuleResp.GroupingL3},
    //EventTags: types.List{},
		//EventTags: types.List{
		//	Elems: make([]attr.Value, len(eventRuleResp.EventTags)),
		//},
	}

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information
func (r resourceEventRule) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var eventRuleId string
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), &eventRuleId)...)
	eventRuleResp, err := r.p.client.GetEventRule(uptycs.EventRule{
		ID: eventRuleId,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading order",
			"Could not get eventRule with ID  "+eventRuleId+": "+err.Error(),
		)
		return
	}

	var result = EventRule{
		ID:          types.String{Value: eventRuleResp.ID},
		Enabled:     types.Bool{Value: eventRuleResp.Enabled},
		Name:        types.String{Value: eventRuleResp.Name},
		Description: types.String{Value: eventRuleResp.Description},
		Code:        types.String{Value: eventRuleResp.Code},
		Type:        types.String{Value: eventRuleResp.Type},
		Rule:        types.String{Value: eventRuleResp.Rule},
		Grouping:    types.String{Value: eventRuleResp.Grouping},
		GroupingL2:  types.String{Value: eventRuleResp.GroupingL2},
		GroupingL3:  types.String{Value: eventRuleResp.GroupingL3},
    //EventTags: types.List{},
		EventTags: types.List{
			Elems: make([]attr.Value, len(eventRuleResp.EventTags)),
		},
	}

	diags := resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

// Update resource
func (r resourceEventRule) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var state EventRule
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	eventRuleID := state.ID.Value

	// Retrieve values from plan
	var plan EventRule
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var tags []string
	plan.EventTags.ElementsAs(ctx, &tags, false)

	eventRuleResp, err := r.p.client.UpdateEventRule(uptycs.EventRule{
		ID:          eventRuleID,
		Name:        plan.Name.Value,
		Code:        plan.Code.Value,
		Custom:      true,
		Description: plan.Description.Value,
		Rule:        plan.Rule.Value,
		Type:        plan.Type.Value,
		Enabled:     plan.Enabled.Value,
		Grouping:    plan.Grouping.Value,
		GroupingL2:  plan.GroupingL2.Value,
		GroupingL3:  plan.GroupingL3.Value,
		EventTags:   tags,
		BuilderConfig: uptycs.BuilderConfig{
			TableName:       plan.BuilderConfig.TableName.Value,
			Added:           plan.BuilderConfig.Added.Value,
			MatchesFilter:   plan.BuilderConfig.MatchesFilter.Value,
			Severity:        plan.BuilderConfig.Severity.Value,
			Key:             plan.BuilderConfig.Key.Value,
			ValueField:      plan.BuilderConfig.ValueField.Value,
			//AutoAlertConfig: uptycs.AutoAlertConfig{},
		},
	})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating eventRule",
			"Could not create eventRule, unexpected error: "+err.Error(),
		)
		return
	}

	var result = EventRule{
		ID:          types.String{Value: eventRuleResp.ID},
		Enabled:     types.Bool{Value: eventRuleResp.Enabled},
		Name:        types.String{Value: eventRuleResp.Name},
		Description: types.String{Value: eventRuleResp.Description},
		Code:        types.String{Value: eventRuleResp.Code},
		Type:        types.String{Value: eventRuleResp.Type},
		Rule:        types.String{Value: eventRuleResp.Rule},
		Grouping:    types.String{Value: eventRuleResp.Grouping},
		GroupingL2:  types.String{Value: eventRuleResp.GroupingL2},
		GroupingL3:  types.String{Value: eventRuleResp.GroupingL3},
		EventTags: types.List{
			Elems: make([]attr.Value, len(eventRuleResp.EventTags)),
		},
		BuilderConfig: BuilderConfig{
			TableName:       types.String{Value: eventRuleResp.BuilderConfig.TableName},
			Added:           types.Bool{Value: eventRuleResp.BuilderConfig.Added},
			MatchesFilter:   types.Bool{Value: eventRuleResp.BuilderConfig.MatchesFilter},
			Severity:        types.String{Value: eventRuleResp.BuilderConfig.Severity},
			Key:             types.String{Value: eventRuleResp.BuilderConfig.Key},
			ValueField:      types.String{Value: eventRuleResp.BuilderConfig.ValueField},
			//AutoAlertConfig: AutoAlertConfig{},
		},
	}

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete resource
func (r resourceEventRule) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state EventRule
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	eventRuleID := state.ID.Value

	_, err := r.p.client.DeleteEventRule(uptycs.EventRule{
		ID: eventRuleID,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting order",
			"Could not delete eventRule with ID  "+eventRuleID+": "+err.Error(),
		)
		return
	}

	// Remove resource from state
	resp.State.RemoveResource(ctx)
}

// Import resource
func (r resourceEventRule) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
