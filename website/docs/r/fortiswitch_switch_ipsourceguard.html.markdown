---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_ipsourceguard"
description: |-
  IP-source-guard table.
---

# fortiswitch_switch_ipsourceguard
IP-source-guard table. Applies to FortiSwitch Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `interface_name` - Ingress interface that Source Guard is bound
* `description` - Description of the source-guard name.
* `binding_entry` - IP-MAC configuration. The structure of `binding_entry` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `binding_entry` block supports:

* `entry_name` - Configure binding pair.
* `ip` - Source ip for this rule.
* `mac` - MAC address value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{interface_name}}.

## Import

Switch IpSourceGuard can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_ipsourceguard.labelname {{interface_name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_ipsourceguard.labelname {{interface_name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
