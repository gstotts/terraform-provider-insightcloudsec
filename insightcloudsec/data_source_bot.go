package insightcloudsec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBot() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudTypesRead,
		Schema: map[string]*schema.Schema{
			"bots": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource_id associated with the bot",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the bot",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description assigned to the bot",
						},
					},
				},
			},
		},
	}
}
