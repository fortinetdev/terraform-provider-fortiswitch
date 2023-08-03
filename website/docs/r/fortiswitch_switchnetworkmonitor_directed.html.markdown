---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchnetworkmonitor_directed"
description: |-
  Configuration of the static entries for network monitoring on the switch.
---

# fortiswitch_switchnetworkmonitor_directed
Configuration of the static entries for network monitoring on the switch.

## Example Usage

```hcl
resource "fortiswitch_switchnetworkmonitor_directed" "name" {
        fswid =  "5"
        monitor_mac = "00:00:5e:00:53:af"
}
```

## Argument Reference

The following arguments are supported:

* `fswid` - Network monitor list entry ID.
* `monitor_mac` - MAC address to be  monitored.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

SwitchNetworkMonitor Directed can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchnetworkmonitor_directed.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchnetworkmonitor_directed.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
