package insightcloudsec

import (
	"context"

	ics "github.com/gstotts/insightcloudsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource id for the specific user",
			},
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the user",
				Required:    true,
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
				Description: "The email address for the user",
				Required:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "The username for the user",
				Required:    true,
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
				Optional:    true,
			},
			"groups": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of groups to which the user belongs",
				Optional:    true,
			},
			"owned_resources": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of resources where the user is identified as an owner",
				Optional:    true,
			},
			"suspended": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if the user is suspended",
			},
			"navigation_blacklist": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Lists any blacklisted navigation links for the user",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
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
	}
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)

	var diags diag.Diagnostics

	resp, err := c.Users.GetUserByID(d.Get("user_id").(int))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("resource_id", resp.ResourceID)
	d.Set("name", resp.Name)
	d.Set("user_id", resp.ID)
	d.Set("organization_admin", resp.OrgAdmin)
	d.Set("domain_admin", resp.DomainAdmin)
	d.Set("domain_viewer", resp.DomainViewer)
	d.Set("email_address", resp.Email)
	d.Set("username", resp.Username)
	d.Set("organization_name", resp.Org)
	d.Set("organization_id", resp.OrgID)
	d.Set("two_factor_enabled", resp.TwoFactorEnabled)
	d.Set("two_factor_required", resp.TwoFactorRequired)
	d.Set("groups", resp.Groups)
	d.Set("owned_resources", resp.OwnedResources)
	d.Set("suspended", resp.Suspended)
	d.Set("navigation_blacklist", resp.NavigationBlacklist)
	d.Set("require_pw_reset", resp.RequirePWReset)
	d.Set("console_access_denied", resp.ConsoleAccessDenied)
	d.Set("active_api_key_present", resp.ActiveAPIKey)
	d.Set("create_date", resp.Created)
	d.SetId(d.Get("resource_id").(string))

	return diags
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)

	var diags diag.Diagnostics
	user := ics.User{
		Name:        d.Get("name").(string),
		Username:    d.Get("username").(string),
		Email:       d.Get("email_address").(string),
		AccessLevel: "BASIC_USER",
	}

	// If 2FA Required, set that value to true, otherwise false
	mfa_req, exists := d.GetOk("two_factor_required")
	if exists {
		user.TwoFactorRequired = mfa_req.(bool)
	} else {
		user.TwoFactorRequired = false
	}

	resp, err := c.Users.Create(user)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("user_id", resp.ID)

	resourceUserRead(ctx, d, m)

	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)

	var diags diag.Diagnostics
	err := c.Users.Delete(d.Get("resource_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
