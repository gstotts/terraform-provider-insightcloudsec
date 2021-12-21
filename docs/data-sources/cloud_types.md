---
page_title: "insightcloudsec_cloud_types Data Source - terraform-provider-insightcloudsec"
subcategory: ""
description: |-
  Gets information about the types of clouds available.
---

# insightcloudsec_cloud_types (Data Source)

This data source retrieves a list of all available cloud types for use.  Note that some of the returned types are not able to be configured by the API at this time per Rapid7's documentation.

## Example Usage

```hcl
data "insightcloudsec_cloud_types" "all_cloud_types" {}
```

## Argument Reference

There are currently no supported arguments.

## Attributes Reference

* `id` - Unique ID assigned to list when queried.
* `name` - The full name for the given cloud.
* `cloud_type` - The given identifier for the cloud type. For example, `AWS`, `AZURE`, etc.
* `cloud_access` - The availability of the cloud - either `private` or `public`.

