// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Global PTP configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchPtpSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchPtpSettingsUpdate,
		Read:   resourceSwitchPtpSettingsRead,
		Update: resourceSwitchPtpSettingsUpdate,
		Delete: resourceSwitchPtpSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchPtpSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchPtpSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchPtpSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchPtpSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchPtpSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchPtpSettings")
	}

	return resourceSwitchPtpSettingsRead(d, m)
}

func resourceSwitchPtpSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchPtpSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchPtpSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchPtpSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchPtpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchPtpSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchPtpSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchPtpSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchPtpSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchPtpSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchPtpSettingsMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchPtpSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mode", flattenSwitchPtpSettingsMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	return nil
}

func flattenSwitchPtpSettingsFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchPtpSettingsMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchPtpSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mode"); ok {
		if setArgNil {
			obj["mode"] = nil
		} else {

			t, err := expandSwitchPtpSettingsMode(d, v, "mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mode"] = t
			}
		}
	}

	return &obj, nil
}
