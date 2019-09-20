package yodeck

import (
	"fmt"
	"log"

	"github.com/easylo/go-yodeck/yodeck"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceYodeckShow() *schema.Resource {
	return &schema.Resource{
		Create: resourceYodeckShowCreate,
		Read:   resourceYodeckShowRead,
		Update: resourceYodeckShowUpdate,
		Delete: resourceYodeckShowDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"workspace": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"regions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"top": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"left": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"height": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"width": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"fit": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"zindex": {
							Type:     schema.TypeInt,
							Default:  2,
							Optional: true,
						},
						"playlists": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"playlist": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"duration": {
										Type:     schema.TypeInt,
										Default:  -1,
										Optional: true,
									},
									"order": {
										Type:     schema.TypeInt,
										Default:  1,
										Optional: true,
									},
								},
							},
						},
						"enable_transparency": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"is_muted": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"res_width": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"res_height": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"background_audio": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceYodeckShowCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	show := buildShowStruct(d)

	log.Printf("[INFO] Creating Yodeck Show %s", show.Name)

	show, _, err := client.Show.Create(show)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprint(show.ID))

	return nil
}

func resourceYodeckShowRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)
	log.Printf("[INFO] Reading Yodeck show %s", d.Id())

	show, _, err := client.Show.Get(d.Id())
	if err != nil {
		return handleNotFoundError(err, d)
	}

	d.Set("name", show.Name)
	d.Set("workspace", show.Workspace)

	if err := d.Set("regions", flattenShowRegions(show.Regions)); err != nil {
		return fmt.Errorf("error setting ShowRegions: %s", err)
	}

	return nil
}

func resourceYodeckShowUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	show := buildShowStruct(d)

	log.Printf("[INFO] Updating Yodeck Show %s", d.Id())

	if _, _, err := client.Show.Update(d.Id(), show); err != nil {
		return err
	}
	return nil
}

func resourceYodeckShowDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	log.Printf("[INFO] Deleting Yodeck Show %s", d.Id())

	if _, err := client.Show.Delete(d.Id()); err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func buildShowStruct(d *schema.ResourceData) *yodeck.Show {
	show := &yodeck.Show{
		Name: d.Get("name").(string),
	}

	if attr, ok := d.GetOk("workspace"); ok {
		show.Workspace = attr.(int)
	}

	if attr, ok := d.GetOk("regions"); ok {
		regions, err := expandShowRegions(attr)
		if err != nil {
			return nil
		}

		show.Regions = regions
	}
	// if attr, ok := d.GetOk("medias"); ok {
	// 	playlist.Media = castSetToSliceStrings(attr.(*schema.Set).List())
	// }
	// if attr, ok := d.GetOk("color"); ok {
	// 	user.Color = attr.(string)
	// }

	// if attr, ok := d.GetOk("time_zone"); ok {
	// 	user.TimeZone = attr.(string)
	// }

	// if attr, ok := d.GetOk("role"); ok {
	// 	user.Role = attr.(string)
	// }

	// if attr, ok := d.GetOk("job_title"); ok {
	// 	user.JobTitle = attr.(string)
	// }

	// if attr, ok := d.GetOk("description"); ok {
	// 	user.Description = attr.(string)
	// }

	return show
}

func expandShowRegionsPlaylists(v interface{}) ([]*yodeck.ShowRegionsPlaylists, error) {
	var showRegionsPlaylists []*yodeck.ShowRegionsPlaylists

	for _, sr := range v.([]interface{}) {
		rsr := sr.(map[string]interface{})
		showRegionsPlaylist := &yodeck.ShowRegionsPlaylists{
			Playlist: rsr["playlist"].(string),
		}
		showRegionsPlaylists = append(showRegionsPlaylists, showRegionsPlaylist)
	}

	return showRegionsPlaylists, nil
}
func expandShowRegions(v interface{}) ([]*yodeck.ShowRegions, error) {
	var showRegions []*yodeck.ShowRegions
	// var pmIndex int
	for _, sr := range v.([]interface{}) {
		rsr := sr.(map[string]interface{})

		// pmIndex++

		// var showRegionsPlaylists []*yodeck.ShowRegionsPlaylists

		showRegion := &yodeck.ShowRegions{

			Left:               rsr["left"].(int),
			Top:                rsr["top"].(int),
			Width:              rsr["width"].(int),
			Height:             rsr["height"].(int),
			Fit:                rsr["fit"].(string),
			EnableTransparency: rsr["enable_transparency"].(bool),
			IsMuted:            rsr["is_muted"].(bool),
			ResWidth:           rsr["res_width"].(int),
			ResHeight:          rsr["res_height"].(int),
			BackgroundAudio:    rsr["background_audio"].(bool),
		}

		if pl, ok := rsr["playlists"]; ok {
			playlists, err := expandShowRegionsPlaylists(pl)
			if err != nil {
				return nil, nil
			}

			showRegion.Playlists = playlists
		}
		// for _, slu := range rsl["users"].([]interface{}) {
		// 	user := &pagerduty.UserReferenceWrapper{
		// 		User: &pagerduty.UserReference{
		// 			ID:   slu.(string),
		// 			Type: "user",
		// 		},
		// 	}
		// 	scheduleLayer.Users = append(scheduleLayer.Users, user)
		// }

		showRegions = append(showRegions, showRegion)
	}

	return showRegions, nil
}

func flattenShowRegions(v []*yodeck.ShowRegions) []map[string]interface{} {
	var showRegions []map[string]interface{}

	for _, sh := range v {

		var playlists = flattenShowRegionsPlaylists(sh.Playlists)

		showRegion := map[string]interface{}{
			"left":                sh.Left,
			"top":                 sh.Top,
			"width":               sh.Width,
			"height":              sh.Height,
			"fit":                 sh.Fit,
			"enable_transparency": sh.EnableTransparency,
			"is_muted":            sh.IsMuted,
			"res_width":           sh.ResWidth,
			"res_height":          sh.ResHeight,
			"background_audio":    sh.BackgroundAudio,
			"zindex":              sh.Zindex,
		}
		showRegion["playlists"] = playlists

		showRegions = append(showRegions, showRegion)
	}

	return showRegions
}

func flattenShowRegionsPlaylists(v []*yodeck.ShowRegionsPlaylists) []map[string]interface{} {
	var showRegionsPlaylists []map[string]interface{}

	for _, shpl := range v {
		showRegionsPlaylist := map[string]interface{}{
			"playlist": shpl.Playlist,
			"duration": shpl.Duration,
			"order":    shpl.Order,
		}

		showRegionsPlaylists = append(showRegionsPlaylists, showRegionsPlaylist)
	}

	return showRegionsPlaylists
}
