---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_virtualwire"
description: |-
  Configure virtual wire.
---

# fortiswitch_switch_virtualwire
Configure virtual wire.

## Argument Reference

The following arguments are supported:

* `first_member` - First member of virtual wire pair.
* `vlan` - VLAN (can be shared between virtual wires and non-virtual wire ports).
* `name` - Virtual wire name.
* `second_member` - Second member of virtual wire pair.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Switch VirtualWire can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_virtualwire.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_virtualwire.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
