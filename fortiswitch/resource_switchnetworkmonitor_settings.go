// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Global configuration of network monitoring on the switch.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchNetworkMonitorSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchNetworkMonitorSettingsUpdate,
		Read:   resourceSwitchNetworkMonitorSettingsRead,
		Update: resourceSwitchNetworkMonitorSettingsUpdate,
		Delete: resourceSwitchNetworkMonitorSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"db_aging_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"survey_mode_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"survey_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchNetworkMonitorSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchNetworkMonitorSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchNetworkMonitorSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchNetworkMonitorSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchNetworkMonitorSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchNetworkMonitorSettings")
	}

	return resourceSwitchNetworkMonitorSettingsRead(d, m)
}

func resourceSwitchNetworkMonitorSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchNetworkMonitorSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchNetworkMonitorSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchNetworkMonitorSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchNetworkMonitorSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchNetworkMonitorSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchNetworkMonitorSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchNetworkMonitorSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchNetworkMonitorSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchNetworkMonitorSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchNetworkMonitorSettingsDbAgingInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchNetworkMonitorSettingsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchNetworkMonitorSettingsSurveyModeInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchNetworkMonitorSettingsSurveyMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchNetworkMonitorSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("db_aging_interval", flattenSwitchNetworkMonitorSettingsDbAgingInterval(o["db-aging-interval"], d, "db_aging_interval", sv)); err != nil {
		if !fortiAPIPatch(o["db-aging-interval"]) {
			return fmt.Errorf("Error reading db_aging_interval: %v", err)
		}
	}

	if err = d.Set("status", flattenSwitchNetworkMonitorSettingsStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("survey_mode_interval", flattenSwitchNetworkMonitorSettingsSurveyModeInterval(o["survey-mode-interval"], d, "survey_mode_interval", sv)); err != nil {
		if !fortiAPIPatch(o["survey-mode-interval"]) {
			return fmt.Errorf("Error reading survey_mode_interval: %v", err)
		}
	}

	if err = d.Set("survey_mode", flattenSwitchNetworkMonitorSettingsSurveyMode(o["survey-mode"], d, "survey_mode", sv)); err != nil {
		if !fortiAPIPatch(o["survey-mode"]) {
			return fmt.Errorf("Error reading survey_mode: %v", err)
		}
	}

	return nil
}

func flattenSwitchNetworkMonitorSettingsFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchNetworkMonitorSettingsDbAgingInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchNetworkMonitorSettingsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchNetworkMonitorSettingsSurveyModeInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchNetworkMonitorSettingsSurveyMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchNetworkMonitorSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("db_aging_interval"); ok {
		if setArgNil {
			obj["db-aging-interval"] = nil
		} else {

			t, err := expandSwitchNetworkMonitorSettingsDbAgingInterval(d, v, "db_aging_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["db-aging-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandSwitchNetworkMonitorSettingsStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("survey_mode_interval"); ok {
		if setArgNil {
			obj["survey-mode-interval"] = nil
		} else {

			t, err := expandSwitchNetworkMonitorSettingsSurveyModeInterval(d, v, "survey_mode_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["survey-mode-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("survey_mode"); ok {
		if setArgNil {
			obj["survey-mode"] = nil
		} else {

			t, err := expandSwitchNetworkMonitorSettingsSurveyMode(d, v, "survey_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["survey-mode"] = t
			}
		}
	}

	return &obj, nil
}
