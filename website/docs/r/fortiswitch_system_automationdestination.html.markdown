---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_automationdestination"
description: |-
  Automation destinations.
---

# fortiswitch_system_automationdestination
Automation destinations. Applies to FortiSwitch Version `>= 7.2.0`.

## Argument Reference

The following arguments are supported:

* `ha_group_id` - Cluster group ID set for this destination (default = 0).
* `destination` - Destinations. The structure of `destination` block is documented below.
* `type` - Destination type. Valid values: `fortigate`, `ha-cluster`.
* `name` - Name.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `destination` block supports:

* `name` - Destination.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationDestination can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_automationdestination.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_automationdestination.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
