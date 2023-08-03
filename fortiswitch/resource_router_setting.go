// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Set rib settings.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterSettingUpdate,
		Read:   resourceRouterSettingRead,
		Update: resourceRouterSettingUpdate,
		Delete: resourceRouterSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"filter_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"route_map": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"hostname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 14),
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

func resourceRouterSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterSetting")
	}

	return resourceRouterSettingRead(d, m)
}

func resourceRouterSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating RouterSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing RouterSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterSetting resource from API: %v", err)
	}
	return nil
}

func flattenRouterSettingFilterList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {

			tmp["route_map"] = flattenRouterSettingFilterListRouteMap(i["route-map"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {

			tmp["protocol"] = flattenRouterSettingFilterListProtocol(i["protocol"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterSettingFilterListId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterSettingFilterListRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingFilterListProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingFilterListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if isImportTable() {
		if err = d.Set("filter_list", flattenRouterSettingFilterList(o["filter-list"], d, "filter_list", sv)); err != nil {
			if !fortiAPIPatch(o["filter-list"]) {
				return fmt.Errorf("Error reading filter_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filter_list"); ok {
			if err = d.Set("filter_list", flattenRouterSettingFilterList(o["filter-list"], d, "filter_list", sv)); err != nil {
				if !fortiAPIPatch(o["filter-list"]) {
					return fmt.Errorf("Error reading filter_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("hostname", flattenRouterSettingHostname(o["hostname"], d, "hostname", sv)); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	return nil
}

func flattenRouterSettingFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandRouterSettingFilterList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map"], _ = expandRouterSettingFilterListRouteMap(d, i["route_map"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["protocol"], _ = expandRouterSettingFilterListProtocol(d, i["protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterSettingFilterListId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterSettingFilterListRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingFilterListProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingFilterListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("filter_list"); ok || d.HasChange("filter_list") {
		if setArgNil {
			obj["filter-list"] = make([]struct{}, 0)
		} else {

			t, err := expandRouterSettingFilterList(d, v, "filter_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["filter-list"] = t
			}
		}
	}

	if v, ok := d.GetOk("hostname"); ok {
		if setArgNil {
			obj["hostname"] = nil
		} else {

			t, err := expandRouterSettingHostname(d, v, "hostname", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostname"] = t
			}
		}
	}

	return &obj, nil
}
