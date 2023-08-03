---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_zone"
description: |-
  Get information on an fortiswitch system zone.
---

# Data Source: fortiswitch_system_zone
Use this data source to get information on an fortiswitch system zone

## Argument Reference

* `name` - (Required) Specify the name of the desired system zone.

## Attribute Reference

The following attributes are exported:

* `interface` - Interfaces belong to the zone. The structure of `interface` block is documented below.
* `intrazone` - Intra-zone traffic.
* `name` - Zone name.

The `interface` block contains:

* `interface_name` - Interface name.

