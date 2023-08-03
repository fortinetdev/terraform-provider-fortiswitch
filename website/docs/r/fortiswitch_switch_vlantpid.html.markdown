---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_vlantpid"
description: |-
  Configure switch global ether-types.
---

# fortiswitch_switch_vlantpid
Configure switch global ether-types.

## Argument Reference

The following arguments are supported:

* `name` - User given name for ether-type.
* `ether_type` - Set switch ether-type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Switch VlanTpid can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_vlantpid.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_vlantpid.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
