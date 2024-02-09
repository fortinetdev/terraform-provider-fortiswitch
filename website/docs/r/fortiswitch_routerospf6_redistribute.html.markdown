---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_routerospf6_redistribute"
description: |-
  Redistribute configuration.
---

# fortiswitch_routerospf6_redistribute
Redistribute configuration.

~> The provider supports the definition of Redistribute in Router Ospf6 `fortiswitch_router_ospf6`, and also allows the definition of separate Redistribute resources `fortiswitch_routerospf6_redistribute`, but do not use a `fortiswitch_router_ospf6` with in-line Redistribute in conjunction with any `fortiswitch_routerospf6_redistribute` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

* `status` - status Valid values: `enable`, `disable`.
* `metric_type` - metric type Valid values: `1`, `2`.
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `name` - Redistribute name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Routerospf6 Redistribute can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_routerospf6_redistribute.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_routerospf6_redistribute.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
