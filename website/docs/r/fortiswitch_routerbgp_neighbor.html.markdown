---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_routerbgp_neighbor"
description: |-
  BGP neighbor table.
---

# fortiswitch_routerbgp_neighbor
BGP neighbor table.

~> The provider supports the definition of Neighbor in Router Bgp `fortiswitch_router_bgp`, and also allows the definition of separate Neighbor resources `fortiswitch_routerbgp_neighbor`, but do not use a `fortiswitch_router_bgp` with in-line Neighbor in conjunction with any `fortiswitch_routerbgp_neighbor` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

* `send_community` - IPv4 Send community attribute to neighbor. Valid values: `standard`, `extended`, `both`, `disable`.
* `activate` - Enable/disable address family IPv4 for this neighbor. Valid values: `enable`, `disable`.
* `filter_list_out6` - BGP filter for IPv6 outbound routes.
* `weight` - Neighbor weight.
* `attribute_unchanged` - IPv4 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.
* `ip` - IP/IPv6 address of neighbor.
* `filter_list_in6` - BGP filter for IPv6 inbound routes.
* `ebgp_multihop_ttl` - EBGP multihop TTL for this peer.
* `default_originate_routemap` - Route map to specify criteria to originate IPv4 default.
* `default_originate_routemap6` - Route map to specify criteria to originate IPv6 default.
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client. Valid values: `enable`, `disable`.
* `route_map_out6` - IPv6 Outbound route map filter.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates. Valid values: `enable`, `disable`.
* `shutdown` - Enable/disable shutdown this neighbor. Valid values: `enable`, `disable`.
* `route_map_in6` - IPv6 Inbound route map filter.
* `route_map_in_evpn` - EVPN Inbound route map filter.
* `unsuppress_map6` - IPv6 Route map to selectively unsuppress suppressed routes.
* `unsuppress_map` - IPv4 Route map to selectively unsuppress suppressed routes.
* `as_override6` - Enable/disable replace peer AS with own AS for IPv6. Valid values: `enable`, `disable`.
* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.
* `ebgp_enforce_multihop` - Enable/disable allow multi-hop next-hops from EBGP neighbors. Valid values: `enable`, `disable`.
* `advertisement_interval` - Minimum interval (seconds) between sending updates.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `capability_default_originate6` - Enable/disable advertise default IPv6 route to this neighbor. Valid values: `enable`, `disable`.
* `capability_extended_nexthop` - Enable/disable extended nexthop capability. Valid values: `enable`, `disable`.
* `bfd` - Enable/disable BFD for this neighbor. Valid values: `enable`, `disable`.
* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor. Valid values: `none`, `receive`, `send`, `both`.
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor. Valid values: `enable`, `disable`.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client. Valid values: `enable`, `disable`.
* `activate_evpn` - Enable/disable address family evpn for this neighbor. Valid values: `enable`, `disable`.
* `dont_capability_negotiate` - Don't negotiate capabilities with this neighbor Valid values: `enable`, `disable`.
* `connect_timer` - Interval (seconds) for connect timer.
* `passive` - Enable/disable sending of open messages to this neighbor. Valid values: `enable`, `disable`.
* `attribute_unchanged_evpn` - EVPN List of attributes that should be unchanged. Valid values: `as-path`, `med`.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `route_server_client` - Enable/disable IPv4 AS route server client. Valid values: `enable`, `disable`.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1-100 percent).
* `filter_list_out` - BGP aspath filter for IPv4 outbound routes.
* `enforce_first_as` - Enable/disable  - Enable to enforce first AS for all(IPV4/IPV6) EBGP routes. Valid values: `enable`, `disable`.
* `soft_reconfiguration_evpn` - Enable/disable allow EVPN inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `keep_alive_timer` - Keepalive timer interval (seconds).
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when threshold is exceeded. Valid values: `enable`, `disable`.
* `description` - Description.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4. Valid values: `enable`, `disable`.
* `route_reflector_client_evpn` - Enable/disable EVPN AS route reflector client. Valid values: `enable`, `disable`.
* `bfd_session_mode` - Single or multihop BFD session to this neighbor. Valid values: `automatic`, `multihop`, `singlehop`.
* `distribute_list_out` - Filter for IPv4 updates to this neighbor.
* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor. Valid values: `none`, `receive`, `send`, `both`.
* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when threshold is exceeded. Valid values: `enable`, `disable`.
* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor. Valid values: `enable`, `disable`.
* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `route_map_out_evpn` - EVPN outbound route map filter.
* `update_source` - Interface to use as source IP/IPv6 address of TCP connections.
* `interface` - Interface.
* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates. Valid values: `enable`, `disable`.
* `password` - Password used in MD5 authentication.
* `holdtime_timer` - Interval (seconds) before peer considered dead.
* `route_map_in` - IPv4 Inbound route map filter.
* `activate6` - Enable/disable address family IPv6 for this neighbor. Valid values: `enable`, `disable`.
* `filter_list_in` - BGP aspath filter for IPv4 inbound routes.
* `maximum_prefix` - Maximum number of IPv4 prefixes to accept from this peer.
* `route_server_client6` - Enable/disable IPv6 AS route server client. Valid values: `enable`, `disable`.
* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor. Valid values: `enable`, `disable`.
* `allowas_in_enable_evpn` - Enable/disable EVPN Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
* `ebgp_ttl_security_hops` - Specify the maximum number of hops to the EBGP peer.
* `distribute_list_in6` - Filter for IPv6 updates from this neighbor.
* `override_capability` - Enable/disable override result of capability negotiation. Valid values: `enable`, `disable`.
* `distribute_list_out6` - Filter for IPv6 updates to this neighbor.
* `send_community6` - IPv6 Send community attribute to neighbor. Valid values: `standard`, `extended`, `both`, `disable`.
* `route_map_out` - IPv4 outbound route map filter.
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor.
* `capability_default_originate` - Enable/disable advertise default IPv4 route to this neighbor. Valid values: `enable`, `disable`.
* `strict_capability_match` - Enable/disable strict capability matching. Valid values: `enable`, `disable`.
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor.
* `distribute_list_in` - Filter for IPv4 updates from this neighbor.
* `remote_as` - AS number of neighbor.
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1-100 percent)
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor.
* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{ip}}.

## Import

Routerbgp Neighbor can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_routerbgp_neighbor.labelname {{ip}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_routerbgp_neighbor.labelname {{ip}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
