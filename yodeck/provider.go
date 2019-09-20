package yodeck

import (
	"fmt"
	"log"

	"github.com/easylo/go-yodeck/yodeck"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider expose supported resource
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("YODECK_USERNAME", nil),
				Description: "Username for YODECK Account.",
			},
			"apikey": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("YODECK_APIKEY", nil),
				Description: "API Key for YODECK",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"yodeck_webpage":   resourceYodeckWebpage(),
			"yodeck_workspace": resourceYodeckWorkspace(),
			"yodeck_playlist":  resourceYodeckPlaylist(),
			"yodeck_show":      resourceYodeckShow(),
			"yodeck_device":    resourceYodeckDevice(),
			// "yodeck_widget":    resourceYodeckWidget(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &yodeck.Config{
		Username: d.Get("username").(string),
		Apikey:   d.Get("apikey").(string),
	}

	return yodeck.NewClient(config)
}

func isErrCode(err error, code int) bool {
	if e, ok := err.(*yodeck.Error); ok && e.ErrorResponse.StatusCode == code {
		return true
	}

	return false
}

func handleNotFoundError(err error, d *schema.ResourceData) error {
	if isErrCode(err, 404) {
		log.Printf("[WARN] Removing %s because it's gone", d.Id())
		d.SetId("")
		return nil
	}

	return fmt.Errorf("Error reading: %s: %s", d.Id(), err)
}
