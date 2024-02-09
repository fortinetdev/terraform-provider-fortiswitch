---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_ospf"
description: |-
  OSPF configuration.
---

# fortiswitch_router_ospf
OSPF configuration.

~> The provider supports the definition of Interface in Router Ospf `fortiswitch_router_ospf`, and also allows the definition of separate Interface resources `fortiswitch_routerospf_interface`, but do not use a `fortiswitch_router_ospf` with in-line Interface in conjunction with any `fortiswitch_routerospf_interface` resources, otherwise conflicts and overwrite will occur.

~> The provider supports the definition of Network in Router Ospf `fortiswitch_router_ospf`, and also allows the definition of separate Network resources `fortiswitch_routerospf_network`, but do not use a `fortiswitch_router_ospf` with in-line Network in conjunction with any `fortiswitch_routerospf_network` resources, otherwise conflicts and overwrite will occur.

~> The provider supports the definition of Area in Router Ospf `fortiswitch_router_ospf`, and also allows the definition of separate Area resources `fortiswitch_routerospf_area`, but do not use a `fortiswitch_router_ospf` with in-line Area in conjunction with any `fortiswitch_routerospf_area` resources, otherwise conflicts and overwrite will occur.



## Example Usage

```hcl
resource "fortiswitch_router_ospf" "name" {
        abr_type = "cisco"
        database_overflow = "enable"
        database_overflow_max_external_lsa = "34"
        database_overflow_time_to_recover = "35"
        default_information_metric = "36"
        default_information_metric_type = "1"
        default_information_originate = "enable"
        distance = "39"
        distance_external = "40"
        distance_inter_area = "41"
        distance_intra_area = "42"
}
```

## Argument Reference

The following arguments are supported:

* `summary_address` - Aggregate address for redistributed routes. The structure of `summary_address` block is documented below.
* `area` - OSPF area configuration. The structure of `area` block is documented below.
* `distance_intra_area` - Administrative intra-area route distance.
* `distance_external` - Administrative external route distance.
* `network` - Enable OSPF on an IP network. The structure of `network` block is documented below.
* `router_id` - Router ID.
* `default_information_metric_type` - Default information metric type. Valid values: `1`, `2`.
* `database_overflow_time_to_recover` - Database overflow time to recover (sec).
* `log_neighbour_changes` - Enable logging of OSPF neighbour's changes Valid values: `enable`, `disable`.
* `rfc1583_compatible` - Enable/disable RFC1583 compatibility. Valid values: `enable`, `disable`.
* `default_information_metric` - Default information metric.
* `default_information_originate` - Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `distribute_list` - Redistribute routes filter. The structure of `distribute_list` block is documented below.
* `distance` - Administrative distance.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `database_overflow` - Enable/disable database overflow. Valid values: `enable`, `disable`.
* `database_overflow_max_external_lsa` - Database overflow maximum External LSAs.
* `name` - Vrf name.
* `spf_timers` - SPF calculation frequency.
* `distance_inter_area` - Administrative inter-area route distance.
* `abr_type` - Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.
* `interface` - OSPF interface configuration. The structure of `interface` block is documented below.
* `vrf` - Enable OSPF on VRF. The structure of `vrf` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `summary_address` block supports:

* `prefix` - Prefix.
* `tag` - Tag value.
* `id` - Summary address entry ID.

The `area` block supports:

* `nssa_translator_role` - NSSA translator role type. Valid values: `candidate`, `never`, `always`.
* `virtual_link` - OSPF virtual link configuration. The structure of `virtual_link` block is documented below.
* `shortcut` - Enable/disable shortcut option. Valid values: `disable`, `enable`, `default`.
* `stub_type` - Stub summary setting. Valid values: `no-summary`, `summary`.
* `range` - OSPF area range configuration. The structure of `range` block is documented below.
* `filter_list` - OSPF area filter-list configuration. The structure of `filter_list` block is documented below.
* `default_cost` - Summary default cost of stub or NSSA area.
* `type` - Area type setting. Valid values: `regular`, `nssa`, `stub`.
* `id` - Area entry IP address.

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

The `network` block supports:

* `prefix` - Prefix.
* `id` - Network entry ID.
* `area` - Attach the network to area.

The `passive_interface` block supports:

* `name` - Passive interface name.

The `distribute_list` block supports:

* `access_list` - Access list name.
* `protocol` - Protocol type. Valid values: `connected`, `static`, `rip`, `bgp`, `isis`.
* `id` - Distribute list entry ID.

