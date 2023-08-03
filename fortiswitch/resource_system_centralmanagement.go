// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Central management configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCentralManagement() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCentralManagementUpdate,
		Read:   resourceSystemCentralManagementRead,
		Update: resourceSystemCentralManagementUpdate,
		Delete: resourceSystemCentralManagementDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmg_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule_config_restore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_pushd_firmware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_push_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_remote_firmware_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule_script_restore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmg": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"allow_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemCentralManagementUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCentralManagement(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCentralManagement resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemCentralManagement(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemCentralManagement resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCentralManagement")
	}

	return resourceSystemCentralManagementRead(d, m)
}

func resourceSystemCentralManagementDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCentralManagement(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemCentralManagement resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCentralManagement(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemCentralManagement resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCentralManagementRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemCentralManagement(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemCentralManagement resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCentralManagement(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCentralManagement resource from API: %v", err)
	}
	return nil
}

func flattenSystemCentralManagementStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementEncAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementFmgSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementScheduleConfigRestore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowPushdFirmware(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowPushConfiguration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowRemoteFirmwareUpgrade(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementSerialNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementScheduleScriptRestore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementFmg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCentralManagement(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemCentralManagementStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenSystemCentralManagementEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("fmg_source_ip", flattenSystemCentralManagementFmgSourceIp(o["fmg-source-ip"], d, "fmg_source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["fmg-source-ip"]) {
			return fmt.Errorf("Error reading fmg_source_ip: %v", err)
		}
	}

	if err = d.Set("schedule_config_restore", flattenSystemCentralManagementScheduleConfigRestore(o["schedule-config-restore"], d, "schedule_config_restore", sv)); err != nil {
		if !fortiAPIPatch(o["schedule-config-restore"]) {
			return fmt.Errorf("Error reading schedule_config_restore: %v", err)
		}
	}

	if err = d.Set("allow_pushd_firmware", flattenSystemCentralManagementAllowPushdFirmware(o["allow-pushd-firmware"], d, "allow_pushd_firmware", sv)); err != nil {
		if !fortiAPIPatch(o["allow-pushd-firmware"]) {
			return fmt.Errorf("Error reading allow_pushd_firmware: %v", err)
		}
	}

	if err = d.Set("allow_push_configuration", flattenSystemCentralManagementAllowPushConfiguration(o["allow-push-configuration"], d, "allow_push_configuration", sv)); err != nil {
		if !fortiAPIPatch(o["allow-push-configuration"]) {
			return fmt.Errorf("Error reading allow_push_configuration: %v", err)
		}
	}

	if err = d.Set("allow_remote_firmware_upgrade", flattenSystemCentralManagementAllowRemoteFirmwareUpgrade(o["allow-remote-firmware-upgrade"], d, "allow_remote_firmware_upgrade", sv)); err != nil {
		if !fortiAPIPatch(o["allow-remote-firmware-upgrade"]) {
			return fmt.Errorf("Error reading allow_remote_firmware_upgrade: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenSystemCentralManagementSerialNumber(o["serial-number"], d, "serial_number", sv)); err != nil {
		if !fortiAPIPatch(o["serial-number"]) {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemCentralManagementMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("schedule_script_restore", flattenSystemCentralManagementScheduleScriptRestore(o["schedule-script-restore"], d, "schedule_script_restore", sv)); err != nil {
		if !fortiAPIPatch(o["schedule-script-restore"]) {
			return fmt.Errorf("Error reading schedule_script_restore: %v", err)
		}
	}

	if err = d.Set("fmg", flattenSystemCentralManagementFmg(o["fmg"], d, "fmg", sv)); err != nil {
		if !fortiAPIPatch(o["fmg"]) {
			return fmt.Errorf("Error reading fmg: %v", err)
		}
	}

	if err = d.Set("allow_monitor", flattenSystemCentralManagementAllowMonitor(o["allow-monitor"], d, "allow_monitor", sv)); err != nil {
		if !fortiAPIPatch(o["allow-monitor"]) {
			return fmt.Errorf("Error reading allow_monitor: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemCentralManagementType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenSystemCentralManagementFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemCentralManagementStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementEncAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementFmgSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementScheduleConfigRestore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowPushdFirmware(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowPushConfiguration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowRemoteFirmwareUpgrade(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementSerialNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementScheduleScriptRestore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementFmg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCentralManagement(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandSystemCentralManagementStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {
		if setArgNil {
			obj["enc-algorithm"] = nil
		} else {

			t, err := expandSystemCentralManagementEncAlgorithm(d, v, "enc_algorithm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["enc-algorithm"] = t
			}
		}
	}

	if v, ok := d.GetOk("fmg_source_ip"); ok {
		if setArgNil {
			obj["fmg-source-ip"] = nil
		} else {

			t, err := expandSystemCentralManagementFmgSourceIp(d, v, "fmg_source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fmg-source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("schedule_config_restore"); ok {
		if setArgNil {
			obj["schedule-config-restore"] = nil
		} else {

			t, err := expandSystemCentralManagementScheduleConfigRestore(d, v, "schedule_config_restore", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["schedule-config-restore"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_pushd_firmware"); ok {
		if setArgNil {
			obj["allow-pushd-firmware"] = nil
		} else {

			t, err := expandSystemCentralManagementAllowPushdFirmware(d, v, "allow_pushd_firmware", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-pushd-firmware"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_push_configuration"); ok {
		if setArgNil {
			obj["allow-push-configuration"] = nil
		} else {

			t, err := expandSystemCentralManagementAllowPushConfiguration(d, v, "allow_push_configuration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-push-configuration"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_remote_firmware_upgrade"); ok {
		if setArgNil {
			obj["allow-remote-firmware-upgrade"] = nil
		} else {

			t, err := expandSystemCentralManagementAllowRemoteFirmwareUpgrade(d, v, "allow_remote_firmware_upgrade", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-remote-firmware-upgrade"] = t
			}
		}
	}

	if v, ok := d.GetOk("serial_number"); ok {
		if setArgNil {
			obj["serial-number"] = nil
		} else {

			t, err := expandSystemCentralManagementSerialNumber(d, v, "serial_number", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["serial-number"] = t
			}
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		if setArgNil {
			obj["mode"] = nil
		} else {

			t, err := expandSystemCentralManagementMode(d, v, "mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("schedule_script_restore"); ok {
		if setArgNil {
			obj["schedule-script-restore"] = nil
		} else {

			t, err := expandSystemCentralManagementScheduleScriptRestore(d, v, "schedule_script_restore", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["schedule-script-restore"] = t
			}
		}
	}

	if v, ok := d.GetOk("fmg"); ok {
		if setArgNil {
			obj["fmg"] = nil
		} else {

			t, err := expandSystemCentralManagementFmg(d, v, "fmg", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fmg"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_monitor"); ok {
		if setArgNil {
			obj["allow-monitor"] = nil
		} else {

			t, err := expandSystemCentralManagementAllowMonitor(d, v, "allow_monitor", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-monitor"] = t
			}
		}
	}

	if v, ok := d.GetOk("type"); ok {
		if setArgNil {
			obj["type"] = nil
		} else {

			t, err := expandSystemCentralManagementType(d, v, "type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["type"] = t
			}
		}
	}

	return &obj, nil
}
