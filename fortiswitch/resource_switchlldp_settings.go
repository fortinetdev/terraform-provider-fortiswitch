// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Global LLDP configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchLldpSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchLldpSettingsUpdate,
		Read:   resourceSwitchLldpSettingsRead,
		Update: resourceSwitchLldpSettingsUpdate,
		Delete: resourceSwitchLldpSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tx_hold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16),
				Optional:     true,
				Computed:     true,
			},
			"tx_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 4095),
				Optional:     true,
				Computed:     true,
			},
			"management_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"fast_start_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchLldpSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchLldpSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchLldpSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchLldpSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchLldpSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchLldpSettings")
	}

	return resourceSwitchLldpSettingsRead(d, m)
}

func resourceSwitchLldpSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchLldpSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchLldpSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchLldpSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchLldpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchLldpSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchLldpSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchLldpSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchLldpSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchLldpSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchLldpSettingsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpSettingsDeviceDetection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpSettingsManagementAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpSettingsTxHold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpSettingsTxInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpSettingsManagementInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchLldpSettingsFastStartInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchLldpSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSwitchLldpSettingsStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("device_detection", flattenSwitchLldpSettingsDeviceDetection(o["device-detection"], d, "device_detection", sv)); err != nil {
		if !fortiAPIPatch(o["device-detection"]) {
			return fmt.Errorf("Error reading device_detection: %v", err)
		}
	}

	if err = d.Set("management_address", flattenSwitchLldpSettingsManagementAddress(o["management-address"], d, "management_address", sv)); err != nil {
		if !fortiAPIPatch(o["management-address"]) {
			return fmt.Errorf("Error reading management_address: %v", err)
		}
	}

	if err = d.Set("tx_hold", flattenSwitchLldpSettingsTxHold(o["tx-hold"], d, "tx_hold", sv)); err != nil {
		if !fortiAPIPatch(o["tx-hold"]) {
			return fmt.Errorf("Error reading tx_hold: %v", err)
		}
	}

	if err = d.Set("tx_interval", flattenSwitchLldpSettingsTxInterval(o["tx-interval"], d, "tx_interval", sv)); err != nil {
		if !fortiAPIPatch(o["tx-interval"]) {
			return fmt.Errorf("Error reading tx_interval: %v", err)
		}
	}

	if err = d.Set("management_interface", flattenSwitchLldpSettingsManagementInterface(o["management-interface"], d, "management_interface", sv)); err != nil {
		if !fortiAPIPatch(o["management-interface"]) {
			return fmt.Errorf("Error reading management_interface: %v", err)
		}
	}

	if err = d.Set("fast_start_interval", flattenSwitchLldpSettingsFastStartInterval(o["fast-start-interval"], d, "fast_start_interval", sv)); err != nil {
		if !fortiAPIPatch(o["fast-start-interval"]) {
			return fmt.Errorf("Error reading fast_start_interval: %v", err)
		}
	}

	return nil
}

func flattenSwitchLldpSettingsFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchLldpSettingsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpSettingsDeviceDetection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpSettingsManagementAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpSettingsTxHold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpSettingsTxInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpSettingsManagementInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchLldpSettingsFastStartInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchLldpSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandSwitchLldpSettingsStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("device_detection"); ok {
		if setArgNil {
			obj["device-detection"] = nil
		} else {

			t, err := expandSwitchLldpSettingsDeviceDetection(d, v, "device_detection", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["device-detection"] = t
			}
		}
	}

	if v, ok := d.GetOk("management_address"); ok {
		if setArgNil {
			obj["management-address"] = nil
		} else {

			t, err := expandSwitchLldpSettingsManagementAddress(d, v, "management_address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["management-address"] = t
			}
		}
	}

	if v, ok := d.GetOk("tx_hold"); ok {
		if setArgNil {
			obj["tx-hold"] = nil
		} else {

			t, err := expandSwitchLldpSettingsTxHold(d, v, "tx_hold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tx-hold"] = t
			}
		}
	}

	if v, ok := d.GetOk("tx_interval"); ok {
		if setArgNil {
			obj["tx-interval"] = nil
		} else {

			t, err := expandSwitchLldpSettingsTxInterval(d, v, "tx_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tx-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("management_interface"); ok {
		if setArgNil {
			obj["management-interface"] = nil
		} else {

			t, err := expandSwitchLldpSettingsManagementInterface(d, v, "management_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["management-interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("fast_start_interval"); ok {
		if setArgNil {
			obj["fast-start-interval"] = nil
		} else {

			t, err := expandSwitchLldpSettingsFastStartInterval(d, v, "fast_start_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fast-start-interval"] = t
			}
		}
	}

	return &obj, nil
}
