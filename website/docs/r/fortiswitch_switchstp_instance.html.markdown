---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchstp_instance"
description: |-
  Stp instances.
---

# fortiswitch_switchstp_instance
Stp instances.

## Argument Reference

The following arguments are supported:

* `priority` - Priority. Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`, `61440`.
* `stp_port` - Port configuration. The structure of `stp_port` block is documented below.
* `fswid` - Instance id.
* `vlan_range` - Vlan-range.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `stp_port` block supports:

* `priority` - Port priority. Valid values: `0`, `16`, `32`, `64`, `80`, `96`, `112`, `128`, `144`, `160`, `176`, `192`, `208`, `224`, `240`.
* `cost` - Port cost, 0 means to select a value for the path cost based on linkspeed.
* `name` - Port name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

SwitchStp Instance can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchstp_instance.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchstp_instance.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
