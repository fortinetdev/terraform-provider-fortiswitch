---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_routerbgp_network6"
description: |-
  BGP IPv6 network table.
---

# fortiswitch_routerbgp_network6
BGP IPv6 network table.

~> The provider supports the definition of Network6 in Router Bgp `fortiswitch_router_bgp`, and also allows the definition of separate Network6 resources `fortiswitch_routerbgp_network6`, but do not use a `fortiswitch_router_bgp` with in-line Network6 in conjunction with any `fortiswitch_routerbgp_network6` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

* `route_map` - Route map to modify generated route.
* `fswid` - ID.
* `prefix6` - Network IPv6 prefix.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

Routerbgp Network6 can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_routerbgp_network6.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_routerbgp_network6.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
