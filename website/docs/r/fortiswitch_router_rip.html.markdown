---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_rip"
description: |-
  RIP configuration.
---

# fortiswitch_router_rip
RIP configuration.

## Example Usage

```hcl
resource "fortiswitch_router_rip" "name" {
        recv_buffer_size = 6553660
}
```

## Argument Reference

The following arguments are supported:

* `distance` - Set admin distance based on route source ip. The structure of `distance` block is documented below.
* `default_metric` - Default metric of redistribute routes (Except connected).
* `recv_buffer_size` - receiving buffer size
* `timeout_timer` - Routing information timeout timer.
* `name` - Vrf name.
* `bfd` - Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
* `offset_list` - Offset list to modify RIP metric. The structure of `offset_list` block is documented below.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `neighbor` - Specify a neighbor router. Required only for non-multicast networks. The structure of `neighbor` block is documented below.
* `version` - RIP version Valid values: `1`, `2`.
* `garbage_timer` - Garbage collection timer.
* `vrf` - Enable RIP on VRF. The structure of `vrf` block is documented below.
* `default_information_originate` - Generate a default route. Valid values: `enable`, `disable`.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `update_timer` - Routing table update timer.
* `interface` - RIP interface configuration The structure of `interface` block is documented below.
* `distribute_list` - Filter networks in routing updates. The structure of `distribute_list` block is documented below.
* `network` - Enable RIP routing on an IP network. The structure of `network` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `distance` block supports:

* `access_list` - Access list for route destination.
* `distance` - Distance.
* `prefix` - IP source prefix.
* `id` - Distance id.

The `offset_list` block supports:

* `status` - Status. Valid values: `enable`, `disable`.
* `direction` - Offset list direction. Valid values: `in`, `out`.
* `interface` - Interface to match.
* `offset` - Metric value.
* `access_list` - Access list name.
* `id` - Offset-list id.

The `redistribute` block supports:

* `status` - status Valid values: `enable`, `disable`.
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `flags` - flags
* `name` - Redistribute name.

The `neighbor` block supports:

* `ip` - IP address.
* `id` - Neighbor entry id.

The `vrf` block supports:

* `distance` - Set admin distance based on route source ip. The structure of `distance` block is documented below.
* `default_metric` - Default metric of redistribute routes (Except connected).
* `recv_buffer_size` - receiving buffer size
* `timeout_timer` - Routing information timeout timer.
* `name` - Vrf name.
* `offset_list` - Offset list to modify RIP metric. The structure of `offset_list` block is documented below.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `neighbor` - Specify a neighbor router. Required only for non-multicast networks. The structure of `neighbor` block is documented below.
* `version` - RIP version Valid values: `1`, `2`.
* `garbage_timer` - Garbage collection timer.
* `default_information_originate` - Generate a default route. Valid values: `enable`, `disable`.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `update_timer` - Routing table update timer.
* `interface` - RIP interface configuration The structure of `interface` block is documented below.
* `distribute_list` - Filter networks in routing updates. The structure of `distribute_list` block is documented below.
* `network` - Enable RIP routing on an IP network. The structure of `network` block is documented below.

The `distance` block supports:

* `access_list` - Access list for route destination.
* `distance` - Distance.
* `prefix` - IP source prefix.
* `id` - Distance id.

The `offset_list` block supports:

* `status` - Status. Valid values: `enable`, `disable`.
* `direction` - Offset list direction. Valid values: `in`, `out`.
* `interface` - Interface to match.
* `offset` - Metric value.
* `access_list` - Access list name.
* `id` - Offset-list id.

The `redistribute` block supports:

* `status` - status Valid values: `enable`, `disable`.
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `flags` - flags
* `name` - Redistribute name.

The `neighbor` block supports:

* `ip` - IP address.
* `id` - Neighbor entry id.

The `passive_interface` block supports:

* `name` - Passive interface name.

The `interface` block supports:

* `split_horizon_status` - Split horizon status. Valid values: `enable`, `disable`.
* `auth_mode` - Authentication mode. Valid values: `none`, `text`, `md5`.
* `name` - interface name
* `send_version2_broadcast` - broadcast version 1 compatible packets Valid values: `disable`, `enable`.
* `send_version` - Send version. Valid values: `global`, `1`, `2`, `both`.
* `auth_keychain` - Authentication keychain name.
* `split_horizon` - Split horizon method. Valid values: `poisoned`, `regular`.
* `flags` - flags
* `auth_string` - Authentication string/password.
* `receive_version` - Receive version. Valid values: `global`, `1`, `2`, `both`.

The `distribute_list` block supports:

* `status` - Status. Valid values: `enable`, `disable`.
* `listname` - Distribute access/prefix list name.
* `direction` - Distribute list direction. Valid values: `in`, `out`.
* `interface` - Distribute list interface name.
* `id` - Distribute-list id.

The `network` block supports:

* `prefix` - Network prefix.
* `id` - Network entry id.

The `passive_interface` block supports:

* `name` - Passive interface name.

The `interface` block supports:

* `split_horizon_status` - Split horizon status. Valid values: `enable`, `disable`.
* `bfd` - Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
* `auth_mode` - Authentication mode. Valid values: `none`, `text`, `md5`.
* `name` - interface name
* `send_version2_broadcast` - broadcast version 1 compatible packets Valid values: `disable`, `enable`.
* `send_version` - Send version. Valid values: `global`, `1`, `2`, `both`.
* `auth_keychain` - Authentication keychain name.
* `split_horizon` - Split horizon method. Valid values: `poisoned`, `regular`.
* `flags` - flags
* `auth_string` - Authentication string/password.
* `receive_version` - Receive version. Valid values: `global`, `1`, `2`, `both`.

The `distribute_list` block supports:

* `status` - Status. Valid values: `enable`, `disable`.
* `listname` - Distribute access/prefix list name.
* `direction` - Distribute list direction. Valid values: `in`, `out`.
* `interface` - Distribute list interface name.
* `id` - Distribute-list id.

The `network` block supports:

* `prefix` - Network prefix.
* `id` - Network entry id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Rip can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_rip.labelname RouterRip

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_rip.labelname RouterRip
$ unset "FORTISWITCH_IMPORT_TABLE"
```
