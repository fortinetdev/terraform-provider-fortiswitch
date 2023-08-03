---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchnetworkmonitor_settings"
description: |-
  Global configuration of network monitoring on the switch.
---

# fortiswitch_switchnetworkmonitor_settings
Global configuration of network monitoring on the switch.

## Argument Reference

The following arguments are supported:

* `db_aging_interval` - Network monitor database aging interval (3600 - 86400 sec, default = 3600 sec, 0 = disable). 
* `status` - Enable/disable network monitor. Valid values: `enable`, `disable`.
* `survey_mode_interval` - Duration for which a network monitor is programmed in hardware in survey mode(120 - 3600 sec, default = 120 sec). 
* `survey_mode` - Enable/disable network monitor survey mode. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchNetworkMonitor Settings can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchnetworkmonitor_settings.labelname SwitchNetworkMonitorSettings

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchnetworkmonitor_settings.labelname SwitchNetworkMonitorSettings
$ unset "FORTISWITCH_IMPORT_TABLE"
```
