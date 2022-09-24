---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "insightcloudsec_custom_insight Resource - terraform-provider-insightcloudsec"
subcategory: ""
description: |-
  Provides details on a custom_insight configuration for InsightCloudSec.
---

# insightcloudsec_custom_insight

Provides details on a custom insight configuration for InsightCloudSec.  This is only for customer created insights -- built-in insights cannot be managed.

## Example Usage
```terraform
resource insightcloudsec_custom_insight "my-insight" {
    name = "My Custom Insight"
    description = "This is my custom insight"
    severity = 1
    resource_types = ["divvyorganizationservice"]
    filter {
            name = "divvy.filter.cloud_trail_in_all_regions"
    }
    
    scopes = ["divvyorganizationservice:0"]
    tags = ["Test"]
    
    badge {
        key = "cloud_org_path"
        value = "/Root"
    }
    badge_filter_operator = "OR"
}
```

## Argument Reference

The following arguments are supported: 

- `filter` (Block List, Min: 1) Filter used with the insight to determine resources (see [below for nested schema](#nestedblock--filter))
- `name` (String) The name of the insight for display in InsightCloudSec
- `resource_types` (List of String) Resource types the insight applies to
- `severity` (Number) The severity associated with the insight represented by an int

### Optional

- `badge` (Block List) Badges used to limit the insight (see [below for nested schema](#nestedblock--badge))
- `badge_filter_operator` (String) The badge filter operator for the insight
- `description` (String) The description to assign to the insight
- `scopes` (List of String) The scope for the insight
- `tags` (List of String) Tags to associate with the insight

### Read-Only

- `id` (String) The ID assigned to the insight.

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Required:

- `name` (String) The name of the filter

Optional:

- `collections` (Map of String) The collections associated with the filter
- `config` (Map of String) The configuration of the filter


<a id="nestedblock--badge"></a>
### Nested Schema for `badge`

Required:

- `key` (String) Key for the badge
- `value` (String) Value for the badge

