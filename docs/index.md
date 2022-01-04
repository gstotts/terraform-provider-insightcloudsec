---
page_title: "insightcloudsec Provider"
subcategory: ""
description: |-
    The InsightCloudSec provider is used to interact with resources supported by Rapid7's InsightCloudSec.  
This provider is not maintained or created directly by Rapid7 but by a community of product users that wish to see further
integration with Infrastructure as Code and the tool.  The provider must be configured with the proper credentials before it 
can be used.  
---

# InsightCloudSec Provider

The InsightCloudSec provider is used to interact with resources supported by Rapid7's InsightCloudSec.  This provider is not maintained or created directly by Rapid7 but by a community of product users that wish to see further integration with Infrastructure as Code and the tool.  The provider must be configured with the proper credentials before it can be used.  

Currently the provider is in intial creation and testing and, as such, not recommended for production workloads at this time.


## Example Usage

```hcl
terraform {
    required_providers {
        insightcloudsec = {
            source = "gstotts/insightcloudsec"
            version = "0.1.0"
        }
    }
}

provider "insightcloudsec" {
    url     = var.insightcloudsec_url
    apikey  = var.insightcloudsec_api_key
}

# Create a cloud
resource "insightcloudsec_cloud" "my_cloud" {
    # ...
}
```

## Argument Reference

The following arguments are supported:

* `apikey` -  (Optional) Api-Key for use with InsightCloudSec API calls.  This can also be specified  with the `INSIGHTCLOUDSEC_API_KEY` environment variable.
* `url` - (Optional) The specific base url path for the InsightCloudSec in use.  This can also be specified with the `INSIGHTCLOUDSEC_BASE_URL` environment variable.
