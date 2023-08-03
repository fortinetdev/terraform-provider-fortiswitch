---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_bgp"
description: |-
  BGP configuration.
---

# fortiswitch_router_bgp
BGP configuration.

~> The provider supports the definition of Neighbor in Router Bgp `fortiswitch_router_bgp`, and also allows the definition of separate Neighbor resources `fortiswitch_routerbgp_neighbor`, but do not use a `fortiswitch_router_bgp` with in-line Neighbor in conjunction with any `fortiswitch_routerbgp_neighbor` resources, otherwise conflicts and overwrite will occur.

~> The provider supports the definition of Network in Router Bgp `fortiswitch_router_bgp`, and also allows the definition of separate Network resources `fortiswitch_routerbgp_network`, but do not use a `fortiswitch_router_bgp` with in-line Network in conjunction with any `fortiswitch_routerbgp_network` resources, otherwise conflicts and overwrite will occur.

~> The provider supports the definition of Network6 in Router Bgp `fortiswitch_router_bgp`, and also allows the definition of separate Network6 resources `fortiswitch_routerbgp_network6`, but do not use a `fortiswitch_router_bgp` with in-line Network6 in conjunction with any `fortiswitch_routerbgp_network6` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

* `confederation_peers` - Confederation peers. The structure of `confederation_peers` block is documented below.
* `distance_internal` - Distance for routes internal to the AS.
* `aggregate_address6` - BGP IPv6 aggregate address table. The structure of `aggregate_address6` block is documented below.
* `dampening_suppress` - Threshold to suppress routes.
* `bestpath_cmp_confed_aspath` - Enable/disable compare federation AS path length. Valid values: `enable`, `disable`.
* `keepalive_timer` - Frequency to send keepalive requests.
* `as` - Router AS number.
* `dampening_reachability_half_life` - Reachability half-life time for penalty (minutes).
* `admin_distance` - Administrative distance modifications. The structure of `admin_distance` block is documented below.
* `ebgp_requires_policy` - Enable/disable require in and out policy for eBGP peers (RFC8212). Valid values: `enable`, `disable`.
* `distance_external` - Distance for routes external to the AS.
* `dampening` - Enable/disable route-flap dampening. Valid values: `enable`, `disable`.
* `network` - BGP network table. The structure of `network` block is documented below.
* `router_id` - Router ID.
* `aggregate_address` - BGP aggregate address table. The structure of `aggregate_address` block is documented below.
* `distance_local` - Distance for routes local to the AS.
* `bestpath_med_confed` - Enable/disable compare MED among confederation paths. Valid values: `enable`, `disable`.
* `default_local_preference` - Default local preference.
* `bestpath_cmp_routerid` - Enable/disable compare router ID for identical EBGP paths. Valid values: `enable`, `disable`.
* `bestpath_aspath_multipath_relax` - Allow load sharing across routes that have different AS paths (but same length). Valid values: `disable`, `enable`.
* `graceful_stalepath_time` - Time to hold stale paths of restarting neighbour(sec).
* `enforce_first_as` - Enable/disable enforce first AS for EBGP routes. Valid values: `enable`, `disable`.
* `cluster_id` - Route reflector cluster ID.
* `scan_time` - Background scanner interval (seconds).
* `bestpath_med_missing_as_worst` - Enable/disable treat missing MED as least preferred. Valid values: `enable`, `disable`.
* `holdtime_timer` - Number of seconds to mark peer as dead.
* `network6` - BGP IPv6 network table. The structure of `network6` block is documented below.
* `dampening_reuse` - Threshold to unsuppress routes.
* `always_compare_med` - Enable/disable always compare MED. Valid values: `enable`, `disable`.
* `maximum_paths_ebgp` - Maximum paths for ebgp ecmp.
* `bestpath_as_path_ignore` - Enable/disable ignore AS path. Valid values: `enable`, `disable`.
* `redistribute` - BGP IPv4 redistribute table. The structure of `redistribute` block is documented below.
* `client_to_client_reflection` - Enable/disable client-to-client route reflection. Valid values: `enable`, `disable`.
* `dampening_max_suppress_time` - Maximum minutes a route can be suppressed.
* `deterministic_med` - Enable/disable enforce deterministic comparison of MED. Valid values: `enable`, `disable`.
* `fast_external_failover` - Enable/disable reset peer BGP session if link goes down. Valid values: `enable`, `disable`.
* `maximum_paths_ibgp` - Maximum paths for ibgp ecmp.
* `redistribute6` - BGP IPv6 redistribute table. The structure of `redistribute6` block is documented below.
* `log_neighbour_changes` - Enable logging of BGP neighbour's changes Valid values: `enable`, `disable`.
* `neighbor` - BGP neighbor table. The structure of `neighbor` block is documented below.
* `route_reflector_allow_outbound_policy` - Enable/disable route reflector to apply a route-map to reflected routes. Valid values: `enable`, `disable`.
* `admin_distance6` - Administrative distance modifications. The structure of `admin_distance6` block is documented below.
* `confederation_identifier` - Confederation identifier.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `confederation_peers` block supports:

