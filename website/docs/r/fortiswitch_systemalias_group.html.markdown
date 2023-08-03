---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemalias_group"
description: |-
  Groups of alias commands.
---

# fortiswitch_systemalias_group
Groups of alias commands.

## Argument Reference

The following arguments are supported:

* `commands` - Alias command list. The structure of `commands` block is documented below.
* `name` - Alias group name.
* `description` - Group description.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `commands` block supports:

* `command_name` - Alias command name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemAlias Group can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemalias_group.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemalias_group.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
