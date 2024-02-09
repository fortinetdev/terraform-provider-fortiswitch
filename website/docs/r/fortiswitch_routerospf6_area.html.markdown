---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_routerospf6_area"
description: |-
  OSPF6 area configuration.
---

# fortiswitch_routerospf6_area
OSPF6 area configuration.

~> The provider supports the definition of Area in Router Ospf6 `fortiswitch_router_ospf6`, and also allows the definition of separate Area resources `fortiswitch_routerospf6_area`, but do not use a `fortiswitch_router_ospf6` with in-line Area in conjunction with any `fortiswitch_routerospf6_area` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

* `stub_type` - Stub summary setting. Valid values: `no-summary`, `summary`.
* `filter_list` - OSPF area filter-list configuration. The structure of `filter_list` block is documented below.
* `range` - OSPF6 area range configuration. The structure of `range` block is documented below.
* `type` - Area type setting. Valid values: `regular`, `stub`.
* `fswid` - Area entry ip address.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `filter_list` block supports:

* `direction` - Direction. Valid values: `in`, `out`.
* `list` - Access-list or prefix-list name.
* `id` - Filter list entry ID.

The `range` block supports:

* `advertise` - Enable/disable advertise status. Valid values: `disable`, `enable`.
* `id` - Range entry id.
* `prefix6` - <prefix6>   IPv6 prefix


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

Routerospf6 Area can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_routerospf6_area.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_routerospf6_area.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
