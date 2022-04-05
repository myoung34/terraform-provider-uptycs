package uptycs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAlertRuleType() *schema.Resource {
	return &schema.Resource{
		Description: "uptycs alert rule resource.",

		CreateContext: resourceAlertRuleCreate,
		ReadContext:   resourceAlertRuleRead,
		DeleteContext: resourceAlertRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Computed:            true,
				Type: schema.TypeString,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"code": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
			},
			"grouping": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"grouping_l2": {
				Type:     schema.TypeString,
				Required: true,
			},
			"grouping_l3": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sql_config": {
				Type:        schema.TypeList,
				Elem:        schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval_seconds": {
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
				Computed:    true,
			},
		},
	}
}


// Create a new resource
func resourceAlertRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceAlertRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceAlertRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
