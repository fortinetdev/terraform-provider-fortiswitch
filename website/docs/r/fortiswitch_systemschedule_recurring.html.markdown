---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemschedule_recurring"
description: |-
  recurring schedule configuration
---

# fortiswitch_systemschedule_recurring
recurring schedule configuration

## Example Usage

```hcl
resource "fortiswitch_systemschedule_recurring" "name" {
        day = "sunday"
        end = "23:59"
        name = "recurring"
        start = "00:00"
}
```

## Argument Reference

The following arguments are supported:

* `start` - Start time.
* `end` - End time.
* `name` - Recurring schedule name.
* `day` - Weekday. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemSchedule Recurring can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemschedule_recurring.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemschedule_recurring.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
