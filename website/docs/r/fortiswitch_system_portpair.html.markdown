---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_portpair"
description: |-
  Port-pair.
---

# fortiswitch_system_portpair
Port-pair.

## Example Usage

```hcl
resource "fortiswitch_system_portpair" "name" {
        member {
            interface_name = "mgmt"
        }
        name = "test"
}
```

## Argument Reference

The following arguments are supported:

* `member` - Interfaces belong to the port pair. The structure of `member` block is documented below.
* `name` - Port pair name.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `member` block supports:

* `interface_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System PortPair can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_portpair.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_portpair.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
