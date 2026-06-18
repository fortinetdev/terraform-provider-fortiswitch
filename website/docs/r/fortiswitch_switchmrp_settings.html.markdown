---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchmrp_settings"
description: |-
  Media Redundancy Protocol configuration.
---

# fortiswitch_switchmrp_settings
Media Redundancy Protocol configuration. Applies to FortiSwitch Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `fswid` - MRP ring ID.
* `status` - Enable/disable MRP. Valid values: `enable`, `disable`.
* `role` - The admin role of MRP (Client/Automanager). Valid values: `client`, `automanager`.
* `domain_id` - A unique UUID that represents the MRP ring.
* `domain_name` - A unique logical name for the MRP domain-id.
* `vlan_id` - The VLAN ID for sending MRP frames.
* `priority` - The value for the manager priority.
* `ring_port1` - The physical port that serves as the first ring port.
* `ring_port2` - The physical port that serves as the second ring port.
* `profile_name` - A unique MRP profile name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

SwitchMrp Settings can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchmrp_settings.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchmrp_settings.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
