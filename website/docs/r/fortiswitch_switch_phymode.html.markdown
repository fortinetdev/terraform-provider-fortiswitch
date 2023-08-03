---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_phymode"
description: |-
  PHY configuration.
---

# fortiswitch_switch_phymode
PHY configuration.

## Argument Reference

The following arguments are supported:

* `port_configuration` - Port configurations for supporting different split-mode options. Valid values: `default`, `disable-port54`, `disable-port41-48`.
* `port54_phy_mode` - Split-mode configuration. Valid values: `1x40G`, `4x10G`.
* `port53_phy_mode` - Split-mode configuration. Valid values: `1x40G`, `4x10G`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Switch PhyMode can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_phymode.labelname SwitchPhyMode

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_phymode.labelname SwitchPhyMode
$ unset "FORTISWITCH_IMPORT_TABLE"
```
