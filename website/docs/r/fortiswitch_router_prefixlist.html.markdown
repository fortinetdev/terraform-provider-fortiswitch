---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_prefixlist"
description: |-
  Prefix list configuration.
---

# fortiswitch_router_prefixlist
Prefix list configuration.

## Example Usage

```hcl
resource "fortiswitch_router_prefixlist" "name" {
        comments = "test"
        name = "name_4"
        rule {
            action = "permit"
            flags = "7"
            ge = "8"
            id =  "9"
            le = "10"
        }
}
```

## Argument Reference

The following arguments are supported:

* `rule` - Rule. The structure of `rule` block is documented below.
* `name` - Name.
* `comments` - Description/comments.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `rule` block supports:

* `le` - Maximum prefix length to be matched.
* `prefix` - Prefix.
* `ge` - Minimum prefix length to be matched.
* `flags` - Flags.
* `action` - Action. Valid values: `permit`, `deny`.
* `id` - Rule id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router PrefixList can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_prefixlist.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_prefixlist.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
