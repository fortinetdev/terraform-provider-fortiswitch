---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_communitylist"
description: |-
  Community list configuration.
---

# fortiswitch_router_communitylist
Community list configuration.

## Example Usage

```hcl
resource "fortiswitch_router_communitylist" "name" {
        name = "default_name_3"
        rule {
            action = "deny"
            id =  "6"
            match = "exact"
            regexp = "test"
        }
        type = "standard"
}
```

## Argument Reference

The following arguments are supported:

* `type` - Community list type. Valid values: `standard`, `expanded`.
* `name` - Community list name.
* `rule` - Community list rule. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `rule` block supports:

* `action` - Community list action. Valid values: `deny`, `permit`.
* `regexp` - An ordered list as a regular expression.
* `id` - Id.
* `match` - Community specification.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router CommunityList can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_communitylist.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_communitylist.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
