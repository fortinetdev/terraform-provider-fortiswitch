// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Stp instances.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchStpInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchStpInstanceCreate,
		Read:   resourceSwitchStpInstanceRead,
		Update: resourceSwitchStpInstanceUpdate,
		Delete: resourceSwitchStpInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stp_port": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"fswid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"vlan_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchStpInstanceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchStpInstance(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchStpInstance resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchStpInstance(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchStpInstance resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchStpInstance")
	}

	return resourceSwitchStpInstanceRead(d, m)
}

func resourceSwitchStpInstanceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchStpInstance(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchStpInstance resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchStpInstance(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchStpInstance resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchStpInstance")
	}

	return resourceSwitchStpInstanceRead(d, m)
}

func resourceSwitchStpInstanceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchStpInstance(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchStpInstance resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchStpInstanceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchStpInstance(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchStpInstance resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchStpInstance(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchStpInstance resource from API: %v", err)
	}
	return nil
}

func flattenSwitchStpInstancePriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpInstanceStpPort(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {

			tmp["priority"] = flattenSwitchStpInstanceStpPortPriority(i["priority"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {

			tmp["cost"] = flattenSwitchStpInstanceStpPortCost(i["cost"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchStpInstanceStpPortName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchStpInstanceStpPortPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpInstanceStpPortCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpInstanceStpPortName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpInstanceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchStpInstanceVlanRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchStpInstance(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("priority", flattenSwitchStpInstancePriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("stp_port", flattenSwitchStpInstanceStpPort(o["stp-port"], d, "stp_port", sv)); err != nil {
			if !fortiAPIPatch(o["stp-port"]) {
				return fmt.Errorf("Error reading stp_port: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("stp_port"); ok {
			if err = d.Set("stp_port", flattenSwitchStpInstanceStpPort(o["stp-port"], d, "stp_port", sv)); err != nil {
				if !fortiAPIPatch(o["stp-port"]) {
					return fmt.Errorf("Error reading stp_port: %v", err)
				}
			}
		}
	}

	if err = d.Set("fswid", flattenSwitchStpInstanceId(o["id"], d, "fswid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fswid: %v", err)
		}
	}

	if err = d.Set("vlan_range", flattenSwitchStpInstanceVlanRange(o["vlan-range"], d, "vlan_range", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-range"]) {
			return fmt.Errorf("Error reading vlan_range: %v", err)
		}
	}

	return nil
}

func flattenSwitchStpInstanceFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchStpInstancePriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpInstanceStpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["priority"], _ = expandSwitchStpInstanceStpPortPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cost"], _ = expandSwitchStpInstanceStpPortCost(d, i["cost"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSwitchStpInstanceStpPortName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchStpInstanceStpPortPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpInstanceStpPortCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpInstanceStpPortName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpInstanceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchStpInstanceVlanRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchStpInstance(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("priority"); ok {

		t, err := expandSwitchStpInstancePriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("stp_port"); ok || d.HasChange("stp_port") {

		t, err := expandSwitchStpInstanceStpPort(d, v, "stp_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-port"] = t
		}
	}

	if v, ok := d.GetOk("fswid"); ok {

		t, err := expandSwitchStpInstanceId(d, v, "fswid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("vlan_range"); ok {

		t, err := expandSwitchStpInstanceVlanRange(d, v, "vlan_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-range"] = t
		}
	}

	return &obj, nil
}
