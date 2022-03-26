package meshapi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// marshalData is used to ensure the data is put into a format Terraform can output
func marshalData(d *schema.ResourceData, vals map[string]interface{}) {
	for k, v := range vals {
		if k == "id" {
			d.SetId(v.(string))
		} else {
			str, ok := v.(string)
			if ok {
				d.Set(k, str)
			} else {
				d.Set(k, v)
			}
		}
	}
}
