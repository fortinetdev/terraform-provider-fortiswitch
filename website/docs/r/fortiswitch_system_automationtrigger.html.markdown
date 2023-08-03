---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_automationtrigger"
description: |-
  Trigger for automation stitches.
---

# fortiswitch_system_automationtrigger
Trigger for automation stitches. Applies to FortiSwitch Version `>= 7.2.0`.

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `trigger_minute` - Minute of the hour on which to trigger (0 - 59, default = 0).
* `fields` - Customized trigger field settings. The structure of `fields` block is documented below.
* `logid` - Log ID to trigger event.
* `trigger_hour` - Hour of the day on which to trigger (0 - 23, default = 1).
* `trigger_type` - Trigger type. Valid values: `event-based`, `scheduled`.
* `event_type` - Event type. Valid values: `config-change`, `event-log`, `reboot`.
* `trigger_frequency` - Scheduled trigger frequency (default = daily). Valid values: `hourly`, `daily`, `weekly`, `monthly`.
* `trigger_weekday` - Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
* `trigger_day` - Day within a month to trigger.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `fields` block supports:

* `id` - Entry ID.
* `value` - Value.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationTrigger can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_automationtrigger.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_automationtrigger.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
