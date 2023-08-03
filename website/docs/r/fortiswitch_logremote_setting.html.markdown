---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_logremote_setting"
description: |-
  Settings for remote logging.
---

# fortiswitch_logremote_setting
Settings for remote logging.

## Example Usage

```hcl
resource "fortiswitch_logremote_setting" "name" {
        destination = "FAZ"
        status = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable remote log upload. Valid values: `enable`, `disable`.
* `destination` - Remote log type. Valid values: `FAZ`, `FAS`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogRemote Setting can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_logremote_setting.labelname LogRemoteSetting

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_logremote_setting.labelname LogRemoteSetting
$ unset "FORTISWITCH_IMPORT_TABLE"
```
