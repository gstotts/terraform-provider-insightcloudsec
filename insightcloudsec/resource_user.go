package insightcloudsec

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	ics "github.com/gstotts/insightcloudsec"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
			"access_level": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The access level to associate with the user",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"BASIC_USER", "ORGANIZATION_ADMIN", "DOMAIN_VIEWER", "DOMAIN_ADMIN"}, false)),
			},
			"organization_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The organization id to which the user belongs",
			},
			"organization_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The organization name to which the user belongs",
			},
			"console_access_denied": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Boolean representing if a user's console access is currently denied",
			},
			"temporary_pw": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Temporary password returned for resets or intial creation",
				Sensitive:   true,
			},
			"temp_pw_expiration": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time and date of temporary password expiration",
			},
		},
	}
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ics.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	user, err := c.Users.GetUserByID(id)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, fmt.Sprintf("\n\nUser Returned from API:\n%v", user))

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
	d.Set("organization_name", user.Org)
	d.Set("console_access_denied", user.ConsoleAccessDenied)
	d.SetId(strconv.Itoa(user.ID))
	return diags
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ics.Client)

	user, err := c.Users.Create(ics.User{
		Name:        d.Get("name").(string),
		Username:    d.Get("username").(string),
		Email:       d.Get("email_address").(string),
		AccessLevel: d.Get("access_level").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("temporary_pw", user.TempPassword)
	d.Set("temp_pw_expiration", user.TempPasswordExpiration)
	tflog.Debug(ctx, fmt.Sprintf("Retrieved Data: %v", user))
	d.SetId(strconv.Itoa(user.ID))
	resourceUserRead(ctx, d, m)

	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ics.Client)

	if d.HasChanges("name", "email_address", "username") {
		_, err := c.Users.UpdateUserInfo(d.Get("user_id").(int), d.Get("name").(string), d.Get("username").(string), d.Get("email_address").(string), d.Get("access_level").(string))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("access_level") {
		access_level := d.Get("access_level").(string)
		if access_level == "BASIC_USER" || access_level == "ORGANIZATION_ADMIN" {
			_, err := c.Users.UpdateUserInfo(d.Get("user_id").(int), d.Get("name").(string), d.Get("username").(string), d.Get("email_address").(string), d.Get("access_level").(string))
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			in_state, desired := d.GetChange("access_level")
			_, err := c.Users.EditAccessLevel(d.Get("user_id").(int), in_state.(string), desired.(string))
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if d.HasChange("console_access_denied") {
		err := c.Users.SetConsoleAccess(d.Get("user_id").(int), d.Get("console_access_denied").(bool))
		if err != nil {
			return diag.FromErr(err)
		}
	}

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
