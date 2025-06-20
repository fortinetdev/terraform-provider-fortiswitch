---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_multicast"
description: |-
  Get information on fortiswitch router multicast.
---

# Data Source: fortiswitch_router_multicast
Use this data source to get information on fortiswitch router multicast

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `comments` - Description/comments.
* `interface` - Pim interfaces. The structure of `interface` block is documented below.
* `multicast_routing` - Enable multicast routing.

The `interface` block contains:

* `hello_interval` - Interval between sending PIM hello messages(sec).
* `name` - Interface name.
* `pim_mode` - PIM operation mode.
* `dr_priority` - DR election priority.
* `multicast_flow` - Config acceptable source for multicast group.
* `igmp` - Igmp configuration options. The structure of `igmp` block is documented below.

The `igmp` block contains:

* `query_interval` - Interval between queries to IGMP hosts(sec).
* `query_max_response_time` - Maximum time to wait for a IGMP query response(sec).

