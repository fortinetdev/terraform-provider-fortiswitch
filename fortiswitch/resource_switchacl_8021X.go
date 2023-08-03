// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: 802-1X Radius Dynamic Ingress Policy configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchAcl8021X() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchAcl8021XCreate,
		Read:   resourceSwitchAcl8021XRead,
		Update: resourceSwitchAcl8021XUpdate,
		Delete: resourceSwitchAcl8021XDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"access_list_entry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
									"drop": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"group": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"filter_id": &schema.Schema{
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
			"description": &schema.Schema{
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

func resourceSwitchAcl8021XCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAcl8021X(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchAcl8021X resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchAcl8021X(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchAcl8021X resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchAcl8021X")
	}

	return resourceSwitchAcl8021XRead(d, m)
}

func resourceSwitchAcl8021XUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAcl8021X(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAcl8021X resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchAcl8021X(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAcl8021X resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchAcl8021X")
	}

	return resourceSwitchAcl8021XRead(d, m)
}

func resourceSwitchAcl8021XDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchAcl8021X(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchAcl8021X resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchAcl8021XRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchAcl8021X(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAcl8021X resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchAcl8021X(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAcl8021X resource from API: %v", err)
	}
	return nil
}

func flattenSwitchAcl8021XAccessListEntry(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenSwitchAcl8021XAccessListEntryAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := i["group"]; ok {

			tmp["group"] = flattenSwitchAcl8021XAccessListEntryGroup(i["group"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSwitchAcl8021XAccessListEntryId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "classifier"
		if _, ok := i["classifier"]; ok {

			tmp["classifier"] = flattenSwitchAcl8021XAccessListEntryClassifier(i["classifier"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchAcl8021XAccessListEntryDescription(i["description"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchAcl8021XAccessListEntryAction(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "count"
	if _, ok := i["count"]; ok {

		result["count"] = flattenSwitchAcl8021XAccessListEntryActionCount(i["count"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drop"
	if _, ok := i["drop"]; ok {

		result["drop"] = flattenSwitchAcl8021XAccessListEntryActionDrop(i["drop"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchAcl8021XAccessListEntryActionCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAcl8021XAccessListEntryActionDrop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAcl8021XAccessListEntryGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAcl8021XAccessListEntryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAcl8021XAccessListEntryClassifier(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dst_mac"
	if _, ok := i["dst-mac"]; ok {

		result["dst_mac"] = flattenSwitchAcl8021XAccessListEntryClassifierDstMac(i["dst-mac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "service"
	if _, ok := i["service"]; ok {

		result["service"] = flattenSwitchAcl8021XAccessListEntryClassifierService(i["service"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dst_ip_prefix"
	if _, ok := i["dst-ip-prefix"]; ok {

		result["dst_ip_prefix"] = flattenSwitchAcl8021XAccessListEntryClassifierDstIpPrefix(i["dst-ip-prefix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_mac"
	if _, ok := i["src-mac"]; ok {

		result["src_mac"] = flattenSwitchAcl8021XAccessListEntryClassifierSrcMac(i["src-mac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ether_type"
	if _, ok := i["ether-type"]; ok {

		result["ether_type"] = flattenSwitchAcl8021XAccessListEntryClassifierEtherType(i["ether-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vlan_id"
	if _, ok := i["vlan-id"]; ok {

		result["vlan_id"] = flattenSwitchAcl8021XAccessListEntryClassifierVlanId(i["vlan-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_ip_prefix"
	if _, ok := i["src-ip-prefix"]; ok {

		result["src_ip_prefix"] = flattenSwitchAcl8021XAccessListEntryClassifierSrcIpPrefix(i["src-ip-prefix"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchAcl8021XAccessListEntryClassifierDstMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAcl8021XAccessListEntryClassifierService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAcl8021XAccessListEntryClassifierDstIpPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchAcl8021XAccessListEntryClassifierSrcMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAcl8021XAccessListEntryClassifierEtherType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAcl8021XAccessListEntryClassifierVlanId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAcl8021XAccessListEntryClassifierSrcIpPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchAcl8021XAccessListEntryDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAcl8021XFilterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAcl8021XId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAcl8021XDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchAcl8021X(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if isImportTable() {
		if err = d.Set("access_list_entry", flattenSwitchAcl8021XAccessListEntry(o["access-list-entry"], d, "access_list_entry", sv)); err != nil {
			if !fortiAPIPatch(o["access-list-entry"]) {
				return fmt.Errorf("Error reading access_list_entry: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("access_list_entry"); ok {
			if err = d.Set("access_list_entry", flattenSwitchAcl8021XAccessListEntry(o["access-list-entry"], d, "access_list_entry", sv)); err != nil {
				if !fortiAPIPatch(o["access-list-entry"]) {
					return fmt.Errorf("Error reading access_list_entry: %v", err)
				}
			}
		}
	}

	if err = d.Set("filter_id", flattenSwitchAcl8021XFilterId(o["filter-id"], d, "filter_id", sv)); err != nil {
		if !fortiAPIPatch(o["filter-id"]) {
			return fmt.Errorf("Error reading filter_id: %v", err)
		}
	}

	if err = d.Set("fswid", flattenSwitchAcl8021XId(o["id"], d, "fswid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fswid: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchAcl8021XDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	return nil
}

func flattenSwitchAcl8021XFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchAcl8021XAccessListEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandSwitchAcl8021XAccessListEntryAction(d, i["action"], pre_append, sv)
		} else {
			tmp["action"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["group"], _ = expandSwitchAcl8021XAccessListEntryGroup(d, i["group"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSwitchAcl8021XAccessListEntryId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "classifier"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["classifier"], _ = expandSwitchAcl8021XAccessListEntryClassifier(d, i["classifier"], pre_append, sv)
		} else {
			tmp["classifier"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandSwitchAcl8021XAccessListEntryDescription(d, i["description"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchAcl8021XAccessListEntryAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "count"
	if _, ok := d.GetOk(pre_append); ok {

		result["count"], _ = expandSwitchAcl8021XAccessListEntryActionCount(d, i["count"], pre_append, sv)
	}
	pre_append = pre + ".0." + "drop"
	if _, ok := d.GetOk(pre_append); ok {

		result["drop"], _ = expandSwitchAcl8021XAccessListEntryActionDrop(d, i["drop"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchAcl8021XAccessListEntryActionCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAcl8021XAccessListEntryActionDrop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAcl8021XAccessListEntryGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAcl8021XAccessListEntryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAcl8021XAccessListEntryClassifier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dst_mac"
	if _, ok := d.GetOk(pre_append); ok {

		result["dst-mac"], _ = expandSwitchAcl8021XAccessListEntryClassifierDstMac(d, i["dst_mac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "service"
	if _, ok := d.GetOk(pre_append); ok {

		result["service"], _ = expandSwitchAcl8021XAccessListEntryClassifierService(d, i["service"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dst_ip_prefix"
	if _, ok := d.GetOk(pre_append); ok {

		result["dst-ip-prefix"], _ = expandSwitchAcl8021XAccessListEntryClassifierDstIpPrefix(d, i["dst_ip_prefix"], pre_append, sv)
	}
	pre_append = pre + ".0." + "src_mac"
	if _, ok := d.GetOk(pre_append); ok {

		result["src-mac"], _ = expandSwitchAcl8021XAccessListEntryClassifierSrcMac(d, i["src_mac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ether_type"
	if _, ok := d.GetOk(pre_append); ok {

		result["ether-type"], _ = expandSwitchAcl8021XAccessListEntryClassifierEtherType(d, i["ether_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vlan_id"
	if _, ok := d.GetOk(pre_append); ok {

		result["vlan-id"], _ = expandSwitchAcl8021XAccessListEntryClassifierVlanId(d, i["vlan_id"], pre_append, sv)
	}
	pre_append = pre + ".0." + "src_ip_prefix"
	if _, ok := d.GetOk(pre_append); ok {

		result["src-ip-prefix"], _ = expandSwitchAcl8021XAccessListEntryClassifierSrcIpPrefix(d, i["src_ip_prefix"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchAcl8021XAccessListEntryClassifierDstMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAcl8021XAccessListEntryClassifierService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAcl8021XAccessListEntryClassifierDstIpPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAcl8021XAccessListEntryClassifierSrcMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAcl8021XAccessListEntryClassifierEtherType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAcl8021XAccessListEntryClassifierVlanId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAcl8021XAccessListEntryClassifierSrcIpPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAcl8021XAccessListEntryDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAcl8021XFilterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAcl8021XId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAcl8021XDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchAcl8021X(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_list_entry"); ok || d.HasChange("access_list_entry") {

		t, err := expandSwitchAcl8021XAccessListEntry(d, v, "access_list_entry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-list-entry"] = t
		}
	}

	if v, ok := d.GetOk("filter_id"); ok {

		t, err := expandSwitchAcl8021XFilterId(d, v, "filter_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-id"] = t
		}
	}

	if v, ok := d.GetOk("fswid"); ok {

		t, err := expandSwitchAcl8021XId(d, v, "fswid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchAcl8021XDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	return &obj, nil
}
