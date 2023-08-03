---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_ripng"
description: |-
  Get information on fortiswitch router ripng.
---

# Data Source: fortiswitch_router_ripng
Use this data source to get information on fortiswitch router ripng

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `default_metric` - Default metric of redistribute routes (Except connected).
* `timeout_timer` - Routing information timeout timer.
* `aggregate_address` - Set aggregate RIPng route announcement. The structure of `aggregate_address` block is documented below.
* `offset_list` - Offset list to modify RIPng metric. The structure of `offset_list` block is documented below.
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `garbage_timer` - Garbage collection timer.
* `default_information_originate` - Generate a default route.
* `interface` - RIPng interface configuration. The structure of `interface` block is documented below.
* `update_timer` - Routing table update timer.
* `distribute_list` - Filter networks in routing updates. The structure of `distribute_list` block is documented below.

The `aggregate_address` block contains:

* `id` - Aggregate-address entry id.
* `prefix6` - Aggregate-address prefix.

The `offset_list` block contains:

* `status` - Status.
* `direction` - Offset list direction.
* `access_list6` - Ipv6 access list name.
* `offset` - Metric offset.
* `interface` - Interface name.
* `id` - Offset-list id.

The `redistribute` block contains:

* `status` - status
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `flags` - Flags
* `name` - Redistribute name.

The `interface` block contains:

* `passive` - Suppress routing updates on an interface.
* `split_horizon_status` - Enable/disable split horizon.
* `split_horizon` - Split horizon type.
* `flags` - Flags.
* `name` - Interface name.

The `distribute_list` block contains:

* `status` - Status.
* `listname` - Distribute access/prefix list name.
* `direction` - Distribute list direction.
* `interface` - Distribute list interface name.
* `id` - Distribute-list id.

