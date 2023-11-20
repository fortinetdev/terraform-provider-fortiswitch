// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: PTP policy configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemPtpInterfacePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPtpInterfacePolicyCreate,
		Read:   resourceSystemPtpInterfacePolicyRead,
		Update: resourceSystemPtpInterfacePolicyUpdate,
		Delete: resourceSystemPtpInterfacePolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vlan": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4094),
				Optional:     true,
				Computed:     true,
			},
			"vlan_pri": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemPtpInterfacePolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemPtpInterfacePolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemPtpInterfacePolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSystemPtpInterfacePolicy(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemPtpInterfacePolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemPtpInterfacePolicy")
	}

	return resourceSystemPtpInterfacePolicyRead(d, m)
}

func resourceSystemPtpInterfacePolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemPtpInterfacePolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemPtpInterfacePolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemPtpInterfacePolicy(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemPtpInterfacePolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemPtpInterfacePolicy")
	}

	return resourceSystemPtpInterfacePolicyRead(d, m)
}

func resourceSystemPtpInterfacePolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemPtpInterfacePolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPtpInterfacePolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPtpInterfacePolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemPtpInterfacePolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemPtpInterfacePolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPtpInterfacePolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemPtpInterfacePolicy resource from API: %v", err)
	}
	return nil
}

func flattenSystemPtpInterfacePolicyVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpInterfacePolicyVlanPri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpInterfacePolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpInterfacePolicyDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemPtpInterfacePolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("vlan", flattenSystemPtpInterfacePolicyVlan(o["vlan"], d, "vlan", sv)); err != nil {
		if !fortiAPIPatch(o["vlan"]) {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	if err = d.Set("vlan_pri", flattenSystemPtpInterfacePolicyVlanPri(o["vlan-pri"], d, "vlan_pri", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-pri"]) {
			return fmt.Errorf("Error reading vlan_pri: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemPtpInterfacePolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemPtpInterfacePolicyDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	return nil
}

func flattenSystemPtpInterfacePolicyFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemPtpInterfacePolicyVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpInterfacePolicyVlanPri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpInterfacePolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpInterfacePolicyDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPtpInterfacePolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("vlan"); ok {

		t, err := expandSystemPtpInterfacePolicyVlan(d, v, "vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	if v, ok := d.GetOkExists("vlan_pri"); ok {

		t, err := expandSystemPtpInterfacePolicyVlanPri(d, v, "vlan_pri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-pri"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemPtpInterfacePolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSystemPtpInterfacePolicyDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	return &obj, nil
}
