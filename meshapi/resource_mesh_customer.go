package meshapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMeshCustomerSchema() *schema.Resource {
	return &schema.Resource{
		Read:   commonMeshCustomerRead,
		Create: resourceMeshCustomerCreateAndUpdate,
		Update: resourceMeshCustomerCreateAndUpdate,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"name": {
				Required: true,
				ForceNew: true,
				Type:     schema.TypeString,
			},
			"display_name": {
				Required: true,
				Type:     schema.TypeString,
			},
			"tags": {
				Optional: true,
				Type:     schema.TypeString,
				Default:  "{}",
			},
		},
	}
}

func resourceMeshCustomerCreateAndUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshobjects.v1+json")
	resourceHeaders.Set("Content-Type", "application/vnd.meshcloud.api.meshobjects.v1+json;charset=UTF-8")

	resourceName := d.Get("name").(string)
	resourceDisplayName := d.Get("display_name").(string)
	resourceTags := d.Get("tags").(string)

	data := fmt.Sprintf(`{"apiVersion":"v1","kind":"meshCustomer","metadata":{"name":"%s"},"spec":{"displayName":"%s","tags":%s}}`, resourceName, resourceDisplayName, resourceTags)

	log.Printf("[DEBUG] MeshCustomer Create: %s", data)
	response, err := client.executePutAPI(client.BaseUrl.String(), string(data), resourceHeaders)
	log.Printf("[DEBUG] MeshCustomer Execute PutAPI Response: %s", response)

	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error creating MeshCustomer: %s", err)
	}

	d.SetId(resourceName)
	d.Set("name", resourceName)
	d.Set("display_name", resourceDisplayName)
	d.Set("tags", resourceTags)
	return
}
