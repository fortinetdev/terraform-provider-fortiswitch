// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: QOS 802.1p configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchQosDot1PMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchQosDot1PMapCreate,
		Read:   resourceSwitchQosDot1PMapRead,
		Update: resourceSwitchQosDot1PMapUpdate,
		Delete: resourceSwitchQosDot1PMapDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"priority_5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_0": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"egress_pri_tagging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchQosDot1PMapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchQosDot1PMap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchQosDot1PMap resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchQosDot1PMap(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchQosDot1PMap resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchQosDot1PMap")
	}

	return resourceSwitchQosDot1PMapRead(d, m)
}

func resourceSwitchQosDot1PMapUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchQosDot1PMap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchQosDot1PMap resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchQosDot1PMap(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchQosDot1PMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchQosDot1PMap")
	}

	return resourceSwitchQosDot1PMapRead(d, m)
}

func resourceSwitchQosDot1PMapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchQosDot1PMap(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchQosDot1PMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchQosDot1PMapRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchQosDot1PMap(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchQosDot1PMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchQosDot1PMap(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchQosDot1PMap resource from API: %v", err)
	}
	return nil
}

func flattenSwitchQosDot1PMapName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosDot1PMapPriority5(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosDot1PMapPriority4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosDot1PMapPriority7(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosDot1PMapPriority6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosDot1PMapPriority1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosDot1PMapPriority0(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosDot1PMapPriority3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosDot1PMapPriority2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosDot1PMapEgressPriTagging(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosDot1PMapDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchQosDot1PMap(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchQosDot1PMapName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("priority_5", flattenSwitchQosDot1PMapPriority5(o["priority-5"], d, "priority_5", sv)); err != nil {
		if !fortiAPIPatch(o["priority-5"]) {
			return fmt.Errorf("Error reading priority_5: %v", err)
		}
	}

	if err = d.Set("priority_4", flattenSwitchQosDot1PMapPriority4(o["priority-4"], d, "priority_4", sv)); err != nil {
		if !fortiAPIPatch(o["priority-4"]) {
			return fmt.Errorf("Error reading priority_4: %v", err)
		}
	}

	if err = d.Set("priority_7", flattenSwitchQosDot1PMapPriority7(o["priority-7"], d, "priority_7", sv)); err != nil {
		if !fortiAPIPatch(o["priority-7"]) {
			return fmt.Errorf("Error reading priority_7: %v", err)
		}
	}

	if err = d.Set("priority_6", flattenSwitchQosDot1PMapPriority6(o["priority-6"], d, "priority_6", sv)); err != nil {
		if !fortiAPIPatch(o["priority-6"]) {
			return fmt.Errorf("Error reading priority_6: %v", err)
		}
	}

	if err = d.Set("priority_1", flattenSwitchQosDot1PMapPriority1(o["priority-1"], d, "priority_1", sv)); err != nil {
		if !fortiAPIPatch(o["priority-1"]) {
			return fmt.Errorf("Error reading priority_1: %v", err)
		}
	}

	if err = d.Set("priority_0", flattenSwitchQosDot1PMapPriority0(o["priority-0"], d, "priority_0", sv)); err != nil {
		if !fortiAPIPatch(o["priority-0"]) {
			return fmt.Errorf("Error reading priority_0: %v", err)
		}
	}

	if err = d.Set("priority_3", flattenSwitchQosDot1PMapPriority3(o["priority-3"], d, "priority_3", sv)); err != nil {
		if !fortiAPIPatch(o["priority-3"]) {
			return fmt.Errorf("Error reading priority_3: %v", err)
		}
	}

	if err = d.Set("priority_2", flattenSwitchQosDot1PMapPriority2(o["priority-2"], d, "priority_2", sv)); err != nil {
		if !fortiAPIPatch(o["priority-2"]) {
			return fmt.Errorf("Error reading priority_2: %v", err)
		}
	}

	if err = d.Set("egress_pri_tagging", flattenSwitchQosDot1PMapEgressPriTagging(o["egress-pri-tagging"], d, "egress_pri_tagging", sv)); err != nil {
		if !fortiAPIPatch(o["egress-pri-tagging"]) {
			return fmt.Errorf("Error reading egress_pri_tagging: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchQosDot1PMapDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	return nil
}

func flattenSwitchQosDot1PMapFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchQosDot1PMapName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosDot1PMapPriority5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosDot1PMapPriority4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosDot1PMapPriority7(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosDot1PMapPriority6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosDot1PMapPriority1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosDot1PMapPriority0(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosDot1PMapPriority3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosDot1PMapPriority2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosDot1PMapEgressPriTagging(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosDot1PMapDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchQosDot1PMap(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchQosDot1PMapName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("priority_5"); ok {

		t, err := expandSwitchQosDot1PMapPriority5(d, v, "priority_5", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-5"] = t
		}
	}

	if v, ok := d.GetOk("priority_4"); ok {

		t, err := expandSwitchQosDot1PMapPriority4(d, v, "priority_4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-4"] = t
		}
	}

	if v, ok := d.GetOk("priority_7"); ok {

		t, err := expandSwitchQosDot1PMapPriority7(d, v, "priority_7", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-7"] = t
		}
	}

	if v, ok := d.GetOk("priority_6"); ok {

		t, err := expandSwitchQosDot1PMapPriority6(d, v, "priority_6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-6"] = t
		}
	}

	if v, ok := d.GetOk("priority_1"); ok {

		t, err := expandSwitchQosDot1PMapPriority1(d, v, "priority_1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-1"] = t
		}
	}

	if v, ok := d.GetOk("priority_0"); ok {

		t, err := expandSwitchQosDot1PMapPriority0(d, v, "priority_0", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-0"] = t
		}
	}

	if v, ok := d.GetOk("priority_3"); ok {

		t, err := expandSwitchQosDot1PMapPriority3(d, v, "priority_3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-3"] = t
		}
	}

	if v, ok := d.GetOk("priority_2"); ok {

		t, err := expandSwitchQosDot1PMapPriority2(d, v, "priority_2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-2"] = t
		}
	}

	if v, ok := d.GetOk("egress_pri_tagging"); ok {

		t, err := expandSwitchQosDot1PMapEgressPriTagging(d, v, "egress_pri_tagging", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["egress-pri-tagging"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchQosDot1PMapDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	return &obj, nil
}
