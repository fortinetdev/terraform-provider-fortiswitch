---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_automationdestination"
description: |-
  Get information on an fortiswitch system automationdestination.
---

# Data Source: fortiswitch_system_automationdestination
Use this data source to get information on an fortiswitch system automationdestination

## Argument Reference

* `name` - (Required) Specify the name of the desired system automationdestination.

## Attribute Reference

The following attributes are exported:

* `ha_group_id` - Cluster group ID set for this destination (default = 0).
* `destination` - Destinations. The structure of `destination` block is documented below.
* `type` - Destination type.
* `name` - Name.

The `destination` block contains:

* `name` - Destination.

