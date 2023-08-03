---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_arptable"
description: |-
  Configure arp table.
---

# fortiswitch_system_arptable
Configure arp table.

## Example Usage

```hcl
resource "fortiswitch_system_arptable" "fsat" {
        fswid =  "2"
        interface = "internal"
        ip = "1.2.3.4"
        mac = "00:00:5e:00:53:af"
}
```

## Argument Reference

The following arguments are supported:

* `interface` - Interface name.
* `ip` - IPv4 address.
* `mac` - MAC address.
* `fswid` - Unique ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

System ArpTable can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_arptable.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_arptable.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
