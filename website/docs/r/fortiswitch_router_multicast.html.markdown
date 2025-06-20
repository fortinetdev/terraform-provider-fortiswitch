---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_multicast"
description: |-
  Router multicast configuration.
---

# fortiswitch_router_multicast
Router multicast configuration.

## Example Usage

```hcl
resource "fortiswitch_router_multicast" "name" {
        interface {
            dr_priority = "4"
            hello_interval = "5"
            name = "port4"
            pim_mode = "ssm-mode"
        }
        multicast_routing = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `comments` - Description/comments.
* `interface` - Pim interfaces. The structure of `interface` block is documented below.
* `multicast_routing` - Enable multicast routing. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `interface` block supports:

* `hello_interval` - Interval between sending PIM hello messages(sec).
* `name` - Interface name.
* `pim_mode` - PIM operation mode. Valid values: `ssm-mode`.
* `dr_priority` - DR election priority.
* `multicast_flow` - Config acceptable source for multicast group.
* `igmp` - Igmp configuration options. The structure of `igmp` block is documented below.

The `igmp` block supports:

* `query_interval` - Interval between queries to IGMP hosts(sec).
* `query_max_response_time` - Maximum time to wait for a IGMP query response(sec).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Multicast can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_multicast.labelname RouterMulticast

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_multicast.labelname RouterMulticast
$ unset "FORTISWITCH_IMPORT_TABLE"
```
