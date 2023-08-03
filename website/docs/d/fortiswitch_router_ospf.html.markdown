---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_ospf"
description: |-
  Get information on fortiswitch router ospf.
---

# Data Source: fortiswitch_router_ospf
Use this data source to get information on fortiswitch router ospf

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `summary_address` - Aggregate address for redistributed routes. The structure of `summary_address` block is documented below.
* `area` - OSPF area configuration. The structure of `area` block is documented below.
* `distance_intra_area` - Administrative intra-area route distance.
* `distance_external` - Administrative external route distance.
* `network` - Enable OSPF on an IP network. The structure of `network` block is documented below.
* `router_id` - Router ID.
* `default_information_metric_type` - Default information metric type.
* `database_overflow_time_to_recover` - Database overflow time to recover (sec).
* `log_neighbour_changes` - Enable logging of OSPF neighbour's changes
* `rfc1583_compatible` - Enable/disable RFC1583 compatibility.
* `default_information_metric` - Default information metric.
* `default_information_originate` - Enable/disable generation of default route.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `distribute_list` - Redistribute routes filter. The structure of `distribute_list` block is documented below.
* `distance` - Administrative distance.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `database_overflow` - Enable/disable database overflow.
* `database_overflow_max_external_lsa` - Database overflow maximum External LSAs.
* `name` - Vrf name.
* `spf_timers` - SPF calculation frequency.
* `distance_inter_area` - Administrative inter-area route distance.
* `abr_type` - Area border router type.
* `interface` - OSPF interface configuration. The structure of `interface` block is documented below.
* `vrf` - Enable OSPF on VRF. The structure of `vrf` block is documented below.

The `summary_address` block contains:

* `prefix` - Prefix.
* `tag` - Tag value.
* `id` - Summary address entry ID.

The `area` block contains:

* `nssa_translator_role` - NSSA translator role type.
* `virtual_link` - OSPF virtual link configuration. The structure of `virtual_link` block is documented below.
* `shortcut` - Enable/disable shortcut option.
* `stub_type` - Stub summary setting.
* `range` - OSPF area range configuration. The structure of `range` block is documented below.
* `filter_list` - OSPF area filter-list configuration. The structure of `filter_list` block is documented below.
* `default_cost` - Summary default cost of stub or NSSA area.
* `type` - Area type setting.
* `id` - Area entry IP address.

The `virtual_link` block contains:

* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `name` - Virtual link entry name.
* `transmit_delay` - Link state transmit delay.
* `retransmit_interval` - Time between retransmitting lost link state advertisements.
* `authentication` - Authentication type.
* `peer` - Peer IP.
* `md5_keys` - OSPF md5 key configuration. Applicable only when authentication field is set to md5. The structure of `md5_keys` block is documented below.
* `authentication_key` - Authentication key.

The `md5_keys` block contains:

* `id` - key-id (1-255).
* `key` - md5-key.

The `range` block contains:

* `substitute` - Substitute prefix.
* `prefix` - Prefix.
* `substitute_status` - Enable/disable substitute status.
* `id` - Range entry ID.
* `advertise` - Enable/disable advertise status.

The `filter_list` block contains:

* `direction` - Direction.
* `list` - Access-list or prefix-list name.
* `id` - Filter list entry ID.

The `network` block contains:

* `prefix` - Prefix.
* `id` - Network entry ID.
* `area` - Attach the network to area.

The `passive_interface` block contains:

* `name` - Passive interface name.

The `distribute_list` block contains:

* `access_list` - Access list name.
* `protocol` - Protocol type.
* `id` - Distribute list entry ID.

The `redistribute` block contains:

* `status` - status
* `name` - Redistribute name.
* `metric_type` - Metric type.
* `tag` - Tag value.
* `routemap` - Route map name.
* `metric` - Redistribute metric setting.

The `interface` block contains:

* `priority` - Router priority.
* `authentication_key` - Authentication key.
* `name` - Interface entry name.
* `dead_interval` - Dead interval. For fast-hello assign value 1.
* `hello_multiplier` - Number of hello packets within dead interval.Valid only for fast-hello.
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `transmit_delay` - Link state transmit delay.
* `ucast_ttl` - Unicast TTL.
* `mtu` - Interface MTU.
* `retransmit_interval` - Time between retransmitting lost link state advertisements.
* `authentication` - Authentication type.
* `cost` - Cost of the interface.
* `hello_interval` - Hello interval.
* `ttl` - TTL.
* `md5_keys` - OSPF md5 key configuration. Applicable only when authentication field is set to md5. The structure of `md5_keys` block is documented below.
* `mtu_ignore` - Disable MTU mismatch detection on this interface.

