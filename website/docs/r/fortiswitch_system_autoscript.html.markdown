---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_autoscript"
description: |-
  Configure auto script.
---

# fortiswitch_system_autoscript
Configure auto script. Applies to FortiSwitch Version `>= 7.2.0`.

## Argument Reference

The following arguments are supported:

* `output_size` - Number of megabytes to limit script output to (10 - 1024, default = 10).
* `repeat` - Number of times to repeat this script (0 = infinite).
* `name` - Auto script name.
* `script` - List of CLI commands to repeat.
* `interval` - Repeat interval in seconds.
* `start` - Script starting mode. Valid values: `manual`, `auto`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutoScript can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_autoscript.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_autoscript.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
