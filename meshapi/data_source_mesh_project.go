package meshapi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMeshProjectSchema() *schema.Resource {
	return &schema.Resource{
		Description: "meshProject Datasource",
		Read:        commonMeshProjectRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Required: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"customer_id": {
				Required: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Computed: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"tags": {
				Computed: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
		},
	}
}
