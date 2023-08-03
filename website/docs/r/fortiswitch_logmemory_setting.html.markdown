---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_logmemory_setting"
description: |-
  Settings for memory buffer.
---

# fortiswitch_logmemory_setting
Settings for memory buffer.

## Example Usage

```hcl
resource "fortiswitch_logmemory_setting" "name" {
        diskfull = "overwrite"
        status = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Whether to enable memory buffer log. Valid values: `enable`, `disable`.
* `diskfull` - What to do when memory is full. Valid values: `overwrite`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogMemory Setting can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_logmemory_setting.labelname LogMemorySetting

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_logmemory_setting.labelname LogMemorySetting
$ unset "FORTISWITCH_IMPORT_TABLE"
```
