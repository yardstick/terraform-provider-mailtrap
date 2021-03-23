package provider

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yardstick/terraform-provider-mailtrap/api"
	"github.com/yardstick/terraform-provider-mailtrap/log"
)

func resourceInbox() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource, also acts as it's unique ID",
			},
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The project to put the inbox into",
			},
			"smtp_host": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"smtp_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"smtp_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: resourceCreateInbox,
		Read:   resourceReadInbox,
		Update: resourceUpdateInbox,
		Delete: resourceDeleteInbox,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateInbox(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*api.Client)

	params := map[string]interface{}{
		"inbox": map[string]string{
			"name": d.Get("name").(string),
		},
	}

	resp, err := apiClient.Post("/companies/"+d.Get("project_id").(string)+"/inboxes", params)
	if err != nil {
		log.Error("Could not post to API properly...")
		log.Error("API Response: " + err.Error())
		return err
	}
	d.SetId(strconv.Itoa(int(resp["id"].(float64))))
	d.Set("smtp_host", resp["domain"].(string))
	d.Set("smtp_username", resp["username"].(string))
	d.Set("smtp_password", resp["password"].(string))

	return nil
}

func resourceReadInbox(d *schema.ResourceData, m interface{}) error {
	internalId := d.Id()

	if internalId != "" {
		apiClient := m.(*api.Client)

		_, err := apiClient.Get("/inboxes/" + internalId)
		if err != nil {
			d.SetId("")
			return nil
		}
	}

	return nil
}

func resourceUpdateInbox(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*api.Client)
	internalId := d.Id()

	params := map[string]interface{}{
		"inbox": map[string]interface{}{
			"name": d.Get("name").(string),
		},
	}

	resp, err := apiClient.Patch("/inboxes/"+internalId, params)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(int(resp["id"].(float64))))
	d.Set("smtp_host", resp["domain"].(string))
	d.Set("smtp_username", resp["username"].(string))
	d.Set("smtp_password", resp["password"].(string))

	return nil
}

func resourceDeleteInbox(d *schema.ResourceData, m interface{}) error {
	internalId := d.Id()

	apiClient := m.(*api.Client)

	_, err := apiClient.Delete("/inboxes/" + internalId)
	if err != nil {
		return err
	}

	d.SetId("")
	d.Set("smtp_host", "")
	d.Set("smtp_username", "")
	d.Set("smtp_password", "")

	return nil
}
