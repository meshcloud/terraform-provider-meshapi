package meshapi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMeshCustomerSchema() *schema.Resource {
	return &schema.Resource{
		Description: "meshCustomer Datasource",
		Read:        commonMeshCustomerRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Required: true,
				Type:     schema.TypeString,
			},
			"display_name": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"tags": {
				Computed: true,
				Type:     schema.TypeString,
			},
		},
	}
}
