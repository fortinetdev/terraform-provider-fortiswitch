---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_log_eventfilter"
description: |-
  Log event filter configuration.
---

# fortiswitch_log_eventfilter
Log event filter configuration.

## Example Usage

```hcl
resource "fortiswitch_log_eventfilter" "name" {
        event = "enable"
        fos_legacy = "enable"
        link = "enable"
        poe = "disable"
        router = "enable"
        spanning_tree = "enable"
        switch = "enable"
        switch_controller = "enable"
        system = "enable"
        user = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `spanning_tree` - Enable/disable logging of spanning_tree activity messages. Valid values: `enable`, `disable`.
* `system` - Enable/disable logging of system activity messages. Valid values: `enable`, `disable`.
* `switch_controller` - Enable/disable logging of switch_controller activity messages. Valid values: `enable`, `disable`.
* `switch` - Enable/disable logging of switch activity messages. Valid values: `enable`, `disable`.
* `link` - Enable/disable logging of link activity messages. Valid values: `enable`, `disable`.
* `user` - Enable/disable logging of user activity messages. Valid values: `enable`, `disable`.
* `router` - Enable/disable logging of router activity messages. Valid values: `enable`, `disable`.
* `fos_legacy` - Enable/disable logging of fos legacy activity messages. Valid values: `enable`, `disable`.
* `poe` - Enable/disable logging of poe activity messages. Valid values: `enable`, `disable`.
* `event` - Enable/disable logging of event messages. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log Eventfilter can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_log_eventfilter.labelname LogEventfilter

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_log_eventfilter.labelname LogEventfilter
$ unset "FORTISWITCH_IMPORT_TABLE"
```
