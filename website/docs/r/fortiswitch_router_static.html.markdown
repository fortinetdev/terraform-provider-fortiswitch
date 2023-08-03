---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_static"
description: |-
  IPv4 static routes configuration.
---

# fortiswitch_router_static
IPv4 static routes configuration.

## Example Usage

```hcl
resource "fortiswitch_router_static" "name" {
        bfd = "global"
        blackhole = "enable"
        comment = "Comment."
        device = "internal"
        distance = "7"
        dst = "10.10.10.0 255.255.255.255"
        dynamic_gateway = "enable"
        gateway = "20.20.20.20"
        priority = "11"
        seq_num = "12"
        status = "enable"
        weight = "15"
}
```

## Argument Reference

The following arguments are supported:

* `comment` - Comment.
* `distance` - Administrative distance (1-255).
* `weight` - Administrative weight (0-255).
* `dynamic_gateway` - Dynamic-gateway. Valid values: `enable`, `disable`.
* `dst` - Destination ip and mask for this route.
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `seq_num` - Entry No.
* `blackhole` - Blackhole. Valid values: `enable`, `disable`.
* `priority` - Administrative priority (0-4294967295).
* `status` - Status. Valid values: `enable`, `disable`.
* `vrf` - VRF.
* `device` - Gateway out interface.
* `gateway` - Gateway ip for this route.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Router Static can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_static.labelname {{seq_num}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_static.labelname {{seq_num}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
