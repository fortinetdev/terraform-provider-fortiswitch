---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_security"
description: |-
  Configure Fortiswitch secure mode.
---

# fortiswitch_system_security
Configure Fortiswitch secure mode. Applies to FortiSwitch Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `mode` - Enable Fortiswitch secure mode. Valid values: `none`, `lince`, `fips-cc`, `fedramp`.
* `reseed_interval` - Entropy reseeding interval (0- 1440 min, default=1440 min).
* `self_test_interval` - System self test interval (0- 1440 min, default=0 min).
* `app_security` - Configure Fortiswitch per application OpenSSL security settings. The structure of `app_security` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `app_security` block supports:

* `process` - Process name.
* `security_level` - Application specific openssl security level setting. (0- 5, default=2).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Security can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_security.labelname SystemSecurity

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_security.labelname SystemSecurity
$ unset "FORTISWITCH_IMPORT_TABLE"
```
