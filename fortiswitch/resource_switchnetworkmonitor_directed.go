// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configuration of the static entries for network monitoring on the switch.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchNetworkMonitorDirected() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchNetworkMonitorDirectedCreate,
		Read:   resourceSwitchNetworkMonitorDirectedRead,
		Update: resourceSwitchNetworkMonitorDirectedUpdate,
		Delete: resourceSwitchNetworkMonitorDirectedDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fswid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"monitor_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchNetworkMonitorDirectedCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchNetworkMonitorDirected(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchNetworkMonitorDirected resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchNetworkMonitorDirected(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchNetworkMonitorDirected resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchNetworkMonitorDirected")
	}

	return resourceSwitchNetworkMonitorDirectedRead(d, m)
}

func resourceSwitchNetworkMonitorDirectedUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchNetworkMonitorDirected(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchNetworkMonitorDirected resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchNetworkMonitorDirected(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchNetworkMonitorDirected resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchNetworkMonitorDirected")
	}

	return resourceSwitchNetworkMonitorDirectedRead(d, m)
}

func resourceSwitchNetworkMonitorDirectedDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchNetworkMonitorDirected(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchNetworkMonitorDirected resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchNetworkMonitorDirectedRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchNetworkMonitorDirected(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchNetworkMonitorDirected resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchNetworkMonitorDirected(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchNetworkMonitorDirected resource from API: %v", err)
	}
	return nil
}

func flattenSwitchNetworkMonitorDirectedId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchNetworkMonitorDirectedMonitorMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchNetworkMonitorDirected(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fswid", flattenSwitchNetworkMonitorDirectedId(o["id"], d, "fswid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fswid: %v", err)
		}
	}

	if err = d.Set("monitor_mac", flattenSwitchNetworkMonitorDirectedMonitorMac(o["monitor-mac"], d, "monitor_mac", sv)); err != nil {
		if !fortiAPIPatch(o["monitor-mac"]) {
			return fmt.Errorf("Error reading monitor_mac: %v", err)
		}
	}

	return nil
}

func flattenSwitchNetworkMonitorDirectedFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchNetworkMonitorDirectedId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchNetworkMonitorDirectedMonitorMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchNetworkMonitorDirected(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fswid"); ok {

		t, err := expandSwitchNetworkMonitorDirectedId(d, v, "fswid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("monitor_mac"); ok {

		t, err := expandSwitchNetworkMonitorDirectedMonitorMac(d, v, "monitor_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-mac"] = t
		}
	}

	return &obj, nil
}
