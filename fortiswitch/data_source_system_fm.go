// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Fm.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemFm() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemFmRead,
		Schema: map[string]*schema.Schema{

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_backup": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"scheduled_config_restore": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fswid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
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

func dataSourceSystemFmRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemFm"

	o, err := c.ReadSystemFm(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemFm: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemFm(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemFm from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemFmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFmAutoBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFmIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFmScheduledConfigRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFmId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFmVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFmIpsec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemFm(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenSystemFmStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("auto_backup", dataSourceFlattenSystemFmAutoBackup(o["auto-backup"], d, "auto_backup")); err != nil {
		if !fortiAPIPatch(o["auto-backup"]) {
			return fmt.Errorf("Error reading auto_backup: %v", err)
		}
	}

	if err = d.Set("ip", dataSourceFlattenSystemFmIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("scheduled_config_restore", dataSourceFlattenSystemFmScheduledConfigRestore(o["scheduled-config-restore"], d, "scheduled_config_restore")); err != nil {
		if !fortiAPIPatch(o["scheduled-config-restore"]) {
			return fmt.Errorf("Error reading scheduled_config_restore: %v", err)
		}
	}

	if err = d.Set("fswid", dataSourceFlattenSystemFmId(o["id"], d, "fswid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fswid: %v", err)
		}
	}

	if err = d.Set("vdom", dataSourceFlattenSystemFmVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("ipsec", dataSourceFlattenSystemFmIpsec(o["ipsec"], d, "ipsec")); err != nil {
		if !fortiAPIPatch(o["ipsec"]) {
			return fmt.Errorf("Error reading ipsec: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemFmFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
