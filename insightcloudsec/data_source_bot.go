package insightcloudsec

import (
	"context"
	"log"

	ics "github.com/gstotts/insightcloudsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBot() *schema.Resource {
	i := schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource_types": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Resource types to which the bot applies",
			},
			// "actions": {
			// 	Type:        schema.TypeSet,
			// 	Computed:    true,
			// 	Description: "Actions the bot will take",
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"name": {
			// 				Type:        schema.TypeString,
			// 				Computed:    true,
			// 				Description: "Name of the action",
			// 			},
			// 			"config": {
			// 				Type:        schema.TypeMap,
			// 				Computed:    true,
			// 				Description: "Map of action configuration settings",
			// 			},
			// 			"run_when_result_is": {
			// 				Type:        schema.TypeBool,
			// 				Computed:    true,
			// 				Description: "Boolean that determins if the actino runs when the condition is met or when it is not met",
			// 			},
			// 		},
			// 	},
			// },
			// "hookpoints": {
			// 	Type:        schema.TypeList,
			// 	Computed:    true,
			// 	Elem:        &schema.Schema{Type: schema.TypeString},
			// 	Description: "Hookpoints associated with when the bot is to run from EDH",
			// },
		},
	}
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
			"notes": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Notes attached to the bot",
			},
			"insight_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Insight id for an associated insight (if exists)",
			},
			"source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Bot source",
			},
			"insight_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of associated insight (if exists)",
			},
			"insight_severity": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Severity of associated insight (if exists)",
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
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of bot creation",
			},
			"date_modified": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of bot creation",
			},
			"category": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Category of the bot",
			},
			"badge_scope_operator": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Operator used for badge scoping - AND or OR",
			},
			"instructions": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      schema.HashResource(&i),
				Elem:     &i,
			},
		},
	}
}

func dataSourceBotRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)
	var diags diag.Diagnostics
	resource_id := d.Get("resource_id").(string)

	bot, err := c.Bots.GetBotByID(resource_id)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Println("\n\n[DEBUG] Bot Returned from API: \n", bot)
	d.SetId(bot.ResourceID)
	d.Set("resource_id", bot.ResourceID)
	d.Set("name", bot.Name)
	d.Set("description", bot.Description)
	d.Set("notes", bot.Notes)
	d.Set("insight_id", bot.InsightID)
	d.Set("source", bot.Source)
	d.Set("insight_name", bot.InsightName)
	d.Set("insight_severity", bot.Severity)
	d.Set("owner", bot.Owner)
	d.Set("owner_name", bot.OwnerName)
	d.Set("state", bot.State)
	d.Set("date_created", bot.DateCreated)
	d.Set("date_modified", bot.DateModified)
	d.Set("category", bot.Category)
	d.Set("badge_scope_operator", bot.BadgeScopeOperator)

	instructs := []map[string]interface{}{}
	instructs = append(instructs, map[string]interface{}{
		"resource_types": bot.Instructions.ResourceTypes,
		// "actions":        bot.Instructions.Actions,
	})

	d.Set("instructions", instructs)

	return diags
}
