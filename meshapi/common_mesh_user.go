package meshapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func commonMeshUserRead(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	resourceUserId := d.Get("name").(string)

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshuser.v1.hal+json")

	// FEAT: we should retrieve all pages with a for-each loop and combine them into one map. this function only works if there are not many users
	response, err := client.executeGetAPI(client.BaseUrl.String(), "api/meshobjects/meshusers", "", resourceHeaders, nil)
	if err != nil {
		return
	}

	outputs, err := flattenResourceMeshUserResponse(resourceUserId, response)
	if err != nil {
		return
	}
	marshalData(d, outputs)

	return
}

func flattenResourceMeshUserResponse(resourceUserId string, b []byte) (outputs map[string]interface{}, err error) {
	var data map[string]interface{}

	err = json.Unmarshal(b, &data)
	if err != nil {
		err = fmt.Errorf("Cannot unmarshal json of API response: %v", err)
		return
	}

	outputs = make(map[string]interface{})

	// Check if the user exists in the data
	var userExists bool = false
	for _, user := range data["_embedded"].(map[string]interface{})["meshUsers"].([]interface{}) {
		if user.(map[string]interface{})["metadata"].(map[string]interface{})["name"].(string) == resourceUserId {
			userExists = true
			outputs["id"] = resourceUserId
			outputs["name"] = resourceUserId
			outputs["email"] = user.(map[string]interface{})["spec"].(map[string]interface{})["email"].(string)
			outputs["first_name"] = user.(map[string]interface{})["spec"].(map[string]interface{})["firstName"].(string)
			outputs["last_name"] = user.(map[string]interface{})["spec"].(map[string]interface{})["lastName"].(string)
			break
		}
	}

	if !userExists {
		outputs["id"] = ""
	}

	return
}
