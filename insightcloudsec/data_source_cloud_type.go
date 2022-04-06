package insightcloudsec

import (
	"context"
	"strconv"
	"time"

	ics "github.com/gstotts/insightcloudsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudTypes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudTypesRead,
		Schema: map[string]*schema.Schema{
			"clouds": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloud_access": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The accessibility of the cloud. Possible values:  public, private",
						},
						"cloud_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The identifier assigned to the cloud type within InsightCloudSec",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The full name of the cloud type",
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudTypesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)

	var diags diag.Diagnostics
	ctypes, _ := c.Clouds.ListTypes()

	typeDetails := make([]interface{}, 0)
	for _, d := range ctypes.CloudTypes {
		typeDetails = append(typeDetails, map[string]interface{}{
			"cloud_type":   d.ID,
			"name":         d.Name,
			"cloud_access": d.Access,
		})
	}

	if err := d.Set("clouds", typeDetails); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
