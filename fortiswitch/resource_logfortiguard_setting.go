// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Settings for FortiGuard Analysis Service.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogFortiguardSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortiguardSettingUpdate,
		Read:   resourceLogFortiguardSettingRead,
		Update: resourceLogFortiguardSettingUpdate,
		Delete: resourceLogFortiguardSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogFortiguardSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogFortiguardSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogFortiguardSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogFortiguardSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogFortiguardSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogFortiguardSetting")
	}

	return resourceLogFortiguardSettingRead(d, m)
}

func resourceLogFortiguardSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogFortiguardSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogFortiguardSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogFortiguardSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing LogFortiguardSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortiguardSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogFortiguardSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogFortiguardSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortiguardSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogFortiguardSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogFortiguardSettingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortiguardSettingUploadInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortiguardSettingUploadTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortiguardSettingUploadDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortiguardSettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogFortiguardSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenLogFortiguardSettingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upload_interval", flattenLogFortiguardSettingUploadInterval(o["upload-interval"], d, "upload_interval", sv)); err != nil {
		if !fortiAPIPatch(o["upload-interval"]) {
			return fmt.Errorf("Error reading upload_interval: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenLogFortiguardSettingUploadTime(o["upload-time"], d, "upload_time", sv)); err != nil {
		if !fortiAPIPatch(o["upload-time"]) {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	if err = d.Set("upload_day", flattenLogFortiguardSettingUploadDay(o["upload-day"], d, "upload_day", sv)); err != nil {
		if !fortiAPIPatch(o["upload-day"]) {
			return fmt.Errorf("Error reading upload_day: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenLogFortiguardSettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	return nil
}

func flattenLogFortiguardSettingFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandLogFortiguardSettingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingUploadInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingUploadTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingUploadDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortiguardSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandLogFortiguardSettingStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_interval"); ok {
		if setArgNil {
			obj["upload-interval"] = nil
		} else {

			t, err := expandLogFortiguardSettingUploadInterval(d, v, "upload_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_time"); ok {
		if setArgNil {
			obj["upload-time"] = nil
		} else {

			t, err := expandLogFortiguardSettingUploadTime(d, v, "upload_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_day"); ok {
		if setArgNil {
			obj["upload-day"] = nil
		} else {

			t, err := expandLogFortiguardSettingUploadDay(d, v, "upload_day", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-day"] = t
			}
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {
		if setArgNil {
			obj["enc-algorithm"] = nil
		} else {

			t, err := expandLogFortiguardSettingEncAlgorithm(d, v, "enc_algorithm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["enc-algorithm"] = t
			}
		}
	}

	return &obj, nil
}
