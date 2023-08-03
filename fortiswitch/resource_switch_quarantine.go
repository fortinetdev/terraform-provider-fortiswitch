// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure quarantine devices on the switch.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchQuarantine() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchQuarantineCreate,
		Read:   resourceSwitchQuarantineRead,
		Update: resourceSwitchQuarantineUpdate,
		Delete: resourceSwitchQuarantineDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"cos_queue": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"drop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"policer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"acl_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchQuarantineCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchQuarantine(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchQuarantine resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchQuarantine(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchQuarantine resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchQuarantine")
	}

	return resourceSwitchQuarantineRead(d, m)
}

func resourceSwitchQuarantineUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchQuarantine(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchQuarantine resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchQuarantine(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchQuarantine resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchQuarantine")
	}

	return resourceSwitchQuarantineRead(d, m)
}

func resourceSwitchQuarantineDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchQuarantine(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchQuarantine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchQuarantineRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchQuarantine(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchQuarantine resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchQuarantine(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchQuarantine resource from API: %v", err)
	}
	return nil
}

func flattenSwitchQuarantineDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQuarantineCosQueue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQuarantineDrop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQuarantineMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQuarantinePolicer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQuarantineAclId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchQuarantine(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("description", flattenSwitchQuarantineDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("cos_queue", flattenSwitchQuarantineCosQueue(o["cos-queue"], d, "cos_queue", sv)); err != nil {
		if !fortiAPIPatch(o["cos-queue"]) {
			return fmt.Errorf("Error reading cos_queue: %v", err)
		}
	}

	if err = d.Set("drop", flattenSwitchQuarantineDrop(o["drop"], d, "drop", sv)); err != nil {
		if !fortiAPIPatch(o["drop"]) {
			return fmt.Errorf("Error reading drop: %v", err)
		}
	}

	if err = d.Set("mac", flattenSwitchQuarantineMac(o["mac"], d, "mac", sv)); err != nil {
		if !fortiAPIPatch(o["mac"]) {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("policer", flattenSwitchQuarantinePolicer(o["policer"], d, "policer", sv)); err != nil {
		if !fortiAPIPatch(o["policer"]) {
			return fmt.Errorf("Error reading policer: %v", err)
		}
	}

	if err = d.Set("acl_id", flattenSwitchQuarantineAclId(o["acl-id"], d, "acl_id", sv)); err != nil {
		if !fortiAPIPatch(o["acl-id"]) {
			return fmt.Errorf("Error reading acl_id: %v", err)
		}
	}

	return nil
}

func flattenSwitchQuarantineFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchQuarantineDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQuarantineCosQueue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQuarantineDrop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQuarantineMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQuarantinePolicer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQuarantineAclId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchQuarantine(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchQuarantineDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOkExists("cos_queue"); ok {

		t, err := expandSwitchQuarantineCosQueue(d, v, "cos_queue", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos-queue"] = t
		}
	}

	if v, ok := d.GetOk("drop"); ok {

		t, err := expandSwitchQuarantineDrop(d, v, "drop", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok {

		t, err := expandSwitchQuarantineMac(d, v, "mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("policer"); ok {

		t, err := expandSwitchQuarantinePolicer(d, v, "policer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policer"] = t
		}
	}

	if v, ok := d.GetOk("acl_id"); ok {

		t, err := expandSwitchQuarantineAclId(d, v, "acl_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acl-id"] = t
		}
	}

	return &obj, nil
}
