package insightcloudsec

import (
	"context"
	"fmt"
	"strconv"

	ics "github.com/gstotts/insightcloudsec"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceInsight() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInsightCreate,
		ReadContext:   resourceInsightRead,
		UpdateContext: resourceInsightUpdate,
		DeleteContext: resourceInsightDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID assigned to the insight.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the insight for display in InsightCloudSec",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The description to assign to the insight",
			},
			"severity": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "The severity associated with the insight represented by an int",
				ValidateFunc: validation.IntBetween(1, 5),
			},
			"scopes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The scope for the insight",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Filter used with the insight to determine resources",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the filter",
						},
						"config": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "The configuration of the filter",
						},
						"collections": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "The collections associated with the filter",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Tags to associate with the insight",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"badges": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Badges used to limit the insight",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"badge_filter_operator": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The badge filter operator for the insight",
				Default:     "OR",
			},
			"resource_types": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Resource types the insight applies to",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceInsightCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)
	var diags diag.Diagnostics

	filters := d.Get("filter").([]interface{})
	fis := []ics.InsightFilter{}

	for _, filter := range filters {
		i := filter.(map[string]interface{})
		fi := ics.InsightFilter{
			Name:        i["name"].(string),
			Config:      i["config"].(map[string]interface{}),
			Collections: i["collections"].(map[string]interface{}),
		}
		fis = append(fis, fi)
	}

	// scopes := interfaceToList(d.Get("scopes").([]interface{}))
	// tags := interfaceToList(d.Get("tags").([]interface{}))
	// badges := interfaceToList(d.Get("badges").([]interface{}))
	resourcetypes := interfaceToList(d.Get("resource_types").([]interface{}))

	tflog.Debug(ctx, fmt.Sprintf("%s", resourcetypes))

	insight := ics.Insight{
		Name:                d.Get("name").(string),
		Description:         d.Get("description").(string),
		Severity:            d.Get("severity").(int),
		ResourceTypes:       resourcetypes,
		Filters:             fis,
		BadgeFilterOperator: d.Get("badge_filter_operator").(string),
	}

	resp, err := c.Insights.Create(insight)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(resp.ID))
	resourceInsightRead(ctx, d, m)
	return diags
}

func resourceInsightRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)
	var diags diag.Diagnostics

	id, _ := strconv.Atoi(d.Get("id").(string))
	insight, err := c.Insights.Get_Insight(id, "custom")
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", insight.Name)
	d.Set("description", insight.Description)
	d.Set("severity", insight.Severity)
	d.Set("filter", insight.Filters)
	// d.Set("scopes", insight.Scopes)
	// d.Set("tags", insight.Tags)
	// d.Set("badges", insight.Badges)
	d.Set("badge_filter_operator", insight.BadgeFilterOperator)
	d.Set("resource_types", insight.ResourceTypes)

	return diags
}

func resourceInsightUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diag diag.Diagnostics
	return diag
}

func resourceInsightDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diag diag.Diagnostics
	return diag
}

func interfaceToList(i []interface{}) []string {
	s := make([]string, 0, len(i))
	for _, item := range i {
		s = append(s, item.(string))
	}
	return s
}
