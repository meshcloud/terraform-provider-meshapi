package meshapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMeshTenantSchema() *schema.Resource {
	return &schema.Resource{
		Read:   commonMeshTenantRead,
		Create: resourceMeshTenantCreateAndUpdate,
		Update: resourceMeshTenantCreateAndUpdate,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
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
			"platform_id": {
				Required: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"tenant_id": {
				Optional: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
			"landing_zone_id": {
				Required: true,
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
			},
		},
	}
}

func resourceMeshTenantCreateAndUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	provider := meta.(ProviderClient)
	client := provider.Client

	resourceHeaders := make(http.Header)
	resourceHeaders.Set("Accept", "application/vnd.meshcloud.api.meshobjects.v1+json")
	resourceHeaders.Set("Content-Type", "application/vnd.meshcloud.api.meshobjects.v1+json;charset=UTF-8")

	resourceCustomerId := d.Get("customer_id").(string)
	resourceProjectId := d.Get("project_id").(string)
	resourcePlatformId := d.Get("platform_id").(string)
	resourceTenantId := d.Get("tenant_id").(string)
	resourceLandingZoneId := d.Get("landing_zone_id").(string)

	// This part must be rewritten. Currently we do not support quota parameter and landing zone parameter is required.
	// TODO: Check, is it possible to send empty string values for optional parameters.
	data := fmt.Sprintf(`{"apiVersion":"v2","kind":"meshTenant","metadata":{"ownedByProject":"%s","ownedByCustomer":"%s","platformIdentifier":"%s"},"spec":{"landingZoneIdentifier":"%s"%s}}`, resourceProjectId, resourceCustomerId, resourcePlatformId, resourceLandingZoneId, `,"localId":"`+resourceTenantId+`"`)

	log.Printf("[DEBUG] MeshTenant Create: %s", data)
	response, err := client.executePutAPI(client.BaseUrl.String(), string(data), resourceHeaders)
	log.Printf("[DEBUG] MeshTenant Execute PutAPI Response: %s", response)

	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error creating MeshTenant: %s", err)
	}

	d.SetId(resourceTenantId)
	d.Set("tenant_id", resourceTenantId)
	d.Set("landing_zone_id", resourceLandingZoneId)
	d.Set("customer_id", resourceCustomerId)
	d.Set("project_id", resourceProjectId)
	d.Set("platform_id", resourcePlatformId)
	return
}
