---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_prefixlist"
description: |-
  Get information on an fortiswitch router prefixlist.
---

# Data Source: fortiswitch_router_prefixlist
Use this data source to get information on an fortiswitch router prefixlist

## Argument Reference

* `name` - (Required) Specify the name of the desired router prefixlist.

## Attribute Reference

The following attributes are exported:

* `rule` - Rule. The structure of `rule` block is documented below.
* `name` - Name.
* `comments` - Description/comments.

The `rule` block contains:

* `le` - Maximum prefix length to be matched.
* `prefix` - Prefix.
* `ge` - Minimum prefix length to be matched.
* `flags` - Flags.
* `action` - Action.
* `id` - Rule id.

