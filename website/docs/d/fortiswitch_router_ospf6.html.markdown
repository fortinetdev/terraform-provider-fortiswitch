---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_ospf6"
description: |-
  Get information on fortiswitch router ospf6.
---

# Data Source: fortiswitch_router_ospf6
Use this data source to get information on fortiswitch router ospf6

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `router_id` - A.B.C.D, in IPv4 address format.
* `spf_timers` - SPF calculation frequency.
* `area` - OSPF6 area configuration. The structure of `area` block is documented below.
* `log_neighbor_changes` - Enable logging of OSPF neighbor's changes.
* `interface` - OSPF6 interface configuration. The structure of `interface` block is documented below.

The `redistribute` block contains:

* `status` - status
* `metric_type` - metric type
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `name` - Redistribute name.

The `area` block contains:

* `stub_type` - Stub summary setting.
* `filter_list` - OSPF area filter-list configuration. The structure of `filter_list` block is documented below.
* `range` - OSPF6 area range configuration. The structure of `range` block is documented below.
* `type` - Area type setting.
* `id` - Area entry ip address.

The `filter_list` block contains:

* `direction` - Direction.
* `list` - Access-list or prefix-list name.
* `id` - Filter list entry ID.

The `range` block contains:

* `advertise` - Enable/disable advertise status.
* `id` - Range entry id.
* `prefix6` - <prefix6>   IPv6 prefix

The `interface` block contains:

* `status` - Enable/disable OSPF6 routing on this interface.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `name` - Interface name.
* `bfd` - Enable/Disable Bidirectional Forwarding Detection (BFD).
* `area_id` - A.B.C.D, in IPv4 address format.
* `transmit_delay` - Link state transmit delay.
* `retransmit_interval` - Time between retransmitting lost link state advertisements.
* `cost` - The cost of the interface.
* `passive` - Enable/disable passive interface.
* `priority` - Router priority.

