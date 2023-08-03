---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_console"
description: |-
  Configure console.
---

# fortiswitch_system_console
Configure console.

## Example Usage

```hcl
resource "fortiswitch_system_console" "sconsole" {
        hostname_display_length = "5"
}
```

## Argument Reference

The following arguments are supported:

* `hostname_display_length` - CLI prompt hostname display length.
* `output` - Console output mode. Valid values: `standard`, `more`.
* `baudrate` - Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
* `mode` - Console mode. Valid values: `batch`, `line`.
* `login` - Enable/disable console login. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Console can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_console.labelname SystemConsole

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_console.labelname SystemConsole
$ unset "FORTISWITCH_IMPORT_TABLE"
```
