package meshapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func commonMeshProjectRead(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshproject.v1.hal+json")

	resourceName := d.Get("name").(string)
	resourceCustomerId := d.Get("customer_id").(string)

	b, err := client.executeGetAPI(client.BaseUrl.String(), "api/meshobjects/meshprojects", fmt.Sprintf("%s.%s", resourceCustomerId, resourceName), resourceHeaders)
	if err != nil {
		return
	}
	outputs, err := flattenResourceMeshProjectResponse(b)
	if err != nil {
		return
	}
	marshalData(d, outputs)

	return
}

func flattenResourceMeshProjectResponse(b []byte) (outputs map[string]interface{}, err error) {
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
	outputs["customer_id"] = data["metadata"].(map[string]interface{})["ownedByCustomer"]
	outputs["display_name"] = data["spec"].(map[string]interface{})["displayName"]
	outputs["tags"] = string(tags.([]byte))
	return
}
