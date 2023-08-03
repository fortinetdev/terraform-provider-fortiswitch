---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_automationtrigger"
description: |-
  Get information on an fortiswitch system automationtrigger.
---

# Data Source: fortiswitch_system_automationtrigger
Use this data source to get information on an fortiswitch system automationtrigger

## Argument Reference

* `name` - (Required) Specify the name of the desired system automationtrigger.

## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `trigger_minute` - Minute of the hour on which to trigger (0 - 59, default = 0).
* `fields` - Customized trigger field settings. The structure of `fields` block is documented below.
* `logid` - Log ID to trigger event.
* `trigger_hour` - Hour of the day on which to trigger (0 - 23, default = 1).
* `trigger_type` - Trigger type.
* `event_type` - Event type.
* `trigger_frequency` - Scheduled trigger frequency (default = daily).
* `trigger_weekday` - Day of week for trigger.
* `trigger_day` - Day within a month to trigger.

The `fields` block contains:

* `id` - Entry ID.
* `value` - Value.
* `name` - Name.

