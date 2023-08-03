---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_routerbgp_network"
description: |-
  BGP network table.
---

# fortiswitch_routerbgp_network
BGP network table.

~> The provider supports the definition of Network in Router Bgp `fortiswitch_router_bgp`, and also allows the definition of separate Network resources `fortiswitch_routerbgp_network`, but do not use a `fortiswitch_router_bgp` with in-line Network in conjunction with any `fortiswitch_routerbgp_network` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

* `backdoor` - Enable/disable route as backdoor. Valid values: `enable`, `disable`.
* `prefix` - Network prefix.
* `route_map` - Route map to modify generated route.
* `fswid` - ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

Routerbgp Network can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_routerbgp_network.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_routerbgp_network.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
