package insightcloudsec

import (
	"context"
	"log"
	"strconv"
	"time"

	ics "github.com/gstotts/insightcloudsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUsersRead,
		Schema: map[string]*schema.Schema{
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource id for the specific user",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the user",
						},
						"user_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The user id",
						},
						"organization_admin": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates if the user is an organization administrator (true) or not (false)",
						},
						"domain_admin": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates if the user is a domain administrator (true) or not (false)",
						},
						"domain_viewer": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates if the user is a domain viewer (true) or not (false)",
						},
						"email_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The email address for the user",
						},
						"username": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The username for the user",
						},
						"organization_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The organization name to which the user belongs",
						},
						"organization_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The organization id for the organization to which the user bleongs",
						},
						"two_factor_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates if two factor is enabled for the user",
						},
						"two_factor_required": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates if two factor is required for the user",
						},
						"groups": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of groups to which the user belongs",
						},
						"owned_resources": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of resources where the user is identified as an owner",
						},
						"consecutive_failed_login_attempts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of consecutive failed login attempts for the user",
						},
						"suspended": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates if the user is suspended",
						},
						"last_login_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last login time for the user",
						},
						"navigation_blacklist": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Lists any blacklisted navigation links for the user",
							Elem:        schema.TypeString,
						},
						"require_pw_reset": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates if the user's password is required to be reset",
						},
						"console_access_denied": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates if the user is denied console access",
						},
						"active_api_key_present": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates if the user has an active api key associated",
						},
						"create_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date the user was created",
						},
					},
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of users returned",
			},
		},
	}
}

func dataSourceUsersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)

	var diags diag.Diagnostics

	users, err := c.ListUsers()
	if err != nil {
		return diag.FromErr(err)
	}

	log.Println("[DEBUG] Users Returned from API: \n%", users)

	d.Set("total_count", users.Count)

	userDetails := make([]interface{}, 0)
	for _, user := range users.Users {
		userDetails = append(userDetails, map[string]interface{}{
			"resource_id":                       user.ResourceID,
			"name":                              user.Name,
			"user_id":                           user.ID,
			"organization_admin":                user.OrgAdmin,
			"domain_admin":                      user.DomainAdmin,
			"domain_viewer":                     user.DomainViewer,
			"email_address":                     user.Email,
			"username":                          user.Username,
			"organization_name":                 user.Org,
			"organization_id":                   user.OrgID,
			"two_factor_enabled":                user.TwoFactorEnabled,
			"two_factor_required":               user.TwoFactorRequired,
			"groups":                            user.Groups,
			"owned_resources":                   user.OwnedResources,
			"consecutive_failed_login_attempts": user.FailedLoginAttempts,
			"suspended":                         user.Suspended,
			"last_login_time":                   user.LastLogin,
			"navigation_blacklist":              user.NavigationBlacklist,
			"require_pw_reset":                  user.RequirePWReset,
			"console_access_denied":             user.ConsoleAccessDenied,
			"active_api_key_present":            user.ActiveAPIKey,
			"create_date":                       user.Created,
		})
	}

	if err := d.Set("users", userDetails); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
