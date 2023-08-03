---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_communitylist"
description: |-
  Get information on an fortiswitch router communitylist.
---

# Data Source: fortiswitch_router_communitylist
Use this data source to get information on an fortiswitch router communitylist

## Argument Reference

* `name` - (Required) Specify the name of the desired router communitylist.

## Attribute Reference

The following attributes are exported:

* `type` - Community list type.
* `name` - Community list name.
* `rule` - Community list rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `action` - Community list action.
* `regexp` - An ordered list as a regular expression.
* `id` - Id.
* `match` - Community specification.

