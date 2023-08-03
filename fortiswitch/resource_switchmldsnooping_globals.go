// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure mld-snooping on Switch.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchMldSnoopingGlobals() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchMldSnoopingGlobalsUpdate,
		Read:   resourceSwitchMldSnoopingGlobalsRead,
		Update: resourceSwitchMldSnoopingGlobalsUpdate,
		Delete: resourceSwitchMldSnoopingGlobalsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"query_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 1200),
				Optional:     true,
				Computed:     true,
			},
			"proxy_report_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 260),
				Optional:     true,
				Computed:     true,
			},
			"query_max_response_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(100, 32768),
				Optional:     true,
				Computed:     true,
			},
			"leave_response_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 5000),
				Optional:     true,
				Computed:     true,
			},
			"aging_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(15, 3600),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchMldSnoopingGlobalsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchMldSnoopingGlobals(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchMldSnoopingGlobals resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchMldSnoopingGlobals(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchMldSnoopingGlobals resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchMldSnoopingGlobals")
	}

	return resourceSwitchMldSnoopingGlobalsRead(d, m)
}

func resourceSwitchMldSnoopingGlobalsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchMldSnoopingGlobals(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchMldSnoopingGlobals resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchMldSnoopingGlobals(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchMldSnoopingGlobals resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchMldSnoopingGlobalsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchMldSnoopingGlobals(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchMldSnoopingGlobals resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchMldSnoopingGlobals(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchMldSnoopingGlobals resource from API: %v", err)
	}
	return nil
}

func flattenSwitchMldSnoopingGlobalsQueryInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMldSnoopingGlobalsProxyReportInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMldSnoopingGlobalsQueryMaxResponseTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMldSnoopingGlobalsLeaveResponseTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMldSnoopingGlobalsAgingTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchMldSnoopingGlobals(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("query_interval", flattenSwitchMldSnoopingGlobalsQueryInterval(o["query-interval"], d, "query_interval", sv)); err != nil {
		if !fortiAPIPatch(o["query-interval"]) {
			return fmt.Errorf("Error reading query_interval: %v", err)
		}
	}

	if err = d.Set("proxy_report_interval", flattenSwitchMldSnoopingGlobalsProxyReportInterval(o["proxy-report-interval"], d, "proxy_report_interval", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-report-interval"]) {
			return fmt.Errorf("Error reading proxy_report_interval: %v", err)
		}
	}

	if err = d.Set("query_max_response_timeout", flattenSwitchMldSnoopingGlobalsQueryMaxResponseTimeout(o["query-max-response-timeout"], d, "query_max_response_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["query-max-response-timeout"]) {
			return fmt.Errorf("Error reading query_max_response_timeout: %v", err)
		}
	}

	if err = d.Set("leave_response_timeout", flattenSwitchMldSnoopingGlobalsLeaveResponseTimeout(o["leave-response-timeout"], d, "leave_response_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["leave-response-timeout"]) {
			return fmt.Errorf("Error reading leave_response_timeout: %v", err)
		}
	}

	if err = d.Set("aging_time", flattenSwitchMldSnoopingGlobalsAgingTime(o["aging-time"], d, "aging_time", sv)); err != nil {
		if !fortiAPIPatch(o["aging-time"]) {
			return fmt.Errorf("Error reading aging_time: %v", err)
		}
	}

	return nil
}

func flattenSwitchMldSnoopingGlobalsFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchMldSnoopingGlobalsQueryInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMldSnoopingGlobalsProxyReportInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMldSnoopingGlobalsQueryMaxResponseTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMldSnoopingGlobalsLeaveResponseTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMldSnoopingGlobalsAgingTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchMldSnoopingGlobals(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("query_interval"); ok {
		if setArgNil {
			obj["query-interval"] = nil
		} else {

			t, err := expandSwitchMldSnoopingGlobalsQueryInterval(d, v, "query_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["query-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("proxy_report_interval"); ok {
		if setArgNil {
			obj["proxy-report-interval"] = nil
		} else {

			t, err := expandSwitchMldSnoopingGlobalsProxyReportInterval(d, v, "proxy_report_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["proxy-report-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("query_max_response_timeout"); ok {
		if setArgNil {
			obj["query-max-response-timeout"] = nil
		} else {

			t, err := expandSwitchMldSnoopingGlobalsQueryMaxResponseTimeout(d, v, "query_max_response_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["query-max-response-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("leave_response_timeout"); ok {
		if setArgNil {
			obj["leave-response-timeout"] = nil
		} else {

			t, err := expandSwitchMldSnoopingGlobalsLeaveResponseTimeout(d, v, "leave_response_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["leave-response-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("aging_time"); ok {
		if setArgNil {
			obj["aging-time"] = nil
		} else {

			t, err := expandSwitchMldSnoopingGlobalsAgingTime(d, v, "aging_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["aging-time"] = t
			}
		}
	}

	return &obj, nil
}
