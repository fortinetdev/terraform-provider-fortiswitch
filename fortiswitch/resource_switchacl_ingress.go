// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Ingress Policy configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchAclIngress() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchAclIngressCreate,
		Read:   resourceSwitchAclIngressRead,
		Update: resourceSwitchAclIngressUpdate,
		Delete: resourceSwitchAclIngressDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ingress_interface_all": &schema.Schema{
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
			"ingress_interface": &schema.Schema{
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
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"service": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"l3_interface": &schema.Schema{
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
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"dst_ip6_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ether_type": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"vlan_id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"src_ip_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"src_ip6_prefix": &schema.Schema{
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
						"redirect_bcast_cpu": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"redirect_physical_port": &schema.Schema{
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
						"cos_queue": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"remark_dscp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"egress_mask": &schema.Schema{
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
						"drop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"redirect_bcast_no_cpu": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"policer": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"remark_cos": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
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
						"cpu_cos_queue": &schema.Schema{
							Type: schema.TypeString,
							//ValidateFunc: validation.IntBetween(17, 25),
							Optional: true,
							Computed: true,
						},
						"outer_vlan_tag": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"group": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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

func resourceSwitchAclIngressCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAclIngress(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchAclIngress resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchAclIngress(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchAclIngress resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(fmt.Sprintf("%v", o["mkey"]))
	} else {
		d.SetId("SwitchAclIngress")
	}

	return resourceSwitchAclIngressRead(d, m)
}

func resourceSwitchAclIngressUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAclIngress(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAclIngress resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchAclIngress(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAclIngress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(fmt.Sprintf("%v", o["mkey"]))
	} else {
		d.SetId("SwitchAclIngress")
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchAclIngress")
	}

	return resourceSwitchAclIngressRead(d, m)
}

func resourceSwitchAclIngressDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchAclIngress(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchAclIngress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchAclIngressRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchAclIngress(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAclIngress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchAclIngress(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAclIngress resource from API: %v", err)
	}
	return nil
}

func flattenSwitchAclIngressStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressIngressInterfaceAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressSchedule(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["schedule_name"] = flattenSwitchAclIngressScheduleScheduleName(i["schedule-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "schedule_name", d)
	return result
}

func flattenSwitchAclIngressScheduleScheduleName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressIngressInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["member_name"] = flattenSwitchAclIngressIngressInterfaceMemberName(i["member-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "member_name", d)
	return result
}

func flattenSwitchAclIngressIngressInterfaceMemberName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressClassifier(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dst_mac"
	if _, ok := i["dst-mac"]; ok {

		result["dst_mac"] = flattenSwitchAclIngressClassifierDstMac(i["dst-mac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cos"
	if _, ok := i["cos"]; ok {

		result["cos"] = flattenSwitchAclIngressClassifierCos(i["cos"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "service"
	if _, ok := i["service"]; ok {

		result["service"] = flattenSwitchAclIngressClassifierService(i["service"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "l3_interface"
	if _, ok := i["l3-interface"]; ok {

		result["l3_interface"] = flattenSwitchAclIngressClassifierL3Interface(i["l3-interface"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dst_ip_prefix"
	if _, ok := i["dst-ip-prefix"]; ok {

		result["dst_ip_prefix"] = flattenSwitchAclIngressClassifierDstIpPrefix(i["dst-ip-prefix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_mac"
	if _, ok := i["src-mac"]; ok {

		result["src_mac"] = flattenSwitchAclIngressClassifierSrcMac(i["src-mac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {

		result["dscp"] = flattenSwitchAclIngressClassifierDscp(i["dscp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dst_ip6_prefix"
	if _, ok := i["dst-ip6-prefix"]; ok {

		result["dst_ip6_prefix"] = flattenSwitchAclIngressClassifierDstIp6Prefix(i["dst-ip6-prefix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ether_type"
	if _, ok := i["ether-type"]; ok {

		result["ether_type"] = flattenSwitchAclIngressClassifierEtherType(i["ether-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vlan_id"
	if _, ok := i["vlan-id"]; ok {

		result["vlan_id"] = flattenSwitchAclIngressClassifierVlanId(i["vlan-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_ip_prefix"
	if _, ok := i["src-ip-prefix"]; ok {

		result["src_ip_prefix"] = flattenSwitchAclIngressClassifierSrcIpPrefix(i["src-ip-prefix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_ip6_prefix"
	if _, ok := i["src-ip6-prefix"]; ok {

		result["src_ip6_prefix"] = flattenSwitchAclIngressClassifierSrcIp6Prefix(i["src-ip6-prefix"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchAclIngressClassifierDstMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressClassifierCos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v == "" || v == "none" || reflect.DeepEqual(v, []interface{}{}) {
		return nil
	}
	return v
}

func flattenSwitchAclIngressClassifierService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressClassifierL3Interface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressClassifierDstIpPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchAclIngressClassifierSrcMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressClassifierDscp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v == "" || v == "none" || reflect.DeepEqual(v, []interface{}{}) {
		return nil
	}

	return v
}

func flattenSwitchAclIngressClassifierDstIp6Prefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressClassifierEtherType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressClassifierVlanId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressClassifierSrcIpPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchAclIngressClassifierSrcIp6Prefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressAction(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "count"
	if _, ok := i["count"]; ok {

		result["count"] = flattenSwitchAclIngressActionCount(i["count"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "redirect"
	if _, ok := i["redirect"]; ok {

		result["redirect"] = flattenSwitchAclIngressActionRedirect(i["redirect"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "redirect_bcast_cpu"
	if _, ok := i["redirect-bcast-cpu"]; ok {

		result["redirect_bcast_cpu"] = flattenSwitchAclIngressActionRedirectBcastCpu(i["redirect-bcast-cpu"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "redirect_physical_port"
	if _, ok := i["redirect-physical-port"]; ok {

		result["redirect_physical_port"] = flattenSwitchAclIngressActionRedirectPhysicalPort(i["redirect-physical-port"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cos_queue"
	if _, ok := i["cos-queue"]; ok {

		result["cos_queue"] = flattenSwitchAclIngressActionCosQueue(i["cos-queue"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "remark_dscp"
	if _, ok := i["remark-dscp"]; ok {

		result["remark_dscp"] = flattenSwitchAclIngressActionRemarkDscp(i["remark-dscp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "egress_mask"
	if _, ok := i["egress-mask"]; ok {

		result["egress_mask"] = flattenSwitchAclIngressActionEgressMask(i["egress-mask"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drop"
	if _, ok := i["drop"]; ok {

		result["drop"] = flattenSwitchAclIngressActionDrop(i["drop"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "redirect_bcast_no_cpu"
	if _, ok := i["redirect-bcast-no-cpu"]; ok {

		result["redirect_bcast_no_cpu"] = flattenSwitchAclIngressActionRedirectBcastNoCpu(i["redirect-bcast-no-cpu"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "policer"
	if _, ok := i["policer"]; ok {

		result["policer"] = flattenSwitchAclIngressActionPolicer(i["policer"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "remark_cos"
	if _, ok := i["remark-cos"]; ok {

		result["remark_cos"] = flattenSwitchAclIngressActionRemarkCos(i["remark-cos"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mirror"
	if _, ok := i["mirror"]; ok {

		result["mirror"] = flattenSwitchAclIngressActionMirror(i["mirror"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "count_type"
	if _, ok := i["count-type"]; ok {

		result["count_type"] = flattenSwitchAclIngressActionCountType(i["count-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cpu_cos_queue"
	if _, ok := i["cpu-cos-queue"]; ok {

		result["cpu_cos_queue"] = flattenSwitchAclIngressActionCpuCosQueue(i["cpu-cos-queue"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "outer_vlan_tag"
	if _, ok := i["outer-vlan-tag"]; ok {

		result["outer_vlan_tag"] = flattenSwitchAclIngressActionOuterVlanTag(i["outer-vlan-tag"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchAclIngressActionCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressActionRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressActionRedirectBcastCpu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressActionRedirectPhysicalPort(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["member_name"] = flattenSwitchAclIngressActionRedirectPhysicalPortMemberName(i["member-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "member_name", d)
	return result
}

func flattenSwitchAclIngressActionRedirectPhysicalPortMemberName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressActionCosQueue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v == "" || v == "none" || reflect.DeepEqual(v, []interface{}{}) {
		return nil
	}

	return v
}

func flattenSwitchAclIngressActionRemarkDscp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v == "" || v == "none" || reflect.DeepEqual(v, []interface{}{}) {
		return nil
	}

	return v
}

func flattenSwitchAclIngressActionEgressMask(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["member_name"] = flattenSwitchAclIngressActionEgressMaskMemberName(i["member-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "member_name", d)
	return result
}

func flattenSwitchAclIngressActionEgressMaskMemberName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressActionDrop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressActionRedirectBcastNoCpu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressActionPolicer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressActionRemarkCos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v == "" || v == "none" || reflect.DeepEqual(v, []interface{}{}) {
		return nil
	}

	return v
}

func flattenSwitchAclIngressActionMirror(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressActionCountType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressActionCpuCosQueue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressActionOuterVlanTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclIngressId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchAclIngress(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSwitchAclIngressStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("ingress_interface_all", flattenSwitchAclIngressIngressInterfaceAll(o["ingress-interface-all"], d, "ingress_interface_all", sv)); err != nil {
		if !fortiAPIPatch(o["ingress-interface-all"]) {
			return fmt.Errorf("Error reading ingress_interface_all: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchAclIngressDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("schedule", flattenSwitchAclIngressSchedule(o["schedule"], d, "schedule", sv)); err != nil {
			if !fortiAPIPatch(o["schedule"]) {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("schedule"); ok {
			if err = d.Set("schedule", flattenSwitchAclIngressSchedule(o["schedule"], d, "schedule", sv)); err != nil {
				if !fortiAPIPatch(o["schedule"]) {
					return fmt.Errorf("Error reading schedule: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ingress_interface", flattenSwitchAclIngressIngressInterface(o["ingress-interface"], d, "ingress_interface", sv)); err != nil {
			if !fortiAPIPatch(o["ingress-interface"]) {
				return fmt.Errorf("Error reading ingress_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ingress_interface"); ok {
			if err = d.Set("ingress_interface", flattenSwitchAclIngressIngressInterface(o["ingress-interface"], d, "ingress_interface", sv)); err != nil {
				if !fortiAPIPatch(o["ingress-interface"]) {
					return fmt.Errorf("Error reading ingress_interface: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("classifier", flattenSwitchAclIngressClassifier(o["classifier"], d, "classifier", sv)); err != nil {
			if !fortiAPIPatch(o["classifier"]) {
				return fmt.Errorf("Error reading classifier: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("classifier"); ok {
			if err = d.Set("classifier", flattenSwitchAclIngressClassifier(o["classifier"], d, "classifier", sv)); err != nil {
				if !fortiAPIPatch(o["classifier"]) {
					return fmt.Errorf("Error reading classifier: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("action", flattenSwitchAclIngressAction(o["action"], d, "action", sv)); err != nil {
			if !fortiAPIPatch(o["action"]) {
				return fmt.Errorf("Error reading action: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("action"); ok {
			if err = d.Set("action", flattenSwitchAclIngressAction(o["action"], d, "action", sv)); err != nil {
				if !fortiAPIPatch(o["action"]) {
					return fmt.Errorf("Error reading action: %v", err)
				}
			}
		}
	}

	if err = d.Set("group", flattenSwitchAclIngressGroup(o["group"], d, "group", sv)); err != nil {
		if !fortiAPIPatch(o["group"]) {
			return fmt.Errorf("Error reading group: %v", err)
		}
	}

	if err = d.Set("fswid", flattenSwitchAclIngressId(o["id"], d, "fswid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fswid: %v", err)
		}
	}

	return nil
}

func flattenSwitchAclIngressFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchAclIngressStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressIngressInterfaceAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressSchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["schedule-name"], _ = expandSwitchAclIngressScheduleScheduleName(d, i["schedule_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchAclIngressScheduleScheduleName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressIngressInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["member-name"], _ = expandSwitchAclIngressIngressInterfaceMemberName(d, i["member_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchAclIngressIngressInterfaceMemberName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressClassifier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dst_mac"
	if _, ok := d.GetOk(pre_append); ok {

		result["dst-mac"], _ = expandSwitchAclIngressClassifierDstMac(d, i["dst_mac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cos"
	if _, ok := d.GetOk(pre_append); ok {

		result["cos"], _ = expandSwitchAclIngressClassifierCos(d, i["cos"], pre_append, sv)
	}
	pre_append = pre + ".0." + "service"
	if _, ok := d.GetOk(pre_append); ok {

		result["service"], _ = expandSwitchAclIngressClassifierService(d, i["service"], pre_append, sv)
	}
	pre_append = pre + ".0." + "l3_interface"
	if _, ok := d.GetOk(pre_append); ok {

		result["l3-interface"], _ = expandSwitchAclIngressClassifierL3Interface(d, i["l3_interface"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dst_ip_prefix"
	if _, ok := d.GetOk(pre_append); ok {

		result["dst-ip-prefix"], _ = expandSwitchAclIngressClassifierDstIpPrefix(d, i["dst_ip_prefix"], pre_append, sv)
	}
	pre_append = pre + ".0." + "src_mac"
	if _, ok := d.GetOk(pre_append); ok {

		result["src-mac"], _ = expandSwitchAclIngressClassifierSrcMac(d, i["src_mac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok {

		result["dscp"], _ = expandSwitchAclIngressClassifierDscp(d, i["dscp"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dst_ip6_prefix"
	if _, ok := d.GetOk(pre_append); ok {

		result["dst-ip6-prefix"], _ = expandSwitchAclIngressClassifierDstIp6Prefix(d, i["dst_ip6_prefix"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ether_type"
	if _, ok := d.GetOk(pre_append); ok {

		result["ether-type"], _ = expandSwitchAclIngressClassifierEtherType(d, i["ether_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vlan_id"
	if _, ok := d.GetOk(pre_append); ok {

		result["vlan-id"], _ = expandSwitchAclIngressClassifierVlanId(d, i["vlan_id"], pre_append, sv)
	}
	pre_append = pre + ".0." + "src_ip_prefix"
	if _, ok := d.GetOk(pre_append); ok {

		result["src-ip-prefix"], _ = expandSwitchAclIngressClassifierSrcIpPrefix(d, i["src_ip_prefix"], pre_append, sv)
	}
	pre_append = pre + ".0." + "src_ip6_prefix"
	if _, ok := d.GetOk(pre_append); ok {

		result["src-ip6-prefix"], _ = expandSwitchAclIngressClassifierSrcIp6Prefix(d, i["src_ip6_prefix"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchAclIngressClassifierDstMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressClassifierCos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressClassifierService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressClassifierL3Interface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressClassifierDstIpPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressClassifierSrcMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressClassifierDscp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressClassifierDstIp6Prefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressClassifierEtherType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressClassifierVlanId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressClassifierSrcIpPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressClassifierSrcIp6Prefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "count"
	if _, ok := d.GetOk(pre_append); ok {

		result["count"], _ = expandSwitchAclIngressActionCount(d, i["count"], pre_append, sv)
	}
	pre_append = pre + ".0." + "redirect"
	if _, ok := d.GetOk(pre_append); ok {

		result["redirect"], _ = expandSwitchAclIngressActionRedirect(d, i["redirect"], pre_append, sv)
	}
	pre_append = pre + ".0." + "redirect_bcast_cpu"
	if _, ok := d.GetOk(pre_append); ok {

		result["redirect-bcast-cpu"], _ = expandSwitchAclIngressActionRedirectBcastCpu(d, i["redirect_bcast_cpu"], pre_append, sv)
	}
	pre_append = pre + ".0." + "redirect_physical_port"
	if _, ok := d.GetOk(pre_append); ok {

		result["redirect-physical-port"], _ = expandSwitchAclIngressActionRedirectPhysicalPort(d, i["redirect_physical_port"], pre_append, sv)
	} else {
		result["redirect-physical-port"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "cos_queue"
	if _, ok := d.GetOk(pre_append); ok {

		result["cos-queue"], _ = expandSwitchAclIngressActionCosQueue(d, i["cos_queue"], pre_append, sv)
	}
	pre_append = pre + ".0." + "remark_dscp"
	if _, ok := d.GetOk(pre_append); ok {

		result["remark-dscp"], _ = expandSwitchAclIngressActionRemarkDscp(d, i["remark_dscp"], pre_append, sv)
	}
	pre_append = pre + ".0." + "egress_mask"
	if _, ok := d.GetOk(pre_append); ok {

		result["egress-mask"], _ = expandSwitchAclIngressActionEgressMask(d, i["egress_mask"], pre_append, sv)
	} else {
		result["egress-mask"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "drop"
	if _, ok := d.GetOk(pre_append); ok {

		result["drop"], _ = expandSwitchAclIngressActionDrop(d, i["drop"], pre_append, sv)
	}
	pre_append = pre + ".0." + "redirect_bcast_no_cpu"
	if _, ok := d.GetOk(pre_append); ok {

		result["redirect-bcast-no-cpu"], _ = expandSwitchAclIngressActionRedirectBcastNoCpu(d, i["redirect_bcast_no_cpu"], pre_append, sv)
	}
	pre_append = pre + ".0." + "policer"
	if _, ok := d.GetOk(pre_append); ok {

		result["policer"], _ = expandSwitchAclIngressActionPolicer(d, i["policer"], pre_append, sv)
	}
	pre_append = pre + ".0." + "remark_cos"
	if _, ok := d.GetOk(pre_append); ok {

		result["remark-cos"], _ = expandSwitchAclIngressActionRemarkCos(d, i["remark_cos"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mirror"
	if _, ok := d.GetOk(pre_append); ok {

		result["mirror"], _ = expandSwitchAclIngressActionMirror(d, i["mirror"], pre_append, sv)
	}
	pre_append = pre + ".0." + "count_type"
	if _, ok := d.GetOk(pre_append); ok {

		result["count-type"], _ = expandSwitchAclIngressActionCountType(d, i["count_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cpu_cos_queue"
	if _, ok := d.GetOk(pre_append); ok {

		result["cpu-cos-queue"], _ = expandSwitchAclIngressActionCpuCosQueue(d, i["cpu_cos_queue"], pre_append, sv)
	}
	pre_append = pre + ".0." + "outer_vlan_tag"
	if _, ok := d.GetOk(pre_append); ok {

		result["outer-vlan-tag"], _ = expandSwitchAclIngressActionOuterVlanTag(d, i["outer_vlan_tag"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchAclIngressActionCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressActionRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressActionRedirectBcastCpu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressActionRedirectPhysicalPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["member-name"], _ = expandSwitchAclIngressActionRedirectPhysicalPortMemberName(d, i["member_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchAclIngressActionRedirectPhysicalPortMemberName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressActionCosQueue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressActionRemarkDscp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressActionEgressMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["member-name"], _ = expandSwitchAclIngressActionEgressMaskMemberName(d, i["member_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchAclIngressActionEgressMaskMemberName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressActionDrop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressActionRedirectBcastNoCpu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressActionPolicer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressActionRemarkCos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressActionMirror(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressActionCountType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressActionCpuCosQueue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressActionOuterVlanTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclIngressId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchAclIngress(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {

		t, err := expandSwitchAclIngressStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("ingress_interface_all"); ok {

		t, err := expandSwitchAclIngressIngressInterfaceAll(d, v, "ingress_interface_all", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-interface-all"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchAclIngressDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {

		t, err := expandSwitchAclIngressSchedule(d, v, "schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("ingress_interface"); ok || d.HasChange("ingress_interface") {

		t, err := expandSwitchAclIngressIngressInterface(d, v, "ingress_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-interface"] = t
		}
	}

	if v, ok := d.GetOk("classifier"); ok {

		t, err := expandSwitchAclIngressClassifier(d, v, "classifier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["classifier"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {

		t, err := expandSwitchAclIngressAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("group"); ok {

		t, err := expandSwitchAclIngressGroup(d, v, "group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group"] = t
		}
	}

	if v, ok := d.GetOk("fswid"); ok {

		t, err := expandSwitchAclIngressId(d, v, "fswid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
