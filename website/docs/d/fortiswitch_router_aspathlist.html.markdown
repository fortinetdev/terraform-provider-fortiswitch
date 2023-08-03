---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_aspathlist"
description: |-
  Get information on an fortiswitch router aspathlist.
---

# Data Source: fortiswitch_router_aspathlist
Use this data source to get information on an fortiswitch router aspathlist

## Argument Reference

* `name` - (Required) Specify the name of the desired router aspathlist.

## Attribute Reference

The following attributes are exported:

* `name` - AS path list name.
* `rule` - AS path list rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `action` - AS path list action.
* `regexp` - Regular-expression to match the BGP AS paths.
* `id` - Id.

