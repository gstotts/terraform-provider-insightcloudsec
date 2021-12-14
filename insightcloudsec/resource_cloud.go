package insightcloudsec

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	ics "github.com/gstotts/insightcloudsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloud() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudCreate,
		ReadContext:   resourceCloudRead,
		UpdateContext: resourceCloudUpdate,
		DeleteContext: resourceCloudDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"creation_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"strategy_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"azure": {
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"aws"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tenant_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"app_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"subscription_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"api_key": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
					},
				},
			},
			"aws": {
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"azure"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"assume_role", "instance_assume_role"}, false),
						},
						"role_arn": {
							Type:     schema.TypeString,
							Required: true,
						},
						"api_key": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"secret_key": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"duration": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  3600,
						},
						"external_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"session_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			// "gce": {
			// 	Type:          schema.TypeList,
			// 	Optional:      true,
			// 	ConflictsWith: []string{"azure", "aws"},
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"project": {
			// 				Type: schema.TypeString,
			// 				Required: true,
			// 			},
			// 			"": {},
			// 		},
			// 	},
			// },
			//},
		},
	}
}

func resourceCloudCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)
	var diags diag.Diagnostics
	var err error
	var cloud ics.Cloud

	// Common Parameters
	params := ics.CloudAccountParameters{
		Name:          d.Get("name").(string),
		AccountNumber: d.Get("account_id").(string),
	}

	// Azure Cloud Accounts
	_, azureOk := d.GetOk("azure")
	if azureOk {
		params.CloudType = "AZURE_ARM"
		params.AuthType = "standard"
		params.TenantID = d.Get("azure.0.tenant_id").(string)
		params.AppID = d.Get("azure.0.app_id").(string)
		params.SubscriptionID = d.Get("azure.0.subscription_id").(string)
		params.ApiKeyOrCert = d.Get("azure.0.api_key").(string)

		cloud, err = c.AddAzureCloud(ics.AzureCloudAccount{CreationParameters: params})
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// AWS Cloud Accounts
	_, awsOk := d.GetOk("aws")
	if awsOk {
		params := ics.CloudAccountParameters{
			RoleArn:     d.Get("aws.0.role_arn").(string),
			Duration:    d.Get("aws.0.duration").(int),
			SessionName: d.Get("aws.0.session_name").(string),
			ExternalID:  d.Get("aws.0.external_id").(string),
			CloudType:   "AWS",
		}

		auth_type := strings.ToLower(d.Get("aws.0.authentication_type").(string))
		params.AuthType = auth_type

		if auth_type == "assume_role" {
			// AWS STS Assume Role (Instance Assume does not require)
			params.ApiKeyOrCert = d.Get("aws.0.api_key").(string)
			params.SecretKey = d.Get("aws.0.secret_key").(string)
		} else if auth_type != "instance_assume_role" {
			return diag.FromErr(fmt.Errorf("[ERROR] Invalid authentication type,  must be assume_role or instance_assume_role for AWS clouds"))
		}

		cloud, err = c.AddAWSCloud(ics.AWSCloudAccount{CreationParameters: params})
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// GCE Cloud Accounts

	d.SetId(strconv.Itoa(cloud.ID))
	d.Set("resource_id", cloud.ResourceID)
	d.Set("group_resource_id", cloud.GroupResourceID)
	d.Set("status", cloud.Status)
	d.Set("creation_time", cloud.Created)
	d.Set("org_resource_id", cloud.CloudOrgID)
	d.Set("strategy_id", cloud.StrategyID)

	resourceCloudRead(ctx, d, m)

	return diags
}

func resourceCloudRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)
	var diags diag.Diagnostics

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	cloud, err := c.GetCloudByID(id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", cloud.Name)
	d.Set("account_id", cloud.AccountID)
	d.Set("resource_id", cloud.ResourceID)
	d.Set("group_resource_id", cloud.GroupResourceID)
	d.Set("status", cloud.Status)
	d.Set("creation_time", cloud.Created)
	d.Set("strategy_id", cloud.StrategyID)
	d.Set("cloud_type", cloud.CloudTypeID)

	if cloud.CloudTypeID == "AWS" {
		d.Set("role_arn", cloud.RoleARN)
	}
	return diags
}

func resourceCloudUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceCloudDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
