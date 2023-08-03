---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_logsyslogd_overridefilter"
description: |-
  Override filters for remote system server.
---

# fortiswitch_logsyslogd_overridefilter
Override filters for remote system server.

## Argument Reference

The following arguments are supported:

* `override` - Override FortiAnalyzer setting or use the global setting. Valid values: `enable`, `disable`.
* `severity` - The least severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogSyslogd OverrideFilter can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_logsyslogd_overridefilter.labelname LogSyslogdOverrideFilter

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_logsyslogd_overridefilter.labelname LogSyslogdOverrideFilter
$ unset "FORTISWITCH_IMPORT_TABLE"
```
