---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_automationstitch"
description: |-
  Automation stitches.
---

# fortiswitch_system_automationstitch
Automation stitches. Applies to FortiSwitch Version `>= 7.2.0`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable this stitch. Valid values: `enable`, `disable`.
* `action` - Action names. The structure of `action` block is documented below.
* `trigger` - Trigger name.
* `name` - Name.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `action` block supports:

* `name` - Action name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationStitch can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_automationstitch.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_automationstitch.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
