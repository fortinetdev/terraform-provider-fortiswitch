// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Filters for memory buffer.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogMemoryFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogMemoryFilterUpdate,
		Read:   resourceLogMemoryFilterRead,
		Update: resourceLogMemoryFilterUpdate,
		Delete: resourceLogMemoryFilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wireless_activity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpu_memory_usage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogMemoryFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogMemoryFilter(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogMemoryFilter resource while getting object: %v", err)
	}

	o, err := c.UpdateLogMemoryFilter(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogMemoryFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogMemoryFilter")
	}

	return resourceLogMemoryFilterRead(d, m)
}

func resourceLogMemoryFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogMemoryFilter(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogMemoryFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogMemoryFilter(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing LogMemoryFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogMemoryFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogMemoryFilter(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogMemoryFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogMemoryFilter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogMemoryFilter resource from API: %v", err)
	}
	return nil
}

func flattenLogMemoryFilterSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterWirelessActivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterSystem(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterCpuMemoryUsage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterDhcp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterHa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterEvent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogMemoryFilter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("severity", flattenLogMemoryFilterSeverity(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("admin", flattenLogMemoryFilterAdmin(o["admin"], d, "admin", sv)); err != nil {
		if !fortiAPIPatch(o["admin"]) {
			return fmt.Errorf("Error reading admin: %v", err)
		}
	}

	if err = d.Set("pattern", flattenLogMemoryFilterPattern(o["pattern"], d, "pattern", sv)); err != nil {
		if !fortiAPIPatch(o["pattern"]) {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("wireless_activity", flattenLogMemoryFilterWirelessActivity(o["wireless-activity"], d, "wireless_activity", sv)); err != nil {
		if !fortiAPIPatch(o["wireless-activity"]) {
			return fmt.Errorf("Error reading wireless_activity: %v", err)
		}
	}

	if err = d.Set("system", flattenLogMemoryFilterSystem(o["system"], d, "system", sv)); err != nil {
		if !fortiAPIPatch(o["system"]) {
			return fmt.Errorf("Error reading system: %v", err)
		}
	}

	if err = d.Set("auth", flattenLogMemoryFilterAuth(o["auth"], d, "auth", sv)); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("cpu_memory_usage", flattenLogMemoryFilterCpuMemoryUsage(o["cpu-memory-usage"], d, "cpu_memory_usage", sv)); err != nil {
		if !fortiAPIPatch(o["cpu-memory-usage"]) {
			return fmt.Errorf("Error reading cpu_memory_usage: %v", err)
		}
	}

	if err = d.Set("override", flattenLogMemoryFilterOverride(o["override"], d, "override", sv)); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("dhcp", flattenLogMemoryFilterDhcp(o["dhcp"], d, "dhcp", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp"]) {
			return fmt.Errorf("Error reading dhcp: %v", err)
		}
	}

	if err = d.Set("ha", flattenLogMemoryFilterHa(o["ha"], d, "ha", sv)); err != nil {
		if !fortiAPIPatch(o["ha"]) {
			return fmt.Errorf("Error reading ha: %v", err)
		}
	}

	if err = d.Set("event", flattenLogMemoryFilterEvent(o["event"], d, "event", sv)); err != nil {
		if !fortiAPIPatch(o["event"]) {
			return fmt.Errorf("Error reading event: %v", err)
		}
	}

	return nil
}

func flattenLogMemoryFilterFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandLogMemoryFilterSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterWirelessActivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterSystem(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterCpuMemoryUsage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterDhcp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterHa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterEvent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogMemoryFilter(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("severity"); ok {
		if setArgNil {
			obj["severity"] = nil
		} else {

			t, err := expandLogMemoryFilterSeverity(d, v, "severity", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["severity"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin"); ok {
		if setArgNil {
			obj["admin"] = nil
		} else {

			t, err := expandLogMemoryFilterAdmin(d, v, "admin", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin"] = t
			}
		}
	}

	if v, ok := d.GetOk("pattern"); ok {
		if setArgNil {
			obj["pattern"] = nil
		} else {

			t, err := expandLogMemoryFilterPattern(d, v, "pattern", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pattern"] = t
			}
		}
	}

	if v, ok := d.GetOk("wireless_activity"); ok {
		if setArgNil {
			obj["wireless-activity"] = nil
		} else {

			t, err := expandLogMemoryFilterWirelessActivity(d, v, "wireless_activity", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wireless-activity"] = t
			}
		}
	}

	if v, ok := d.GetOk("system"); ok {
		if setArgNil {
			obj["system"] = nil
		} else {

			t, err := expandLogMemoryFilterSystem(d, v, "system", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["system"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth"); ok {
		if setArgNil {
			obj["auth"] = nil
		} else {

			t, err := expandLogMemoryFilterAuth(d, v, "auth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth"] = t
			}
		}
	}

	if v, ok := d.GetOk("cpu_memory_usage"); ok {
		if setArgNil {
			obj["cpu-memory-usage"] = nil
		} else {

			t, err := expandLogMemoryFilterCpuMemoryUsage(d, v, "cpu_memory_usage", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cpu-memory-usage"] = t
			}
		}
	}

	if v, ok := d.GetOk("override"); ok {
		if setArgNil {
			obj["override"] = nil
		} else {

			t, err := expandLogMemoryFilterOverride(d, v, "override", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["override"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp"); ok {
		if setArgNil {
			obj["dhcp"] = nil
		} else {

			t, err := expandLogMemoryFilterDhcp(d, v, "dhcp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp"] = t
			}
		}
	}

	if v, ok := d.GetOk("ha"); ok {
		if setArgNil {
			obj["ha"] = nil
		} else {

			t, err := expandLogMemoryFilterHa(d, v, "ha", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ha"] = t
			}
		}
	}

	if v, ok := d.GetOk("event"); ok {
		if setArgNil {
			obj["event"] = nil
		} else {

			t, err := expandLogMemoryFilterEvent(d, v, "event", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["event"] = t
			}
		}
	}

	return &obj, nil
}
