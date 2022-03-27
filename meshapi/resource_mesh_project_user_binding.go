package meshapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMeshProjectUserBindingSchema() *schema.Resource {
	return &schema.Resource{
		Read:   resourceMeshProjectUserBindingRead,
		Create: resourceMeshProjectUserBindingCreateAndUpdate,
		Update: resourceMeshProjectUserBindingCreateAndUpdate,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"role_name": {
				Required: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"customer_id": {
				Required: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"project_id": {
				Required: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"user_id": {
				Required: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
		},
	}
}

func resourceMeshProjectUserBindingRead(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	resourceRoleName := d.Get("role_name").(string)
	resourceCustomerId := d.Get("customer_id").(string)
	resourceProjectId := d.Get("project_id").(string)
	resourceUserId := d.Get("user_id").(string)

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshuser.v1.hal+json")

	resourceQueries := make(map[string]string)
	resourceQueries["customerIdentifier"] = resourceCustomerId
	resourceQueries["projectIdentifier"] = resourceProjectId
	resourceQueries["projectRole"] = resourceRoleName

	response, err := client.executeGetAPI(client.BaseUrl.String(), "api/meshobjects/meshusers", "", resourceHeaders, resourceQueries)
	if err != nil {
		return
	}

	var data map[string]interface{}

	err = json.Unmarshal(response, &data)
	if err != nil {
		err = fmt.Errorf("Cannot unmarshal json of API response: %v", err)
		return
	}

	// Check if the user exists in the data
	var userExists bool = false
	for _, user := range data["_embedded"].(map[string]interface{})["meshUsers"].([]interface{}) {
		if user.(map[string]interface{})["metadata"].(map[string]interface{})["name"].(string) == resourceUserId {
			userExists = true
			break
		}
	}

	if userExists {
		d.SetId(resourceCustomerId + "/" + resourceProjectId + "/" + resourceUserId)
		d.Set("user_id", resourceUserId)
		d.Set("customer_id", resourceCustomerId)
		d.Set("project_id", resourceProjectId)
		d.Set("role_name", resourceRoleName)
	} else {
		d.SetId("")
	}

	return
}

func resourceMeshProjectUserBindingCreateAndUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshobjects.v1+json")
	resourceHeaders.Set("Content-Type", "application/vnd.meshcloud.api.meshobjects.v1+json;charset=UTF-8")

	resourceRoleName := d.Get("role_name").(string)
	resourceCustomerId := d.Get("customer_id").(string)
	resourceProjectId := d.Get("project_id").(string)
	resourceUserId := d.Get("user_id").(string)

	data := fmt.Sprintf(`{"apiVersion":"v1","kind":"meshProjectUserBinding","roleRef":{"name":"%s"},"targetRef":{"name":"%s","ownedByCustomer":"%s"},"subjects":[{"name":"%s"}]}`, resourceRoleName, resourceProjectId, resourceCustomerId, resourceUserId)

	log.Printf("[DEBUG] MeshProjectUserBinding Create: %s", data)
	response, err := client.executePutAPI(client.BaseUrl.String(), string(data), resourceHeaders)
	log.Printf("[DEBUG] MeshProjectUserBinding Execute PutAPI Response: %s", response)

	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error creating MeshProjectUserBinding: %s", err)
	}

	d.SetId(resourceCustomerId + "/" + resourceProjectId + "/" + resourceUserId)
	d.Set("role_name", resourceRoleName)
	d.Set("customer_id", resourceCustomerId)
	d.Set("project_id", resourceProjectId)
	d.Set("user_id", resourceUserId)
	return
}
