---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_stormcontrol"
description: |-
  Configure excess switch traffic (storm control).
---

# fortiswitch_switch_stormcontrol
Configure excess switch traffic (storm control).

## Argument Reference

The following arguments are supported:

* `broadcast` - Enable/disable broadcast storm control. Valid values: `enable`, `disable`.
* `unknown_unicast` - Enable/disable unknown unicast storm control. Valid values: `enable`, `disable`.
* `unknown_multicast` - Enable/disable unknown multicast storm control. Valid values: `enable`, `disable`.
* `rate` - Storm control traffic rate.
* `burst_size_level` - Storm control burst size level 0-4.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Switch StormControl can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_stormcontrol.labelname SwitchStormControl

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_stormcontrol.labelname SwitchStormControl
$ unset "FORTISWITCH_IMPORT_TABLE"
```
