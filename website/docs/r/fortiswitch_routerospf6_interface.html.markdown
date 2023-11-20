---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_routerospf6_interface"
description: |-
  OSPF6 interface configuration.
---

# fortiswitch_routerospf6_interface
OSPF6 interface configuration.

~> The provider supports the definition of Interface in Router Ospf6 `fortiswitch_router_ospf6`, and also allows the definition of separate Interface resources `fortiswitch_routerospf6_interface`, but do not use a `fortiswitch_router_ospf6` with in-line Interface in conjunction with any `fortiswitch_routerospf6_interface` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

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
* `id` - an identifier for the resource with format {{name}}.

## Import

Routerospf6 Interface can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_routerospf6_interface.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_routerospf6_interface.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
