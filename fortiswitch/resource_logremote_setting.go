// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Settings for remote logging.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogRemoteSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogRemoteSettingUpdate,
		Read:   resourceLogRemoteSettingRead,
		Update: resourceLogRemoteSettingUpdate,
		Delete: resourceLogRemoteSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogRemoteSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogRemoteSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogRemoteSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogRemoteSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogRemoteSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogRemoteSetting")
	}

	return resourceLogRemoteSettingRead(d, m)
}

func resourceLogRemoteSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogRemoteSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogRemoteSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogRemoteSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing LogRemoteSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogRemoteSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogRemoteSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogRemoteSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogRemoteSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogRemoteSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogRemoteSettingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogRemoteSettingDestination(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogRemoteSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenLogRemoteSettingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("destination", flattenLogRemoteSettingDestination(o["destination"], d, "destination", sv)); err != nil {
		if !fortiAPIPatch(o["destination"]) {
			return fmt.Errorf("Error reading destination: %v", err)
		}
	}

	return nil
}

func flattenLogRemoteSettingFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandLogRemoteSettingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogRemoteSettingDestination(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogRemoteSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandLogRemoteSettingStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("destination"); ok {
		if setArgNil {
			obj["destination"] = nil
		} else {

			t, err := expandLogRemoteSettingDestination(d, v, "destination", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["destination"] = t
			}
		}
	}

	return &obj, nil
}