* `peer` - Peer ID.

The `aggregate_address6` block supports:

* `summary_only` - Enable/disable filter more specific routes from updates. Valid values: `enable`, `disable`.
* `id` - ID.
* `prefix6` - Aggregate IPv6 prefix.

The `admin_distance` block supports:

* `distance` - Administrative distance to apply (1 - 255).
* `neighbour_prefix` - Neighbor address prefix.
* `route_list` - Access list of routes to apply new distance to.
* `id` - ID.

The `network` block supports:

* `backdoor` - Enable/disable route as backdoor. Valid values: `enable`, `disable`.
* `prefix` - Network prefix.
* `route_map` - Route map to modify generated route.
* `id` - ID.

The `aggregate_address` block supports:

* `as_set` - Enable/disable generate AS set path information. Valid values: `enable`, `disable`.
* `prefix` - Aggregate prefix.
* `id` - ID.
* `summary_only` - Enable/disable filter more specific routes from updates. Valid values: `enable`, `disable`.

The `network6` block supports:

* `route_map` - Route map to modify generated route.
* `id` - ID.
* `prefix6` - Network IPv6 prefix.

The `redistribute` block supports:

* `status` - Status Valid values: `enable`, `disable`.
* `route_map` - Route map name.
* `name` - Redistribute protocol name.

The `redistribute6` block supports:

* `status` - Status Valid values: `enable`, `disable`.
* `route_map` - Route map name.
* `name` - Distribute list entry name.

The `neighbor` block supports:

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
* `unsuppress_map6` - IPv6 Route map to selectively unsuppress suppressed routes.
* `unsuppress_map` - IPv4 Route map to selectively unsuppress suppressed routes.
* `as_override6` - Enable/disable replace peer AS with own AS for IPv6. Valid values: `enable`, `disable`.
* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.
* `ebgp_enforce_multihop` - Enable/disable allow multi-hop next-hops from EBGP neighbors. Valid values: `enable`, `disable`.
* `advertisement_interval` - Minimum interval (seconds) between sending updates.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `capability_default_originate6` - Enable/disable advertise default IPv6 route to this neighbor. Valid values: `enable`, `disable`.
* `bfd` - Enable/disable BFD for this neighbor. Valid values: `enable`, `disable`.
* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor. Valid values: `none`, `receive`, `send`, `both`.
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor. Valid values: `enable`, `disable`.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client. Valid values: `enable`, `disable`.
* `dont_capability_negotiate` - Don't negotiate capabilities with this neighbor Valid values: `enable`, `disable`.
* `connect_timer` - Interval (seconds) for connect timer.
* `passive` - Enable/disable sending of open messages to this neighbor. Valid values: `enable`, `disable`.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `route_server_client` - Enable/disable IPv4 AS route server client. Valid values: `enable`, `disable`.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1-100 percent).
* `filter_list_out` - BGP aspath filter for IPv4 outbound routes.
* `enforce_first_as` - Enable/disable  - Enable to enforce first AS for all(IPV4/IPV6) EBGP routes. Valid values: `enable`, `disable`.
* `keep_alive_timer` - Keepalive timer interval (seconds).
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when threshold is exceeded. Valid values: `enable`, `disable`.
* `description` - Description.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4. Valid values: `enable`, `disable`.
* `bfd_session_mode` - Single or multihop BFD session to this neighbor. Valid values: `automatic`, `multihop`, `singlehop`.
* `distribute_list_out` - Filter for IPv4 updates to this neighbor.
* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor. Valid values: `none`, `receive`, `send`, `both`.
* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when threshold is exceeded. Valid values: `enable`, `disable`.
* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor. Valid values: `enable`, `disable`.
* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
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

The `admin_distance6` block supports:

* `neighbour_prefix6` - Neighbor IPV6 prefix.
* `distance` - Administrative distance to apply (1 - 255).
* `id` - ID.
* `route6_list` - Access list of routes to apply new distance to.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Bgp can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_bgp.labelname RouterBgp

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_bgp.labelname RouterBgp
$ unset "FORTISWITCH_IMPORT_TABLE"
```
