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
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the cloud to retrieve",
			},
			"id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID associated the the data source for this cloud",
			},
			"cloud_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The identifier for the type of cloud utilized.  Examples:  AWS, AZURE_ARM, etc.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The identifier for the account associated with the cloud.  In the case of AWS, this is the account ID.  In Azure, this is the subscription ID",
			},
			"resource_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource ID provided by the console for the cloud",
			},
			"strategy_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The harvesting strategy ID for the cloud",
			},
			"cloud_organization_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The organization ID for the cloud",
			},
			"group_resource_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The group resource ID for the cloud",
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
	d.Set("account_id", cloud.AccountID)
	d.Set("resource_id", cloud.ResourceID)
	d.Set("strategy_id", cloud.StrategyID)
	d.Set("cloud_organization_id", cloud.CloudOrgID)
	d.Set("group_resource_id", cloud.GroupResourceID)
	d.Set("resource_id", cloud.ResourceID)

	d.SetId(strconv.Itoa(cloud.ID))

	return diags
}
