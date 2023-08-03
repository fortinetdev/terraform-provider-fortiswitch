---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_ipmacbinding"
description: |-
  Ip-mac-binding table.
---

# fortiswitch_switch_ipmacbinding
Ip-mac-binding table.

## Argument Reference

The following arguments are supported:

* `status` - Is the binding now enabled. Valid values: `enable`, `disable`.
* `ip` - Source ip and mask for this rule.
* `mac` - MAC address value.
* `seq_num` - Entry No.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Switch IpMacBinding can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_ipmacbinding.labelname {{seq_num}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_ipmacbinding.labelname {{seq_num}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
