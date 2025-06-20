---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_policy"
description: |-
  Get information on an fortiswitch router policy.
---

# Data Source: fortiswitch_router_policy
Use this data source to get information on an fortiswitch router policy

## Argument Reference

* `seq_num` - (Required) Specify the seq_num of the desired router policy.

## Attribute Reference

The following attributes are exported:

* `interface` - Interface configuration. The structure of `interface` block is documented below.
* `nexthop_group` - Nexthop group (ECMP) configuration. The structure of `nexthop_group` block is documented below.
* `pbr_map` - PBR map configuration. The structure of `pbr_map` block is documented below.
* `comments` - Description/comments.
* `src` - Source ip and mask.
* `output_device` - Outgoing interface name.
* `protocol` - Protocol number.
* `end_port` - End port number.
* `dst` - Destination ip and mask.
* `seq_num` - Sequence number.
* `tos_mask` - Terms of service evaluated bits.
* `input_device` - Incoming interface name.
* `tos` - Terms of service bit pattern.
* `gateway` - IP address of gateway.
* `start_port` - Start port number.

The `interface` block contains:

* `name` - Interface name
* `pbr_map_name` - PBR policy map name.

The `nexthop_group` block contains:

* `nexthop` - Nexthop configuration. The structure of `nexthop` block is documented below.
* `name` - Name.

The `nexthop` block contains:

* `nexthop_vrf_name` - VRF name.
* `nexthop_ip` - IP address of nexthop.
* `id` - Id (1-64).

The `pbr_map` block contains:

* `rule` - Rule. The structure of `rule` block is documented below.
* `name` - Name.
* `comments` - Description/comments.

The `rule` block contains:

* `src` - Source ip and mask.
* `nexthop_vrf_name` - Nexthop vrf name.
* `nexthop_group_name` - Nexthop group name. Used for ECMP.
* `dst` - Destination ip and mask.
* `nexthop_ip` - IP address of nexthop.
* `seq_num` - Rule seq-num (1-10000).

