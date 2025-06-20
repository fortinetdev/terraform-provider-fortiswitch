// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Vlan Pruning.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchVlanPruning() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchVlanPruningUpdate,
		Read:   resourceSwitchVlanPruningRead,
		Update: resourceSwitchVlanPruningUpdate,
		Delete: resourceSwitchVlanPruningDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"join_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(20, 1000),
				Optional:     true,
				Computed:     true,
			},
			"leave_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(600, 3000),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchVlanPruningUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchVlanPruning(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchVlanPruning resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchVlanPruning(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchVlanPruning resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchVlanPruning")
	}

	return resourceSwitchVlanPruningRead(d, m)
}

func resourceSwitchVlanPruningDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchVlanPruning(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchVlanPruning resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchVlanPruning(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchVlanPruning resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchVlanPruningRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchVlanPruning(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchVlanPruning resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchVlanPruning(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchVlanPruning resource from API: %v", err)
	}
	return nil
}

func flattenSwitchVlanPruningJoinTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanPruningLeaveTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchVlanPruning(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("join_timer", flattenSwitchVlanPruningJoinTimer(o["join-timer"], d, "join_timer", sv)); err != nil {
		if !fortiAPIPatch(o["join-timer"]) {
			return fmt.Errorf("Error reading join_timer: %v", err)
		}
	}

	if err = d.Set("leave_timer", flattenSwitchVlanPruningLeaveTimer(o["leave-timer"], d, "leave_timer", sv)); err != nil {
		if !fortiAPIPatch(o["leave-timer"]) {
			return fmt.Errorf("Error reading leave_timer: %v", err)
		}
	}

	return nil
}

func flattenSwitchVlanPruningFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchVlanPruningJoinTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanPruningLeaveTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchVlanPruning(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("join_timer"); ok {
		if setArgNil {
			obj["join-timer"] = nil
		} else {

			t, err := expandSwitchVlanPruningJoinTimer(d, v, "join_timer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["join-timer"] = t
			}
		}
	}

	if v, ok := d.GetOk("leave_timer"); ok {
		if setArgNil {
			obj["leave-timer"] = nil
		} else {

			t, err := expandSwitchVlanPruningLeaveTimer(d, v, "leave_timer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["leave-timer"] = t
			}
		}
	}

	return &obj, nil
}
