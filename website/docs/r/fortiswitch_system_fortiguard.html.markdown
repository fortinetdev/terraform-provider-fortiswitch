---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_fortiguard"
description: |-
  Configure FortiGuard services.
---

# fortiswitch_system_fortiguard
Configure FortiGuard services.

## Example Usage

```hcl
resource "fortiswitch_system_fortiguard" "name" {
        analysis_service = "enable"
        antispam_cache = "enable"
        antispam_cache_mpercent = "5"
        antispam_cache_ttl = "6"
        antispam_expiration = "7"
        antispam_force_off = "enable"
        antispam_license = "9"
        antispam_score_threshold = "10"
        antispam_timeout = "11"
        avquery_cache = "enable"
        avquery_cache_mpercent = "13"
        avquery_cache_ttl = "14"
        avquery_expiration = "15"
        avquery_force_off = "enable"
        avquery_license = "17"
}
```

## Argument Reference

The following arguments are supported:

* `webfilter_license` - License type.
* `webfilter_cache` - Enable/disable the cache. Valid values: `enable`, `disable`.
* `avquery_license` - License type.
* `antispam_cache_mpercent` - The maximum percent of memory the cache is allowed to use (1-15%).
* `webfilter_timeout` - Query time out (1-30 seconds).
* `antispam_cache` - Enable/disable the cache. Valid values: `enable`, `disable`.
* `webfilter_expiration` - When license will expire.
* `load_balance_servers` - Number of servers to alternate between as first Fortiguard option.
* `antispam_force_off` - Forcibly disable the service. Valid values: `enable`, `disable`.
* `avquery_force_off` - Forcibly disable the service. Valid values: `enable`, `disable`.
* `webfilter_force_off` - Forcibly disable the service. Valid values: `enable`, `disable`.
* `avquery_timeout` - Query time out (1-30 seconds).
* `client_override_status` - Enable or disable the client override IP. Valid values: `enable`, `disable`.
* `antispam_score_threshold` - Antispam score threshold [50,100].
* `avquery_cache` - Enable/disable the cache. Valid values: `enable`, `disable`.
* `avquery_cache_mpercent` - The maximum percent of memory the cache is allowed to use (1-15%).
* `avquery_cache_ttl` - The time-to-live for cache entries in seconds (300-86400).
* `avquery_expiration` - When license will expire.
* `hostname` - Hostname or IP of the FortiGuard server.
* `antispam_license` - License type.
* `port` - Port used to communicate with the FortiGuard servers. Valid values: `53`, `8888`.
* `antispam_cache_ttl` - The time-to-live for cache entries in seconds (300-86400).
* `analysis_service` - Enable or disable the analysis service. Valid values: `enable`, `disable`.
* `client_override_ip` - Client override IP address.
* `srv_ovrd` - Enable or disable the server override list. Valid values: `enable`, `disable`.
* `webfilter_cache_ttl` - The time-to-live for cache entries in seconds (300-86400).
* `service_account_id` - Service account id.
* `antispam_expiration` - When license will expire.
* `antispam_timeout` - Query time out (1-30 seconds).
* `srv_ovrd_list` - Configure the server override list. The structure of `srv_ovrd_list` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `srv_ovrd_list` block supports:

* `ip` - Override server IP address.
* `ip6` - Override server IP6 address.
* `addr_type` - Type of address. Valid values: `ipv4`, `ipv6`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fortiguard can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_fortiguard.labelname SystemFortiguard

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_fortiguard.labelname SystemFortiguard
$ unset "FORTISWITCH_IMPORT_TABLE"
```
