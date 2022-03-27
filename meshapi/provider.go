package meshapi

import (
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: providerConfigure,

		Schema: map[string]*schema.Schema{
			"url": {
				Required: true,
				Type:     schema.TypeString,
			},

			"headers": {
				Optional: true,
				Type:     schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"meshapi_mesh_customer": dataSourceMeshCustomerSchema(),
			"meshapi_mesh_project":  dataSourceMeshProjectSchema(),
			"meshapi_mesh_tenant":   dataSourceMeshTenantSchema(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"meshapi_mesh_project":               resourceMeshProjectSchema(),
			"meshapi_mesh_customer":              resourceMeshCustomerSchema(),
			"meshapi_mesh_customer_user_binding": resourceMeshCustomerUserBindingSchema(),
			"meshapi_mesh_project_user_binding":  resourceMeshProjectUserBindingSchema(),
			"meshapi_mesh_tenant":                resourceMeshTenantSchema(),
		},
	}
}

// ProviderClient holds metadata / config for use by Terraform resources
type ProviderClient struct {
	Url    string
	Client *Client
}

// newProviderClient is a factory for creating ProviderClient structs
func newProviderClient(url string, headers http.Header) (ProviderClient, error) {
	p := ProviderClient{
		Url: url,
	}
	p.Client = NewClient(url, 443, headers)

	return p, nil
}

// providerConfigure parses the config into the Terraform provider meta object
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("url").(string)
	if url == "" {
		log.Println("Set Url to the localhost.")
		url = "localhost"
	}

	h := make(http.Header)

	headers, exists := d.GetOk("headers")
	if exists {
		for k, v := range headers.(map[string]interface{}) {
			h.Set(k, v.(string))
		}
	}

	return newProviderClient(url, h)
}
