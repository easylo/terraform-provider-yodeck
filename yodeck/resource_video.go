package yodeck

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceYodeckVideo() *schema.Resource {
	return &schema.Resource{
		Create: resourceYodeckVideoCreate,
		Read:   resourceYodeckVideoRead,
		Update: resourceYodeckVideoUpdate,
		Delete: resourceYodeckVideoDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceYodeckVideoCreate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckVideoRead(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckVideoUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckVideoDelete(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}