The `md5_keys` block contains:

* `id` - key-id (1-255).
* `key` - md5-key.

The `vrf` block contains:

* `summary_address` - Aggregate address for redistributed routes. The structure of `summary_address` block is documented below.
* `area` - OSPF area configuration. The structure of `area` block is documented below.
* `distance_intra_area` - Administrative intra-area route distance.
* `distance_external` - Administrative external route distance.
* `network` - Enable OSPF on an IP network. The structure of `network` block is documented below.
* `router_id` - Router ID.
* `default_information_metric_type` - Default information metric type.
* `database_overflow_time_to_recover` - Database overflow time to recover (sec).
* `log_neighbour_changes` - Enable logging of OSPF neighbour's changes
* `rfc1583_compatible` - Enable/disable RFC1583 compatibility.
* `default_information_metric` - Default information metric.
* `default_information_originate` - Enable/disable generation of default route.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `distribute_list` - Redistribute routes filter. The structure of `distribute_list` block is documented below.
* `distance` - Administrative distance.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `database_overflow` - Enable/disable database overflow.
* `database_overflow_max_external_lsa` - Database overflow maximum External LSAs.
* `name` - Vrf name.
* `spf_timers` - SPF calculation frequency.
* `distance_inter_area` - Administrative inter-area route distance.
* `abr_type` - Area border router type.
* `interface` - OSPF interface configuration. The structure of `interface` block is documented below.

The `summary_address` block contains:

* `prefix` - Prefix.
* `tag` - Tag value.
* `id` - Summary address entry ID.

The `area` block contains:

* `nssa_translator_role` - NSSA translator role type.
* `virtual_link` - OSPF virtual link configuration. The structure of `virtual_link` block is documented below.
* `shortcut` - Enable/disable shortcut option.
* `stub_type` - Stub summary setting.
* `range` - OSPF area range configuration. The structure of `range` block is documented below.
* `filter_list` - OSPF area filter-list configuration. The structure of `filter_list` block is documented below.
* `default_cost` - Summary default cost of stub or NSSA area.
* `type` - Area type setting.
* `id` - Area entry IP address.

The `virtual_link` block contains:

* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `name` - Virtual link entry name.
* `transmit_delay` - Link state transmit delay.
* `retransmit_interval` - Time between retransmitting lost link state advertisements.
* `authentication` - Authentication type.
* `peer` - Peer IP.
* `authentication_key` - Authentication key.

The `range` block contains:

* `substitute` - Substitute prefix.
* `prefix` - Prefix.
* `substitute_status` - Enable/disable substitute status.
* `id` - Range entry ID.
* `advertise` - Enable/disable advertise status.

The `filter_list` block contains:

* `direction` - Direction.
* `list` - Access-list or prefix-list name.
* `id` - Filter list entry ID.

The `network` block contains:

* `prefix` - Prefix.
* `id` - Network entry ID.
* `area` - Attach the network to area.

The `passive_interface` block contains:

* `name` - Passive interface name.

The `distribute_list` block contains:

* `access_list` - Access list name.
* `protocol` - Protocol type.
* `id` - Distribute list entry ID.

The `redistribute` block contains:

* `status` - status
* `name` - Redistribute name.
* `metric_type` - Metric type.
* `tag` - Tag value.
* `routemap` - Route map name.
* `metric` - Redistribute metric setting.

The `interface` block contains:

* `priority` - Router priority.
* `authentication_key` - Authentication key.
* `name` - Interface entry name.
* `dead_interval` - Dead interval. For fast-hello assign value 1.
* `hello_multiplier` - Number of hello packets within dead interval.Valid only for fast-hello.
* `transmit_delay` - Link state transmit delay.
* `ucast_ttl` - Unicast TTL.
* `mtu` - Interface MTU.
* `retransmit_interval` - Time between retransmitting lost link state advertisements.
* `authentication` - Authentication type.
* `cost` - Cost of the interface.
* `hello_interval` - Hello interval.
* `ttl` - TTL.
* `md5_keys` - OSPF md5 key configuration. Applicable only when authentication field is set to md5. The structure of `md5_keys` block is documented below.
* `mtu_ignore` - Disable MTU mismatch detection on this interface.

The `md5_keys` block contains:

* `id` - key-id (1-255).
* `key` - md5-key.

