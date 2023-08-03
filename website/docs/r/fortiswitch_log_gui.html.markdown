---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_log_gui"
description: |-
  Logging device to display in GUI.
---

# fortiswitch_log_gui
Logging device to display in GUI.

## Example Usage

```hcl
resource "fortiswitch_log_gui" "name" {
        log_device = "memory"
}
```

## Argument Reference

The following arguments are supported:

* `log_device` - Log device.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log Gui can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_log_gui.labelname LogGui

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_log_gui.labelname LogGui
$ unset "FORTISWITCH_IMPORT_TABLE"
```
