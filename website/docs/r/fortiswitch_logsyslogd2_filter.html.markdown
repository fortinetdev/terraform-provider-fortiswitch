---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_logsyslogd2_filter"
description: |-
  Filters for remote system server.
---

# fortiswitch_logsyslogd2_filter
Filters for remote system server.

## Example Usage

```hcl
resource "fortiswitch_logsyslogd2_filter" "name" {
        override = "test"
        severity = "emergency"
}
```

## Argument Reference

The following arguments are supported:

* `override` - Override FortiAnalyzer setting or use the global setting.
* `severity` - The least severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogSyslogd2 Filter can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_logsyslogd2_filter.labelname LogSyslogd2Filter

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_logsyslogd2_filter.labelname LogSyslogd2Filter
$ unset "FORTISWITCH_IMPORT_TABLE"
```
