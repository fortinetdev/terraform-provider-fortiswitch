---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_fortiguard"
description: |-
  Get information on fortiswitch system fortiguard.
---

# Data Source: fortiswitch_system_fortiguard
Use this data source to get information on fortiswitch system fortiguard

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `webfilter_license` - License type.
* `webfilter_cache` - Enable/disable the cache.
* `avquery_license` - License type.
* `antispam_cache_mpercent` - The maximum percent of memory the cache is allowed to use (1-15%).
* `webfilter_timeout` - Query time out (1-30 seconds).
* `antispam_cache` - Enable/disable the cache.
* `webfilter_expiration` - When license will expire.
* `load_balance_servers` - Number of servers to alternate between as first Fortiguard option.
* `antispam_force_off` - Forcibly disable the service.
* `avquery_force_off` - Forcibly disable the service.
* `webfilter_force_off` - Forcibly disable the service.
* `avquery_timeout` - Query time out (1-30 seconds).
* `client_override_status` - Enable or disable the client override IP.
* `antispam_score_threshold` - Antispam score threshold [50,100].
* `avquery_cache` - Enable/disable the cache.
* `avquery_cache_mpercent` - The maximum percent of memory the cache is allowed to use (1-15%).
* `avquery_cache_ttl` - The time-to-live for cache entries in seconds (300-86400).
* `avquery_expiration` - When license will expire.
* `hostname` - Hostname or IP of the FortiGuard server.
* `antispam_license` - License type.
* `port` - Port used to communicate with the FortiGuard servers.
* `antispam_cache_ttl` - The time-to-live for cache entries in seconds (300-86400).
* `analysis_service` - Enable or disable the analysis service.
* `client_override_ip` - Client override IP address.
* `srv_ovrd` - Enable or disable the server override list.
* `webfilter_cache_ttl` - The time-to-live for cache entries in seconds (300-86400).
* `service_account_id` - Service account id.
* `antispam_expiration` - When license will expire.
* `antispam_timeout` - Query time out (1-30 seconds).
* `srv_ovrd_list` - Configure the server override list. The structure of `srv_ovrd_list` block is documented below.

The `srv_ovrd_list` block contains:

* `ip` - Override server IP address.
* `ip6` - Override server IP6 address.
* `addr_type` - Type of address.

