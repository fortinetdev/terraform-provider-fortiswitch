---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_rip"
description: |-
  Get information on fortiswitch router rip.
---

# Data Source: fortiswitch_router_rip
Use this data source to get information on fortiswitch router rip

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `distance` - Set admin distance based on route source ip. The structure of `distance` block is documented below.
* `default_metric` - Default metric of redistribute routes (Except connected).
* `recv_buffer_size` - receiving buffer size
* `timeout_timer` - Routing information timeout timer.
* `name` - Vrf name.
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `offset_list` - Offset list to modify RIP metric. The structure of `offset_list` block is documented below.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `neighbor` - Specify a neighbor router. Required only for non-multicast networks. The structure of `neighbor` block is documented below.
* `version` - RIP version
* `garbage_timer` - Garbage collection timer.
* `vrf` - Enable RIP on VRF. The structure of `vrf` block is documented below.
* `default_information_originate` - Generate a default route.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `update_timer` - Routing table update timer.
* `interface` - RIP interface configuration The structure of `interface` block is documented below.
* `distribute_list` - Filter networks in routing updates. The structure of `distribute_list` block is documented below.
* `network` - Enable RIP routing on an IP network. The structure of `network` block is documented below.

The `distance` block contains:

* `access_list` - Access list for route destination.
* `distance` - Distance.
* `prefix` - IP source prefix.
* `id` - Distance id.

The `offset_list` block contains:

* `status` - Status.
* `direction` - Offset list direction.
* `interface` - Interface to match.
* `offset` - Metric value.
* `access_list` - Access list name.
* `id` - Offset-list id.

The `redistribute` block contains:

* `status` - status
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `flags` - flags
* `name` - Redistribute name.

The `neighbor` block contains:

* `ip` - IP address.
* `id` - Neighbor entry id.

The `vrf` block contains:

* `distance` - Set admin distance based on route source ip. The structure of `distance` block is documented below.
* `default_metric` - Default metric of redistribute routes (Except connected).
* `recv_buffer_size` - receiving buffer size
* `timeout_timer` - Routing information timeout timer.
* `name` - Vrf name.
* `offset_list` - Offset list to modify RIP metric. The structure of `offset_list` block is documented below.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `neighbor` - Specify a neighbor router. Required only for non-multicast networks. The structure of `neighbor` block is documented below.
* `version` - RIP version
* `garbage_timer` - Garbage collection timer.
* `default_information_originate` - Generate a default route.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `update_timer` - Routing table update timer.
* `interface` - RIP interface configuration The structure of `interface` block is documented below.
* `distribute_list` - Filter networks in routing updates. The structure of `distribute_list` block is documented below.
* `network` - Enable RIP routing on an IP network. The structure of `network` block is documented below.

The `distance` block contains:

* `access_list` - Access list for route destination.
* `distance` - Distance.
* `prefix` - IP source prefix.
* `id` - Distance id.

The `offset_list` block contains:

* `status` - Status.
* `direction` - Offset list direction.
* `interface` - Interface to match.
* `offset` - Metric value.
* `access_list` - Access list name.
* `id` - Offset-list id.

The `redistribute` block contains:

* `status` - status
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `flags` - flags
* `name` - Redistribute name.

The `neighbor` block contains:

* `ip` - IP address.
* `id` - Neighbor entry id.

The `passive_interface` block contains:

* `name` - Passive interface name.

The `interface` block contains:

* `split_horizon_status` - Split horizon status.
* `auth_mode` - Authentication mode.
* `name` - interface name
* `send_version2_broadcast` - broadcast version 1 compatible packets
* `send_version` - Send version.
* `auth_keychain` - Authentication keychain name.
* `split_horizon` - Split horizon method.
* `flags` - flags
* `auth_string` - Authentication string/password.
* `receive_version` - Receive version.

The `distribute_list` block contains:

* `status` - Status.
* `listname` - Distribute access/prefix list name.
* `direction` - Distribute list direction.
* `interface` - Distribute list interface name.
* `id` - Distribute-list id.

The `network` block contains:

* `prefix` - Network prefix.
* `id` - Network entry id.

The `passive_interface` block contains:

* `name` - Passive interface name.

The `interface` block contains:

* `split_horizon_status` - Split horizon status.
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `auth_mode` - Authentication mode.
* `name` - interface name
* `send_version2_broadcast` - broadcast version 1 compatible packets
* `send_version` - Send version.
* `auth_keychain` - Authentication keychain name.
* `split_horizon` - Split horizon method.
* `flags` - flags
* `auth_string` - Authentication string/password.
* `receive_version` - Receive version.

The `distribute_list` block contains:

* `status` - Status.
* `listname` - Distribute access/prefix list name.
* `direction` - Distribute list direction.
* `interface` - Distribute list interface name.
* `id` - Distribute-list id.

The `network` block contains:

* `prefix` - Network prefix.
* `id` - Network entry id.

