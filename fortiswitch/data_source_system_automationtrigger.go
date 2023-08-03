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

func dataSourceSystemAutomationTrigger() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAutomationTriggerRead,
		Schema: map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"trigger_minute": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"fields": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"logid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trigger_hour": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"trigger_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"event_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trigger_frequency": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trigger_weekday": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trigger_day": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemAutomationTriggerRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemAutomationTrigger: type error")
	}

	o, err := c.ReadSystemAutomationTrigger(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemAutomationTrigger: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemAutomationTrigger(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemAutomationTrigger from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemAutomationTriggerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationTriggerTriggerMinute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationTriggerFields(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
			tmp["id"] = dataSourceFlattenSystemAutomationTriggerFieldsId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = dataSourceFlattenSystemAutomationTriggerFieldsValue(i["value"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenSystemAutomationTriggerFieldsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAutomationTriggerFieldsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationTriggerFieldsValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationTriggerFieldsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationTriggerLogid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationTriggerTriggerHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationTriggerTriggerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationTriggerEventType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationTriggerTriggerFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationTriggerTriggerWeekday(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationTriggerTriggerDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAutomationTrigger(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemAutomationTriggerName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("trigger_minute", dataSourceFlattenSystemAutomationTriggerTriggerMinute(o["trigger-minute"], d, "trigger_minute")); err != nil {
		if !fortiAPIPatch(o["trigger-minute"]) {
			return fmt.Errorf("Error reading trigger_minute: %v", err)
		}
	}

	if err = d.Set("fields", dataSourceFlattenSystemAutomationTriggerFields(o["fields"], d, "fields")); err != nil {
		if !fortiAPIPatch(o["fields"]) {
			return fmt.Errorf("Error reading fields: %v", err)
		}
	}

	if err = d.Set("logid", dataSourceFlattenSystemAutomationTriggerLogid(o["logid"], d, "logid")); err != nil {
		if !fortiAPIPatch(o["logid"]) {
			return fmt.Errorf("Error reading logid: %v", err)
		}
	}

	if err = d.Set("trigger_hour", dataSourceFlattenSystemAutomationTriggerTriggerHour(o["trigger-hour"], d, "trigger_hour")); err != nil {
		if !fortiAPIPatch(o["trigger-hour"]) {
			return fmt.Errorf("Error reading trigger_hour: %v", err)
		}
	}

	if err = d.Set("trigger_type", dataSourceFlattenSystemAutomationTriggerTriggerType(o["trigger-type"], d, "trigger_type")); err != nil {
		if !fortiAPIPatch(o["trigger-type"]) {
			return fmt.Errorf("Error reading trigger_type: %v", err)
		}
	}

	if err = d.Set("event_type", dataSourceFlattenSystemAutomationTriggerEventType(o["event-type"], d, "event_type")); err != nil {
		if !fortiAPIPatch(o["event-type"]) {
			return fmt.Errorf("Error reading event_type: %v", err)
		}
	}

	if err = d.Set("trigger_frequency", dataSourceFlattenSystemAutomationTriggerTriggerFrequency(o["trigger-frequency"], d, "trigger_frequency")); err != nil {
		if !fortiAPIPatch(o["trigger-frequency"]) {
			return fmt.Errorf("Error reading trigger_frequency: %v", err)
		}
	}

	if err = d.Set("trigger_weekday", dataSourceFlattenSystemAutomationTriggerTriggerWeekday(o["trigger-weekday"], d, "trigger_weekday")); err != nil {
		if !fortiAPIPatch(o["trigger-weekday"]) {
			return fmt.Errorf("Error reading trigger_weekday: %v", err)
		}
	}

	if err = d.Set("trigger_day", dataSourceFlattenSystemAutomationTriggerTriggerDay(o["trigger-day"], d, "trigger_day")); err != nil {
		if !fortiAPIPatch(o["trigger-day"]) {
			return fmt.Errorf("Error reading trigger_day: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemAutomationTriggerFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
