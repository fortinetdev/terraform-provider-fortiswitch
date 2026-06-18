// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: IP-source-guard table.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchIpSourceGuard() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchIpSourceGuardCreate,
		Read:   resourceSwitchIpSourceGuardRead,
		Update: resourceSwitchIpSourceGuardUpdate,
		Delete: resourceSwitchIpSourceGuardDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"interface_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"binding_entry": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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

func resourceSwitchIpSourceGuardCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchIpSourceGuard(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchIpSourceGuard resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchIpSourceGuard(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchIpSourceGuard resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchIpSourceGuard")
	}

	return resourceSwitchIpSourceGuardRead(d, m)
}

func resourceSwitchIpSourceGuardUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchIpSourceGuard(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchIpSourceGuard resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchIpSourceGuard(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchIpSourceGuard resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchIpSourceGuard")
	}

	return resourceSwitchIpSourceGuardRead(d, m)
}

func resourceSwitchIpSourceGuardDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchIpSourceGuard(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchIpSourceGuard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchIpSourceGuardRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchIpSourceGuard(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchIpSourceGuard resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchIpSourceGuard(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchIpSourceGuard resource from API: %v", err)
	}
	return nil
}

func flattenSwitchIpSourceGuardInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchIpSourceGuardDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchIpSourceGuardBindingEntry(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_name"
		if _, ok := i["entry-name"]; ok {

			tmp["entry_name"] = flattenSwitchIpSourceGuardBindingEntryEntryName(i["entry-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenSwitchIpSourceGuardBindingEntryIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {

			tmp["mac"] = flattenSwitchIpSourceGuardBindingEntryMac(i["mac"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "entry_name", d)
	return result
}

func flattenSwitchIpSourceGuardBindingEntryEntryName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchIpSourceGuardBindingEntryIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchIpSourceGuardBindingEntryMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchIpSourceGuard(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("interface_name", flattenSwitchIpSourceGuardInterfaceName(o["interface-name"], d, "interface_name", sv)); err != nil {
		if !fortiAPIPatch(o["interface-name"]) {
			return fmt.Errorf("Error reading interface_name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchIpSourceGuardDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("binding_entry", flattenSwitchIpSourceGuardBindingEntry(o["binding-entry"], d, "binding_entry", sv)); err != nil {
			if !fortiAPIPatch(o["binding-entry"]) {
				return fmt.Errorf("Error reading binding_entry: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("binding_entry"); ok {
			if err = d.Set("binding_entry", flattenSwitchIpSourceGuardBindingEntry(o["binding-entry"], d, "binding_entry", sv)); err != nil {
				if !fortiAPIPatch(o["binding-entry"]) {
					return fmt.Errorf("Error reading binding_entry: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchIpSourceGuardFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchIpSourceGuardInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchIpSourceGuardDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchIpSourceGuardBindingEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["entry-name"], _ = expandSwitchIpSourceGuardBindingEntryEntryName(d, i["entry_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandSwitchIpSourceGuardBindingEntryIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mac"], _ = expandSwitchIpSourceGuardBindingEntryMac(d, i["mac"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchIpSourceGuardBindingEntryEntryName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchIpSourceGuardBindingEntryIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchIpSourceGuardBindingEntryMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchIpSourceGuard(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("interface_name"); ok {

		t, err := expandSwitchIpSourceGuardInterfaceName(d, v, "interface_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchIpSourceGuardDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("binding_entry"); ok || d.HasChange("binding_entry") {

		t, err := expandSwitchIpSourceGuardBindingEntry(d, v, "binding_entry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["binding-entry"] = t
		}
	}

	return &obj, nil
}
