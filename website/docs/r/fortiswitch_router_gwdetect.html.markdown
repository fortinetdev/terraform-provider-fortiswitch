---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_gwdetect"
description: |-
  Gwdetect.
---

# fortiswitch_router_gwdetect
Gwdetect.

## Example Usage

```hcl
resource "fortiswitch_router_gwdetect" "name" {
        failtime = "3"
        ha_priority = "4"
        interface = "port6"
        interval = "6"
        protocol = "ping"
        server {
            address = "2.3.4.5"
        }
        source_ip = "84.230.14.43"
        status = "11"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Server status.
* `protocol` - Protocols used to detect the server. Valid values: `ping`, `tcp-echo`, `udp-echo`.
* `source_ip` - Source IP used to ping the server.
* `interval` - Detection interval.
* `failtime` - Fail-times for ping server lost .
* `interface` - Interface name.
* `ha_priority` - HA election priority (1-50).
* `server` - Server address(es). The structure of `server` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `server` block supports:

* `address` - Server address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{interface}}.

## Import

Router Gwdetect can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_gwdetect.labelname {{interface}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_gwdetect.labelname {{interface}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
