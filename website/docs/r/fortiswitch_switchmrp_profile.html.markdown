---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchmrp_profile"
description: |-
  MRP configuration profiles.
---

# fortiswitch_switchmrp_profile
MRP configuration profiles. Applies to FortiSwitch Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `name` - Profile name.
* `topology_change_interval` - The interval for sending MRP_TopologyChange.
* `short_test_interval` - The short interval for sending MRP_Test after link changes in the ring.
* `default_test_interval` - The default interval for sending MRP_Test.
* `topology_change_repeat_count` - The count for repeated MRP_TopologyChange transmissions.
* `test_monitoring_count` - The count for monitoring MRP_Test received.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchMrp Profile can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchmrp_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchmrp_profile.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
