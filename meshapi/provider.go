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
			"api_version": {
				Optional: true,
				Type:     schema.TypeString,
				Default:  "v1",
			},

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
		},

		ResourcesMap: map[string]*schema.Resource{
			"meshapi_mesh_project": resourceMeshProjectSchema(),
		},
	}
}

// ProviderClient holds metadata / config for use by Terraform resources
type ProviderClient struct {
	ApiVersion string
	Hostname   string
	Client     *Client
}

// newProviderClient is a factory for creating ProviderClient structs
func newProviderClient(apiVersion, hostname string, headers http.Header) (ProviderClient, error) {
	p := ProviderClient{
		ApiVersion: apiVersion,
		Hostname:   hostname,
	}
	p.Client = NewClient(headers, 443, hostname, apiVersion)

	return p, nil
}

// providerConfigure parses the config into the Terraform provider meta object
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	apiVersion := d.Get("api_version").(string)
	if apiVersion == "" {
		log.Println("Defaulting environment in URL config to use API default version...")
	}

	hostname := d.Get("url").(string)
	if hostname == "" {
		log.Println("Defaulting environment in URL config to use API default hostname...")
		hostname = "localhost"
	}

	h := make(http.Header)

	headers, exists := d.GetOk("headers")
	if exists {
		for k, v := range headers.(map[string]interface{}) {
			h.Set(k, v.(string))
		}
	}

	return newProviderClient(apiVersion, hostname, h)
}
