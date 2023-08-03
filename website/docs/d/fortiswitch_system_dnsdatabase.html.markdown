---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_dnsdatabase"
description: |-
  Get information on an fortiswitch system dnsdatabase.
---

# Data Source: fortiswitch_system_dnsdatabase
Use this data source to get information on an fortiswitch system dnsdatabase

## Argument Reference

* `name` - (Required) Specify the name of the desired system dnsdatabase.

## Attribute Reference

The following attributes are exported:

* `status` - Dns zone status.
* `domain` - Domain name.
* `name` - Zone name.
* `authoritative` - Authoritative zone.
* `primary_name` - Domain name of the default DNS server for this zone.
* `ip_master` - IP address of master DNS server to import entries of this zone.
* `dns_entry` - Dns entry. The structure of `dns_entry` block is documented below.
* `contact` - Email address of the administrator for this zone
		you can specify only the username (e.g. admin)or full email address (e.g. admin.ca@test.com) 
		when using simple username, the domain of the email will be this zone.
* `ttl` - Default time-to-live value in units of seconds for the entries of this zone, range 0 to 2147483647.
* `forwarder` - Dns zone forwarder ip address list.
* `allow_transfer` - Dns zone transfer ip address list.
* `type` - Zone type ('master' to manage entries directly, 'slave' to import entries from outside).
* `source_ip` - Source IP for forwarding to DNS server.
* `view` - Zone view ('public' to server public clients, 'shadow' to serve internal clients).

The `dns_entry` block contains:

* `status` - Resource record status.
* `hostname` - Hostname.
* `ip` - IPv4 address.
* `ipv6` - IPv6 address.
* `canonical_name` - Canonical name.
* `preference` - 0 for the highest preference, range 0 to 65535.
* `ttl` - Time-to-live value in units of seconds for this entry, range 0 to 2147483647.
* `type` - Resource record type.
* `id` - Dns entry id.

