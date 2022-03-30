package meshapi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMeshUserSchema() *schema.Resource {
	return &schema.Resource{
		Description: "meshUser Datasource",
		Read:        commonMeshUserRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Required: true,
				Type:     schema.TypeString,
			},
			"email": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"first_name": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"last_name": {
				Computed: true,
				Type:     schema.TypeString,
			},
		},
	}
}
