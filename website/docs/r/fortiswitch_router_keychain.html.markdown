---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_keychain"
description: |-
  Key-chain configuration.
---

# fortiswitch_router_keychain
Key-chain configuration.

## Example Usage

```hcl
resource "fortiswitch_router_keychain" "name" {
        key {
            id =  "5"
            key_string = "test"
        }
        name = "default_name_8"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Key-chain name.
* `key` - Key. The structure of `key` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `key` block supports:

* `accept_lifetime` - Lifetime of received authentication key.
* `key_string` - Password for the key.
* `id` - Key id <0-2147483647>.
* `send_lifetime` - Lifetime of sent authentication key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router KeyChain can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_keychain.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_keychain.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
