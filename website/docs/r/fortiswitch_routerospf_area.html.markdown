---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_routerospf_area"
description: |-
  OSPF area configuration.
---

# fortiswitch_routerospf_area
OSPF area configuration.

~> The provider supports the definition of Area in Router Ospf `fortiswitch_router_ospf`, and also allows the definition of separate Area resources `fortiswitch_routerospf_area`, but do not use a `fortiswitch_router_ospf` with in-line Area in conjunction with any `fortiswitch_routerospf_area` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

* `nssa_translator_role` - NSSA translator role type. Valid values: `candidate`, `never`, `always`.
* `virtual_link` - OSPF virtual link configuration. The structure of `virtual_link` block is documented below.
* `shortcut` - Enable/disable shortcut option. Valid values: `disable`, `enable`, `default`.
* `stub_type` - Stub summary setting. Valid values: `no-summary`, `summary`.
* `range` - OSPF area range configuration. The structure of `range` block is documented below.
* `filter_list` - OSPF area filter-list configuration. The structure of `filter_list` block is documented below.
* `default_cost` - Summary default cost of stub or NSSA area.
* `type` - Area type setting. Valid values: `regular`, `nssa`, `stub`.
* `fswid` - Area entry IP address.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `virtual_link` block supports:

* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `name` - Virtual link entry name.
* `transmit_delay` - Link state transmit delay.
* `retransmit_interval` - Time between retransmitting lost link state advertisements.
* `authentication` - Authentication type. Valid values: `none`, `text`, `md5`.
* `peer` - Peer IP.
* `md5_keys` - OSPF md5 key configuration. Applicable only when authentication field is set to md5. The structure of `md5_keys` block is documented below.
* `authentication_key` - Authentication key.

The `md5_keys` block supports:

* `id` - key-id (1-255).
* `key` - md5-key.

The `range` block supports:

* `substitute` - Substitute prefix.
* `prefix` - Prefix.
* `substitute_status` - Enable/disable substitute status. Valid values: `enable`, `disable`.
* `id` - Range entry ID.
* `advertise` - Enable/disable advertise status. Valid values: `disable`, `enable`.

The `filter_list` block supports:

* `direction` - Direction. Valid values: `in`, `out`.
* `list` - Access-list or prefix-list name.
* `id` - Filter list entry ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

Routerospf Area can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_routerospf_area.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_routerospf_area.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
