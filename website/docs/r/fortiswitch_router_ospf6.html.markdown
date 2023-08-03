---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_ospf6"
description: |-
  Router OSPF6 configuration.
---

# fortiswitch_router_ospf6
Router OSPF6 configuration.

## Example Usage

```hcl
resource "fortiswitch_router_ospf6" "name" {
        area {
            id =  "8"
            range {
                advertise = "disable"
                id =  "11"
                prefix6 = "ffff::"
            }
            stub_type = "no_summary"
            type = "regular"
        }
        interface {
            area_id = "12"
            bfd = "enable"
            cost = "18"
            dead_interval = "19"
            hello_interval = "20"
            name = "name_21"
            passive = "enable"
            priority = "23"
            retransmit_interval = "24"
            status = "disable"
            transmit_delay = "26"
        }
        log_neighbor_changes = "enable"
        redistribute {
            metric = "29"
            metric_type = "1"
            name = "name"
            routemap = "test"
            status = "enable"
        }
        router_id = "33"
}
```

## Argument Reference

The following arguments are supported:

* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `router_id` - A.B.C.D, in IPv4 address format.
* `spf_timers` - SPF calculation frequency.
* `area` - OSPF6 area configuration. The structure of `area` block is documented below.
* `log_neighbor_changes` - Enable logging of OSPF neighbor's changes. Valid values: `enable`, `disable`.
* `interface` - OSPF6 interface configuration. The structure of `interface` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `redistribute` block supports:

* `status` - status Valid values: `enable`, `disable`.
* `metric_type` - metric type Valid values: `1`, `2`.
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `name` - Redistribute name.

The `area` block supports:

* `stub_type` - Stub summary setting. Valid values: `no-summary`, `summary`.
* `filter_list` - OSPF area filter-list configuration. The structure of `filter_list` block is documented below.
* `range` - OSPF6 area range configuration. The structure of `range` block is documented below.
* `type` - Area type setting. Valid values: `regular`, `stub`.
* `id` - Area entry ip address.

The `filter_list` block supports:

* `direction` - Direction. Valid values: `in`, `out`.
* `list` - Access-list or prefix-list name.
* `id` - Filter list entry ID.

The `range` block supports:

* `advertise` - Enable/disable advertise status. Valid values: `disable`, `enable`.
* `id` - Range entry id.
* `prefix6` - <prefix6>   IPv6 prefix

The `interface` block supports:

* `status` - Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `name` - Interface name.
* `bfd` - Enable/Disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
* `area_id` - A.B.C.D, in IPv4 address format.
* `transmit_delay` - Link state transmit delay.
* `retransmit_interval` - Time between retransmitting lost link state advertisements.
* `cost` - The cost of the interface.
* `passive` - Enable/disable passive interface. Valid values: `enable`, `disable`.
* `priority` - Router priority.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Ospf6 can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_ospf6.labelname RouterOspf6

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_ospf6.labelname RouterOspf6
$ unset "FORTISWITCH_IMPORT_TABLE"
```
