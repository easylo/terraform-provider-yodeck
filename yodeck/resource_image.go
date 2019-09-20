package yodeck

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceYodeckImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceYodeckImageCreate,
		Read:   resourceYodeckImageRead,
		Update: resourceYodeckImageUpdate,
		Delete: resourceYodeckImageDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceYodeckImageCreate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckImageRead(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckImageUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckImageDelete(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}
