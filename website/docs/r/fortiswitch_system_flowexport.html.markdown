---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_flowexport"
description: |-
  System Flow Export settings.
---

# fortiswitch_system_flowexport
System Flow Export settings.

## Example Usage

```hcl
resource "fortiswitch_system_flowexport" "name" {
        timeout_udp = "604800"
}
```

## Argument Reference

The following arguments are supported:

* `timeout_tcp` - Flow Session TCP Timeout (60-604800, default=3600 seconds).
* `timeout_udp` - Flow Session UDP Timeout (60-604800, default=300 seconds).
* `collectors` - Collectors. The structure of `collectors` block is documented below.
* `level` - Export Level (vlan|ip|port|protocol|mac). Valid values: `mac`, `ip`, `proto`, `port`, `vlan`.
* `timeout_tcp_rst` - Flow Session TCP Reset Timeout (60-604800, default=120 seconds).
* `aggregates` - Aggregates. The structure of `aggregates` block is documented below.
* `format` - Export Format (netflow1|netflow5|netflow9|ipfix). Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
* `timeout_tcp_fin` - Flow Session TCP Fin Timeout (60-604800, default=300 seconds).
* `template_export_period` - Template export period in minutes (1-60).
* `filter` - Filter (BPF).
* `timeout_icmp` - Flow Session ICMP Timeout (60-604800, default=300 seconds).
* `max_export_pkt_size` - Max Export Packet Size (512-9216, default=512 bytes).
* `timeout_general` - Flow Session General Timeout (60-604800, default=3600 seconds).
* `identity` - Set identity of switch (0x00000000-0xFFFFFFFF default=0x00000000).
* `timeout_max` - Flow Session MAX Timeout (60-604800, default=604800 seconds).
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `collectors` block supports:

* `ip` - IP address.
* `name` - Collector name.
* `transport` - Export transport (udp|tcp|sctp). Valid values: `udp`, `tcp`, `sctp`.
* `port` - Port number (0-65535, default=0 (netflow:2055, ipfix:4739).

The `aggregates` block supports:

* `ip` - Aggregate's IP and Mask.
* `id` - Aggregate id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FlowExport can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_flowexport.labelname SystemFlowExport

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_flowexport.labelname SystemFlowExport
$ unset "FORTISWITCH_IMPORT_TABLE"
```
