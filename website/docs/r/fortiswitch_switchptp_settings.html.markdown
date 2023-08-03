---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchptp_settings"
description: |-
  Global PTP configuration.
---

# fortiswitch_switchptp_settings
Global PTP configuration.

## Argument Reference

The following arguments are supported:

* `mode` - Disable/enable PTP mode. Valid values: `disable`, `transparent-e2e`, `transparent-p2p`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchPtp Settings can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchptp_settings.labelname SwitchPtpSettings

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchptp_settings.labelname SwitchPtpSettings
$ unset "FORTISWITCH_IMPORT_TABLE"
```
