// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure access-control lists global settings on Switch.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchAclSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchAclSettingsUpdate,
		Read:   resourceSwitchAclSettingsRead,
		Update: resourceSwitchAclSettingsUpdate,
		Delete: resourceSwitchAclSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"density_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trunk_load_balance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchAclSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAclSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAclSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchAclSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAclSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchAclSettings")
	}

	return resourceSwitchAclSettingsRead(d, m)
}

func resourceSwitchAclSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAclSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchAclSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchAclSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchAclSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchAclSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchAclSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAclSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchAclSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAclSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchAclSettingsDensityMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclSettingsTrunkLoadBalance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchAclSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("density_mode", flattenSwitchAclSettingsDensityMode(o["density-mode"], d, "density_mode", sv)); err != nil {
		if !fortiAPIPatch(o["density-mode"]) {
			return fmt.Errorf("Error reading density_mode: %v", err)
		}
	}

	if err = d.Set("trunk_load_balance", flattenSwitchAclSettingsTrunkLoadBalance(o["trunk-load-balance"], d, "trunk_load_balance", sv)); err != nil {
		if !fortiAPIPatch(o["trunk-load-balance"]) {
			return fmt.Errorf("Error reading trunk_load_balance: %v", err)
		}
	}

	return nil
}

func flattenSwitchAclSettingsFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchAclSettingsDensityMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclSettingsTrunkLoadBalance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchAclSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("density_mode"); ok {
		if setArgNil {
			obj["density-mode"] = nil
		} else {

			t, err := expandSwitchAclSettingsDensityMode(d, v, "density_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["density-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("trunk_load_balance"); ok {
		if setArgNil {
			obj["trunk-load-balance"] = nil
		} else {

			t, err := expandSwitchAclSettingsTrunkLoadBalance(d, v, "trunk_load_balance", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trunk-load-balance"] = t
			}
		}
	}

	return &obj, nil
}
