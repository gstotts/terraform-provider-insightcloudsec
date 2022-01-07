package insightcloudsec

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var (
	// For use in ConflictsWith statements
	AZR_ORG_ONLY_ATTR = []string{}
	AWS_ORG_ONLY_ATTR = []string{"remove_suspended"}
	GCE_ORG_ONLY_ATTR = []string{"auto_add", "auto_badge"}

	// Combinations of Attributes
	AZR_AND_GCP_ORG_ATTR = append(AZR_ORG_ONLY_ATTR, GCE_ORG_ONLY_ATTR...)
	AZR_AND_AWS_ORG_ATTR = append(AZR_ORG_ONLY_ATTR, AWS_ORG_ONLY_ATTR...)
	AZR_AND_GCE_ORG_ATTR = append(AWS_ORG_ONLY_ATTR, GCE_ORG_ONLY_ATTR...)
)

func resourceCloudOrg() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudOrgCreate,
		ReadContext:   resourceCloudOrgRead,
		UpdateContext: resourceCloudOrgUpdate,
		DeleteContext: resourceCloudOrgDelete,
		Schema: map[string]*schema.Schema{
			"nickname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_type_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"AWS", "AZURE_ARM", "GCE"}, false),
			},
			"credentials": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"auto_add": {
				Type:          schema.TypeBool,
				Optional:      true,
				RequiredWith:  GCE_ORG_ONLY_ATTR,
				ConflictsWith: AZR_AND_AWS_ORG_ATTR,
			},
			"auto_badge": {
				Type:          schema.TypeBool,
				Optional:      true,
				RequiredWith:  GCE_ORG_ONLY_ATTR,
				ConflictsWith: AZR_AND_AWS_ORG_ATTR,
			},
			"auto_remove": {
				Type:          schema.TypeBool,
				Optional:      true,
				RequiredWith:  AZR_AND_GCE_ORG_ATTR,
				ConflictsWith: AWS_ORG_ONLY_ATTR,
			},
			"parent_folder_id": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"remove_suspended": {
				Type:          schema.TypeBool,
				Optional:      true,
				RequiredWith:  AWS_ORG_ONLY_ATTR,
				ConflictsWith: AZR_AND_GCE_ORG_ATTR,
			},
			"skip_prefixes": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"domain_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"organization_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"projects": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"failures": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"added_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCloudOrgCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceCloudOrgRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceCloudOrgUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceCloudOrgDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
