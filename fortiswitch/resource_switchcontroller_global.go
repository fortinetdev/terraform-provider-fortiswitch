// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Switch-controller global configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerGlobalUpdate,
		Read:   resourceSwitchControllerGlobalRead,
		Update: resourceSwitchControllerGlobalUpdate,
		Delete: resourceSwitchControllerGlobalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ac_data_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1024, 49150),
				Optional:     true,
				Computed:     true,
			},
			"echo_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 600),
				Optional:     true,
				Computed:     true,
			},
			"ac_dhcp_option_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_discoveries": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"ac_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"location": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ac_discovery_mc_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ac_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1024, 49150),
				Optional:     true,
				Computed:     true,
			},
			"max_retransmit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"ac_discovery_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mgmt_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_mode": &schema.Schema{
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

func resourceSwitchControllerGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerGlobal(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerGlobal resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerGlobal(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerGlobal")
	}

	return resourceSwitchControllerGlobalRead(d, m)
}

func resourceSwitchControllerGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerGlobal(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerGlobal(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchControllerGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerGlobalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerGlobal(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerGlobal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerGlobal resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerGlobalName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalAcDataPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalEchoInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalAcDhcpOptionCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalMaxDiscoveries(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalAcList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_address"
		if _, ok := i["ipv6-address"]; ok {

			tmp["ipv6_address"] = flattenSwitchControllerGlobalAcListIpv6Address(i["ipv6-address"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_address"
		if _, ok := i["ipv4-address"]; ok {

			tmp["ipv4_address"] = flattenSwitchControllerGlobalAcListIpv4Address(i["ipv4-address"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSwitchControllerGlobalAcListId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchControllerGlobalAcListIpv6Address(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalAcListIpv4Address(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalAcListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalAcDiscoveryMcAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalAcPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalMaxRetransmit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalAcDiscoveryType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalMgmtMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalSourceIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalTunnelMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerGlobal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerGlobalName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ac_data_port", flattenSwitchControllerGlobalAcDataPort(o["ac-data-port"], d, "ac_data_port", sv)); err != nil {
		if !fortiAPIPatch(o["ac-data-port"]) {
			return fmt.Errorf("Error reading ac_data_port: %v", err)
		}
	}

	if err = d.Set("echo_interval", flattenSwitchControllerGlobalEchoInterval(o["echo-interval"], d, "echo_interval", sv)); err != nil {
		if !fortiAPIPatch(o["echo-interval"]) {
			return fmt.Errorf("Error reading echo_interval: %v", err)
		}
	}

	if err = d.Set("ac_dhcp_option_code", flattenSwitchControllerGlobalAcDhcpOptionCode(o["ac-dhcp-option-code"], d, "ac_dhcp_option_code", sv)); err != nil {
		if !fortiAPIPatch(o["ac-dhcp-option-code"]) {
			return fmt.Errorf("Error reading ac_dhcp_option_code: %v", err)
		}
	}

	if err = d.Set("max_discoveries", flattenSwitchControllerGlobalMaxDiscoveries(o["max-discoveries"], d, "max_discoveries", sv)); err != nil {
		if !fortiAPIPatch(o["max-discoveries"]) {
			return fmt.Errorf("Error reading max_discoveries: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ac_list", flattenSwitchControllerGlobalAcList(o["ac-list"], d, "ac_list", sv)); err != nil {
			if !fortiAPIPatch(o["ac-list"]) {
				return fmt.Errorf("Error reading ac_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ac_list"); ok {
			if err = d.Set("ac_list", flattenSwitchControllerGlobalAcList(o["ac-list"], d, "ac_list", sv)); err != nil {
				if !fortiAPIPatch(o["ac-list"]) {
					return fmt.Errorf("Error reading ac_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("location", flattenSwitchControllerGlobalLocation(o["location"], d, "location", sv)); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("ac_discovery_mc_addr", flattenSwitchControllerGlobalAcDiscoveryMcAddr(o["ac-discovery-mc-addr"], d, "ac_discovery_mc_addr", sv)); err != nil {
		if !fortiAPIPatch(o["ac-discovery-mc-addr"]) {
			return fmt.Errorf("Error reading ac_discovery_mc_addr: %v", err)
		}
	}

	if err = d.Set("ac_port", flattenSwitchControllerGlobalAcPort(o["ac-port"], d, "ac_port", sv)); err != nil {
		if !fortiAPIPatch(o["ac-port"]) {
			return fmt.Errorf("Error reading ac_port: %v", err)
		}
	}

	if err = d.Set("max_retransmit", flattenSwitchControllerGlobalMaxRetransmit(o["max-retransmit"], d, "max_retransmit", sv)); err != nil {
		if !fortiAPIPatch(o["max-retransmit"]) {
			return fmt.Errorf("Error reading max_retransmit: %v", err)
		}
	}

	if err = d.Set("ac_discovery_type", flattenSwitchControllerGlobalAcDiscoveryType(o["ac-discovery-type"], d, "ac_discovery_type", sv)); err != nil {
		if !fortiAPIPatch(o["ac-discovery-type"]) {
			return fmt.Errorf("Error reading ac_discovery_type: %v", err)
		}
	}

	if err = d.Set("mgmt_mode", flattenSwitchControllerGlobalMgmtMode(o["mgmt-mode"], d, "mgmt_mode", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-mode"]) {
			return fmt.Errorf("Error reading mgmt_mode: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSwitchControllerGlobalSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenSwitchControllerGlobalSourceIp6(o["source-ip6"], d, "source_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("tunnel_mode", flattenSwitchControllerGlobalTunnelMode(o["tunnel-mode"], d, "tunnel_mode", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-mode"]) {
			return fmt.Errorf("Error reading tunnel_mode: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerGlobalFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchControllerGlobalName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalAcDataPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalEchoInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalAcDhcpOptionCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalMaxDiscoveries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalAcList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_address"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ipv6-address"], _ = expandSwitchControllerGlobalAcListIpv6Address(d, i["ipv6_address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_address"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ipv4-address"], _ = expandSwitchControllerGlobalAcListIpv4Address(d, i["ipv4_address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSwitchControllerGlobalAcListId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerGlobalAcListIpv6Address(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalAcListIpv4Address(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalAcListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalAcDiscoveryMcAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalAcPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalMaxRetransmit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalAcDiscoveryType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalMgmtMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalSourceIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalTunnelMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerGlobal(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		if setArgNil {
			obj["name"] = nil
		} else {

			t, err := expandSwitchControllerGlobalName(d, v, "name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["name"] = t
			}
		}
	}

	if v, ok := d.GetOk("ac_data_port"); ok {
		if setArgNil {
			obj["ac-data-port"] = nil
		} else {

			t, err := expandSwitchControllerGlobalAcDataPort(d, v, "ac_data_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ac-data-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("echo_interval"); ok {
		if setArgNil {
			obj["echo-interval"] = nil
		} else {

			t, err := expandSwitchControllerGlobalEchoInterval(d, v, "echo_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["echo-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("ac_dhcp_option_code"); ok {
		if setArgNil {
			obj["ac-dhcp-option-code"] = nil
		} else {

			t, err := expandSwitchControllerGlobalAcDhcpOptionCode(d, v, "ac_dhcp_option_code", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ac-dhcp-option-code"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("max_discoveries"); ok {
		if setArgNil {
			obj["max-discoveries"] = nil
		} else {

			t, err := expandSwitchControllerGlobalMaxDiscoveries(d, v, "max_discoveries", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-discoveries"] = t
			}
		}
	}

	if v, ok := d.GetOk("ac_list"); ok || d.HasChange("ac_list") {
		if setArgNil {
			obj["ac-list"] = make([]struct{}, 0)
		} else {

			t, err := expandSwitchControllerGlobalAcList(d, v, "ac_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ac-list"] = t
			}
		}
	}

	if v, ok := d.GetOk("location"); ok {
		if setArgNil {
			obj["location"] = nil
		} else {

			t, err := expandSwitchControllerGlobalLocation(d, v, "location", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["location"] = t
			}
		}
	}

	if v, ok := d.GetOk("ac_discovery_mc_addr"); ok {
		if setArgNil {
			obj["ac-discovery-mc-addr"] = nil
		} else {

			t, err := expandSwitchControllerGlobalAcDiscoveryMcAddr(d, v, "ac_discovery_mc_addr", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ac-discovery-mc-addr"] = t
			}
		}
	}

	if v, ok := d.GetOk("ac_port"); ok {
		if setArgNil {
			obj["ac-port"] = nil
		} else {

			t, err := expandSwitchControllerGlobalAcPort(d, v, "ac_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ac-port"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("max_retransmit"); ok {
		if setArgNil {
			obj["max-retransmit"] = nil
		} else {

			t, err := expandSwitchControllerGlobalMaxRetransmit(d, v, "max_retransmit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-retransmit"] = t
			}
		}
	}

	if v, ok := d.GetOk("ac_discovery_type"); ok {
		if setArgNil {
			obj["ac-discovery-type"] = nil
		} else {

			t, err := expandSwitchControllerGlobalAcDiscoveryType(d, v, "ac_discovery_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ac-discovery-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("mgmt_mode"); ok {
		if setArgNil {
			obj["mgmt-mode"] = nil
		} else {

			t, err := expandSwitchControllerGlobalMgmtMode(d, v, "mgmt_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mgmt-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {

			t, err := expandSwitchControllerGlobalSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {
		if setArgNil {
			obj["source-ip6"] = nil
		} else {

			t, err := expandSwitchControllerGlobalSourceIp6(d, v, "source_ip6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip6"] = t
			}
		}
	}

	if v, ok := d.GetOk("tunnel_mode"); ok {
		if setArgNil {
			obj["tunnel-mode"] = nil
		} else {

			t, err := expandSwitchControllerGlobalTunnelMode(d, v, "tunnel_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tunnel-mode"] = t
			}
		}
	}

	return &obj, nil
}
