---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_dnsdatabase"
description: |-
  Dns-database.
---

# fortiswitch_system_dnsdatabase
Dns-database.

## Argument Reference

The following arguments are supported:

* `status` - Dns zone status. Valid values: `enable`, `disable`.
* `domain` - Domain name.
* `name` - Zone name.
* `authoritative` - Authoritative zone. Valid values: `enable`, `disable`.
* `primary_name` - Domain name of the default DNS server for this zone.
* `ip_master` - IP address of master DNS server to import entries of this zone.
* `dns_entry` - Dns entry. The structure of `dns_entry` block is documented below.
* `contact` - Email address of the administrator for this zone
		you can specify only the username (e.g. admin)or full email address (e.g. admin.ca@test.com) 
		when using simple username, the domain of the email will be this zone.
* `ttl` - Default time-to-live value in units of seconds for the entries of this zone, range 0 to 2147483647.
* `forwarder` - Dns zone forwarder ip address list.
* `allow_transfer` - Dns zone transfer ip address list.
* `type` - Zone type ('master' to manage entries directly, 'slave' to import entries from outside). Valid values: `master`, `slave`.
* `source_ip` - Source IP for forwarding to DNS server.
* `view` - Zone view ('public' to server public clients, 'shadow' to serve internal clients). Valid values: `shadow`, `public`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `dns_entry` block supports:

* `status` - Resource record status. Valid values: `enable`, `disable`.
* `hostname` - Hostname.
* `ip` - IPv4 address.
* `ipv6` - IPv6 address.
* `canonical_name` - Canonical name.
* `preference` - 0 for the highest preference, range 0 to 65535.
* `ttl` - Time-to-live value in units of seconds for this entry, range 0 to 2147483647.
* `type` - Resource record type. Valid values: `A`, `NS`, `CNAME`, `MX`, `AAAA`, `PTR`, `PTR_V6`.
* `id` - Dns entry id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System DnsDatabase can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_dnsdatabase.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_dnsdatabase.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
