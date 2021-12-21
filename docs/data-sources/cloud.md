---
page_title: "insightcloudsec_cloud Data Source - terraform-provider-insightcloudsec"
subcategory: ""
description: |-
  Gets information on an InsightCloudSec cloud account.
---

# insightcloudsec_cloud (Data Source)

This data source retrieves a specific InsightCloudSec cloud account with the given name.  


## Example Usage

```hcl
# Look up a cloud named "My Cloud 123"
data "insightcloudsec_cloud" "my_cloud" {
    name = "My Cloud 123"
}
```

## Argument Reference

The following arguments are supported:
* `name` - (Required) The name of the specific cloud account to lookup.

## Attributes Reference

* `id` - The ID of the resource as an integer.
* `name` - The name of the cloud account.
* `account` - The account or tenant ID associated with the given cloud account.
* `cloud_type` - Type of cloud associated with the account. For example, `AWS`, `AZURE_ARM`, `GCE`, etc.
* `group_resource_id` - The group resource ID.
* `resource_id` - The resource ID.
* `strategy_id` - The ID for the harevesting strategy in use.


