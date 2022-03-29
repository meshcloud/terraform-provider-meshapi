package meshapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMeshUserSchema() *schema.Resource {
	return &schema.Resource{
		Read:   commonMeshUserRead,
		Create: resourceMeshUserCreateAndUpdate,
		Update: resourceMeshUserCreateAndUpdate,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"name": {
				Required: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"email": {
				Required: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"euid": {
				Required: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"first_name": {
				Required: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"last_name": {
				Required: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"tags": {
				Optional: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
				Default:  "{}",
			},
		},
	}
}

func resourceMeshUserCreateAndUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshobjects.v1+json")
	resourceHeaders.Set("Content-Type", "application/vnd.meshcloud.api.meshobjects.v1+json;charset=UTF-8")

	resourceName := d.Get("name").(string)
	resourceEmail := d.Get("email").(string)
	resourceEuid := d.Get("euid").(string)
	resourceFirstName := d.Get("first_name").(string)
	resourceLastName := d.Get("last_name").(string)
	resourceTags := d.Get("tags").(string)

	data := fmt.Sprintf(`{"apiVersion":"v1","kind":"meshUser","metadata":{"name":"%s"},"spec":{"email":"%s","firstName":"%s","lastName":"%s","euid":"%s","tags":%s}}`, resourceName, resourceEmail, resourceFirstName, resourceLastName, resourceEuid, resourceTags)

	log.Printf("[DEBUG] MeshUser Create: %s", data)
	response, err := client.executePutAPI(client.BaseUrl.String(), string(data), resourceHeaders)
	log.Printf("[DEBUG] MeshUser Execute PutAPI Response: %s", response)

	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error creating MeshUser: %s", err)
	}

	d.SetId(resourceName)
	d.Set("name", resourceName)
	d.Set("email", resourceEmail)
	d.Set("euid", resourceEuid)
	d.Set("first_name", resourceFirstName)
	d.Set("last_name", resourceLastName)
	d.Set("tags", resourceTags)
	return
}
