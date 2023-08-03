---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchqos_ipdscpmap"
description: |-
  QOS IP precedence/DSCP configuration.
---

# fortiswitch_switchqos_ipdscpmap
QOS IP precedence/DSCP configuration.

## Argument Reference

The following arguments are supported:

* `map` - Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
* `name` - DSCP map name.
* `description` - Description of the map name.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `map` block supports:

* `diffserv` - Differentiated service. Valid values: `CS0`, `CS1`, `AF11`, `AF12`, `AF13`, `CS2`, `AF21`, `AF22`, `AF23`, `CS3`, `AF31`, `AF32`, `AF33`, `CS4`, `AF41`, `AF42`, `AF43`, `CS5`, `EF`, `CS6`, `CS7`.
* `cos_queue` - COS queue number.
* `value` - Raw values of DSCP (0 - 63).
* `ip_precedence` - IP precedence. Valid values: `Network-Control`, `Internetwork-Control`, `Critic/ECP`, `FlashOverride`, `Flash`, `Immediate`, `Priority`, `Routine`.
* `entry_name` - Mapping entry.
* `type` - type


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchQos IpDscpMap can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchqos_ipdscpmap.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchqos_ipdscpmap.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
