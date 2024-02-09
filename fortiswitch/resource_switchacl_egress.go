// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Egress Policy configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchAclEgress() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchAclEgressCreate,
		Read:   resourceSwitchAclEgressRead,
		Update: resourceSwitchAclEgressUpdate,
		Delete: resourceSwitchAclEgressDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group": &schema.Schema{
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
			"schedule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"schedule_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"classifier": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"service": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"dst_ip_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"src_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ether_type": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"vlan_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"src_ip_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"action": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"redirect": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"remark_dscp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"drop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"policer": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mirror": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"count_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"outer_vlan_tag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"fswid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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

func resourceSwitchAclEgressCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAclEgress(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchAclEgress resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchAclEgress(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchAclEgress resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchAclEgress")
	}

	return resourceSwitchAclEgressRead(d, m)
}

func resourceSwitchAclEgressUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAclEgress(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAclEgress resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchAclEgress(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAclEgress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchAclEgress")
	}

	return resourceSwitchAclEgressRead(d, m)
}

func resourceSwitchAclEgressDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchAclEgress(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchAclEgress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchAclEgressRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchAclEgress(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAclEgress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchAclEgress(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAclEgress resource from API: %v", err)
	}
	return nil
}

func flattenSwitchAclEgressStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressSchedule(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "schedule_name"
		if _, ok := i["schedule-name"]; ok {

			tmp["schedule_name"] = flattenSwitchAclEgressScheduleScheduleName(i["schedule-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "schedule_name", d)
	return result
}

func flattenSwitchAclEgressScheduleScheduleName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressClassifier(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dst_mac"
	if _, ok := i["dst-mac"]; ok {

		result["dst_mac"] = flattenSwitchAclEgressClassifierDstMac(i["dst-mac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cos"
	if _, ok := i["cos"]; ok {

		result["cos"] = flattenSwitchAclEgressClassifierCos(i["cos"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "service"
	if _, ok := i["service"]; ok {

		result["service"] = flattenSwitchAclEgressClassifierService(i["service"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dst_ip_prefix"
	if _, ok := i["dst-ip-prefix"]; ok {

		result["dst_ip_prefix"] = flattenSwitchAclEgressClassifierDstIpPrefix(i["dst-ip-prefix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_mac"
	if _, ok := i["src-mac"]; ok {

		result["src_mac"] = flattenSwitchAclEgressClassifierSrcMac(i["src-mac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {

		result["dscp"] = flattenSwitchAclEgressClassifierDscp(i["dscp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ether_type"
	if _, ok := i["ether-type"]; ok {

		result["ether_type"] = flattenSwitchAclEgressClassifierEtherType(i["ether-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vlan_id"
	if _, ok := i["vlan-id"]; ok {

		result["vlan_id"] = flattenSwitchAclEgressClassifierVlanId(i["vlan-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_ip_prefix"
	if _, ok := i["src-ip-prefix"]; ok {

		result["src_ip_prefix"] = flattenSwitchAclEgressClassifierSrcIpPrefix(i["src-ip-prefix"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchAclEgressClassifierDstMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressClassifierCos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v == "" {
		return nil
	}
	return v
}

func flattenSwitchAclEgressClassifierService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressClassifierDstIpPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchAclEgressClassifierSrcMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressClassifierDscp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v == "" {
		return nil
	}
	return v
}

func flattenSwitchAclEgressClassifierEtherType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressClassifierVlanId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressClassifierSrcIpPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchAclEgressAction(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "count"
	if _, ok := i["count"]; ok {

		result["count"] = flattenSwitchAclEgressActionCount(i["count"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "redirect"
	if _, ok := i["redirect"]; ok {

		result["redirect"] = flattenSwitchAclEgressActionRedirect(i["redirect"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "remark_dscp"
	if _, ok := i["remark-dscp"]; ok {

		result["remark_dscp"] = flattenSwitchAclEgressActionRemarkDscp(i["remark-dscp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drop"
	if _, ok := i["drop"]; ok {

		result["drop"] = flattenSwitchAclEgressActionDrop(i["drop"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "policer"
	if _, ok := i["policer"]; ok {

		result["policer"] = flattenSwitchAclEgressActionPolicer(i["policer"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mirror"
	if _, ok := i["mirror"]; ok {

		result["mirror"] = flattenSwitchAclEgressActionMirror(i["mirror"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "count_type"
	if _, ok := i["count-type"]; ok {

		result["count_type"] = flattenSwitchAclEgressActionCountType(i["count-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "outer_vlan_tag"
	if _, ok := i["outer-vlan-tag"]; ok {

		result["outer_vlan_tag"] = flattenSwitchAclEgressActionOuterVlanTag(i["outer-vlan-tag"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchAclEgressActionCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressActionRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressActionRemarkDscp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v == "" {
		return nil
	}
	return v
}

func flattenSwitchAclEgressActionDrop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressActionPolicer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressActionMirror(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressActionCountType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressActionOuterVlanTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclEgressId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchAclEgress(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSwitchAclEgressStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("group", flattenSwitchAclEgressGroup(o["group"], d, "group", sv)); err != nil {
		if !fortiAPIPatch(o["group"]) {
			return fmt.Errorf("Error reading group: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchAclEgressDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("schedule", flattenSwitchAclEgressSchedule(o["schedule"], d, "schedule", sv)); err != nil {
			if !fortiAPIPatch(o["schedule"]) {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("schedule"); ok {
			if err = d.Set("schedule", flattenSwitchAclEgressSchedule(o["schedule"], d, "schedule", sv)); err != nil {
				if !fortiAPIPatch(o["schedule"]) {
					return fmt.Errorf("Error reading schedule: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("classifier", flattenSwitchAclEgressClassifier(o["classifier"], d, "classifier", sv)); err != nil {
			if !fortiAPIPatch(o["classifier"]) {
				return fmt.Errorf("Error reading classifier: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("classifier"); ok {
			if err = d.Set("classifier", flattenSwitchAclEgressClassifier(o["classifier"], d, "classifier", sv)); err != nil {
				if !fortiAPIPatch(o["classifier"]) {
					return fmt.Errorf("Error reading classifier: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("action", flattenSwitchAclEgressAction(o["action"], d, "action", sv)); err != nil {
			if !fortiAPIPatch(o["action"]) {
				return fmt.Errorf("Error reading action: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("action"); ok {
			if err = d.Set("action", flattenSwitchAclEgressAction(o["action"], d, "action", sv)); err != nil {
				if !fortiAPIPatch(o["action"]) {
					return fmt.Errorf("Error reading action: %v", err)
				}
			}
		}
	}

	if err = d.Set("interface", flattenSwitchAclEgressInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("fswid", flattenSwitchAclEgressId(o["id"], d, "fswid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fswid: %v", err)
		}
	}

	return nil
}

func flattenSwitchAclEgressFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchAclEgressStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressSchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "schedule_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["schedule-name"], _ = expandSwitchAclEgressScheduleScheduleName(d, i["schedule_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchAclEgressScheduleScheduleName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressClassifier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dst_mac"
	if _, ok := d.GetOk(pre_append); ok {

		result["dst-mac"], _ = expandSwitchAclEgressClassifierDstMac(d, i["dst_mac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cos"
	if _, ok := d.GetOk(pre_append); ok {

		result["cos"], _ = expandSwitchAclEgressClassifierCos(d, i["cos"], pre_append, sv)
	}
	pre_append = pre + ".0." + "service"
	if _, ok := d.GetOk(pre_append); ok {

		result["service"], _ = expandSwitchAclEgressClassifierService(d, i["service"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dst_ip_prefix"
	if _, ok := d.GetOk(pre_append); ok {

		result["dst-ip-prefix"], _ = expandSwitchAclEgressClassifierDstIpPrefix(d, i["dst_ip_prefix"], pre_append, sv)
	}
	pre_append = pre + ".0." + "src_mac"
	if _, ok := d.GetOk(pre_append); ok {

		result["src-mac"], _ = expandSwitchAclEgressClassifierSrcMac(d, i["src_mac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok {

		result["dscp"], _ = expandSwitchAclEgressClassifierDscp(d, i["dscp"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ether_type"
	if _, ok := d.GetOk(pre_append); ok {

		result["ether-type"], _ = expandSwitchAclEgressClassifierEtherType(d, i["ether_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vlan_id"
	if _, ok := d.GetOk(pre_append); ok {

		result["vlan-id"], _ = expandSwitchAclEgressClassifierVlanId(d, i["vlan_id"], pre_append, sv)
	}
	pre_append = pre + ".0." + "src_ip_prefix"
	if _, ok := d.GetOk(pre_append); ok {

		result["src-ip-prefix"], _ = expandSwitchAclEgressClassifierSrcIpPrefix(d, i["src_ip_prefix"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchAclEgressClassifierDstMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressClassifierCos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressClassifierService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressClassifierDstIpPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressClassifierSrcMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressClassifierDscp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressClassifierEtherType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressClassifierVlanId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressClassifierSrcIpPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "count"
	if _, ok := d.GetOk(pre_append); ok {

		result["count"], _ = expandSwitchAclEgressActionCount(d, i["count"], pre_append, sv)
	}
	pre_append = pre + ".0." + "redirect"
	if _, ok := d.GetOk(pre_append); ok {

		result["redirect"], _ = expandSwitchAclEgressActionRedirect(d, i["redirect"], pre_append, sv)
	}
	pre_append = pre + ".0." + "remark_dscp"
	if _, ok := d.GetOk(pre_append); ok {

		result["remark-dscp"], _ = expandSwitchAclEgressActionRemarkDscp(d, i["remark_dscp"], pre_append, sv)
	}
	pre_append = pre + ".0." + "drop"
	if _, ok := d.GetOk(pre_append); ok {

		result["drop"], _ = expandSwitchAclEgressActionDrop(d, i["drop"], pre_append, sv)
	}
	pre_append = pre + ".0." + "policer"
	if _, ok := d.GetOk(pre_append); ok {

		result["policer"], _ = expandSwitchAclEgressActionPolicer(d, i["policer"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mirror"
	if _, ok := d.GetOk(pre_append); ok {

		result["mirror"], _ = expandSwitchAclEgressActionMirror(d, i["mirror"], pre_append, sv)
	}
	pre_append = pre + ".0." + "count_type"
	if _, ok := d.GetOk(pre_append); ok {

		result["count-type"], _ = expandSwitchAclEgressActionCountType(d, i["count_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "outer_vlan_tag"
	if _, ok := d.GetOk(pre_append); ok {

		result["outer-vlan-tag"], _ = expandSwitchAclEgressActionOuterVlanTag(d, i["outer_vlan_tag"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchAclEgressActionCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressActionRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressActionRemarkDscp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressActionDrop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressActionPolicer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressActionMirror(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressActionCountType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressActionOuterVlanTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclEgressId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchAclEgress(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {

		t, err := expandSwitchAclEgressStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("group"); ok {

		t, err := expandSwitchAclEgressGroup(d, v, "group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchAclEgressDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {

		t, err := expandSwitchAclEgressSchedule(d, v, "schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("classifier"); ok {

		t, err := expandSwitchAclEgressClassifier(d, v, "classifier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["classifier"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {

		t, err := expandSwitchAclEgressAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandSwitchAclEgressInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("fswid"); ok {

		t, err := expandSwitchAclEgressId(d, v, "fswid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
