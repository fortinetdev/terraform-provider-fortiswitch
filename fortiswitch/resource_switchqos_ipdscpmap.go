// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: QOS IP precedence/DSCP configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchQosIpDscpMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchQosIpDscpMapCreate,
		Read:   resourceSwitchQosIpDscpMapRead,
		Update: resourceSwitchQosIpDscpMapUpdate,
		Delete: resourceSwitchQosIpDscpMapDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"map": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"diffserv": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cos_queue": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip_precedence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"entry_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchQosIpDscpMapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchQosIpDscpMap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchQosIpDscpMap resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchQosIpDscpMap(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchQosIpDscpMap resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchQosIpDscpMap")
	}

	return resourceSwitchQosIpDscpMapRead(d, m)
}

func resourceSwitchQosIpDscpMapUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchQosIpDscpMap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchQosIpDscpMap resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchQosIpDscpMap(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchQosIpDscpMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchQosIpDscpMap")
	}

	return resourceSwitchQosIpDscpMapRead(d, m)
}

func resourceSwitchQosIpDscpMapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchQosIpDscpMap(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchQosIpDscpMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchQosIpDscpMapRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchQosIpDscpMap(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchQosIpDscpMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchQosIpDscpMap(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchQosIpDscpMap resource from API: %v", err)
	}
	return nil
}

func flattenSwitchQosIpDscpMapMap(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffserv"
		if _, ok := i["diffserv"]; ok {

			tmp["diffserv"] = flattenSwitchQosIpDscpMapMapDiffserv(i["diffserv"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos_queue"
		if _, ok := i["cos-queue"]; ok {

			tmp["cos_queue"] = flattenSwitchQosIpDscpMapMapCosQueue(i["cos-queue"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {

			tmp["value"] = flattenSwitchQosIpDscpMapMapValue(i["value"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_precedence"
		if _, ok := i["ip-precedence"]; ok {

			tmp["ip_precedence"] = flattenSwitchQosIpDscpMapMapIpPrecedence(i["ip-precedence"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_name"
		if _, ok := i["entry-name"]; ok {

			tmp["entry_name"] = flattenSwitchQosIpDscpMapMapEntryName(i["entry-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenSwitchQosIpDscpMapMapType(i["type"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "entry_name", d)
	return result
}

func flattenSwitchQosIpDscpMapMapDiffserv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosIpDscpMapMapCosQueue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosIpDscpMapMapValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosIpDscpMapMapIpPrecedence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosIpDscpMapMapEntryName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosIpDscpMapMapType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosIpDscpMapName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchQosIpDscpMapDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchQosIpDscpMap(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if isImportTable() {
		if err = d.Set("map", flattenSwitchQosIpDscpMapMap(o["map"], d, "map", sv)); err != nil {
			if !fortiAPIPatch(o["map"]) {
				return fmt.Errorf("Error reading map: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("map"); ok {
			if err = d.Set("map", flattenSwitchQosIpDscpMapMap(o["map"], d, "map", sv)); err != nil {
				if !fortiAPIPatch(o["map"]) {
					return fmt.Errorf("Error reading map: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenSwitchQosIpDscpMapName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchQosIpDscpMapDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	return nil
}

func flattenSwitchQosIpDscpMapFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchQosIpDscpMapMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffserv"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["diffserv"], _ = expandSwitchQosIpDscpMapMapDiffserv(d, i["diffserv"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos_queue"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cos-queue"], _ = expandSwitchQosIpDscpMapMapCosQueue(d, i["cos_queue"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["value"], _ = expandSwitchQosIpDscpMapMapValue(d, i["value"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_precedence"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip-precedence"], _ = expandSwitchQosIpDscpMapMapIpPrecedence(d, i["ip_precedence"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["entry-name"], _ = expandSwitchQosIpDscpMapMapEntryName(d, i["entry_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandSwitchQosIpDscpMapMapType(d, i["type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchQosIpDscpMapMapDiffserv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosIpDscpMapMapCosQueue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosIpDscpMapMapValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosIpDscpMapMapIpPrecedence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosIpDscpMapMapEntryName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosIpDscpMapMapType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosIpDscpMapName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchQosIpDscpMapDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchQosIpDscpMap(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("map"); ok || d.HasChange("map") {

		t, err := expandSwitchQosIpDscpMapMap(d, v, "map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["map"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchQosIpDscpMapName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchQosIpDscpMapDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	return &obj, nil
}
