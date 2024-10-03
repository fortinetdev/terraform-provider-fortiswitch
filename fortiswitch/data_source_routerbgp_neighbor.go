// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: BGP neighbor table.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceRouterbgpNeighbor() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterbgpNeighborRead,
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
				Required: true,
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
			"route_map_in_evpn": &schema.Schema{
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
			"capability_extended_nexthop": &schema.Schema{
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
			"activate_evpn": &schema.Schema{
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
			"attribute_unchanged_evpn": &schema.Schema{
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
			"soft_reconfiguration_evpn": &schema.Schema{
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
			"route_reflector_client_evpn": &schema.Schema{
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
			"route_map_out_evpn": &schema.Schema{
				Type:     schema.TypeString,
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
			"allowas_in_enable_evpn": &schema.Schema{
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
	}
}

func dataSourceRouterbgpNeighborRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("ip")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterbgpNeighbor: type error")
	}

	o, err := c.ReadRouterbgpNeighbor(mkey)
	if err != nil {
		return fmt.Errorf("Error describing RouterbgpNeighbor: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterbgpNeighbor(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterbgpNeighbor from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterbgpNeighborSendCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborActivate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborFilterListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAttributeUnchanged(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborFilterListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborEbgpMultihopTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDefaultOriginateRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDefaultOriginateRoutemap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteReflectorClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteMapOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRemovePrivateAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborShutdown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteMapIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteMapInEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborUnsuppressMap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborUnsuppressMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAsOverride6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAttributeUnchanged6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborEbgpEnforceMultihop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAdvertisementInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborPrefixListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborCapabilityDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborCapabilityExtendedNexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborCapabilityOrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborNextHopSelf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAllowasInEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteReflectorClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborActivateEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDontCapabilityNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborConnectTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborPassive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAttributeUnchangedEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAllowasIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborMaximumPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteServerClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborMaximumPrefixThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborFilterListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborEnforceFirstAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborSoftReconfigurationEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborKeepAliveTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborMaximumPrefixWarningOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAsOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteReflectorClientEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborBfdSessionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDistributeListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborCapabilityOrf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborSoftReconfiguration6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborMaximumPrefixWarningOnly6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborNextHopSelf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAllowasInEnable6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAllowasIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteMapOutEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborUpdateSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRemovePrivateAs6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteMapIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborActivate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborFilterListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborMaximumPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteServerClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborCapabilityDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborAllowasInEnableEvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborEbgpTtlSecurityHops(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDistributeListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborOverrideCapability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDistributeListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborSendCommunity6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRouteMapOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborPrefixListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborCapabilityDefaultOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborStrictCapabilityMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborPrefixListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborDistributeListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborRemoteAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborMaximumPrefixThreshold6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborPrefixListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterbgpNeighborSoftReconfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterbgpNeighbor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("send_community", dataSourceFlattenRouterbgpNeighborSendCommunity(o["send-community"], d, "send_community")); err != nil {
		if !fortiAPIPatch(o["send-community"]) {
			return fmt.Errorf("Error reading send_community: %v", err)
		}
	}

	if err = d.Set("activate", dataSourceFlattenRouterbgpNeighborActivate(o["activate"], d, "activate")); err != nil {
		if !fortiAPIPatch(o["activate"]) {
			return fmt.Errorf("Error reading activate: %v", err)
		}
	}

	if err = d.Set("filter_list_out6", dataSourceFlattenRouterbgpNeighborFilterListOut6(o["filter-list-out6"], d, "filter_list_out6")); err != nil {
		if !fortiAPIPatch(o["filter-list-out6"]) {
			return fmt.Errorf("Error reading filter_list_out6: %v", err)
		}
	}

	if err = d.Set("weight", dataSourceFlattenRouterbgpNeighborWeight(o["weight"], d, "weight")); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged", dataSourceFlattenRouterbgpNeighborAttributeUnchanged(o["attribute-unchanged"], d, "attribute_unchanged")); err != nil {
		if !fortiAPIPatch(o["attribute-unchanged"]) {
			return fmt.Errorf("Error reading attribute_unchanged: %v", err)
		}
	}

	if err = d.Set("ip", dataSourceFlattenRouterbgpNeighborIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("filter_list_in6", dataSourceFlattenRouterbgpNeighborFilterListIn6(o["filter-list-in6"], d, "filter_list_in6")); err != nil {
		if !fortiAPIPatch(o["filter-list-in6"]) {
			return fmt.Errorf("Error reading filter_list_in6: %v", err)
		}
	}

	if err = d.Set("ebgp_multihop_ttl", dataSourceFlattenRouterbgpNeighborEbgpMultihopTtl(o["ebgp-multihop-ttl"], d, "ebgp_multihop_ttl")); err != nil {
		if !fortiAPIPatch(o["ebgp-multihop-ttl"]) {
			return fmt.Errorf("Error reading ebgp_multihop_ttl: %v", err)
		}
	}

	if err = d.Set("default_originate_routemap", dataSourceFlattenRouterbgpNeighborDefaultOriginateRoutemap(o["default-originate-routemap"], d, "default_originate_routemap")); err != nil {
		if !fortiAPIPatch(o["default-originate-routemap"]) {
			return fmt.Errorf("Error reading default_originate_routemap: %v", err)
		}
	}

	if err = d.Set("default_originate_routemap6", dataSourceFlattenRouterbgpNeighborDefaultOriginateRoutemap6(o["default-originate-routemap6"], d, "default_originate_routemap6")); err != nil {
		if !fortiAPIPatch(o["default-originate-routemap6"]) {
			return fmt.Errorf("Error reading default_originate_routemap6: %v", err)
		}
	}

	if err = d.Set("route_reflector_client", dataSourceFlattenRouterbgpNeighborRouteReflectorClient(o["route-reflector-client"], d, "route_reflector_client")); err != nil {
		if !fortiAPIPatch(o["route-reflector-client"]) {
			return fmt.Errorf("Error reading route_reflector_client: %v", err)
		}
	}

	if err = d.Set("route_map_out6", dataSourceFlattenRouterbgpNeighborRouteMapOut6(o["route-map-out6"], d, "route_map_out6")); err != nil {
		if !fortiAPIPatch(o["route-map-out6"]) {
			return fmt.Errorf("Error reading route_map_out6: %v", err)
		}
	}

	if err = d.Set("remove_private_as", dataSourceFlattenRouterbgpNeighborRemovePrivateAs(o["remove-private-as"], d, "remove_private_as")); err != nil {
		if !fortiAPIPatch(o["remove-private-as"]) {
			return fmt.Errorf("Error reading remove_private_as: %v", err)
		}
	}

	if err = d.Set("shutdown", dataSourceFlattenRouterbgpNeighborShutdown(o["shutdown"], d, "shutdown")); err != nil {
		if !fortiAPIPatch(o["shutdown"]) {
			return fmt.Errorf("Error reading shutdown: %v", err)
		}
	}

	if err = d.Set("route_map_in6", dataSourceFlattenRouterbgpNeighborRouteMapIn6(o["route-map-in6"], d, "route_map_in6")); err != nil {
		if !fortiAPIPatch(o["route-map-in6"]) {
			return fmt.Errorf("Error reading route_map_in6: %v", err)
		}
	}

	if err = d.Set("route_map_in_evpn", dataSourceFlattenRouterbgpNeighborRouteMapInEvpn(o["route-map-in-evpn"], d, "route_map_in_evpn")); err != nil {
		if !fortiAPIPatch(o["route-map-in-evpn"]) {
			return fmt.Errorf("Error reading route_map_in_evpn: %v", err)
		}
	}

	if err = d.Set("unsuppress_map6", dataSourceFlattenRouterbgpNeighborUnsuppressMap6(o["unsuppress-map6"], d, "unsuppress_map6")); err != nil {
		if !fortiAPIPatch(o["unsuppress-map6"]) {
			return fmt.Errorf("Error reading unsuppress_map6: %v", err)
		}
	}

	if err = d.Set("unsuppress_map", dataSourceFlattenRouterbgpNeighborUnsuppressMap(o["unsuppress-map"], d, "unsuppress_map")); err != nil {
		if !fortiAPIPatch(o["unsuppress-map"]) {
			return fmt.Errorf("Error reading unsuppress_map: %v", err)
		}
	}

	if err = d.Set("as_override6", dataSourceFlattenRouterbgpNeighborAsOverride6(o["as-override6"], d, "as_override6")); err != nil {
		if !fortiAPIPatch(o["as-override6"]) {
			return fmt.Errorf("Error reading as_override6: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged6", dataSourceFlattenRouterbgpNeighborAttributeUnchanged6(o["attribute-unchanged6"], d, "attribute_unchanged6")); err != nil {
		if !fortiAPIPatch(o["attribute-unchanged6"]) {
			return fmt.Errorf("Error reading attribute_unchanged6: %v", err)
		}
	}

	if err = d.Set("ebgp_enforce_multihop", dataSourceFlattenRouterbgpNeighborEbgpEnforceMultihop(o["ebgp-enforce-multihop"], d, "ebgp_enforce_multihop")); err != nil {
		if !fortiAPIPatch(o["ebgp-enforce-multihop"]) {
			return fmt.Errorf("Error reading ebgp_enforce_multihop: %v", err)
		}
	}

	if err = d.Set("advertisement_interval", dataSourceFlattenRouterbgpNeighborAdvertisementInterval(o["advertisement-interval"], d, "advertisement_interval")); err != nil {
		if !fortiAPIPatch(o["advertisement-interval"]) {
			return fmt.Errorf("Error reading advertisement_interval: %v", err)
		}
	}

	if err = d.Set("prefix_list_in6", dataSourceFlattenRouterbgpNeighborPrefixListIn6(o["prefix-list-in6"], d, "prefix_list_in6")); err != nil {
		if !fortiAPIPatch(o["prefix-list-in6"]) {
			return fmt.Errorf("Error reading prefix_list_in6: %v", err)
		}
	}

	if err = d.Set("capability_default_originate6", dataSourceFlattenRouterbgpNeighborCapabilityDefaultOriginate6(o["capability-default-originate6"], d, "capability_default_originate6")); err != nil {
		if !fortiAPIPatch(o["capability-default-originate6"]) {
			return fmt.Errorf("Error reading capability_default_originate6: %v", err)
		}
	}

	if err = d.Set("capability_extended_nexthop", dataSourceFlattenRouterbgpNeighborCapabilityExtendedNexthop(o["capability-extended-nexthop"], d, "capability_extended_nexthop")); err != nil {
		if !fortiAPIPatch(o["capability-extended-nexthop"]) {
			return fmt.Errorf("Error reading capability_extended_nexthop: %v", err)
		}
	}

	if err = d.Set("bfd", dataSourceFlattenRouterbgpNeighborBfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("capability_orf", dataSourceFlattenRouterbgpNeighborCapabilityOrf(o["capability-orf"], d, "capability_orf")); err != nil {
		if !fortiAPIPatch(o["capability-orf"]) {
			return fmt.Errorf("Error reading capability_orf: %v", err)
		}
	}

	if err = d.Set("next_hop_self", dataSourceFlattenRouterbgpNeighborNextHopSelf(o["next-hop-self"], d, "next_hop_self")); err != nil {
		if !fortiAPIPatch(o["next-hop-self"]) {
			return fmt.Errorf("Error reading next_hop_self: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable", dataSourceFlattenRouterbgpNeighborAllowasInEnable(o["allowas-in-enable"], d, "allowas_in_enable")); err != nil {
		if !fortiAPIPatch(o["allowas-in-enable"]) {
			return fmt.Errorf("Error reading allowas_in_enable: %v", err)
		}
	}

	if err = d.Set("route_reflector_client6", dataSourceFlattenRouterbgpNeighborRouteReflectorClient6(o["route-reflector-client6"], d, "route_reflector_client6")); err != nil {
		if !fortiAPIPatch(o["route-reflector-client6"]) {
			return fmt.Errorf("Error reading route_reflector_client6: %v", err)
		}
	}

	if err = d.Set("activate_evpn", dataSourceFlattenRouterbgpNeighborActivateEvpn(o["activate-evpn"], d, "activate_evpn")); err != nil {
		if !fortiAPIPatch(o["activate-evpn"]) {
			return fmt.Errorf("Error reading activate_evpn: %v", err)
		}
	}

	if err = d.Set("dont_capability_negotiate", dataSourceFlattenRouterbgpNeighborDontCapabilityNegotiate(o["dont-capability-negotiate"], d, "dont_capability_negotiate")); err != nil {
		if !fortiAPIPatch(o["dont-capability-negotiate"]) {
			return fmt.Errorf("Error reading dont_capability_negotiate: %v", err)
		}
	}

	if err = d.Set("connect_timer", dataSourceFlattenRouterbgpNeighborConnectTimer(o["connect-timer"], d, "connect_timer")); err != nil {
		if !fortiAPIPatch(o["connect-timer"]) {
			return fmt.Errorf("Error reading connect_timer: %v", err)
		}
	}

	if err = d.Set("passive", dataSourceFlattenRouterbgpNeighborPassive(o["passive"], d, "passive")); err != nil {
		if !fortiAPIPatch(o["passive"]) {
			return fmt.Errorf("Error reading passive: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged_evpn", dataSourceFlattenRouterbgpNeighborAttributeUnchangedEvpn(o["attribute-unchanged-evpn"], d, "attribute_unchanged_evpn")); err != nil {
		if !fortiAPIPatch(o["attribute-unchanged-evpn"]) {
			return fmt.Errorf("Error reading attribute_unchanged_evpn: %v", err)
		}
	}

	if err = d.Set("allowas_in", dataSourceFlattenRouterbgpNeighborAllowasIn(o["allowas-in"], d, "allowas_in")); err != nil {
		if !fortiAPIPatch(o["allowas-in"]) {
			return fmt.Errorf("Error reading allowas_in: %v", err)
		}
	}

	if err = d.Set("maximum_prefix6", dataSourceFlattenRouterbgpNeighborMaximumPrefix6(o["maximum-prefix6"], d, "maximum_prefix6")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix6"]) {
			return fmt.Errorf("Error reading maximum_prefix6: %v", err)
		}
	}

	if err = d.Set("route_server_client", dataSourceFlattenRouterbgpNeighborRouteServerClient(o["route-server-client"], d, "route_server_client")); err != nil {
		if !fortiAPIPatch(o["route-server-client"]) {
			return fmt.Errorf("Error reading route_server_client: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold", dataSourceFlattenRouterbgpNeighborMaximumPrefixThreshold(o["maximum-prefix-threshold"], d, "maximum_prefix_threshold")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-threshold"]) {
			return fmt.Errorf("Error reading maximum_prefix_threshold: %v", err)
		}
	}

	if err = d.Set("filter_list_out", dataSourceFlattenRouterbgpNeighborFilterListOut(o["filter-list-out"], d, "filter_list_out")); err != nil {
		if !fortiAPIPatch(o["filter-list-out"]) {
			return fmt.Errorf("Error reading filter_list_out: %v", err)
		}
	}

	if err = d.Set("enforce_first_as", dataSourceFlattenRouterbgpNeighborEnforceFirstAs(o["enforce-first-as"], d, "enforce_first_as")); err != nil {
		if !fortiAPIPatch(o["enforce-first-as"]) {
			return fmt.Errorf("Error reading enforce_first_as: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration_evpn", dataSourceFlattenRouterbgpNeighborSoftReconfigurationEvpn(o["soft-reconfiguration-evpn"], d, "soft_reconfiguration_evpn")); err != nil {
		if !fortiAPIPatch(o["soft-reconfiguration-evpn"]) {
			return fmt.Errorf("Error reading soft_reconfiguration_evpn: %v", err)
		}
	}

	if err = d.Set("keep_alive_timer", dataSourceFlattenRouterbgpNeighborKeepAliveTimer(o["keep-alive-timer"], d, "keep_alive_timer")); err != nil {
		if !fortiAPIPatch(o["keep-alive-timer"]) {
			return fmt.Errorf("Error reading keep_alive_timer: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only", dataSourceFlattenRouterbgpNeighborMaximumPrefixWarningOnly(o["maximum-prefix-warning-only"], d, "maximum_prefix_warning_only")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-warning-only"]) {
			return fmt.Errorf("Error reading maximum_prefix_warning_only: %v", err)
		}
	}

	if err = d.Set("description", dataSourceFlattenRouterbgpNeighborDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("as_override", dataSourceFlattenRouterbgpNeighborAsOverride(o["as-override"], d, "as_override")); err != nil {
		if !fortiAPIPatch(o["as-override"]) {
			return fmt.Errorf("Error reading as_override: %v", err)
		}
	}

	if err = d.Set("route_reflector_client_evpn", dataSourceFlattenRouterbgpNeighborRouteReflectorClientEvpn(o["route-reflector-client-evpn"], d, "route_reflector_client_evpn")); err != nil {
		if !fortiAPIPatch(o["route-reflector-client-evpn"]) {
			return fmt.Errorf("Error reading route_reflector_client_evpn: %v", err)
		}
	}

	if err = d.Set("bfd_session_mode", dataSourceFlattenRouterbgpNeighborBfdSessionMode(o["bfd-session-mode"], d, "bfd_session_mode")); err != nil {
		if !fortiAPIPatch(o["bfd-session-mode"]) {
			return fmt.Errorf("Error reading bfd_session_mode: %v", err)
		}
	}

	if err = d.Set("distribute_list_out", dataSourceFlattenRouterbgpNeighborDistributeListOut(o["distribute-list-out"], d, "distribute_list_out")); err != nil {
		if !fortiAPIPatch(o["distribute-list-out"]) {
			return fmt.Errorf("Error reading distribute_list_out: %v", err)
		}
	}

	if err = d.Set("capability_orf6", dataSourceFlattenRouterbgpNeighborCapabilityOrf6(o["capability-orf6"], d, "capability_orf6")); err != nil {
		if !fortiAPIPatch(o["capability-orf6"]) {
			return fmt.Errorf("Error reading capability_orf6: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration6", dataSourceFlattenRouterbgpNeighborSoftReconfiguration6(o["soft-reconfiguration6"], d, "soft_reconfiguration6")); err != nil {
		if !fortiAPIPatch(o["soft-reconfiguration6"]) {
			return fmt.Errorf("Error reading soft_reconfiguration6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only6", dataSourceFlattenRouterbgpNeighborMaximumPrefixWarningOnly6(o["maximum-prefix-warning-only6"], d, "maximum_prefix_warning_only6")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-warning-only6"]) {
			return fmt.Errorf("Error reading maximum_prefix_warning_only6: %v", err)
		}
	}

	if err = d.Set("next_hop_self6", dataSourceFlattenRouterbgpNeighborNextHopSelf6(o["next-hop-self6"], d, "next_hop_self6")); err != nil {
		if !fortiAPIPatch(o["next-hop-self6"]) {
			return fmt.Errorf("Error reading next_hop_self6: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable6", dataSourceFlattenRouterbgpNeighborAllowasInEnable6(o["allowas-in-enable6"], d, "allowas_in_enable6")); err != nil {
		if !fortiAPIPatch(o["allowas-in-enable6"]) {
			return fmt.Errorf("Error reading allowas_in_enable6: %v", err)
		}
	}

	if err = d.Set("allowas_in6", dataSourceFlattenRouterbgpNeighborAllowasIn6(o["allowas-in6"], d, "allowas_in6")); err != nil {
		if !fortiAPIPatch(o["allowas-in6"]) {
			return fmt.Errorf("Error reading allowas_in6: %v", err)
		}
	}

	if err = d.Set("route_map_out_evpn", dataSourceFlattenRouterbgpNeighborRouteMapOutEvpn(o["route-map-out-evpn"], d, "route_map_out_evpn")); err != nil {
		if !fortiAPIPatch(o["route-map-out-evpn"]) {
			return fmt.Errorf("Error reading route_map_out_evpn: %v", err)
		}
	}

	if err = d.Set("update_source", dataSourceFlattenRouterbgpNeighborUpdateSource(o["update-source"], d, "update_source")); err != nil {
		if !fortiAPIPatch(o["update-source"]) {
			return fmt.Errorf("Error reading update_source: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenRouterbgpNeighborInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("remove_private_as6", dataSourceFlattenRouterbgpNeighborRemovePrivateAs6(o["remove-private-as6"], d, "remove_private_as6")); err != nil {
		if !fortiAPIPatch(o["remove-private-as6"]) {
			return fmt.Errorf("Error reading remove_private_as6: %v", err)
		}
	}

	if err = d.Set("password", dataSourceFlattenRouterbgpNeighborPassword(o["password"], d, "password")); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("holdtime_timer", dataSourceFlattenRouterbgpNeighborHoldtimeTimer(o["holdtime-timer"], d, "holdtime_timer")); err != nil {
		if !fortiAPIPatch(o["holdtime-timer"]) {
			return fmt.Errorf("Error reading holdtime_timer: %v", err)
		}
	}

	if err = d.Set("route_map_in", dataSourceFlattenRouterbgpNeighborRouteMapIn(o["route-map-in"], d, "route_map_in")); err != nil {
		if !fortiAPIPatch(o["route-map-in"]) {
			return fmt.Errorf("Error reading route_map_in: %v", err)
		}
	}

	if err = d.Set("activate6", dataSourceFlattenRouterbgpNeighborActivate6(o["activate6"], d, "activate6")); err != nil {
		if !fortiAPIPatch(o["activate6"]) {
			return fmt.Errorf("Error reading activate6: %v", err)
		}
	}

	if err = d.Set("filter_list_in", dataSourceFlattenRouterbgpNeighborFilterListIn(o["filter-list-in"], d, "filter_list_in")); err != nil {
		if !fortiAPIPatch(o["filter-list-in"]) {
			return fmt.Errorf("Error reading filter_list_in: %v", err)
		}
	}

	if err = d.Set("maximum_prefix", dataSourceFlattenRouterbgpNeighborMaximumPrefix(o["maximum-prefix"], d, "maximum_prefix")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix"]) {
			return fmt.Errorf("Error reading maximum_prefix: %v", err)
		}
	}

	if err = d.Set("route_server_client6", dataSourceFlattenRouterbgpNeighborRouteServerClient6(o["route-server-client6"], d, "route_server_client6")); err != nil {
		if !fortiAPIPatch(o["route-server-client6"]) {
			return fmt.Errorf("Error reading route_server_client6: %v", err)
		}
	}

	if err = d.Set("capability_dynamic", dataSourceFlattenRouterbgpNeighborCapabilityDynamic(o["capability-dynamic"], d, "capability_dynamic")); err != nil {
		if !fortiAPIPatch(o["capability-dynamic"]) {
			return fmt.Errorf("Error reading capability_dynamic: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable_evpn", dataSourceFlattenRouterbgpNeighborAllowasInEnableEvpn(o["allowas-in-enable-evpn"], d, "allowas_in_enable_evpn")); err != nil {
		if !fortiAPIPatch(o["allowas-in-enable-evpn"]) {
			return fmt.Errorf("Error reading allowas_in_enable_evpn: %v", err)
		}
	}

	if err = d.Set("ebgp_ttl_security_hops", dataSourceFlattenRouterbgpNeighborEbgpTtlSecurityHops(o["ebgp-ttl-security-hops"], d, "ebgp_ttl_security_hops")); err != nil {
		if !fortiAPIPatch(o["ebgp-ttl-security-hops"]) {
			return fmt.Errorf("Error reading ebgp_ttl_security_hops: %v", err)
		}
	}

	if err = d.Set("distribute_list_in6", dataSourceFlattenRouterbgpNeighborDistributeListIn6(o["distribute-list-in6"], d, "distribute_list_in6")); err != nil {
		if !fortiAPIPatch(o["distribute-list-in6"]) {
			return fmt.Errorf("Error reading distribute_list_in6: %v", err)
		}
	}

	if err = d.Set("override_capability", dataSourceFlattenRouterbgpNeighborOverrideCapability(o["override-capability"], d, "override_capability")); err != nil {
		if !fortiAPIPatch(o["override-capability"]) {
			return fmt.Errorf("Error reading override_capability: %v", err)
		}
	}

	if err = d.Set("distribute_list_out6", dataSourceFlattenRouterbgpNeighborDistributeListOut6(o["distribute-list-out6"], d, "distribute_list_out6")); err != nil {
		if !fortiAPIPatch(o["distribute-list-out6"]) {
			return fmt.Errorf("Error reading distribute_list_out6: %v", err)
		}
	}

	if err = d.Set("send_community6", dataSourceFlattenRouterbgpNeighborSendCommunity6(o["send-community6"], d, "send_community6")); err != nil {
		if !fortiAPIPatch(o["send-community6"]) {
			return fmt.Errorf("Error reading send_community6: %v", err)
		}
	}

	if err = d.Set("route_map_out", dataSourceFlattenRouterbgpNeighborRouteMapOut(o["route-map-out"], d, "route_map_out")); err != nil {
		if !fortiAPIPatch(o["route-map-out"]) {
			return fmt.Errorf("Error reading route_map_out: %v", err)
		}
	}

	if err = d.Set("prefix_list_out6", dataSourceFlattenRouterbgpNeighborPrefixListOut6(o["prefix-list-out6"], d, "prefix_list_out6")); err != nil {
		if !fortiAPIPatch(o["prefix-list-out6"]) {
			return fmt.Errorf("Error reading prefix_list_out6: %v", err)
		}
	}

	if err = d.Set("capability_default_originate", dataSourceFlattenRouterbgpNeighborCapabilityDefaultOriginate(o["capability-default-originate"], d, "capability_default_originate")); err != nil {
		if !fortiAPIPatch(o["capability-default-originate"]) {
			return fmt.Errorf("Error reading capability_default_originate: %v", err)
		}
	}

	if err = d.Set("strict_capability_match", dataSourceFlattenRouterbgpNeighborStrictCapabilityMatch(o["strict-capability-match"], d, "strict_capability_match")); err != nil {
		if !fortiAPIPatch(o["strict-capability-match"]) {
			return fmt.Errorf("Error reading strict_capability_match: %v", err)
		}
	}

	if err = d.Set("prefix_list_in", dataSourceFlattenRouterbgpNeighborPrefixListIn(o["prefix-list-in"], d, "prefix_list_in")); err != nil {
		if !fortiAPIPatch(o["prefix-list-in"]) {
			return fmt.Errorf("Error reading prefix_list_in: %v", err)
		}
	}

	if err = d.Set("distribute_list_in", dataSourceFlattenRouterbgpNeighborDistributeListIn(o["distribute-list-in"], d, "distribute_list_in")); err != nil {
		if !fortiAPIPatch(o["distribute-list-in"]) {
			return fmt.Errorf("Error reading distribute_list_in: %v", err)
		}
	}

	if err = d.Set("remote_as", dataSourceFlattenRouterbgpNeighborRemoteAs(o["remote-as"], d, "remote_as")); err != nil {
		if !fortiAPIPatch(o["remote-as"]) {
			return fmt.Errorf("Error reading remote_as: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold6", dataSourceFlattenRouterbgpNeighborMaximumPrefixThreshold6(o["maximum-prefix-threshold6"], d, "maximum_prefix_threshold6")); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-threshold6"]) {
			return fmt.Errorf("Error reading maximum_prefix_threshold6: %v", err)
		}
	}

	if err = d.Set("prefix_list_out", dataSourceFlattenRouterbgpNeighborPrefixListOut(o["prefix-list-out"], d, "prefix_list_out")); err != nil {
		if !fortiAPIPatch(o["prefix-list-out"]) {
			return fmt.Errorf("Error reading prefix_list_out: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration", dataSourceFlattenRouterbgpNeighborSoftReconfiguration(o["soft-reconfiguration"], d, "soft_reconfiguration")); err != nil {
		if !fortiAPIPatch(o["soft-reconfiguration"]) {
			return fmt.Errorf("Error reading soft_reconfiguration: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterbgpNeighborFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
