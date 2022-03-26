package meshapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMeshProjectSchema() *schema.Resource {
	return &schema.Resource{
		Read:   commonMeshProjectRead,
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
