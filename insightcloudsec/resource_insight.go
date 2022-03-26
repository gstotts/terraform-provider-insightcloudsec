package insightcloudsec

import (
	ics "github.com/gstotts/insightcloudsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceInsight() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInsightCreate,
		ReadContext: resourceInsightRead,
		UpdateContext: resourceInsightUpdate,
		DeleteContext: resourceInsightDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
				Description: "The name of the insight for display in InsightCloudSec",
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
				Default: '',
				Description: "The description to assign to the insight"
			},
			"severity": {
				Type: schema.TypeString,
				Required: true,
				Description: "The severity associated with the instance.  Must be one of: critical, severe, major, moderate, minor"
				ValidateFunc: validation.StringInSlice([]{"critical", "severe", "major", "moderate", "minor"})
			},
			"filters": {
				Type: schema.TypeList
				Required: true,
				Description: "Filter used with the insight to determine resources"
				Elem: 
			}
		}
	}
}

func resourceInsightCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)
	var diags diag.Diagnostics
	var err error
	var insight ics.Insight 

	filter := ics.InsightFilter{
		Name: ,
	}

}