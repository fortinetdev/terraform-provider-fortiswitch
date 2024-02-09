---
subcategory: "FortiSwitch Switch-Controller"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchcontroller_global"
description: |-
  Switch-controller global configuration.
---

# fortiswitch_switchcontroller_global
Switch-controller global configuration.

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `ac_data_port` - Switch controller data port [1024, 49150].
* `echo_interval` - Interval before SWTP sends Echo Request after joining AC. [1, 600] default = 30s.
* `ac_dhcp_option_code` - DHCP option code for CAPUTP AC.
* `max_discoveries` - The maximum # of Discovery Request messages every round.
* `ac_list` - AC list. The structure of `ac_list` block is documented below.
* `location` - Location.
* `ac_discovery_mc_addr` - Discovery multicast address
* `ac_port` - Switch controller ctl port [1024, 49150].
* `max_retransmit` - The maximum # of retransmissions for tunnel packet.
* `ac_discovery_type` - AC discovery type.
* `mgmt_mode` - FortiLink management mode. Valid values: `capwap`, `https`.
* `tunnel_mode` - Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `ac_list` block supports:

* `ipv6_address` - IPv6 address.
* `ipv4_address` - IP addr.
* `id` - Id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController Global can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchcontroller_global.labelname SwitchControllerGlobal

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchcontroller_global.labelname SwitchControllerGlobal
$ unset "FORTISWITCH_IMPORT_TABLE"
```
