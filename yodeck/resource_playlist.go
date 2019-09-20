package yodeck

import (
	"fmt"
	"log"

	"github.com/easylo/go-yodeck/yodeck"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceYodeckPlaylist() *schema.Resource {
	return &schema.Resource{
		Create: resourceYodeckPlaylistCreate,
		Read:   resourceYodeckPlaylistRead,
		Update: resourceYodeckPlaylistUpdate,
		Delete: resourceYodeckPlaylistDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"workspace": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"media": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"media": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"duration": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceYodeckPlaylistCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	playlist := buildPlaylistStruct(d)

	log.Printf("[INFO] Creating Yodeck Playlist %s", playlist.Name)

	playlistResponse, _, err := client.Playlist.Create(playlist)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprint(playlistResponse.ID))

	return nil
}

func resourceYodeckPlaylistRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)
	log.Printf("[INFO] Reading PagerDuty user %s", d.Id())

	playlist, _, err := client.Playlist.Get(d.Id())
	if err != nil {
		return handleNotFoundError(err, d)
	}

	d.Set("name", playlist.Name)
	d.Set("description", playlist.Description)
	d.Set("workspace", playlist.Workspace)

	if err := d.Set("media", flattenPlaylistMedia(playlist.Media)); err != nil {
		return fmt.Errorf("error setting teams: %s", err)
	}

	return nil
}

func resourceYodeckPlaylistUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	playlist := buildPlaylistStruct(d)

	log.Printf("[INFO] Updating Yodeck Playlist %s", d.Id())

	if _, _, err := client.Playlist.Update(d.Id(), playlist); err != nil {
		return err
	}
	return nil
}

func resourceYodeckPlaylistDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*yodeck.Client)

	log.Printf("[INFO] Deleting Yodeck Playlist %s", d.Id())

	if _, err := client.Playlist.Delete(d.Id()); err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func buildPlaylistStruct(d *schema.ResourceData) *yodeck.Playlist {
	playlist := &yodeck.Playlist{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	if attr, ok := d.GetOk("workspace"); ok {
		playlist.Workspace = attr.(int)
	}

	if attr, ok := d.GetOk("media"); ok {
		media, err := expandPlaylistMedia(attr)
		if err != nil {
			return nil
		}

		playlist.Media = media
	}

	return playlist
}

func expandPlaylistMedia(v interface{}) ([]*yodeck.PlaylistMedia, error) {
	var playlistMedias []*yodeck.PlaylistMedia
	var pmIndex int
	for _, pm := range v.([]interface{}) {
		rpm := pm.(map[string]interface{})

		pmIndex++

		playlistMedia := &yodeck.PlaylistMedia{
			Priority: pmIndex,
			Media:    rpm["media"].(int),
			Duration: rpm["duration"].(int),
		}

		playlistMedias = append(playlistMedias, playlistMedia)
	}

	return playlistMedias, nil
}

func flattenPlaylistMedia(v []*yodeck.PlaylistMediaResponse) []map[string]interface{} {
	var playlistMedias []map[string]interface{}

	for _, pm := range v {
		playlistMedia := map[string]interface{}{
			"priority": pm.Priority,
			"media":    pm.Media.ID,
			"duration": pm.Duration,
		}

		playlistMedias = append(playlistMedias, playlistMedia)
	}

	return playlistMedias
}

// func castSetToSliceStrings(configured []interface{}) []string {
// 	res := make([]string, len(configured))

// 	for i, element := range configured {
// 		res[i] = element.(string)
// 	}
// 	return res
// }
