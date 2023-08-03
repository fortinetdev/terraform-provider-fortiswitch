// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: LLDP configuration profiles.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchLldpProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchLldpProfileCreate,
		Read:   resourceSwitchLldpProfileRead,
		Update: resourceSwitchLldpProfileUpdate,
		Delete: resourceSwitchLldpProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vlan_name_map": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_mclag_icl": &schema.Schema{
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"n8023_tlvs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_isl_receive_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 300),
				Optional:     true,
				Computed:     true,
			},
			"auto_isl_port_group": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 9),
				Optional:     true,
				Computed:     true,
			},
			"custom_tlvs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subtype": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"oui": &schema.Schema{
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
						"information_string": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"med_location_service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sys_location_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"n8021_tlvs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"med_tlvs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"med_network_policy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
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
						"vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4094),
							Optional:     true,
							Computed:     true,
						},
						"dscp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"assign_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"auto_isl_hello_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
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

func resourceSwitchLldpProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchLldpProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchLldpProfile resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchLldpProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchLldpProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchLldpProfile")
	}

	return resourceSwitchLldpProfileRead(d, m)
}

func resourceSwitchLldpProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchLldpProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchLldpProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchLldpProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchLldpProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchLldpProfile")
	}

	return resourceSwitchLldpProfileRead(d, m)
}

func resourceSwitchLldpProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchLldpProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchLldpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchLldpProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchLldpProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchLldpProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchLldpProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchLldpProfile resource from API: %v", err)
	}
	return nil
}

func flattenSwitchLldpProfileVlanNameMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileAutoMclagIcl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileAutoIsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfile8023Tlvs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileAutoIslReceiveTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileAutoIslPortGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileCustomTlvs(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subtype"
		if _, ok := i["subtype"]; ok {

			tmp["subtype"] = flattenSwitchLldpProfileCustomTlvsSubtype(i["subtype"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui"
		if _, ok := i["oui"]; ok {

			tmp["oui"] = flattenSwitchLldpProfileCustomTlvsOui(i["oui"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchLldpProfileCustomTlvsName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "information_string"
		if _, ok := i["information-string"]; ok {

			tmp["information_string"] = flattenSwitchLldpProfileCustomTlvsInformationString(i["information-string"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchLldpProfileCustomTlvsSubtype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileCustomTlvsOui(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileCustomTlvsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileCustomTlvsInformationString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileMedLocationService(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenSwitchLldpProfileMedLocationServiceStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sys_location_id"
		if _, ok := i["sys-location-id"]; ok {

			tmp["sys_location_id"] = flattenSwitchLldpProfileMedLocationServiceSysLocationId(i["sys-location-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchLldpProfileMedLocationServiceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchLldpProfileMedLocationServiceStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileMedLocationServiceSysLocationId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileMedLocationServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfile8021Tlvs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileMedTlvs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileMedNetworkPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenSwitchLldpProfileMedNetworkPolicyStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchLldpProfileMedNetworkPolicyName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := i["vlan"]; ok {

			tmp["vlan"] = flattenSwitchLldpProfileMedNetworkPolicyVlan(i["vlan"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := i["dscp"]; ok {

			tmp["dscp"] = flattenSwitchLldpProfileMedNetworkPolicyDscp(i["dscp"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {

			tmp["priority"] = flattenSwitchLldpProfileMedNetworkPolicyPriority(i["priority"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assign_vlan"
		if _, ok := i["assign-vlan"]; ok {

			tmp["assign_vlan"] = flattenSwitchLldpProfileMedNetworkPolicyAssignVlan(i["assign-vlan"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchLldpProfileMedNetworkPolicyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileMedNetworkPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileMedNetworkPolicyVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileMedNetworkPolicyDscp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileMedNetworkPolicyPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileMedNetworkPolicyAssignVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpProfileAutoIslHelloTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchLldpProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("vlan_name_map", flattenSwitchLldpProfileVlanNameMap(o["vlan-name-map"], d, "vlan_name_map", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-name-map"]) {
			return fmt.Errorf("Error reading vlan_name_map: %v", err)
		}
	}

	if err = d.Set("auto_mclag_icl", flattenSwitchLldpProfileAutoMclagIcl(o["auto-mclag-icl"], d, "auto_mclag_icl", sv)); err != nil {
		if !fortiAPIPatch(o["auto-mclag-icl"]) {
			return fmt.Errorf("Error reading auto_mclag_icl: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchLldpProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("auto_isl", flattenSwitchLldpProfileAutoIsl(o["auto-isl"], d, "auto_isl", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl"]) {
			return fmt.Errorf("Error reading auto_isl: %v", err)
		}
	}

	if err = d.Set("n8023_tlvs", flattenSwitchLldpProfile8023Tlvs(o["802.3-tlvs"], d, "n8023_tlvs", sv)); err != nil {
		if !fortiAPIPatch(o["802.3-tlvs"]) {
			return fmt.Errorf("Error reading n8023_tlvs: %v", err)
		}
	}

	if err = d.Set("auto_isl_receive_timeout", flattenSwitchLldpProfileAutoIslReceiveTimeout(o["auto-isl-receive-timeout"], d, "auto_isl_receive_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl-receive-timeout"]) {
			return fmt.Errorf("Error reading auto_isl_receive_timeout: %v", err)
		}
	}

	if err = d.Set("auto_isl_port_group", flattenSwitchLldpProfileAutoIslPortGroup(o["auto-isl-port-group"], d, "auto_isl_port_group", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl-port-group"]) {
			return fmt.Errorf("Error reading auto_isl_port_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_tlvs", flattenSwitchLldpProfileCustomTlvs(o["custom-tlvs"], d, "custom_tlvs", sv)); err != nil {
			if !fortiAPIPatch(o["custom-tlvs"]) {
				return fmt.Errorf("Error reading custom_tlvs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_tlvs"); ok {
			if err = d.Set("custom_tlvs", flattenSwitchLldpProfileCustomTlvs(o["custom-tlvs"], d, "custom_tlvs", sv)); err != nil {
				if !fortiAPIPatch(o["custom-tlvs"]) {
					return fmt.Errorf("Error reading custom_tlvs: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("med_location_service", flattenSwitchLldpProfileMedLocationService(o["med-location-service"], d, "med_location_service", sv)); err != nil {
			if !fortiAPIPatch(o["med-location-service"]) {
				return fmt.Errorf("Error reading med_location_service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("med_location_service"); ok {
			if err = d.Set("med_location_service", flattenSwitchLldpProfileMedLocationService(o["med-location-service"], d, "med_location_service", sv)); err != nil {
				if !fortiAPIPatch(o["med-location-service"]) {
					return fmt.Errorf("Error reading med_location_service: %v", err)
				}
			}
		}
	}

	if err = d.Set("n8021_tlvs", flattenSwitchLldpProfile8021Tlvs(o["802.1-tlvs"], d, "n8021_tlvs", sv)); err != nil {
		if !fortiAPIPatch(o["802.1-tlvs"]) {
			return fmt.Errorf("Error reading n8021_tlvs: %v", err)
		}
	}

	if err = d.Set("med_tlvs", flattenSwitchLldpProfileMedTlvs(o["med-tlvs"], d, "med_tlvs", sv)); err != nil {
		if !fortiAPIPatch(o["med-tlvs"]) {
			return fmt.Errorf("Error reading med_tlvs: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("med_network_policy", flattenSwitchLldpProfileMedNetworkPolicy(o["med-network-policy"], d, "med_network_policy", sv)); err != nil {
			if !fortiAPIPatch(o["med-network-policy"]) {
				return fmt.Errorf("Error reading med_network_policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("med_network_policy"); ok {
			if err = d.Set("med_network_policy", flattenSwitchLldpProfileMedNetworkPolicy(o["med-network-policy"], d, "med_network_policy", sv)); err != nil {
				if !fortiAPIPatch(o["med-network-policy"]) {
					return fmt.Errorf("Error reading med_network_policy: %v", err)
				}
			}
		}
	}

	if err = d.Set("auto_isl_hello_timer", flattenSwitchLldpProfileAutoIslHelloTimer(o["auto-isl-hello-timer"], d, "auto_isl_hello_timer", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl-hello-timer"]) {
			return fmt.Errorf("Error reading auto_isl_hello_timer: %v", err)
		}
	}

	return nil
}

func flattenSwitchLldpProfileFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchLldpProfileVlanNameMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileAutoMclagIcl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileAutoIsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfile8023Tlvs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileAutoIslReceiveTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileAutoIslPortGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileCustomTlvs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subtype"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["subtype"], _ = expandSwitchLldpProfileCustomTlvsSubtype(d, i["subtype"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["oui"], _ = expandSwitchLldpProfileCustomTlvsOui(d, i["oui"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSwitchLldpProfileCustomTlvsName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "information_string"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["information-string"], _ = expandSwitchLldpProfileCustomTlvsInformationString(d, i["information_string"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchLldpProfileCustomTlvsSubtype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileCustomTlvsOui(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileCustomTlvsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileCustomTlvsInformationString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileMedLocationService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandSwitchLldpProfileMedLocationServiceStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sys_location_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sys-location-id"], _ = expandSwitchLldpProfileMedLocationServiceSysLocationId(d, i["sys_location_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSwitchLldpProfileMedLocationServiceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchLldpProfileMedLocationServiceStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileMedLocationServiceSysLocationId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileMedLocationServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfile8021Tlvs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileMedTlvs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileMedNetworkPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandSwitchLldpProfileMedNetworkPolicyStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSwitchLldpProfileMedNetworkPolicyName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["vlan"], _ = expandSwitchLldpProfileMedNetworkPolicyVlan(d, i["vlan"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dscp"], _ = expandSwitchLldpProfileMedNetworkPolicyDscp(d, i["dscp"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["priority"], _ = expandSwitchLldpProfileMedNetworkPolicyPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assign_vlan"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["assign-vlan"], _ = expandSwitchLldpProfileMedNetworkPolicyAssignVlan(d, i["assign_vlan"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchLldpProfileMedNetworkPolicyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileMedNetworkPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileMedNetworkPolicyVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileMedNetworkPolicyDscp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileMedNetworkPolicyPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileMedNetworkPolicyAssignVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpProfileAutoIslHelloTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchLldpProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("vlan_name_map"); ok {

		t, err := expandSwitchLldpProfileVlanNameMap(d, v, "vlan_name_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-name-map"] = t
		}
	}

	if v, ok := d.GetOk("auto_mclag_icl"); ok {

		t, err := expandSwitchLldpProfileAutoMclagIcl(d, v, "auto_mclag_icl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-mclag-icl"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchLldpProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl"); ok {

		t, err := expandSwitchLldpProfileAutoIsl(d, v, "auto_isl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl"] = t
		}
	}

	if v, ok := d.GetOk("n8023_tlvs"); ok {

		t, err := expandSwitchLldpProfile8023Tlvs(d, v, "n8023_tlvs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802.3-tlvs"] = t
		}
	}

	if v, ok := d.GetOkExists("auto_isl_receive_timeout"); ok {

		t, err := expandSwitchLldpProfileAutoIslReceiveTimeout(d, v, "auto_isl_receive_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-receive-timeout"] = t
		}
	}

	if v, ok := d.GetOkExists("auto_isl_port_group"); ok {

		t, err := expandSwitchLldpProfileAutoIslPortGroup(d, v, "auto_isl_port_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-port-group"] = t
		}
	}

	if v, ok := d.GetOk("custom_tlvs"); ok || d.HasChange("custom_tlvs") {

		t, err := expandSwitchLldpProfileCustomTlvs(d, v, "custom_tlvs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("med_location_service"); ok || d.HasChange("med_location_service") {

		t, err := expandSwitchLldpProfileMedLocationService(d, v, "med_location_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["med-location-service"] = t
		}
	}

	if v, ok := d.GetOk("n8021_tlvs"); ok {

		t, err := expandSwitchLldpProfile8021Tlvs(d, v, "n8021_tlvs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802.1-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("med_tlvs"); ok {

		t, err := expandSwitchLldpProfileMedTlvs(d, v, "med_tlvs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["med-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("med_network_policy"); ok || d.HasChange("med_network_policy") {

		t, err := expandSwitchLldpProfileMedNetworkPolicy(d, v, "med_network_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["med-network-policy"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_hello_timer"); ok {

		t, err := expandSwitchLldpProfileAutoIslHelloTimer(d, v, "auto_isl_hello_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-hello-timer"] = t
		}
	}

	return &obj, nil
}
