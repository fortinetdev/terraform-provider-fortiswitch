---
subcategory: "FortiSwitch User"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_user_group"
description: |-
  User group configuration.
---

# fortiswitch_user_group
User group configuration.

## Example Usage

```hcl
resource "fortiswitch_user_group" "ugroup" {
        authtimeout = "3"
        group_type = "firewall"
        member {
            name = "test89"
        }
        name = "group2"
}
```

## Argument Reference

The following arguments are supported:

* `group_type` - Type of user group. Valid values: `firewall`.
* `name` - Group name.
* `authtimeout` - Authentication timeout (0 - 480).
* `member` - Group members. The structure of `member` block is documented below.
* `http_digest_realm` - Realm attribute for MD5-digest authentication.
* `match` - Set group matches. The structure of `match` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `member` block supports:

* `name` - Group member name.

The `match` block supports:

* `server_name` - Name of remote authentication server.
* `group_name` - Name of matching group on remote authentication server.
* `id` - ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Group can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_user_group.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_user_group.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
