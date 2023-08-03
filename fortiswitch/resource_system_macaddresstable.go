// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Mac address table.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemMacAddressTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemMacAddressTableCreate,
		Read:   resourceSystemMacAddressTableRead,
		Update: resourceSystemMacAddressTableUpdate,
		Delete: resourceSystemMacAddressTableDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemMacAddressTableCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemMacAddressTable(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemMacAddressTable resource while getting object: %v", err)
	}

	o, err := c.CreateSystemMacAddressTable(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemMacAddressTable resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemMacAddressTable")
	}

	return resourceSystemMacAddressTableRead(d, m)
}

func resourceSystemMacAddressTableUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemMacAddressTable(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemMacAddressTable resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemMacAddressTable(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemMacAddressTable resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemMacAddressTable")
	}

	return resourceSystemMacAddressTableRead(d, m)
}

func resourceSystemMacAddressTableDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemMacAddressTable(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemMacAddressTable resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemMacAddressTableRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemMacAddressTable(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemMacAddressTable resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemMacAddressTable(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemMacAddressTable resource from API: %v", err)
	}
	return nil
}

func flattenSystemMacAddressTableInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMacAddressTableMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemMacAddressTable(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("interface", flattenSystemMacAddressTableInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("mac", flattenSystemMacAddressTableMac(o["mac"], d, "mac", sv)); err != nil {
		if !fortiAPIPatch(o["mac"]) {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	return nil
}

func flattenSystemMacAddressTableFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemMacAddressTableInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMacAddressTableMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemMacAddressTable(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandSystemMacAddressTableInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok {

		t, err := expandSystemMacAddressTableMac(d, v, "mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	return &obj, nil
}
