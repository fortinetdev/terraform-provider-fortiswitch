---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_static"
description: |-
  Get information on an fortiswitch router static.
---

# Data Source: fortiswitch_router_static
Use this data source to get information on an fortiswitch router static

## Argument Reference

* `seq_num` - (Required) Specify the seq_num of the desired router static.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `distance` - Administrative distance (1-255).
* `weight` - Administrative weight (0-255).
* `dynamic_gateway` - Dynamic-gateway.
* `dst` - Destination ip and mask for this route.
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `seq_num` - Entry No.
* `blackhole` - Blackhole.
* `priority` - Administrative priority (0-4294967295).
* `status` - Status.
* `vrf` - VRF.
* `device` - Gateway out interface.
* `gateway` - Gateway ip for this route.

