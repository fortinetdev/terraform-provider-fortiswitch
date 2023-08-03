---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchstp_settings"
description: |-
  Switch-global stp settings.
---

# fortiswitch_switchstp_settings
Switch-global stp settings.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable spanning tree. Valid values: `enable`, `disable`.
* `hello_time` - Hello time.
* `name` - Name.
* `mclag_stp_mac` - MCLAG STP mac address.
* `pending_timer` - Pending time.
* `forward_time` - Forwarding delay.
* `flood` - Enable/disable stp bpdu flood. Valid values: `enable`, `disable`.
* `mclag_stp_bpdu` - MCLAG STP bpdu. Valid values: `both`, `single`.
* `max_hops` - Maximum number of hops.
* `max_age` - Maximum aging time.
* `revision` - Stp revision.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchStp Settings can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchstp_settings.labelname SwitchStpSettings

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchstp_settings.labelname SwitchStpSettings
$ unset "FORTISWITCH_IMPORT_TABLE"
```
