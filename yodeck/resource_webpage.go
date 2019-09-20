package yodeck

import (
	"fmt"
	"log"

	"github.com/easylo/go-yodeck/yodeck"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceYodeckWebpage() *schema.Resource {
	return &schema.Resource{
		Create: resourceYodeckWebpageCreate,
		Read:   resourceYodeckWebpageRead,
		Update: resourceYodeckWebpageUpdate,
		Delete: resourceYodeckWebpageDelete,

		Schema: map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"workspace": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceYodeckWebpageCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	webpage := buildWebpageStruct(d)

	log.Printf("[INFO] Creating Yodeck Webpage %s", webpage.Name)

	webpage, _, err := client.Webpage.Create(webpage)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprint(webpage.ID))

	return nil
}

func resourceYodeckWebpageRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	log.Printf("[INFO] Reading YODECK Webpage %s", d.Id())

	webpage, _, err := client.Webpage.Get(d.Id())
	if err != nil {
		return handleNotFoundError(err, d)
	}

	d.Set("name", webpage.Name)
	d.Set("url", webpage.URL)

	d.Set("description", webpage.Description)
	d.Set("duration", webpage.Duration)
	d.Set("workspace", webpage.Workspace)

	return nil
}

func resourceYodeckWebpageUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	webpage := buildWebpageStruct(d)

	log.Printf("[INFO] Updating Yodeck Webpage %s", d.Id())

	if _, _, err := client.Webpage.Update(d.Id(), webpage); err != nil {
		return err
	}
	return nil
}

func resourceYodeckWebpageDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	log.Printf("[INFO] Deleting Yodeck Webpage %s", d.Id())

	if _, err := client.Webpage.Delete(d.Id()); err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func buildWebpageStruct(d *schema.ResourceData) *yodeck.Webpage {
	webpage := &yodeck.Webpage{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		URL:         d.Get("url").(string),
		Duration:    d.Get("duration").(int),
	}

	if attr, ok := d.GetOk("workspace"); ok {
		webpage.Workspace = attr.(int)
	}

	return webpage
}
