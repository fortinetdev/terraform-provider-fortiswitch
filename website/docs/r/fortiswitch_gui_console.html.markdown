---
subcategory: "FortiSwitch Gui"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_gui_console"
description: |-
  Dashboard CLI console configuration.
---

# fortiswitch_gui_console
Dashboard CLI console configuration.

## Argument Reference

The following arguments are supported:

* `preferences` - Preferences.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Gui Console can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_gui_console.labelname GuiConsole

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_gui_console.labelname GuiConsole
$ unset "FORTISWITCH_IMPORT_TABLE"
```
