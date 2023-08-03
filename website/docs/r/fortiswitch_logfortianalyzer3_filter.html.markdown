---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_logfortianalyzer3_filter"
description: |-
  Filters for FortiAnalyzer.
---

# fortiswitch_logfortianalyzer3_filter
Filters for FortiAnalyzer.

## Example Usage

```hcl
resource "fortiswitch_logfortianalyzer3_filter" "name" {
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

LogFortianalyzer3 Filter can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_logfortianalyzer3_filter.labelname LogFortianalyzer3Filter

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_logfortianalyzer3_filter.labelname LogFortianalyzer3Filter
$ unset "FORTISWITCH_IMPORT_TABLE"
```
