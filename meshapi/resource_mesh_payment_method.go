package meshapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMeshPaymentMethodSchema() *schema.Resource {
	return &schema.Resource{
		Read:   commonMeshPaymentMethodRead,
		Create: resourceMeshPaymentMethodCreateAndUpdate,
		Update: resourceMeshPaymentMethodCreateAndUpdate,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"name": {
				Required: true,
				ForceNew: true,
				Type:     schema.TypeString,
			},
			"customer_id": {
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
			"amount": {
				Optional: true,
				Type:     schema.TypeInt,
			},
			"expiration_date": {
				Optional: true,
				Type:     schema.TypeString,
			},
		},
	}
}

func resourceMeshPaymentMethodCreateAndUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshobjects.v1+json")
	resourceHeaders.Set("Content-Type", "application/vnd.meshcloud.api.meshobjects.v1+json;charset=UTF-8")

	resourceName := d.Get("name").(string)
	resourceDisplayName := d.Get("display_name").(string)
	resourceCustomerId := d.Get("customer_id").(string)
	resourceTags := d.Get("tags").(string)
	resourceAmount := d.Get("amount").(int)
	resourceExpirationDate := d.Get("expiration_date").(string)

	// ADD EXPIRATION DATE
	data := fmt.Sprintf(`{"apiVersion":"v1","kind":"meshPaymentMethod","metadata":{"name":"%s","ownedByCustomer":"%s"},"spec":{"displayName":"%s","amount":%d,"tags":%s}}`, resourceName, resourceCustomerId, resourceDisplayName, resourceAmount, resourceTags)

	log.Printf("[DEBUG] MeshPaymentMethod Create: %s", data)
	response, err := client.executePutAPI(client.BaseUrl.String(), string(data), resourceHeaders)
	log.Printf("[DEBUG] MeshPaymentMethod Execute PutAPI Response: %s", response)

	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error creating MeshPaymentMethod: %s", err)
	}

	d.SetId(resourceName)
	d.Set("name", resourceName)
	d.Set("display_name", resourceDisplayName)
	d.Set("customer_id", resourceCustomerId)
	d.Set("tags", resourceTags)
	d.Set("amount", resourceAmount)
	d.Set("expiration_date", resourceExpirationDate)
	return
}
