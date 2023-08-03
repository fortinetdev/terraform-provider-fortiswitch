---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemschedule_onetime"
description: |-
  onetime schedule configuration
---

# fortiswitch_systemschedule_onetime
onetime schedule configuration

## Example Usage

```hcl
resource "fortiswitch_systemschedule_onetime" "name" {
        end = "12:12 2020/12/13"
        name = "recurring"
        start = "12:12 2020/12/12"
}
```

## Argument Reference

The following arguments are supported:

* `start` - Start time and date.
* `end` - End time and date.
* `name` - Onetime schedule name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemSchedule Onetime can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemschedule_onetime.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemschedule_onetime.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
