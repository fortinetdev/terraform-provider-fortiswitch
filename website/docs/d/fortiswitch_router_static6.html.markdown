---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_static6"
description: |-
  Get information on an fortiswitch router static6.
---

# Data Source: fortiswitch_router_static6
Use this data source to get information on an fortiswitch router static6

## Argument Reference

* `seq_num` - (Required) Specify the seq_num of the desired router static6.

## Attribute Reference

The following attributes are exported:

* `status` - Status.
* `distance` - Administrative distance (1-255).
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `dst` - Destination IPv6 prefix for this route.
* `devindex` - devindex
* `seq_num` - entry No.
* `blackhole` - Enable/disable black hole.
* `vrf` - VRF.
* `device` - Gateway out interface.
* `gateway` - Gateway IPv6 address for this route.
* `comment` - comment

