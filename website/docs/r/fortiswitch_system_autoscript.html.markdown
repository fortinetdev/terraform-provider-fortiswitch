---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_autoscript"
description: |-
  Configure auto script.
---

# fortiswitch_system_autoscript
Configure auto script. Applies to FortiSwitch Version `7.2.0,7.2.1,7.2.2,7.2.3,7.2.4,7.2.5,7.4.0,7.4.1,7.4.2,7.4.3,7.6.0,7.6.1,7.6.2,7.6.3,7.6.4,8.0.0`.

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
