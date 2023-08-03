// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Link-aggregation.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchTrunk() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchTrunkCreate,
		Read:   resourceSwitchTrunkRead,
		Update: resourceSwitchTrunkUpdate,
		Delete: resourceSwitchTrunkDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"hb_src_udp_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hb_verify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hb_src_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_bundle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hb_dst_udp_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"member_withdrawal_behavior": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mclag_mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"isl_fortilink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_miss_heartbeats": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"aggregator_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_selection_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lacp_speed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mclag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trunk_id": &schema.Schema{
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
			"static_isl_auto_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bundle": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_bundle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mclag_icl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hb_in_vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"hb_dst_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hb_out_vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fortilink": &schema.Schema{
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
			"auto_isl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port_extension": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"static_isl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_extension_trigger": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSwitchTrunkCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchTrunk(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchTrunk resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchTrunk(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchTrunk resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchTrunk")
	}

	return resourceSwitchTrunkRead(d, m)
}

func resourceSwitchTrunkUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchTrunk(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchTrunk resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchTrunk(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchTrunk resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchTrunk")
	}

	return resourceSwitchTrunkRead(d, m)
}

func resourceSwitchTrunkDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchTrunk(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchTrunk resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchTrunkRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchTrunk(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchTrunk resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchTrunk(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchTrunk resource from API: %v", err)
	}
	return nil
}

func flattenSwitchTrunkHbSrcUdpPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkHbVerify(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkHbSrcIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkMinBundle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkHbDstUdpPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkMemberWithdrawalBehavior(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkMclagMacAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkIslFortilink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkMaxMissHeartbeats(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkAggregatorMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkPortSelectionCriteria(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkLacpSpeed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkMclag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkTrunkId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkStaticIslAutoVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkBundle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkMaxBundle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkMclagIcl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkHbInVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkMembers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["member_name"] = flattenSwitchTrunkMembersMemberName(i["member-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "member_name", d)
	return result
}

func flattenSwitchTrunkMembersMemberName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkHbDstIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkHbOutVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkFortilink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkAutoIsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkPortExtension(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkStaticIsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchTrunkPortExtensionTrigger(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchTrunk(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("hb_src_udp_port", flattenSwitchTrunkHbSrcUdpPort(o["hb-src-udp-port"], d, "hb_src_udp_port", sv)); err != nil {
		if !fortiAPIPatch(o["hb-src-udp-port"]) {
			return fmt.Errorf("Error reading hb_src_udp_port: %v", err)
		}
	}

	if err = d.Set("hb_verify", flattenSwitchTrunkHbVerify(o["hb-verify"], d, "hb_verify", sv)); err != nil {
		if !fortiAPIPatch(o["hb-verify"]) {
			return fmt.Errorf("Error reading hb_verify: %v", err)
		}
	}

	if err = d.Set("hb_src_ip", flattenSwitchTrunkHbSrcIp(o["hb-src-ip"], d, "hb_src_ip", sv)); err != nil {
		if !fortiAPIPatch(o["hb-src-ip"]) {
			return fmt.Errorf("Error reading hb_src_ip: %v", err)
		}
	}

	if err = d.Set("min_bundle", flattenSwitchTrunkMinBundle(o["min-bundle"], d, "min_bundle", sv)); err != nil {
		if !fortiAPIPatch(o["min-bundle"]) {
			return fmt.Errorf("Error reading min_bundle: %v", err)
		}
	}

	if err = d.Set("hb_dst_udp_port", flattenSwitchTrunkHbDstUdpPort(o["hb-dst-udp-port"], d, "hb_dst_udp_port", sv)); err != nil {
		if !fortiAPIPatch(o["hb-dst-udp-port"]) {
			return fmt.Errorf("Error reading hb_dst_udp_port: %v", err)
		}
	}

	if err = d.Set("member_withdrawal_behavior", flattenSwitchTrunkMemberWithdrawalBehavior(o["member-withdrawal-behavior"], d, "member_withdrawal_behavior", sv)); err != nil {
		if !fortiAPIPatch(o["member-withdrawal-behavior"]) {
			return fmt.Errorf("Error reading member_withdrawal_behavior: %v", err)
		}
	}

	if err = d.Set("mclag_mac_address", flattenSwitchTrunkMclagMacAddress(o["mclag-mac-address"], d, "mclag_mac_address", sv)); err != nil {
		if !fortiAPIPatch(o["mclag-mac-address"]) {
			return fmt.Errorf("Error reading mclag_mac_address: %v", err)
		}
	}

	if err = d.Set("isl_fortilink", flattenSwitchTrunkIslFortilink(o["isl-fortilink"], d, "isl_fortilink", sv)); err != nil {
		if !fortiAPIPatch(o["isl-fortilink"]) {
			return fmt.Errorf("Error reading isl_fortilink: %v", err)
		}
	}

	if err = d.Set("max_miss_heartbeats", flattenSwitchTrunkMaxMissHeartbeats(o["max-miss-heartbeats"], d, "max_miss_heartbeats", sv)); err != nil {
		if !fortiAPIPatch(o["max-miss-heartbeats"]) {
			return fmt.Errorf("Error reading max_miss_heartbeats: %v", err)
		}
	}

	if err = d.Set("aggregator_mode", flattenSwitchTrunkAggregatorMode(o["aggregator-mode"], d, "aggregator_mode", sv)); err != nil {
		if !fortiAPIPatch(o["aggregator-mode"]) {
			return fmt.Errorf("Error reading aggregator_mode: %v", err)
		}
	}

	if err = d.Set("port_selection_criteria", flattenSwitchTrunkPortSelectionCriteria(o["port-selection-criteria"], d, "port_selection_criteria", sv)); err != nil {
		if !fortiAPIPatch(o["port-selection-criteria"]) {
			return fmt.Errorf("Error reading port_selection_criteria: %v", err)
		}
	}

	if err = d.Set("lacp_speed", flattenSwitchTrunkLacpSpeed(o["lacp-speed"], d, "lacp_speed", sv)); err != nil {
		if !fortiAPIPatch(o["lacp-speed"]) {
			return fmt.Errorf("Error reading lacp_speed: %v", err)
		}
	}

	if err = d.Set("mclag", flattenSwitchTrunkMclag(o["mclag"], d, "mclag", sv)); err != nil {
		if !fortiAPIPatch(o["mclag"]) {
			return fmt.Errorf("Error reading mclag: %v", err)
		}
	}

	if err = d.Set("trunk_id", flattenSwitchTrunkTrunkId(o["trunk-id"], d, "trunk_id", sv)); err != nil {
		if !fortiAPIPatch(o["trunk-id"]) {
			return fmt.Errorf("Error reading trunk_id: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchTrunkDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("static_isl_auto_vlan", flattenSwitchTrunkStaticIslAutoVlan(o["static-isl-auto-vlan"], d, "static_isl_auto_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["static-isl-auto-vlan"]) {
			return fmt.Errorf("Error reading static_isl_auto_vlan: %v", err)
		}
	}

	if err = d.Set("bundle", flattenSwitchTrunkBundle(o["bundle"], d, "bundle", sv)); err != nil {
		if !fortiAPIPatch(o["bundle"]) {
			return fmt.Errorf("Error reading bundle: %v", err)
		}
	}

	if err = d.Set("max_bundle", flattenSwitchTrunkMaxBundle(o["max-bundle"], d, "max_bundle", sv)); err != nil {
		if !fortiAPIPatch(o["max-bundle"]) {
			return fmt.Errorf("Error reading max_bundle: %v", err)
		}
	}

	if err = d.Set("mclag_icl", flattenSwitchTrunkMclagIcl(o["mclag-icl"], d, "mclag_icl", sv)); err != nil {
		if !fortiAPIPatch(o["mclag-icl"]) {
			return fmt.Errorf("Error reading mclag_icl: %v", err)
		}
	}

	if err = d.Set("hb_in_vlan", flattenSwitchTrunkHbInVlan(o["hb-in-vlan"], d, "hb_in_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["hb-in-vlan"]) {
			return fmt.Errorf("Error reading hb_in_vlan: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("members", flattenSwitchTrunkMembers(o["members"], d, "members", sv)); err != nil {
			if !fortiAPIPatch(o["members"]) {
				return fmt.Errorf("Error reading members: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("members"); ok {
			if err = d.Set("members", flattenSwitchTrunkMembers(o["members"], d, "members", sv)); err != nil {
				if !fortiAPIPatch(o["members"]) {
					return fmt.Errorf("Error reading members: %v", err)
				}
			}
		}
	}

	if err = d.Set("hb_dst_ip", flattenSwitchTrunkHbDstIp(o["hb-dst-ip"], d, "hb_dst_ip", sv)); err != nil {
		if !fortiAPIPatch(o["hb-dst-ip"]) {
			return fmt.Errorf("Error reading hb_dst_ip: %v", err)
		}
	}

	if err = d.Set("hb_out_vlan", flattenSwitchTrunkHbOutVlan(o["hb-out-vlan"], d, "hb_out_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["hb-out-vlan"]) {
			return fmt.Errorf("Error reading hb_out_vlan: %v", err)
		}
	}

	if err = d.Set("fortilink", flattenSwitchTrunkFortilink(o["fortilink"], d, "fortilink", sv)); err != nil {
		if !fortiAPIPatch(o["fortilink"]) {
			return fmt.Errorf("Error reading fortilink: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchTrunkName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("auto_isl", flattenSwitchTrunkAutoIsl(o["auto-isl"], d, "auto_isl", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl"]) {
			return fmt.Errorf("Error reading auto_isl: %v", err)
		}
	}

	if err = d.Set("port_extension", flattenSwitchTrunkPortExtension(o["port-extension"], d, "port_extension", sv)); err != nil {
		if !fortiAPIPatch(o["port-extension"]) {
			return fmt.Errorf("Error reading port_extension: %v", err)
		}
	}

	if err = d.Set("mode", flattenSwitchTrunkMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("static_isl", flattenSwitchTrunkStaticIsl(o["static-isl"], d, "static_isl", sv)); err != nil {
		if !fortiAPIPatch(o["static-isl"]) {
			return fmt.Errorf("Error reading static_isl: %v", err)
		}
	}

	if err = d.Set("port_extension_trigger", flattenSwitchTrunkPortExtensionTrigger(o["port-extension-trigger"], d, "port_extension_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["port-extension-trigger"]) {
			return fmt.Errorf("Error reading port_extension_trigger: %v", err)
		}
	}

	return nil
}

func flattenSwitchTrunkFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchTrunkHbSrcUdpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkHbVerify(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkHbSrcIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkMinBundle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkHbDstUdpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkMemberWithdrawalBehavior(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkMclagMacAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkIslFortilink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkMaxMissHeartbeats(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkAggregatorMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkPortSelectionCriteria(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkLacpSpeed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkMclag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkTrunkId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkStaticIslAutoVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkBundle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkMaxBundle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkMclagIcl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkHbInVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["member-name"], _ = expandSwitchTrunkMembersMemberName(d, i["member_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchTrunkMembersMemberName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkHbDstIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkHbOutVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkFortilink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkAutoIsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkPortExtension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkStaticIsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchTrunkPortExtensionTrigger(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchTrunk(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("hb_src_udp_port"); ok {

		t, err := expandSwitchTrunkHbSrcUdpPort(d, v, "hb_src_udp_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-src-udp-port"] = t
		}
	}

	if v, ok := d.GetOk("hb_verify"); ok {

		t, err := expandSwitchTrunkHbVerify(d, v, "hb_verify", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-verify"] = t
		}
	}

	if v, ok := d.GetOk("hb_src_ip"); ok {

		t, err := expandSwitchTrunkHbSrcIp(d, v, "hb_src_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-src-ip"] = t
		}
	}

	if v, ok := d.GetOk("min_bundle"); ok {

		t, err := expandSwitchTrunkMinBundle(d, v, "min_bundle", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-bundle"] = t
		}
	}

	if v, ok := d.GetOk("hb_dst_udp_port"); ok {

		t, err := expandSwitchTrunkHbDstUdpPort(d, v, "hb_dst_udp_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-dst-udp-port"] = t
		}
	}

	if v, ok := d.GetOk("member_withdrawal_behavior"); ok {

		t, err := expandSwitchTrunkMemberWithdrawalBehavior(d, v, "member_withdrawal_behavior", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-withdrawal-behavior"] = t
		}
	}

	if v, ok := d.GetOk("mclag_mac_address"); ok {

		t, err := expandSwitchTrunkMclagMacAddress(d, v, "mclag_mac_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mclag-mac-address"] = t
		}
	}

	if v, ok := d.GetOk("isl_fortilink"); ok {

		t, err := expandSwitchTrunkIslFortilink(d, v, "isl_fortilink", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isl-fortilink"] = t
		}
	}

	if v, ok := d.GetOk("max_miss_heartbeats"); ok {

		t, err := expandSwitchTrunkMaxMissHeartbeats(d, v, "max_miss_heartbeats", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-miss-heartbeats"] = t
		}
	}

	if v, ok := d.GetOk("aggregator_mode"); ok {

		t, err := expandSwitchTrunkAggregatorMode(d, v, "aggregator_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregator-mode"] = t
		}
	}

	if v, ok := d.GetOk("port_selection_criteria"); ok {

		t, err := expandSwitchTrunkPortSelectionCriteria(d, v, "port_selection_criteria", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-selection-criteria"] = t
		}
	}

	if v, ok := d.GetOk("lacp_speed"); ok {

		t, err := expandSwitchTrunkLacpSpeed(d, v, "lacp_speed", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lacp-speed"] = t
		}
	}

	if v, ok := d.GetOk("mclag"); ok {

		t, err := expandSwitchTrunkMclag(d, v, "mclag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mclag"] = t
		}
	}

	if v, ok := d.GetOk("trunk_id"); ok {

		t, err := expandSwitchTrunkTrunkId(d, v, "trunk_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trunk-id"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchTrunkDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("static_isl_auto_vlan"); ok {

		t, err := expandSwitchTrunkStaticIslAutoVlan(d, v, "static_isl_auto_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["static-isl-auto-vlan"] = t
		}
	}

	if v, ok := d.GetOk("bundle"); ok {

		t, err := expandSwitchTrunkBundle(d, v, "bundle", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bundle"] = t
		}
	}

	if v, ok := d.GetOk("max_bundle"); ok {

		t, err := expandSwitchTrunkMaxBundle(d, v, "max_bundle", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-bundle"] = t
		}
	}

	if v, ok := d.GetOk("mclag_icl"); ok {

		t, err := expandSwitchTrunkMclagIcl(d, v, "mclag_icl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mclag-icl"] = t
		}
	}

	if v, ok := d.GetOk("hb_in_vlan"); ok {

		t, err := expandSwitchTrunkHbInVlan(d, v, "hb_in_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-in-vlan"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok || d.HasChange("members") {

		t, err := expandSwitchTrunkMembers(d, v, "members", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	}

	if v, ok := d.GetOk("hb_dst_ip"); ok {

		t, err := expandSwitchTrunkHbDstIp(d, v, "hb_dst_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-dst-ip"] = t
		}
	}

	if v, ok := d.GetOk("hb_out_vlan"); ok {

		t, err := expandSwitchTrunkHbOutVlan(d, v, "hb_out_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-out-vlan"] = t
		}
	}

	if v, ok := d.GetOk("fortilink"); ok {

		t, err := expandSwitchTrunkFortilink(d, v, "fortilink", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchTrunkName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl"); ok {

		t, err := expandSwitchTrunkAutoIsl(d, v, "auto_isl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl"] = t
		}
	}

	if v, ok := d.GetOk("port_extension"); ok {

		t, err := expandSwitchTrunkPortExtension(d, v, "port_extension", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-extension"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {

		t, err := expandSwitchTrunkMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("static_isl"); ok {

		t, err := expandSwitchTrunkStaticIsl(d, v, "static_isl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["static-isl"] = t
		}
	}

	if v, ok := d.GetOk("port_extension_trigger"); ok {

		t, err := expandSwitchTrunkPortExtensionTrigger(d, v, "port_extension_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-extension-trigger"] = t
		}
	}

	return &obj, nil
}
