---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_routerospf_network"
description: |-
  Enable OSPF on an IP network.
---

# fortiswitch_routerospf_network
Enable OSPF on an IP network.

~> The provider supports the definition of Network in Router Ospf `fortiswitch_router_ospf`, and also allows the definition of separate Network resources `fortiswitch_routerospf_network`, but do not use a `fortiswitch_router_ospf` with in-line Network in conjunction with any `fortiswitch_routerospf_network` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

* `prefix` - Prefix.
* `fswid` - Network entry ID.
* `area` - Attach the network to area.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

Routerospf Network can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_routerospf_network.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_routerospf_network.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
