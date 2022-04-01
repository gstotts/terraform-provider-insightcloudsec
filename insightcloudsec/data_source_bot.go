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
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Counts of any event failures for the bot",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
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
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Map representing the information for scheduling of the bot",
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
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Resource types to which the bot applies",
							Elem:        schema.TypeString,
						},
						"actions": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Actions associated with the bot",
							Elem: &schema.Resource{
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
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "The filters utilized in identifying resources for the bot",
							Elem: &schema.Resource{
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

	log.Println("[DEBUG] Bot Returned from API: \n", bot)

	log.Println("[DEBUG] Schedule: \n", bot.Schedule)
	d.SetId(bot.ResourceID)
	d.Set("name", bot.Name)
	d.Set("description", bot.Description)
	d.Set("owner", bot.Owner)
	d.Set("owner_name", bot.OwnerName)
	d.Set("state", bot.State)
	d.Set("event_failures", map[string]int{
		"errors":        bot.EventFailures.Errors,
		"timeouts":      bot.EventFailures.Timeouts,
		"invalid_perms": bot.EventFailures.InvalidPerms,
	})
	d.Set("valid", bot.Valid)

	// Schedule
	schedule := make(map[string]interface{})
	schedule["_type"] = bot.Schedule.Type
	if schedule["_type"] == "Hourly" {
		schedule["minute_of_hour"] = bot.Schedule.MinuteOfHour
		schedule["second_of_hour"] = bot.Schedule.SecondOfHour
	} else if schedule["_type"] == "Daily" {
		schedule["time_of_day"] = map[string]int{
			"minute": bot.Schedule.TimeOfDay.Minute,
			"hour":   bot.Schedule.TimeOfDay.Hour,
		}
		schedule["exclude_days"] = bot.Schedule.ExcludeDays
	} else if schedule["_type"] == "Weekly" {
		schedule["time_of_day"] = map[string]int{
			"minute": bot.Schedule.TimeOfDay.Minute,
			"hour":   bot.Schedule.TimeOfDay.Hour,
		}
		schedule["day_of_week"] = bot.Schedule.DayOfWeek
	} else if schedule["_type"] == "Monthly" {
		schedule["time_of_day"] = map[string]int{
			"minute": bot.Schedule.TimeOfDay.Minute,
			"hour":   bot.Schedule.TimeOfDay.Hour,
		}
		schedule["day_of_month"] = bot.Schedule.DayOfMonth
	}
	d.Set("schedule", schedule)

	d.Set("next_scheduled_run", bot.NextScheduled)
	d.Set("hookpoint_created", bot.HookpointCreated)
	d.Set("hookpoint_modified", bot.HookpointModified)
	d.Set("hookpoint_destroyed", bot.HookpointDestroyed)
	d.Set("hookpoint_tags_modified", bot.HookpointTagsModified)
	d.Set("creation_timestamp", bot.DateCreated)
	d.Set("modified_timestamp", bot.DateModified)
	d.Set("category", bot.Category)
	d.Set("severity", bot.Severity)
	d.Set("detailed_logging", bot.DetailedLogging)
	d.Set("resource_types", bot.Instructions.ResourceTypes)
	instructions := make([]interface{}, 0)
	instructions = append(instructions, map[string]interface{}{
		"resource_types": bot.Instructions.ResourceTypes,
		"actions":        flattenBotActionsData(&bot.Instructions.Actions),
		"filters":        flattenBotFiltersData(&bot.Instructions.Filters),
	})
	d.Set("source", bot.Source)
	d.Set("insight_id", bot.InsightID)
	d.Set("exemptions_count", bot.ExemptionsCount)
	d.Set("notes", bot.Notes)
	d.Set("version", bot.Version)

	return diags
}

func flattenBotFiltersData(filters *[]ics.BotFilter) []interface{} {
	if filters != nil {
		data := make([]interface{}, len(*filters), len(*filters))

		for i, filter := range *filters {
			data_b := make(map[string]interface{})
			data_b["name"] = filter.Name
			data_b["filter_config"] = filter.Config
			data[i] = data_b
		}

		return data
	}

	return make([]interface{}, 0)
}

func flattenBotActionsData(actions *[]ics.BotAction) []interface{} {
	if actions != nil {
		data := make([]interface{}, len(*actions), len(*actions))

		for i, action := range *actions {
			data_b := make(map[string]interface{})
			data_b["name"] = action.Name
			data_b["action_config"] = action.Config
			data_b["run_when_result_is"] = action.RunWhenResultIs
			data[i] = data_b
		}

		return data
	}

	return make([]interface{}, 0)
}
