// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure optional per-VLAN settings.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchVlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchVlanCreate,
		Read:   resourceSwitchVlanRead,
		Update: resourceSwitchVlanUpdate,
		Delete: resourceSwitchVlanDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"community_vlans": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lan_subvlans": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp_inspection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"igmp_snooping_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcp6_snooping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lan_segment_primary_vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"learning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fswid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"igmp_snooping_querier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_snooping_static_client": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"switch_interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"member_by_mac": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"lan_segment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary_vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"igmp_snooping_querier_version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cos_queue": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"member_by_ipv6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"mld_snooping_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mld_snooping_fast_leave": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"igmp_snooping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mld_snooping_querier_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_server_access_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"server_ip6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"learning_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mrouter_ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"mld_snooping_static_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ignore_reports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mcast_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"members": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"member_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"igmp_snooping_querier_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rspan_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private_vlan_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mld_snooping_querier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"igmp_snooping_fast_leave": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_snooping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lan_segment_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"igmp_snooping_static_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ignore_reports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mcast_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"members": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"member_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
					},
				},
			},
			"dhcp_snooping_verify_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"isolated_vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcp_snooping_option82": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"member_by_ipv4": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"private_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"member_by_proto": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"frametypes": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"access_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mld_snooping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchVlanCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchVlan(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchVlan resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchVlan(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchVlan resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchVlan")
	}

	return resourceSwitchVlanRead(d, m)
}

func resourceSwitchVlanUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchVlan(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchVlan resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchVlan(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchVlan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchVlan")
	}

	return resourceSwitchVlanRead(d, m)
}

func resourceSwitchVlanDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchVlan(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchVlan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchVlanRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchVlan(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchVlan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchVlan(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchVlan resource from API: %v", err)
	}
	return nil
}

func flattenSwitchVlanCommunityVlans(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanLanSubvlans(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanArpInspection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanIgmpSnoopingProxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanPolicer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanDhcp6Snooping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanLanSegmentPrimaryVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanLearning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanIgmpSnoopingQuerier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanDhcpSnoopingStaticClient(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_addr"
		if _, ok := i["ip-addr"]; ok {

			tmp["ip_addr"] = flattenSwitchVlanDhcpSnoopingStaticClientIpAddr(i["ip-addr"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr"
		if _, ok := i["mac-addr"]; ok {

			tmp["mac_addr"] = flattenSwitchVlanDhcpSnoopingStaticClientMacAddr(i["mac-addr"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchVlanDhcpSnoopingStaticClientName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_interface"
		if _, ok := i["switch-interface"]; ok {

			tmp["switch_interface"] = flattenSwitchVlanDhcpSnoopingStaticClientSwitchInterface(i["switch-interface"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchVlanDhcpSnoopingStaticClientIpAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanDhcpSnoopingStaticClientMacAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanDhcpSnoopingStaticClientName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanDhcpSnoopingStaticClientSwitchInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMemberByMac(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {

			tmp["mac"] = flattenSwitchVlanMemberByMacMac(i["mac"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSwitchVlanMemberByMacId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchVlanMemberByMacDescription(i["description"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchVlanMemberByMacMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMemberByMacId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMemberByMacDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanLanSegment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanPrimaryVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanIgmpSnoopingQuerierVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanCosQueue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMemberByIpv6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = flattenSwitchVlanMemberByIpv6Prefix(i["prefix"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSwitchVlanMemberByIpv6Id(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchVlanMemberByIpv6Description(i["description"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchVlanMemberByIpv6Prefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMemberByIpv6Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMemberByIpv6Description(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMldSnoopingProxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMldSnoopingFastLeave(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanIgmpSnooping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMldSnoopingQuerierAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanDhcpServerAccessList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_ip"
		if _, ok := i["server-ip"]; ok {

			tmp["server_ip"] = flattenSwitchVlanDhcpServerAccessListServerIp(i["server-ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchVlanDhcpServerAccessListName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_ip6"
		if _, ok := i["server-ip6"]; ok {

			tmp["server_ip6"] = flattenSwitchVlanDhcpServerAccessListServerIp6(i["server-ip6"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchVlanDhcpServerAccessListServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanDhcpServerAccessListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanDhcpServerAccessListServerIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanLearningLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMrouterPorts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_name"
		if _, ok := i["member-name"]; ok {

			tmp["member_name"] = flattenSwitchVlanMrouterPortsMemberName(i["member-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "member_name", d)
	return result
}

func flattenSwitchVlanMrouterPortsMemberName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMldSnoopingStaticGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_reports"
		if _, ok := i["ignore-reports"]; ok {

			tmp["ignore_reports"] = flattenSwitchVlanMldSnoopingStaticGroupIgnoreReports(i["ignore-reports"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcast_addr"
		if _, ok := i["mcast-addr"]; ok {

			tmp["mcast_addr"] = flattenSwitchVlanMldSnoopingStaticGroupMcastAddr(i["mcast-addr"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchVlanMldSnoopingStaticGroupName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {

			tmp["members"] = flattenSwitchVlanMldSnoopingStaticGroupMembers(i["members"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchVlanMldSnoopingStaticGroupIgnoreReports(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMldSnoopingStaticGroupMcastAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMldSnoopingStaticGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMldSnoopingStaticGroupMembers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_name"
		if _, ok := i["member-name"]; ok {

			tmp["member_name"] = flattenSwitchVlanMldSnoopingStaticGroupMembersMemberName(i["member-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "member_name", d)
	return result
}

func flattenSwitchVlanMldSnoopingStaticGroupMembersMemberName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanIgmpSnoopingQuerierAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanRspanMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanPrivateVlanType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMldSnoopingQuerier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanIgmpSnoopingFastLeave(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanDhcpSnooping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanLanSegmentType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanIgmpSnoopingStaticGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_reports"
		if _, ok := i["ignore-reports"]; ok {

			tmp["ignore_reports"] = flattenSwitchVlanIgmpSnoopingStaticGroupIgnoreReports(i["ignore-reports"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcast_addr"
		if _, ok := i["mcast-addr"]; ok {

			tmp["mcast_addr"] = flattenSwitchVlanIgmpSnoopingStaticGroupMcastAddr(i["mcast-addr"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchVlanIgmpSnoopingStaticGroupName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {

			tmp["members"] = flattenSwitchVlanIgmpSnoopingStaticGroupMembers(i["members"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchVlanIgmpSnoopingStaticGroupIgnoreReports(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanIgmpSnoopingStaticGroupMcastAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanIgmpSnoopingStaticGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanIgmpSnoopingStaticGroupMembers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_name"
		if _, ok := i["member-name"]; ok {

			tmp["member_name"] = flattenSwitchVlanIgmpSnoopingStaticGroupMembersMemberName(i["member-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "member_name", d)
	return result
}

func flattenSwitchVlanIgmpSnoopingStaticGroupMembersMemberName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanDhcpSnoopingVerifyMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanIsolatedVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanDhcpSnoopingOption82(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMemberByIpv4(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchVlanMemberByIpv4Description(i["description"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSwitchVlanMemberByIpv4Id(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {

			tmp["address"] = flattenSwitchVlanMemberByIpv4Address(i["address"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchVlanMemberByIpv4Description(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMemberByIpv4Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMemberByIpv4Address(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchVlanPrivateVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMemberByProto(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {

			tmp["protocol"] = flattenSwitchVlanMemberByProtoProtocol(i["protocol"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "frametypes"
		if _, ok := i["frametypes"]; ok {

			tmp["frametypes"] = flattenSwitchVlanMemberByProtoFrametypes(i["frametypes"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSwitchVlanMemberByProtoId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchVlanMemberByProtoDescription(i["description"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchVlanMemberByProtoProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMemberByProtoFrametypes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMemberByProtoId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMemberByProtoDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanAccessVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanMldSnooping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchVlan(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("community_vlans", flattenSwitchVlanCommunityVlans(o["community-vlans"], d, "community_vlans", sv)); err != nil {
		if !fortiAPIPatch(o["community-vlans"]) {
			return fmt.Errorf("Error reading community_vlans: %v", err)
		}
	}

	if err = d.Set("lan_subvlans", flattenSwitchVlanLanSubvlans(o["lan-subvlans"], d, "lan_subvlans", sv)); err != nil {
		if !fortiAPIPatch(o["lan-subvlans"]) {
			return fmt.Errorf("Error reading lan_subvlans: %v", err)
		}
	}

	if err = d.Set("arp_inspection", flattenSwitchVlanArpInspection(o["arp-inspection"], d, "arp_inspection", sv)); err != nil {
		if !fortiAPIPatch(o["arp-inspection"]) {
			return fmt.Errorf("Error reading arp_inspection: %v", err)
		}
	}

	if err = d.Set("igmp_snooping_proxy", flattenSwitchVlanIgmpSnoopingProxy(o["igmp-snooping-proxy"], d, "igmp_snooping_proxy", sv)); err != nil {
		if !fortiAPIPatch(o["igmp-snooping-proxy"]) {
			return fmt.Errorf("Error reading igmp_snooping_proxy: %v", err)
		}
	}

	if err = d.Set("policer", flattenSwitchVlanPolicer(o["policer"], d, "policer", sv)); err != nil {
		if !fortiAPIPatch(o["policer"]) {
			return fmt.Errorf("Error reading policer: %v", err)
		}
	}

	if err = d.Set("dhcp6_snooping", flattenSwitchVlanDhcp6Snooping(o["dhcp6-snooping"], d, "dhcp6_snooping", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp6-snooping"]) {
			return fmt.Errorf("Error reading dhcp6_snooping: %v", err)
		}
	}

	if err = d.Set("lan_segment_primary_vlan", flattenSwitchVlanLanSegmentPrimaryVlan(o["lan-segment-primary-vlan"], d, "lan_segment_primary_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["lan-segment-primary-vlan"]) {
			return fmt.Errorf("Error reading lan_segment_primary_vlan: %v", err)
		}
	}

	if err = d.Set("learning", flattenSwitchVlanLearning(o["learning"], d, "learning", sv)); err != nil {
		if !fortiAPIPatch(o["learning"]) {
			return fmt.Errorf("Error reading learning: %v", err)
		}
	}

	if err = d.Set("fswid", flattenSwitchVlanId(o["id"], d, "fswid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fswid: %v", err)
		}
	}

	if err = d.Set("igmp_snooping_querier", flattenSwitchVlanIgmpSnoopingQuerier(o["igmp-snooping-querier"], d, "igmp_snooping_querier", sv)); err != nil {
		if !fortiAPIPatch(o["igmp-snooping-querier"]) {
			return fmt.Errorf("Error reading igmp_snooping_querier: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dhcp_snooping_static_client", flattenSwitchVlanDhcpSnoopingStaticClient(o["dhcp-snooping-static-client"], d, "dhcp_snooping_static_client", sv)); err != nil {
			if !fortiAPIPatch(o["dhcp-snooping-static-client"]) {
				return fmt.Errorf("Error reading dhcp_snooping_static_client: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dhcp_snooping_static_client"); ok {
			if err = d.Set("dhcp_snooping_static_client", flattenSwitchVlanDhcpSnoopingStaticClient(o["dhcp-snooping-static-client"], d, "dhcp_snooping_static_client", sv)); err != nil {
				if !fortiAPIPatch(o["dhcp-snooping-static-client"]) {
					return fmt.Errorf("Error reading dhcp_snooping_static_client: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("member_by_mac", flattenSwitchVlanMemberByMac(o["member-by-mac"], d, "member_by_mac", sv)); err != nil {
			if !fortiAPIPatch(o["member-by-mac"]) {
				return fmt.Errorf("Error reading member_by_mac: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("member_by_mac"); ok {
			if err = d.Set("member_by_mac", flattenSwitchVlanMemberByMac(o["member-by-mac"], d, "member_by_mac", sv)); err != nil {
				if !fortiAPIPatch(o["member-by-mac"]) {
					return fmt.Errorf("Error reading member_by_mac: %v", err)
				}
			}
		}
	}

	if err = d.Set("lan_segment", flattenSwitchVlanLanSegment(o["lan-segment"], d, "lan_segment", sv)); err != nil {
		if !fortiAPIPatch(o["lan-segment"]) {
			return fmt.Errorf("Error reading lan_segment: %v", err)
		}
	}

	if err = d.Set("primary_vlan", flattenSwitchVlanPrimaryVlan(o["primary-vlan"], d, "primary_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["primary-vlan"]) {
			return fmt.Errorf("Error reading primary_vlan: %v", err)
		}
	}

	if err = d.Set("igmp_snooping_querier_version", flattenSwitchVlanIgmpSnoopingQuerierVersion(o["igmp-snooping-querier-version"], d, "igmp_snooping_querier_version", sv)); err != nil {
		if !fortiAPIPatch(o["igmp-snooping-querier-version"]) {
			return fmt.Errorf("Error reading igmp_snooping_querier_version: %v", err)
		}
	}

	if err = d.Set("cos_queue", flattenSwitchVlanCosQueue(o["cos-queue"], d, "cos_queue", sv)); err != nil {
		if !fortiAPIPatch(o["cos-queue"]) {
			return fmt.Errorf("Error reading cos_queue: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("member_by_ipv6", flattenSwitchVlanMemberByIpv6(o["member-by-ipv6"], d, "member_by_ipv6", sv)); err != nil {
			if !fortiAPIPatch(o["member-by-ipv6"]) {
				return fmt.Errorf("Error reading member_by_ipv6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("member_by_ipv6"); ok {
			if err = d.Set("member_by_ipv6", flattenSwitchVlanMemberByIpv6(o["member-by-ipv6"], d, "member_by_ipv6", sv)); err != nil {
				if !fortiAPIPatch(o["member-by-ipv6"]) {
					return fmt.Errorf("Error reading member_by_ipv6: %v", err)
				}
			}
		}
	}

	if err = d.Set("mld_snooping_proxy", flattenSwitchVlanMldSnoopingProxy(o["mld-snooping-proxy"], d, "mld_snooping_proxy", sv)); err != nil {
		if !fortiAPIPatch(o["mld-snooping-proxy"]) {
			return fmt.Errorf("Error reading mld_snooping_proxy: %v", err)
		}
	}

	if err = d.Set("mld_snooping_fast_leave", flattenSwitchVlanMldSnoopingFastLeave(o["mld-snooping-fast-leave"], d, "mld_snooping_fast_leave", sv)); err != nil {
		if !fortiAPIPatch(o["mld-snooping-fast-leave"]) {
			return fmt.Errorf("Error reading mld_snooping_fast_leave: %v", err)
		}
	}

	if err = d.Set("igmp_snooping", flattenSwitchVlanIgmpSnooping(o["igmp-snooping"], d, "igmp_snooping", sv)); err != nil {
		if !fortiAPIPatch(o["igmp-snooping"]) {
			return fmt.Errorf("Error reading igmp_snooping: %v", err)
		}
	}

	if err = d.Set("mld_snooping_querier_addr", flattenSwitchVlanMldSnoopingQuerierAddr(o["mld-snooping-querier-addr"], d, "mld_snooping_querier_addr", sv)); err != nil {
		if !fortiAPIPatch(o["mld-snooping-querier-addr"]) {
			return fmt.Errorf("Error reading mld_snooping_querier_addr: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dhcp_server_access_list", flattenSwitchVlanDhcpServerAccessList(o["dhcp-server-access-list"], d, "dhcp_server_access_list", sv)); err != nil {
			if !fortiAPIPatch(o["dhcp-server-access-list"]) {
				return fmt.Errorf("Error reading dhcp_server_access_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dhcp_server_access_list"); ok {
			if err = d.Set("dhcp_server_access_list", flattenSwitchVlanDhcpServerAccessList(o["dhcp-server-access-list"], d, "dhcp_server_access_list", sv)); err != nil {
				if !fortiAPIPatch(o["dhcp-server-access-list"]) {
					return fmt.Errorf("Error reading dhcp_server_access_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("learning_limit", flattenSwitchVlanLearningLimit(o["learning-limit"], d, "learning_limit", sv)); err != nil {
		if !fortiAPIPatch(o["learning-limit"]) {
			return fmt.Errorf("Error reading learning_limit: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mrouter_ports", flattenSwitchVlanMrouterPorts(o["mrouter-ports"], d, "mrouter_ports", sv)); err != nil {
			if !fortiAPIPatch(o["mrouter-ports"]) {
				return fmt.Errorf("Error reading mrouter_ports: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mrouter_ports"); ok {
			if err = d.Set("mrouter_ports", flattenSwitchVlanMrouterPorts(o["mrouter-ports"], d, "mrouter_ports", sv)); err != nil {
				if !fortiAPIPatch(o["mrouter-ports"]) {
					return fmt.Errorf("Error reading mrouter_ports: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mld_snooping_static_group", flattenSwitchVlanMldSnoopingStaticGroup(o["mld-snooping-static-group"], d, "mld_snooping_static_group", sv)); err != nil {
			if !fortiAPIPatch(o["mld-snooping-static-group"]) {
				return fmt.Errorf("Error reading mld_snooping_static_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mld_snooping_static_group"); ok {
			if err = d.Set("mld_snooping_static_group", flattenSwitchVlanMldSnoopingStaticGroup(o["mld-snooping-static-group"], d, "mld_snooping_static_group", sv)); err != nil {
				if !fortiAPIPatch(o["mld-snooping-static-group"]) {
					return fmt.Errorf("Error reading mld_snooping_static_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("description", flattenSwitchVlanDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("igmp_snooping_querier_addr", flattenSwitchVlanIgmpSnoopingQuerierAddr(o["igmp-snooping-querier-addr"], d, "igmp_snooping_querier_addr", sv)); err != nil {
		if !fortiAPIPatch(o["igmp-snooping-querier-addr"]) {
			return fmt.Errorf("Error reading igmp_snooping_querier_addr: %v", err)
		}
	}

	if err = d.Set("rspan_mode", flattenSwitchVlanRspanMode(o["rspan-mode"], d, "rspan_mode", sv)); err != nil {
		if !fortiAPIPatch(o["rspan-mode"]) {
			return fmt.Errorf("Error reading rspan_mode: %v", err)
		}
	}

	if err = d.Set("private_vlan_type", flattenSwitchVlanPrivateVlanType(o["private-vlan-type"], d, "private_vlan_type", sv)); err != nil {
		if !fortiAPIPatch(o["private-vlan-type"]) {
			return fmt.Errorf("Error reading private_vlan_type: %v", err)
		}
	}

	if err = d.Set("mld_snooping_querier", flattenSwitchVlanMldSnoopingQuerier(o["mld-snooping-querier"], d, "mld_snooping_querier", sv)); err != nil {
		if !fortiAPIPatch(o["mld-snooping-querier"]) {
			return fmt.Errorf("Error reading mld_snooping_querier: %v", err)
		}
	}

	if err = d.Set("igmp_snooping_fast_leave", flattenSwitchVlanIgmpSnoopingFastLeave(o["igmp-snooping-fast-leave"], d, "igmp_snooping_fast_leave", sv)); err != nil {
		if !fortiAPIPatch(o["igmp-snooping-fast-leave"]) {
			return fmt.Errorf("Error reading igmp_snooping_fast_leave: %v", err)
		}
	}

	if err = d.Set("dhcp_snooping", flattenSwitchVlanDhcpSnooping(o["dhcp-snooping"], d, "dhcp_snooping", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-snooping"]) {
			return fmt.Errorf("Error reading dhcp_snooping: %v", err)
		}
	}

	if err = d.Set("lan_segment_type", flattenSwitchVlanLanSegmentType(o["lan-segment-type"], d, "lan_segment_type", sv)); err != nil {
		if !fortiAPIPatch(o["lan-segment-type"]) {
			return fmt.Errorf("Error reading lan_segment_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("igmp_snooping_static_group", flattenSwitchVlanIgmpSnoopingStaticGroup(o["igmp-snooping-static-group"], d, "igmp_snooping_static_group", sv)); err != nil {
			if !fortiAPIPatch(o["igmp-snooping-static-group"]) {
				return fmt.Errorf("Error reading igmp_snooping_static_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("igmp_snooping_static_group"); ok {
			if err = d.Set("igmp_snooping_static_group", flattenSwitchVlanIgmpSnoopingStaticGroup(o["igmp-snooping-static-group"], d, "igmp_snooping_static_group", sv)); err != nil {
				if !fortiAPIPatch(o["igmp-snooping-static-group"]) {
					return fmt.Errorf("Error reading igmp_snooping_static_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("dhcp_snooping_verify_mac", flattenSwitchVlanDhcpSnoopingVerifyMac(o["dhcp-snooping-verify-mac"], d, "dhcp_snooping_verify_mac", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-snooping-verify-mac"]) {
			return fmt.Errorf("Error reading dhcp_snooping_verify_mac: %v", err)
		}
	}

	if err = d.Set("isolated_vlan", flattenSwitchVlanIsolatedVlan(o["isolated-vlan"], d, "isolated_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["isolated-vlan"]) {
			return fmt.Errorf("Error reading isolated_vlan: %v", err)
		}
	}

	if err = d.Set("dhcp_snooping_option82", flattenSwitchVlanDhcpSnoopingOption82(o["dhcp-snooping-option82"], d, "dhcp_snooping_option82", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-snooping-option82"]) {
			return fmt.Errorf("Error reading dhcp_snooping_option82: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("member_by_ipv4", flattenSwitchVlanMemberByIpv4(o["member-by-ipv4"], d, "member_by_ipv4", sv)); err != nil {
			if !fortiAPIPatch(o["member-by-ipv4"]) {
				return fmt.Errorf("Error reading member_by_ipv4: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("member_by_ipv4"); ok {
			if err = d.Set("member_by_ipv4", flattenSwitchVlanMemberByIpv4(o["member-by-ipv4"], d, "member_by_ipv4", sv)); err != nil {
				if !fortiAPIPatch(o["member-by-ipv4"]) {
					return fmt.Errorf("Error reading member_by_ipv4: %v", err)
				}
			}
		}
	}

	if err = d.Set("private_vlan", flattenSwitchVlanPrivateVlan(o["private-vlan"], d, "private_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["private-vlan"]) {
			return fmt.Errorf("Error reading private_vlan: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("member_by_proto", flattenSwitchVlanMemberByProto(o["member-by-proto"], d, "member_by_proto", sv)); err != nil {
			if !fortiAPIPatch(o["member-by-proto"]) {
				return fmt.Errorf("Error reading member_by_proto: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("member_by_proto"); ok {
			if err = d.Set("member_by_proto", flattenSwitchVlanMemberByProto(o["member-by-proto"], d, "member_by_proto", sv)); err != nil {
				if !fortiAPIPatch(o["member-by-proto"]) {
					return fmt.Errorf("Error reading member_by_proto: %v", err)
				}
			}
		}
	}

	if err = d.Set("access_vlan", flattenSwitchVlanAccessVlan(o["access-vlan"], d, "access_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["access-vlan"]) {
			return fmt.Errorf("Error reading access_vlan: %v", err)
		}
	}

	if err = d.Set("mld_snooping", flattenSwitchVlanMldSnooping(o["mld-snooping"], d, "mld_snooping", sv)); err != nil {
		if !fortiAPIPatch(o["mld-snooping"]) {
			return fmt.Errorf("Error reading mld_snooping: %v", err)
		}
	}

	return nil
}

func flattenSwitchVlanFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchVlanCommunityVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanLanSubvlans(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanArpInspection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanIgmpSnoopingProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanPolicer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanDhcp6Snooping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanLanSegmentPrimaryVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanLearning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanIgmpSnoopingQuerier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanDhcpSnoopingStaticClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_addr"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip-addr"], _ = expandSwitchVlanDhcpSnoopingStaticClientIpAddr(d, i["ip_addr"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mac-addr"], _ = expandSwitchVlanDhcpSnoopingStaticClientMacAddr(d, i["mac_addr"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSwitchVlanDhcpSnoopingStaticClientName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_interface"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["switch-interface"], _ = expandSwitchVlanDhcpSnoopingStaticClientSwitchInterface(d, i["switch_interface"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchVlanDhcpSnoopingStaticClientIpAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanDhcpSnoopingStaticClientMacAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanDhcpSnoopingStaticClientName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanDhcpSnoopingStaticClientSwitchInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMemberByMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mac"], _ = expandSwitchVlanMemberByMacMac(d, i["mac"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSwitchVlanMemberByMacId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandSwitchVlanMemberByMacDescription(d, i["description"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchVlanMemberByMacMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMemberByMacId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMemberByMacDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanLanSegment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanPrimaryVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanIgmpSnoopingQuerierVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanCosQueue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMemberByIpv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix"], _ = expandSwitchVlanMemberByIpv6Prefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSwitchVlanMemberByIpv6Id(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandSwitchVlanMemberByIpv6Description(d, i["description"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchVlanMemberByIpv6Prefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMemberByIpv6Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMemberByIpv6Description(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMldSnoopingProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMldSnoopingFastLeave(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanIgmpSnooping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMldSnoopingQuerierAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanDhcpServerAccessList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["server-ip"], _ = expandSwitchVlanDhcpServerAccessListServerIp(d, i["server_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSwitchVlanDhcpServerAccessListName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_ip6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["server-ip6"], _ = expandSwitchVlanDhcpServerAccessListServerIp6(d, i["server_ip6"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchVlanDhcpServerAccessListServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanDhcpServerAccessListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanDhcpServerAccessListServerIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanLearningLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMrouterPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["member-name"], _ = expandSwitchVlanMrouterPortsMemberName(d, i["member_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchVlanMrouterPortsMemberName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMldSnoopingStaticGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_reports"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ignore-reports"], _ = expandSwitchVlanMldSnoopingStaticGroupIgnoreReports(d, i["ignore_reports"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcast_addr"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mcast-addr"], _ = expandSwitchVlanMldSnoopingStaticGroupMcastAddr(d, i["mcast_addr"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSwitchVlanMldSnoopingStaticGroupName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["members"], _ = expandSwitchVlanMldSnoopingStaticGroupMembers(d, i["members"], pre_append, sv)
		} else {
			tmp["members"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchVlanMldSnoopingStaticGroupIgnoreReports(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMldSnoopingStaticGroupMcastAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMldSnoopingStaticGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMldSnoopingStaticGroupMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["member-name"], _ = expandSwitchVlanMldSnoopingStaticGroupMembersMemberName(d, i["member_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchVlanMldSnoopingStaticGroupMembersMemberName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanIgmpSnoopingQuerierAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanRspanMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanPrivateVlanType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMldSnoopingQuerier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanIgmpSnoopingFastLeave(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanDhcpSnooping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanLanSegmentType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanIgmpSnoopingStaticGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_reports"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ignore-reports"], _ = expandSwitchVlanIgmpSnoopingStaticGroupIgnoreReports(d, i["ignore_reports"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcast_addr"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mcast-addr"], _ = expandSwitchVlanIgmpSnoopingStaticGroupMcastAddr(d, i["mcast_addr"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSwitchVlanIgmpSnoopingStaticGroupName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["members"], _ = expandSwitchVlanIgmpSnoopingStaticGroupMembers(d, i["members"], pre_append, sv)
		} else {
			tmp["members"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchVlanIgmpSnoopingStaticGroupIgnoreReports(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanIgmpSnoopingStaticGroupMcastAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanIgmpSnoopingStaticGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanIgmpSnoopingStaticGroupMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["member-name"], _ = expandSwitchVlanIgmpSnoopingStaticGroupMembersMemberName(d, i["member_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchVlanIgmpSnoopingStaticGroupMembersMemberName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanDhcpSnoopingVerifyMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanIsolatedVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanDhcpSnoopingOption82(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMemberByIpv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandSwitchVlanMemberByIpv4Description(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSwitchVlanMemberByIpv4Id(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["address"], _ = expandSwitchVlanMemberByIpv4Address(d, i["address"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchVlanMemberByIpv4Description(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMemberByIpv4Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMemberByIpv4Address(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanPrivateVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMemberByProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["protocol"], _ = expandSwitchVlanMemberByProtoProtocol(d, i["protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "frametypes"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["frametypes"], _ = expandSwitchVlanMemberByProtoFrametypes(d, i["frametypes"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSwitchVlanMemberByProtoId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandSwitchVlanMemberByProtoDescription(d, i["description"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchVlanMemberByProtoProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMemberByProtoFrametypes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMemberByProtoId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMemberByProtoDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanAccessVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanMldSnooping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchVlan(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("community_vlans"); ok {

		t, err := expandSwitchVlanCommunityVlans(d, v, "community_vlans", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["community-vlans"] = t
		}
	}

	if v, ok := d.GetOk("lan_subvlans"); ok {

		t, err := expandSwitchVlanLanSubvlans(d, v, "lan_subvlans", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-subvlans"] = t
		}
	}

	if v, ok := d.GetOk("arp_inspection"); ok {

		t, err := expandSwitchVlanArpInspection(d, v, "arp_inspection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-inspection"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping_proxy"); ok {

		t, err := expandSwitchVlanIgmpSnoopingProxy(d, v, "igmp_snooping_proxy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping-proxy"] = t
		}
	}

	if v, ok := d.GetOk("policer"); ok {

		t, err := expandSwitchVlanPolicer(d, v, "policer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policer"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_snooping"); ok {

		t, err := expandSwitchVlanDhcp6Snooping(d, v, "dhcp6_snooping", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-snooping"] = t
		}
	}

	if v, ok := d.GetOk("lan_segment_primary_vlan"); ok {

		t, err := expandSwitchVlanLanSegmentPrimaryVlan(d, v, "lan_segment_primary_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-segment-primary-vlan"] = t
		}
	}

	if v, ok := d.GetOk("learning"); ok {

		t, err := expandSwitchVlanLearning(d, v, "learning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learning"] = t
		}
	}

	if v, ok := d.GetOk("fswid"); ok {

		t, err := expandSwitchVlanId(d, v, "fswid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping_querier"); ok {

		t, err := expandSwitchVlanIgmpSnoopingQuerier(d, v, "igmp_snooping_querier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping-querier"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snooping_static_client"); ok || d.HasChange("dhcp_snooping_static_client") {

		t, err := expandSwitchVlanDhcpSnoopingStaticClient(d, v, "dhcp_snooping_static_client", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snooping-static-client"] = t
		}
	}

	if v, ok := d.GetOk("member_by_mac"); ok || d.HasChange("member_by_mac") {

		t, err := expandSwitchVlanMemberByMac(d, v, "member_by_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-by-mac"] = t
		}
	}

	if v, ok := d.GetOk("lan_segment"); ok {

		t, err := expandSwitchVlanLanSegment(d, v, "lan_segment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-segment"] = t
		}
	}

	if v, ok := d.GetOk("primary_vlan"); ok {

		t, err := expandSwitchVlanPrimaryVlan(d, v, "primary_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-vlan"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping_querier_version"); ok {

		t, err := expandSwitchVlanIgmpSnoopingQuerierVersion(d, v, "igmp_snooping_querier_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping-querier-version"] = t
		}
	}

	if v, ok := d.GetOkExists("cos_queue"); ok {

		t, err := expandSwitchVlanCosQueue(d, v, "cos_queue", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos-queue"] = t
		}
	}

	if v, ok := d.GetOk("member_by_ipv6"); ok || d.HasChange("member_by_ipv6") {

		t, err := expandSwitchVlanMemberByIpv6(d, v, "member_by_ipv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-by-ipv6"] = t
		}
	}

	if v, ok := d.GetOk("mld_snooping_proxy"); ok {

		t, err := expandSwitchVlanMldSnoopingProxy(d, v, "mld_snooping_proxy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mld-snooping-proxy"] = t
		}
	}

	if v, ok := d.GetOk("mld_snooping_fast_leave"); ok {

		t, err := expandSwitchVlanMldSnoopingFastLeave(d, v, "mld_snooping_fast_leave", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mld-snooping-fast-leave"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping"); ok {

		t, err := expandSwitchVlanIgmpSnooping(d, v, "igmp_snooping", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("mld_snooping_querier_addr"); ok {

		t, err := expandSwitchVlanMldSnoopingQuerierAddr(d, v, "mld_snooping_querier_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mld-snooping-querier-addr"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_server_access_list"); ok || d.HasChange("dhcp_server_access_list") {

		t, err := expandSwitchVlanDhcpServerAccessList(d, v, "dhcp_server_access_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-server-access-list"] = t
		}
	}

	if v, ok := d.GetOk("learning_limit"); ok {

		t, err := expandSwitchVlanLearningLimit(d, v, "learning_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learning-limit"] = t
		}
	}

	if v, ok := d.GetOk("mrouter_ports"); ok || d.HasChange("mrouter_ports") {

		t, err := expandSwitchVlanMrouterPorts(d, v, "mrouter_ports", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mrouter-ports"] = t
		}
	}

	if v, ok := d.GetOk("mld_snooping_static_group"); ok || d.HasChange("mld_snooping_static_group") {

		t, err := expandSwitchVlanMldSnoopingStaticGroup(d, v, "mld_snooping_static_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mld-snooping-static-group"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchVlanDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping_querier_addr"); ok {

		t, err := expandSwitchVlanIgmpSnoopingQuerierAddr(d, v, "igmp_snooping_querier_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping-querier-addr"] = t
		}
	}

	if v, ok := d.GetOk("rspan_mode"); ok {

		t, err := expandSwitchVlanRspanMode(d, v, "rspan_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rspan-mode"] = t
		}
	}

	if v, ok := d.GetOk("private_vlan_type"); ok {

		t, err := expandSwitchVlanPrivateVlanType(d, v, "private_vlan_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-vlan-type"] = t
		}
	}

	if v, ok := d.GetOk("mld_snooping_querier"); ok {

		t, err := expandSwitchVlanMldSnoopingQuerier(d, v, "mld_snooping_querier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mld-snooping-querier"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping_fast_leave"); ok {

		t, err := expandSwitchVlanIgmpSnoopingFastLeave(d, v, "igmp_snooping_fast_leave", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping-fast-leave"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snooping"); ok {

		t, err := expandSwitchVlanDhcpSnooping(d, v, "dhcp_snooping", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("lan_segment_type"); ok {

		t, err := expandSwitchVlanLanSegmentType(d, v, "lan_segment_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-segment-type"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping_static_group"); ok || d.HasChange("igmp_snooping_static_group") {

		t, err := expandSwitchVlanIgmpSnoopingStaticGroup(d, v, "igmp_snooping_static_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping-static-group"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snooping_verify_mac"); ok {

		t, err := expandSwitchVlanDhcpSnoopingVerifyMac(d, v, "dhcp_snooping_verify_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snooping-verify-mac"] = t
		}
	}

	if v, ok := d.GetOk("isolated_vlan"); ok {

		t, err := expandSwitchVlanIsolatedVlan(d, v, "isolated_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isolated-vlan"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snooping_option82"); ok {

		t, err := expandSwitchVlanDhcpSnoopingOption82(d, v, "dhcp_snooping_option82", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snooping-option82"] = t
		}
	}

	if v, ok := d.GetOk("member_by_ipv4"); ok || d.HasChange("member_by_ipv4") {

		t, err := expandSwitchVlanMemberByIpv4(d, v, "member_by_ipv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-by-ipv4"] = t
		}
	}

	if v, ok := d.GetOk("private_vlan"); ok {

		t, err := expandSwitchVlanPrivateVlan(d, v, "private_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-vlan"] = t
		}
	}

	if v, ok := d.GetOk("member_by_proto"); ok || d.HasChange("member_by_proto") {

		t, err := expandSwitchVlanMemberByProto(d, v, "member_by_proto", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-by-proto"] = t
		}
	}

	if v, ok := d.GetOk("access_vlan"); ok {

		t, err := expandSwitchVlanAccessVlan(d, v, "access_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-vlan"] = t
		}
	}

	if v, ok := d.GetOk("mld_snooping"); ok {

		t, err := expandSwitchVlanMldSnooping(d, v, "mld_snooping", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mld-snooping"] = t
		}
	}

	return &obj, nil
}
