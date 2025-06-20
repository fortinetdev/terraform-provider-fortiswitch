---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_policy"
description: |-
  Policy routing configuration.
---

# fortiswitch_router_policy
Policy routing configuration.

## Example Usage

```hcl
resource "fortiswitch_router_policy" "name" {
        comments = "test"
        dst = "1.2.0.0/16"
        end_port = "5"
        gateway = "192.160.44.33"
        input_device = "port5"
        protocol = "28"
        seq_num = "29"
        start_port = "31"
}
```

## Argument Reference

The following arguments are supported:

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
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `interface` block supports:

* `name` - Interface name
* `pbr_map_name` - PBR policy map name.

The `nexthop_group` block supports:

* `nexthop` - Nexthop configuration. The structure of `nexthop` block is documented below.
* `name` - Name.

The `nexthop` block supports:

* `nexthop_vrf_name` - VRF name.
* `nexthop_ip` - IP address of nexthop.
* `id` - Id (1-64).

The `pbr_map` block supports:

* `rule` - Rule. The structure of `rule` block is documented below.
* `name` - Name.
* `comments` - Description/comments.

The `rule` block supports:

* `src` - Source ip and mask.
* `nexthop_vrf_name` - Nexthop vrf name.
* `nexthop_group_name` - Nexthop group name. Used for ECMP.
* `dst` - Destination ip and mask.
* `nexthop_ip` - IP address of nexthop.
* `seq_num` - Rule seq-num (1-10000).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Router Policy can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_policy.labelname {{seq_num}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_policy.labelname {{seq_num}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
