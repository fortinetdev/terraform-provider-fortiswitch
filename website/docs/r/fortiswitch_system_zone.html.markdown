---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_zone"
description: |-
  Zone configuration.
---

# fortiswitch_system_zone
Zone configuration.

## Example Usage

```hcl
resource "fortiswitch_system_zone" "name" {
        interface {
            interface_name = "port5"
        }
        intrazone = "allow"
        name = "default_name_6"
}
```

## Argument Reference

The following arguments are supported:

* `interface` - Interfaces belong to the zone. The structure of `interface` block is documented below.
* `intrazone` - Intra-zone traffic. Valid values: `allow`, `deny`.
* `name` - Zone name.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `interface` block supports:

* `interface_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Zone can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_zone.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_zone.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
