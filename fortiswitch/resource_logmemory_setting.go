// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Settings for memory buffer.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogMemorySetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogMemorySettingUpdate,
		Read:   resourceLogMemorySettingRead,
		Update: resourceLogMemorySettingUpdate,
		Delete: resourceLogMemorySettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diskfull": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogMemorySettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogMemorySetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogMemorySetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogMemorySetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogMemorySetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogMemorySetting")
	}

	return resourceLogMemorySettingRead(d, m)
}

func resourceLogMemorySettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogMemorySetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogMemorySetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogMemorySetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing LogMemorySetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogMemorySettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogMemorySetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogMemorySetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogMemorySetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogMemorySetting resource from API: %v", err)
	}
	return nil
}

func flattenLogMemorySettingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemorySettingDiskfull(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogMemorySetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenLogMemorySettingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("diskfull", flattenLogMemorySettingDiskfull(o["diskfull"], d, "diskfull", sv)); err != nil {
		if !fortiAPIPatch(o["diskfull"]) {
			return fmt.Errorf("Error reading diskfull: %v", err)
		}
	}

	return nil
}

func flattenLogMemorySettingFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandLogMemorySettingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemorySettingDiskfull(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogMemorySetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandLogMemorySettingStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("diskfull"); ok {
		if setArgNil {
			obj["diskfull"] = nil
		} else {

			t, err := expandLogMemorySettingDiskfull(d, v, "diskfull", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["diskfull"] = t
			}
		}
	}

	return &obj, nil
}
