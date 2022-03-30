package meshapi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMeshTenantSchema() *schema.Resource {
	return &schema.Resource{
		Description: "meshTenant Datasource",
		Read:        commonMeshTenantRead,

		Schema: map[string]*schema.Schema{
			"customer_id": {
				Required: true,
				Type:     schema.TypeString,
			},
			"project_id": {
				Required: true,
				Type:     schema.TypeString,
			},
			"platform_id": {
				Required: true,
				Type:     schema.TypeString,
			},
			"tenant_id": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"landing_zone_id": {
				Computed: true,
				Type:     schema.TypeString,
			},
		},
	}
}
