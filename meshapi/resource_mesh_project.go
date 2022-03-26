package meshapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMeshProjectSchema() *schema.Resource {
	return &schema.Resource{
		Read:   resourceMeshProjectRead,
		Create: resourceMeshProjectCreate,
		Update: resourceMeshProjectUpdate,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"name": {
				Required: true,
				ForceNew: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"customer_id": {
				Required: true,
				ForceNew: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Required: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"tags": {
				Optional: true,
				Type:     schema.TypeString,
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

func resourceMeshProjectCreate(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshobjects.v1+json")
	resourceHeaders.Set("Content-Type", "application/vnd.meshcloud.api.meshobjects.v1+json;charset=UTF-8")

	resourceName := d.Get("name").(string)
	resourceDisplayName := d.Get("display_name").(string)
	resourceCustomerId := d.Get("customer_id").(string)
	resourceTags := d.Get("tags").(string)

	data := fmt.Sprintf(`{"apiVersion":"v1","kind":"meshProject","metadata":{"name":"%s","ownedByCustomer":"%s"},"spec":{"displayName":"%s","tags":%s}}`, resourceName, resourceCustomerId, resourceDisplayName, resourceTags)

	log.Printf("[DEBUG] MeshProject Create: %s", data)
	response, err := client.executePutAPI(client.BaseUrl.String(), string(data), resourceHeaders)

	log.Printf("[DEBUG] MeshProject Execute PutAPI Response: %s", response)

	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error creating MeshProject: %s", err)
	}

	d.SetId(resourceName)
	d.Set("name", resourceName)
	d.Set("display_name", resourceDisplayName)
	d.Set("customer_id", resourceCustomerId)
	d.Set("tags", resourceTags)
	return
}

func resourceMeshProjectUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshobjects.v1+json")
	resourceHeaders.Set("Content-Type", "application/vnd.meshcloud.api.meshobjects.v1+json;charset=UTF-8")

	resourceName := d.Get("name").(string)
	resourceDisplayName := d.Get("display_name").(string)
	resourceCustomerId := d.Get("customer_id").(string)
	resourceTags := d.Get("tags").(string)

	data := fmt.Sprintf(`{"apiVersion":"v1","kind":"meshProject","metadata":{"name":"%s","ownedByCustomer":"%s"},"spec":{"displayName":"%s","tags":%s}}`, resourceName, resourceCustomerId, resourceDisplayName, resourceTags)

	log.Printf("[DEBUG] MeshProject Create: %s", data)
	response, err := client.executePutAPI(client.BaseUrl.String(), string(data), resourceHeaders)

	log.Printf("[DEBUG] MeshProject Execute PutAPI Response: %s", response)

	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error creating MeshProject: %s", err)
	}

	d.SetId(resourceName)
	d.Set("name", resourceName)
	d.Set("display_name", resourceDisplayName)
	d.Set("customer_id", resourceCustomerId)
	d.Set("tags", resourceTags)
	return

}
