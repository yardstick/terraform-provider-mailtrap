package provider

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yardstick/terraform-provider-mailtrap/api"
	"github.com/yardstick/terraform-provider-mailtrap/log"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the project",
				ValidateFunc: validateProjectName,
			},
		},
		Create: resourceCreateProject,
		Read:   resourceReadProject,
		Update: resourceUpdateProject,
		Delete: resourceDeleteProject,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func validateProjectName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string

	v, err := v.(string)
	if !err {
		errs = append(errs, fmt.Errorf("name is not a string value"))
		return warns, errs
	}

	if v == "" {
		errs = append(errs, fmt.Errorf("name is an empty string value"))
		return warns, errs
	}

	return warns, errs

}

func resourceCreateProject(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*api.Client)

	params := map[string]interface{}{
		"company": map[string]string{
			"name": d.Get("name").(string),
		},
	}

	resp, err := apiClient.Post("/companies", params)
	if err != nil {
		log.Error("Could not post to API properly...")
		log.Error("API Response: " + err.Error())
		return err
	}
	id := strconv.Itoa(int(resp["id"].(float64)))
	log.Info("Returned ID = " + id)
	d.SetId(id)

	return nil
}

func resourceReadProject(d *schema.ResourceData, m interface{}) error {
	internalId := d.Id()

	if internalId != "" {
		apiClient := m.(*api.Client)

		_, err := apiClient.Get("/companies/" + internalId)
		if err != nil {
			d.SetId("")
			return nil
		}
	}

	return nil
}

func resourceUpdateProject(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*api.Client)
	internalId := d.Id()

	params := map[string]interface{}{
		"company": map[string]interface{}{
			"name": d.Get("name").(string),
		},
	}

	_, err := apiClient.Patch("/companies/"+internalId, params)
	if err != nil {
		return err
	}

	return nil
}

func resourceDeleteProject(d *schema.ResourceData, m interface{}) error {
	internalId := d.Id()

	apiClient := m.(*api.Client)

	_, err := apiClient.Delete("/companies/" + internalId)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
