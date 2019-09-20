package yodeck

import (
	"fmt"
	"log"

	"github.com/easylo/go-yodeck/yodeck"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceYodeckDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceYodeckDeviceCreate,
		Read:   resourceYodeckDeviceRead,
		Update: resourceYodeckDeviceUpdate,
		Delete: resourceYodeckDeviceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"workspace": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"default_show": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wifi_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wifi_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wifi_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceYodeckDeviceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	device := buildDeviceStruct(d)

	log.Printf("[INFO] Creating Yodeck Device %s", device.Name)

	device, _, err := client.Device.Create(device)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprint(device.ID))

	return nil
}

func resourceYodeckDeviceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)
	log.Printf("[INFO] Reading PagerDuty user %s", d.Id())

	device, _, err := client.Device.Get(d.Id())
	if err != nil {
		return handleNotFoundError(err, d)
	}

	d.Set("name", device.Name)
	d.Set("workspace", device.Workspace)
	d.Set("default_show", device.DefaultShow.SourceID)
	d.Set("wifi_ssid", device.WifiSSID)
	d.Set("wifi_key", device.WifiKey)
	d.Set("wifi_mode", device.WifiMode)

	return nil
}

func resourceYodeckDeviceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	device := buildDeviceStruct(d)

	log.Printf("[INFO] Updating Yodeck Device %s", d.Id())

	if _, _, err := client.Device.Update(d.Id(), device); err != nil {
		return err
	}
	return nil
}

func resourceYodeckDeviceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	log.Printf("[INFO] Deleting Yodeck Device %s", d.Id())

	if _, err := client.Device.Delete(d.Id()); err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func buildDeviceStruct(d *schema.ResourceData) *yodeck.Device {
	device := &yodeck.Device{
		Name: d.Get("name").(string),
	}

	if attr, ok := d.GetOk("workspace"); ok {
		device.Workspace = attr.(int)
	}

	if attr, ok := d.GetOk("default_show"); ok {
		device.DefaultShow.SourceID = attr.(int)
		device.DefaultShow.SourceType = "show"
	}

	if attr, ok := d.GetOk("wifi_ssid"); ok {
		device.WifiSSID = attr.(string)
	}
	if attr, ok := d.GetOk("wifi_key"); ok {
		device.WifiKey = attr.(string)
	}

	if attr, ok := d.GetOk("wifi_mode"); ok {
		device.WifiMode = attr.(string)
	}
	return device
}
