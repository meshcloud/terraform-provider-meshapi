package meshapi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMeshPaymentMethodSchema() *schema.Resource {
	return &schema.Resource{
		Description: "meshPaymentMethod Datasource",
		Read:        commonMeshPaymentMethodRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Required: true,
				Type:     schema.TypeString,
			},
			"customer_id": {
				Computed: true,
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
			"amount": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"expiration_date": {
				Computed: true,
				Type:     schema.TypeString,
			},
		},
	}
}
