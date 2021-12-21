package insightcloudsec

import (
	"context"
	"log"
	"strconv"

	ics "github.com/gstotts/insightcloudsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasSourceCloud() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cloud_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"strategy_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"group_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCloudRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)
	var diags diag.Diagnostics
	cloud_name := d.Get("name").(string)

	cloud, err := c.GetCloudByName(cloud_name)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Println("[DEBUG] Cloud Returned from API: \n%", cloud)

	d.Set("name", cloud.Name)
	d.Set("cloud_type", cloud.CloudTypeID)
	d.Set("account", cloud.AccountID)
	d.Set("resource_id", cloud.ResourceID)
	d.Set("strategy_id", cloud.StrategyID)
	d.Set("group_resource_id", cloud.GroupResourceID)
	d.Set("resource_id", cloud.ResourceID)

	d.SetId(strconv.Itoa(cloud.ID))

	return diags
}
