package meshapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func commonMeshTenantRead(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshtenant.v2.hal+json")

	resourceCustomerId := d.Get("customer_id").(string)
	resourceProjectId := d.Get("project_id").(string)
	resourcePlatformId := d.Get("platform_id").(string)

	b, err := client.executeGetAPI(client.BaseUrl.String(), "api/meshobjects/meshtenants", fmt.Sprintf("%s.%s.%s", resourceCustomerId, resourceProjectId, resourcePlatformId), resourceHeaders, nil)
	if err != nil {
		return
	}
	outputs, err := flattenResourceMeshTenantResponse(b)
	if err != nil {
		return
	}
	marshalData(d, outputs)

	return
}

func flattenResourceMeshTenantResponse(b []byte) (outputs map[string]interface{}, err error) {
	var data map[string]interface{}

	err = json.Unmarshal(b, &data)
	if err != nil {
		err = fmt.Errorf("Cannot unmarshal json of API response: %v", err)
		return
	}

	outputs = make(map[string]interface{})
	outputs["id"] = data["spec"].(map[string]interface{})["localId"]
	outputs["tenant_id"] = data["spec"].(map[string]interface{})["localId"]
	outputs["landing_zone_id"] = data["spec"].(map[string]interface{})["landingZoneIdentifier"]
	outputs["customer_id"] = data["metadata"].(map[string]interface{})["ownedByCustomer"]
	outputs["project_id"] = data["metadata"].(map[string]interface{})["ownedByProject"]
	outputs["platform_id"] = data["metadata"].(map[string]interface{})["platformIdentifier"]
	return
}
