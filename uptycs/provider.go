package uptycs

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/myoung34/uptycs-client-go/uptycs"
)


type config struct {
	client *uptycs.Client
}

// New returns the provider instance.
func New() *schema.Provider {
	var p *schema.Provider
	p = &schema.Provider{
		Schema: map[string]*schema.Schema{
      "api_key": {
        Description: "uptycs Application Key ID (UPTYCS_API_KEY env)",
        Type:        schema.TypeString,
        Optional:    true,
        Sensitive:   true,
        DefaultFunc: schema.EnvDefaultFunc("UPTYCS_API_KEY", nil),
      },
      "api_secret": {
        Description: "uptycs Application Key secret (UPTYCS_API_SECRET env)",
        Type:        schema.TypeString,
        Optional:    true,
        Sensitive:   true,
        DefaultFunc: schema.EnvDefaultFunc("UPTYCS_API_SECRET", nil),
      },
      "customer_id": {
        Description: "uptycs customer id (UPTYCS_CUSTOMER_ID env)",
        Type:        schema.TypeString,
        Optional:    true,
        Sensitive:   true,
        DefaultFunc: schema.EnvDefaultFunc("UPTYCS_CUSTOMER_ID", nil),
      },
      "host": {
        Description: "uptycs endpoint - production or custom URL (UPTYCS_HOST env)",
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("UPTYCS_HOST", "https://thor.uptycs.io"),
      },
		},

		DataSourcesMap: map[string]*schema.Resource{
      //"uptycs_example":           dataSourceUptycsExample(),
		},
    ResourcesMap: map[string]*schema.Resource{
		  "uptycs_alert_rule": resourceAlertRuleType(),
    },
	}
	p.ConfigureContextFunc = providerConfigure(p)
	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		uptycsClient, _ := uptycs.NewClient(uptycs.UptycsConfig{
			Host:       d.Get("host").(string),
			ApiKey:     d.Get("api_key").(string),
			ApiSecret:  d.Get("api_secret").(string),
			CustomerID: d.Get("customer_id").(string),
		})


		c := &config{
			client: uptycsClient,
		}

		return c, nil
	}
}
