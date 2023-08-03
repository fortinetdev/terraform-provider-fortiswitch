---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_setting"
description: |-
  Set rib settings.
---

# fortiswitch_router_setting
Set rib settings.

## Example Usage

```hcl
resource "fortiswitch_router_setting" "name" {
       hostname = "myhostname"
}
```

## Argument Reference

The following arguments are supported:

* `filter_list` - Route filter list configuration. The structure of `filter_list` block is documented below.
* `hostname` - Set hostname for this virtual domain router.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `filter_list` block supports:

* `route_map` - Route map name.
* `protocol` - Protocol type. Valid values: `static`, `static6`, `rip`, `ripng`, `ospf`, `ospf6`, `isis`, `isis6`, `bgp`, `bgp6`, `any`, `any6`.
* `id` - Filter list entry ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Setting can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_setting.labelname RouterSetting

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_setting.labelname RouterSetting
$ unset "FORTISWITCH_IMPORT_TABLE"
```
