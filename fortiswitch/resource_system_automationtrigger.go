// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Trigger for automation stitches.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutomationTrigger() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationTriggerCreate,
		Read:   resourceSystemAutomationTriggerRead,
		Update: resourceSystemAutomationTriggerUpdate,
		Delete: resourceSystemAutomationTriggerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"trigger_minute": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 59),
				Optional:     true,
				Computed:     true,
			},
			"fields": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"logid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trigger_hour": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 23),
				Optional:     true,
				Computed:     true,
			},
			"trigger_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trigger_frequency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trigger_weekday": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trigger_day": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 31),
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

func resourceSystemAutomationTriggerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAutomationTrigger(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationTrigger resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAutomationTrigger(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationTrigger resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutomationTrigger")
	}

	return resourceSystemAutomationTriggerRead(d, m)
}

func resourceSystemAutomationTriggerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAutomationTrigger(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationTrigger resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutomationTrigger(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationTrigger resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutomationTrigger")
	}

	return resourceSystemAutomationTriggerRead(d, m)
}

func resourceSystemAutomationTriggerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAutomationTrigger(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationTrigger resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationTriggerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAutomationTrigger(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationTrigger resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationTrigger(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationTrigger resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationTriggerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerMinute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFields(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSystemAutomationTriggerFieldsId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {

			tmp["value"] = flattenSystemAutomationTriggerFieldsValue(i["value"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemAutomationTriggerFieldsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemAutomationTriggerFieldsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFieldsValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFieldsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerLogid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerHour(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerEventType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerFrequency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerWeekday(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAutomationTrigger(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemAutomationTriggerName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("trigger_minute", flattenSystemAutomationTriggerTriggerMinute(o["trigger-minute"], d, "trigger_minute", sv)); err != nil {
		if !fortiAPIPatch(o["trigger-minute"]) {
			return fmt.Errorf("Error reading trigger_minute: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fields", flattenSystemAutomationTriggerFields(o["fields"], d, "fields", sv)); err != nil {
			if !fortiAPIPatch(o["fields"]) {
				return fmt.Errorf("Error reading fields: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fields"); ok {
			if err = d.Set("fields", flattenSystemAutomationTriggerFields(o["fields"], d, "fields", sv)); err != nil {
				if !fortiAPIPatch(o["fields"]) {
					return fmt.Errorf("Error reading fields: %v", err)
				}
			}
		}
	}

	if err = d.Set("logid", flattenSystemAutomationTriggerLogid(o["logid"], d, "logid", sv)); err != nil {
		if !fortiAPIPatch(o["logid"]) {
			return fmt.Errorf("Error reading logid: %v", err)
		}
	}

	if err = d.Set("trigger_hour", flattenSystemAutomationTriggerTriggerHour(o["trigger-hour"], d, "trigger_hour", sv)); err != nil {
		if !fortiAPIPatch(o["trigger-hour"]) {
			return fmt.Errorf("Error reading trigger_hour: %v", err)
		}
	}

	if err = d.Set("trigger_type", flattenSystemAutomationTriggerTriggerType(o["trigger-type"], d, "trigger_type", sv)); err != nil {
		if !fortiAPIPatch(o["trigger-type"]) {
			return fmt.Errorf("Error reading trigger_type: %v", err)
		}
	}

	if err = d.Set("event_type", flattenSystemAutomationTriggerEventType(o["event-type"], d, "event_type", sv)); err != nil {
		if !fortiAPIPatch(o["event-type"]) {
			return fmt.Errorf("Error reading event_type: %v", err)
		}
	}

	if err = d.Set("trigger_frequency", flattenSystemAutomationTriggerTriggerFrequency(o["trigger-frequency"], d, "trigger_frequency", sv)); err != nil {
		if !fortiAPIPatch(o["trigger-frequency"]) {
			return fmt.Errorf("Error reading trigger_frequency: %v", err)
		}
	}

	if err = d.Set("trigger_weekday", flattenSystemAutomationTriggerTriggerWeekday(o["trigger-weekday"], d, "trigger_weekday", sv)); err != nil {
		if !fortiAPIPatch(o["trigger-weekday"]) {
			return fmt.Errorf("Error reading trigger_weekday: %v", err)
		}
	}

	if err = d.Set("trigger_day", flattenSystemAutomationTriggerTriggerDay(o["trigger-day"], d, "trigger_day", sv)); err != nil {
		if !fortiAPIPatch(o["trigger-day"]) {
			return fmt.Errorf("Error reading trigger_day: %v", err)
		}
	}

	return nil
}

func flattenSystemAutomationTriggerFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemAutomationTriggerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerMinute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFields(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSystemAutomationTriggerFieldsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["value"], _ = expandSystemAutomationTriggerFieldsValue(d, i["value"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemAutomationTriggerFieldsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationTriggerFieldsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFieldsValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFieldsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerLogid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerHour(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerEventType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerFrequency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerWeekday(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutomationTrigger(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemAutomationTriggerName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("trigger_minute"); ok {

		t, err := expandSystemAutomationTriggerTriggerMinute(d, v, "trigger_minute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-minute"] = t
		}
	}

	if v, ok := d.GetOk("fields"); ok || d.HasChange("fields") {

		t, err := expandSystemAutomationTriggerFields(d, v, "fields", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fields"] = t
		}
	}

	if v, ok := d.GetOk("logid"); ok {

		t, err := expandSystemAutomationTriggerLogid(d, v, "logid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logid"] = t
		}
	}

	if v, ok := d.GetOkExists("trigger_hour"); ok {

		t, err := expandSystemAutomationTriggerTriggerHour(d, v, "trigger_hour", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-hour"] = t
		}
	}

	if v, ok := d.GetOk("trigger_type"); ok {

		t, err := expandSystemAutomationTriggerTriggerType(d, v, "trigger_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-type"] = t
		}
	}

	if v, ok := d.GetOk("event_type"); ok {

		t, err := expandSystemAutomationTriggerEventType(d, v, "event_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event-type"] = t
		}
	}

	if v, ok := d.GetOk("trigger_frequency"); ok {

		t, err := expandSystemAutomationTriggerTriggerFrequency(d, v, "trigger_frequency", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-frequency"] = t
		}
	}

	if v, ok := d.GetOk("trigger_weekday"); ok {

		t, err := expandSystemAutomationTriggerTriggerWeekday(d, v, "trigger_weekday", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-weekday"] = t
		}
	}

	if v, ok := d.GetOk("trigger_day"); ok {

		t, err := expandSystemAutomationTriggerTriggerDay(d, v, "trigger_day", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-day"] = t
		}
	}

	return &obj, nil
}
