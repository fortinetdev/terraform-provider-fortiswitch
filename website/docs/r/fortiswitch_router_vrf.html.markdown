---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_vrf"
description: |-
  VRF configuration.
---

# fortiswitch_router_vrf
VRF configuration.

## Example Usage

```hcl
resource "fortiswitch_router_vrf" "name" {
        name = "default_name_3"
        vrfid = "4"
}
```

## Argument Reference

The following arguments are supported:

* `vrfid` - VRF ID.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router Vrf can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_vrf.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_vrf.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
