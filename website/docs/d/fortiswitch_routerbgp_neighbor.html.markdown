---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_routerbgp_neighbor"
description: |-
  Get information on an fortiswitch routerbgp neighbor.
---

# Data Source: fortiswitch_routerbgp_neighbor
Use this data source to get information on an fortiswitch routerbgp neighbor

## Argument Reference

* `ip` - (Required) Specify the ip of the desired routerbgp neighbor.

## Attribute Reference

The following attributes are exported:

* `send_community` - IPv4 Send community attribute to neighbor.
* `activate` - Enable/disable address family IPv4 for this neighbor.
* `filter_list_out6` - BGP filter for IPv6 outbound routes.
* `weight` - Neighbor weight.
* `attribute_unchanged` - IPv4 List of attributes that should be unchanged.
* `ip` - IP/IPv6 address of neighbor.
* `filter_list_in6` - BGP filter for IPv6 inbound routes.
* `ebgp_multihop_ttl` - EBGP multihop TTL for this peer.
* `default_originate_routemap` - Route map to specify criteria to originate IPv4 default.
* `default_originate_routemap6` - Route map to specify criteria to originate IPv6 default.
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client.
* `route_map_out6` - IPv6 Outbound route map filter.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates.
* `shutdown` - Enable/disable shutdown this neighbor.
* `route_map_in6` - IPv6 Inbound route map filter.
* `route_map_in_evpn` - EVPN Inbound route map filter.
* `unsuppress_map6` - IPv6 Route map to selectively unsuppress suppressed routes.
* `unsuppress_map` - IPv4 Route map to selectively unsuppress suppressed routes.
* `as_override6` - Enable/disable replace peer AS with own AS for IPv6.
* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged.
* `ebgp_enforce_multihop` - Enable/disable allow multi-hop next-hops from EBGP neighbors.
* `advertisement_interval` - Minimum interval (seconds) between sending updates.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `capability_default_originate6` - Enable/disable advertise default IPv6 route to this neighbor.
* `bfd` - Enable/disable BFD for this neighbor.
* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor.
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path.
* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client.
* `activate_evpn` - Enable/disable address family evpn for this neighbor.
* `dont_capability_negotiate` - Don't negotiate capabilities with this neighbor
* `connect_timer` - Interval (seconds) for connect timer.
* `passive` - Enable/disable sending of open messages to this neighbor.
* `attribute_unchanged_evpn` - EVPN List of attributes that should be unchanged.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `route_server_client` - Enable/disable IPv4 AS route server client.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1-100 percent).
* `filter_list_out` - BGP aspath filter for IPv4 outbound routes.
* `enforce_first_as` - Enable/disable  - Enable to enforce first AS for all(IPV4/IPV6) EBGP routes.
* `soft_reconfiguration_evpn` - Enable/disable allow EVPN inbound soft reconfiguration.
* `keep_alive_timer` - Keepalive timer interval (seconds).
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when threshold is exceeded.
* `description` - Description.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4.
* `route_reflector_client_evpn` - Enable/disable EVPN AS route reflector client.
* `bfd_session_mode` - Single or multihop BFD session to this neighbor.
* `distribute_list_out` - Filter for IPv4 updates to this neighbor.
* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor.
* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration.
* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when threshold is exceeded.
* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor.
* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `route_map_out_evpn` - EVPN outbound route map filter.
* `update_source` - Interface to use as source IP/IPv6 address of TCP connections.
* `interface` - Interface.
* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates.
* `password` - Password used in MD5 authentication.
* `holdtime_timer` - Interval (seconds) before peer considered dead.
* `route_map_in` - IPv4 Inbound route map filter.
* `activate6` - Enable/disable address family IPv6 for this neighbor.
* `filter_list_in` - BGP aspath filter for IPv4 inbound routes.
* `maximum_prefix` - Maximum number of IPv4 prefixes to accept from this peer.
* `route_server_client6` - Enable/disable IPv6 AS route server client.
* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor.
* `allowas_in_enable_evpn` - Enable/disable EVPN Enable to allow my AS in AS path.
* `ebgp_ttl_security_hops` - Specify the maximum number of hops to the EBGP peer.
* `distribute_list_in6` - Filter for IPv6 updates from this neighbor.
* `override_capability` - Enable/disable override result of capability negotiation.
* `distribute_list_out6` - Filter for IPv6 updates to this neighbor.
* `send_community6` - IPv6 Send community attribute to neighbor.
* `route_map_out` - IPv4 outbound route map filter.
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor.
* `capability_default_originate` - Enable/disable advertise default IPv4 route to this neighbor.
* `strict_capability_match` - Enable/disable strict capability matching.
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor.
* `distribute_list_in` - Filter for IPv4 updates from this neighbor.
* `remote_as` - AS number of neighbor.
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1-100 percent)
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor.
* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration.

