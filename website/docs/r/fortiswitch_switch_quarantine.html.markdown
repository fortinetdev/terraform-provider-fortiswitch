---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_quarantine"
description: |-
  Configure quarantine devices on the switch.
---

# fortiswitch_switch_quarantine
Configure quarantine devices on the switch.

## Argument Reference

The following arguments are supported:

* `description` - Description for the quarantine MAC.
* `cos_queue` - COS queue for the quarantine device traffic from range (0 - 7) or unset to disable.
* `drop` - Drop setting for the quarantine device traffic. Valid values: `disable`, `enable`.
* `mac` - Quarantine MAC.
* `policer` - ACL policer for the quarantine device traffic.
* `acl_id` - Associated ingress acl for the quarantine MAC.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mac}}.

## Import

Switch Quarantine can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_quarantine.labelname {{mac}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_quarantine.labelname {{mac}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
