---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_ipv6neighborcache"
description: |-
  Configure IPv6 neighbor cache table.
---

# fortiswitch_system_ipv6neighborcache
Configure IPv6 neighbor cache table.

## Example Usage

```hcl
resource "fortiswitch_system_ipv6neighborcache" "name" {
        fswid =  "3"
        interface = "internal"
        ipv6 = "ffff:ffff:ffff:ffff:ffff:0000::"
        mac = "ff:ff:ff:ff:ff:ff"
}
```

## Argument Reference

The following arguments are supported:

* `interface` - Select the associated interface name from available options.
* `mac` - MAC address (format: xx:xx:xx:xx:xx:xx).
* `fswid` - Unique integer ID of the entry.
* `ipv6` - IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

System Ipv6NeighborCache can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_ipv6neighborcache.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_ipv6neighborcache.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
