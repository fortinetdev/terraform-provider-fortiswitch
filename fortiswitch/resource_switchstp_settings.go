// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Switch-global stp settings.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchStpSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchStpSettingsUpdate,
		Read:   resourceSwitchStpSettingsRead,
		Update: resourceSwitchStpSettingsUpdate,
		Delete: resourceSwitchStpSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hello_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"mclag_stp_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pending_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"forward_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mclag_stp_bpdu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_hops": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"revision": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchStpSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchStpSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchStpSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchStpSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchStpSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchStpSettings")
	}

	return resourceSwitchStpSettingsRead(d, m)
}

func resourceSwitchStpSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchStpSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchStpSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchStpSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchStpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchStpSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchStpSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchStpSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchStpSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchStpSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchStpSettingsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpSettingsHelloTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpSettingsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpSettingsMclagStpMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpSettingsPendingTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpSettingsForwardTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpSettingsFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpSettingsMclagStpBpdu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpSettingsMaxHops(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpSettingsMaxAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpSettingsRevision(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchStpSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSwitchStpSettingsStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("hello_time", flattenSwitchStpSettingsHelloTime(o["hello-time"], d, "hello_time", sv)); err != nil {
		if !fortiAPIPatch(o["hello-time"]) {
			return fmt.Errorf("Error reading hello_time: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchStpSettingsName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("mclag_stp_mac", flattenSwitchStpSettingsMclagStpMac(o["mclag-stp-mac"], d, "mclag_stp_mac", sv)); err != nil {
		if !fortiAPIPatch(o["mclag-stp-mac"]) {
			return fmt.Errorf("Error reading mclag_stp_mac: %v", err)
		}
	}

	if err = d.Set("pending_timer", flattenSwitchStpSettingsPendingTimer(o["pending-timer"], d, "pending_timer", sv)); err != nil {
		if !fortiAPIPatch(o["pending-timer"]) {
			return fmt.Errorf("Error reading pending_timer: %v", err)
		}
	}

	if err = d.Set("forward_time", flattenSwitchStpSettingsForwardTime(o["forward-time"], d, "forward_time", sv)); err != nil {
		if !fortiAPIPatch(o["forward-time"]) {
			return fmt.Errorf("Error reading forward_time: %v", err)
		}
	}

	if err = d.Set("flood", flattenSwitchStpSettingsFlood(o["flood"], d, "flood", sv)); err != nil {
		if !fortiAPIPatch(o["flood"]) {
			return fmt.Errorf("Error reading flood: %v", err)
		}
	}

	if err = d.Set("mclag_stp_bpdu", flattenSwitchStpSettingsMclagStpBpdu(o["mclag-stp-bpdu"], d, "mclag_stp_bpdu", sv)); err != nil {
		if !fortiAPIPatch(o["mclag-stp-bpdu"]) {
			return fmt.Errorf("Error reading mclag_stp_bpdu: %v", err)
		}
	}

	if err = d.Set("max_hops", flattenSwitchStpSettingsMaxHops(o["max-hops"], d, "max_hops", sv)); err != nil {
		if !fortiAPIPatch(o["max-hops"]) {
			return fmt.Errorf("Error reading max_hops: %v", err)
		}
	}

	if err = d.Set("max_age", flattenSwitchStpSettingsMaxAge(o["max-age"], d, "max_age", sv)); err != nil {
		if !fortiAPIPatch(o["max-age"]) {
			return fmt.Errorf("Error reading max_age: %v", err)
		}
	}

	if err = d.Set("revision", flattenSwitchStpSettingsRevision(o["revision"], d, "revision", sv)); err != nil {
		if !fortiAPIPatch(o["revision"]) {
			return fmt.Errorf("Error reading revision: %v", err)
		}
	}

	return nil
}

func flattenSwitchStpSettingsFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchStpSettingsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpSettingsHelloTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpSettingsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpSettingsMclagStpMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpSettingsPendingTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpSettingsForwardTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpSettingsFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpSettingsMclagStpBpdu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpSettingsMaxHops(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpSettingsMaxAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpSettingsRevision(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchStpSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandSwitchStpSettingsStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("hello_time"); ok {
		if setArgNil {
			obj["hello-time"] = nil
		} else {

			t, err := expandSwitchStpSettingsHelloTime(d, v, "hello_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hello-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("name"); ok {
		if setArgNil {
			obj["name"] = nil
		} else {

			t, err := expandSwitchStpSettingsName(d, v, "name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["name"] = t
			}
		}
	}

	if v, ok := d.GetOk("mclag_stp_mac"); ok {
		if setArgNil {
			obj["mclag-stp-mac"] = nil
		} else {

			t, err := expandSwitchStpSettingsMclagStpMac(d, v, "mclag_stp_mac", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mclag-stp-mac"] = t
			}
		}
	}

	if v, ok := d.GetOk("pending_timer"); ok {
		if setArgNil {
			obj["pending-timer"] = nil
		} else {

			t, err := expandSwitchStpSettingsPendingTimer(d, v, "pending_timer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pending-timer"] = t
			}
		}
	}

	if v, ok := d.GetOk("forward_time"); ok {
		if setArgNil {
			obj["forward-time"] = nil
		} else {

			t, err := expandSwitchStpSettingsForwardTime(d, v, "forward_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["forward-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("flood"); ok {
		if setArgNil {
			obj["flood"] = nil
		} else {

			t, err := expandSwitchStpSettingsFlood(d, v, "flood", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["flood"] = t
			}
		}
	}

	if v, ok := d.GetOk("mclag_stp_bpdu"); ok {
		if setArgNil {
			obj["mclag-stp-bpdu"] = nil
		} else {

			t, err := expandSwitchStpSettingsMclagStpBpdu(d, v, "mclag_stp_bpdu", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mclag-stp-bpdu"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_hops"); ok {
		if setArgNil {
			obj["max-hops"] = nil
		} else {

			t, err := expandSwitchStpSettingsMaxHops(d, v, "max_hops", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-hops"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_age"); ok {
		if setArgNil {
			obj["max-age"] = nil
		} else {

			t, err := expandSwitchStpSettingsMaxAge(d, v, "max_age", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-age"] = t
			}
		}
	}

	if v, ok := d.GetOk("revision"); ok {
		if setArgNil {
			obj["revision"] = nil
		} else {

			t, err := expandSwitchStpSettingsRevision(d, v, "revision", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["revision"] = t
			}
		}
	}

	return &obj, nil
}
