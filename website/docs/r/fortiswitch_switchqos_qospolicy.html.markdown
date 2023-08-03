---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchqos_qospolicy"
description: |-
  QOS egress policy.
---

# fortiswitch_switchqos_qospolicy
QOS egress policy.

## Argument Reference

The following arguments are supported:

* `rate_by` - COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
* `name` - QOS policy name.
* `cos_queue` - COS queue configuration. The structure of `cos_queue` block is documented below.
* `schedule` - COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `cos_queue` block supports:

* `max_rate_percent` - Maximum rate (% of link speed).
* `name` - COS queue ID.
* `weight` - Weight of weighted round robin scheduling.
* `min_rate_percent` - Minimum rate (% of link speed).
* `ecn` - Update frame IP ECN field in lieu of packet drop. Valid values: `disable`, `enable`.
* `min_rate` - Minimum rate (kbps). 0 to disable.
* `wred_slope` - Slope of WRED drop probability.
* `max_rate` - Maximum rate (kbps). 0 to disable.
* `drop_policy` - COS queue drop policy. Valid values: `taildrop`, `weighted-random-early-detection`.
* `description` - Description of the COS queue.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchQos QosPolicy can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchqos_qospolicy.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchqos_qospolicy.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
