package yodeck

import (
	"fmt"
	"log"

	"github.com/easylo/go-yodeck/yodeck"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceYodeckWorkspace() *schema.Resource {
	return &schema.Resource{
		Create: resourceYodeckWorkspaceCreate,
		Read:   resourceYodeckWorkspaceRead,
		Update: resourceYodeckWorkspaceUpdate,
		Delete: resourceYodeckWorkspaceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceYodeckWorkspaceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	workspace := buildWorkspaceStruct(d)

	log.Printf("[INFO] Creating Yodeck Workspace %s", workspace.Name)

	Workspace, _, err := client.Workspace.Create(workspace)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprint(Workspace.ID))

	return nil
}

func resourceYodeckWorkspaceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)
	log.Printf("[INFO] Reading PagerDuty user %s", d.Id())

	Workspace, _, err := client.Workspace.Get(d.Id())
	if err != nil {
		return handleNotFoundError(err, d)
	}

	d.Set("name", Workspace.Name)
	d.Set("description", Workspace.Description)

	return nil
}

func resourceYodeckWorkspaceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	Workspace := buildWorkspaceStruct(d)

	log.Printf("[INFO] Updating Yodeck Workspace %s", d.Id())

	if _, _, err := client.Workspace.Update(d.Id(), Workspace); err != nil {
		return err
	}
	return nil
}

func resourceYodeckWorkspaceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	log.Printf("[INFO] Deleting Yodeck Workspace %s", d.Id())

	if _, err := client.Workspace.Delete(d.Id()); err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func buildWorkspaceStruct(d *schema.ResourceData) *yodeck.Workspace {
	workspace := &yodeck.Workspace{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	return workspace
}
