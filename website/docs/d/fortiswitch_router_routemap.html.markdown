---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_routemap"
description: |-
  Get information on an fortiswitch router routemap.
---

# Data Source: fortiswitch_router_routemap
Use this data source to get information on an fortiswitch router routemap

## Argument Reference

* `name` - (Required) Specify the name of the desired router routemap.

## Attribute Reference

The following attributes are exported:

* `rule` - Rule. The structure of `rule` block is documented below.
* `protocol` - Route-map type.
* `name` - Name.
* `comments` - Description/comments.

The `rule` block contains:

* `match_community_exact` - Do exact matching of communities.
* `set_community` - Set BGP community attribute. The structure of `set_community` block is documented below.
* `set_metric` - Set the metric value.
* `set_atomic_aggregate` - BGP atomic aggregate attribute.
* `match_ip6_address` - Match ipv6 address permitted by access-list6 or prefix-list6.
* `match_origin` - Match BGP origin code.
* `match_metric` - Match metric for redistribute routes.
* `id` - Rule id.
* `match_flags` - Match-flags.
* `match_ip_address` - Match ip address permitted by access-list or prefix-list.
* `set_origin` - Set BGP origin code.
* `set_extcommunity_soo` - Set Site-of-Origin extended community. The structure of `set_extcommunity_soo` block is documented below.
* `set_flags` - Set-flags.
* `set_ip6_nexthop` - Set ipv6 global address of next hop.
* `match_as_path` - Match BGP AS path list.
* `set_extcommunity_rt` - Set Route Target extended community. The structure of `set_extcommunity_rt` block is documented below.
* `set_ip_nexthop` - Set ip address of next hop.
* `set_tag` - Set the tag value.
* `set_aggregator_as` - Set BGP aggregator AS.
* `set_weight` - Set BGP weight for routing table.
* `match_route_type` - Match route type.
* `set_community_delete` - Delete communities matching community list.
* `match_community` - Match BGP community list.
* `set_metric_type` - Set the metric type.
* `set_community_additive` - Add set-community to existing community.
* `match_ip_nexthop` - Match next hop ip address passed by access-list or prefix-list.
* `set_originator_id` - Set BGP originator ID attribute.
* `match_tag` - Match tag.
* `set_aggregator_ip` - Set BGP aggregator IP.
* `match_interface` - Match interface configuration.
* `set_aspath` - Prepend BGP AS path attribute. The structure of `set_aspath` block is documented below.
* `action` - Action.
* `set_ip6_nexthop_local` - Set ipv6 local address of next hop.
* `set_local_preference` - Set BGP local preference path attribute.

The `set_community` block contains:

* `community` - AA|AA:NN|internet|local-AS|no-advertise|no-export.

The `set_extcommunity_soo` block contains:

* `community` - AA:NN.

The `set_extcommunity_rt` block contains:

* `community` - AA:NN.

The `set_aspath` block contains:

* `as` - AS number, value range from 0 to 4294967295
NOTE: Use quotes for repeating numbers, e.g.: "1 1 2"
.

