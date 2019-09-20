package yodeck

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceYodeckWidget() *schema.Resource {
	return &schema.Resource{
		Create: resourceYodeckWidgetCreate,
		Read:   resourceYodeckWidgetRead,
		Update: resourceYodeckWidgetUpdate,
		Delete: resourceYodeckWidgetDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceYodeckWidgetCreate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckWidgetRead(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckWidgetUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckWidgetDelete(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}
