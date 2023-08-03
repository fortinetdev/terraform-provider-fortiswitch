---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_macaddresstable"
description: |-
  Mac address table.
---

# fortiswitch_system_macaddresstable
Mac address table.

## Example Usage

```hcl
resource "fortiswitch_system_macaddresstable" "name" {
        interface = "internal"
        mac = "00:00:5e:00:53:af"
}
```

## Argument Reference

The following arguments are supported:

* `interface` - Interface name.
* `mac` - MAC address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mac}}.

## Import

System MacAddressTable can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_macaddresstable.labelname {{mac}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_macaddresstable.labelname {{mac}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
