---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_mirror"
description: |-
  Packet mirror.
---

# fortiswitch_switch_mirror
Packet mirror.

## Example Usage

```hcl
resource "fortiswitch_switch_mirror" "name" {
        dst = "port3"
        mode = "SPAN"
        name = "test"
        src_egress {
            name = "port1"
        }
        src_ingress {
            name = "port2"
        }
        status = "active"
        switching_packet = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Status. Valid values: `active`, `inactive`.
* `encap_vlan` - Control the tagged/untagged status of ERSPAN encapsulation headers. Valid values: `tagged`, `untagged`.
* `encap_ipv4_tos` - TOS, or DSCP and ECN, values in the ERSPAN IP header.
* `erspan_collector_ip` - ERSPAN collector IP address.
* `switching_packet` - Enable/disable switching functionality when mirroring. Valid values: `enable`, `disable`.
* `encap_vlan_priority` - Priority code point value in the ERSPAN or RSPAN VLAN header.
* `dst` - Destination interface.
* `encap_gre_protocol` - Protocol value in the ERSPAN GRE header.
* `src_egress` - Source egress interfaces. The structure of `src_egress` block is documented below.
* `encap_ipv4_ttl` - IPv4 time-to-live value in the ERSPAN IP header.
* `encap_vlan_id` - VLAN ID in the ERSPAN or RSPAN VLAN header.
* `encap_vlan_tpid` - TPID in the ERSPAN or RSPAN VLAN header.
* `mode` - Mirroring mode.
* `encap_mac_dst` - Nexthop/Gateway MAC address on the path to the ERSPAN collector IP.
* `src_ingress` - Source ingress interfaces. The structure of `src_ingress` block is documented below.
* `rspan_ip` - RSPAN destination IP address.
* `encap_mac_src` - Source MAC address in the ERSPAN ethernet header.
* `strip_mirrored_traffic_tags` - Enable/disable stripping of VLAN tags from mirrored traffic. Valid values: `enable`, `disable`.
* `encap_ipv4_src` - IPv4 source address in the ERSPAN IP header.
* `encap_vlan_cfi` - CFI or DEI bit in the ERSPAN or RSPAN VLAN header.
* `name` - Mirror session name.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `src_egress` block supports:

* `name` - Interface name.

The `src_ingress` block supports:

* `name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Switch Mirror can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_mirror.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_mirror.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
