---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_aspathlist"
description: |-
  AS path list configuration.
---

# fortiswitch_router_aspathlist
AS path list configuration.

## Example Usage

```hcl
resource "fortiswitch_router_aspathlist" "name" {
        name = "default_name_3"
        rule {
            action = "deny"
            id =  "6"
            regexp = "test"
        }
}
```

## Argument Reference

The following arguments are supported:

* `name` - AS path list name.
* `rule` - AS path list rule. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `rule` block supports:

* `action` - AS path list action. Valid values: `deny`, `permit`.
* `regexp` - Regular-expression to match the BGP AS paths.
* `id` - Id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router AspathList can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_aspathlist.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_aspathlist.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
