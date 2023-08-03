// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Logging device to display in GUI.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogGui() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogGuiUpdate,
		Read:   resourceLogGuiRead,
		Update: resourceLogGuiUpdate,
		Delete: resourceLogGuiDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"log_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogGuiUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogGui(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogGui resource while getting object: %v", err)
	}

	o, err := c.UpdateLogGui(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogGui resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogGui")
	}

	return resourceLogGuiRead(d, m)
}

func resourceLogGuiDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogGui(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogGui resource while getting object: %v", err)
	}

	_, err = c.UpdateLogGui(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing LogGui resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogGuiRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogGui(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogGui resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogGui(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogGui resource from API: %v", err)
	}
	return nil
}

func flattenLogGuiLogDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogGui(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("log_device", flattenLogGuiLogDevice(o["log-device"], d, "log_device", sv)); err != nil {
		if !fortiAPIPatch(o["log-device"]) {
			return fmt.Errorf("Error reading log_device: %v", err)
		}
	}

	return nil
}

func flattenLogGuiFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandLogGuiLogDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogGui(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("log_device"); ok {
		if setArgNil {
			obj["log-device"] = nil
		} else {

			t, err := expandLogGuiLogDevice(d, v, "log_device", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-device"] = t
			}
		}
	}

	return &obj, nil
}
