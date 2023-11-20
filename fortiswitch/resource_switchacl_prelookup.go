// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Prelookup Policy configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchAclPrelookup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchAclPrelookupCreate,
		Read:   resourceSwitchAclPrelookupRead,
		Update: resourceSwitchAclPrelookupUpdate,
		Delete: resourceSwitchAclPrelookupDelete,

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
						"remark_cos": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"drop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"outer_vlan_tag": &schema.Schema{
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
					},
				},
			},
			"interface_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceSwitchAclPrelookupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAclPrelookup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchAclPrelookup resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchAclPrelookup(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchAclPrelookup resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchAclPrelookup")
	}

	return resourceSwitchAclPrelookupRead(d, m)
}

func resourceSwitchAclPrelookupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAclPrelookup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAclPrelookup resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchAclPrelookup(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAclPrelookup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchAclPrelookup")
	}

	return resourceSwitchAclPrelookupRead(d, m)
}

func resourceSwitchAclPrelookupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchAclPrelookup(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchAclPrelookup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchAclPrelookupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchAclPrelookup(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAclPrelookup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchAclPrelookup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAclPrelookup resource from API: %v", err)
	}
	return nil
}

func flattenSwitchAclPrelookupStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupSchedule(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["schedule_name"] = flattenSwitchAclPrelookupScheduleScheduleName(i["schedule-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "schedule_name", d)
	return result
}

func flattenSwitchAclPrelookupScheduleScheduleName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupClassifier(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dst_mac"
	if _, ok := i["dst-mac"]; ok {

		result["dst_mac"] = flattenSwitchAclPrelookupClassifierDstMac(i["dst-mac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cos"
	if _, ok := i["cos"]; ok {

		result["cos"] = flattenSwitchAclPrelookupClassifierCos(i["cos"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "service"
	if _, ok := i["service"]; ok {

		result["service"] = flattenSwitchAclPrelookupClassifierService(i["service"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dst_ip_prefix"
	if _, ok := i["dst-ip-prefix"]; ok {

		result["dst_ip_prefix"] = flattenSwitchAclPrelookupClassifierDstIpPrefix(i["dst-ip-prefix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_mac"
	if _, ok := i["src-mac"]; ok {

		result["src_mac"] = flattenSwitchAclPrelookupClassifierSrcMac(i["src-mac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {

		result["dscp"] = flattenSwitchAclPrelookupClassifierDscp(i["dscp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ether_type"
	if _, ok := i["ether-type"]; ok {

		result["ether_type"] = flattenSwitchAclPrelookupClassifierEtherType(i["ether-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vlan_id"
	if _, ok := i["vlan-id"]; ok {

		result["vlan_id"] = flattenSwitchAclPrelookupClassifierVlanId(i["vlan-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_ip_prefix"
	if _, ok := i["src-ip-prefix"]; ok {

		result["src_ip_prefix"] = flattenSwitchAclPrelookupClassifierSrcIpPrefix(i["src-ip-prefix"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchAclPrelookupClassifierDstMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupClassifierCos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupClassifierService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupClassifierDstIpPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchAclPrelookupClassifierSrcMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupClassifierDscp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupClassifierEtherType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupClassifierVlanId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupClassifierSrcIpPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchAclPrelookupAction(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "count"
	if _, ok := i["count"]; ok {

		result["count"] = flattenSwitchAclPrelookupActionCount(i["count"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "remark_cos"
	if _, ok := i["remark-cos"]; ok {

		result["remark_cos"] = flattenSwitchAclPrelookupActionRemarkCos(i["remark-cos"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drop"
	if _, ok := i["drop"]; ok {

		result["drop"] = flattenSwitchAclPrelookupActionDrop(i["drop"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "outer_vlan_tag"
	if _, ok := i["outer-vlan-tag"]; ok {

		result["outer_vlan_tag"] = flattenSwitchAclPrelookupActionOuterVlanTag(i["outer-vlan-tag"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cos_queue"
	if _, ok := i["cos-queue"]; ok {

		result["cos_queue"] = flattenSwitchAclPrelookupActionCosQueue(i["cos-queue"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchAclPrelookupActionCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupActionRemarkCos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupActionDrop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupActionOuterVlanTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupActionCosQueue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupInterfaceAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPrelookupId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchAclPrelookup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSwitchAclPrelookupStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("group", flattenSwitchAclPrelookupGroup(o["group"], d, "group", sv)); err != nil {
		if !fortiAPIPatch(o["group"]) {
			return fmt.Errorf("Error reading group: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchAclPrelookupDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("schedule", flattenSwitchAclPrelookupSchedule(o["schedule"], d, "schedule", sv)); err != nil {
			if !fortiAPIPatch(o["schedule"]) {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("schedule"); ok {
			if err = d.Set("schedule", flattenSwitchAclPrelookupSchedule(o["schedule"], d, "schedule", sv)); err != nil {
				if !fortiAPIPatch(o["schedule"]) {
					return fmt.Errorf("Error reading schedule: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("classifier", flattenSwitchAclPrelookupClassifier(o["classifier"], d, "classifier", sv)); err != nil {
			if !fortiAPIPatch(o["classifier"]) {
				return fmt.Errorf("Error reading classifier: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("classifier"); ok {
			if err = d.Set("classifier", flattenSwitchAclPrelookupClassifier(o["classifier"], d, "classifier", sv)); err != nil {
				if !fortiAPIPatch(o["classifier"]) {
					return fmt.Errorf("Error reading classifier: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("action", flattenSwitchAclPrelookupAction(o["action"], d, "action", sv)); err != nil {
			if !fortiAPIPatch(o["action"]) {
				return fmt.Errorf("Error reading action: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("action"); ok {
			if err = d.Set("action", flattenSwitchAclPrelookupAction(o["action"], d, "action", sv)); err != nil {
				if !fortiAPIPatch(o["action"]) {
					return fmt.Errorf("Error reading action: %v", err)
				}
			}
		}
	}

	if err = d.Set("interface_all", flattenSwitchAclPrelookupInterfaceAll(o["interface-all"], d, "interface_all", sv)); err != nil {
		if !fortiAPIPatch(o["interface-all"]) {
			return fmt.Errorf("Error reading interface_all: %v", err)
		}
	}

	if err = d.Set("interface", flattenSwitchAclPrelookupInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("fswid", flattenSwitchAclPrelookupId(o["id"], d, "fswid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fswid: %v", err)
		}
	}

	return nil
}

func flattenSwitchAclPrelookupFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchAclPrelookupStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupSchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["schedule-name"], _ = expandSwitchAclPrelookupScheduleScheduleName(d, i["schedule_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchAclPrelookupScheduleScheduleName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupClassifier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dst_mac"
	if _, ok := d.GetOk(pre_append); ok {

		result["dst-mac"], _ = expandSwitchAclPrelookupClassifierDstMac(d, i["dst_mac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cos"
	if _, ok := d.GetOk(pre_append); ok {

		result["cos"], _ = expandSwitchAclPrelookupClassifierCos(d, i["cos"], pre_append, sv)
	}
	pre_append = pre + ".0." + "service"
	if _, ok := d.GetOk(pre_append); ok {

		result["service"], _ = expandSwitchAclPrelookupClassifierService(d, i["service"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dst_ip_prefix"
	if _, ok := d.GetOk(pre_append); ok {

		result["dst-ip-prefix"], _ = expandSwitchAclPrelookupClassifierDstIpPrefix(d, i["dst_ip_prefix"], pre_append, sv)
	}
	pre_append = pre + ".0." + "src_mac"
	if _, ok := d.GetOk(pre_append); ok {

		result["src-mac"], _ = expandSwitchAclPrelookupClassifierSrcMac(d, i["src_mac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok {

		result["dscp"], _ = expandSwitchAclPrelookupClassifierDscp(d, i["dscp"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ether_type"
	if _, ok := d.GetOk(pre_append); ok {

		result["ether-type"], _ = expandSwitchAclPrelookupClassifierEtherType(d, i["ether_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vlan_id"
	if _, ok := d.GetOk(pre_append); ok {

		result["vlan-id"], _ = expandSwitchAclPrelookupClassifierVlanId(d, i["vlan_id"], pre_append, sv)
	}
	pre_append = pre + ".0." + "src_ip_prefix"
	if _, ok := d.GetOk(pre_append); ok {

		result["src-ip-prefix"], _ = expandSwitchAclPrelookupClassifierSrcIpPrefix(d, i["src_ip_prefix"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchAclPrelookupClassifierDstMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupClassifierCos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupClassifierService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupClassifierDstIpPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupClassifierSrcMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupClassifierDscp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupClassifierEtherType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupClassifierVlanId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupClassifierSrcIpPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "count"
	if _, ok := d.GetOk(pre_append); ok {

		result["count"], _ = expandSwitchAclPrelookupActionCount(d, i["count"], pre_append, sv)
	}
	pre_append = pre + ".0." + "remark_cos"
	if _, ok := d.GetOk(pre_append); ok {

		result["remark-cos"], _ = expandSwitchAclPrelookupActionRemarkCos(d, i["remark_cos"], pre_append, sv)
	}
	pre_append = pre + ".0." + "drop"
	if _, ok := d.GetOk(pre_append); ok {

		result["drop"], _ = expandSwitchAclPrelookupActionDrop(d, i["drop"], pre_append, sv)
	}
	pre_append = pre + ".0." + "outer_vlan_tag"
	if _, ok := d.GetOk(pre_append); ok {

		result["outer-vlan-tag"], _ = expandSwitchAclPrelookupActionOuterVlanTag(d, i["outer_vlan_tag"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cos_queue"
	if _, ok := d.GetOk(pre_append); ok {

		result["cos-queue"], _ = expandSwitchAclPrelookupActionCosQueue(d, i["cos_queue"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchAclPrelookupActionCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupActionRemarkCos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupActionDrop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupActionOuterVlanTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupActionCosQueue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupInterfaceAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPrelookupId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchAclPrelookup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {

		t, err := expandSwitchAclPrelookupStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("group"); ok {

		t, err := expandSwitchAclPrelookupGroup(d, v, "group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchAclPrelookupDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {

		t, err := expandSwitchAclPrelookupSchedule(d, v, "schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("classifier"); ok {

		t, err := expandSwitchAclPrelookupClassifier(d, v, "classifier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["classifier"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {

		t, err := expandSwitchAclPrelookupAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("interface_all"); ok {

		t, err := expandSwitchAclPrelookupInterfaceAll(d, v, "interface_all", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-all"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandSwitchAclPrelookupInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("fswid"); ok {

		t, err := expandSwitchAclPrelookupId(d, v, "fswid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
