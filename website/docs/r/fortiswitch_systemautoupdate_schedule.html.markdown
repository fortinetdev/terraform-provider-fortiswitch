---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemautoupdate_schedule"
description: |-
  Configure update schedule.
---

# fortiswitch_systemautoupdate_schedule
Configure update schedule.

## Example Usage

```hcl
resource "fortiswitch_systemautoupdate_schedule" "saos" {
        day = "Sunday"
        frequency = "every"
        status = "enable"
        time = "11:13"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable scheduled updates. Valid values: `enable`, `disable`.
* `frequency` - Configure the update frequency. Valid values: `every`, `daily`, `weekly`.
* `day` - Configure the update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
* `time` - Configure the update time.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SystemAutoupdate Schedule can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemautoupdate_schedule.labelname SystemAutoupdateSchedule

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemautoupdate_schedule.labelname SystemAutoupdateSchedule
$ unset "FORTISWITCH_IMPORT_TABLE"
```
