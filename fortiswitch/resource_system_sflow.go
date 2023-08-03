// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure sFlow.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSflow() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSflowUpdate,
		Read:   resourceSystemSflowRead,
		Update: resourceSystemSflowUpdate,
		Delete: resourceSystemSflowDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"dummy": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"collectors": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemSflowUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSflow(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSflow resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSflow(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemSflow resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSflow")
	}

	return resourceSystemSflowRead(d, m)
}

func resourceSystemSflowDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSflow(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemSflow resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSflow(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemSflow resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSflowRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemSflow(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemSflow resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSflow(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSflow resource from API: %v", err)
	}
	return nil
}

func flattenSystemSflowDummy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSflowCollectors(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenSystemSflowCollectorsIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemSflowCollectorsName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenSystemSflowCollectorsPort(i["port"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSflowCollectorsIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSflowCollectorsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSflowCollectorsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSflow(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("dummy", flattenSystemSflowDummy(o["dummy"], d, "dummy", sv)); err != nil {
		if !fortiAPIPatch(o["dummy"]) {
			return fmt.Errorf("Error reading dummy: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("collectors", flattenSystemSflowCollectors(o["collectors"], d, "collectors", sv)); err != nil {
			if !fortiAPIPatch(o["collectors"]) {
				return fmt.Errorf("Error reading collectors: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("collectors"); ok {
			if err = d.Set("collectors", flattenSystemSflowCollectors(o["collectors"], d, "collectors", sv)); err != nil {
				if !fortiAPIPatch(o["collectors"]) {
					return fmt.Errorf("Error reading collectors: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemSflowFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemSflowDummy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSflowCollectors(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandSystemSflowCollectorsIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemSflowCollectorsName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandSystemSflowCollectorsPort(d, i["port"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSflowCollectorsIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSflowCollectorsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSflowCollectorsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSflow(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dummy"); ok {
		if setArgNil {
			obj["dummy"] = nil
		} else {

			t, err := expandSystemSflowDummy(d, v, "dummy", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dummy"] = t
			}
		}
	}

	if v, ok := d.GetOk("collectors"); ok || d.HasChange("collectors") {
		if setArgNil {
			obj["collectors"] = make([]struct{}, 0)
		} else {

			t, err := expandSystemSflowCollectors(d, v, "collectors", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["collectors"] = t
			}
		}
	}

	return &obj, nil
}
