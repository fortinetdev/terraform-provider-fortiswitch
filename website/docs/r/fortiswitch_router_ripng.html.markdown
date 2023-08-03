---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_ripng"
description: |-
  router ripng configuration
---

# fortiswitch_router_ripng
router ripng configuration

## Example Usage

```hcl
resource "fortiswitch_router_ripng" "name" {
        bfd = "enable"
        default_information_originate = "enable"
        interface {
            flags = "17"
            name = "port2"
            passive = "enable"
        }
}
```

## Argument Reference

The following arguments are supported:

* `default_metric` - Default metric of redistribute routes (Except connected).
* `timeout_timer` - Routing information timeout timer.
* `aggregate_address` - Set aggregate RIPng route announcement. The structure of `aggregate_address` block is documented below.
* `offset_list` - Offset list to modify RIPng metric. The structure of `offset_list` block is documented below.
* `bfd` - Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `garbage_timer` - Garbage collection timer.
* `default_information_originate` - Generate a default route. Valid values: `enable`, `disable`.
* `interface` - RIPng interface configuration. The structure of `interface` block is documented below.
* `update_timer` - Routing table update timer.
* `distribute_list` - Filter networks in routing updates. The structure of `distribute_list` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `aggregate_address` block supports:

* `id` - Aggregate-address entry id.
* `prefix6` - Aggregate-address prefix.

The `offset_list` block supports:

* `status` - Status. Valid values: `enable`, `disable`.
* `direction` - Offset list direction. Valid values: `in`, `out`.
* `access_list6` - Ipv6 access list name.
* `offset` - Metric offset.
* `interface` - Interface name.
* `id` - Offset-list id.

The `redistribute` block supports:

* `status` - status Valid values: `enable`, `disable`.
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `flags` - Flags
* `name` - Redistribute name.

The `interface` block supports:

* `passive` - Suppress routing updates on an interface. Valid values: `enable`, `disable`.
* `split_horizon_status` - Enable/disable split horizon. Valid values: `enable`, `disable`.
* `split_horizon` - Split horizon type. Valid values: `poisoned`, `regular`.
* `flags` - Flags.
* `name` - Interface name.

The `distribute_list` block supports:

* `status` - Status. Valid values: `enable`, `disable`.
* `listname` - Distribute access/prefix list name.
* `direction` - Distribute list direction. Valid values: `in`, `out`.
* `interface` - Distribute list interface name.
* `id` - Distribute-list id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Ripng can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_ripng.labelname RouterRipng

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_ripng.labelname RouterRipng
$ unset "FORTISWITCH_IMPORT_TABLE"
```
