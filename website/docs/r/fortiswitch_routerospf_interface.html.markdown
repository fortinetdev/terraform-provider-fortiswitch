---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_routerospf_interface"
description: |-
  OSPF interface configuration.
---

# fortiswitch_routerospf_interface
OSPF interface configuration.

~> The provider supports the definition of Interface in Router Ospf `fortiswitch_router_ospf`, and also allows the definition of separate Interface resources `fortiswitch_routerospf_interface`, but do not use a `fortiswitch_router_ospf` with in-line Interface in conjunction with any `fortiswitch_routerospf_interface` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

* `priority` - Router priority.
* `authentication_key` - Authentication key.
* `name` - Interface entry name.
* `dead_interval` - Dead interval. For fast-hello assign value 1.
* `hello_multiplier` - Number of hello packets within dead interval.Valid only for fast-hello.
* `bfd` - Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
* `transmit_delay` - Link state transmit delay.
* `ucast_ttl` - Unicast TTL.
* `mtu` - Interface MTU.
* `retransmit_interval` - Time between retransmitting lost link state advertisements.
* `authentication` - Authentication type. Valid values: `none`, `text`, `md5`.
* `cost` - Cost of the interface.
* `hello_interval` - Hello interval.
* `ttl` - TTL.
* `md5_keys` - OSPF md5 key configuration. Applicable only when authentication field is set to md5. The structure of `md5_keys` block is documented below.
* `mtu_ignore` - Disable MTU mismatch detection on this interface. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `md5_keys` block supports:

* `id` - key-id (1-255).
* `key` - md5-key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Routerospf Interface can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_routerospf_interface.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_routerospf_interface.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
