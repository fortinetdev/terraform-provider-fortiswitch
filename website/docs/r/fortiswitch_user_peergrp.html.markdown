---
subcategory: "FortiSwitch User"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_user_peergrp"
description: |-
  config peer's user group
---

# fortiswitch_user_peergrp
config peer's user group

## Example Usage

```hcl
resource "fortiswitch_user_peergrp" "upg" {
        member {
            name = "peer"
        }
        name = "group_peer"
}
```

## Argument Reference

The following arguments are supported:

* `member` - set peer group members The structure of `member` block is documented below.
* `name` - peer group name
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `member` block supports:

* `name` - peer group member name


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Peergrp can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_user_peergrp.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_user_peergrp.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
