---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_vlanpruning"
description: |-
  Vlan Pruning.
---

# fortiswitch_switch_vlanpruning
Vlan Pruning. Applies to FortiSwitch Version `>= 7.6.1`.

## Argument Reference

The following arguments are supported:

* `join_timer` - Vlan Pruning Join Timer (in millisecond).
* `leave_timer` - Vlan Pruning Leave Timer (in millisecond).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Switch VlanPruning can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_vlanpruning.labelname SwitchVlanPruning

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_vlanpruning.labelname SwitchVlanPruning
$ unset "FORTISWITCH_IMPORT_TABLE"
```
