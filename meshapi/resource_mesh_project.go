package meshapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceMeshProjectSchema() *schema.Resource {
	return &schema.Resource{
		Read:   resourceMeshProjectRead,
		Create: resourceMeshProjectCreate,
		Update: resourceMeshProjectUpdate,
		Delete: resourceMeshProjectDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Elem:     schema.TypeString,
			},
			"customer_id": {
				Type:     schema.TypeString,
				Required: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				Elem:     schema.TypeString,
			},
			"tags": {
				Type:     schema.TypeString,
				Optional: true,
				Elem:     schema.TypeString,
			},
		},
	}
}

type MeshProject struct {
	name         string
	display_name string
}

func resourceMeshProjectRead(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshproject.v1.hal+json")
	resourceName := d.Get("name").(string)
	if resourceName == "" {
		return fmt.Errorf("Project name is required!")
	}

	resourceCustomerId := d.Get("customer_id").(string)
	if resourceCustomerId == "" {
		return fmt.Errorf("Customer Id is required!")
	}
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

func resourceMeshProjectCreate(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshobjects.v1+json")
	resourceHeaders.Set("Content-Type", "application/vnd.meshcloud.api.meshobjects.v1+json;charset=UTF-8")

	resourceName := d.Get("name").(string)
	if resourceName == "" {
		return fmt.Errorf("Project name is required!")
	}

	resourceDisplayName := d.Get("display_name").(string)
	if resourceDisplayName == "" {
		return fmt.Errorf("Project display_name is required!")
	}

	resourceCustomerId := d.Get("customer_id").(string)
	if resourceCustomerId == "" {
		return fmt.Errorf("Customer Id is required!")
	}

	data := fmt.Sprintf(`{
		"apiVersion": "v1",
		"kind": "meshProject",
		"metadata": {
		  "name": "%s",
		  "ownedByCustomer": "%s"
		},
		"spec": {
		  "displayName": "%s",
		  "tags": {
			"Environment": [ "Playground" ],
			"AccessLevel: [ "internal" ],
			"ProjectContact": "dakinci@meshcloud.io"
		  }
		}
	  }`, resourceName, resourceCustomerId, resourceDisplayName)

	b, err := client.executePutAPI(client.BaseUrl.String(), string(data), resourceHeaders)
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

func resourceMeshProjectUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshobjects.v1+json")
	resourceHeaders.Set("Content-Type", "application/vnd.meshcloud.api.meshobjects.v1+json;charset=UTF-8")

	resourceName := d.Get("name").(string)
	if resourceName == "" {
		return fmt.Errorf("Project name is required!")
	}
	resourceDisplayName := d.Get("display_name").(string)
	if resourceDisplayName == "" {
		return fmt.Errorf("Project display_name is required!")
	}

	m := MeshProject{resourceName, resourceDisplayName}
	data, err := json.Marshal(m)

	b, err := client.executePutAPI(client.BaseUrl.String(), string(data), resourceHeaders)
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

func resourceMeshProjectDelete(d *schema.ResourceData, meta interface{}) (err error) {
	return nil
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
	outputs["display_name"] = data["spec"].(map[string]interface{})["displayName"]
	outputs["tags"] = string(tags.([]byte))
	return
}
