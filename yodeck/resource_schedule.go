package yodeck

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceYodeckSchedule() *schema.Resource {
	return &schema.Resource{
		Create: resourceYodeckScheduleCreate,
		Read:   resourceYodeckScheduleRead,
		Update: resourceYodeckScheduleUpdate,
		Delete: resourceYodeckScheduleDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceYodeckScheduleCreate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckScheduleRead(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckScheduleUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckScheduleDelete(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}
