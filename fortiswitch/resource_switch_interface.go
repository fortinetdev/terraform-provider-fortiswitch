// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Usable interfaces (trunks and ports).

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchInterfaceCreate,
		Read:   resourceSwitchInterfaceRead,
		Update: resourceSwitchInterfaceUpdate,
		Delete: resourceSwitchInterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"allow_arp_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortilink_l3_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"igmp_snooping_flood_reports": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_fortilink": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_dot1p_map": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"discard_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qos_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"arp_inspection_trust": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"private_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sub_vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"allowed_sub_vlans": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sample_direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_snoop_option82_trust": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"edge_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stp_bpdu_guard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_snoop_option82_override": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"circuit_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 254),
							Optional:     true,
							Computed:     true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"remote_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 254),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"ptp_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loop_guard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qnq": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan_mapping": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"new_s_vlan": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 4094),
										Optional:     true,
										Computed:     true,
									},
									"match_c_vlan": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 4094),
										Optional:     true,
										Computed:     true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"description": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"stp_qnq_admin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"native_c_vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"remove_inner": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan_mapping_miss_drop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untagged_s_vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"add_inner": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"allowed_c_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"edge_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"s_tag_priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"allowed_vlans": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_fortilink_packet_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 300),
				Optional:     true,
				Computed:     true,
			},
			"primary_vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"match_s_vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"new_s_vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_c_vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"raguard": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_list": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"raguard_policy": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"stp_bpdu_guard_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"stp_loop_protection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"learning_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"learning_limit_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loop_guard_mac_move_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_mac_event": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_snoop_learning_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmp_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcp_snoop_learning_limit_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_port_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"stp_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_traffic_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sticky_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"packet_sampler": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loop_guard_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rpvst_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_mac_binding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mld_snooping_flood_reports": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sflow_counter_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"native_vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vlan_mapping_miss_drop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"untagged_vlans": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private_vlan_port_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcp_snooping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter_sub_vlans": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stp_root_guard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_ip_dscp_map": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"port_security": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"macsec_pae_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_fail_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"macsec_profile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"allow_mac_move_to": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_fail_vlanid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"port_security_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authserver_timeout_tagged": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mab_eapol_request": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"dacl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authserver_timeout_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"allow_mac_move": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"guest_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"guest_auth_delay": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"client_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 20),
							Optional:     true,
							Computed:     true,
						},
						"auth_order": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"framevid_apply": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"eap_auto_untagged_vlans": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authserver_timeout_tagged_lldp_voice_vlanid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"authserver_timeout_tagged_vlanid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mac_auth_bypass": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"radius_timeout_overwrite": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authserver_timeout_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"open_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authserver_timeout_vlanid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"quarantine_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"guest_vlanid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"eap_egress_tagged": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"eap_passthru": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"default_cos": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mcast_snooping_flood_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_tpid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"packet_sample_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ptp_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchInterface resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchInterface(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchInterface resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchInterface")
	}

	return resourceSwitchInterfaceRead(d, m)
}

func resourceSwitchInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchInterface resource while getting object: %v", err)
	}

	// Create the request body in the required format
	data := map[string]interface{}{
		"json": obj, // Assuming `obj` contains the request body data
	}

	// Pass the `data` to the `UpdateSwitchInterface` function
	o, err := c.UpdateSwitchInterface(&data, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchInterface")
	}

	return resourceSwitchInterfaceRead(d, m)
}

func resourceSwitchInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchInterface(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchInterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchInterface(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchInterface(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchInterface resource from API: %v", err)
	}
	return nil
}

func flattenSwitchInterfaceAllowArpMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceFortilinkL3Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceIgmpSnoopingFloodReports(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceInterfaceMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceAutoDiscoveryFortilink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceTrustDot1PMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceDiscardMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceQosPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceArpInspectionTrust(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceSecurityGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchInterfaceSecurityGroupsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchInterfaceSecurityGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePrivateVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceSubVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceAllowedSubVlans(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceSampleDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceDhcpSnoopOption82Trust(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceEdgePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceStpBpduGuard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceDhcpSnoopOption82Override(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := i["circuit-id"]; ok {

			tmp["circuit_id"] = flattenSwitchInterfaceDhcpSnoopOption82OverrideCircuitId(i["circuit-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSwitchInterfaceDhcpSnoopOption82OverrideId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := i["remote-id"]; ok {

			tmp["remote_id"] = flattenSwitchInterfaceDhcpSnoopOption82OverrideRemoteId(i["remote-id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchInterfaceDhcpSnoopOption82OverrideCircuitId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceDhcpSnoopOption82OverrideId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceDhcpSnoopOption82OverrideRemoteId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePtpStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceLoopGuard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceQnq(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {

		result["status"] = flattenSwitchInterfaceQnqStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vlan_mapping"
	if _, ok := i["vlan-mapping"]; ok {

		result["vlan_mapping"] = flattenSwitchInterfaceQnqVlanMapping(i["vlan-mapping"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "stp_qnq_admin"
	if _, ok := i["stp-qnq-admin"]; ok {

		result["stp_qnq_admin"] = flattenSwitchInterfaceQnqStpQnqAdmin(i["stp-qnq-admin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "native_c_vlan"
	if _, ok := i["native-c-vlan"]; ok {

		result["native_c_vlan"] = flattenSwitchInterfaceQnqNativeCVlan(i["native-c-vlan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "remove_inner"
	if _, ok := i["remove-inner"]; ok {

		result["remove_inner"] = flattenSwitchInterfaceQnqRemoveInner(i["remove-inner"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {

		result["priority"] = flattenSwitchInterfaceQnqPriority(i["priority"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vlan_mapping_miss_drop"
	if _, ok := i["vlan-mapping-miss-drop"]; ok {

		result["vlan_mapping_miss_drop"] = flattenSwitchInterfaceQnqVlanMappingMissDrop(i["vlan-mapping-miss-drop"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "untagged_s_vlan"
	if _, ok := i["untagged-s-vlan"]; ok {

		result["untagged_s_vlan"] = flattenSwitchInterfaceQnqUntaggedSVlan(i["untagged-s-vlan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "add_inner"
	if _, ok := i["add-inner"]; ok {

		result["add_inner"] = flattenSwitchInterfaceQnqAddInner(i["add-inner"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "allowed_c_vlan"
	if _, ok := i["allowed-c-vlan"]; ok {

		result["allowed_c_vlan"] = flattenSwitchInterfaceQnqAllowedCVlan(i["allowed-c-vlan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "edge_type"
	if _, ok := i["edge-type"]; ok {

		result["edge_type"] = flattenSwitchInterfaceQnqEdgeType(i["edge-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "s_tag_priority"
	if _, ok := i["s-tag-priority"]; ok {

		result["s_tag_priority"] = flattenSwitchInterfaceQnqSTagPriority(i["s-tag-priority"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchInterfaceQnqStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceQnqVlanMapping(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "new_s_vlan"
		if _, ok := i["new-s-vlan"]; ok {

			tmp["new_s_vlan"] = flattenSwitchInterfaceQnqVlanMappingNewSVlan(i["new-s-vlan"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_c_vlan"
		if _, ok := i["match-c-vlan"]; ok {

			tmp["match_c_vlan"] = flattenSwitchInterfaceQnqVlanMappingMatchCVlan(i["match-c-vlan"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSwitchInterfaceQnqVlanMappingId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchInterfaceQnqVlanMappingDescription(i["description"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchInterfaceQnqVlanMappingNewSVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceQnqVlanMappingMatchCVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceQnqVlanMappingId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceQnqVlanMappingDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceQnqStpQnqAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceQnqNativeCVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v == "" || v == "none" {
		return nil
	}
	return v
}

func flattenSwitchInterfaceQnqRemoveInner(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceQnqPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceQnqVlanMappingMissDrop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceQnqUntaggedSVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceQnqAddInner(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v == "" || v == "none" {
		return nil
	}
	return v
}

func flattenSwitchInterfaceQnqAllowedCVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceQnqEdgeType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceQnqSTagPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceAllowedVlans(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceNac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceAutoDiscoveryFortilinkPacketInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePrimaryVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceVlanMapping(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {

			tmp["direction"] = flattenSwitchInterfaceVlanMappingDirection(i["direction"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchInterfaceVlanMappingDescription(i["description"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_s_vlan"
		if _, ok := i["match-s-vlan"]; ok {

			tmp["match_s_vlan"] = flattenSwitchInterfaceVlanMappingMatchSVlan(i["match-s-vlan"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "new_s_vlan"
		if _, ok := i["new-s-vlan"]; ok {

			tmp["new_s_vlan"] = flattenSwitchInterfaceVlanMappingNewSVlan(i["new-s-vlan"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenSwitchInterfaceVlanMappingAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_c_vlan"
		if _, ok := i["match-c-vlan"]; ok {

			tmp["match_c_vlan"] = flattenSwitchInterfaceVlanMappingMatchCVlan(i["match-c-vlan"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSwitchInterfaceVlanMappingId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchInterfaceVlanMappingDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceVlanMappingDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceVlanMappingMatchSVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceVlanMappingNewSVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceVlanMappingAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceVlanMappingMatchCVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceVlanMappingId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceRaguard(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_list"
		if _, ok := i["vlan-list"]; ok {

			tmp["vlan_list"] = flattenSwitchInterfaceRaguardVlanList(i["vlan-list"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSwitchInterfaceRaguardId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "raguard_policy"
		if _, ok := i["raguard-policy"]; ok {

			tmp["raguard_policy"] = flattenSwitchInterfaceRaguardRaguardPolicy(i["raguard-policy"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchInterfaceRaguardVlanList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceRaguardId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceRaguardRaguardPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceStpBpduGuardTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceStpLoopProtection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceLearningLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceLearningLimitAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceLoopGuardMacMoveThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceLogMacEvent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceDhcpSnoopLearningLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceSnmpIndex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceDhcpSnoopLearningLimitCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceSwitchPortMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceStpState(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceVlanTrafficType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceStickyMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePacketSampler(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceLoopGuardTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceRpvstPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceIpMacBinding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceMldSnoopingFloodReports(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceSflowCounterInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceNativeVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceVlanMappingMissDrop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceUntaggedVlans(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePrivateVlanPortType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceDhcpSnooping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceFilterSubVlans(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceStpRootGuard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceTrustIpDscpMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurity(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "macsec_pae_mode"
	if _, ok := i["macsec-pae-mode"]; ok {

		result["macsec_pae_mode"] = flattenSwitchInterfacePortSecurityMacsecPaeMode(i["macsec-pae-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auth_fail_vlan"
	if _, ok := i["auth-fail-vlan"]; ok {

		result["auth_fail_vlan"] = flattenSwitchInterfacePortSecurityAuthFailVlan(i["auth-fail-vlan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "macsec_profile"
	if _, ok := i["macsec-profile"]; ok {

		result["macsec_profile"] = flattenSwitchInterfacePortSecurityMacsecProfile(i["macsec-profile"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "allow_mac_move_to"
	if _, ok := i["allow-mac-move-to"]; ok {

		result["allow_mac_move_to"] = flattenSwitchInterfacePortSecurityAllowMacMoveTo(i["allow-mac-move-to"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auth_fail_vlanid"
	if _, ok := i["auth-fail-vlanid"]; ok {

		result["auth_fail_vlanid"] = flattenSwitchInterfacePortSecurityAuthFailVlanid(i["auth-fail-vlanid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port_security_mode"
	if _, ok := i["port-security-mode"]; ok {

		result["port_security_mode"] = flattenSwitchInterfacePortSecurityPortSecurityMode(i["port-security-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "authserver_timeout_tagged"
	if _, ok := i["authserver-timeout-tagged"]; ok {

		result["authserver_timeout_tagged"] = flattenSwitchInterfacePortSecurityAuthserverTimeoutTagged(i["authserver-timeout-tagged"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mab_eapol_request"
	if _, ok := i["mab-eapol-request"]; ok {

		result["mab_eapol_request"] = flattenSwitchInterfacePortSecurityMabEapolRequest(i["mab-eapol-request"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dacl"
	if _, ok := i["dacl"]; ok {

		result["dacl"] = flattenSwitchInterfacePortSecurityDacl(i["dacl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "authserver_timeout_vlan"
	if _, ok := i["authserver-timeout-vlan"]; ok {

		result["authserver_timeout_vlan"] = flattenSwitchInterfacePortSecurityAuthserverTimeoutVlan(i["authserver-timeout-vlan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "allow_mac_move"
	if _, ok := i["allow-mac-move"]; ok {

		result["allow_mac_move"] = flattenSwitchInterfacePortSecurityAllowMacMove(i["allow-mac-move"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "guest_vlan"
	if _, ok := i["guest-vlan"]; ok {

		result["guest_vlan"] = flattenSwitchInterfacePortSecurityGuestVlan(i["guest-vlan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "guest_auth_delay"
	if _, ok := i["guest-auth-delay"]; ok {

		result["guest_auth_delay"] = flattenSwitchInterfacePortSecurityGuestAuthDelay(i["guest-auth-delay"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "client_limit"
	if _, ok := i["client-limit"]; ok {

		result["client_limit"] = flattenSwitchInterfacePortSecurityClientLimit(i["client-limit"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auth_order"
	if _, ok := i["auth-order"]; ok {

		result["auth_order"] = flattenSwitchInterfacePortSecurityAuthOrder(i["auth-order"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "framevid_apply"
	if _, ok := i["framevid-apply"]; ok {

		result["framevid_apply"] = flattenSwitchInterfacePortSecurityFramevidApply(i["framevid-apply"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "eap_auto_untagged_vlans"
	if _, ok := i["eap-auto-untagged-vlans"]; ok {

		result["eap_auto_untagged_vlans"] = flattenSwitchInterfacePortSecurityEapAutoUntaggedVlans(i["eap-auto-untagged-vlans"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "authserver_timeout_tagged_lldp_voice_vlanid"
	if _, ok := i["authserver-timeout-tagged-lldp-voice-vlanid"]; ok {

		result["authserver_timeout_tagged_lldp_voice_vlanid"] = flattenSwitchInterfacePortSecurityAuthserverTimeoutTaggedLldpVoiceVlanid(i["authserver-timeout-tagged-lldp-voice-vlanid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "authserver_timeout_tagged_vlanid"
	if _, ok := i["authserver-timeout-tagged-vlanid"]; ok {

		result["authserver_timeout_tagged_vlanid"] = flattenSwitchInterfacePortSecurityAuthserverTimeoutTaggedVlanid(i["authserver-timeout-tagged-vlanid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mac_auth_bypass"
	if _, ok := i["mac-auth-bypass"]; ok {

		result["mac_auth_bypass"] = flattenSwitchInterfacePortSecurityMacAuthBypass(i["mac-auth-bypass"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auth_priority"
	if _, ok := i["auth-priority"]; ok {

		result["auth_priority"] = flattenSwitchInterfacePortSecurityAuthPriority(i["auth-priority"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "radius_timeout_overwrite"
	if _, ok := i["radius-timeout-overwrite"]; ok {

		result["radius_timeout_overwrite"] = flattenSwitchInterfacePortSecurityRadiusTimeoutOverwrite(i["radius-timeout-overwrite"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "authserver_timeout_period"
	if _, ok := i["authserver-timeout-period"]; ok {

		result["authserver_timeout_period"] = flattenSwitchInterfacePortSecurityAuthserverTimeoutPeriod(i["authserver-timeout-period"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "open_auth"
	if _, ok := i["open-auth"]; ok {

		result["open_auth"] = flattenSwitchInterfacePortSecurityOpenAuth(i["open-auth"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "authserver_timeout_vlanid"
	if _, ok := i["authserver-timeout-vlanid"]; ok {

		result["authserver_timeout_vlanid"] = flattenSwitchInterfacePortSecurityAuthserverTimeoutVlanid(i["authserver-timeout-vlanid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "quarantine_vlan"
	if _, ok := i["quarantine-vlan"]; ok {

		result["quarantine_vlan"] = flattenSwitchInterfacePortSecurityQuarantineVlan(i["quarantine-vlan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "guest_vlanid"
	if _, ok := i["guest-vlanid"]; ok {

		result["guest_vlanid"] = flattenSwitchInterfacePortSecurityGuestVlanid(i["guest-vlanid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "eap_egress_tagged"
	if _, ok := i["eap-egress-tagged"]; ok {

		result["eap_egress_tagged"] = flattenSwitchInterfacePortSecurityEapEgressTagged(i["eap-egress-tagged"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "eap_passthru"
	if _, ok := i["eap-passthru"]; ok {

		result["eap_passthru"] = flattenSwitchInterfacePortSecurityEapPassthru(i["eap-passthru"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchInterfacePortSecurityMacsecPaeMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityAuthFailVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityMacsecProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityAllowMacMoveTo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityAuthFailVlanid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityPortSecurityMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityAuthserverTimeoutTagged(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityMabEapolRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityDacl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityAuthserverTimeoutVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityAllowMacMove(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityGuestVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityGuestAuthDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityClientLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityAuthOrder(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityFramevidApply(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityEapAutoUntaggedVlans(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityAuthserverTimeoutTaggedLldpVoiceVlanid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityAuthserverTimeoutTaggedVlanid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityMacAuthBypass(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityAuthPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityRadiusTimeoutOverwrite(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityAuthserverTimeoutPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityOpenAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityAuthserverTimeoutVlanid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityQuarantineVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityGuestVlanid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityEapEgressTagged(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePortSecurityEapPassthru(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceDefaultCos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceMcastSnoopingFloodTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfaceVlanTpid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePacketSampleRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchInterfacePtpPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchInterface(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("allow_arp_monitor", flattenSwitchInterfaceAllowArpMonitor(o["allow-arp-monitor"], d, "allow_arp_monitor", sv)); err != nil {
		if !fortiAPIPatch(o["allow-arp-monitor"]) {
			return fmt.Errorf("Error reading allow_arp_monitor: %v", err)
		}
	}

	if err = d.Set("fortilink_l3_mode", flattenSwitchInterfaceFortilinkL3Mode(o["fortilink-l3-mode"], d, "fortilink_l3_mode", sv)); err != nil {
		if !fortiAPIPatch(o["fortilink-l3-mode"]) {
			return fmt.Errorf("Error reading fortilink_l3_mode: %v", err)
		}
	}

	if err = d.Set("igmp_snooping_flood_reports", flattenSwitchInterfaceIgmpSnoopingFloodReports(o["igmp-snooping-flood-reports"], d, "igmp_snooping_flood_reports", sv)); err != nil {
		if !fortiAPIPatch(o["igmp-snooping-flood-reports"]) {
			return fmt.Errorf("Error reading igmp_snooping_flood_reports: %v", err)
		}
	}

	if err = d.Set("interface_mode", flattenSwitchInterfaceInterfaceMode(o["interface-mode"], d, "interface_mode", sv)); err != nil {
		if !fortiAPIPatch(o["interface-mode"]) {
			return fmt.Errorf("Error reading interface_mode: %v", err)
		}
	}

	if err = d.Set("auto_discovery_fortilink", flattenSwitchInterfaceAutoDiscoveryFortilink(o["auto-discovery-fortilink"], d, "auto_discovery_fortilink", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discovery-fortilink"]) {
			return fmt.Errorf("Error reading auto_discovery_fortilink: %v", err)
		}
	}

	if err = d.Set("trust_dot1p_map", flattenSwitchInterfaceTrustDot1PMap(o["trust-dot1p-map"], d, "trust_dot1p_map", sv)); err != nil {
		if !fortiAPIPatch(o["trust-dot1p-map"]) {
			return fmt.Errorf("Error reading trust_dot1p_map: %v", err)
		}
	}

	if err = d.Set("discard_mode", flattenSwitchInterfaceDiscardMode(o["discard-mode"], d, "discard_mode", sv)); err != nil {
		if !fortiAPIPatch(o["discard-mode"]) {
			return fmt.Errorf("Error reading discard_mode: %v", err)
		}
	}

	if err = d.Set("qos_policy", flattenSwitchInterfaceQosPolicy(o["qos-policy"], d, "qos_policy", sv)); err != nil {
		if !fortiAPIPatch(o["qos-policy"]) {
			return fmt.Errorf("Error reading qos_policy: %v", err)
		}
	}

	if err = d.Set("arp_inspection_trust", flattenSwitchInterfaceArpInspectionTrust(o["arp-inspection-trust"], d, "arp_inspection_trust", sv)); err != nil {
		if !fortiAPIPatch(o["arp-inspection-trust"]) {
			return fmt.Errorf("Error reading arp_inspection_trust: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("security_groups", flattenSwitchInterfaceSecurityGroups(o["security-groups"], d, "security_groups", sv)); err != nil {
			if !fortiAPIPatch(o["security-groups"]) {
				return fmt.Errorf("Error reading security_groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("security_groups"); ok {
			if err = d.Set("security_groups", flattenSwitchInterfaceSecurityGroups(o["security-groups"], d, "security_groups", sv)); err != nil {
				if !fortiAPIPatch(o["security-groups"]) {
					return fmt.Errorf("Error reading security_groups: %v", err)
				}
			}
		}
	}

	if err = d.Set("private_vlan", flattenSwitchInterfacePrivateVlan(o["private-vlan"], d, "private_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["private-vlan"]) {
			return fmt.Errorf("Error reading private_vlan: %v", err)
		}
	}

	if err = d.Set("sub_vlan", flattenSwitchInterfaceSubVlan(o["sub-vlan"], d, "sub_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["sub-vlan"]) {
			return fmt.Errorf("Error reading sub_vlan: %v", err)
		}
	}

	if err = d.Set("allowed_sub_vlans", flattenSwitchInterfaceAllowedSubVlans(o["allowed-sub-vlans"], d, "allowed_sub_vlans", sv)); err != nil {
		if !fortiAPIPatch(o["allowed-sub-vlans"]) {
			return fmt.Errorf("Error reading allowed_sub_vlans: %v", err)
		}
	}

	if err = d.Set("sample_direction", flattenSwitchInterfaceSampleDirection(o["sample-direction"], d, "sample_direction", sv)); err != nil {
		if !fortiAPIPatch(o["sample-direction"]) {
			return fmt.Errorf("Error reading sample_direction: %v", err)
		}
	}

	if err = d.Set("dhcp_snoop_option82_trust", flattenSwitchInterfaceDhcpSnoopOption82Trust(o["dhcp-snoop-option82-trust"], d, "dhcp_snoop_option82_trust", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-snoop-option82-trust"]) {
			return fmt.Errorf("Error reading dhcp_snoop_option82_trust: %v", err)
		}
	}

	if err = d.Set("edge_port", flattenSwitchInterfaceEdgePort(o["edge-port"], d, "edge_port", sv)); err != nil {
		if !fortiAPIPatch(o["edge-port"]) {
			return fmt.Errorf("Error reading edge_port: %v", err)
		}
	}

	if err = d.Set("stp_bpdu_guard", flattenSwitchInterfaceStpBpduGuard(o["stp-bpdu-guard"], d, "stp_bpdu_guard", sv)); err != nil {
		if !fortiAPIPatch(o["stp-bpdu-guard"]) {
			return fmt.Errorf("Error reading stp_bpdu_guard: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dhcp_snoop_option82_override", flattenSwitchInterfaceDhcpSnoopOption82Override(o["dhcp-snoop-option82-override"], d, "dhcp_snoop_option82_override", sv)); err != nil {
			if !fortiAPIPatch(o["dhcp-snoop-option82-override"]) {
				return fmt.Errorf("Error reading dhcp_snoop_option82_override: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dhcp_snoop_option82_override"); ok {
			if err = d.Set("dhcp_snoop_option82_override", flattenSwitchInterfaceDhcpSnoopOption82Override(o["dhcp-snoop-option82-override"], d, "dhcp_snoop_option82_override", sv)); err != nil {
				if !fortiAPIPatch(o["dhcp-snoop-option82-override"]) {
					return fmt.Errorf("Error reading dhcp_snoop_option82_override: %v", err)
				}
			}
		}
	}

	if err = d.Set("ptp_status", flattenSwitchInterfacePtpStatus(o["ptp-status"], d, "ptp_status", sv)); err != nil {
		if !fortiAPIPatch(o["ptp-status"]) {
			return fmt.Errorf("Error reading ptp_status: %v", err)
		}
	}

	if err = d.Set("loop_guard", flattenSwitchInterfaceLoopGuard(o["loop-guard"], d, "loop_guard", sv)); err != nil {
		if !fortiAPIPatch(o["loop-guard"]) {
			return fmt.Errorf("Error reading loop_guard: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("qnq", flattenSwitchInterfaceQnq(o["qnq"], d, "qnq", sv)); err != nil {
			if !fortiAPIPatch(o["qnq"]) {
				return fmt.Errorf("Error reading qnq: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("qnq"); ok {
			if err = d.Set("qnq", flattenSwitchInterfaceQnq(o["qnq"], d, "qnq", sv)); err != nil {
				if !fortiAPIPatch(o["qnq"]) {
					return fmt.Errorf("Error reading qnq: %v", err)
				}
			}
		}
	}

	if err = d.Set("allowed_vlans", flattenSwitchInterfaceAllowedVlans(o["allowed-vlans"], d, "allowed_vlans", sv)); err != nil {
		if !fortiAPIPatch(o["allowed-vlans"]) {
			return fmt.Errorf("Error reading allowed_vlans: %v", err)
		}
	}

	if err = d.Set("nac", flattenSwitchInterfaceNac(o["nac"], d, "nac", sv)); err != nil {
		if !fortiAPIPatch(o["nac"]) {
			return fmt.Errorf("Error reading nac: %v", err)
		}
	}

	if err = d.Set("auto_discovery_fortilink_packet_interval", flattenSwitchInterfaceAutoDiscoveryFortilinkPacketInterval(o["auto-discovery-fortilink-packet-interval"], d, "auto_discovery_fortilink_packet_interval", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discovery-fortilink-packet-interval"]) {
			return fmt.Errorf("Error reading auto_discovery_fortilink_packet_interval: %v", err)
		}
	}

	if err = d.Set("primary_vlan", flattenSwitchInterfacePrimaryVlan(o["primary-vlan"], d, "primary_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["primary-vlan"]) {
			return fmt.Errorf("Error reading primary_vlan: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vlan_mapping", flattenSwitchInterfaceVlanMapping(o["vlan-mapping"], d, "vlan_mapping", sv)); err != nil {
			if !fortiAPIPatch(o["vlan-mapping"]) {
				return fmt.Errorf("Error reading vlan_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vlan_mapping"); ok {
			if err = d.Set("vlan_mapping", flattenSwitchInterfaceVlanMapping(o["vlan-mapping"], d, "vlan_mapping", sv)); err != nil {
				if !fortiAPIPatch(o["vlan-mapping"]) {
					return fmt.Errorf("Error reading vlan_mapping: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("raguard", flattenSwitchInterfaceRaguard(o["raguard"], d, "raguard", sv)); err != nil {
			if !fortiAPIPatch(o["raguard"]) {
				return fmt.Errorf("Error reading raguard: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("raguard"); ok {
			if err = d.Set("raguard", flattenSwitchInterfaceRaguard(o["raguard"], d, "raguard", sv)); err != nil {
				if !fortiAPIPatch(o["raguard"]) {
					return fmt.Errorf("Error reading raguard: %v", err)
				}
			}
		}
	}

	if err = d.Set("stp_bpdu_guard_timeout", flattenSwitchInterfaceStpBpduGuardTimeout(o["stp-bpdu-guard-timeout"], d, "stp_bpdu_guard_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["stp-bpdu-guard-timeout"]) {
			return fmt.Errorf("Error reading stp_bpdu_guard_timeout: %v", err)
		}
	}

	if err = d.Set("stp_loop_protection", flattenSwitchInterfaceStpLoopProtection(o["stp-loop-protection"], d, "stp_loop_protection", sv)); err != nil {
		if !fortiAPIPatch(o["stp-loop-protection"]) {
			return fmt.Errorf("Error reading stp_loop_protection: %v", err)
		}
	}

	if err = d.Set("learning_limit", flattenSwitchInterfaceLearningLimit(o["learning-limit"], d, "learning_limit", sv)); err != nil {
		if !fortiAPIPatch(o["learning-limit"]) {
			return fmt.Errorf("Error reading learning_limit: %v", err)
		}
	}

	if err = d.Set("learning_limit_action", flattenSwitchInterfaceLearningLimitAction(o["learning-limit-action"], d, "learning_limit_action", sv)); err != nil {
		if !fortiAPIPatch(o["learning-limit-action"]) {
			return fmt.Errorf("Error reading learning_limit_action: %v", err)
		}
	}

	if err = d.Set("loop_guard_mac_move_threshold", flattenSwitchInterfaceLoopGuardMacMoveThreshold(o["loop-guard-mac-move-threshold"], d, "loop_guard_mac_move_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["loop-guard-mac-move-threshold"]) {
			return fmt.Errorf("Error reading loop_guard_mac_move_threshold: %v", err)
		}
	}

	if err = d.Set("log_mac_event", flattenSwitchInterfaceLogMacEvent(o["log-mac-event"], d, "log_mac_event", sv)); err != nil {
		if !fortiAPIPatch(o["log-mac-event"]) {
			return fmt.Errorf("Error reading log_mac_event: %v", err)
		}
	}

	if err = d.Set("dhcp_snoop_learning_limit", flattenSwitchInterfaceDhcpSnoopLearningLimit(o["dhcp-snoop-learning-limit"], d, "dhcp_snoop_learning_limit", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-snoop-learning-limit"]) {
			return fmt.Errorf("Error reading dhcp_snoop_learning_limit: %v", err)
		}
	}

	if err = d.Set("type", flattenSwitchInterfaceType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("snmp_index", flattenSwitchInterfaceSnmpIndex(o["snmp-index"], d, "snmp_index", sv)); err != nil {
		if !fortiAPIPatch(o["snmp-index"]) {
			return fmt.Errorf("Error reading snmp_index: %v", err)
		}
	}

	if err = d.Set("dhcp_snoop_learning_limit_check", flattenSwitchInterfaceDhcpSnoopLearningLimitCheck(o["dhcp-snoop-learning-limit-check"], d, "dhcp_snoop_learning_limit_check", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-snoop-learning-limit-check"]) {
			return fmt.Errorf("Error reading dhcp_snoop_learning_limit_check: %v", err)
		}
	}

	if err = d.Set("switch_port_mode", flattenSwitchInterfaceSwitchPortMode(o["switch-port-mode"], d, "switch_port_mode", sv)); err != nil {
		if !fortiAPIPatch(o["switch-port-mode"]) {
			return fmt.Errorf("Error reading switch_port_mode: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchInterfaceDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("stp_state", flattenSwitchInterfaceStpState(o["stp-state"], d, "stp_state", sv)); err != nil {
		if !fortiAPIPatch(o["stp-state"]) {
			return fmt.Errorf("Error reading stp_state: %v", err)
		}
	}

	if err = d.Set("vlan_traffic_type", flattenSwitchInterfaceVlanTrafficType(o["vlan-traffic-type"], d, "vlan_traffic_type", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-traffic-type"]) {
			return fmt.Errorf("Error reading vlan_traffic_type: %v", err)
		}
	}

	if err = d.Set("sticky_mac", flattenSwitchInterfaceStickyMac(o["sticky-mac"], d, "sticky_mac", sv)); err != nil {
		if !fortiAPIPatch(o["sticky-mac"]) {
			return fmt.Errorf("Error reading sticky_mac: %v", err)
		}
	}

	if err = d.Set("packet_sampler", flattenSwitchInterfacePacketSampler(o["packet-sampler"], d, "packet_sampler", sv)); err != nil {
		if !fortiAPIPatch(o["packet-sampler"]) {
			return fmt.Errorf("Error reading packet_sampler: %v", err)
		}
	}

	if err = d.Set("loop_guard_timeout", flattenSwitchInterfaceLoopGuardTimeout(o["loop-guard-timeout"], d, "loop_guard_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["loop-guard-timeout"]) {
			return fmt.Errorf("Error reading loop_guard_timeout: %v", err)
		}
	}

	if err = d.Set("rpvst_port", flattenSwitchInterfaceRpvstPort(o["rpvst-port"], d, "rpvst_port", sv)); err != nil {
		if !fortiAPIPatch(o["rpvst-port"]) {
			return fmt.Errorf("Error reading rpvst_port: %v", err)
		}
	}

	if err = d.Set("ip_mac_binding", flattenSwitchInterfaceIpMacBinding(o["ip-mac-binding"], d, "ip_mac_binding", sv)); err != nil {
		if !fortiAPIPatch(o["ip-mac-binding"]) {
			return fmt.Errorf("Error reading ip_mac_binding: %v", err)
		}
	}

	if err = d.Set("mld_snooping_flood_reports", flattenSwitchInterfaceMldSnoopingFloodReports(o["mld-snooping-flood-reports"], d, "mld_snooping_flood_reports", sv)); err != nil {
		if !fortiAPIPatch(o["mld-snooping-flood-reports"]) {
			return fmt.Errorf("Error reading mld_snooping_flood_reports: %v", err)
		}
	}

	if err = d.Set("sflow_counter_interval", flattenSwitchInterfaceSflowCounterInterval(o["sflow-counter-interval"], d, "sflow_counter_interval", sv)); err != nil {
		if !fortiAPIPatch(o["sflow-counter-interval"]) {
			return fmt.Errorf("Error reading sflow_counter_interval: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchInterfaceName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("native_vlan", flattenSwitchInterfaceNativeVlan(o["native-vlan"], d, "native_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["native-vlan"]) {
			return fmt.Errorf("Error reading native_vlan: %v", err)
		}
	}

	if err = d.Set("vlan_mapping_miss_drop", flattenSwitchInterfaceVlanMappingMissDrop(o["vlan-mapping-miss-drop"], d, "vlan_mapping_miss_drop", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-mapping-miss-drop"]) {
			return fmt.Errorf("Error reading vlan_mapping_miss_drop: %v", err)
		}
	}

	if err = d.Set("untagged_vlans", flattenSwitchInterfaceUntaggedVlans(o["untagged-vlans"], d, "untagged_vlans", sv)); err != nil {
		if !fortiAPIPatch(o["untagged-vlans"]) {
			return fmt.Errorf("Error reading untagged_vlans: %v", err)
		}
	}

	if err = d.Set("private_vlan_port_type", flattenSwitchInterfacePrivateVlanPortType(o["private-vlan-port-type"], d, "private_vlan_port_type", sv)); err != nil {
		if !fortiAPIPatch(o["private-vlan-port-type"]) {
			return fmt.Errorf("Error reading private_vlan_port_type: %v", err)
		}
	}

	if err = d.Set("dhcp_snooping", flattenSwitchInterfaceDhcpSnooping(o["dhcp-snooping"], d, "dhcp_snooping", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-snooping"]) {
			return fmt.Errorf("Error reading dhcp_snooping: %v", err)
		}
	}

	if err = d.Set("filter_sub_vlans", flattenSwitchInterfaceFilterSubVlans(o["filter-sub-vlans"], d, "filter_sub_vlans", sv)); err != nil {
		if !fortiAPIPatch(o["filter-sub-vlans"]) {
			return fmt.Errorf("Error reading filter_sub_vlans: %v", err)
		}
	}

	if err = d.Set("stp_root_guard", flattenSwitchInterfaceStpRootGuard(o["stp-root-guard"], d, "stp_root_guard", sv)); err != nil {
		if !fortiAPIPatch(o["stp-root-guard"]) {
			return fmt.Errorf("Error reading stp_root_guard: %v", err)
		}
	}

	if err = d.Set("trust_ip_dscp_map", flattenSwitchInterfaceTrustIpDscpMap(o["trust-ip-dscp-map"], d, "trust_ip_dscp_map", sv)); err != nil {
		if !fortiAPIPatch(o["trust-ip-dscp-map"]) {
			return fmt.Errorf("Error reading trust_ip_dscp_map: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("port_security", flattenSwitchInterfacePortSecurity(o["port-security"], d, "port_security", sv)); err != nil {
			if !fortiAPIPatch(o["port-security"]) {
				return fmt.Errorf("Error reading port_security: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("port_security"); ok {
			if err = d.Set("port_security", flattenSwitchInterfacePortSecurity(o["port-security"], d, "port_security", sv)); err != nil {
				if !fortiAPIPatch(o["port-security"]) {
					return fmt.Errorf("Error reading port_security: %v", err)
				}
			}
		}
	}

	if err = d.Set("default_cos", flattenSwitchInterfaceDefaultCos(o["default-cos"], d, "default_cos", sv)); err != nil {
		if !fortiAPIPatch(o["default-cos"]) {
			return fmt.Errorf("Error reading default_cos: %v", err)
		}
	}

	if err = d.Set("mcast_snooping_flood_traffic", flattenSwitchInterfaceMcastSnoopingFloodTraffic(o["mcast-snooping-flood-traffic"], d, "mcast_snooping_flood_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["mcast-snooping-flood-traffic"]) {
			return fmt.Errorf("Error reading mcast_snooping_flood_traffic: %v", err)
		}
	}

	if err = d.Set("vlan_tpid", flattenSwitchInterfaceVlanTpid(o["vlan-tpid"], d, "vlan_tpid", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-tpid"]) {
			return fmt.Errorf("Error reading vlan_tpid: %v", err)
		}
	}

	if err = d.Set("packet_sample_rate", flattenSwitchInterfacePacketSampleRate(o["packet-sample-rate"], d, "packet_sample_rate", sv)); err != nil {
		if !fortiAPIPatch(o["packet-sample-rate"]) {
			return fmt.Errorf("Error reading packet_sample_rate: %v", err)
		}
	}

	if err = d.Set("ptp_policy", flattenSwitchInterfacePtpPolicy(o["ptp-policy"], d, "ptp_policy", sv)); err != nil {
		if !fortiAPIPatch(o["ptp-policy"]) {
			return fmt.Errorf("Error reading ptp_policy: %v", err)
		}
	}

	return nil
}

func flattenSwitchInterfaceFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchInterfaceAllowArpMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceFortilinkL3Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceIgmpSnoopingFloodReports(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceInterfaceMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceAutoDiscoveryFortilink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceTrustDot1PMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceDiscardMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQosPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceArpInspectionTrust(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceSecurityGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSwitchInterfaceSecurityGroupsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchInterfaceSecurityGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePrivateVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceSubVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceAllowedSubVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceSampleDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceDhcpSnoopOption82Trust(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceEdgePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceStpBpduGuard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceDhcpSnoopOption82Override(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["circuit-id"], _ = expandSwitchInterfaceDhcpSnoopOption82OverrideCircuitId(d, i["circuit_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSwitchInterfaceDhcpSnoopOption82OverrideId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["remote-id"], _ = expandSwitchInterfaceDhcpSnoopOption82OverrideRemoteId(d, i["remote_id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchInterfaceDhcpSnoopOption82OverrideCircuitId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceDhcpSnoopOption82OverrideId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceDhcpSnoopOption82OverrideRemoteId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePtpStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceLoopGuard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnq(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {

		result["status"], _ = expandSwitchInterfaceQnqStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vlan_mapping"
	if _, ok := d.GetOk(pre_append); ok {

		result["vlan-mapping"], _ = expandSwitchInterfaceQnqVlanMapping(d, i["vlan_mapping"], pre_append, sv)
	} else {
		result["vlan-mapping"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "stp_qnq_admin"
	if _, ok := d.GetOk(pre_append); ok {

		result["stp-qnq-admin"], _ = expandSwitchInterfaceQnqStpQnqAdmin(d, i["stp_qnq_admin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "native_c_vlan"
	if _, ok := d.GetOk(pre_append); ok {

		result["native-c-vlan"], _ = expandSwitchInterfaceQnqNativeCVlan(d, i["native_c_vlan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "remove_inner"
	if _, ok := d.GetOk(pre_append); ok {

		result["remove-inner"], _ = expandSwitchInterfaceQnqRemoveInner(d, i["remove_inner"], pre_append, sv)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok {

		result["priority"], _ = expandSwitchInterfaceQnqPriority(d, i["priority"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vlan_mapping_miss_drop"
	if _, ok := d.GetOk(pre_append); ok {

		result["vlan-mapping-miss-drop"], _ = expandSwitchInterfaceQnqVlanMappingMissDrop(d, i["vlan_mapping_miss_drop"], pre_append, sv)
	}
	pre_append = pre + ".0." + "untagged_s_vlan"
	if _, ok := d.GetOk(pre_append); ok {

		result["untagged-s-vlan"], _ = expandSwitchInterfaceQnqUntaggedSVlan(d, i["untagged_s_vlan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "add_inner"
	if _, ok := d.GetOk(pre_append); ok {

		result["add-inner"], _ = expandSwitchInterfaceQnqAddInner(d, i["add_inner"], pre_append, sv)
	}
	pre_append = pre + ".0." + "allowed_c_vlan"
	if _, ok := d.GetOk(pre_append); ok {

		result["allowed-c-vlan"], _ = expandSwitchInterfaceQnqAllowedCVlan(d, i["allowed_c_vlan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "edge_type"
	if _, ok := d.GetOk(pre_append); ok {

		result["edge-type"], _ = expandSwitchInterfaceQnqEdgeType(d, i["edge_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "s_tag_priority"
	if _, ok := d.GetOk(pre_append); ok {

		result["s-tag-priority"], _ = expandSwitchInterfaceQnqSTagPriority(d, i["s_tag_priority"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchInterfaceQnqStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnqVlanMapping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "new_s_vlan"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["new-s-vlan"], _ = expandSwitchInterfaceQnqVlanMappingNewSVlan(d, i["new_s_vlan"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_c_vlan"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["match-c-vlan"], _ = expandSwitchInterfaceQnqVlanMappingMatchCVlan(d, i["match_c_vlan"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSwitchInterfaceQnqVlanMappingId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandSwitchInterfaceQnqVlanMappingDescription(d, i["description"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchInterfaceQnqVlanMappingNewSVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnqVlanMappingMatchCVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnqVlanMappingId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnqVlanMappingDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnqStpQnqAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnqNativeCVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnqRemoveInner(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnqPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnqVlanMappingMissDrop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnqUntaggedSVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnqAddInner(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnqAllowedCVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnqEdgeType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceQnqSTagPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceAllowedVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceNac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceAutoDiscoveryFortilinkPacketInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePrimaryVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceVlanMapping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["direction"], _ = expandSwitchInterfaceVlanMappingDirection(d, i["direction"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandSwitchInterfaceVlanMappingDescription(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_s_vlan"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["match-s-vlan"], _ = expandSwitchInterfaceVlanMappingMatchSVlan(d, i["match_s_vlan"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "new_s_vlan"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["new-s-vlan"], _ = expandSwitchInterfaceVlanMappingNewSVlan(d, i["new_s_vlan"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandSwitchInterfaceVlanMappingAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_c_vlan"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["match-c-vlan"], _ = expandSwitchInterfaceVlanMappingMatchCVlan(d, i["match_c_vlan"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSwitchInterfaceVlanMappingId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchInterfaceVlanMappingDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceVlanMappingDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceVlanMappingMatchSVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceVlanMappingNewSVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceVlanMappingAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceVlanMappingMatchCVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceVlanMappingId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceRaguard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_list"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["vlan-list"], _ = expandSwitchInterfaceRaguardVlanList(d, i["vlan_list"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSwitchInterfaceRaguardId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "raguard_policy"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["raguard-policy"], _ = expandSwitchInterfaceRaguardRaguardPolicy(d, i["raguard_policy"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchInterfaceRaguardVlanList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceRaguardId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceRaguardRaguardPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceStpBpduGuardTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceStpLoopProtection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceLearningLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceLearningLimitAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceLoopGuardMacMoveThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceLogMacEvent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceDhcpSnoopLearningLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceSnmpIndex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceDhcpSnoopLearningLimitCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceSwitchPortMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceStpState(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceVlanTrafficType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceStickyMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePacketSampler(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceLoopGuardTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceRpvstPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceIpMacBinding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceMldSnoopingFloodReports(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceSflowCounterInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceNativeVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceVlanMappingMissDrop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceUntaggedVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePrivateVlanPortType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceDhcpSnooping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceFilterSubVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceStpRootGuard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceTrustIpDscpMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "macsec_pae_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["macsec-pae-mode"], _ = expandSwitchInterfacePortSecurityMacsecPaeMode(d, i["macsec_pae_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auth_fail_vlan"
	if _, ok := d.GetOk(pre_append); ok {

		result["auth-fail-vlan"], _ = expandSwitchInterfacePortSecurityAuthFailVlan(d, i["auth_fail_vlan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "macsec_profile"
	if _, ok := d.GetOk(pre_append); ok {

		result["macsec-profile"], _ = expandSwitchInterfacePortSecurityMacsecProfile(d, i["macsec_profile"], pre_append, sv)
	}
	pre_append = pre + ".0." + "allow_mac_move_to"
	if _, ok := d.GetOk(pre_append); ok {

		result["allow-mac-move-to"], _ = expandSwitchInterfacePortSecurityAllowMacMoveTo(d, i["allow_mac_move_to"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auth_fail_vlanid"
	if _, ok := d.GetOk(pre_append); ok {

		result["auth-fail-vlanid"], _ = expandSwitchInterfacePortSecurityAuthFailVlanid(d, i["auth_fail_vlanid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port_security_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["port-security-mode"], _ = expandSwitchInterfacePortSecurityPortSecurityMode(d, i["port_security_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "authserver_timeout_tagged"
	if _, ok := d.GetOk(pre_append); ok {

		result["authserver-timeout-tagged"], _ = expandSwitchInterfacePortSecurityAuthserverTimeoutTagged(d, i["authserver_timeout_tagged"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mab_eapol_request"
	if _, ok := d.GetOk(pre_append); ok {

		result["mab-eapol-request"], _ = expandSwitchInterfacePortSecurityMabEapolRequest(d, i["mab_eapol_request"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dacl"
	if _, ok := d.GetOk(pre_append); ok {

		result["dacl"], _ = expandSwitchInterfacePortSecurityDacl(d, i["dacl"], pre_append, sv)
	}
	pre_append = pre + ".0." + "authserver_timeout_vlan"
	if _, ok := d.GetOk(pre_append); ok {

		result["authserver-timeout-vlan"], _ = expandSwitchInterfacePortSecurityAuthserverTimeoutVlan(d, i["authserver_timeout_vlan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "allow_mac_move"
	if _, ok := d.GetOk(pre_append); ok {

		result["allow-mac-move"], _ = expandSwitchInterfacePortSecurityAllowMacMove(d, i["allow_mac_move"], pre_append, sv)
	}
	pre_append = pre + ".0." + "guest_vlan"
	if _, ok := d.GetOk(pre_append); ok {

		result["guest-vlan"], _ = expandSwitchInterfacePortSecurityGuestVlan(d, i["guest_vlan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "guest_auth_delay"
	if _, ok := d.GetOk(pre_append); ok {

		result["guest-auth-delay"], _ = expandSwitchInterfacePortSecurityGuestAuthDelay(d, i["guest_auth_delay"], pre_append, sv)
	}
	pre_append = pre + ".0." + "client_limit"
	if _, ok := d.GetOk(pre_append); ok {

		result["client-limit"], _ = expandSwitchInterfacePortSecurityClientLimit(d, i["client_limit"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auth_order"
	if _, ok := d.GetOk(pre_append); ok {

		result["auth-order"], _ = expandSwitchInterfacePortSecurityAuthOrder(d, i["auth_order"], pre_append, sv)
	}
	pre_append = pre + ".0." + "framevid_apply"
	if _, ok := d.GetOk(pre_append); ok {

		result["framevid-apply"], _ = expandSwitchInterfacePortSecurityFramevidApply(d, i["framevid_apply"], pre_append, sv)
	}
	pre_append = pre + ".0." + "eap_auto_untagged_vlans"
	if _, ok := d.GetOk(pre_append); ok {

		result["eap-auto-untagged-vlans"], _ = expandSwitchInterfacePortSecurityEapAutoUntaggedVlans(d, i["eap_auto_untagged_vlans"], pre_append, sv)
	}
	pre_append = pre + ".0." + "authserver_timeout_tagged_lldp_voice_vlanid"
	if _, ok := d.GetOk(pre_append); ok {

		result["authserver-timeout-tagged-lldp-voice-vlanid"], _ = expandSwitchInterfacePortSecurityAuthserverTimeoutTaggedLldpVoiceVlanid(d, i["authserver_timeout_tagged_lldp_voice_vlanid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "authserver_timeout_tagged_vlanid"
	if _, ok := d.GetOk(pre_append); ok {

		result["authserver-timeout-tagged-vlanid"], _ = expandSwitchInterfacePortSecurityAuthserverTimeoutTaggedVlanid(d, i["authserver_timeout_tagged_vlanid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mac_auth_bypass"
	if _, ok := d.GetOk(pre_append); ok {

		result["mac-auth-bypass"], _ = expandSwitchInterfacePortSecurityMacAuthBypass(d, i["mac_auth_bypass"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auth_priority"
	if _, ok := d.GetOk(pre_append); ok {

		result["auth-priority"], _ = expandSwitchInterfacePortSecurityAuthPriority(d, i["auth_priority"], pre_append, sv)
	}
	pre_append = pre + ".0." + "radius_timeout_overwrite"
	if _, ok := d.GetOk(pre_append); ok {

		result["radius-timeout-overwrite"], _ = expandSwitchInterfacePortSecurityRadiusTimeoutOverwrite(d, i["radius_timeout_overwrite"], pre_append, sv)
	}
	pre_append = pre + ".0." + "authserver_timeout_period"
	if _, ok := d.GetOk(pre_append); ok {

		result["authserver-timeout-period"], _ = expandSwitchInterfacePortSecurityAuthserverTimeoutPeriod(d, i["authserver_timeout_period"], pre_append, sv)
	}
	pre_append = pre + ".0." + "open_auth"
	if _, ok := d.GetOk(pre_append); ok {

		result["open-auth"], _ = expandSwitchInterfacePortSecurityOpenAuth(d, i["open_auth"], pre_append, sv)
	}
	pre_append = pre + ".0." + "authserver_timeout_vlanid"
	if _, ok := d.GetOk(pre_append); ok {

		result["authserver-timeout-vlanid"], _ = expandSwitchInterfacePortSecurityAuthserverTimeoutVlanid(d, i["authserver_timeout_vlanid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "quarantine_vlan"
	if _, ok := d.GetOk(pre_append); ok {

		result["quarantine-vlan"], _ = expandSwitchInterfacePortSecurityQuarantineVlan(d, i["quarantine_vlan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "guest_vlanid"
	if _, ok := d.GetOk(pre_append); ok {

		result["guest-vlanid"], _ = expandSwitchInterfacePortSecurityGuestVlanid(d, i["guest_vlanid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "eap_egress_tagged"
	if _, ok := d.GetOk(pre_append); ok {

		result["eap-egress-tagged"], _ = expandSwitchInterfacePortSecurityEapEgressTagged(d, i["eap_egress_tagged"], pre_append, sv)
	}
	pre_append = pre + ".0." + "eap_passthru"
	if _, ok := d.GetOk(pre_append); ok {

		result["eap-passthru"], _ = expandSwitchInterfacePortSecurityEapPassthru(d, i["eap_passthru"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchInterfacePortSecurityMacsecPaeMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityAuthFailVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityMacsecProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityAllowMacMoveTo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityAuthFailVlanid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityPortSecurityMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityAuthserverTimeoutTagged(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityMabEapolRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityDacl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityAuthserverTimeoutVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityAllowMacMove(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityGuestVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityGuestAuthDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityClientLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityAuthOrder(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityFramevidApply(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityEapAutoUntaggedVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityAuthserverTimeoutTaggedLldpVoiceVlanid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityAuthserverTimeoutTaggedVlanid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityMacAuthBypass(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityAuthPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityRadiusTimeoutOverwrite(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityAuthserverTimeoutPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityOpenAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityAuthserverTimeoutVlanid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityQuarantineVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityGuestVlanid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityEapEgressTagged(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePortSecurityEapPassthru(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceDefaultCos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceMcastSnoopingFloodTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfaceVlanTpid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePacketSampleRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchInterfacePtpPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchInterface(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allow_arp_monitor"); ok {

		t, err := expandSwitchInterfaceAllowArpMonitor(d, v, "allow_arp_monitor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-arp-monitor"] = t
		}
	}

	if v, ok := d.GetOk("fortilink_l3_mode"); ok {

		t, err := expandSwitchInterfaceFortilinkL3Mode(d, v, "fortilink_l3_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink-l3-mode"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping_flood_reports"); ok {

		t, err := expandSwitchInterfaceIgmpSnoopingFloodReports(d, v, "igmp_snooping_flood_reports", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping-flood-reports"] = t
		}
	}

	if v, ok := d.GetOk("interface_mode"); ok {

		t, err := expandSwitchInterfaceInterfaceMode(d, v, "interface_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-mode"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_fortilink"); ok {

		t, err := expandSwitchInterfaceAutoDiscoveryFortilink(d, v, "auto_discovery_fortilink", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-fortilink"] = t
		}
	}

	if v, ok := d.GetOk("trust_dot1p_map"); ok {

		t, err := expandSwitchInterfaceTrustDot1PMap(d, v, "trust_dot1p_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-dot1p-map"] = t
		}
	}

	if v, ok := d.GetOk("discard_mode"); ok {

		t, err := expandSwitchInterfaceDiscardMode(d, v, "discard_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["discard-mode"] = t
		}
	}

	if v, ok := d.GetOk("qos_policy"); ok {

		t, err := expandSwitchInterfaceQosPolicy(d, v, "qos_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-policy"] = t
		}
	}

	if v, ok := d.GetOk("arp_inspection_trust"); ok {

		t, err := expandSwitchInterfaceArpInspectionTrust(d, v, "arp_inspection_trust", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-inspection-trust"] = t
		}
	}

	if v, ok := d.GetOk("security_groups"); ok || d.HasChange("security_groups") {

		t, err := expandSwitchInterfaceSecurityGroups(d, v, "security_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-groups"] = t
		}
	}

	if v, ok := d.GetOk("private_vlan"); ok {

		t, err := expandSwitchInterfacePrivateVlan(d, v, "private_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-vlan"] = t
		}
	}

	if v, ok := d.GetOk("sub_vlan"); ok {

		t, err := expandSwitchInterfaceSubVlan(d, v, "sub_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sub-vlan"] = t
		}
	}

	if v, ok := d.GetOk("allowed_sub_vlans"); ok {

		t, err := expandSwitchInterfaceAllowedSubVlans(d, v, "allowed_sub_vlans", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-sub-vlans"] = t
		}
	}

	if v, ok := d.GetOk("sample_direction"); ok {

		t, err := expandSwitchInterfaceSampleDirection(d, v, "sample_direction", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sample-direction"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snoop_option82_trust"); ok {

		t, err := expandSwitchInterfaceDhcpSnoopOption82Trust(d, v, "dhcp_snoop_option82_trust", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snoop-option82-trust"] = t
		}
	}

	if v, ok := d.GetOk("edge_port"); ok {

		t, err := expandSwitchInterfaceEdgePort(d, v, "edge_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["edge-port"] = t
		}
	}

	if v, ok := d.GetOk("stp_bpdu_guard"); ok {

		t, err := expandSwitchInterfaceStpBpduGuard(d, v, "stp_bpdu_guard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-bpdu-guard"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snoop_option82_override"); ok || d.HasChange("dhcp_snoop_option82_override") {

		t, err := expandSwitchInterfaceDhcpSnoopOption82Override(d, v, "dhcp_snoop_option82_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snoop-option82-override"] = t
		}
	}

	if v, ok := d.GetOk("ptp_status"); ok {

		t, err := expandSwitchInterfacePtpStatus(d, v, "ptp_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptp-status"] = t
		}
	}

	if v, ok := d.GetOk("loop_guard"); ok {

		t, err := expandSwitchInterfaceLoopGuard(d, v, "loop_guard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loop-guard"] = t
		}
	}

	if v, ok := d.GetOk("qnq"); ok {

		t, err := expandSwitchInterfaceQnq(d, v, "qnq", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qnq"] = t
		}
	}

	if v, ok := d.GetOk("allowed_vlans"); ok {

		t, err := expandSwitchInterfaceAllowedVlans(d, v, "allowed_vlans", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-vlans"] = t
		}
	}

	if v, ok := d.GetOk("nac"); ok {

		t, err := expandSwitchInterfaceNac(d, v, "nac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_fortilink_packet_interval"); ok {

		t, err := expandSwitchInterfaceAutoDiscoveryFortilinkPacketInterval(d, v, "auto_discovery_fortilink_packet_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-fortilink-packet-interval"] = t
		}
	}

	if v, ok := d.GetOk("primary_vlan"); ok {

		t, err := expandSwitchInterfacePrimaryVlan(d, v, "primary_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-vlan"] = t
		}
	}

	if v, ok := d.GetOk("vlan_mapping"); ok || d.HasChange("vlan_mapping") {

		t, err := expandSwitchInterfaceVlanMapping(d, v, "vlan_mapping", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-mapping"] = t
		}
	}

	if v, ok := d.GetOk("raguard"); ok || d.HasChange("raguard") {

		t, err := expandSwitchInterfaceRaguard(d, v, "raguard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["raguard"] = t
		}
	}

	if v, ok := d.GetOk("stp_bpdu_guard_timeout"); ok {

		t, err := expandSwitchInterfaceStpBpduGuardTimeout(d, v, "stp_bpdu_guard_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-bpdu-guard-timeout"] = t
		}
	}

	if v, ok := d.GetOk("stp_loop_protection"); ok {

		t, err := expandSwitchInterfaceStpLoopProtection(d, v, "stp_loop_protection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-loop-protection"] = t
		}
	}

	if v, ok := d.GetOk("learning_limit"); ok {

		t, err := expandSwitchInterfaceLearningLimit(d, v, "learning_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learning-limit"] = t
		}
	}

	if v, ok := d.GetOk("learning_limit_action"); ok {

		t, err := expandSwitchInterfaceLearningLimitAction(d, v, "learning_limit_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learning-limit-action"] = t
		}
	}

	if v, ok := d.GetOk("loop_guard_mac_move_threshold"); ok {

		t, err := expandSwitchInterfaceLoopGuardMacMoveThreshold(d, v, "loop_guard_mac_move_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loop-guard-mac-move-threshold"] = t
		}
	}

	if v, ok := d.GetOk("log_mac_event"); ok {

		t, err := expandSwitchInterfaceLogMacEvent(d, v, "log_mac_event", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-mac-event"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snoop_learning_limit"); ok {

		t, err := expandSwitchInterfaceDhcpSnoopLearningLimit(d, v, "dhcp_snoop_learning_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snoop-learning-limit"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandSwitchInterfaceType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("snmp_index"); ok {

		t, err := expandSwitchInterfaceSnmpIndex(d, v, "snmp_index", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snmp-index"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snoop_learning_limit_check"); ok {

		t, err := expandSwitchInterfaceDhcpSnoopLearningLimitCheck(d, v, "dhcp_snoop_learning_limit_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snoop-learning-limit-check"] = t
		}
	}

	if v, ok := d.GetOk("switch_port_mode"); ok {

		t, err := expandSwitchInterfaceSwitchPortMode(d, v, "switch_port_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-port-mode"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchInterfaceDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("stp_state"); ok {

		t, err := expandSwitchInterfaceStpState(d, v, "stp_state", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-state"] = t
		}
	}

	if v, ok := d.GetOk("vlan_traffic_type"); ok {

		t, err := expandSwitchInterfaceVlanTrafficType(d, v, "vlan_traffic_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-traffic-type"] = t
		}
	}

	if v, ok := d.GetOk("sticky_mac"); ok {

		t, err := expandSwitchInterfaceStickyMac(d, v, "sticky_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-mac"] = t
		}
	}

	if v, ok := d.GetOk("packet_sampler"); ok {

		t, err := expandSwitchInterfacePacketSampler(d, v, "packet_sampler", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-sampler"] = t
		}
	}

	if v, ok := d.GetOk("loop_guard_timeout"); ok {

		t, err := expandSwitchInterfaceLoopGuardTimeout(d, v, "loop_guard_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loop-guard-timeout"] = t
		}
	}

	if v, ok := d.GetOk("rpvst_port"); ok {

		t, err := expandSwitchInterfaceRpvstPort(d, v, "rpvst_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpvst-port"] = t
		}
	}

	if v, ok := d.GetOk("ip_mac_binding"); ok {

		t, err := expandSwitchInterfaceIpMacBinding(d, v, "ip_mac_binding", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-mac-binding"] = t
		}
	}

	if v, ok := d.GetOk("mld_snooping_flood_reports"); ok {

		t, err := expandSwitchInterfaceMldSnoopingFloodReports(d, v, "mld_snooping_flood_reports", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mld-snooping-flood-reports"] = t
		}
	}

	if v, ok := d.GetOkExists("sflow_counter_interval"); ok {

		t, err := expandSwitchInterfaceSflowCounterInterval(d, v, "sflow_counter_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sflow-counter-interval"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchInterfaceName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("native_vlan"); ok {

		t, err := expandSwitchInterfaceNativeVlan(d, v, "native_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["native-vlan"] = t
		}
	}

	if v, ok := d.GetOk("vlan_mapping_miss_drop"); ok {

		t, err := expandSwitchInterfaceVlanMappingMissDrop(d, v, "vlan_mapping_miss_drop", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-mapping-miss-drop"] = t
		}
	}

	if v, ok := d.GetOk("untagged_vlans"); ok {

		t, err := expandSwitchInterfaceUntaggedVlans(d, v, "untagged_vlans", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["untagged-vlans"] = t
		}
	}

	if v, ok := d.GetOk("private_vlan_port_type"); ok {

		t, err := expandSwitchInterfacePrivateVlanPortType(d, v, "private_vlan_port_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-vlan-port-type"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snooping"); ok {

		t, err := expandSwitchInterfaceDhcpSnooping(d, v, "dhcp_snooping", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("filter_sub_vlans"); ok {

		t, err := expandSwitchInterfaceFilterSubVlans(d, v, "filter_sub_vlans", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-sub-vlans"] = t
		}
	}

	if v, ok := d.GetOk("stp_root_guard"); ok {

		t, err := expandSwitchInterfaceStpRootGuard(d, v, "stp_root_guard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-root-guard"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip_dscp_map"); ok {

		t, err := expandSwitchInterfaceTrustIpDscpMap(d, v, "trust_ip_dscp_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip-dscp-map"] = t
		}
	}

	if v, ok := d.GetOk("port_security"); ok {

		t, err := expandSwitchInterfacePortSecurity(d, v, "port_security", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-security"] = t
		}
	}

	if v, ok := d.GetOk("default_cos"); ok {

		t, err := expandSwitchInterfaceDefaultCos(d, v, "default_cos", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-cos"] = t
		}
	}

	if v, ok := d.GetOk("mcast_snooping_flood_traffic"); ok {

		t, err := expandSwitchInterfaceMcastSnoopingFloodTraffic(d, v, "mcast_snooping_flood_traffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcast-snooping-flood-traffic"] = t
		}
	}

	if v, ok := d.GetOk("vlan_tpid"); ok {

		t, err := expandSwitchInterfaceVlanTpid(d, v, "vlan_tpid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-tpid"] = t
		}
	}

	if v, ok := d.GetOk("packet_sample_rate"); ok {

		t, err := expandSwitchInterfacePacketSampleRate(d, v, "packet_sample_rate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-sample-rate"] = t
		}
	}

	if v, ok := d.GetOk("ptp_policy"); ok {

		t, err := expandSwitchInterfacePtpPolicy(d, v, "ptp_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptp-policy"] = t
		}
	}

	return &obj, nil
}
