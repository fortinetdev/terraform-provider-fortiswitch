// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: OSPF6 interface configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterospf6Interface() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterospf6InterfaceCreate,
		Read:   resourceRouterospf6InterfaceRead,
		Update: resourceRouterospf6InterfaceUpdate,
		Delete: resourceRouterospf6InterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dead_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hello_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"area_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transmit_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retransmit_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cost": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"passive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterospf6InterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterospf6Interface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating Routerospf6Interface resource while getting object: %v", err)
	}

	o, err := c.CreateRouterospf6Interface(obj)

	if err != nil {
		return fmt.Errorf("Error creating Routerospf6Interface resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("Routerospf6Interface")
	}

	return resourceRouterospf6InterfaceRead(d, m)
}

func resourceRouterospf6InterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterospf6Interface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating Routerospf6Interface resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterospf6Interface(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Routerospf6Interface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("Routerospf6Interface")
	}

	return resourceRouterospf6InterfaceRead(d, m)
}

func resourceRouterospf6InterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterospf6Interface(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting Routerospf6Interface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterospf6InterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterospf6Interface(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Routerospf6Interface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterospf6Interface(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading Routerospf6Interface resource from API: %v", err)
	}
	return nil
}

func flattenRouterospf6InterfaceStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospf6InterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospf6InterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospf6InterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospf6InterfaceBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospf6InterfaceAreaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospf6InterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospf6InterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospf6InterfaceCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospf6InterfacePassive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospf6InterfacePriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterospf6Interface(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenRouterospf6InterfaceStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("dead_interval", flattenRouterospf6InterfaceDeadInterval(o["dead-interval"], d, "dead_interval", sv)); err != nil {
		if !fortiAPIPatch(o["dead-interval"]) {
			return fmt.Errorf("Error reading dead_interval: %v", err)
		}
	}

	if err = d.Set("hello_interval", flattenRouterospf6InterfaceHelloInterval(o["hello-interval"], d, "hello_interval", sv)); err != nil {
		if !fortiAPIPatch(o["hello-interval"]) {
			return fmt.Errorf("Error reading hello_interval: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterospf6InterfaceName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterospf6InterfaceBfd(o["bfd"], d, "bfd", sv)); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("area_id", flattenRouterospf6InterfaceAreaId(o["area-id"], d, "area_id", sv)); err != nil {
		if !fortiAPIPatch(o["area-id"]) {
			return fmt.Errorf("Error reading area_id: %v", err)
		}
	}

	if err = d.Set("transmit_delay", flattenRouterospf6InterfaceTransmitDelay(o["transmit-delay"], d, "transmit_delay", sv)); err != nil {
		if !fortiAPIPatch(o["transmit-delay"]) {
			return fmt.Errorf("Error reading transmit_delay: %v", err)
		}
	}

	if err = d.Set("retransmit_interval", flattenRouterospf6InterfaceRetransmitInterval(o["retransmit-interval"], d, "retransmit_interval", sv)); err != nil {
		if !fortiAPIPatch(o["retransmit-interval"]) {
			return fmt.Errorf("Error reading retransmit_interval: %v", err)
		}
	}

	if err = d.Set("cost", flattenRouterospf6InterfaceCost(o["cost"], d, "cost", sv)); err != nil {
		if !fortiAPIPatch(o["cost"]) {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	if err = d.Set("passive", flattenRouterospf6InterfacePassive(o["passive"], d, "passive", sv)); err != nil {
		if !fortiAPIPatch(o["passive"]) {
			return fmt.Errorf("Error reading passive: %v", err)
		}
	}

	if err = d.Set("priority", flattenRouterospf6InterfacePriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	return nil
}

func flattenRouterospf6InterfaceFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandRouterospf6InterfaceStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospf6InterfaceDeadInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospf6InterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospf6InterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospf6InterfaceBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospf6InterfaceAreaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospf6InterfaceTransmitDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospf6InterfaceRetransmitInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospf6InterfaceCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospf6InterfacePassive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospf6InterfacePriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterospf6Interface(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {

		t, err := expandRouterospf6InterfaceStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("dead_interval"); ok {

		t, err := expandRouterospf6InterfaceDeadInterval(d, v, "dead_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dead-interval"] = t
		}
	}

	if v, ok := d.GetOk("hello_interval"); ok {

		t, err := expandRouterospf6InterfaceHelloInterval(d, v, "hello_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-interval"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandRouterospf6InterfaceName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok {

		t, err := expandRouterospf6InterfaceBfd(d, v, "bfd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("area_id"); ok {

		t, err := expandRouterospf6InterfaceAreaId(d, v, "area_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["area-id"] = t
		}
	}

	if v, ok := d.GetOk("transmit_delay"); ok {

		t, err := expandRouterospf6InterfaceTransmitDelay(d, v, "transmit_delay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transmit-delay"] = t
		}
	}

	if v, ok := d.GetOk("retransmit_interval"); ok {

		t, err := expandRouterospf6InterfaceRetransmitInterval(d, v, "retransmit_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retransmit-interval"] = t
		}
	}

	if v, ok := d.GetOk("cost"); ok {

		t, err := expandRouterospf6InterfaceCost(d, v, "cost", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	if v, ok := d.GetOk("passive"); ok {

		t, err := expandRouterospf6InterfacePassive(d, v, "passive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {

		t, err := expandRouterospf6InterfacePriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	return &obj, nil
}
