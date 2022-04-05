package main

import (
	"terraform-provider-uptycs/uptycs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: uptycs.New,
	})
}
