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

func resourceSystemFm() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFmUpdate,
		Read:   resourceSystemFmRead,
		Update: resourceSystemFmUpdate,
		Delete: resourceSystemFmDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_backup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scheduled_config_restore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fswid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 11),
				Optional:     true,
				Computed:     true,
			},
			"ipsec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemFmUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemFm(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemFm resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFm(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemFm resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFm")
	}

	return resourceSystemFmRead(d, m)
}

func resourceSystemFmDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemFm(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemFm resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFm(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemFm resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFmRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemFm(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemFm resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFm(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemFm resource from API: %v", err)
	}
	return nil
}

func flattenSystemFmStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFmAutoBackup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFmIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFmScheduledConfigRestore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFmId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFmVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFmIpsec(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemFm(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemFmStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("auto_backup", flattenSystemFmAutoBackup(o["auto-backup"], d, "auto_backup", sv)); err != nil {
		if !fortiAPIPatch(o["auto-backup"]) {
			return fmt.Errorf("Error reading auto_backup: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemFmIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("scheduled_config_restore", flattenSystemFmScheduledConfigRestore(o["scheduled-config-restore"], d, "scheduled_config_restore", sv)); err != nil {
		if !fortiAPIPatch(o["scheduled-config-restore"]) {
			return fmt.Errorf("Error reading scheduled_config_restore: %v", err)
		}
	}

	if err = d.Set("fswid", flattenSystemFmId(o["id"], d, "fswid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fswid: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemFmVdom(o["vdom"], d, "vdom", sv)); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("ipsec", flattenSystemFmIpsec(o["ipsec"], d, "ipsec", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec"]) {
			return fmt.Errorf("Error reading ipsec: %v", err)
		}
	}

	return nil
}

func flattenSystemFmFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemFmStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFmAutoBackup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFmIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFmScheduledConfigRestore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFmId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFmVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFmIpsec(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFm(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandSystemFmStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_backup"); ok {
		if setArgNil {
			obj["auto-backup"] = nil
		} else {

			t, err := expandSystemFmAutoBackup(d, v, "auto_backup", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-backup"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		if setArgNil {
			obj["ip"] = nil
		} else {

			t, err := expandSystemFmIp(d, v, "ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("scheduled_config_restore"); ok {
		if setArgNil {
			obj["scheduled-config-restore"] = nil
		} else {

			t, err := expandSystemFmScheduledConfigRestore(d, v, "scheduled_config_restore", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["scheduled-config-restore"] = t
			}
		}
	}

	if v, ok := d.GetOk("fswid"); ok {
		if setArgNil {
			obj["id"] = nil
		} else {

			t, err := expandSystemFmId(d, v, "fswid", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["id"] = t
			}
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		if setArgNil {
			obj["vdom"] = nil
		} else {

			t, err := expandSystemFmVdom(d, v, "vdom", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vdom"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec"); ok {
		if setArgNil {
			obj["ipsec"] = nil
		} else {

			t, err := expandSystemFmIpsec(d, v, "ipsec", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec"] = t
			}
		}
	}

	return &obj, nil
}
