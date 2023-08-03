---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_logmemory_globalsetting"
description: |-
  Global settings for memory log.
---

# fortiswitch_logmemory_globalsetting
Global settings for memory log.

## Example Usage

```hcl
resource "fortiswitch_logmemory_globalsetting" "name" {
        full_final_warning_threshold = "92"
        full_first_warning_threshold = "90"
        full_second_warning_threshold = "91"
        hourly_upload = "enable"
        max_size = "98305"
}
```

## Argument Reference

The following arguments are supported:

* `max_size` - Maximum size(byte) in memory buffer log.
* `full_final_warning_threshold` - Log full final warning threshold(3-100), the default is 95.
* `full_first_warning_threshold` - Log full first warning threshold(1-98), the default is 75.
* `full_second_warning_threshold` - Log full second warning threshold(2-99), the default is 90.
* `hourly_upload` - Whether to enable scheduled upload hourly. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogMemory GlobalSetting can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_logmemory_globalsetting.labelname LogMemoryGlobalSetting

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_logmemory_globalsetting.labelname LogMemoryGlobalSetting
$ unset "FORTISWITCH_IMPORT_TABLE"
```
