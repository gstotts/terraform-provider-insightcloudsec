---
page_title: "insightcloudsec_cloud Resource - terraform-provider-insightcloudsec"
subcategory: ""
description: |-
  
---

# insightcloudsec_cloud (Resource)



## Example Usage

```hcl
# AWS Cloud
resource "insightcloudsec_cloud" "my_aws_cloud" {
    name                = "My AWS Cloud"
    cloud_type          = "AWS"
    authentication_type = "assume_role"
    account             = "123123123123"
    role_arn            = "arn:aws:iam::123123123123:role/insightcloudsec_sts_assume_role"
    api_key             = var.my_aws_api_key
    secret_key          = var.my_aws_secret_key
    duration            = 3600
    session_name        = "InsightCloudSec-Role"  // Name for Cloudtrail logs
    external_id         = "MYAWS" // Optional name for external identities.
}

# Azure Cloud
resource "insightcloudsec_cloud" "my_azure_cloud" {
    name            = "My Azure Cloud"
    cloud_type      = "AZURE_ARM"
    tenant_id       = var.my_azure_tenant_id
    subscription_id = var.my_azure_subscription_id
    app_id          = var.my_azure_app_id
    api_key         = var.my_azure_api_key
}

# GCP Cloud
resource "insightcloudsec_cloud" "my_gcp_cloud" {
    name            = "My GCP Cloud"
    cloud_type      = "GCE"
    project         = "my_gcp_project_name"
    api_credentials {
        type            = "service_account"
        project_id      = "my_gcp_project_name"
        private_key_id  = var.my_gcp_private_key_id
        private_key     = var.my_gcp_private_key
        client_email    = "myadmin@bond-007.iam.gserviceaccount.com"
        client_id       = var.my_gcp_client_id
        auth_uri        = "https://accounts.google.com/o/oauth2/auth"
        token_uri       = "https://accounts.google.com/o/oauth2/token"

        auth_provider_x509_cert_url = "https://www.googleapis.com/oauth2/v1/certs"
        client_x509_cert_url = "https://www.googleapis.com/robot/v1/metadata/x509/myadmin%4-my_gcp_project_name.iam.gserviceaccount.com"
    }
}
```


## Argument Reference

Shared arguments for all cloud types:
* `name` - (Required) Name for cloud account.
* `cloud_type` - (Required) Must be one of `AWS`, `AZURE_ARM`, or `GCE` as determined by the cloud being utilized.

Some arguments vary by supported cloud type:

### AWS Argument Reference
* `account` (Required) AWS account number. 
* `authentication_type` - (Required) Type of authentication to use.  Must be either `assume_role` or `instance_assume_role`.
* `role_arn` - (Required) Role ARN of the role to be asssumed.  
* `api_key` - (Required) Api-Key for use with AWS APIs.
* `secret_key` - The secret key that corresponds to the api-key if using `assume_role`.
* `duration` - 
* `external_id` - (Optional)
* `session_name` - (Required) Name to use for the session in Cloudtrail logs.

### Azure Argument Reference
* `tenant_id` - (Required) Azure tenant ID.
* `subscription_id` - (Required) Azure subscription ID.
* `app_id` - (Required) Azure application ID.
* `api_key` - (Required) Api-Key for use with Azure.

### GCP Argument Reference
* `project` - 
* ``

## Attributes Reference
