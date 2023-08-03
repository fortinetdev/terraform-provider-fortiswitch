// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure switch global ether-types.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchVlanTpid() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchVlanTpidCreate,
		Read:   resourceSwitchVlanTpidRead,
		Update: resourceSwitchVlanTpidUpdate,
		Delete: resourceSwitchVlanTpidDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"ether_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchVlanTpidCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchVlanTpid(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchVlanTpid resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchVlanTpid(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchVlanTpid resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchVlanTpid")
	}

	return resourceSwitchVlanTpidRead(d, m)
}

func resourceSwitchVlanTpidUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchVlanTpid(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchVlanTpid resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchVlanTpid(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchVlanTpid resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchVlanTpid")
	}

	return resourceSwitchVlanTpidRead(d, m)
}

func resourceSwitchVlanTpidDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchVlanTpid(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchVlanTpid resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchVlanTpidRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchVlanTpid(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchVlanTpid resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchVlanTpid(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchVlanTpid resource from API: %v", err)
	}
	return nil
}

func flattenSwitchVlanTpidName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchVlanTpidEtherType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchVlanTpid(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchVlanTpidName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ether_type", flattenSwitchVlanTpidEtherType(o["ether-type"], d, "ether_type", sv)); err != nil {
		if !fortiAPIPatch(o["ether-type"]) {
			return fmt.Errorf("Error reading ether_type: %v", err)
		}
	}

	return nil
}

func flattenSwitchVlanTpidFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchVlanTpidName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchVlanTpidEtherType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchVlanTpid(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchVlanTpidName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ether_type"); ok {

		t, err := expandSwitchVlanTpidEtherType(d, v, "ether_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ether-type"] = t
		}
	}

	return &obj, nil
}
