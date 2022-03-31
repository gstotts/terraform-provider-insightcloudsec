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
							Type:          schema.TypeString,
							Optional:      true,
							Description:   "The resource_id associated with the bot",
							ConflictsWith: []string{"name"},
						},
						"name": {
							Type:          schema.TypeString,
							Optional:      true,
							Description:   "The name of the bot",
							ConflictsWith: []string{"resource_id"},
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
					},
				},
			},
			"count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Bot count",
			},
		},
	}
}
