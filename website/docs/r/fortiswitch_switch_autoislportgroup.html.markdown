---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_autoislportgroup"
description: |-
  Auto ISL port group.
---

# fortiswitch_switch_autoislportgroup
Auto ISL port group.

## Example Usage

```hcl
resource "fortiswitch_switch_autoislportgroup" "name" {
        members {
            member_name = "port1"
        }
        name = "test"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Auto ISL port group name.
* `members` - Auto ISL port group. The structure of `members` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `members` block supports:

* `member_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Switch AutoIslPortGroup can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_autoislportgroup.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_autoislportgroup.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
