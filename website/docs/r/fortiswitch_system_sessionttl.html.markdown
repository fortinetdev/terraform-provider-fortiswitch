---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_sessionttl"
description: |-
  Session ttl configuration.
---

# fortiswitch_system_sessionttl
Session ttl configuration.

## Example Usage

```hcl
resource "fortiswitch_system_sessionttl" "name" {
        default = "56"
        port {
            end_port = "5"
            id =  "6"
            protocol = "7"
            start_port = "8"
            timeout = "23"
        }
}
```

## Argument Reference

The following arguments are supported:

* `default` - Default timeout.
* `port` - Configure session ttl port. The structure of `port` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `port` block supports:

* `end_port` - End port number.
* `timeout` - Configure timeout.
* `protocol` - Protocol (0-255).
* `id` - Table entry ID.
* `start_port` - Start port number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SessionTtl can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_sessionttl.labelname SystemSessionTtl

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_sessionttl.labelname SystemSessionTtl
$ unset "FORTISWITCH_IMPORT_TABLE"
```
