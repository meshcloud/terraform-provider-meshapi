package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/meshcloud/terraform-provider-meshapi/meshapi"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: meshapi.Provider})
}
