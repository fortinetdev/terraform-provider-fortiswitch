---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_bgp"
description: |-
  Get information on fortiswitch router bgp.
---

# Data Source: fortiswitch_router_bgp
Use this data source to get information on fortiswitch router bgp

## Example Usage

```hcl
data "fortiswitch_router_bgp" sample1 {
}

output output1 {
  value = data.fortiswitch_router_bgp.sample1
}
```

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `confederation_peers` - Confederation peers. The structure of `confederation_peers` block is documented below.
* `distance_internal` - Distance for routes internal to the AS.
* `aggregate_address6` - BGP IPv6 aggregate address table. The structure of `aggregate_address6` block is documented below.
* `dampening_suppress` - Threshold to suppress routes.
* `bestpath_cmp_confed_aspath` - Enable/disable compare federation AS path length.
* `keepalive_timer` - Frequency to send keepalive requests.
* `as` - Router AS number.
* `dampening_reachability_half_life` - Reachability half-life time for penalty (minutes).
* `admin_distance` - Administrative distance modifications. The structure of `admin_distance` block is documented below.
* `ebgp_requires_policy` - Enable/disable require in and out policy for eBGP peers (RFC8212).
* `distance_external` - Distance for routes external to the AS.
* `dampening` - Enable/disable route-flap dampening.
* `network` - BGP network table. The structure of `network` block is documented below.
* `router_id` - Router ID.
* `aggregate_address` - BGP aggregate address table. The structure of `aggregate_address` block is documented below.
* `distance_local` - Distance for routes local to the AS.
* `bestpath_med_confed` - Enable/disable compare MED among confederation paths.
* `default_local_preference` - Default local preference.
* `bestpath_cmp_routerid` - Enable/disable compare router ID for identical EBGP paths.
* `bestpath_aspath_multipath_relax` - Allow load sharing across routes that have different AS paths (but same length).
* `graceful_stalepath_time` - Time to hold stale paths of restarting neighbour(sec).
* `enforce_first_as` - Enable/disable enforce first AS for EBGP routes.
* `cluster_id` - Route reflector cluster ID.
* `scan_time` - Background scanner interval (seconds).
* `bestpath_med_missing_as_worst` - Enable/disable treat missing MED as least preferred.
* `holdtime_timer` - Number of seconds to mark peer as dead.
* `network6` - BGP IPv6 network table. The structure of `network6` block is documented below.
* `dampening_reuse` - Threshold to unsuppress routes.
* `always_compare_med` - Enable/disable always compare MED.
* `maximum_paths_ebgp` - Maximum paths for ebgp ecmp.
* `bestpath_as_path_ignore` - Enable/disable ignore AS path.
* `redistribute` - BGP IPv4 redistribute table. The structure of `redistribute` block is documented below.
* `client_to_client_reflection` - Enable/disable client-to-client route reflection.
* `dampening_max_suppress_time` - Maximum minutes a route can be suppressed.
* `deterministic_med` - Enable/disable enforce deterministic comparison of MED.
* `fast_external_failover` - Enable/disable reset peer BGP session if link goes down.
* `maximum_paths_ibgp` - Maximum paths for ibgp ecmp.
* `redistribute6` - BGP IPv6 redistribute table. The structure of `redistribute6` block is documented below.
* `log_neighbour_changes` - Enable logging of BGP neighbour's changes
* `neighbor` - BGP neighbor table. The structure of `neighbor` block is documented below.
* `neighbor_group` - BGP neighbor group table. The structure of `neighbor_group` block is documented below.
* `route_reflector_allow_outbound_policy` - Enable/disable route reflector to apply a route-map to reflected routes.
* `admin_distance6` - Administrative distance modifications. The structure of `admin_distance6` block is documented below.
* `confederation_identifier` - Confederation identifier.

The `confederation_peers` block contains:

* `peer` - Peer ID.

The `aggregate_address6` block contains:

* `summary_only` - Enable/disable filter more specific routes from updates.
* `id` - ID.
* `prefix6` - Aggregate IPv6 prefix.

The `admin_distance` block contains:

* `distance` - Administrative distance to apply (1 - 255).
* `neighbour_prefix` - Neighbor address prefix.
* `route_list` - Access list of routes to apply new distance to.
* `id` - ID.

The `network` block contains:

* `backdoor` - Enable/disable route as backdoor.
* `prefix` - Network prefix.
* `route_map` - Route map to modify generated route.
* `id` - ID.

The `aggregate_address` block contains:

* `as_set` - Enable/disable generate AS set path information.
* `prefix` - Aggregate prefix.
* `id` - ID.
* `summary_only` - Enable/disable filter more specific routes from updates.

The `network6` block contains:

* `route_map` - Route map to modify generated route.
* `id` - ID.
* `prefix6` - Network IPv6 prefix.

The `redistribute` block contains:

* `status` - Status
* `route_map` - Route map name.
* `name` - Redistribute protocol name.

The `redistribute6` block contains:

