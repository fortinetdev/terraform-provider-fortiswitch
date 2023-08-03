---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchptp_policy"
description: |-
  PTP policy configuration.
---

# fortiswitch_switchptp_policy
PTP policy configuration.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable PTP policy. Valid values: `enable`, `disable`.
* `name` - Policy name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchPtp Policy can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchptp_policy.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchptp_policy.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
