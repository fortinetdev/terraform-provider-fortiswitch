---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_accesslist"
description: |-
  Get information on an fortiswitch router accesslist.
---

# Data Source: fortiswitch_router_accesslist
Use this data source to get information on an fortiswitch router accesslist

## Example Usage

```hcl
data "fortiswitch_router_accesslist" sample1 {
  name = "test"
}

output output1 {
  value = data.fortiswitch_router_accesslist.sample1
}
```

## Argument Reference

* `name` - (Required) Specify the name of the desired router accesslist.

## Attribute Reference

The following attributes are exported:

* `rule` - Rule. The structure of `rule` block is documented below.
* `name` - Name. For creating Cisco wild-card access-list, name should be a digit within range <1-99>.
* `comments` - Description/comments.

The `rule` block contains:

* `exact_match` - Exact match.
* `prefix` - Prefix to define regular filter criteria, such as any or subnets.
* `flags` - Flags.
* `wildcard` - Wildcard to define cisco-style wildcard filter criteria.
* `action` - Action.
* `id` - Rule id.

