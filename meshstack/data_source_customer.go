package meshstack

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func getCustomerSchema() *schema.Resource {
	return &schema.Resource{
		Read: getCustomerDataSourceRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"tags": {
				Type:     schema.TypeString,
				Computed: true,
				Elem:     schema.TypeString,
			},
		},
	}
}

func getCustomerDataSourceRead(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	// Accept = "application/vnd.meshcloud.api.meshcustomer.v1.hal+json"
	// Content-Type = "application/vnd.meshcloud.api.meshobjectcollection.v1+json;charset=UTF-8"

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshcustomer.v1.hal+json")
	resourceHeaders.Set("Content-Type", "application/vnd.meshcloud.api.meshobjectcollection.v1+json;charset=UTF-8")

	resourceName := d.Get("name").(string)
	if resourceName == "" {
		return fmt.Errorf("Customer name is required!")
	}
	b, err := client.executeGetAPI(client.BaseUrl.String(), "api/meshobjects/meshcustomers", resourceName, resourceHeaders)
	if err != nil {
		return
	}
	outputs, err := flattenCustomerResponse(b)
	if err != nil {
		return
	}
	marshalData(d, outputs)

	return
}

func flattenCustomerResponse(b []byte) (outputs map[string]interface{}, err error) {
	var data map[string]interface{}
	var tags interface{}

	err = json.Unmarshal(b, &data)
	if err != nil {
		err = fmt.Errorf("Cannot unmarshal json of API response: %v", err)
		return
	}

	if data["spec"].(map[string]interface{})["tags"] != nil {
		tags, err = json.Marshal(data["spec"].(map[string]interface{})["tags"])
		if err != nil {
			err = fmt.Errorf("Cannot marshal tags: %v", err)
			return
		}
	}

	outputs = make(map[string]interface{})
	outputs["id"] = data["metadata"].(map[string]interface{})["name"]
	outputs["name"] = data["metadata"].(map[string]interface{})["name"]
	outputs["display_name"] = data["spec"].(map[string]interface{})["displayName"]
	outputs["tags"] = string(tags.([]byte))
	return
}
