---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_sflow"
description: |-
  Configure sFlow.
---

# fortiswitch_system_sflow
Configure sFlow.

## Example Usage

```hcl
resource "fortiswitch_system_sflow" "name" {
        collectors {
            ip = "1.1.1.1"
            name = "default_name_5"
            port = "6"
        }
        dummy = "7"
}
```

## Argument Reference

The following arguments are supported:

* `dummy` - SFlow dummy.
* `collectors` - Collectors. The structure of `collectors` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `collectors` block supports:

* `ip` - Collector IP.
* `name` - Collector name.
* `port` - SFlow collector port (0 - 65535).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Sflow can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_sflow.labelname SystemSflow

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_sflow.labelname SystemSflow
$ unset "FORTISWITCH_IMPORT_TABLE"
```
