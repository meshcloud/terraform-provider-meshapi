package main

import (
	"github.com/DorukAkinci/terraform-provider-meshapi/meshapi"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: meshapi.Provider})
}