* `status` - Status
* `route_map` - Route map name.
* `name` - Distribute list entry name.

The `neighbor` block contains:

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
* `capability_extended_nexthop` - Enable/disable extended nexthop capability.
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

The `neighbor_group` block contains:

* `name` - Neighbor group name.
* `advertisement_interval` - Minimum interval (seconds) between sending updates.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path.
* `allowas_in_enable_evpn` - Enable/disable EVPN Enable to allow my AS in AS path.
* `allowas_in_enable6` - Enable/disable IPv6 - Enable to allow my AS in AS path.
* `enforce_first_as` - Enable/disable  - Enable to enforce first AS for all(IPV4/IPV6) EBGP routes.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `attribute_unchanged` - IPv4 List of attributes that should be unchanged.
* `attribute_unchanged_evpn` - EVPN List of attributes that should be unchanged.
* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged.
* `activate` - Enable/disable address family IPv4 for this neighbor.
* `activate6` - Enable/disable address family IPv6 for this neighbor.
* `activate_evpn` - Enable/disable address family evpn for this neighbor.
* `bfd` - Enable/disable BFD for this neighbor.
* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor.
* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor.
* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor.
* `capability_default_originate` - Enable/disable advertise default IPv4 route to this neighbor.
* `capability_default_originate6` - Enable/disable advertise default IPv6 route to this neighbor.
* `capability_extended_nexthop` - Enable/disable extended nexthop capability.
* `dont_capability_negotiate` - Don't negotiate capabilities with this neighbor
* `ebgp_enforce_multihop` - Enable/disable allow multi-hop next-hops from EBGP neighbors.
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor.
* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor.
* `override_capability` - Enable/disable override result of capability negotiation.
* `passive` - Enable/disable sending of open messages to this neighbor.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates.
* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates.
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client.
* `route_reflector_client_evpn` - Enable/disable EVPN AS route reflector client.
* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client.
* `route_server_client` - Enable/disable IPv4 AS route server client.
* `route_server_client6` - Enable/disable IPv6 AS route server client.
* `shutdown` - Enable/disable shutdown this neighbor.
* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration.
* `soft_reconfiguration_evpn` - Enable/disable allow EVPN inbound soft reconfiguration.
* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4.
* `as_override6` - Enable/disable replace peer AS with own AS for IPv6.
* `strict_capability_match` - Enable/disable strict capability matching.
* `default_originate_routemap` - Route map to specify criteria to originate IPv4 default.
* `default_originate_routemap6` - Route map to specify criteria to originate IPv6 default.
* `description` - Description.
* `distribute_list_in` - Filter for IPv4 updates from this neighbor.
* `distribute_list_in6` - Filter for IPv6 updates from this neighbor.
* `distribute_list_out` - Filter for IPv4 updates to this neighbor.
* `distribute_list_out6` - Filter for IPv6 updates to this neighbor.
* `ebgp_multihop_ttl` - EBGP multihop TTL for this peer.
* `ebgp_ttl_security_hops` - Specify the maximum number of hops to the EBGP peer.
* `filter_list_in` - BGP aspath filter for IPv4 inbound routes.
* `filter_list_in6` - BGP filter for IPv6 inbound routes.
* `filter_list_out` - BGP aspath filter for IPv4 outbound routes.
* `filter_list_out6` - BGP filter for IPv6 outbound routes.
* `interface` - Interface(s). The structure of `interface` block is documented below.
* `maximum_prefix` - Maximum number of IPv4 prefixes to accept from this peer.
* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1-100 percent).
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1-100 percent)
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when threshold is exceeded.
* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when threshold is exceeded.
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor.
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor.
* `remote_as` - AS number of neighbor.
* `route_map_in` - IPv4 Inbound route map filter.
* `route_map_in_evpn` - EVPN Inbound route map filter.
* `route_map_in6` - IPv6 Inbound route map filter.
* `route_map_out` - IPv4 outbound route map filter.
* `route_map_out_evpn` - EVPN outbound route map filter.
* `route_map_out6` - IPv6 Outbound route map filter.
* `send_community` - IPv4 Send community attribute to neighbor.
* `send_community6` - IPv6 Send community attribute to neighbor.
* `keep_alive_timer` - Keepalive timer interval (seconds).
* `holdtime_timer` - Interval (seconds) before peer considered dead.
* `connect_timer` - Interval (seconds) for connect timer.
* `unsuppress_map` - IPv4 Route map to selectively unsuppress suppressed routes.
* `unsuppress_map6` - IPv6 Route map to selectively unsuppress suppressed routes.
* `update_source` - Interface to use as source IP/IPv6 address of TCP connections.
* `weight` - Neighbor weight.
* `password` - Password used in MD5 authentication.

The `interface` block contains:

* `interface_name` - RVI interface name(s).

The `admin_distance6` block contains:

* `neighbour_prefix6` - Neighbor IPV6 prefix.
* `distance` - Administrative distance to apply (1 - 255).
* `id` - ID.
* `route6_list` - Access list of routes to apply new distance to.

