// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure global settings.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchGlobalUpdate,
		Read:   resourceSwitchGlobalRead,
		Update: resourceSwitchGlobalUpdate,
		Delete: resourceSwitchGlobalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fortilink_heartbeat_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mclag_split_brain_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_mac_limit_violations": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trunk_hash_unkunicast_src_dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"poe_power_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_mac_binding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vxlan_stp_virtual_root": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bpdu_learn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vxlan_dport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mac_address_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_wire_tpid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trunk_hash_unicast_src_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_fortilink_discovery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_path_in_ecmp_group": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trunk_hash_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortilink_p2p_tpid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"flood_vtp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_violation_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mirror_qos": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mclag_port_base": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"flapguard_retain_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mclag_split_brain_all_ports_down": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_address": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mclag_stp_aware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortilink_p2p_native_vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fortilink_vlan_optimization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vxlan_stp_virtual_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loop_guard_tx_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_isl_port_group": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vxlan_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"poe_alarm_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"poe_guard_band": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_stp_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"poe_power_budget": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vxlan_sport": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"mclag_peer_info_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mclag_igmpsnooping_aware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_vlan_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"poe_pre_standard_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_snooping_database_export": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"auto_isl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dmi_global_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2_memory_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mclag_split_brain_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forti_trunk_dmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_aging_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port_security": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mac_called_station_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mab_reauth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_reauth_attempt": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mac_case": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_password_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"link_down_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mab_entry_as": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_username_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"reauth_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mac_calling_station_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tx_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"flood_unknown_multicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2_memory_check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchGlobal(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchGlobal resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchGlobal(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchGlobal")
	}

	return resourceSwitchGlobalRead(d, m)
}

func resourceSwitchGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchGlobal(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchGlobal(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchGlobalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchGlobal(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchGlobal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchGlobal resource from API: %v", err)
	}
	return nil
}

func flattenSwitchGlobalFortilinkHeartbeatTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalMclagSplitBrainPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalLogMacLimitViolations(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalTrunkHashUnkunicastSrcDst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPoePowerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalIpMacBinding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalVxlanStpVirtualRoot(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalBpduLearn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalVxlanDport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalMacAddressAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalVirtualWireTpid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalTrunkHashUnicastSrcPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalAutoFortilinkDiscovery(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalMaxPathInEcmpGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalTrunkHashMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalFortilinkP2PTpid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalFloodVtp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalMacViolationTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalMirrorQos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalMclagPortBase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalFlapguardRetainTrigger(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalMclagSplitBrainAllPortsDown(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalMacAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalMclagStpAware(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalFortilinkP2PNativeVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalFortilinkVlanOptimization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalVxlanStpVirtualMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalLoopGuardTxInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalAutoIslPortGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalVxlanPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPoeAlarmThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPoeGuardBand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalAutoStpPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPoePowerBudget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalVxlanSport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalMclagPeerInfoTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalMclagIgmpsnoopingAware(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalAccessVlanMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPoePreStandardDetect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalDhcpSnoopingDatabaseExport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalAutoIsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalDmiGlobalAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalL2MemoryCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalMclagSplitBrainDetect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalFortiTrunkDmac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalMacAgingInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPortSecurity(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mac_called_station_delimiter"
	if _, ok := i["mac-called-station-delimiter"]; ok {

		result["mac_called_station_delimiter"] = flattenSwitchGlobalPortSecurityMacCalledStationDelimiter(i["mac-called-station-delimiter"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mab_reauth"
	if _, ok := i["mab-reauth"]; ok {

		result["mab_reauth"] = flattenSwitchGlobalPortSecurityMabReauth(i["mab-reauth"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_reauth_attempt"
	if _, ok := i["max-reauth-attempt"]; ok {

		result["max_reauth_attempt"] = flattenSwitchGlobalPortSecurityMaxReauthAttempt(i["max-reauth-attempt"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mac_case"
	if _, ok := i["mac-case"]; ok {

		result["mac_case"] = flattenSwitchGlobalPortSecurityMacCase(i["mac-case"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mac_password_delimiter"
	if _, ok := i["mac-password-delimiter"]; ok {

		result["mac_password_delimiter"] = flattenSwitchGlobalPortSecurityMacPasswordDelimiter(i["mac-password-delimiter"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "link_down_auth"
	if _, ok := i["link-down-auth"]; ok {

		result["link_down_auth"] = flattenSwitchGlobalPortSecurityLinkDownAuth(i["link-down-auth"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mab_entry_as"
	if _, ok := i["mab-entry-as"]; ok {

		result["mab_entry_as"] = flattenSwitchGlobalPortSecurityMabEntryAs(i["mab-entry-as"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mac_username_delimiter"
	if _, ok := i["mac-username-delimiter"]; ok {

		result["mac_username_delimiter"] = flattenSwitchGlobalPortSecurityMacUsernameDelimiter(i["mac-username-delimiter"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "reauth_period"
	if _, ok := i["reauth-period"]; ok {

		result["reauth_period"] = flattenSwitchGlobalPortSecurityReauthPeriod(i["reauth-period"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mac_calling_station_delimiter"
	if _, ok := i["mac-calling-station-delimiter"]; ok {

		result["mac_calling_station_delimiter"] = flattenSwitchGlobalPortSecurityMacCallingStationDelimiter(i["mac-calling-station-delimiter"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "quarantine_vlan"
	if _, ok := i["quarantine-vlan"]; ok {

		result["quarantine_vlan"] = flattenSwitchGlobalPortSecurityQuarantineVlan(i["quarantine-vlan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tx_period"
	if _, ok := i["tx-period"]; ok {

		result["tx_period"] = flattenSwitchGlobalPortSecurityTxPeriod(i["tx-period"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchGlobalPortSecurityMacCalledStationDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPortSecurityMabReauth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPortSecurityMaxReauthAttempt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPortSecurityMacCase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPortSecurityMacPasswordDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPortSecurityLinkDownAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPortSecurityMabEntryAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPortSecurityMacUsernameDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPortSecurityReauthPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPortSecurityMacCallingStationDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPortSecurityQuarantineVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalPortSecurityTxPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalFloodUnknownMulticast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchGlobalL2MemoryCheckInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchGlobal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fortilink_heartbeat_timeout", flattenSwitchGlobalFortilinkHeartbeatTimeout(o["fortilink-heartbeat-timeout"], d, "fortilink_heartbeat_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["fortilink-heartbeat-timeout"]) {
			return fmt.Errorf("Error reading fortilink_heartbeat_timeout: %v", err)
		}
	}

	if err = d.Set("mclag_split_brain_priority", flattenSwitchGlobalMclagSplitBrainPriority(o["mclag-split-brain-priority"], d, "mclag_split_brain_priority", sv)); err != nil {
		if !fortiAPIPatch(o["mclag-split-brain-priority"]) {
			return fmt.Errorf("Error reading mclag_split_brain_priority: %v", err)
		}
	}

	if err = d.Set("log_mac_limit_violations", flattenSwitchGlobalLogMacLimitViolations(o["log-mac-limit-violations"], d, "log_mac_limit_violations", sv)); err != nil {
		if !fortiAPIPatch(o["log-mac-limit-violations"]) {
			return fmt.Errorf("Error reading log_mac_limit_violations: %v", err)
		}
	}

	if err = d.Set("trunk_hash_unkunicast_src_dst", flattenSwitchGlobalTrunkHashUnkunicastSrcDst(o["trunk-hash-unkunicast-src-dst"], d, "trunk_hash_unkunicast_src_dst", sv)); err != nil {
		if !fortiAPIPatch(o["trunk-hash-unkunicast-src-dst"]) {
			return fmt.Errorf("Error reading trunk_hash_unkunicast_src_dst: %v", err)
		}
	}

	if err = d.Set("poe_power_mode", flattenSwitchGlobalPoePowerMode(o["poe-power-mode"], d, "poe_power_mode", sv)); err != nil {
		if !fortiAPIPatch(o["poe-power-mode"]) {
			return fmt.Errorf("Error reading poe_power_mode: %v", err)
		}
	}

	if err = d.Set("ip_mac_binding", flattenSwitchGlobalIpMacBinding(o["ip-mac-binding"], d, "ip_mac_binding", sv)); err != nil {
		if !fortiAPIPatch(o["ip-mac-binding"]) {
			return fmt.Errorf("Error reading ip_mac_binding: %v", err)
		}
	}

	if err = d.Set("vxlan_stp_virtual_root", flattenSwitchGlobalVxlanStpVirtualRoot(o["vxlan-stp-virtual-root"], d, "vxlan_stp_virtual_root", sv)); err != nil {
		if !fortiAPIPatch(o["vxlan-stp-virtual-root"]) {
			return fmt.Errorf("Error reading vxlan_stp_virtual_root: %v", err)
		}
	}

	if err = d.Set("bpdu_learn", flattenSwitchGlobalBpduLearn(o["bpdu-learn"], d, "bpdu_learn", sv)); err != nil {
		if !fortiAPIPatch(o["bpdu-learn"]) {
			return fmt.Errorf("Error reading bpdu_learn: %v", err)
		}
	}

	if err = d.Set("vxlan_dport", flattenSwitchGlobalVxlanDport(o["vxlan-dport"], d, "vxlan_dport", sv)); err != nil {
		if !fortiAPIPatch(o["vxlan-dport"]) {
			return fmt.Errorf("Error reading vxlan_dport: %v", err)
		}
	}

	if err = d.Set("mac_address_algorithm", flattenSwitchGlobalMacAddressAlgorithm(o["mac-address-algorithm"], d, "mac_address_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["mac-address-algorithm"]) {
			return fmt.Errorf("Error reading mac_address_algorithm: %v", err)
		}
	}

	if err = d.Set("virtual_wire_tpid", flattenSwitchGlobalVirtualWireTpid(o["virtual-wire-tpid"], d, "virtual_wire_tpid", sv)); err != nil {
		if !fortiAPIPatch(o["virtual-wire-tpid"]) {
			return fmt.Errorf("Error reading virtual_wire_tpid: %v", err)
		}
	}

	if err = d.Set("trunk_hash_unicast_src_port", flattenSwitchGlobalTrunkHashUnicastSrcPort(o["trunk-hash-unicast-src-port"], d, "trunk_hash_unicast_src_port", sv)); err != nil {
		if !fortiAPIPatch(o["trunk-hash-unicast-src-port"]) {
			return fmt.Errorf("Error reading trunk_hash_unicast_src_port: %v", err)
		}
	}

	if err = d.Set("auto_fortilink_discovery", flattenSwitchGlobalAutoFortilinkDiscovery(o["auto-fortilink-discovery"], d, "auto_fortilink_discovery", sv)); err != nil {
		if !fortiAPIPatch(o["auto-fortilink-discovery"]) {
			return fmt.Errorf("Error reading auto_fortilink_discovery: %v", err)
		}
	}

	if err = d.Set("max_path_in_ecmp_group", flattenSwitchGlobalMaxPathInEcmpGroup(o["max-path-in-ecmp-group"], d, "max_path_in_ecmp_group", sv)); err != nil {
		if !fortiAPIPatch(o["max-path-in-ecmp-group"]) {
			return fmt.Errorf("Error reading max_path_in_ecmp_group: %v", err)
		}
	}

	if err = d.Set("trunk_hash_mode", flattenSwitchGlobalTrunkHashMode(o["trunk-hash-mode"], d, "trunk_hash_mode", sv)); err != nil {
		if !fortiAPIPatch(o["trunk-hash-mode"]) {
			return fmt.Errorf("Error reading trunk_hash_mode: %v", err)
		}
	}

	if err = d.Set("fortilink_p2p_tpid", flattenSwitchGlobalFortilinkP2PTpid(o["fortilink-p2p-tpid"], d, "fortilink_p2p_tpid", sv)); err != nil {
		if !fortiAPIPatch(o["fortilink-p2p-tpid"]) {
			return fmt.Errorf("Error reading fortilink_p2p_tpid: %v", err)
		}
	}

	if err = d.Set("flood_vtp", flattenSwitchGlobalFloodVtp(o["flood-vtp"], d, "flood_vtp", sv)); err != nil {
		if !fortiAPIPatch(o["flood-vtp"]) {
			return fmt.Errorf("Error reading flood_vtp: %v", err)
		}
	}

	if err = d.Set("mac_violation_timer", flattenSwitchGlobalMacViolationTimer(o["mac-violation-timer"], d, "mac_violation_timer", sv)); err != nil {
		if !fortiAPIPatch(o["mac-violation-timer"]) {
			return fmt.Errorf("Error reading mac_violation_timer: %v", err)
		}
	}

	if err = d.Set("mirror_qos", flattenSwitchGlobalMirrorQos(o["mirror-qos"], d, "mirror_qos", sv)); err != nil {
		if !fortiAPIPatch(o["mirror-qos"]) {
			return fmt.Errorf("Error reading mirror_qos: %v", err)
		}
	}

	if err = d.Set("mclag_port_base", flattenSwitchGlobalMclagPortBase(o["mclag-port-base"], d, "mclag_port_base", sv)); err != nil {
		if !fortiAPIPatch(o["mclag-port-base"]) {
			return fmt.Errorf("Error reading mclag_port_base: %v", err)
		}
	}

	if err = d.Set("flapguard_retain_trigger", flattenSwitchGlobalFlapguardRetainTrigger(o["flapguard-retain-trigger"], d, "flapguard_retain_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["flapguard-retain-trigger"]) {
			return fmt.Errorf("Error reading flapguard_retain_trigger: %v", err)
		}
	}

	if err = d.Set("mclag_split_brain_all_ports_down", flattenSwitchGlobalMclagSplitBrainAllPortsDown(o["mclag-split-brain-all-ports-down"], d, "mclag_split_brain_all_ports_down", sv)); err != nil {
		if !fortiAPIPatch(o["mclag-split-brain-all-ports-down"]) {
			return fmt.Errorf("Error reading mclag_split_brain_all_ports_down: %v", err)
		}
	}

	if err = d.Set("mac_address", flattenSwitchGlobalMacAddress(o["mac-address"], d, "mac_address", sv)); err != nil {
		if !fortiAPIPatch(o["mac-address"]) {
			return fmt.Errorf("Error reading mac_address: %v", err)
		}
	}

	if err = d.Set("mclag_stp_aware", flattenSwitchGlobalMclagStpAware(o["mclag-stp-aware"], d, "mclag_stp_aware", sv)); err != nil {
		if !fortiAPIPatch(o["mclag-stp-aware"]) {
			return fmt.Errorf("Error reading mclag_stp_aware: %v", err)
		}
	}

	if err = d.Set("fortilink_p2p_native_vlan", flattenSwitchGlobalFortilinkP2PNativeVlan(o["fortilink-p2p-native-vlan"], d, "fortilink_p2p_native_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["fortilink-p2p-native-vlan"]) {
			return fmt.Errorf("Error reading fortilink_p2p_native_vlan: %v", err)
		}
	}

	if err = d.Set("fortilink_vlan_optimization", flattenSwitchGlobalFortilinkVlanOptimization(o["fortilink-vlan-optimization"], d, "fortilink_vlan_optimization", sv)); err != nil {
		if !fortiAPIPatch(o["fortilink-vlan-optimization"]) {
			return fmt.Errorf("Error reading fortilink_vlan_optimization: %v", err)
		}
	}

	if err = d.Set("vxlan_stp_virtual_mac", flattenSwitchGlobalVxlanStpVirtualMac(o["vxlan-stp-virtual-mac"], d, "vxlan_stp_virtual_mac", sv)); err != nil {
		if !fortiAPIPatch(o["vxlan-stp-virtual-mac"]) {
			return fmt.Errorf("Error reading vxlan_stp_virtual_mac: %v", err)
		}
	}

	if err = d.Set("loop_guard_tx_interval", flattenSwitchGlobalLoopGuardTxInterval(o["loop-guard-tx-interval"], d, "loop_guard_tx_interval", sv)); err != nil {
		if !fortiAPIPatch(o["loop-guard-tx-interval"]) {
			return fmt.Errorf("Error reading loop_guard_tx_interval: %v", err)
		}
	}

	if err = d.Set("auto_isl_port_group", flattenSwitchGlobalAutoIslPortGroup(o["auto-isl-port-group"], d, "auto_isl_port_group", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl-port-group"]) {
			return fmt.Errorf("Error reading auto_isl_port_group: %v", err)
		}
	}

	if err = d.Set("vxlan_port", flattenSwitchGlobalVxlanPort(o["vxlan-port"], d, "vxlan_port", sv)); err != nil {
		if !fortiAPIPatch(o["vxlan-port"]) {
			return fmt.Errorf("Error reading vxlan_port: %v", err)
		}
	}

	if err = d.Set("poe_alarm_threshold", flattenSwitchGlobalPoeAlarmThreshold(o["poe-alarm-threshold"], d, "poe_alarm_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["poe-alarm-threshold"]) {
			return fmt.Errorf("Error reading poe_alarm_threshold: %v", err)
		}
	}

	if err = d.Set("poe_guard_band", flattenSwitchGlobalPoeGuardBand(o["poe-guard-band"], d, "poe_guard_band", sv)); err != nil {
		if !fortiAPIPatch(o["poe-guard-band"]) {
			return fmt.Errorf("Error reading poe_guard_band: %v", err)
		}
	}

	if err = d.Set("auto_stp_priority", flattenSwitchGlobalAutoStpPriority(o["auto-stp-priority"], d, "auto_stp_priority", sv)); err != nil {
		if !fortiAPIPatch(o["auto-stp-priority"]) {
			return fmt.Errorf("Error reading auto_stp_priority: %v", err)
		}
	}

	if err = d.Set("poe_power_budget", flattenSwitchGlobalPoePowerBudget(o["poe-power-budget"], d, "poe_power_budget", sv)); err != nil {
		if !fortiAPIPatch(o["poe-power-budget"]) {
			return fmt.Errorf("Error reading poe_power_budget: %v", err)
		}
	}

	if err = d.Set("vxlan_sport", flattenSwitchGlobalVxlanSport(o["vxlan-sport"], d, "vxlan_sport", sv)); err != nil {
		if !fortiAPIPatch(o["vxlan-sport"]) {
			return fmt.Errorf("Error reading vxlan_sport: %v", err)
		}
	}

	if err = d.Set("mclag_peer_info_timeout", flattenSwitchGlobalMclagPeerInfoTimeout(o["mclag-peer-info-timeout"], d, "mclag_peer_info_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["mclag-peer-info-timeout"]) {
			return fmt.Errorf("Error reading mclag_peer_info_timeout: %v", err)
		}
	}

	if err = d.Set("mclag_igmpsnooping_aware", flattenSwitchGlobalMclagIgmpsnoopingAware(o["mclag-igmpsnooping-aware"], d, "mclag_igmpsnooping_aware", sv)); err != nil {
		if !fortiAPIPatch(o["mclag-igmpsnooping-aware"]) {
			return fmt.Errorf("Error reading mclag_igmpsnooping_aware: %v", err)
		}
	}

	if err = d.Set("access_vlan_mode", flattenSwitchGlobalAccessVlanMode(o["access-vlan-mode"], d, "access_vlan_mode", sv)); err != nil {
		if !fortiAPIPatch(o["access-vlan-mode"]) {
			return fmt.Errorf("Error reading access_vlan_mode: %v", err)
		}
	}

	if err = d.Set("poe_pre_standard_detect", flattenSwitchGlobalPoePreStandardDetect(o["poe-pre-standard-detect"], d, "poe_pre_standard_detect", sv)); err != nil {
		if !fortiAPIPatch(o["poe-pre-standard-detect"]) {
			return fmt.Errorf("Error reading poe_pre_standard_detect: %v", err)
		}
	}

	if err = d.Set("dhcp_snooping_database_export", flattenSwitchGlobalDhcpSnoopingDatabaseExport(o["dhcp-snooping-database-export"], d, "dhcp_snooping_database_export", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-snooping-database-export"]) {
			return fmt.Errorf("Error reading dhcp_snooping_database_export: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchGlobalName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("auto_isl", flattenSwitchGlobalAutoIsl(o["auto-isl"], d, "auto_isl", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl"]) {
			return fmt.Errorf("Error reading auto_isl: %v", err)
		}
	}

	if err = d.Set("dmi_global_all", flattenSwitchGlobalDmiGlobalAll(o["dmi-global-all"], d, "dmi_global_all", sv)); err != nil {
		if !fortiAPIPatch(o["dmi-global-all"]) {
			return fmt.Errorf("Error reading dmi_global_all: %v", err)
		}
	}

	if err = d.Set("l2_memory_check", flattenSwitchGlobalL2MemoryCheck(o["l2-memory-check"], d, "l2_memory_check", sv)); err != nil {
		if !fortiAPIPatch(o["l2-memory-check"]) {
			return fmt.Errorf("Error reading l2_memory_check: %v", err)
		}
	}

	if err = d.Set("mclag_split_brain_detect", flattenSwitchGlobalMclagSplitBrainDetect(o["mclag-split-brain-detect"], d, "mclag_split_brain_detect", sv)); err != nil {
		if !fortiAPIPatch(o["mclag-split-brain-detect"]) {
			return fmt.Errorf("Error reading mclag_split_brain_detect: %v", err)
		}
	}

	if err = d.Set("forti_trunk_dmac", flattenSwitchGlobalFortiTrunkDmac(o["forti-trunk-dmac"], d, "forti_trunk_dmac", sv)); err != nil {
		if !fortiAPIPatch(o["forti-trunk-dmac"]) {
			return fmt.Errorf("Error reading forti_trunk_dmac: %v", err)
		}
	}

	if err = d.Set("mac_aging_interval", flattenSwitchGlobalMacAgingInterval(o["mac-aging-interval"], d, "mac_aging_interval", sv)); err != nil {
		if !fortiAPIPatch(o["mac-aging-interval"]) {
			return fmt.Errorf("Error reading mac_aging_interval: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("port_security", flattenSwitchGlobalPortSecurity(o["port-security"], d, "port_security", sv)); err != nil {
			if !fortiAPIPatch(o["port-security"]) {
				return fmt.Errorf("Error reading port_security: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("port_security"); ok {
			if err = d.Set("port_security", flattenSwitchGlobalPortSecurity(o["port-security"], d, "port_security", sv)); err != nil {
				if !fortiAPIPatch(o["port-security"]) {
					return fmt.Errorf("Error reading port_security: %v", err)
				}
			}
		}
	}

	if err = d.Set("flood_unknown_multicast", flattenSwitchGlobalFloodUnknownMulticast(o["flood-unknown-multicast"], d, "flood_unknown_multicast", sv)); err != nil {
		if !fortiAPIPatch(o["flood-unknown-multicast"]) {
			return fmt.Errorf("Error reading flood_unknown_multicast: %v", err)
		}
	}

	if err = d.Set("l2_memory_check_interval", flattenSwitchGlobalL2MemoryCheckInterval(o["l2-memory-check-interval"], d, "l2_memory_check_interval", sv)); err != nil {
		if !fortiAPIPatch(o["l2-memory-check-interval"]) {
			return fmt.Errorf("Error reading l2_memory_check_interval: %v", err)
		}
	}

	return nil
}

func flattenSwitchGlobalFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchGlobalFortilinkHeartbeatTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalMclagSplitBrainPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalLogMacLimitViolations(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalTrunkHashUnkunicastSrcDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPoePowerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalIpMacBinding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalVxlanStpVirtualRoot(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalBpduLearn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalVxlanDport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalMacAddressAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalVirtualWireTpid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalTrunkHashUnicastSrcPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalAutoFortilinkDiscovery(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalMaxPathInEcmpGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalTrunkHashMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalFortilinkP2PTpid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalFloodVtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalMacViolationTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalMirrorQos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalMclagPortBase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalFlapguardRetainTrigger(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalMclagSplitBrainAllPortsDown(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalMacAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalMclagStpAware(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalFortilinkP2PNativeVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalFortilinkVlanOptimization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalVxlanStpVirtualMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalLoopGuardTxInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalAutoIslPortGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalVxlanPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPoeAlarmThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPoeGuardBand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalAutoStpPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPoePowerBudget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalVxlanSport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalMclagPeerInfoTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalMclagIgmpsnoopingAware(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalAccessVlanMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPoePreStandardDetect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalDhcpSnoopingDatabaseExport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalAutoIsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalDmiGlobalAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalL2MemoryCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalMclagSplitBrainDetect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalFortiTrunkDmac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalMacAgingInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPortSecurity(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mac_called_station_delimiter"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mac-called-station-delimiter"] = nil
		} else {

			result["mac-called-station-delimiter"], _ = expandSwitchGlobalPortSecurityMacCalledStationDelimiter(d, i["mac_called_station_delimiter"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "mab_reauth"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mab-reauth"] = nil
		} else {

			result["mab-reauth"], _ = expandSwitchGlobalPortSecurityMabReauth(d, i["mab_reauth"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "max_reauth_attempt"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["max-reauth-attempt"] = nil
		} else {

			result["max-reauth-attempt"], _ = expandSwitchGlobalPortSecurityMaxReauthAttempt(d, i["max_reauth_attempt"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "mac_case"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mac-case"] = nil
		} else {

			result["mac-case"], _ = expandSwitchGlobalPortSecurityMacCase(d, i["mac_case"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "mac_password_delimiter"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mac-password-delimiter"] = nil
		} else {

			result["mac-password-delimiter"], _ = expandSwitchGlobalPortSecurityMacPasswordDelimiter(d, i["mac_password_delimiter"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "link_down_auth"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["link-down-auth"] = nil
		} else {

			result["link-down-auth"], _ = expandSwitchGlobalPortSecurityLinkDownAuth(d, i["link_down_auth"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "mab_entry_as"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mab-entry-as"] = nil
		} else {

			result["mab-entry-as"], _ = expandSwitchGlobalPortSecurityMabEntryAs(d, i["mab_entry_as"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "mac_username_delimiter"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mac-username-delimiter"] = nil
		} else {

			result["mac-username-delimiter"], _ = expandSwitchGlobalPortSecurityMacUsernameDelimiter(d, i["mac_username_delimiter"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "reauth_period"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["reauth-period"] = nil
		} else {

			result["reauth-period"], _ = expandSwitchGlobalPortSecurityReauthPeriod(d, i["reauth_period"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "mac_calling_station_delimiter"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mac-calling-station-delimiter"] = nil
		} else {

			result["mac-calling-station-delimiter"], _ = expandSwitchGlobalPortSecurityMacCallingStationDelimiter(d, i["mac_calling_station_delimiter"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "quarantine_vlan"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["quarantine-vlan"] = nil
		} else {

			result["quarantine-vlan"], _ = expandSwitchGlobalPortSecurityQuarantineVlan(d, i["quarantine_vlan"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "tx_period"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["tx-period"] = nil
		} else {

			result["tx-period"], _ = expandSwitchGlobalPortSecurityTxPeriod(d, i["tx_period"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSwitchGlobalPortSecurityMacCalledStationDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPortSecurityMabReauth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPortSecurityMaxReauthAttempt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPortSecurityMacCase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPortSecurityMacPasswordDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPortSecurityLinkDownAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPortSecurityMabEntryAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPortSecurityMacUsernameDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPortSecurityReauthPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPortSecurityMacCallingStationDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPortSecurityQuarantineVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalPortSecurityTxPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalFloodUnknownMulticast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchGlobalL2MemoryCheckInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchGlobal(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fortilink_heartbeat_timeout"); ok {
		if setArgNil {
			obj["fortilink-heartbeat-timeout"] = nil
		} else {

			t, err := expandSwitchGlobalFortilinkHeartbeatTimeout(d, v, "fortilink_heartbeat_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fortilink-heartbeat-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("mclag_split_brain_priority"); ok {
		if setArgNil {
			obj["mclag-split-brain-priority"] = nil
		} else {

			t, err := expandSwitchGlobalMclagSplitBrainPriority(d, v, "mclag_split_brain_priority", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mclag-split-brain-priority"] = t
			}
		}
	}

	if v, ok := d.GetOk("log_mac_limit_violations"); ok {
		if setArgNil {
			obj["log-mac-limit-violations"] = nil
		} else {

			t, err := expandSwitchGlobalLogMacLimitViolations(d, v, "log_mac_limit_violations", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-mac-limit-violations"] = t
			}
		}
	}

	if v, ok := d.GetOk("trunk_hash_unkunicast_src_dst"); ok {
		if setArgNil {
			obj["trunk-hash-unkunicast-src-dst"] = nil
		} else {

			t, err := expandSwitchGlobalTrunkHashUnkunicastSrcDst(d, v, "trunk_hash_unkunicast_src_dst", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trunk-hash-unkunicast-src-dst"] = t
			}
		}
	}

	if v, ok := d.GetOk("poe_power_mode"); ok {
		if setArgNil {
			obj["poe-power-mode"] = nil
		} else {

			t, err := expandSwitchGlobalPoePowerMode(d, v, "poe_power_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["poe-power-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip_mac_binding"); ok {
		if setArgNil {
			obj["ip-mac-binding"] = nil
		} else {

			t, err := expandSwitchGlobalIpMacBinding(d, v, "ip_mac_binding", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip-mac-binding"] = t
			}
		}
	}

	if v, ok := d.GetOk("vxlan_stp_virtual_root"); ok {
		if setArgNil {
			obj["vxlan-stp-virtual-root"] = nil
		} else {

			t, err := expandSwitchGlobalVxlanStpVirtualRoot(d, v, "vxlan_stp_virtual_root", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vxlan-stp-virtual-root"] = t
			}
		}
	}

	if v, ok := d.GetOk("bpdu_learn"); ok {
		if setArgNil {
			obj["bpdu-learn"] = nil
		} else {

			t, err := expandSwitchGlobalBpduLearn(d, v, "bpdu_learn", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bpdu-learn"] = t
			}
		}
	}

	if v, ok := d.GetOk("vxlan_dport"); ok {
		if setArgNil {
			obj["vxlan-dport"] = nil
		} else {

			t, err := expandSwitchGlobalVxlanDport(d, v, "vxlan_dport", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vxlan-dport"] = t
			}
		}
	}

	if v, ok := d.GetOk("mac_address_algorithm"); ok {
		if setArgNil {
			obj["mac-address-algorithm"] = nil
		} else {

			t, err := expandSwitchGlobalMacAddressAlgorithm(d, v, "mac_address_algorithm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mac-address-algorithm"] = t
			}
		}
	}

	if v, ok := d.GetOk("virtual_wire_tpid"); ok {
		if setArgNil {
			obj["virtual-wire-tpid"] = nil
		} else {

			t, err := expandSwitchGlobalVirtualWireTpid(d, v, "virtual_wire_tpid", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["virtual-wire-tpid"] = t
			}
		}
	}

	if v, ok := d.GetOk("trunk_hash_unicast_src_port"); ok {
		if setArgNil {
			obj["trunk-hash-unicast-src-port"] = nil
		} else {

			t, err := expandSwitchGlobalTrunkHashUnicastSrcPort(d, v, "trunk_hash_unicast_src_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trunk-hash-unicast-src-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_fortilink_discovery"); ok {
		if setArgNil {
			obj["auto-fortilink-discovery"] = nil
		} else {

			t, err := expandSwitchGlobalAutoFortilinkDiscovery(d, v, "auto_fortilink_discovery", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-fortilink-discovery"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_path_in_ecmp_group"); ok {
		if setArgNil {
			obj["max-path-in-ecmp-group"] = nil
		} else {

			t, err := expandSwitchGlobalMaxPathInEcmpGroup(d, v, "max_path_in_ecmp_group", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-path-in-ecmp-group"] = t
			}
		}
	}

	if v, ok := d.GetOk("trunk_hash_mode"); ok {
		if setArgNil {
			obj["trunk-hash-mode"] = nil
		} else {

			t, err := expandSwitchGlobalTrunkHashMode(d, v, "trunk_hash_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trunk-hash-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("fortilink_p2p_tpid"); ok {
		if setArgNil {
			obj["fortilink-p2p-tpid"] = nil
		} else {

			t, err := expandSwitchGlobalFortilinkP2PTpid(d, v, "fortilink_p2p_tpid", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fortilink-p2p-tpid"] = t
			}
		}
	}

	if v, ok := d.GetOk("flood_vtp"); ok {
		if setArgNil {
			obj["flood-vtp"] = nil
		} else {

			t, err := expandSwitchGlobalFloodVtp(d, v, "flood_vtp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["flood-vtp"] = t
			}
		}
	}

	if v, ok := d.GetOk("mac_violation_timer"); ok {
		if setArgNil {
			obj["mac-violation-timer"] = nil
		} else {

			t, err := expandSwitchGlobalMacViolationTimer(d, v, "mac_violation_timer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mac-violation-timer"] = t
			}
		}
	}

	if v, ok := d.GetOk("mirror_qos"); ok {
		if setArgNil {
			obj["mirror-qos"] = nil
		} else {

			t, err := expandSwitchGlobalMirrorQos(d, v, "mirror_qos", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mirror-qos"] = t
			}
		}
	}

	if v, ok := d.GetOk("mclag_port_base"); ok {
		if setArgNil {
			obj["mclag-port-base"] = nil
		} else {

			t, err := expandSwitchGlobalMclagPortBase(d, v, "mclag_port_base", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mclag-port-base"] = t
			}
		}
	}

	if v, ok := d.GetOk("flapguard_retain_trigger"); ok {
		if setArgNil {
			obj["flapguard-retain-trigger"] = nil
		} else {

			t, err := expandSwitchGlobalFlapguardRetainTrigger(d, v, "flapguard_retain_trigger", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["flapguard-retain-trigger"] = t
			}
		}
	}

	if v, ok := d.GetOk("mclag_split_brain_all_ports_down"); ok {
		if setArgNil {
			obj["mclag-split-brain-all-ports-down"] = nil
		} else {

			t, err := expandSwitchGlobalMclagSplitBrainAllPortsDown(d, v, "mclag_split_brain_all_ports_down", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mclag-split-brain-all-ports-down"] = t
			}
		}
	}

	if v, ok := d.GetOk("mac_address"); ok {
		if setArgNil {
			obj["mac-address"] = nil
		} else {

			t, err := expandSwitchGlobalMacAddress(d, v, "mac_address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mac-address"] = t
			}
		}
	}

	if v, ok := d.GetOk("mclag_stp_aware"); ok {
		if setArgNil {
			obj["mclag-stp-aware"] = nil
		} else {

			t, err := expandSwitchGlobalMclagStpAware(d, v, "mclag_stp_aware", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mclag-stp-aware"] = t
			}
		}
	}

	if v, ok := d.GetOk("fortilink_p2p_native_vlan"); ok {
		if setArgNil {
			obj["fortilink-p2p-native-vlan"] = nil
		} else {

			t, err := expandSwitchGlobalFortilinkP2PNativeVlan(d, v, "fortilink_p2p_native_vlan", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fortilink-p2p-native-vlan"] = t
			}
		}
	}

	if v, ok := d.GetOk("fortilink_vlan_optimization"); ok {
		if setArgNil {
			obj["fortilink-vlan-optimization"] = nil
		} else {

			t, err := expandSwitchGlobalFortilinkVlanOptimization(d, v, "fortilink_vlan_optimization", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fortilink-vlan-optimization"] = t
			}
		}
	}

	if v, ok := d.GetOk("vxlan_stp_virtual_mac"); ok {
		if setArgNil {
			obj["vxlan-stp-virtual-mac"] = nil
		} else {

			t, err := expandSwitchGlobalVxlanStpVirtualMac(d, v, "vxlan_stp_virtual_mac", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vxlan-stp-virtual-mac"] = t
			}
		}
	}

	if v, ok := d.GetOk("loop_guard_tx_interval"); ok {
		if setArgNil {
			obj["loop-guard-tx-interval"] = nil
		} else {

			t, err := expandSwitchGlobalLoopGuardTxInterval(d, v, "loop_guard_tx_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["loop-guard-tx-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_isl_port_group"); ok {
		if setArgNil {
			obj["auto-isl-port-group"] = nil
		} else {

			t, err := expandSwitchGlobalAutoIslPortGroup(d, v, "auto_isl_port_group", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-isl-port-group"] = t
			}
		}
	}

	if v, ok := d.GetOk("vxlan_port"); ok {
		if setArgNil {
			obj["vxlan-port"] = nil
		} else {

			t, err := expandSwitchGlobalVxlanPort(d, v, "vxlan_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vxlan-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("poe_alarm_threshold"); ok {
		if setArgNil {
			obj["poe-alarm-threshold"] = nil
		} else {

			t, err := expandSwitchGlobalPoeAlarmThreshold(d, v, "poe_alarm_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["poe-alarm-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("poe_guard_band"); ok {
		if setArgNil {
			obj["poe-guard-band"] = nil
		} else {

			t, err := expandSwitchGlobalPoeGuardBand(d, v, "poe_guard_band", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["poe-guard-band"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_stp_priority"); ok {
		if setArgNil {
			obj["auto-stp-priority"] = nil
		} else {

			t, err := expandSwitchGlobalAutoStpPriority(d, v, "auto_stp_priority", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-stp-priority"] = t
			}
		}
	}

	if v, ok := d.GetOk("poe_power_budget"); ok {
		if setArgNil {
			obj["poe-power-budget"] = nil
		} else {

			t, err := expandSwitchGlobalPoePowerBudget(d, v, "poe_power_budget", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["poe-power-budget"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("vxlan_sport"); ok {
		if setArgNil {
			obj["vxlan-sport"] = nil
		} else {

			t, err := expandSwitchGlobalVxlanSport(d, v, "vxlan_sport", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vxlan-sport"] = t
			}
		}
	}

	if v, ok := d.GetOk("mclag_peer_info_timeout"); ok {
		if setArgNil {
			obj["mclag-peer-info-timeout"] = nil
		} else {

			t, err := expandSwitchGlobalMclagPeerInfoTimeout(d, v, "mclag_peer_info_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mclag-peer-info-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("mclag_igmpsnooping_aware"); ok {
		if setArgNil {
			obj["mclag-igmpsnooping-aware"] = nil
		} else {

			t, err := expandSwitchGlobalMclagIgmpsnoopingAware(d, v, "mclag_igmpsnooping_aware", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mclag-igmpsnooping-aware"] = t
			}
		}
	}

	if v, ok := d.GetOk("access_vlan_mode"); ok {
		if setArgNil {
			obj["access-vlan-mode"] = nil
		} else {

			t, err := expandSwitchGlobalAccessVlanMode(d, v, "access_vlan_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["access-vlan-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("poe_pre_standard_detect"); ok {
		if setArgNil {
			obj["poe-pre-standard-detect"] = nil
		} else {

			t, err := expandSwitchGlobalPoePreStandardDetect(d, v, "poe_pre_standard_detect", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["poe-pre-standard-detect"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_snooping_database_export"); ok {
		if setArgNil {
			obj["dhcp-snooping-database-export"] = nil
		} else {

			t, err := expandSwitchGlobalDhcpSnoopingDatabaseExport(d, v, "dhcp_snooping_database_export", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-snooping-database-export"] = t
			}
		}
	}

	if v, ok := d.GetOk("name"); ok {
		if setArgNil {
			obj["name"] = nil
		} else {

			t, err := expandSwitchGlobalName(d, v, "name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["name"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_isl"); ok {
		if setArgNil {
			obj["auto-isl"] = nil
		} else {

			t, err := expandSwitchGlobalAutoIsl(d, v, "auto_isl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-isl"] = t
			}
		}
	}

	if v, ok := d.GetOk("dmi_global_all"); ok {
		if setArgNil {
			obj["dmi-global-all"] = nil
		} else {

			t, err := expandSwitchGlobalDmiGlobalAll(d, v, "dmi_global_all", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dmi-global-all"] = t
			}
		}
	}

	if v, ok := d.GetOk("l2_memory_check"); ok {
		if setArgNil {
			obj["l2-memory-check"] = nil
		} else {

			t, err := expandSwitchGlobalL2MemoryCheck(d, v, "l2_memory_check", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["l2-memory-check"] = t
			}
		}
	}

	if v, ok := d.GetOk("mclag_split_brain_detect"); ok {
		if setArgNil {
			obj["mclag-split-brain-detect"] = nil
		} else {

			t, err := expandSwitchGlobalMclagSplitBrainDetect(d, v, "mclag_split_brain_detect", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mclag-split-brain-detect"] = t
			}
		}
	}

	if v, ok := d.GetOk("forti_trunk_dmac"); ok {
		if setArgNil {
			obj["forti-trunk-dmac"] = nil
		} else {

			t, err := expandSwitchGlobalFortiTrunkDmac(d, v, "forti_trunk_dmac", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["forti-trunk-dmac"] = t
			}
		}
	}

	if v, ok := d.GetOk("mac_aging_interval"); ok {
		if setArgNil {
			obj["mac-aging-interval"] = nil
		} else {

			t, err := expandSwitchGlobalMacAgingInterval(d, v, "mac_aging_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mac-aging-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("port_security"); ok {

		t, err := expandSwitchGlobalPortSecurity(d, v, "port_security", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-security"] = t
		}
	}

	if v, ok := d.GetOk("flood_unknown_multicast"); ok {
		if setArgNil {
			obj["flood-unknown-multicast"] = nil
		} else {

			t, err := expandSwitchGlobalFloodUnknownMulticast(d, v, "flood_unknown_multicast", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["flood-unknown-multicast"] = t
			}
		}
	}

	if v, ok := d.GetOk("l2_memory_check_interval"); ok {
		if setArgNil {
			obj["l2-memory-check-interval"] = nil
		} else {

			t, err := expandSwitchGlobalL2MemoryCheckInterval(d, v, "l2_memory_check_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["l2-memory-check-interval"] = t
			}
		}
	}

	return &obj, nil
}
