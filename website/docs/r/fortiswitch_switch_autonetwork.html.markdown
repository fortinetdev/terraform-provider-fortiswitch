---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_autonetwork"
description: |-
  Auto network.
---

# fortiswitch_switch_autonetwork
Auto network.

## Example Usage

```hcl
resource "fortiswitch_switch_autonetwork" "name" {
        mgmt_vlan = "3"
        status =  "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable auto network status. Valid values: `enable`, `disable`.
* `mgmt_vlan` - Auto network management VLAN.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Switch AutoNetwork can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_autonetwork.labelname SwitchAutoNetwork

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_autonetwork.labelname SwitchAutoNetwork
$ unset "FORTISWITCH_IMPORT_TABLE"
```
