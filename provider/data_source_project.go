package provider

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yardstick/terraform-provider-mailtrap/api"
)

func datasourceProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceReadProject,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceReadProject(d *schema.ResourceData, m interface{}) error {

	name := d.Get("name")
	apiClient := m.(*api.Client)

	resp, err := apiClient.GetArray("/companies")
	if err != nil {
		return err
	}

	for _, project := range resp {
		if project["name"] == name {
			d.SetId(strconv.Itoa(int(project["id"].(float64))))
			return nil
		}
	}

	return nil
}
