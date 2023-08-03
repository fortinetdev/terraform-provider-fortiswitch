// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure excess switch traffic (storm control).

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchStormControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchStormControlUpdate,
		Read:   resourceSwitchStormControlRead,
		Update: resourceSwitchStormControlUpdate,
		Delete: resourceSwitchStormControlDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"broadcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_unicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_multicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"burst_size_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchStormControlUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchStormControl(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchStormControl resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchStormControl(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchStormControl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchStormControl")
	}

	return resourceSwitchStormControlRead(d, m)
}

func resourceSwitchStormControlDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchStormControl(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchStormControl resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchStormControl(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchStormControl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchStormControlRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchStormControl(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchStormControl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchStormControl(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchStormControl resource from API: %v", err)
	}
	return nil
}

func flattenSwitchStormControlBroadcast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStormControlUnknownUnicast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStormControlUnknownMulticast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStormControlRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStormControlBurstSizeLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchStormControl(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("broadcast", flattenSwitchStormControlBroadcast(o["broadcast"], d, "broadcast", sv)); err != nil {
		if !fortiAPIPatch(o["broadcast"]) {
			return fmt.Errorf("Error reading broadcast: %v", err)
		}
	}

	if err = d.Set("unknown_unicast", flattenSwitchStormControlUnknownUnicast(o["unknown-unicast"], d, "unknown_unicast", sv)); err != nil {
		if !fortiAPIPatch(o["unknown-unicast"]) {
			return fmt.Errorf("Error reading unknown_unicast: %v", err)
		}
	}

	if err = d.Set("unknown_multicast", flattenSwitchStormControlUnknownMulticast(o["unknown-multicast"], d, "unknown_multicast", sv)); err != nil {
		if !fortiAPIPatch(o["unknown-multicast"]) {
			return fmt.Errorf("Error reading unknown_multicast: %v", err)
		}
	}

	if err = d.Set("rate", flattenSwitchStormControlRate(o["rate"], d, "rate", sv)); err != nil {
		if !fortiAPIPatch(o["rate"]) {
			return fmt.Errorf("Error reading rate: %v", err)
		}
	}

	if err = d.Set("burst_size_level", flattenSwitchStormControlBurstSizeLevel(o["burst-size-level"], d, "burst_size_level", sv)); err != nil {
		if !fortiAPIPatch(o["burst-size-level"]) {
			return fmt.Errorf("Error reading burst_size_level: %v", err)
		}
	}

	return nil
}

func flattenSwitchStormControlFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchStormControlBroadcast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStormControlUnknownUnicast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStormControlUnknownMulticast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStormControlRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStormControlBurstSizeLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchStormControl(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("broadcast"); ok {
		if setArgNil {
			obj["broadcast"] = nil
		} else {

			t, err := expandSwitchStormControlBroadcast(d, v, "broadcast", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["broadcast"] = t
			}
		}
	}

	if v, ok := d.GetOk("unknown_unicast"); ok {
		if setArgNil {
			obj["unknown-unicast"] = nil
		} else {

			t, err := expandSwitchStormControlUnknownUnicast(d, v, "unknown_unicast", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["unknown-unicast"] = t
			}
		}
	}

	if v, ok := d.GetOk("unknown_multicast"); ok {
		if setArgNil {
			obj["unknown-multicast"] = nil
		} else {

			t, err := expandSwitchStormControlUnknownMulticast(d, v, "unknown_multicast", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["unknown-multicast"] = t
			}
		}
	}

	if v, ok := d.GetOk("rate"); ok {
		if setArgNil {
			obj["rate"] = nil
		} else {

			t, err := expandSwitchStormControlRate(d, v, "rate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rate"] = t
			}
		}
	}

	if v, ok := d.GetOk("burst_size_level"); ok {
		if setArgNil {
			obj["burst-size-level"] = nil
		} else {

			t, err := expandSwitchStormControlBurstSizeLevel(d, v, "burst_size_level", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["burst-size-level"] = t
			}
		}
	}

	return &obj, nil
}
