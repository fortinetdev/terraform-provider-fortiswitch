---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_trunk"
description: |-
  Link-aggregation.
---

# fortiswitch_switch_trunk
Link-aggregation.

## Argument Reference

The following arguments are supported:

* `hb_src_udp_port` - Source UDP port of heartbeat packet.
* `hb_verify` - Enable/disable heartbeat packet strict validation. Valid values: `enable`, `disable`.
* `hb_src_ip` - Source IP address of heartbeat packet.
* `min_bundle` - Minimum size of bundle.
* `hb_dst_udp_port` - Destination UDP port of heartbeat packet.
* `member_withdrawal_behavior` - Port behaviors after it withdraws because of loss of control packets. Valid values: `forward`, `block`.
* `fallback_port` - LACP fallback port.
* `mclag_mac_address` - MCLAG MAC address.
* `isl_fortilink` - ISL fortiLink trunk.
* `max_miss_heartbeats` - Maximum tolerant missed heartbeats.
* `restricted` - Restricted ISL ICL trunk.
* `aggregator_mode` - LACP Member Select Mode. Valid values: `bandwidth`, `count`.
* `port_selection_criteria` - Algorithm for aggregate port selection. Valid values: `src-mac`, `dst-mac`, `src-dst-mac`, `src-ip`, `dst-ip`, `src-dst-ip`.
* `lacp_speed` - LACP speed. Valid values: `slow`, `fast`.
* `mclag` - Multi Chassis LAG. Valid values: `enable`, `disable`.
* `trunk_id` - Internal id.
* `description` - Description.
* `static_isl_auto_vlan` - User ISL auto VLAN. Valid values: `enable`, `disable`.
* `bundle` - Enable bundle. Valid values: `enable`, `disable`.
* `max_bundle` - Maximum size of bundle.
* `mclag_icl` - MCLAG inter-chassis-link. Valid values: `enable`, `disable`.
* `hb_in_vlan` - Receive VLAN ID in heartbeat packet.
* `members` - Aggregated interfaces. The structure of `members` block is documented below.
* `hb_dst_ip` - Destination IP address of heartbeat packet.
* `hb_out_vlan` - Transmit VLAN ID in heartbeat packet.
* `fortilink` - FortiLink trunk.
* `name` - Trunk name.
* `auto_isl` - Trunk with auto-isl.
* `port_extension` - Port extension enable. Valid values: `enable`, `disable`.
* `mode` - Link Aggreation mode. Valid values: `static`, `lacp-passive`, `lacp-active`, `fortinet-trunk`.
* `static_isl` - Static ISL. Valid values: `enable`, `disable`.
* `port_extension_trigger` - Number of failed port to trigger the whole trunk down.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `members` block supports:

* `member_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Switch Trunk can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_trunk.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_trunk.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
