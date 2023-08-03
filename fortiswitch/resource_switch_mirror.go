// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Packet mirror.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchMirror() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchMirrorCreate,
		Read:   resourceSwitchMirrorRead,
		Update: resourceSwitchMirrorUpdate,
		Delete: resourceSwitchMirrorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_ipv4_tos": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"erspan_collector_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switching_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_vlan_priority": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"dst": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"encap_gre_protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"src_egress": &schema.Schema{
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
			"encap_ipv4_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"encap_vlan_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4094),
				Optional:     true,
				Computed:     true,
			},
			"encap_vlan_tpid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_mac_dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_ingress": &schema.Schema{
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
			"rspan_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_mac_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strip_mirrored_traffic_tags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_ipv4_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_vlan_cfi": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1),
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchMirrorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchMirror(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchMirror resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchMirror(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchMirror resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchMirror")
	}

	return resourceSwitchMirrorRead(d, m)
}

func resourceSwitchMirrorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchMirror(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchMirror resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchMirror(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchMirror resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchMirror")
	}

	return resourceSwitchMirrorRead(d, m)
}

func resourceSwitchMirrorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchMirror(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchMirror resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchMirrorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchMirror(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchMirror resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchMirror(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchMirror resource from API: %v", err)
	}
	return nil
}

func flattenSwitchMirrorStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorEncapVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorEncapIpv4Tos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorErspanCollectorIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorSwitchingPacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorEncapVlanPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorDst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorEncapGreProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorSrcEgress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSwitchMirrorSrcEgressName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchMirrorSrcEgressName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorEncapIpv4Ttl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorEncapVlanId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorEncapVlanTpid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorEncapMacDst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorSrcIngress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSwitchMirrorSrcIngressName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchMirrorSrcIngressName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorRspanIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorEncapMacSrc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorStripMirroredTrafficTags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorEncapIpv4Src(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorEncapVlanCfi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMirrorName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchMirror(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSwitchMirrorStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("encap_vlan", flattenSwitchMirrorEncapVlan(o["encap-vlan"], d, "encap_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["encap-vlan"]) {
			return fmt.Errorf("Error reading encap_vlan: %v", err)
		}
	}

	if err = d.Set("encap_ipv4_tos", flattenSwitchMirrorEncapIpv4Tos(o["encap-ipv4-tos"], d, "encap_ipv4_tos", sv)); err != nil {
		if !fortiAPIPatch(o["encap-ipv4-tos"]) {
			return fmt.Errorf("Error reading encap_ipv4_tos: %v", err)
		}
	}

	if err = d.Set("erspan_collector_ip", flattenSwitchMirrorErspanCollectorIp(o["erspan-collector-ip"], d, "erspan_collector_ip", sv)); err != nil {
		if !fortiAPIPatch(o["erspan-collector-ip"]) {
			return fmt.Errorf("Error reading erspan_collector_ip: %v", err)
		}
	}

	if err = d.Set("switching_packet", flattenSwitchMirrorSwitchingPacket(o["switching-packet"], d, "switching_packet", sv)); err != nil {
		if !fortiAPIPatch(o["switching-packet"]) {
			return fmt.Errorf("Error reading switching_packet: %v", err)
		}
	}

	if err = d.Set("encap_vlan_priority", flattenSwitchMirrorEncapVlanPriority(o["encap-vlan-priority"], d, "encap_vlan_priority", sv)); err != nil {
		if !fortiAPIPatch(o["encap-vlan-priority"]) {
			return fmt.Errorf("Error reading encap_vlan_priority: %v", err)
		}
	}

	if err = d.Set("dst", flattenSwitchMirrorDst(o["dst"], d, "dst", sv)); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("encap_gre_protocol", flattenSwitchMirrorEncapGreProtocol(o["encap-gre-protocol"], d, "encap_gre_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["encap-gre-protocol"]) {
			return fmt.Errorf("Error reading encap_gre_protocol: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("src_egress", flattenSwitchMirrorSrcEgress(o["src-egress"], d, "src_egress", sv)); err != nil {
			if !fortiAPIPatch(o["src-egress"]) {
				return fmt.Errorf("Error reading src_egress: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("src_egress"); ok {
			if err = d.Set("src_egress", flattenSwitchMirrorSrcEgress(o["src-egress"], d, "src_egress", sv)); err != nil {
				if !fortiAPIPatch(o["src-egress"]) {
					return fmt.Errorf("Error reading src_egress: %v", err)
				}
			}
		}
	}

	if err = d.Set("encap_ipv4_ttl", flattenSwitchMirrorEncapIpv4Ttl(o["encap-ipv4-ttl"], d, "encap_ipv4_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["encap-ipv4-ttl"]) {
			return fmt.Errorf("Error reading encap_ipv4_ttl: %v", err)
		}
	}

	if err = d.Set("encap_vlan_id", flattenSwitchMirrorEncapVlanId(o["encap-vlan-id"], d, "encap_vlan_id", sv)); err != nil {
		if !fortiAPIPatch(o["encap-vlan-id"]) {
			return fmt.Errorf("Error reading encap_vlan_id: %v", err)
		}
	}

	if err = d.Set("encap_vlan_tpid", flattenSwitchMirrorEncapVlanTpid(o["encap-vlan-tpid"], d, "encap_vlan_tpid", sv)); err != nil {
		if !fortiAPIPatch(o["encap-vlan-tpid"]) {
			return fmt.Errorf("Error reading encap_vlan_tpid: %v", err)
		}
	}

	if err = d.Set("mode", flattenSwitchMirrorMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("encap_mac_dst", flattenSwitchMirrorEncapMacDst(o["encap-mac-dst"], d, "encap_mac_dst", sv)); err != nil {
		if !fortiAPIPatch(o["encap-mac-dst"]) {
			return fmt.Errorf("Error reading encap_mac_dst: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("src_ingress", flattenSwitchMirrorSrcIngress(o["src-ingress"], d, "src_ingress", sv)); err != nil {
			if !fortiAPIPatch(o["src-ingress"]) {
				return fmt.Errorf("Error reading src_ingress: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("src_ingress"); ok {
			if err = d.Set("src_ingress", flattenSwitchMirrorSrcIngress(o["src-ingress"], d, "src_ingress", sv)); err != nil {
				if !fortiAPIPatch(o["src-ingress"]) {
					return fmt.Errorf("Error reading src_ingress: %v", err)
				}
			}
		}
	}

	if err = d.Set("rspan_ip", flattenSwitchMirrorRspanIp(o["rspan-ip"], d, "rspan_ip", sv)); err != nil {
		if !fortiAPIPatch(o["rspan-ip"]) {
			return fmt.Errorf("Error reading rspan_ip: %v", err)
		}
	}

	if err = d.Set("encap_mac_src", flattenSwitchMirrorEncapMacSrc(o["encap-mac-src"], d, "encap_mac_src", sv)); err != nil {
		if !fortiAPIPatch(o["encap-mac-src"]) {
			return fmt.Errorf("Error reading encap_mac_src: %v", err)
		}
	}

	if err = d.Set("strip_mirrored_traffic_tags", flattenSwitchMirrorStripMirroredTrafficTags(o["strip-mirrored-traffic-tags"], d, "strip_mirrored_traffic_tags", sv)); err != nil {
		if !fortiAPIPatch(o["strip-mirrored-traffic-tags"]) {
			return fmt.Errorf("Error reading strip_mirrored_traffic_tags: %v", err)
		}
	}

	if err = d.Set("encap_ipv4_src", flattenSwitchMirrorEncapIpv4Src(o["encap-ipv4-src"], d, "encap_ipv4_src", sv)); err != nil {
		if !fortiAPIPatch(o["encap-ipv4-src"]) {
			return fmt.Errorf("Error reading encap_ipv4_src: %v", err)
		}
	}

	if err = d.Set("encap_vlan_cfi", flattenSwitchMirrorEncapVlanCfi(o["encap-vlan-cfi"], d, "encap_vlan_cfi", sv)); err != nil {
		if !fortiAPIPatch(o["encap-vlan-cfi"]) {
			return fmt.Errorf("Error reading encap_vlan_cfi: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchMirrorName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSwitchMirrorFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchMirrorStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorEncapVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorEncapIpv4Tos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorErspanCollectorIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorSwitchingPacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorEncapVlanPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorEncapGreProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorSrcEgress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSwitchMirrorSrcEgressName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchMirrorSrcEgressName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorEncapIpv4Ttl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorEncapVlanId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorEncapVlanTpid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorEncapMacDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorSrcIngress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSwitchMirrorSrcIngressName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchMirrorSrcIngressName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorRspanIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorEncapMacSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorStripMirroredTrafficTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorEncapIpv4Src(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorEncapVlanCfi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMirrorName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchMirror(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {

		t, err := expandSwitchMirrorStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("encap_vlan"); ok {

		t, err := expandSwitchMirrorEncapVlan(d, v, "encap_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-vlan"] = t
		}
	}

	if v, ok := d.GetOk("encap_ipv4_tos"); ok {

		t, err := expandSwitchMirrorEncapIpv4Tos(d, v, "encap_ipv4_tos", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-ipv4-tos"] = t
		}
	}

	if v, ok := d.GetOk("erspan_collector_ip"); ok {

		t, err := expandSwitchMirrorErspanCollectorIp(d, v, "erspan_collector_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["erspan-collector-ip"] = t
		}
	}

	if v, ok := d.GetOk("switching_packet"); ok {

		t, err := expandSwitchMirrorSwitchingPacket(d, v, "switching_packet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switching-packet"] = t
		}
	}

	if v, ok := d.GetOkExists("encap_vlan_priority"); ok {

		t, err := expandSwitchMirrorEncapVlanPriority(d, v, "encap_vlan_priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-vlan-priority"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok {

		t, err := expandSwitchMirrorDst(d, v, "dst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("encap_gre_protocol"); ok {

		t, err := expandSwitchMirrorEncapGreProtocol(d, v, "encap_gre_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-gre-protocol"] = t
		}
	}

	if v, ok := d.GetOk("src_egress"); ok || d.HasChange("src_egress") {

		t, err := expandSwitchMirrorSrcEgress(d, v, "src_egress", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-egress"] = t
		}
	}

	if v, ok := d.GetOkExists("encap_ipv4_ttl"); ok {

		t, err := expandSwitchMirrorEncapIpv4Ttl(d, v, "encap_ipv4_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-ipv4-ttl"] = t
		}
	}

	if v, ok := d.GetOk("encap_vlan_id"); ok {

		t, err := expandSwitchMirrorEncapVlanId(d, v, "encap_vlan_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("encap_vlan_tpid"); ok {

		t, err := expandSwitchMirrorEncapVlanTpid(d, v, "encap_vlan_tpid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-vlan-tpid"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {

		t, err := expandSwitchMirrorMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("encap_mac_dst"); ok {

		t, err := expandSwitchMirrorEncapMacDst(d, v, "encap_mac_dst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-mac-dst"] = t
		}
	}

	if v, ok := d.GetOk("src_ingress"); ok || d.HasChange("src_ingress") {

		t, err := expandSwitchMirrorSrcIngress(d, v, "src_ingress", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-ingress"] = t
		}
	}

	if v, ok := d.GetOk("rspan_ip"); ok {

		t, err := expandSwitchMirrorRspanIp(d, v, "rspan_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rspan-ip"] = t
		}
	}

	if v, ok := d.GetOk("encap_mac_src"); ok {

		t, err := expandSwitchMirrorEncapMacSrc(d, v, "encap_mac_src", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-mac-src"] = t
		}
	}

	if v, ok := d.GetOk("strip_mirrored_traffic_tags"); ok {

		t, err := expandSwitchMirrorStripMirroredTrafficTags(d, v, "strip_mirrored_traffic_tags", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strip-mirrored-traffic-tags"] = t
		}
	}

	if v, ok := d.GetOk("encap_ipv4_src"); ok {

		t, err := expandSwitchMirrorEncapIpv4Src(d, v, "encap_ipv4_src", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-ipv4-src"] = t
		}
	}

	if v, ok := d.GetOkExists("encap_vlan_cfi"); ok {

		t, err := expandSwitchMirrorEncapVlanCfi(d, v, "encap_vlan_cfi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-vlan-cfi"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchMirrorName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
