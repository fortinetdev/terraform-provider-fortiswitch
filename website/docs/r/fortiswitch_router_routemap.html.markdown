---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_routemap"
description: |-
  Route map configuration.
---

# fortiswitch_router_routemap
Route map configuration.

## Example Usage

```hcl
resource "fortiswitch_router_routemap" "name" {
        comments = "test"
        name = "default_name_4"
        protocol = "ospf"
        rule {
            action = "permit"
            id =  "8"
            match_community_exact = "enable"
            match_flags = "12"
            match_metric = "17"
            match_origin = "none"
            match_route_type = "1"
            match_tag = "20"
            set_aggregator_as = "21"
            set_atomic_aggregate = "enable"
            set_community_additive = "enable"
            set_flags = "34"
            set_local_preference = "38"
            set_metric = "39"
            set_metric_type = "1"
            set_origin = "none"
            set_originator_id = "12"
            set_tag = "43"
            set_weight = "44"
        }
}
```

## Argument Reference

The following arguments are supported:

* `rule` - Rule. The structure of `rule` block is documented below.
* `protocol` - Route-map type. Valid values: `ospf`, `ospf6`, `rip`, `bgp`, `isis`, `zebra`, `ripng`, `isis6`.
* `name` - Name.
* `comments` - Description/comments.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `rule` block supports:

* `match_community_exact` - Do exact matching of communities. Valid values: `enable`, `disable`.
* `set_community` - Set BGP community attribute. The structure of `set_community` block is documented below.
* `set_metric` - Set the metric value.
* `set_atomic_aggregate` - BGP atomic aggregate attribute. Valid values: `enable`, `disable`.
* `match_ip6_address` - Match ipv6 address permitted by access-list6 or prefix-list6.
* `match_origin` - Match BGP origin code. Valid values: `none`, `egp`, `igp`, `incomplete`.
* `match_metric` - Match metric for redistribute routes.
* `id` - Rule id.
* `match_flags` - Match-flags.
* `match_ip_address` - Match ip address permitted by access-list or prefix-list.
* `set_origin` - Set BGP origin code. Valid values: `none`, `egp`, `igp`, `incomplete`.
* `set_extcommunity_soo` - Set Site-of-Origin extended community. The structure of `set_extcommunity_soo` block is documented below.
* `set_flags` - Set-flags.
* `set_ip6_nexthop` - Set ipv6 global address of next hop.
* `match_as_path` - Match BGP AS path list.
* `set_extcommunity_rt` - Set Route Target extended community. The structure of `set_extcommunity_rt` block is documented below.
* `set_ip_nexthop` - Set ip address of next hop.
* `set_tag` - Set the tag value.
* `set_aggregator_as` - Set BGP aggregator AS.
* `set_weight` - Set BGP weight for routing table.
* `match_route_type` - Match route type. Valid values: `1`, `2`.
* `set_community_delete` - Delete communities matching community list.
* `match_community` - Match BGP community list.
* `set_metric_type` - Set the metric type. Valid values: `1`, `2`.
* `set_community_additive` - Add set-community to existing community. Valid values: `enable`, `disable`.
* `match_ip_nexthop` - Match next hop ip address passed by access-list or prefix-list.
* `set_originator_id` - Set BGP originator ID attribute.
* `match_tag` - Match tag.
* `set_aggregator_ip` - Set BGP aggregator IP.
* `match_interface` - Match interface configuration.
* `set_aspath` - Prepend BGP AS path attribute. The structure of `set_aspath` block is documented below.
* `action` - Action. Valid values: `permit`, `deny`.
* `set_ip6_nexthop_local` - Set ipv6 local address of next hop.
* `set_local_preference` - Set BGP local preference path attribute.

The `set_community` block supports:

* `community` - AA|AA:NN|internet|local-AS|no-advertise|no-export.

The `set_extcommunity_soo` block supports:

* `community` - AA:NN.

The `set_extcommunity_rt` block supports:

* `community` - AA:NN.

The `set_aspath` block supports:

* `as` - AS number, value range from 0 to 4294967295
NOTE: Use quotes for repeating numbers, e.g.: "1 1 2"
.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router RouteMap can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_routemap.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_routemap.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
