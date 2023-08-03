---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_setting"
description: |-
  Get information on fortiswitch router setting.
---

# Data Source: fortiswitch_router_setting
Use this data source to get information on fortiswitch router setting

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `filter_list` - Route filter list configuration. The structure of `filter_list` block is documented below.
* `hostname` - Set hostname for this virtual domain router.

The `filter_list` block contains:

* `route_map` - Route map name.
* `protocol` - Protocol type.
* `id` - Filter list entry ID.

