---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemptp_interfacepolicy"
description: |-
  PTP policy configuration.
---

# fortiswitch_systemptp_interfacepolicy
PTP policy configuration. Applies to FortiSwitch Version `>= 7.2.5`.

## Argument Reference

The following arguments are supported:

* `vlan` - PTP Vlan (0-4094)
* `vlan_pri` - PTP Vlan Priority (0-7)
* `name` - Policy name.
* `description` - Description.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemPtp InterfacePolicy can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemptp_interfacepolicy.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemptp_interfacepolicy.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
