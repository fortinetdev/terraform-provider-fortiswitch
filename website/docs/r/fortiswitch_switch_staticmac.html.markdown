---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_staticmac"
description: |-
  Switch static mac address entries.
---

# fortiswitch_switch_staticmac
Switch static mac address entries.

## Argument Reference

The following arguments are supported:

* `description` - Description.
* `seq_num` - Sequence number.
* `vlan_id` - VLAN ID.
* `action` - Action. Valid values: `allow`, `drop`.
* `mac` - Static mac address.
* `interface` - Interface name.
* `type` - Type. Valid values: `static`, `sticky`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Switch StaticMac can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_staticmac.labelname {{seq_num}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_staticmac.labelname {{seq_num}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