The `redistribute` block supports:

* `status` - status Valid values: `enable`, `disable`.
* `name` - Redistribute name.
* `metric_type` - Metric type. Valid values: `1`, `2`.
* `tag` - Tag value.
* `routemap` - Route map name.
* `metric` - Redistribute metric setting.

The `interface` block supports:

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

The `md5_keys` block supports:

* `id` - key-id (1-255).
* `key` - md5-key.

The `vrf` block supports:

* `summary_address` - Aggregate address for redistributed routes. The structure of `summary_address` block is documented below.
* `area` - OSPF area configuration. The structure of `area` block is documented below.
* `distance_intra_area` - Administrative intra-area route distance.
* `distance_external` - Administrative external route distance.
* `network` - Enable OSPF on an IP network. The structure of `network` block is documented below.
* `router_id` - Router ID.
* `default_information_metric_type` - Default information metric type. Valid values: `1`, `2`.
* `database_overflow_time_to_recover` - Database overflow time to recover (sec).
* `log_neighbour_changes` - Enable logging of OSPF neighbour's changes Valid values: `enable`, `disable`.
* `rfc1583_compatible` - Enable/disable RFC1583 compatibility. Valid values: `enable`, `disable`.
* `default_information_metric` - Default information metric.
* `default_information_originate` - Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `distribute_list` - Redistribute routes filter. The structure of `distribute_list` block is documented below.
* `distance` - Administrative distance.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `database_overflow` - Enable/disable database overflow. Valid values: `enable`, `disable`.
* `database_overflow_max_external_lsa` - Database overflow maximum External LSAs.
* `name` - Vrf name.
* `spf_timers` - SPF calculation frequency.
* `distance_inter_area` - Administrative inter-area route distance.
* `abr_type` - Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.
* `interface` - OSPF interface configuration. The structure of `interface` block is documented below.

The `summary_address` block supports:

* `prefix` - Prefix.
* `tag` - Tag value.
* `id` - Summary address entry ID.

The `area` block supports:

* `nssa_translator_role` - NSSA translator role type. Valid values: `candidate`, `never`, `always`.
* `virtual_link` - OSPF virtual link configuration. The structure of `virtual_link` block is documented below.
* `shortcut` - Enable/disable shortcut option. Valid values: `disable`, `enable`, `default`.
* `stub_type` - Stub summary setting. Valid values: `no-summary`, `summary`.
* `range` - OSPF area range configuration. The structure of `range` block is documented below.
* `filter_list` - OSPF area filter-list configuration. The structure of `filter_list` block is documented below.
* `default_cost` - Summary default cost of stub or NSSA area.
* `type` - Area type setting. Valid values: `regular`, `nssa`, `stub`.
* `id` - Area entry IP address.

The `virtual_link` block supports:

* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `name` - Virtual link entry name.
* `transmit_delay` - Link state transmit delay.
* `retransmit_interval` - Time between retransmitting lost link state advertisements.
* `authentication` - Authentication type. Valid values: `none`, `text`.
* `peer` - Peer IP.
* `authentication_key` - Authentication key.

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

The `network` block supports:

* `prefix` - Prefix.
* `id` - Network entry ID.
* `area` - Attach the network to area.

The `passive_interface` block supports:

* `name` - Passive interface name.

The `distribute_list` block supports:

* `access_list` - Access list name.
* `protocol` - Protocol type. Valid values: `connected`, `static`, `rip`, `bgp`, `isis`.
* `id` - Distribute list entry ID.

The `redistribute` block supports:

* `status` - status Valid values: `enable`, `disable`.
* `name` - Redistribute name.
* `metric_type` - Metric type. Valid values: `1`, `2`.
* `tag` - Tag value.
* `routemap` - Route map name.
* `metric` - Redistribute metric setting.

The `interface` block supports:

* `priority` - Router priority.
* `authentication_key` - Authentication key.
* `name` - Interface entry name.
* `dead_interval` - Dead interval. For fast-hello assign value 1.
* `hello_multiplier` - Number of hello packets within dead interval.Valid only for fast-hello.
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

The `md5_keys` block supports:

* `id` - key-id (1-255).
* `key` - md5-key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Ospf can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_ospf.labelname RouterOspf

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_ospf.labelname RouterOspf
$ unset "FORTISWITCH_IMPORT_TABLE"
```
