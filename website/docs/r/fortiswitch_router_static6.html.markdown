---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_static6"
description: |-
  Ipv6 static routes configuration.
---

# fortiswitch_router_static6
Ipv6 static routes configuration.

## Argument Reference

The following arguments are supported:

* `status` - Status. Valid values: `enable`, `disable`.
* `distance` - Administrative distance (1-255).
* `bfd` - Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
* `dst` - Destination IPv6 prefix for this route.
* `devindex` - devindex
* `seq_num` - entry No.
* `blackhole` - Enable/disable black hole. Valid values: `enable`, `disable`.
* `vrf` - VRF.
* `device` - Gateway out interface.
* `gateway` - Gateway IPv6 address for this route.
* `comment` - comment


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Router Static6 can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_static6.labelname {{seq_num}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_static6.labelname {{seq_num}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
