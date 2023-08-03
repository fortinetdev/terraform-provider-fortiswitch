---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_logfortiguard_setting"
description: |-
  Settings for FortiGuard Analysis Service.
---

# fortiswitch_logfortiguard_setting
Settings for FortiGuard Analysis Service.

## Example Usage

```hcl
resource "fortiswitch_logfortiguard_setting" "name" {
        enc_algorithm = "default"
        status = "enable"
        upload_day = "3"
        upload_interval = "daily"
        upload_time = "12:23"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Whether to enable FortiGuard Analysis Service. Valid values: `enable`, `disable`.
* `upload_interval` - Frequency to check log file for upload. Valid values: `daily`, `weekly`, `monthly`.
* `upload_time` - Time to roll logs [hh:mm].
* `upload_day` - Days of week to roll logs.
* `enc_algorithm` - Whether to send FortiAnalyzer log data with SSL encryption. Valid values: `default`, `high`, `low`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogFortiguard Setting can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_logfortiguard_setting.labelname LogFortiguardSetting

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_logfortiguard_setting.labelname LogFortiguardSetting
$ unset "FORTISWITCH_IMPORT_TABLE"
```
