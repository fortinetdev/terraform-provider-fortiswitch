---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_multicastflow"
description: |-
  Multicast-flow configuration.
---

# fortiswitch_router_multicastflow
Multicast-flow configuration.

## Argument Reference

The following arguments are supported:

* `flows` - Config multicast-flow entries. The structure of `flows` block is documented below.
* `name` - Name.
* `comments` - Description/comments.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `flows` block supports:

* `group_addr` - Multicast group address.
* `source_addr` - Multicast source address.
* `id` - ID.
* `group_addr_end` - Multicast group end address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router MulticastFlow can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_multicastflow.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_multicastflow.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
