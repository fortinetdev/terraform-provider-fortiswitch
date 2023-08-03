---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_accesslist6"
description: |-
  Get information on an fortiswitch router accesslist6.
---

# Data Source: fortiswitch_router_accesslist6
Use this data source to get information on an fortiswitch router accesslist6

## Argument Reference

* `name` - (Required) Specify the name of the desired router accesslist6.

## Attribute Reference

The following attributes are exported:

* `rule` - Rule. The structure of `rule` block is documented below.
* `name` - Name.
* `comments` - Description/comments.

The `rule` block contains:

* `action` - Action.
* `exact_match` - Exact match.
* `flags` - Flags.
* `id` - Rule id.
* `prefix6` - IPv6 prefix to define regular filter criteria, such as any or X:X::X:X/M.

