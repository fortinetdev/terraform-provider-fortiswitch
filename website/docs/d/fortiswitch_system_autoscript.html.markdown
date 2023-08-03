---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_autoscript"
description: |-
  Get information on an fortiswitch system autoscript.
---

# Data Source: fortiswitch_system_autoscript
Use this data source to get information on an fortiswitch system autoscript

## Argument Reference

* `name` - (Required) Specify the name of the desired system autoscript.

## Attribute Reference

The following attributes are exported:

* `output_size` - Number of megabytes to limit script output to (10 - 1024, default = 10).
* `repeat` - Number of times to repeat this script (0 = infinite).
* `name` - Auto script name.
* `script` - List of CLI commands to repeat.
* `interval` - Repeat interval in seconds.
* `start` - Script starting mode.

