---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchqos_dot1pmap"
description: |-
  QOS 802.1p configuration.
---

# fortiswitch_switchqos_dot1pmap
QOS 802.1p configuration.

## Argument Reference

The following arguments are supported:

* `name` - Dot1p map name.
* `priority_5` - COS queue mapped by priority number (PCP in 802.1q). Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `priority_4` - COS queue mapped by priority number (PCP in 802.1q). Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `priority_7` - COS queue mapped by priority number (PCP in 802.1q). Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `priority_6` - COS queue mapped by priority number (PCP in 802.1q). Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `priority_1` - COS queue mapped by priority number (PCP in 802.1q). Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `priority_0` - COS queue mapped by priority number (PCP in 802.1q). Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `priority_3` - COS queue mapped by priority number (PCP in 802.1q). Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `priority_2` - COS queue mapped by priority number (PCP in 802.1q). Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `egress_pri_tagging` - Enable/disable egress priority-tag frame. Valid values: `disable`, `enable`.
* `description` - Description of the 802.1p name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchQos Dot1PMap can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchqos_dot1pmap.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchqos_dot1pmap.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
