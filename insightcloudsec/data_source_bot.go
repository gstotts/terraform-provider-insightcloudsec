package insightcloudsec

import (
	"context"
	"log"

	ics "github.com/gstotts/insightcloudsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBot() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBotRead,
		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:        schema.TypeString,
				Required:    true,
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
			"owner": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The owner for the bot given as a user id",
			},
			"owner_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the owner for the bot",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current state of the bot",
			},
			"event_failures": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Counts of any event failures for the bot",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"errors": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Error count for the bot",
						},
						"timeouts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Timeout count for the bot",
						},
						"invalid_perms": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Invalid permissions count for the bot",
						},
					},
				},
			},
			"valid": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the configuration is valid or not for the bot",
			},
			"hookpoint_created": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if a hookpoint has been created for the bot",
			},
			"hookpoint_tags_modified": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if hookpoint tags have been modified for the bot",
			},
			"hookpoint_modified": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if the hookpoint has been modified for the bot",
			},
			"hookpoint_destroyed": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if the hookpoint has been destroyed for the bot",
			},
			"schedule": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "String representing the information for scheduling of the bot",
			},
			"next_scheduled_run": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "String represnting the next scheduled run time",
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp for bot creation",
			},
			"modified_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp for last bot modification",
			},
			"category": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Category assigned to the bot",
			},
			"severity": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Severity assigned to the bot",
			},
			"detailed_logging": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates status of detailed logging",
			},
			"instructions": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Instructions for bot configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_types": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Resource types to which the bot applies",
							Elem:        schema.TypeString,
						},
						"actions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Actions associated with the bot",
							Elem: schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the bot action",
									},
									"config": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "A map of the configurations for the bot action",
									},
									"run_when_result_is": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the bot runs when the config is true or false",
									},
								},
							},
						},
						"filters": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filters utilized in identifying resources for the bot",
							Elem: schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the filter being used",
									},
									"config": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "A map of the configurations for the filter",
									},
								},
							},
						},
					},
				},
			},
			"source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Bot",
			},
			"insight_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Insight id for an associated insight",
			},
			"exemptions_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of exemptions associated with the bot or bot insight",
			},
			"notes": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Notes attached to the bot",
			},
			"version": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Version number of the bot",
			},
			"count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Bot count",
			},
		},
	}
}

func dataSourceBotRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)
	var diags diag.Diagnostics
	resource_id := d.Get("resource_id").(string)

	bot, err := c.GetBotByID(resource_id)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Println("[DEBUG] Bot Returned from API: \n%", bot)
	d.SetId(bot.ResourceID)
	d.Set("name", bot.Name)
	d.Set("description", bot.Description)
	d.Set("owner", bot.Owner)
	d.Set("owner_name", bot.OwnerName)
	d.Set("state", bot.State)

	// Flatten Event Failures?

	d.Set("valid", bot.Valid)

	//Hookpoints, Schedule / Next Scheduled no longer part of instructions?

	d.Set("creation_timestamp", bot.DateCreated)
	d.Set("modified_timestamp", bot.DateModified)
	d.Set("category", bot.Category)
	d.Set("severity", bot.Severity)
	d.Set("detailed_logging", bot.DetailedLogging)

	// Flatten Instructions?

	d.Set("source", bot.Source)
	d.Set("insight_id", bot.InsightID)

	// Exemptions Count not in client?
	// Consider revising client to return raw responses so can just apply json decoding?

	d.Set("notes", bot.Notes)
	// Versions not in client?

	return diags
}
