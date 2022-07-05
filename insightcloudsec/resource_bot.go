package insightcloudsec

import (
	"context"
	"log"

	ics "github.com/gstotts/insightcloudsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var (
	// Bot Setting Options
	VALID_BOT_STATES          = []string{"paused", "running", "archived", "scanning"}
	VALID_BOT_CATEGORIES      = []string{"Security", "Optimization", "Best Practices", "Curation", "Miscellaneous"}
	VALID_BOT_SEVERITIES      = []string{"low", "medium", "high"}
	VALID_BOT_BADGE_OPERATORS = []string{"OR", "AND"}
	VALID_BOT_HOOKPOINTS      = []string{"divvycloud.resource.created", "divvycloud.resource.tags_modified", "divvycloud.resource.modified", "divvycloud.resource.destroyed", "divvycloud.resource.threat_finding_discovered"}
)

func resourceBot() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBotCreate,
		ReadContext:   resourceBotRead,
		UpdateContext: resourceBotUpdate,
		DeleteContext: resourceBotDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name for the bot",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The description for the bot",
			},
			"state": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "PAUSED",
				Description: "The current state of the bot.  Defaults to PAUSED",
			},
			"badge_state_operator": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "OR",
				ValidateFunc: validation.StringInSlice(VALID_BOT_BADGE_OPERATORS, false),
			},
			"category": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(VALID_BOT_CATEGORIES, false),
				Description:  "The category assigned to the bot",
			},
			"severity": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(VALID_BOT_SEVERITIES, false),
				Description:  "The severity assigned to the bot",
			},
			"instructions": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"groups": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     schema.TypeString,
						},
						"badges": {
							Type:     schema.TypeList,
							Optional: true,
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
						"resource_types": {
							Type:     schema.TypeList,
							Required: true,
							Elem:     schema.TypeString,
						},
						"filters": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"config": {
										Type:     schema.TypeList,
										Required: true,
										Elem:     schema.TypeString,
									},
								},
							},
						},
						"actions": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"config": {
										Type:     schema.TypeList,
										Required: true,
										Elem:     schema.TypeString,
									},
									"run_when_result_is": {
										Type:     schema.TypeBool,
										Required: true,
									},
								},
							},
						},
						"hookpoints": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     schema.TypeString,
						},
						"schedule": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"_type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"monthly, weekly, daily, hourly"}, false),
									},
									"time_of_day": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"_type": {
													Type:     schema.TypeString,
													Required: true,
												},
												"second": {
													Type:         schema.TypeInt,
													Required:     true,
													ValidateFunc: validation.IntBetween(0, 59),
												},
												"minute": {
													Type:         schema.TypeInt,
													Required:     true,
													ValidateFunc: validation.IntBetween(0, 59),
												},
												"hour": {
													Type:         schema.TypeInt,
													Required:     true,
													ValidateFunc: validation.IntBetween(0, 24),
												},
											},
										},
									},
									"day_of_month": {
										Type:         schema.TypeInt,
										Optional:     true,
										ValidateFunc: validation.IntBetween(1, 31),
									},
									"exclude_days": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"insight_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceBotCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// c := m.(*ics.Client)
	var diags diag.Diagnostics

	return diags
}

func resourceBotRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)
	var diags diag.Diagnostics

	resource_id := d.Get("resource_id").(string)
	bot, err := c.Bots.GetBotByID(resource_id)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Println("[DEBUG] Bot Returned from API: \n%", bot)

	// Set attributes
	d.Set("name", bot.Name)
	d.Set("description", bot.Description)
	d.Set("notes", bot.Notes)
	d.Set("state", bot.State)
	d.Set("badge_state_operator", bot.BadgeScopeOperator)
	d.Set("category", bot.Category)
	d.Set("severity", bot.Severity)
	d.Set("groups", bot.Instructions.Groups)

	// Still need more here -- also need to add to above.

	return diags
}

func resourceBotUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// c := m.(*ics.Client)
	var diags diag.Diagnostics

	return diags
}

func resourceBotDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)
	var diags diag.Diagnostics

	err := c.Bots.ArchiveBot(d.Get("resource_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
