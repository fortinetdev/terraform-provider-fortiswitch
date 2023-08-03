// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure virtual wire.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchVirtualWire() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchVirtualWireCreate,
		Read:   resourceSwitchVirtualWireRead,
		Update: resourceSwitchVirtualWireUpdate,
		Delete: resourceSwitchVirtualWireDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"first_member": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"second_member": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchVirtualWireCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchVirtualWire(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchVirtualWire resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchVirtualWire(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchVirtualWire resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchVirtualWire")
	}

	return resourceSwitchVirtualWireRead(d, m)
}

func resourceSwitchVirtualWireUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchVirtualWire(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchVirtualWire resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchVirtualWire(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchVirtualWire resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchVirtualWire")
	}

	return resourceSwitchVirtualWireRead(d, m)
}

func resourceSwitchVirtualWireDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchVirtualWire(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchVirtualWire resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchVirtualWireRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchVirtualWire(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchVirtualWire resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchVirtualWire(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchVirtualWire resource from API: %v", err)
	}
	return nil
}

func flattenSwitchVirtualWireFirstMember(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVirtualWireVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVirtualWireName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVirtualWireSecondMember(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchVirtualWire(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("first_member", flattenSwitchVirtualWireFirstMember(o["first-member"], d, "first_member", sv)); err != nil {
		if !fortiAPIPatch(o["first-member"]) {
			return fmt.Errorf("Error reading first_member: %v", err)
		}
	}

	if err = d.Set("vlan", flattenSwitchVirtualWireVlan(o["vlan"], d, "vlan", sv)); err != nil {
		if !fortiAPIPatch(o["vlan"]) {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchVirtualWireName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("second_member", flattenSwitchVirtualWireSecondMember(o["second-member"], d, "second_member", sv)); err != nil {
		if !fortiAPIPatch(o["second-member"]) {
			return fmt.Errorf("Error reading second_member: %v", err)
		}
	}

	return nil
}

func flattenSwitchVirtualWireFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchVirtualWireFirstMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVirtualWireVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVirtualWireName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVirtualWireSecondMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchVirtualWire(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("first_member"); ok {

		t, err := expandSwitchVirtualWireFirstMember(d, v, "first_member", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["first-member"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok {

		t, err := expandSwitchVirtualWireVlan(d, v, "vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchVirtualWireName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("second_member"); ok {

		t, err := expandSwitchVirtualWireSecondMember(d, v, "second_member", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["second-member"] = t
		}
	}

	return &obj, nil
}
