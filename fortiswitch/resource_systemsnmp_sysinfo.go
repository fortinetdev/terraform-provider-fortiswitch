// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: SNMP system info configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSnmpSysinfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpSysinfoUpdate,
		Read:   resourceSystemSnmpSysinfoRead,
		Update: resourceSystemSnmpSysinfoUpdate,
		Delete: resourceSystemSnmpSysinfoDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_temp_alarm_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"contact_info": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"trap_high_cpu_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_log_full_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_low_memory_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"engine_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 24),
				Optional:     true,
				Computed:     true,
			},
			"location": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"trap_temp_warning_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_high_cpu_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSnmpSysinfoUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSnmpSysinfo(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpSysinfo resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSnmpSysinfo(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpSysinfo resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSnmpSysinfo")
	}

	return resourceSystemSnmpSysinfoRead(d, m)
}

func resourceSystemSnmpSysinfoDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSnmpSysinfo(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpSysinfo resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSnmpSysinfo(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemSnmpSysinfo resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpSysinfoRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemSnmpSysinfo(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpSysinfo resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpSysinfo(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpSysinfo resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnmpSysinfoStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapTempAlarmThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoContactInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapHighCpuThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapLogFullThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapLowMemoryThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoEngineId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapTempWarningThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapHighCpuInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSnmpSysinfo(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemSnmpSysinfoStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trap_temp_alarm_threshold", flattenSystemSnmpSysinfoTrapTempAlarmThreshold(o["trap-temp-alarm-threshold"], d, "trap_temp_alarm_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["trap-temp-alarm-threshold"]) {
			return fmt.Errorf("Error reading trap_temp_alarm_threshold: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemSnmpSysinfoDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("contact_info", flattenSystemSnmpSysinfoContactInfo(o["contact-info"], d, "contact_info", sv)); err != nil {
		if !fortiAPIPatch(o["contact-info"]) {
			return fmt.Errorf("Error reading contact_info: %v", err)
		}
	}

	if err = d.Set("trap_high_cpu_threshold", flattenSystemSnmpSysinfoTrapHighCpuThreshold(o["trap-high-cpu-threshold"], d, "trap_high_cpu_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["trap-high-cpu-threshold"]) {
			return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
		}
	}

	if err = d.Set("trap_log_full_threshold", flattenSystemSnmpSysinfoTrapLogFullThreshold(o["trap-log-full-threshold"], d, "trap_log_full_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["trap-log-full-threshold"]) {
			return fmt.Errorf("Error reading trap_log_full_threshold: %v", err)
		}
	}

	if err = d.Set("trap_low_memory_threshold", flattenSystemSnmpSysinfoTrapLowMemoryThreshold(o["trap-low-memory-threshold"], d, "trap_low_memory_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["trap-low-memory-threshold"]) {
			return fmt.Errorf("Error reading trap_low_memory_threshold: %v", err)
		}
	}

	if err = d.Set("engine_id", flattenSystemSnmpSysinfoEngineId(o["engine-id"], d, "engine_id", sv)); err != nil {
		if !fortiAPIPatch(o["engine-id"]) {
			return fmt.Errorf("Error reading engine_id: %v", err)
		}
	}

	if err = d.Set("location", flattenSystemSnmpSysinfoLocation(o["location"], d, "location", sv)); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("trap_temp_warning_threshold", flattenSystemSnmpSysinfoTrapTempWarningThreshold(o["trap-temp-warning-threshold"], d, "trap_temp_warning_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["trap-temp-warning-threshold"]) {
			return fmt.Errorf("Error reading trap_temp_warning_threshold: %v", err)
		}
	}

	if err = d.Set("trap_high_cpu_interval", flattenSystemSnmpSysinfoTrapHighCpuInterval(o["trap-high-cpu-interval"], d, "trap_high_cpu_interval", sv)); err != nil {
		if !fortiAPIPatch(o["trap-high-cpu-interval"]) {
			return fmt.Errorf("Error reading trap_high_cpu_interval: %v", err)
		}
	}

	return nil
}

func flattenSystemSnmpSysinfoFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemSnmpSysinfoStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapTempAlarmThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoContactInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapHighCpuThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapLogFullThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapLowMemoryThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoEngineId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapTempWarningThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapHighCpuInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpSysinfo(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandSystemSnmpSysinfoStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("trap_temp_alarm_threshold"); ok {
		if setArgNil {
			obj["trap-temp-alarm-threshold"] = nil
		} else {

			t, err := expandSystemSnmpSysinfoTrapTempAlarmThreshold(d, v, "trap_temp_alarm_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trap-temp-alarm-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("description"); ok {
		if setArgNil {
			obj["description"] = nil
		} else {

			t, err := expandSystemSnmpSysinfoDescription(d, v, "description", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["description"] = t
			}
		}
	}

	if v, ok := d.GetOk("contact_info"); ok {
		if setArgNil {
			obj["contact-info"] = nil
		} else {

			t, err := expandSystemSnmpSysinfoContactInfo(d, v, "contact_info", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["contact-info"] = t
			}
		}
	}

	if v, ok := d.GetOk("trap_high_cpu_threshold"); ok {
		if setArgNil {
			obj["trap-high-cpu-threshold"] = nil
		} else {

			t, err := expandSystemSnmpSysinfoTrapHighCpuThreshold(d, v, "trap_high_cpu_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trap-high-cpu-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("trap_log_full_threshold"); ok {
		if setArgNil {
			obj["trap-log-full-threshold"] = nil
		} else {

			t, err := expandSystemSnmpSysinfoTrapLogFullThreshold(d, v, "trap_log_full_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trap-log-full-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("trap_low_memory_threshold"); ok {
		if setArgNil {
			obj["trap-low-memory-threshold"] = nil
		} else {

			t, err := expandSystemSnmpSysinfoTrapLowMemoryThreshold(d, v, "trap_low_memory_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trap-low-memory-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("engine_id"); ok {
		if setArgNil {
			obj["engine-id"] = nil
		} else {

			t, err := expandSystemSnmpSysinfoEngineId(d, v, "engine_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["engine-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("location"); ok {
		if setArgNil {
			obj["location"] = nil
		} else {

			t, err := expandSystemSnmpSysinfoLocation(d, v, "location", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["location"] = t
			}
		}
	}

	if v, ok := d.GetOk("trap_temp_warning_threshold"); ok {
		if setArgNil {
			obj["trap-temp-warning-threshold"] = nil
		} else {

			t, err := expandSystemSnmpSysinfoTrapTempWarningThreshold(d, v, "trap_temp_warning_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trap-temp-warning-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("trap_high_cpu_interval"); ok {
		if setArgNil {
			obj["trap-high-cpu-interval"] = nil
		} else {

			t, err := expandSystemSnmpSysinfoTrapHighCpuInterval(d, v, "trap_high_cpu_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trap-high-cpu-interval"] = t
			}
		}
	}

	return &obj, nil
}
