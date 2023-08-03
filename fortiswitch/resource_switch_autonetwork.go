// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Auto network.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchAutoNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchAutoNetworkUpdate,
		Read:   resourceSwitchAutoNetworkRead,
		Update: resourceSwitchAutoNetworkUpdate,
		Delete: resourceSwitchAutoNetworkDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mgmt_vlan": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4094),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchAutoNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAutoNetwork(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAutoNetwork resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchAutoNetwork(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAutoNetwork resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchAutoNetwork")
	}

	return resourceSwitchAutoNetworkRead(d, m)
}

func resourceSwitchAutoNetworkDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAutoNetwork(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchAutoNetwork resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchAutoNetwork(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchAutoNetwork resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchAutoNetworkRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchAutoNetwork(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAutoNetwork resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchAutoNetwork(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAutoNetwork resource from API: %v", err)
	}
	return nil
}

func flattenSwitchAutoNetworkStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAutoNetworkMgmtVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchAutoNetwork(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSwitchAutoNetworkStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("mgmt_vlan", flattenSwitchAutoNetworkMgmtVlan(o["mgmt-vlan"], d, "mgmt_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-vlan"]) {
			return fmt.Errorf("Error reading mgmt_vlan: %v", err)
		}
	}

	return nil
}

func flattenSwitchAutoNetworkFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchAutoNetworkStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAutoNetworkMgmtVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchAutoNetwork(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandSwitchAutoNetworkStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("mgmt_vlan"); ok {
		if setArgNil {
			obj["mgmt-vlan"] = nil
		} else {

			t, err := expandSwitchAutoNetworkMgmtVlan(d, v, "mgmt_vlan", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mgmt-vlan"] = t
			}
		}
	}

	return &obj, nil
}
