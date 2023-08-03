// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: QOS egress policy.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchQosQosPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchQosQosPolicyCreate,
		Read:   resourceSwitchQosQosPolicyRead,
		Update: resourceSwitchQosQosPolicyUpdate,
		Delete: resourceSwitchQosQosPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"rate_by": &schema.Schema{
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
			"cos_queue": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_rate_percent": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"min_rate_percent": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ecn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"min_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"wred_slope": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"drop_policy": &schema.Schema{
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
					},
				},
			},
			"schedule": &schema.Schema{
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

func resourceSwitchQosQosPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchQosQosPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchQosQosPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchQosQosPolicy(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchQosQosPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchQosQosPolicy")
	}

	return resourceSwitchQosQosPolicyRead(d, m)
}

func resourceSwitchQosQosPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchQosQosPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchQosQosPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchQosQosPolicy(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchQosQosPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchQosQosPolicy")
	}

	return resourceSwitchQosQosPolicyRead(d, m)
}

func resourceSwitchQosQosPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchQosQosPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchQosQosPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchQosQosPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchQosQosPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchQosQosPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchQosQosPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchQosQosPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchQosQosPolicyRateBy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosQosPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosQosPolicyCosQueue(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate_percent"
		if _, ok := i["max-rate-percent"]; ok {

			tmp["max_rate_percent"] = flattenSwitchQosQosPolicyCosQueueMaxRatePercent(i["max-rate-percent"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchQosQosPolicyCosQueueName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {

			tmp["weight"] = flattenSwitchQosQosPolicyCosQueueWeight(i["weight"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate_percent"
		if _, ok := i["min-rate-percent"]; ok {

			tmp["min_rate_percent"] = flattenSwitchQosQosPolicyCosQueueMinRatePercent(i["min-rate-percent"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ecn"
		if _, ok := i["ecn"]; ok {

			tmp["ecn"] = flattenSwitchQosQosPolicyCosQueueEcn(i["ecn"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate"
		if _, ok := i["min-rate"]; ok {

			tmp["min_rate"] = flattenSwitchQosQosPolicyCosQueueMinRate(i["min-rate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wred_slope"
		if _, ok := i["wred-slope"]; ok {

			tmp["wred_slope"] = flattenSwitchQosQosPolicyCosQueueWredSlope(i["wred-slope"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate"
		if _, ok := i["max-rate"]; ok {

			tmp["max_rate"] = flattenSwitchQosQosPolicyCosQueueMaxRate(i["max-rate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop_policy"
		if _, ok := i["drop-policy"]; ok {

			tmp["drop_policy"] = flattenSwitchQosQosPolicyCosQueueDropPolicy(i["drop-policy"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchQosQosPolicyCosQueueDescription(i["description"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchQosQosPolicyCosQueueMaxRatePercent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosQosPolicyCosQueueName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosQosPolicyCosQueueWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosQosPolicyCosQueueMinRatePercent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosQosPolicyCosQueueEcn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosQosPolicyCosQueueMinRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosQosPolicyCosQueueWredSlope(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosQosPolicyCosQueueMaxRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosQosPolicyCosQueueDropPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosQosPolicyCosQueueDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosQosPolicySchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchQosQosPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("rate_by", flattenSwitchQosQosPolicyRateBy(o["rate-by"], d, "rate_by", sv)); err != nil {
		if !fortiAPIPatch(o["rate-by"]) {
			return fmt.Errorf("Error reading rate_by: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchQosQosPolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("cos_queue", flattenSwitchQosQosPolicyCosQueue(o["cos-queue"], d, "cos_queue", sv)); err != nil {
			if !fortiAPIPatch(o["cos-queue"]) {
				return fmt.Errorf("Error reading cos_queue: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cos_queue"); ok {
			if err = d.Set("cos_queue", flattenSwitchQosQosPolicyCosQueue(o["cos-queue"], d, "cos_queue", sv)); err != nil {
				if !fortiAPIPatch(o["cos-queue"]) {
					return fmt.Errorf("Error reading cos_queue: %v", err)
				}
			}
		}
	}

	if err = d.Set("schedule", flattenSwitchQosQosPolicySchedule(o["schedule"], d, "schedule", sv)); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	return nil
}

func flattenSwitchQosQosPolicyFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchQosQosPolicyRateBy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosQosPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosQosPolicyCosQueue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate_percent"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["max-rate-percent"], _ = expandSwitchQosQosPolicyCosQueueMaxRatePercent(d, i["max_rate_percent"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSwitchQosQosPolicyCosQueueName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["weight"], _ = expandSwitchQosQosPolicyCosQueueWeight(d, i["weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate_percent"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["min-rate-percent"], _ = expandSwitchQosQosPolicyCosQueueMinRatePercent(d, i["min_rate_percent"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ecn"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ecn"], _ = expandSwitchQosQosPolicyCosQueueEcn(d, i["ecn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["min-rate"], _ = expandSwitchQosQosPolicyCosQueueMinRate(d, i["min_rate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wred_slope"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["wred-slope"], _ = expandSwitchQosQosPolicyCosQueueWredSlope(d, i["wred_slope"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["max-rate"], _ = expandSwitchQosQosPolicyCosQueueMaxRate(d, i["max_rate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop_policy"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["drop-policy"], _ = expandSwitchQosQosPolicyCosQueueDropPolicy(d, i["drop_policy"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandSwitchQosQosPolicyCosQueueDescription(d, i["description"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchQosQosPolicyCosQueueMaxRatePercent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosQosPolicyCosQueueName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosQosPolicyCosQueueWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosQosPolicyCosQueueMinRatePercent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosQosPolicyCosQueueEcn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosQosPolicyCosQueueMinRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosQosPolicyCosQueueWredSlope(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosQosPolicyCosQueueMaxRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosQosPolicyCosQueueDropPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosQosPolicyCosQueueDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosQosPolicySchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchQosQosPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("rate_by"); ok {

		t, err := expandSwitchQosQosPolicyRateBy(d, v, "rate_by", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-by"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchQosQosPolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("cos_queue"); ok || d.HasChange("cos_queue") {

		t, err := expandSwitchQosQosPolicyCosQueue(d, v, "cos_queue", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos-queue"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {

		t, err := expandSwitchQosQosPolicySchedule(d, v, "schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	return &obj, nil
}
