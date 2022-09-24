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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
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
							Computed:    true,
							Description: "The configuration of the filter",
						},
						"collections": {
							Type:        schema.TypeMap,
							Optional:    true,
							Computed:    true,
							Description: "The collections associated with the filter",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Tags to associate with the insight",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"badge": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Badges used to limit the insight",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Key for the badge",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Value for the badge",
						},
					},
				},
			},
			"badge_filter_operator": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The badge filter operator for the insight",
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

	insight := prepareInsight(d)

	tflog.Debug(ctx, fmt.Sprintf(
		"Insight Details to Create:\n%v\n", insight,
	))

	resp, err := c.Insights.Create(insight)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, fmt.Sprintf(
		"Response Data from API:\n%v\n", resp,
	))

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

	for _, filter := range insight.Filters {
		if filter.Config == nil {
			filter.Config = make(map[string]interface{}, 0)
		}
		if filter.Collections == nil {
			filter.Collections = make(map[string]interface{}, 0)
		}
	}

	d.Set("filter", insight.Filters)
	d.Set("scopes", insight.Scopes)
	d.Set("tags", insight.Tags)
	d.Set("badges", insight.Badges)
	d.Set("badge_filter_operator", insight.BadgeFilterOperator)
	d.Set("resource_types", insight.ResourceTypes)

	return diags
}

func resourceInsightUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)
	var diags diag.Diagnostics

	id, _ := strconv.Atoi(d.Get("id").(string))
	insight := prepareInsight(d)
	insight.ID = id
	err := c.Insights.Edit(insight)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceInsightDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)
	var diags diag.Diagnostics

	id, _ := strconv.Atoi(d.Get("id").(string))
	err := c.Insights.Delete(id)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func interfaceToList(i []interface{}) []string {
	s := make([]string, 0, len(i))
	for _, item := range i {
		s = append(s, item.(string))
	}
	return s
}

func prepareInsight(d *schema.ResourceData) ics.Insight {
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

	badges := d.Get("badge").([]interface{})
	bis := []ics.Badge{}
	for _, badge := range badges {
		b := badge.(map[string]interface{})
		bi := ics.Badge{
			Key:   b["key"].(string),
			Value: b["value"].(string),
		}
		bis = append(bis, bi)
	}

	return ics.Insight{
		Name:                d.Get("name").(string),
		Description:         d.Get("description").(string),
		Severity:            d.Get("severity").(int),
		ResourceTypes:       interfaceToList(d.Get("resource_types").([]interface{})),
		Filters:             fis,
		Tags:                interfaceToList(d.Get("tags").([]interface{})),
		Scopes:              interfaceToList(d.Get("scopes").([]interface{})),
		Badges:              bis,
		BadgeFilterOperator: d.Get("badge_filter_operator").(string),
	}

}
