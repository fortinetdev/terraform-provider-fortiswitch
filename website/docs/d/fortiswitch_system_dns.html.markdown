---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_dns"
description: |-
  Get information on fortiswitch system dns.
---

# Data Source: fortiswitch_system_dns
Use this data source to get information on fortiswitch system dns

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `domain` - Local domain name.
* `dns_cache_ttl` - TTL in DNS cache (60 - 86400).
* `primary` - IP address for primary DNS server.
* `dns_cache_limit` - Maximum number of entries in DNS cache.
* `cache_notfound_responses` - Enable/disable caching of NOTFOUND responses from DNS server.
* `ip6_secondary` - IPv6 address for secondary DNS server.
* `ip6_primary` - IPv6 address for primary DNS server.
* `source_ip` - Source IP for DNS queries.
* `secondary` - IP address for secondary DNS server.

