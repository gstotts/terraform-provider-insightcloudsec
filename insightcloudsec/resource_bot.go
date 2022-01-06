package insightcloudsec

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBot() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBotCreate,
		ReadContext:   resourceBotRead,
		UpdateContext: resourceBotUpdate,
		DeleteContext: resourceBotDelete,
		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource ID for the bot",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name for the bot",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the bot",
			},
			"owner": {
				Type:        schema.TypeString,
				Description: "The owner ID for whom the bot should belong",
			},
			"owner_name": {
				Type:        schema.TypeString,
				Description: "The owner of the bot's name",
			},
			"state": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The current state of the bot",
				Default:     "PAUSED",
			},
			"event_failures":          {},
			"valid":                   {Type: schema.TypeBool},
			"hookpoint_created":       {Type: schema.TypeBool},
			"hookpoint_tags_modified": {Type: schema.TypeBool},
			"hookpoint_modified":      {Type: schema.TypeBool},
			"hookpoint_destroyed":     {Type: schema.TypeBool},
			"schedule":                {},
			"next_scheduled_run":      {},
			"creation_timestamp":      {Type: schema.TypeString},
			"modified_timestamp":      {Type: schema.TypeString},
			"category":                {Type: schema.TypeString},
			"severity":                {Type: schema.TypeString},
			"detailed_logging":        {Type: schema.TypeBool},
			"instructions": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_types": {Type: schema.TypeList},
						"filters":        {Type: schema.TypeList},
						"actions":        {Type: schema.TypeList},
					},
				},
			},
			"source": {Type: schema.TypeString},
			"insight_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The insight ID associated with the bot",
			},
			"exemptions_count": {},
			"notes": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Notes associated with the bot",
			},
			"version": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Integer represnting the version of the bot",
			},
		},
	}
}

func resourceBotCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceBotRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceBotUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceBotDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
