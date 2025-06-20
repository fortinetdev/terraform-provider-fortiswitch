// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Physical port specific configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchPhysicalPort() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchPhysicalPortCreate,
		Read:   resourceSwitchPhysicalPortRead,
		Update: resourceSwitchPhysicalPortUpdate,
		Delete: resourceSwitchPhysicalPortDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"macsec_pae_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2_learning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dmi_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storm_control_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"speed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortilink_p2p": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flap_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"macsec_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"egress_drop_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loopback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_based_flow_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owning_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"energy_efficient_ethernet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qsfp_low_power_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"poe_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"medium": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cdp_status": &schema.Schema{
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
			"flapguard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flap_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eee_tx_wake_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"storm_control": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"broadcast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unknown_unicast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unknown_multicast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"burst_size_level": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flapguard_state": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 23),
				Optional:     true,
				Computed:     true,
			},
			"pause_resume": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2_sa_unknown": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lldp_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eee_tx_idle_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"lldp_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"flap_trig": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"link_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flap_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"poe_port_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"poe_max_power_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_frame_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"flow_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pause_meter_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"poe_port_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchPhysicalPortCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchPhysicalPort(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchPhysicalPort resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchPhysicalPort(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchPhysicalPort resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchPhysicalPort")
	}

	return resourceSwitchPhysicalPortRead(d, m)
}

func resourceSwitchPhysicalPortUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchPhysicalPort(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchPhysicalPort resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchPhysicalPort(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchPhysicalPort resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchPhysicalPort")
	}

	return resourceSwitchPhysicalPortRead(d, m)
}

func resourceSwitchPhysicalPortDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchPhysicalPort(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchPhysicalPort resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchPhysicalPortRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchPhysicalPort(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchPhysicalPort resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchPhysicalPort(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchPhysicalPort resource from API: %v", err)
	}
	return nil
}

func flattenSwitchPhysicalPortMacsecPaeMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortL2Learning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortDmiStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortStormControlMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortSpeed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortFortilinkP2P(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortFlapRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortMacsecProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortEgressDropMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortLoopback(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortPriorityBasedFlowControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortOwningInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortEnergyEfficientEthernet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortQsfpLowPowerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortPoeStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortMedium(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortCdpStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortFlapguard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortFlapDuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortEeeTxWakeTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortStormControl(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "broadcast"
	if _, ok := i["broadcast"]; ok {

		result["broadcast"] = flattenSwitchPhysicalPortStormControlBroadcast(i["broadcast"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unknown_unicast"
	if _, ok := i["unknown-unicast"]; ok {

		result["unknown_unicast"] = flattenSwitchPhysicalPortStormControlUnknownUnicast(i["unknown-unicast"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unknown_multicast"
	if _, ok := i["unknown-multicast"]; ok {

		result["unknown_multicast"] = flattenSwitchPhysicalPortStormControlUnknownMulticast(i["unknown-multicast"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rate"
	if _, ok := i["rate"]; ok {

		result["rate"] = flattenSwitchPhysicalPortStormControlRate(i["rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "burst_size_level"
	if _, ok := i["burst-size-level"]; ok {

		result["burst_size_level"] = flattenSwitchPhysicalPortStormControlBurstSizeLevel(i["burst-size-level"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchPhysicalPortStormControlBroadcast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortStormControlUnknownUnicast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortStormControlUnknownMulticast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortStormControlRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortStormControlBurstSizeLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortSecurityMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortFlapguardState(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortPauseResume(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortL2SaUnknown(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortLldpStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortEeeTxIdleTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortLldpProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortFlapTrig(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortLinkStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortFlapTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortPoePortMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortPoeMaxPowerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortPortIndex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortMaxFrameSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortFlowControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortPauseMeterRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhysicalPortPoePortPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchPhysicalPort(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("macsec_pae_mode", flattenSwitchPhysicalPortMacsecPaeMode(o["macsec-pae-mode"], d, "macsec_pae_mode", sv)); err != nil {
		if !fortiAPIPatch(o["macsec-pae-mode"]) {
			return fmt.Errorf("Error reading macsec_pae_mode: %v", err)
		}
	}

	if err = d.Set("l2_learning", flattenSwitchPhysicalPortL2Learning(o["l2-learning"], d, "l2_learning", sv)); err != nil {
		if !fortiAPIPatch(o["l2-learning"]) {
			return fmt.Errorf("Error reading l2_learning: %v", err)
		}
	}

	if err = d.Set("dmi_status", flattenSwitchPhysicalPortDmiStatus(o["dmi-status"], d, "dmi_status", sv)); err != nil {
		if !fortiAPIPatch(o["dmi-status"]) {
			return fmt.Errorf("Error reading dmi_status: %v", err)
		}
	}

	if err = d.Set("storm_control_mode", flattenSwitchPhysicalPortStormControlMode(o["storm-control-mode"], d, "storm_control_mode", sv)); err != nil {
		if !fortiAPIPatch(o["storm-control-mode"]) {
			return fmt.Errorf("Error reading storm_control_mode: %v", err)
		}
	}

	if err = d.Set("speed", flattenSwitchPhysicalPortSpeed(o["speed"], d, "speed", sv)); err != nil {
		if !fortiAPIPatch(o["speed"]) {
			return fmt.Errorf("Error reading speed: %v", err)
		}
	}

	if err = d.Set("fortilink_p2p", flattenSwitchPhysicalPortFortilinkP2P(o["fortilink-p2p"], d, "fortilink_p2p", sv)); err != nil {
		if !fortiAPIPatch(o["fortilink-p2p"]) {
			return fmt.Errorf("Error reading fortilink_p2p: %v", err)
		}
	}

	if err = d.Set("flap_rate", flattenSwitchPhysicalPortFlapRate(o["flap-rate"], d, "flap_rate", sv)); err != nil {
		if !fortiAPIPatch(o["flap-rate"]) {
			return fmt.Errorf("Error reading flap_rate: %v", err)
		}
	}

	if err = d.Set("macsec_profile", flattenSwitchPhysicalPortMacsecProfile(o["macsec-profile"], d, "macsec_profile", sv)); err != nil {
		if !fortiAPIPatch(o["macsec-profile"]) {
			return fmt.Errorf("Error reading macsec_profile: %v", err)
		}
	}

	if err = d.Set("egress_drop_mode", flattenSwitchPhysicalPortEgressDropMode(o["egress-drop-mode"], d, "egress_drop_mode", sv)); err != nil {
		if !fortiAPIPatch(o["egress-drop-mode"]) {
			return fmt.Errorf("Error reading egress_drop_mode: %v", err)
		}
	}

	if err = d.Set("loopback", flattenSwitchPhysicalPortLoopback(o["loopback"], d, "loopback", sv)); err != nil {
		if !fortiAPIPatch(o["loopback"]) {
			return fmt.Errorf("Error reading loopback: %v", err)
		}
	}

	if err = d.Set("priority_based_flow_control", flattenSwitchPhysicalPortPriorityBasedFlowControl(o["priority-based-flow-control"], d, "priority_based_flow_control", sv)); err != nil {
		if !fortiAPIPatch(o["priority-based-flow-control"]) {
			return fmt.Errorf("Error reading priority_based_flow_control: %v", err)
		}
	}

	if err = d.Set("owning_interface", flattenSwitchPhysicalPortOwningInterface(o["owning-interface"], d, "owning_interface", sv)); err != nil {
		if !fortiAPIPatch(o["owning-interface"]) {
			return fmt.Errorf("Error reading owning_interface: %v", err)
		}
	}

	if err = d.Set("energy_efficient_ethernet", flattenSwitchPhysicalPortEnergyEfficientEthernet(o["energy-efficient-ethernet"], d, "energy_efficient_ethernet", sv)); err != nil {
		if !fortiAPIPatch(o["energy-efficient-ethernet"]) {
			return fmt.Errorf("Error reading energy_efficient_ethernet: %v", err)
		}
	}

	if err = d.Set("qsfp_low_power_mode", flattenSwitchPhysicalPortQsfpLowPowerMode(o["qsfp-low-power-mode"], d, "qsfp_low_power_mode", sv)); err != nil {
		if !fortiAPIPatch(o["qsfp-low-power-mode"]) {
			return fmt.Errorf("Error reading qsfp_low_power_mode: %v", err)
		}
	}

	if err = d.Set("poe_status", flattenSwitchPhysicalPortPoeStatus(o["poe-status"], d, "poe_status", sv)); err != nil {
		if !fortiAPIPatch(o["poe-status"]) {
			return fmt.Errorf("Error reading poe_status: %v", err)
		}
	}

	if err = d.Set("status", flattenSwitchPhysicalPortStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("medium", flattenSwitchPhysicalPortMedium(o["medium"], d, "medium", sv)); err != nil {
		if !fortiAPIPatch(o["medium"]) {
			return fmt.Errorf("Error reading medium: %v", err)
		}
	}

	if err = d.Set("cdp_status", flattenSwitchPhysicalPortCdpStatus(o["cdp-status"], d, "cdp_status", sv)); err != nil {
		if !fortiAPIPatch(o["cdp-status"]) {
			return fmt.Errorf("Error reading cdp_status: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchPhysicalPortDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("flapguard", flattenSwitchPhysicalPortFlapguard(o["flapguard"], d, "flapguard", sv)); err != nil {
		if !fortiAPIPatch(o["flapguard"]) {
			return fmt.Errorf("Error reading flapguard: %v", err)
		}
	}

	if err = d.Set("flap_duration", flattenSwitchPhysicalPortFlapDuration(o["flap-duration"], d, "flap_duration", sv)); err != nil {
		if !fortiAPIPatch(o["flap-duration"]) {
			return fmt.Errorf("Error reading flap_duration: %v", err)
		}
	}

	if err = d.Set("eee_tx_wake_time", flattenSwitchPhysicalPortEeeTxWakeTime(o["eee-tx-wake-time"], d, "eee_tx_wake_time", sv)); err != nil {
		if !fortiAPIPatch(o["eee-tx-wake-time"]) {
			return fmt.Errorf("Error reading eee_tx_wake_time: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("storm_control", flattenSwitchPhysicalPortStormControl(o["storm-control"], d, "storm_control", sv)); err != nil {
			if !fortiAPIPatch(o["storm-control"]) {
				return fmt.Errorf("Error reading storm_control: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("storm_control"); ok {
			if err = d.Set("storm_control", flattenSwitchPhysicalPortStormControl(o["storm-control"], d, "storm_control", sv)); err != nil {
				if !fortiAPIPatch(o["storm-control"]) {
					return fmt.Errorf("Error reading storm_control: %v", err)
				}
			}
		}
	}

	if err = d.Set("security_mode", flattenSwitchPhysicalPortSecurityMode(o["security-mode"], d, "security_mode", sv)); err != nil {
		if !fortiAPIPatch(o["security-mode"]) {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("flapguard_state", flattenSwitchPhysicalPortFlapguardState(o["flapguard-state"], d, "flapguard_state", sv)); err != nil {
		if !fortiAPIPatch(o["flapguard-state"]) {
			return fmt.Errorf("Error reading flapguard_state: %v", err)
		}
	}

	if err = d.Set("pause_resume", flattenSwitchPhysicalPortPauseResume(o["pause-resume"], d, "pause_resume", sv)); err != nil {
		if !fortiAPIPatch(o["pause-resume"]) {
			return fmt.Errorf("Error reading pause_resume: %v", err)
		}
	}

	if err = d.Set("l2_sa_unknown", flattenSwitchPhysicalPortL2SaUnknown(o["l2-sa-unknown"], d, "l2_sa_unknown", sv)); err != nil {
		if !fortiAPIPatch(o["l2-sa-unknown"]) {
			return fmt.Errorf("Error reading l2_sa_unknown: %v", err)
		}
	}

	if err = d.Set("lldp_status", flattenSwitchPhysicalPortLldpStatus(o["lldp-status"], d, "lldp_status", sv)); err != nil {
		if !fortiAPIPatch(o["lldp-status"]) {
			return fmt.Errorf("Error reading lldp_status: %v", err)
		}
	}

	if err = d.Set("eee_tx_idle_time", flattenSwitchPhysicalPortEeeTxIdleTime(o["eee-tx-idle-time"], d, "eee_tx_idle_time", sv)); err != nil {
		if !fortiAPIPatch(o["eee-tx-idle-time"]) {
			return fmt.Errorf("Error reading eee_tx_idle_time: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchPhysicalPortName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("lldp_profile", flattenSwitchPhysicalPortLldpProfile(o["lldp-profile"], d, "lldp_profile", sv)); err != nil {
		if !fortiAPIPatch(o["lldp-profile"]) {
			return fmt.Errorf("Error reading lldp_profile: %v", err)
		}
	}

	if err = d.Set("flap_trig", flattenSwitchPhysicalPortFlapTrig(o["flap-trig"], d, "flap_trig", sv)); err != nil {
		if !fortiAPIPatch(o["flap-trig"]) {
			return fmt.Errorf("Error reading flap_trig: %v", err)
		}
	}

	if err = d.Set("link_status", flattenSwitchPhysicalPortLinkStatus(o["link-status"], d, "link_status", sv)); err != nil {
		if !fortiAPIPatch(o["link-status"]) {
			return fmt.Errorf("Error reading link_status: %v", err)
		}
	}

	if err = d.Set("flap_timeout", flattenSwitchPhysicalPortFlapTimeout(o["flap-timeout"], d, "flap_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["flap-timeout"]) {
			return fmt.Errorf("Error reading flap_timeout: %v", err)
		}
	}

	if err = d.Set("poe_port_mode", flattenSwitchPhysicalPortPoePortMode(o["poe-port-mode"], d, "poe_port_mode", sv)); err != nil {
		if !fortiAPIPatch(o["poe-port-mode"]) {
			return fmt.Errorf("Error reading poe_port_mode: %v", err)
		}
	}

	if err = d.Set("poe_max_power_mode", flattenSwitchPhysicalPortPoeMaxPowerMode(o["poe-max-power-mode"], d, "poe_max_power_mode", sv)); err != nil {
		if !fortiAPIPatch(o["poe-max-power-mode"]) {
			return fmt.Errorf("Error reading poe_max_power_mode: %v", err)
		}
	}

	if err = d.Set("port_index", flattenSwitchPhysicalPortPortIndex(o["port-index"], d, "port_index", sv)); err != nil {
		if !fortiAPIPatch(o["port-index"]) {
			return fmt.Errorf("Error reading port_index: %v", err)
		}
	}

	if err = d.Set("max_frame_size", flattenSwitchPhysicalPortMaxFrameSize(o["max-frame-size"], d, "max_frame_size", sv)); err != nil {
		if !fortiAPIPatch(o["max-frame-size"]) {
			return fmt.Errorf("Error reading max_frame_size: %v", err)
		}
	}

	if err = d.Set("flow_control", flattenSwitchPhysicalPortFlowControl(o["flow-control"], d, "flow_control", sv)); err != nil {
		if !fortiAPIPatch(o["flow-control"]) {
			return fmt.Errorf("Error reading flow_control: %v", err)
		}
	}

	if err = d.Set("pause_meter_rate", flattenSwitchPhysicalPortPauseMeterRate(o["pause-meter-rate"], d, "pause_meter_rate", sv)); err != nil {
		if !fortiAPIPatch(o["pause-meter-rate"]) {
			return fmt.Errorf("Error reading pause_meter_rate: %v", err)
		}
	}

	if err = d.Set("poe_port_priority", flattenSwitchPhysicalPortPoePortPriority(o["poe-port-priority"], d, "poe_port_priority", sv)); err != nil {
		if !fortiAPIPatch(o["poe-port-priority"]) {
			return fmt.Errorf("Error reading poe_port_priority: %v", err)
		}
	}

	return nil
}

func flattenSwitchPhysicalPortFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchPhysicalPortMacsecPaeMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortL2Learning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortDmiStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortStormControlMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortSpeed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortFortilinkP2P(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortFlapRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortMacsecProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortEgressDropMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortLoopback(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortPriorityBasedFlowControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortOwningInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortEnergyEfficientEthernet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortQsfpLowPowerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortPoeStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortMedium(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortCdpStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortFlapguard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortFlapDuration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortEeeTxWakeTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortStormControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "broadcast"
	if _, ok := d.GetOk(pre_append); ok {

		result["broadcast"], _ = expandSwitchPhysicalPortStormControlBroadcast(d, i["broadcast"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unknown_unicast"
	if _, ok := d.GetOk(pre_append); ok {

		result["unknown-unicast"], _ = expandSwitchPhysicalPortStormControlUnknownUnicast(d, i["unknown_unicast"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unknown_multicast"
	if _, ok := d.GetOk(pre_append); ok {

		result["unknown-multicast"], _ = expandSwitchPhysicalPortStormControlUnknownMulticast(d, i["unknown_multicast"], pre_append, sv)
	}
	pre_append = pre + ".0." + "rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["rate"], _ = expandSwitchPhysicalPortStormControlRate(d, i["rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "burst_size_level"
	if _, ok := d.GetOk(pre_append); ok {

		result["burst-size-level"], _ = expandSwitchPhysicalPortStormControlBurstSizeLevel(d, i["burst_size_level"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchPhysicalPortStormControlBroadcast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortStormControlUnknownUnicast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortStormControlUnknownMulticast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortStormControlRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortStormControlBurstSizeLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortSecurityMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortFlapguardState(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortPauseResume(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortL2SaUnknown(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortLldpStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortEeeTxIdleTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortLldpProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortFlapTrig(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortLinkStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortFlapTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortPoePortMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortPoeMaxPowerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortPortIndex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortMaxFrameSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortFlowControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortPauseMeterRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhysicalPortPoePortPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchPhysicalPort(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("macsec_pae_mode"); ok {

		t, err := expandSwitchPhysicalPortMacsecPaeMode(d, v, "macsec_pae_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macsec-pae-mode"] = t
		}
	}

	if v, ok := d.GetOk("l2_learning"); ok {

		t, err := expandSwitchPhysicalPortL2Learning(d, v, "l2_learning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2-learning"] = t
		}
	}

	if v, ok := d.GetOk("dmi_status"); ok {

		t, err := expandSwitchPhysicalPortDmiStatus(d, v, "dmi_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dmi-status"] = t
		}
	}

	if v, ok := d.GetOk("storm_control_mode"); ok {

		t, err := expandSwitchPhysicalPortStormControlMode(d, v, "storm_control_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["storm-control-mode"] = t
		}
	}

	if v, ok := d.GetOk("speed"); ok {

		t, err := expandSwitchPhysicalPortSpeed(d, v, "speed", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["speed"] = t
		}
	}

	if v, ok := d.GetOk("fortilink_p2p"); ok {

		t, err := expandSwitchPhysicalPortFortilinkP2P(d, v, "fortilink_p2p", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink-p2p"] = t
		}
	}

	if v, ok := d.GetOk("flap_rate"); ok {

		t, err := expandSwitchPhysicalPortFlapRate(d, v, "flap_rate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flap-rate"] = t
		}
	}

	if v, ok := d.GetOk("macsec_profile"); ok {

		t, err := expandSwitchPhysicalPortMacsecProfile(d, v, "macsec_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macsec-profile"] = t
		}
	}

	if v, ok := d.GetOk("egress_drop_mode"); ok {

		t, err := expandSwitchPhysicalPortEgressDropMode(d, v, "egress_drop_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["egress-drop-mode"] = t
		}
	}

	if v, ok := d.GetOk("loopback"); ok {

		t, err := expandSwitchPhysicalPortLoopback(d, v, "loopback", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loopback"] = t
		}
	}

	if v, ok := d.GetOk("priority_based_flow_control"); ok {

		t, err := expandSwitchPhysicalPortPriorityBasedFlowControl(d, v, "priority_based_flow_control", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-based-flow-control"] = t
		}
	}

	if v, ok := d.GetOk("owning_interface"); ok {

		t, err := expandSwitchPhysicalPortOwningInterface(d, v, "owning_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owning-interface"] = t
		}
	}

	if v, ok := d.GetOk("energy_efficient_ethernet"); ok {

		t, err := expandSwitchPhysicalPortEnergyEfficientEthernet(d, v, "energy_efficient_ethernet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["energy-efficient-ethernet"] = t
		}
	}

	if v, ok := d.GetOk("qsfp_low_power_mode"); ok {

		t, err := expandSwitchPhysicalPortQsfpLowPowerMode(d, v, "qsfp_low_power_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qsfp-low-power-mode"] = t
		}
	}

	if v, ok := d.GetOk("poe_status"); ok {

		t, err := expandSwitchPhysicalPortPoeStatus(d, v, "poe_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-status"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandSwitchPhysicalPortStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("medium"); ok {

		t, err := expandSwitchPhysicalPortMedium(d, v, "medium", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["medium"] = t
		}
	}

	if v, ok := d.GetOk("cdp_status"); ok {

		t, err := expandSwitchPhysicalPortCdpStatus(d, v, "cdp_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cdp-status"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchPhysicalPortDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("flapguard"); ok {

		t, err := expandSwitchPhysicalPortFlapguard(d, v, "flapguard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flapguard"] = t
		}
	}

	if v, ok := d.GetOk("flap_duration"); ok {

		t, err := expandSwitchPhysicalPortFlapDuration(d, v, "flap_duration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flap-duration"] = t
		}
	}

	if v, ok := d.GetOk("eee_tx_wake_time"); ok {

		t, err := expandSwitchPhysicalPortEeeTxWakeTime(d, v, "eee_tx_wake_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eee-tx-wake-time"] = t
		}
	}

	if v, ok := d.GetOk("storm_control"); ok {

		t, err := expandSwitchPhysicalPortStormControl(d, v, "storm_control", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["storm-control"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok {

		t, err := expandSwitchPhysicalPortSecurityMode(d, v, "security_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("flapguard_state"); ok {

		t, err := expandSwitchPhysicalPortFlapguardState(d, v, "flapguard_state", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flapguard-state"] = t
		}
	}

	if v, ok := d.GetOk("pause_resume"); ok {

		t, err := expandSwitchPhysicalPortPauseResume(d, v, "pause_resume", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pause-resume"] = t
		}
	}

	if v, ok := d.GetOk("l2_sa_unknown"); ok {

		t, err := expandSwitchPhysicalPortL2SaUnknown(d, v, "l2_sa_unknown", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2-sa-unknown"] = t
		}
	}

	if v, ok := d.GetOk("lldp_status"); ok {

		t, err := expandSwitchPhysicalPortLldpStatus(d, v, "lldp_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-status"] = t
		}
	}

	if v, ok := d.GetOk("eee_tx_idle_time"); ok {

		t, err := expandSwitchPhysicalPortEeeTxIdleTime(d, v, "eee_tx_idle_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eee-tx-idle-time"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchPhysicalPortName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("lldp_profile"); ok {

		t, err := expandSwitchPhysicalPortLldpProfile(d, v, "lldp_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-profile"] = t
		}
	}

	if v, ok := d.GetOk("flap_trig"); ok {

		t, err := expandSwitchPhysicalPortFlapTrig(d, v, "flap_trig", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flap-trig"] = t
		}
	}

	if v, ok := d.GetOk("link_status"); ok {

		t, err := expandSwitchPhysicalPortLinkStatus(d, v, "link_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-status"] = t
		}
	}

	if v, ok := d.GetOk("flap_timeout"); ok {

		t, err := expandSwitchPhysicalPortFlapTimeout(d, v, "flap_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flap-timeout"] = t
		}
	}

	if v, ok := d.GetOk("poe_port_mode"); ok {

		t, err := expandSwitchPhysicalPortPoePortMode(d, v, "poe_port_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-port-mode"] = t
		}
	}

	if v, ok := d.GetOk("poe_max_power_mode"); ok {

		t, err := expandSwitchPhysicalPortPoeMaxPowerMode(d, v, "poe_max_power_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-max-power-mode"] = t
		}
	}

	if v, ok := d.GetOk("port_index"); ok {

		t, err := expandSwitchPhysicalPortPortIndex(d, v, "port_index", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-index"] = t
		}
	}

	if v, ok := d.GetOk("max_frame_size"); ok {

		t, err := expandSwitchPhysicalPortMaxFrameSize(d, v, "max_frame_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-frame-size"] = t
		}
	}

	if v, ok := d.GetOk("flow_control"); ok {

		t, err := expandSwitchPhysicalPortFlowControl(d, v, "flow_control", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flow-control"] = t
		}
	}

	if v, ok := d.GetOk("pause_meter_rate"); ok {

		t, err := expandSwitchPhysicalPortPauseMeterRate(d, v, "pause_meter_rate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pause-meter-rate"] = t
		}
	}

	if v, ok := d.GetOk("poe_port_priority"); ok {

		t, err := expandSwitchPhysicalPortPoePortPriority(d, v, "poe_port_priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-port-priority"] = t
		}
	}

	return &obj, nil
}
