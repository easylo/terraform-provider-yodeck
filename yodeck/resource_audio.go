package yodeck

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceYodeckAudio() *schema.Resource {
	return &schema.Resource{
		Create: resourceYodeckAudioCreate,
		Read:   resourceYodeckAudioRead,
		Update: resourceYodeckAudioUpdate,
		Delete: resourceYodeckAudioDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceYodeckAudioCreate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckAudioRead(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckAudioUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}

func resourceYodeckAudioDelete(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*yodeck.Client)
	return nil
}
