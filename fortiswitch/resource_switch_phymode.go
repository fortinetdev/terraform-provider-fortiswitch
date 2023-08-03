// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: PHY configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchPhyMode() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchPhyModeUpdate,
		Read:   resourceSwitchPhyModeRead,
		Update: resourceSwitchPhyModeUpdate,
		Delete: resourceSwitchPhyModeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"port_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port54_phy_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port53_phy_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchPhyModeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchPhyMode(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchPhyMode resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchPhyMode(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchPhyMode resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchPhyMode")
	}

	return resourceSwitchPhyModeRead(d, m)
}

func resourceSwitchPhyModeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchPhyMode(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchPhyMode resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchPhyMode(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchPhyMode resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchPhyModeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchPhyMode(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchPhyMode resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchPhyMode(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchPhyMode resource from API: %v", err)
	}
	return nil
}

func flattenSwitchPhyModePortConfiguration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhyModePort54PhyMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchPhyModePort53PhyMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchPhyMode(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("port_configuration", flattenSwitchPhyModePortConfiguration(o["port-configuration"], d, "port_configuration", sv)); err != nil {
		if !fortiAPIPatch(o["port-configuration"]) {
			return fmt.Errorf("Error reading port_configuration: %v", err)
		}
	}

	if err = d.Set("port54_phy_mode", flattenSwitchPhyModePort54PhyMode(o["port54-phy-mode"], d, "port54_phy_mode", sv)); err != nil {
		if !fortiAPIPatch(o["port54-phy-mode"]) {
			return fmt.Errorf("Error reading port54_phy_mode: %v", err)
		}
	}

	if err = d.Set("port53_phy_mode", flattenSwitchPhyModePort53PhyMode(o["port53-phy-mode"], d, "port53_phy_mode", sv)); err != nil {
		if !fortiAPIPatch(o["port53-phy-mode"]) {
			return fmt.Errorf("Error reading port53_phy_mode: %v", err)
		}
	}

	return nil
}

func flattenSwitchPhyModeFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchPhyModePortConfiguration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhyModePort54PhyMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchPhyModePort53PhyMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchPhyMode(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("port_configuration"); ok {
		if setArgNil {
			obj["port-configuration"] = nil
		} else {

			t, err := expandSwitchPhyModePortConfiguration(d, v, "port_configuration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["port-configuration"] = t
			}
		}
	}

	if v, ok := d.GetOk("port54_phy_mode"); ok {
		if setArgNil {
			obj["port54-phy-mode"] = nil
		} else {

			t, err := expandSwitchPhyModePort54PhyMode(d, v, "port54_phy_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["port54-phy-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("port53_phy_mode"); ok {
		if setArgNil {
			obj["port53-phy-mode"] = nil
		} else {

			t, err := expandSwitchPhyModePort53PhyMode(d, v, "port53_phy_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["port53-phy-mode"] = t
			}
		}
	}

	return &obj, nil
}
