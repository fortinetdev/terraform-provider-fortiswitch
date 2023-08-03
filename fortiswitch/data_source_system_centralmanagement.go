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

func dataSourceSystemCentralManagement() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemCentralManagementRead,
		Schema: map[string]*schema.Schema{

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fmg_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule_config_restore": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_pushd_firmware": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_push_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_remote_firmware_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule_script_restore": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fmg": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemCentralManagementRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemCentralManagement"

	o, err := c.ReadSystemCentralManagement(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemCentralManagement: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemCentralManagement(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemCentralManagement from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemCentralManagementStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementFmgSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementScheduleConfigRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementAllowPushdFirmware(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementAllowPushConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementAllowRemoteFirmwareUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementScheduleScriptRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementFmg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementAllowMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCentralManagementType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemCentralManagement(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenSystemCentralManagementStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", dataSourceFlattenSystemCentralManagementEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("fmg_source_ip", dataSourceFlattenSystemCentralManagementFmgSourceIp(o["fmg-source-ip"], d, "fmg_source_ip")); err != nil {
		if !fortiAPIPatch(o["fmg-source-ip"]) {
			return fmt.Errorf("Error reading fmg_source_ip: %v", err)
		}
	}

	if err = d.Set("schedule_config_restore", dataSourceFlattenSystemCentralManagementScheduleConfigRestore(o["schedule-config-restore"], d, "schedule_config_restore")); err != nil {
		if !fortiAPIPatch(o["schedule-config-restore"]) {
			return fmt.Errorf("Error reading schedule_config_restore: %v", err)
		}
	}

	if err = d.Set("allow_pushd_firmware", dataSourceFlattenSystemCentralManagementAllowPushdFirmware(o["allow-pushd-firmware"], d, "allow_pushd_firmware")); err != nil {
		if !fortiAPIPatch(o["allow-pushd-firmware"]) {
			return fmt.Errorf("Error reading allow_pushd_firmware: %v", err)
		}
	}

	if err = d.Set("allow_push_configuration", dataSourceFlattenSystemCentralManagementAllowPushConfiguration(o["allow-push-configuration"], d, "allow_push_configuration")); err != nil {
		if !fortiAPIPatch(o["allow-push-configuration"]) {
			return fmt.Errorf("Error reading allow_push_configuration: %v", err)
		}
	}

	if err = d.Set("allow_remote_firmware_upgrade", dataSourceFlattenSystemCentralManagementAllowRemoteFirmwareUpgrade(o["allow-remote-firmware-upgrade"], d, "allow_remote_firmware_upgrade")); err != nil {
		if !fortiAPIPatch(o["allow-remote-firmware-upgrade"]) {
			return fmt.Errorf("Error reading allow_remote_firmware_upgrade: %v", err)
		}
	}

	if err = d.Set("serial_number", dataSourceFlattenSystemCentralManagementSerialNumber(o["serial-number"], d, "serial_number")); err != nil {
		if !fortiAPIPatch(o["serial-number"]) {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("mode", dataSourceFlattenSystemCentralManagementMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("schedule_script_restore", dataSourceFlattenSystemCentralManagementScheduleScriptRestore(o["schedule-script-restore"], d, "schedule_script_restore")); err != nil {
		if !fortiAPIPatch(o["schedule-script-restore"]) {
			return fmt.Errorf("Error reading schedule_script_restore: %v", err)
		}
	}

	if err = d.Set("fmg", dataSourceFlattenSystemCentralManagementFmg(o["fmg"], d, "fmg")); err != nil {
		if !fortiAPIPatch(o["fmg"]) {
			return fmt.Errorf("Error reading fmg: %v", err)
		}
	}

	if err = d.Set("allow_monitor", dataSourceFlattenSystemCentralManagementAllowMonitor(o["allow-monitor"], d, "allow_monitor")); err != nil {
		if !fortiAPIPatch(o["allow-monitor"]) {
			return fmt.Errorf("Error reading allow_monitor: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemCentralManagementType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemCentralManagementFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
