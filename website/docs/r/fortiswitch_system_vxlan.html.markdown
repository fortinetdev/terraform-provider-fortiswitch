---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_vxlan"
description: |-
  Configure VXLAN devices.
---

# fortiswitch_system_vxlan
Configure VXLAN devices. Applies to FortiSwitch Version `>= 7.2.0`.

## Argument Reference

The following arguments are supported:

* `remote_ip` - IPv4 address of the VXLAN interface. The structure of `remote_ip` block is documented below.
* `name` - VXLAN interface name.
* `multicast_ttl` - VXLAN multicast TTL (1-255, default = 0).
* `tagged_vlans` - Tagged VLANs.
* `vni` - VXLAN network ID.
* `ip_version` - IP version to use for the VXLAN interface. Valid values: `ipv4-unicast`, `ipv4-multicast`.
* `vlanid` - VXLAN VNI mapped VLAN ID.
* `interface` - Interface binding for tunnel initiator (source IP) / termination.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `remote_ip` block supports:

* `ip` - IPv4 address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Vxlan can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_vxlan.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_vxlan.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
