package main

import (
	"github.com/DorukAkinci/terraform-provider-meshstack/meshstack"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: meshstack.Provider})
}
