// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: FortiManagerconfiguration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemFortimanager() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemFortimanagerRead,
		Schema: map[string]*schema.Schema{

			"central_mgmt_auto_backup": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"central_mgmt_schedule_config_restore": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"central_mgmt_schedule_script_restore": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipsec": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemFortimanagerRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemFortimanager"

	o, err := c.ReadSystemFortimanager(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemFortimanager: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemFortimanager(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemFortimanager from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemFortimanagerCentralMgmtAutoBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortimanagerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortimanagerCentralMgmtScheduleConfigRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortimanagerCentralMgmtScheduleScriptRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortimanagerIpsec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemFortimanager(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("central_mgmt_auto_backup", dataSourceFlattenSystemFortimanagerCentralMgmtAutoBackup(o["central-mgmt-auto-backup"], d, "central_mgmt_auto_backup")); err != nil {
		if !fortiAPIPatch(o["central-mgmt-auto-backup"]) {
			return fmt.Errorf("Error reading central_mgmt_auto_backup: %v", err)
		}
	}

	if err = d.Set("ip", dataSourceFlattenSystemFortimanagerIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("central_mgmt_schedule_config_restore", dataSourceFlattenSystemFortimanagerCentralMgmtScheduleConfigRestore(o["central-mgmt-schedule-config-restore"], d, "central_mgmt_schedule_config_restore")); err != nil {
		if !fortiAPIPatch(o["central-mgmt-schedule-config-restore"]) {
			return fmt.Errorf("Error reading central_mgmt_schedule_config_restore: %v", err)
		}
	}

	if err = d.Set("central_mgmt_schedule_script_restore", dataSourceFlattenSystemFortimanagerCentralMgmtScheduleScriptRestore(o["central-mgmt-schedule-script-restore"], d, "central_mgmt_schedule_script_restore")); err != nil {
		if !fortiAPIPatch(o["central-mgmt-schedule-script-restore"]) {
			return fmt.Errorf("Error reading central_mgmt_schedule_script_restore: %v", err)
		}
	}

	if err = d.Set("ipsec", dataSourceFlattenSystemFortimanagerIpsec(o["ipsec"], d, "ipsec")); err != nil {
		if !fortiAPIPatch(o["ipsec"]) {
			return fmt.Errorf("Error reading ipsec: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemFortimanagerFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
