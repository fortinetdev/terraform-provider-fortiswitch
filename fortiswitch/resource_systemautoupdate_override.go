// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure override FDS server.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutoupdateOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoupdateOverrideUpdate,
		Read:   resourceSystemAutoupdateOverrideRead,
		Update: resourceSystemAutoupdateOverrideUpdate,
		Delete: resourceSystemAutoupdateOverrideDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fail_over": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemAutoupdateOverrideUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAutoupdateOverride(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdateOverride resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutoupdateOverride(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdateOverride resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutoupdateOverride")
	}

	return resourceSystemAutoupdateOverrideRead(d, m)
}

func resourceSystemAutoupdateOverrideDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAutoupdateOverride(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdateOverride resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoupdateOverride(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemAutoupdateOverride resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoupdateOverrideRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAutoupdateOverride(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoupdateOverride resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoupdateOverride(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoupdateOverride resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutoupdateOverrideStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoupdateOverrideFailOver(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoupdateOverrideAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAutoupdateOverride(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemAutoupdateOverrideStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("fail_over", flattenSystemAutoupdateOverrideFailOver(o["fail-over"], d, "fail_over", sv)); err != nil {
		if !fortiAPIPatch(o["fail-over"]) {
			return fmt.Errorf("Error reading fail_over: %v", err)
		}
	}

	if err = d.Set("address", flattenSystemAutoupdateOverrideAddress(o["address"], d, "address", sv)); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	return nil
}

func flattenSystemAutoupdateOverrideFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemAutoupdateOverrideStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdateOverrideFailOver(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdateOverrideAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoupdateOverride(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandSystemAutoupdateOverrideStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("fail_over"); ok {
		if setArgNil {
			obj["fail-over"] = nil
		} else {

			t, err := expandSystemAutoupdateOverrideFailOver(d, v, "fail_over", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fail-over"] = t
			}
		}
	}

	if v, ok := d.GetOk("address"); ok {
		if setArgNil {
			obj["address"] = nil
		} else {

			t, err := expandSystemAutoupdateOverrideAddress(d, v, "address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["address"] = t
			}
		}
	}

	return &obj, nil
}
