package insightcloudsec

import (
	"context"
	"regexp"

	ics "github.com/gstotts/insightcloudsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID assigned to the user",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the user",
			},
			"email_address": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The email for the user",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringMatch(regexp.MustCompile(`[\w+=,.-]+@[\w.-]+\.[\w]+`), "must be a valid email address")),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The username for the user",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The password for the user",
			},
			"access_level": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The access level to associate with the user",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"BASIC_USER", "ORGANIZATION_ADMIN", "DOMAIN_VIEWER", "DOMAIN_ADMIN"}, false)),
			},
			"two_factor_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Default:     false,
				Description: "Boolean representing if 2FA is enabled for the user",
			},
			"two_factor_required": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Boolean representing whether 2FA is required for this user",
			},
			"organization_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The organization id to which the user belongs",
			},
			"organization_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The organization name to which the user belongs",
			},
			"domain_admin": {
				Type:        schema.TypeBool,
				Computed:    true,
				Default:     false,
				Description: "Boolean representing if the user is a domain admin",
			},
			// Both attributes below are only computed -- cannot be configured.
			// For now, these are not included as tracking them would result in changes uncontrollable to the user directly.
			// --------------
			// "groups": {
			// 	Type:        schema.TypeInt,
			// 	Computed:    true,
			// 	Description: "Int representing the number of groups associated",
			// },
			// "owned_resources": {
			// 	Type:        schema.TypeInt,
			// 	Computed:    true,
			// 	Description: "Int representing the number of owned resources",
			// },
			"resource_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource_id for the user",
			},
			"suspended": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Default:     false,
				Description: "Boolean representing if user is suspended",
			},
			"require_pw_reset": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Default:     false,
				Description: "Boolean representing if a user's password is required to be reset",
			},
			"console_access_denied": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Default:     false,
				Description: "Boolean representing if a user's console access is currently denied",
			},
		},
	}
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ics.Client)
	user, err := c.Users.GetUserByID(d.Get("user_id").(int))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", user.Name)
	d.Set("email_address", user.Email)
	d.Set("username", user.Username)
	if user.DomainAdmin {
		d.Set("access_level", "DOMAIN_ADMIN")
	} else if user.OrgAdmin {
		d.Set("access_level", "ORGANIZATION_ADMIN")
	} else if user.DomainViewer {
		d.Set("access_level", "DOMAIN_VIEWER")
	} else {
		d.Set("access_level", "BASIC_USER")
	}
	d.Set("two_factor_enabled", user.TwoFactorEnabled)
	d.Set("two_factor_required", user.TwoFactorRequired)
	d.Set("organization_name", user.Org)
	d.Set("domain_admin", user.DomainAdmin)
	d.Set("resource_id", user.ResourceID)
	d.Set("suspended", user.Suspended)
	d.Set("require_pw_reset", user.RequirePWReset)
	d.Set("console_access_denied", user.ConsoleAccessDenied)

	return diags
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}
