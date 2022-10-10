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
				ForceNew:    true,
			},
			"access_level": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The access level to associate with the user",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"BASIC_USER", "ORGANIZATION_ADMIN", "DOMAIN_VIEWER", "DOMAIN_ADMIN"}, false)),
			},
			"two_factor_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
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
			"resource_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource_id for the user",
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
	d.Set("resource_id", user.ResourceID)
	d.Set("console_access_denied", user.ConsoleAccessDenied)

	return diags
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ics.Client)

	user, err := c.Users.Create(ics.User{
		Name:              d.Get("name").(string),
		Username:          d.Get("username").(string),
		Email:             d.Get("email_address").(string),
		Password:          d.Get("password").(string),
		ConfirmPassword:   d.Get("password").(string),
		TwoFactorRequired: d.Get("two_factor_required").(bool),
		AccessLevel:       d.Get("access_level").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("user_id", user.ID)
	resourceUserRead(ctx, d, m)

	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ics.Client)

	d.Partial(true)

	if d.HasChanges("name", "email_address", "username", "access_level") {
		_, err := c.Users.UpdateUserInfo(d.Get("user_id").(int), d.Get("name").(string), d.Get("username").(string), d.Get("email_address").(string), d.Get("access_level").(string))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("two_factor_enabled") {
		if d.Get("two_factor_enabled").(bool) {
			err := c.Users.Disable2FA(d.Get("user_id").(int32))
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Cannot Enable 2FA",
				Detail:   "2FA can only be enabled by the user.  This attribute is only for disabling.",
			})
			return diags
		}
	}

	if d.HasChange("console_access_denied") {
		err := c.Users.SetConsoleAccess(d.Get("user_id").(int), d.Get("console_access_denied").(bool))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.Partial(false)

	resourceUserRead(ctx, d, m)

	return diags
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ics.Client)
	err := c.Users.Delete(d.Get("resource_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
