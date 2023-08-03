// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: BGP configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceRouterBgp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterBgpRead,
		Schema: map[string]*schema.Schema{

			"confederation_peers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"peer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"distance_internal": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"aggregate_address6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"summary_only": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dampening_suppress": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bestpath_cmp_confed_aspath": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"keepalive_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"as": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dampening_reachability_half_life": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_distance": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"distance": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"neighbour_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"ebgp_requires_policy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"distance_external": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dampening": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backdoor": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"aggregate_address": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"as_set": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"summary_only": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"distance_local": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bestpath_med_confed": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_local_preference": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bestpath_cmp_routerid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bestpath_aspath_multipath_relax": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"graceful_stalepath_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"enforce_first_as": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cluster_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"scan_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bestpath_med_missing_as_worst": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"holdtime_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"network6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"route_map": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dampening_reuse": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"always_compare_med": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"maximum_paths_ebgp": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bestpath_as_path_ignore": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"client_to_client_reflection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dampening_max_suppress_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"deterministic_med": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fast_external_failover": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"maximum_paths_ibgp": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"redistribute6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"log_neighbour_changes": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"send_community": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"activate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"filter_list_out6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"attribute_unchanged": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"filter_list_in6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ebgp_multihop_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"default_originate_routemap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"default_originate_routemap6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_reflector_client": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map_out6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"remove_private_as": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"shutdown": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map_in6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"unsuppress_map6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"unsuppress_map": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"as_override6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_unchanged6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ebgp_enforce_multihop": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"advertisement_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix_list_in6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_default_originate6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_orf": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"next_hop_self": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allowas_in_enable": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_reflector_client6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dont_capability_negotiate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"connect_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"passive": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allowas_in": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"maximum_prefix6": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"route_server_client": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"maximum_prefix_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"filter_list_out": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enforce_first_as": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"keep_alive_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"maximum_prefix_warning_only": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"as_override": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bfd_session_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distribute_list_out": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_orf6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"soft_reconfiguration6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"maximum_prefix_warning_only6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"next_hop_self6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allowas_in_enable6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allowas_in6": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"update_source": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"remove_private_as6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"holdtime_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"route_map_in": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"activate6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"filter_list_in": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"maximum_prefix": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"route_server_client6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_dynamic": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ebgp_ttl_security_hops": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"distribute_list_in6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"override_capability": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distribute_list_out6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"send_community6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map_out": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_list_out6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_default_originate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"strict_capability_match": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_list_in": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distribute_list_in": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_as": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"maximum_prefix_threshold6": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix_list_out": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"soft_reconfiguration": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"route_reflector_allow_outbound_policy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_distance6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"neighbour_prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distance": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"route6_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"confederation_identifier": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceRouterBgpRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "RouterBgp"

	o, err := c.ReadRouterBgp(mkey)
	if err != nil {
		return fmt.Errorf("Error describing RouterBgp: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterBgp(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterBgp from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterBgpConfederationPeers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {
			tmp["peer"] = dataSourceFlattenRouterBgpConfederationPeersPeer(i["peer"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpConfederationPeersPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDistanceInternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAggregateAddress6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if _, ok := i["summary-only"]; ok {
			tmp["summary_only"] = dataSourceFlattenRouterBgpAggregateAddress6SummaryOnly(i["summary-only"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterBgpAggregateAddress6Id(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			tmp["prefix6"] = dataSourceFlattenRouterBgpAggregateAddress6Prefix6(i["prefix6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpAggregateAddress6SummaryOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAggregateAddress6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAggregateAddress6Prefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDampeningSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpBestpathCmpConfedAspath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpKeepaliveTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDampeningReachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAdminDistance(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {
			tmp["distance"] = dataSourceFlattenRouterBgpAdminDistanceDistance(i["distance"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbour_prefix"
		if _, ok := i["neighbour-prefix"]; ok {
			tmp["neighbour_prefix"] = dataSourceFlattenRouterBgpAdminDistanceNeighbourPrefix(i["neighbour-prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_list"
		if _, ok := i["route-list"]; ok {
			tmp["route_list"] = dataSourceFlattenRouterBgpAdminDistanceRouteList(i["route-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterBgpAdminDistanceId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpAdminDistanceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAdminDistanceNeighbourPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterBgpAdminDistanceRouteList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAdminDistanceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpEbgpRequiresPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDistanceExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDampening(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if _, ok := i["backdoor"]; ok {
			tmp["backdoor"] = dataSourceFlattenRouterBgpNetworkBackdoor(i["backdoor"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterBgpNetworkPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			tmp["route_map"] = dataSourceFlattenRouterBgpNetworkRouteMap(i["route-map"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterBgpNetworkId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpNetworkBackdoor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterBgpNetworkRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAggregateAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if _, ok := i["as-set"]; ok {
			tmp["as_set"] = dataSourceFlattenRouterBgpAggregateAddressAsSet(i["as-set"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterBgpAggregateAddressPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterBgpAggregateAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if _, ok := i["summary-only"]; ok {
			tmp["summary_only"] = dataSourceFlattenRouterBgpAggregateAddressSummaryOnly(i["summary-only"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpAggregateAddressAsSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAggregateAddressPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterBgpAggregateAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAggregateAddressSummaryOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDistanceLocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpBestpathMedConfed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDefaultLocalPreference(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpBestpathCmpRouterid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpBestpathAspathMultipathRelax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpGracefulStalepathTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpEnforceFirstAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpClusterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpScanTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpBestpathMedMissingAsWorst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNetwork6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			tmp["route_map"] = dataSourceFlattenRouterBgpNetwork6RouteMap(i["route-map"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterBgpNetwork6Id(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			tmp["prefix6"] = dataSourceFlattenRouterBgpNetwork6Prefix6(i["prefix6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpNetwork6RouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNetwork6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNetwork6Prefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDampeningReuse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAlwaysCompareMed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpMaximumPathsEbgp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpBestpathAsPathIgnore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterBgpRedistributeStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			tmp["route_map"] = dataSourceFlattenRouterBgpRedistributeRouteMap(i["route-map"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterBgpRedistributeName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistributeRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpClientToClientReflection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDampeningMaxSuppressTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDeterministicMed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpFastExternalFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpMaximumPathsIbgp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistribute6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterBgpRedistribute6Status(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			tmp["route_map"] = dataSourceFlattenRouterBgpRedistribute6RouteMap(i["route-map"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterBgpRedistribute6Name(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpRedistribute6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistribute6RouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistribute6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpLogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if _, ok := i["send-community"]; ok {
			tmp["send_community"] = dataSourceFlattenRouterBgpNeighborSendCommunity(i["send-community"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if _, ok := i["activate"]; ok {
			tmp["activate"] = dataSourceFlattenRouterBgpNeighborActivate(i["activate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if _, ok := i["filter-list-out6"]; ok {
			tmp["filter_list_out6"] = dataSourceFlattenRouterBgpNeighborFilterListOut6(i["filter-list-out6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			tmp["weight"] = dataSourceFlattenRouterBgpNeighborWeight(i["weight"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if _, ok := i["attribute-unchanged"]; ok {
			tmp["attribute_unchanged"] = dataSourceFlattenRouterBgpNeighborAttributeUnchanged(i["attribute-unchanged"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenRouterBgpNeighborIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if _, ok := i["filter-list-in6"]; ok {
			tmp["filter_list_in6"] = dataSourceFlattenRouterBgpNeighborFilterListIn6(i["filter-list-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if _, ok := i["ebgp-multihop-ttl"]; ok {
			tmp["ebgp_multihop_ttl"] = dataSourceFlattenRouterBgpNeighborEbgpMultihopTtl(i["ebgp-multihop-ttl"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if _, ok := i["default-originate-routemap"]; ok {
			tmp["default_originate_routemap"] = dataSourceFlattenRouterBgpNeighborDefaultOriginateRoutemap(i["default-originate-routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if _, ok := i["default-originate-routemap6"]; ok {
			tmp["default_originate_routemap6"] = dataSourceFlattenRouterBgpNeighborDefaultOriginateRoutemap6(i["default-originate-routemap6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if _, ok := i["route-reflector-client"]; ok {
			tmp["route_reflector_client"] = dataSourceFlattenRouterBgpNeighborRouteReflectorClient(i["route-reflector-client"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if _, ok := i["route-map-out6"]; ok {
			tmp["route_map_out6"] = dataSourceFlattenRouterBgpNeighborRouteMapOut6(i["route-map-out6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if _, ok := i["remove-private-as"]; ok {
			tmp["remove_private_as"] = dataSourceFlattenRouterBgpNeighborRemovePrivateAs(i["remove-private-as"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if _, ok := i["shutdown"]; ok {
			tmp["shutdown"] = dataSourceFlattenRouterBgpNeighborShutdown(i["shutdown"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if _, ok := i["route-map-in6"]; ok {
			tmp["route_map_in6"] = dataSourceFlattenRouterBgpNeighborRouteMapIn6(i["route-map-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if _, ok := i["unsuppress-map6"]; ok {
			tmp["unsuppress_map6"] = dataSourceFlattenRouterBgpNeighborUnsuppressMap6(i["unsuppress-map6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if _, ok := i["unsuppress-map"]; ok {
			tmp["unsuppress_map"] = dataSourceFlattenRouterBgpNeighborUnsuppressMap(i["unsuppress-map"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if _, ok := i["as-override6"]; ok {
			tmp["as_override6"] = dataSourceFlattenRouterBgpNeighborAsOverride6(i["as-override6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if _, ok := i["attribute-unchanged6"]; ok {
			tmp["attribute_unchanged6"] = dataSourceFlattenRouterBgpNeighborAttributeUnchanged6(i["attribute-unchanged6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if _, ok := i["ebgp-enforce-multihop"]; ok {
			tmp["ebgp_enforce_multihop"] = dataSourceFlattenRouterBgpNeighborEbgpEnforceMultihop(i["ebgp-enforce-multihop"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if _, ok := i["advertisement-interval"]; ok {
			tmp["advertisement_interval"] = dataSourceFlattenRouterBgpNeighborAdvertisementInterval(i["advertisement-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if _, ok := i["prefix-list-in6"]; ok {
			tmp["prefix_list_in6"] = dataSourceFlattenRouterBgpNeighborPrefixListIn6(i["prefix-list-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if _, ok := i["capability-default-originate6"]; ok {
			tmp["capability_default_originate6"] = dataSourceFlattenRouterBgpNeighborCapabilityDefaultOriginate6(i["capability-default-originate6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			tmp["bfd"] = dataSourceFlattenRouterBgpNeighborBfd(i["bfd"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if _, ok := i["capability-orf"]; ok {
			tmp["capability_orf"] = dataSourceFlattenRouterBgpNeighborCapabilityOrf(i["capability-orf"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if _, ok := i["next-hop-self"]; ok {
			tmp["next_hop_self"] = dataSourceFlattenRouterBgpNeighborNextHopSelf(i["next-hop-self"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if _, ok := i["allowas-in-enable"]; ok {
			tmp["allowas_in_enable"] = dataSourceFlattenRouterBgpNeighborAllowasInEnable(i["allowas-in-enable"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if _, ok := i["route-reflector-client6"]; ok {
			tmp["route_reflector_client6"] = dataSourceFlattenRouterBgpNeighborRouteReflectorClient6(i["route-reflector-client6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if _, ok := i["dont-capability-negotiate"]; ok {
			tmp["dont_capability_negotiate"] = dataSourceFlattenRouterBgpNeighborDontCapabilityNegotiate(i["dont-capability-negotiate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if _, ok := i["connect-timer"]; ok {
			tmp["connect_timer"] = dataSourceFlattenRouterBgpNeighborConnectTimer(i["connect-timer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := i["passive"]; ok {
			tmp["passive"] = dataSourceFlattenRouterBgpNeighborPassive(i["passive"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if _, ok := i["allowas-in"]; ok {
			tmp["allowas_in"] = dataSourceFlattenRouterBgpNeighborAllowasIn(i["allowas-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if _, ok := i["maximum-prefix6"]; ok {
			tmp["maximum_prefix6"] = dataSourceFlattenRouterBgpNeighborMaximumPrefix6(i["maximum-prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if _, ok := i["route-server-client"]; ok {
			tmp["route_server_client"] = dataSourceFlattenRouterBgpNeighborRouteServerClient(i["route-server-client"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if _, ok := i["maximum-prefix-threshold"]; ok {
			tmp["maximum_prefix_threshold"] = dataSourceFlattenRouterBgpNeighborMaximumPrefixThreshold(i["maximum-prefix-threshold"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if _, ok := i["filter-list-out"]; ok {
			tmp["filter_list_out"] = dataSourceFlattenRouterBgpNeighborFilterListOut(i["filter-list-out"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "enforce_first_as"
		if _, ok := i["enforce-first-as"]; ok {
			tmp["enforce_first_as"] = dataSourceFlattenRouterBgpNeighborEnforceFirstAs(i["enforce-first-as"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if _, ok := i["keep-alive-timer"]; ok {
			tmp["keep_alive_timer"] = dataSourceFlattenRouterBgpNeighborKeepAliveTimer(i["keep-alive-timer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if _, ok := i["maximum-prefix-warning-only"]; ok {
			tmp["maximum_prefix_warning_only"] = dataSourceFlattenRouterBgpNeighborMaximumPrefixWarningOnly(i["maximum-prefix-warning-only"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			tmp["description"] = dataSourceFlattenRouterBgpNeighborDescription(i["description"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if _, ok := i["as-override"]; ok {
			tmp["as_override"] = dataSourceFlattenRouterBgpNeighborAsOverride(i["as-override"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_session_mode"
		if _, ok := i["bfd-session-mode"]; ok {
			tmp["bfd_session_mode"] = dataSourceFlattenRouterBgpNeighborBfdSessionMode(i["bfd-session-mode"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if _, ok := i["distribute-list-out"]; ok {
			tmp["distribute_list_out"] = dataSourceFlattenRouterBgpNeighborDistributeListOut(i["distribute-list-out"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if _, ok := i["capability-orf6"]; ok {
			tmp["capability_orf6"] = dataSourceFlattenRouterBgpNeighborCapabilityOrf6(i["capability-orf6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if _, ok := i["soft-reconfiguration6"]; ok {
			tmp["soft_reconfiguration6"] = dataSourceFlattenRouterBgpNeighborSoftReconfiguration6(i["soft-reconfiguration6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if _, ok := i["maximum-prefix-warning-only6"]; ok {
			tmp["maximum_prefix_warning_only6"] = dataSourceFlattenRouterBgpNeighborMaximumPrefixWarningOnly6(i["maximum-prefix-warning-only6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if _, ok := i["next-hop-self6"]; ok {
			tmp["next_hop_self6"] = dataSourceFlattenRouterBgpNeighborNextHopSelf6(i["next-hop-self6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if _, ok := i["allowas-in-enable6"]; ok {
			tmp["allowas_in_enable6"] = dataSourceFlattenRouterBgpNeighborAllowasInEnable6(i["allowas-in-enable6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if _, ok := i["allowas-in6"]; ok {
			tmp["allowas_in6"] = dataSourceFlattenRouterBgpNeighborAllowasIn6(i["allowas-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if _, ok := i["update-source"]; ok {
			tmp["update_source"] = dataSourceFlattenRouterBgpNeighborUpdateSource(i["update-source"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterBgpNeighborInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if _, ok := i["remove-private-as6"]; ok {
			tmp["remove_private_as6"] = dataSourceFlattenRouterBgpNeighborRemovePrivateAs6(i["remove-private-as6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			tmp["password"] = dataSourceFlattenRouterBgpNeighborPassword(i["password"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if _, ok := i["holdtime-timer"]; ok {
			tmp["holdtime_timer"] = dataSourceFlattenRouterBgpNeighborHoldtimeTimer(i["holdtime-timer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if _, ok := i["route-map-in"]; ok {
			tmp["route_map_in"] = dataSourceFlattenRouterBgpNeighborRouteMapIn(i["route-map-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if _, ok := i["activate6"]; ok {
			tmp["activate6"] = dataSourceFlattenRouterBgpNeighborActivate6(i["activate6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if _, ok := i["filter-list-in"]; ok {
			tmp["filter_list_in"] = dataSourceFlattenRouterBgpNeighborFilterListIn(i["filter-list-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if _, ok := i["maximum-prefix"]; ok {
			tmp["maximum_prefix"] = dataSourceFlattenRouterBgpNeighborMaximumPrefix(i["maximum-prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if _, ok := i["route-server-client6"]; ok {
			tmp["route_server_client6"] = dataSourceFlattenRouterBgpNeighborRouteServerClient6(i["route-server-client6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if _, ok := i["capability-dynamic"]; ok {
			tmp["capability_dynamic"] = dataSourceFlattenRouterBgpNeighborCapabilityDynamic(i["capability-dynamic"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_ttl_security_hops"
		if _, ok := i["ebgp-ttl-security-hops"]; ok {
			tmp["ebgp_ttl_security_hops"] = dataSourceFlattenRouterBgpNeighborEbgpTtlSecurityHops(i["ebgp-ttl-security-hops"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if _, ok := i["distribute-list-in6"]; ok {
			tmp["distribute_list_in6"] = dataSourceFlattenRouterBgpNeighborDistributeListIn6(i["distribute-list-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if _, ok := i["override-capability"]; ok {
			tmp["override_capability"] = dataSourceFlattenRouterBgpNeighborOverrideCapability(i["override-capability"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if _, ok := i["distribute-list-out6"]; ok {
			tmp["distribute_list_out6"] = dataSourceFlattenRouterBgpNeighborDistributeListOut6(i["distribute-list-out6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if _, ok := i["send-community6"]; ok {
			tmp["send_community6"] = dataSourceFlattenRouterBgpNeighborSendCommunity6(i["send-community6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if _, ok := i["route-map-out"]; ok {
			tmp["route_map_out"] = dataSourceFlattenRouterBgpNeighborRouteMapOut(i["route-map-out"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if _, ok := i["prefix-list-out6"]; ok {
			tmp["prefix_list_out6"] = dataSourceFlattenRouterBgpNeighborPrefixListOut6(i["prefix-list-out6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if _, ok := i["capability-default-originate"]; ok {
			tmp["capability_default_originate"] = dataSourceFlattenRouterBgpNeighborCapabilityDefaultOriginate(i["capability-default-originate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if _, ok := i["strict-capability-match"]; ok {
			tmp["strict_capability_match"] = dataSourceFlattenRouterBgpNeighborStrictCapabilityMatch(i["strict-capability-match"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if _, ok := i["prefix-list-in"]; ok {
			tmp["prefix_list_in"] = dataSourceFlattenRouterBgpNeighborPrefixListIn(i["prefix-list-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if _, ok := i["distribute-list-in"]; ok {
			tmp["distribute_list_in"] = dataSourceFlattenRouterBgpNeighborDistributeListIn(i["distribute-list-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if _, ok := i["remote-as"]; ok {
			tmp["remote_as"] = dataSourceFlattenRouterBgpNeighborRemoteAs(i["remote-as"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if _, ok := i["maximum-prefix-threshold6"]; ok {
			tmp["maximum_prefix_threshold6"] = dataSourceFlattenRouterBgpNeighborMaximumPrefixThreshold6(i["maximum-prefix-threshold6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if _, ok := i["prefix-list-out"]; ok {
			tmp["prefix_list_out"] = dataSourceFlattenRouterBgpNeighborPrefixListOut(i["prefix-list-out"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if _, ok := i["soft-reconfiguration"]; ok {
			tmp["soft_reconfiguration"] = dataSourceFlattenRouterBgpNeighborSoftReconfiguration(i["soft-reconfiguration"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpNeighborSendCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborActivate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborFilterListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAttributeUnchanged(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborFilterListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborEbgpMultihopTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDefaultOriginateRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDefaultOriginateRoutemap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteReflectorClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteMapOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRemovePrivateAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborShutdown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteMapIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborUnsuppressMap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborUnsuppressMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAsOverride6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAttributeUnchanged6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborEbgpEnforceMultihop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAdvertisementInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborPrefixListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborCapabilityDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborCapabilityOrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborNextHopSelf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAllowasInEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteReflectorClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDontCapabilityNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborConnectTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborPassive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAllowasIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborMaximumPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteServerClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborMaximumPrefixThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborFilterListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborEnforceFirstAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborKeepAliveTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborMaximumPrefixWarningOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAsOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborBfdSessionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDistributeListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborCapabilityOrf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborSoftReconfiguration6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborMaximumPrefixWarningOnly6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborNextHopSelf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAllowasInEnable6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAllowasIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborUpdateSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRemovePrivateAs6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteMapIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborActivate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborFilterListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborMaximumPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteServerClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborCapabilityDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborEbgpTtlSecurityHops(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDistributeListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborOverrideCapability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDistributeListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborSendCommunity6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteMapOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborPrefixListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborCapabilityDefaultOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborStrictCapabilityMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborPrefixListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDistributeListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRemoteAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborMaximumPrefixThreshold6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborPrefixListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborSoftReconfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRouteReflectorAllowOutboundPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAdminDistance6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbour_prefix6"
		if _, ok := i["neighbour-prefix6"]; ok {
			tmp["neighbour_prefix6"] = dataSourceFlattenRouterBgpAdminDistance6NeighbourPrefix6(i["neighbour-prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {
			tmp["distance"] = dataSourceFlattenRouterBgpAdminDistance6Distance(i["distance"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterBgpAdminDistance6Id(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route6_list"
		if _, ok := i["route6-list"]; ok {
			tmp["route6_list"] = dataSourceFlattenRouterBgpAdminDistance6Route6List(i["route6-list"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpAdminDistance6NeighbourPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAdminDistance6Distance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAdminDistance6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAdminDistance6Route6List(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpConfederationIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterBgp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("confederation_peers", dataSourceFlattenRouterBgpConfederationPeers(o["confederation-peers"], d, "confederation_peers")); err != nil {
		if !fortiAPIPatch(o["confederation-peers"]) {
			return fmt.Errorf("Error reading confederation_peers: %v", err)
		}
	}

	if err = d.Set("distance_internal", dataSourceFlattenRouterBgpDistanceInternal(o["distance-internal"], d, "distance_internal")); err != nil {
		if !fortiAPIPatch(o["distance-internal"]) {
			return fmt.Errorf("Error reading distance_internal: %v", err)
		}
	}

	if err = d.Set("aggregate_address6", dataSourceFlattenRouterBgpAggregateAddress6(o["aggregate-address6"], d, "aggregate_address6")); err != nil {
		if !fortiAPIPatch(o["aggregate-address6"]) {
			return fmt.Errorf("Error reading aggregate_address6: %v", err)
		}
	}

	if err = d.Set("dampening_suppress", dataSourceFlattenRouterBgpDampeningSuppress(o["dampening-suppress"], d, "dampening_suppress")); err != nil {
		if !fortiAPIPatch(o["dampening-suppress"]) {
			return fmt.Errorf("Error reading dampening_suppress: %v", err)
		}
	}

	if err = d.Set("bestpath_cmp_confed_aspath", dataSourceFlattenRouterBgpBestpathCmpConfedAspath(o["bestpath-cmp-confed-aspath"], d, "bestpath_cmp_confed_aspath")); err != nil {
		if !fortiAPIPatch(o["bestpath-cmp-confed-aspath"]) {
			return fmt.Errorf("Error reading bestpath_cmp_confed_aspath: %v", err)
		}
	}

	if err = d.Set("keepalive_timer", dataSourceFlattenRouterBgpKeepaliveTimer(o["keepalive-timer"], d, "keepalive_timer")); err != nil {
		if !fortiAPIPatch(o["keepalive-timer"]) {
			return fmt.Errorf("Error reading keepalive_timer: %v", err)
		}
	}

	if err = d.Set("as", dataSourceFlattenRouterBgpAs(o["as"], d, "as")); err != nil {
		if !fortiAPIPatch(o["as"]) {
			return fmt.Errorf("Error reading as: %v", err)
		}
	}

	if err = d.Set("dampening_reachability_half_life", dataSourceFlattenRouterBgpDampeningReachabilityHalfLife(o["dampening-reachability-half-life"], d, "dampening_reachability_half_life")); err != nil {
		if !fortiAPIPatch(o["dampening-reachability-half-life"]) {
			return fmt.Errorf("Error reading dampening_reachability_half_life: %v", err)
		}
	}

	if err = d.Set("admin_distance", dataSourceFlattenRouterBgpAdminDistance(o["admin-distance"], d, "admin_distance")); err != nil {
		if !fortiAPIPatch(o["admin-distance"]) {
			return fmt.Errorf("Error reading admin_distance: %v", err)
		}
	}

	if err = d.Set("ebgp_requires_policy", dataSourceFlattenRouterBgpEbgpRequiresPolicy(o["ebgp-requires-policy"], d, "ebgp_requires_policy")); err != nil {
		if !fortiAPIPatch(o["ebgp-requires-policy"]) {
			return fmt.Errorf("Error reading ebgp_requires_policy: %v", err)
		}
	}

	if err = d.Set("distance_external", dataSourceFlattenRouterBgpDistanceExternal(o["distance-external"], d, "distance_external")); err != nil {
		if !fortiAPIPatch(o["distance-external"]) {
			return fmt.Errorf("Error reading distance_external: %v", err)
		}
	}

	if err = d.Set("dampening", dataSourceFlattenRouterBgpDampening(o["dampening"], d, "dampening")); err != nil {
		if !fortiAPIPatch(o["dampening"]) {
			return fmt.Errorf("Error reading dampening: %v", err)
		}
	}

	if err = d.Set("network", dataSourceFlattenRouterBgpNetwork(o["network"], d, "network")); err != nil {
		if !fortiAPIPatch(o["network"]) {
			return fmt.Errorf("Error reading network: %v", err)
		}
	}

	if err = d.Set("router_id", dataSourceFlattenRouterBgpRouterId(o["router-id"], d, "router_id")); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("aggregate_address", dataSourceFlattenRouterBgpAggregateAddress(o["aggregate-address"], d, "aggregate_address")); err != nil {
		if !fortiAPIPatch(o["aggregate-address"]) {
			return fmt.Errorf("Error reading aggregate_address: %v", err)
		}
	}

	if err = d.Set("distance_local", dataSourceFlattenRouterBgpDistanceLocal(o["distance-local"], d, "distance_local")); err != nil {
		if !fortiAPIPatch(o["distance-local"]) {
			return fmt.Errorf("Error reading distance_local: %v", err)
		}
	}

	if err = d.Set("bestpath_med_confed", dataSourceFlattenRouterBgpBestpathMedConfed(o["bestpath-med-confed"], d, "bestpath_med_confed")); err != nil {
		if !fortiAPIPatch(o["bestpath-med-confed"]) {
			return fmt.Errorf("Error reading bestpath_med_confed: %v", err)
		}
	}

	if err = d.Set("default_local_preference", dataSourceFlattenRouterBgpDefaultLocalPreference(o["default-local-preference"], d, "default_local_preference")); err != nil {
		if !fortiAPIPatch(o["default-local-preference"]) {
			return fmt.Errorf("Error reading default_local_preference: %v", err)
		}
	}

	if err = d.Set("bestpath_cmp_routerid", dataSourceFlattenRouterBgpBestpathCmpRouterid(o["bestpath-cmp-routerid"], d, "bestpath_cmp_routerid")); err != nil {
		if !fortiAPIPatch(o["bestpath-cmp-routerid"]) {
			return fmt.Errorf("Error reading bestpath_cmp_routerid: %v", err)
		}
	}

	if err = d.Set("bestpath_aspath_multipath_relax", dataSourceFlattenRouterBgpBestpathAspathMultipathRelax(o["bestpath-aspath-multipath-relax"], d, "bestpath_aspath_multipath_relax")); err != nil {
		if !fortiAPIPatch(o["bestpath-aspath-multipath-relax"]) {
			return fmt.Errorf("Error reading bestpath_aspath_multipath_relax: %v", err)
		}
	}

	if err = d.Set("graceful_stalepath_time", dataSourceFlattenRouterBgpGracefulStalepathTime(o["graceful-stalepath-time"], d, "graceful_stalepath_time")); err != nil {
		if !fortiAPIPatch(o["graceful-stalepath-time"]) {
			return fmt.Errorf("Error reading graceful_stalepath_time: %v", err)
		}
	}

	if err = d.Set("enforce_first_as", dataSourceFlattenRouterBgpEnforceFirstAs(o["enforce-first-as"], d, "enforce_first_as")); err != nil {
		if !fortiAPIPatch(o["enforce-first-as"]) {
			return fmt.Errorf("Error reading enforce_first_as: %v", err)
		}
	}

	if err = d.Set("cluster_id", dataSourceFlattenRouterBgpClusterId(o["cluster-id"], d, "cluster_id")); err != nil {
		if !fortiAPIPatch(o["cluster-id"]) {
			return fmt.Errorf("Error reading cluster_id: %v", err)
		}
	}

	if err = d.Set("scan_time", dataSourceFlattenRouterBgpScanTime(o["scan-time"], d, "scan_time")); err != nil {
		if !fortiAPIPatch(o["scan-time"]) {
			return fmt.Errorf("Error reading scan_time: %v", err)
		}
	}

	if err = d.Set("bestpath_med_missing_as_worst", dataSourceFlattenRouterBgpBestpathMedMissingAsWorst(o["bestpath-med-missing-as-worst"], d, "bestpath_med_missing_as_worst")); err != nil {
		if !fortiAPIPatch(o["bestpath-med-missing-as-worst"]) {
			return fmt.Errorf("Error reading bestpath_med_missing_as_worst: %v", err)
		}
	}

	if err = d.Set("holdtime_timer", dataSourceFlattenRouterBgpHoldtimeTimer(o["holdtime-timer"], d, "holdtime_timer")); err != nil {
		if !fortiAPIPatch(o["holdtime-timer"]) {
			return fmt.Errorf("Error reading holdtime_timer: %v", err)
		}
	}

	if err = d.Set("network6", dataSourceFlattenRouterBgpNetwork6(o["network6"], d, "network6")); err != nil {
		if !fortiAPIPatch(o["network6"]) {
			return fmt.Errorf("Error reading network6: %v", err)
		}
	}

	if err = d.Set("dampening_reuse", dataSourceFlattenRouterBgpDampeningReuse(o["dampening-reuse"], d, "dampening_reuse")); err != nil {
		if !fortiAPIPatch(o["dampening-reuse"]) {
			return fmt.Errorf("Error reading dampening_reuse: %v", err)
		}
	}

	if err = d.Set("always_compare_med", dataSourceFlattenRouterBgpAlwaysCompareMed(o["always-compare-med"], d, "always_compare_med")); err != nil {
		if !fortiAPIPatch(o["always-compare-med"]) {
			return fmt.Errorf("Error reading always_compare_med: %v", err)
		}
	}

	if err = d.Set("maximum_paths_ebgp", dataSourceFlattenRouterBgpMaximumPathsEbgp(o["maximum-paths-ebgp"], d, "maximum_paths_ebgp")); err != nil {
		if !fortiAPIPatch(o["maximum-paths-ebgp"]) {
			return fmt.Errorf("Error reading maximum_paths_ebgp: %v", err)
		}
	}

	if err = d.Set("bestpath_as_path_ignore", dataSourceFlattenRouterBgpBestpathAsPathIgnore(o["bestpath-as-path-ignore"], d, "bestpath_as_path_ignore")); err != nil {
		if !fortiAPIPatch(o["bestpath-as-path-ignore"]) {
			return fmt.Errorf("Error reading bestpath_as_path_ignore: %v", err)
		}
	}

	if err = d.Set("redistribute", dataSourceFlattenRouterBgpRedistribute(o["redistribute"], d, "redistribute")); err != nil {
		if !fortiAPIPatch(o["redistribute"]) {
			return fmt.Errorf("Error reading redistribute: %v", err)
		}
	}

	if err = d.Set("client_to_client_reflection", dataSourceFlattenRouterBgpClientToClientReflection(o["client-to-client-reflection"], d, "client_to_client_reflection")); err != nil {
		if !fortiAPIPatch(o["client-to-client-reflection"]) {
			return fmt.Errorf("Error reading client_to_client_reflection: %v", err)
		}
	}

	if err = d.Set("dampening_max_suppress_time", dataSourceFlattenRouterBgpDampeningMaxSuppressTime(o["dampening-max-suppress-time"], d, "dampening_max_suppress_time")); err != nil {
		if !fortiAPIPatch(o["dampening-max-suppress-time"]) {
			return fmt.Errorf("Error reading dampening_max_suppress_time: %v", err)
		}
	}

	if err = d.Set("deterministic_med", dataSourceFlattenRouterBgpDeterministicMed(o["deterministic-med"], d, "deterministic_med")); err != nil {
		if !fortiAPIPatch(o["deterministic-med"]) {
			return fmt.Errorf("Error reading deterministic_med: %v", err)
		}
	}

	if err = d.Set("fast_external_failover", dataSourceFlattenRouterBgpFastExternalFailover(o["fast-external-failover"], d, "fast_external_failover")); err != nil {
		if !fortiAPIPatch(o["fast-external-failover"]) {
			return fmt.Errorf("Error reading fast_external_failover: %v", err)
		}
	}

	if err = d.Set("maximum_paths_ibgp", dataSourceFlattenRouterBgpMaximumPathsIbgp(o["maximum-paths-ibgp"], d, "maximum_paths_ibgp")); err != nil {
		if !fortiAPIPatch(o["maximum-paths-ibgp"]) {
			return fmt.Errorf("Error reading maximum_paths_ibgp: %v", err)
		}
	}

	if err = d.Set("redistribute6", dataSourceFlattenRouterBgpRedistribute6(o["redistribute6"], d, "redistribute6")); err != nil {
		if !fortiAPIPatch(o["redistribute6"]) {
			return fmt.Errorf("Error reading redistribute6: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", dataSourceFlattenRouterBgpLogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes")); err != nil {
		if !fortiAPIPatch(o["log-neighbour-changes"]) {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	if err = d.Set("neighbor", dataSourceFlattenRouterBgpNeighbor(o["neighbor"], d, "neighbor")); err != nil {
		if !fortiAPIPatch(o["neighbor"]) {
			return fmt.Errorf("Error reading neighbor: %v", err)
		}
	}

	if err = d.Set("route_reflector_allow_outbound_policy", dataSourceFlattenRouterBgpRouteReflectorAllowOutboundPolicy(o["route-reflector-allow-outbound-policy"], d, "route_reflector_allow_outbound_policy")); err != nil {
		if !fortiAPIPatch(o["route-reflector-allow-outbound-policy"]) {
			return fmt.Errorf("Error reading route_reflector_allow_outbound_policy: %v", err)
		}
	}

	if err = d.Set("admin_distance6", dataSourceFlattenRouterBgpAdminDistance6(o["admin-distance6"], d, "admin_distance6")); err != nil {
		if !fortiAPIPatch(o["admin-distance6"]) {
			return fmt.Errorf("Error reading admin_distance6: %v", err)
		}
	}

	if err = d.Set("confederation_identifier", dataSourceFlattenRouterBgpConfederationIdentifier(o["confederation-identifier"], d, "confederation_identifier")); err != nil {
		if !fortiAPIPatch(o["confederation-identifier"]) {
			return fmt.Errorf("Error reading confederation_identifier: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterBgpFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
