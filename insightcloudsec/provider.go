package insightcloudsec

import (
	"context"
	"net/http"
	"regexp"

	ics "github.com/gstotts/insightcloudsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INSIGHTCLOUDSEC_BASE_URL", nil),
				Description: "An instance of InsightCloudSec to point at",
			},
			"apikey": {
				Type:         schema.TypeString,
				Required:     true,
				Sensitive:    true,
				DefaultFunc:  schema.EnvDefaultFunc("INSIGHTCLOUDSEC_API_KEY", nil),
				Description:  "ApiKey for use with InsightCloudSec API calls",
				ValidateFunc: validation.StringMatch(regexp.MustCompile("[A-Za-z0-9-_]{51}"), "API keys must only contain charachters a-z, A-Z, 0-9, hyphens and underscores"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"insightcloudsec_cloud": resourceCloud(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"insightcloudsec_cloud":       datasSourceCloud(),
			"insightcloudsec_cloud_types": dataSourceCloudTypes(),
			"insightcloudsec_users":       dataSourceUsers(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	url := d.Get("url").(string)
	apiKey := d.Get("apikey").(string)

	var diags diag.Diagnostics

	if (url != "") && (apiKey != "") {
		return &ics.Client{
			APIKey:     apiKey,
			BaseURL:    url,
			HttpClient: http.DefaultClient,
		}, diags
	}

	return &ics.Client{}, diags
}
