---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_dns"
description: |-
  Configure DNS.
---

# fortiswitch_system_dns
Configure DNS.

## Example Usage

```hcl
resource "fortiswitch_system_dns" "sdns" {
        dns_cache_limit = "4"
        dns_cache_ttl = "60"
        domain = "baidu"
        primary = "2.3.4.5"
}
```

## Argument Reference

The following arguments are supported:

* `domain` - Local domain name.
* `dns_cache_ttl` - TTL in DNS cache (60 - 86400).
* `primary` - IP address for primary DNS server.
* `dns_cache_limit` - Maximum number of entries in DNS cache.
* `cache_notfound_responses` - Enable/disable caching of NOTFOUND responses from DNS server. Valid values: `disable`, `enable`.
* `ip6_secondary` - IPv6 address for secondary DNS server.
* `ip6_primary` - IPv6 address for primary DNS server.
* `source_ip` - Source IP for DNS queries.
* `secondary` - IP address for secondary DNS server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Dns can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_dns.labelname SystemDns

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_dns.labelname SystemDns
$ unset "FORTISWITCH_IMPORT_TABLE"
```
