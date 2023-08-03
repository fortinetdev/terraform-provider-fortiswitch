---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_prefixlist6"
description: |-
  Get information on an fortiswitch router prefixlist6.
---

# Data Source: fortiswitch_router_prefixlist6
Use this data source to get information on an fortiswitch router prefixlist6

## Argument Reference

* `name` - (Required) Specify the name of the desired router prefixlist6.

## Attribute Reference

The following attributes are exported:

* `rule` - Rule. The structure of `rule` block is documented below.
* `name` - Name.
* `comments` - Description/comments.

The `rule` block contains:

* `le` - Maximum prefix length to be matched.
* `prefix6` - Prefix6.
* `ge` - Minimum prefix length to be matched.
* `flags` - Flags.
* `action` - Action.
* `id` - Rule id.

