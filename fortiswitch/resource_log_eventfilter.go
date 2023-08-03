// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Log event filter configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogEventfilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogEventfilterUpdate,
		Read:   resourceLogEventfilterRead,
		Update: resourceLogEventfilterUpdate,
		Delete: resourceLogEventfilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"spanning_tree": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"router": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fos_legacy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"poe": &schema.Schema{
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

func resourceLogEventfilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogEventfilter(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogEventfilter resource while getting object: %v", err)
	}

	o, err := c.UpdateLogEventfilter(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogEventfilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogEventfilter")
	}

	return resourceLogEventfilterRead(d, m)
}

func resourceLogEventfilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogEventfilter(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogEventfilter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogEventfilter(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing LogEventfilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogEventfilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogEventfilter(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogEventfilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogEventfilter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogEventfilter resource from API: %v", err)
	}
	return nil
}

func flattenLogEventfilterSpanning_Tree(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterSystem(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterSwitch_Controller(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterLink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterRouter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterFos_Legacy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterPoe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterEvent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogEventfilter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("spanning_tree", flattenLogEventfilterSpanning_Tree(o["spanning_tree"], d, "spanning_tree", sv)); err != nil {
		if !fortiAPIPatch(o["spanning_tree"]) {
			return fmt.Errorf("Error reading spanning_tree: %v", err)
		}
	}

	if err = d.Set("system", flattenLogEventfilterSystem(o["system"], d, "system", sv)); err != nil {
		if !fortiAPIPatch(o["system"]) {
			return fmt.Errorf("Error reading system: %v", err)
		}
	}

	if err = d.Set("switch_controller", flattenLogEventfilterSwitch_Controller(o["switch_controller"], d, "switch_controller", sv)); err != nil {
		if !fortiAPIPatch(o["switch_controller"]) {
			return fmt.Errorf("Error reading switch_controller: %v", err)
		}
	}

	if err = d.Set("switch", flattenLogEventfilterSwitch(o["switch"], d, "switch", sv)); err != nil {
		if !fortiAPIPatch(o["switch"]) {
			return fmt.Errorf("Error reading switch: %v", err)
		}
	}

	if err = d.Set("link", flattenLogEventfilterLink(o["link"], d, "link", sv)); err != nil {
		if !fortiAPIPatch(o["link"]) {
			return fmt.Errorf("Error reading link: %v", err)
		}
	}

	if err = d.Set("user", flattenLogEventfilterUser(o["user"], d, "user", sv)); err != nil {
		if !fortiAPIPatch(o["user"]) {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("router", flattenLogEventfilterRouter(o["router"], d, "router", sv)); err != nil {
		if !fortiAPIPatch(o["router"]) {
			return fmt.Errorf("Error reading router: %v", err)
		}
	}

	if err = d.Set("fos_legacy", flattenLogEventfilterFos_Legacy(o["fos_legacy"], d, "fos_legacy", sv)); err != nil {
		if !fortiAPIPatch(o["fos_legacy"]) {
			return fmt.Errorf("Error reading fos_legacy: %v", err)
		}
	}

	if err = d.Set("poe", flattenLogEventfilterPoe(o["poe"], d, "poe", sv)); err != nil {
		if !fortiAPIPatch(o["poe"]) {
			return fmt.Errorf("Error reading poe: %v", err)
		}
	}

	if err = d.Set("event", flattenLogEventfilterEvent(o["event"], d, "event", sv)); err != nil {
		if !fortiAPIPatch(o["event"]) {
			return fmt.Errorf("Error reading event: %v", err)
		}
	}

	return nil
}

func flattenLogEventfilterFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandLogEventfilterSpanning_Tree(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterSystem(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterSwitch_Controller(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterRouter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterFos_Legacy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterPoe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterEvent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogEventfilter(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("spanning_tree"); ok {
		if setArgNil {
			obj["spanning_tree"] = nil
		} else {

			t, err := expandLogEventfilterSpanning_Tree(d, v, "spanning_tree", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["spanning_tree"] = t
			}
		}
	}

	if v, ok := d.GetOk("system"); ok {
		if setArgNil {
			obj["system"] = nil
		} else {

			t, err := expandLogEventfilterSystem(d, v, "system", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["system"] = t
			}
		}
	}

	if v, ok := d.GetOk("switch_controller"); ok {
		if setArgNil {
			obj["switch_controller"] = nil
		} else {

			t, err := expandLogEventfilterSwitch_Controller(d, v, "switch_controller", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["switch_controller"] = t
			}
		}
	}

	if v, ok := d.GetOk("switch"); ok {
		if setArgNil {
			obj["switch"] = nil
		} else {

			t, err := expandLogEventfilterSwitch(d, v, "switch", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["switch"] = t
			}
		}
	}

	if v, ok := d.GetOk("link"); ok {
		if setArgNil {
			obj["link"] = nil
		} else {

			t, err := expandLogEventfilterLink(d, v, "link", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["link"] = t
			}
		}
	}

	if v, ok := d.GetOk("user"); ok {
		if setArgNil {
			obj["user"] = nil
		} else {

			t, err := expandLogEventfilterUser(d, v, "user", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["user"] = t
			}
		}
	}

	if v, ok := d.GetOk("router"); ok {
		if setArgNil {
			obj["router"] = nil
		} else {

			t, err := expandLogEventfilterRouter(d, v, "router", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["router"] = t
			}
		}
	}

	if v, ok := d.GetOk("fos_legacy"); ok {
		if setArgNil {
			obj["fos_legacy"] = nil
		} else {

			t, err := expandLogEventfilterFos_Legacy(d, v, "fos_legacy", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fos_legacy"] = t
			}
		}
	}

	if v, ok := d.GetOk("poe"); ok {
		if setArgNil {
			obj["poe"] = nil
		} else {

			t, err := expandLogEventfilterPoe(d, v, "poe", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["poe"] = t
			}
		}
	}

	if v, ok := d.GetOk("event"); ok {
		if setArgNil {
			obj["event"] = nil
		} else {

			t, err := expandLogEventfilterEvent(d, v, "event", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["event"] = t
			}
		}
	}

	return &obj, nil
}
