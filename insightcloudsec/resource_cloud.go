package insightcloudsec

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	ics "github.com/gstotts/insightcloudsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var (
	// For use in ConflictsWith statements
	AZR_ONLY_ATTR = []string{"tenant_id", "subscription_id", "app_id"}
	AWS_ONLY_ATTR = []string{"account", "role_arn", "duration", "external_id", "session_name", "secret_key"}
	GCE_ONLY_ATTR = []string{"project", "api_credentials"}

	// Combinations of Attributes
	AZR_AND_GCP_ATTR = append(AZR_ONLY_ATTR, GCE_ONLY_ATTR...)
	AZR_AND_AWS_ATTR = append(AZR_ONLY_ATTR, AWS_ONLY_ATTR...)
	AZR_AND_GCE_ATTR = append(AWS_ONLY_ATTR, GCE_ONLY_ATTR...)
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
			"creation_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
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
			"cloud_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"AWS", "AZURE_ARM", "GCE"}, false),
			},
			"tenant_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: AZR_AND_GCE_ATTR,
				RequiredWith:  AZR_ONLY_ATTR,
			},
			"app_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: AZR_AND_GCE_ATTR,
				RequiredWith:  AZR_ONLY_ATTR,
			},
			"subscription_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: AZR_AND_GCE_ATTR,
				RequiredWith:  AZR_ONLY_ATTR,
			},
			// Used in multiple clouds so must careful use ConflictsWith and RequiredWith
			"api_key": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{},
				RequiredWith:  []string{},
			},
			"account": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: AZR_AND_GCP_ATTR,
				RequiredWith:  AWS_ONLY_ATTR,
			},
			"authentication_type": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: AZR_AND_GCP_ATTR,
				RequiredWith:  AWS_ONLY_ATTR,
				ValidateFunc:  validation.StringInSlice([]string{"assume_role", "instance_assume_role"}, false),
			},
			"role_arn": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: AZR_AND_GCP_ATTR,
				RequiredWith:  AWS_ONLY_ATTR,
			},
			// Not required for use with assume_role authentication method
			"secret_key": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: AZR_AND_GCP_ATTR,
			},
			"duration": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: AZR_AND_GCP_ATTR,
			},
			"external_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: AZR_AND_GCP_ATTR,
				RequiredWith:  AWS_ONLY_ATTR,
			},
			"session_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: AZR_AND_GCP_ATTR,
				RequiredWith:  AWS_ONLY_ATTR,
			},
			"api_credentials": {
				Type:          schema.TypeSet,
				Optional:      true,
				ConflictsWith: AZR_AND_AWS_ATTR,
				RequiredWith:  GCE_ONLY_ATTR,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "service_account",
						},
						"project_id": {
							Type:         schema.TypeString,
							Optional:     true,
							RequiredWith: GCE_ONLY_ATTR,
						},
						"private_key_id": {
							Type:         schema.TypeString,
							Optional:     true,
							RequiredWith: GCE_ONLY_ATTR,
						},
						"private_key": {
							Type:         schema.TypeString,
							Optional:     true,
							Sensitive:    true,
							RequiredWith: GCE_ONLY_ATTR,
						},
						"client_email": {
							Type:         schema.TypeString,
							Optional:     true,
							RequiredWith: GCE_ONLY_ATTR,
						},
						"auth_uri": {
							Type:         schema.TypeString,
							Optional:     true,
							RequiredWith: GCE_ONLY_ATTR,
						},
						"token_uri": {
							Type:         schema.TypeString,
							Optional:     true,
							RequiredWith: GCE_ONLY_ATTR,
						},
						"auth_provider_x509_cert_url": {
							Type:         schema.TypeString,
							Optional:     true,
							RequiredWith: GCE_ONLY_ATTR,
						},
						"client_x509_cert_url": {
							Type:         schema.TypeString,
							Optional:     true,
							RequiredWith: GCE_ONLY_ATTR,
						},
					},
				},
			},
			"project": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: AZR_AND_AWS_ATTR,
				RequiredWith:  GCE_ONLY_ATTR,
			},
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
		Name:      d.Get("name").(string),
		CloudType: d.Get("cloud_type").(string),
	}

	// Azure Cloud Accounts
	if params.CloudType == "AZURE_ARM" {
		params.AuthType = "standard"
		params.TenantID = d.Get("tenant_id").(string)
		params.AppID = d.Get("app_id").(string)
		params.SubscriptionID = d.Get("subscription_id").(string)
		params.ApiKeyOrCert = d.Get("api_key").(string)

		cloud, err = c.AddAzureCloud(ics.AzureCloudAccount{CreationParameters: params})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Error Adding Azure Cloud",
				Detail: fmt.Sprintf("%s\n%s\n\n%s\n%s",
					"An error was returned when attempting to add an Azure Cloud to InsightCloudSec.",
					"This could be the result of an incorrect tenant_id, subscription_id or app_id.",
					"Error from API:", err),
			})
			return diags
		}

		log.Println("[DEBUG] Azure Cloud Returned from API: \n%", cloud)
	}

	// AWS Cloud Accounts
	if params.CloudType == "AWS" {
		params.RoleArn = d.Get("role_arn").(string)
		params.Duration = d.Get("duration").(int)
		params.SessionName = d.Get("session_name").(string)
		params.ExternalID = d.Get("external_id").(string)

		auth_type := strings.ToLower(d.Get("authentication_type").(string))
		params.AuthType = auth_type

		if auth_type == "assume_role" {
			// AWS STS Assume Role (Instance Assume does not require)
			params.ApiKeyOrCert = d.Get("api_key").(string)
			params.SecretKey = d.Get("secret_key").(string)
			log.Println("[DEBUG] Setting up Assume Role for: ", params.Name)
		} else if auth_type != "instance_assume_role" {
			return diag.FromErr(fmt.Errorf("[ERROR] Invalid authentication type,  must be assume_role or instance_assume_role for AWS clouds"))
		}

		cloud, err = c.AddAWSCloud(ics.AWSCloudAccount{CreationParameters: params})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Error Adding AWS Cloud",
				Detail: fmt.Sprintf("%s\n%s\n\n%s\n%s",
					"An error was returned when attempting to add an AWS Cloud to InsightCloudSec.",
					"This could be the result of an incorrect role_arn, api_key, secret_key, etc.",
					"Error from API:", err),
			})
			return diags
		}

		log.Println("[DEBUG] AWS Cloud Returned from API: \n%", cloud)
	}

	// GCE Cloud Accounts

	if params.CloudType == "GCE" {
		params.GCPAuth.Type = d.Get("api_credentials.type").(string)
		params.GCPAuth.ProjectID = d.Get("api_credentials.project_id").(string)
		params.GCPAuth.PrivateKeyID = d.Get("api_credentials.private_key_id").(string)
		params.GCPAuth.PrivateKey = d.Get("api_credentials.private_key").(string)
		params.GCPAuth.ClientEmail = d.Get("api_credentials.client_email").(string)
		params.GCPAuth.ClientID = d.Get("api_credentials.client_id").(string)
		params.GCPAuth.AuthURI = d.Get("api_credentials.auth_uri").(string)
		params.GCPAuth.TokenURI = d.Get("api_credentials.token_uri").(string)
		params.GCPAuth.AuthProviderx509CertURL = d.Get("api_credentials.auth_provider_x509_cert_url").(string)
		params.GCPAuth.Clientx509CertUrl = d.Get("api_credentials.client_x509_cert_url").(string)

		cloud, err = c.AddGCPCloud(ics.GCPCloudAccount{CreationParameters: params})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Error Adding GCP Cloud",
				Detail: fmt.Sprintf("%s\n%s\n\n%s\n%s",
					"An error was returned when attempting to add a GCP Cloud to InsightCloudSec.",
					"This could be the result of an incorrect credentials or project settings.",
					"Error from API:", err),
			})
			return diags
		}

		log.Println("[DEBUG] GCP Cloud Returned from API: \n%", cloud)
	}

	d.SetId(strconv.Itoa(cloud.ID))
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
	d.Set("account", cloud.AccountID)
	d.Set("resource_id", cloud.ResourceID)
	d.Set("group_resource_id", cloud.GroupResourceID)
	d.Set("status", cloud.Status)
	d.Set("creation_time", cloud.Created)
	d.Set("strategy_id", cloud.StrategyID)
	d.Set("cloud_type", cloud.CloudTypeID)

	return diags
}

func resourceCloudUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ics.Client)

	// Common Parameters
	params := ics.CloudAccountParameters{
		Name:      d.Get("name").(string),
		CloudType: d.Get("cloud_type").(string),
	}

	// Azure Cloud Accounts Parameters
	if params.CloudType == "AZURE_ARM" {
		params.AuthType = "standard"
		params.TenantID = d.Get("tenant_id").(string)
		params.AppID = d.Get("app_id").(string)
		params.SubscriptionID = d.Get("subscription_id").(string)
		params.ApiKeyOrCert = d.Get("api_key").(string)
	}

	// AWS Cloud Account Parameters
	if params.CloudType == "AWS" {
		params.RoleArn = d.Get("aws.0.role_arn").(string)
		params.Duration = d.Get("aws.0.duration").(int)
		params.SessionName = d.Get("aws.0.session_name").(string)
		params.ExternalID = d.Get("aws.0.external_id").(string)
		params.CloudType = "AWS"

		auth_type := strings.ToLower(d.Get("authentication_type").(string))
		params.AuthType = auth_type

		if auth_type == "assume_role" {
			// AWS STS Assume Role (Instance Assume does not require)
			params.ApiKeyOrCert = d.Get("api_key").(string)
			params.SecretKey = d.Get("secret_key").(string)
		} else if auth_type != "instance_assume_role" {
			return diag.FromErr(fmt.Errorf("[ERROR] Invalid authentication type,  must be assume_role or instance_assume_role for AWS clouds"))
		}
	}

	// GCE Cloud Accounts
	if params.CloudType == "GCE" {
		params.GCPAuth.Type = d.Get("api_credentials.type").(string)
		params.GCPAuth.ProjectID = d.Get("api_credentials.project_id").(string)
		params.GCPAuth.PrivateKeyID = d.Get("api_credentials.private_key_id").(string)
		params.GCPAuth.PrivateKey = d.Get("api_credentials.private_key").(string)
		params.GCPAuth.ClientEmail = d.Get("api_credentials.client_email").(string)
		params.GCPAuth.ClientID = d.Get("api_credentials.client_id").(string)
		params.GCPAuth.AuthURI = d.Get("api_credentials.auth_uri").(string)
		params.GCPAuth.TokenURI = d.Get("api_credentials.token_uri").(string)
		params.GCPAuth.AuthProviderx509CertURL = d.Get("api_credentials.auth_provider_x509_cert_url").(string)
		params.GCPAuth.Clientx509CertUrl = d.Get("api_credentials.client_509x_cert_url").(string)
	}

	id, _ := strconv.Atoi(d.Id())
	log.Println("[DEBUG] Updating Cloud ID: ", id)
	_, err := c.UpdateCloud(id, params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("last_updated", time.Now().Format(time.RFC850))
	resourceCloudRead(ctx, d, m)

	return diags
}

func resourceCloudDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ics.Client)
	var diags diag.Diagnostics

	err := c.DeleteCloud(d.Get("resource_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}