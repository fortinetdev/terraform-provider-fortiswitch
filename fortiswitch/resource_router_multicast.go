// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Router multicast configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticast() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticastUpdate,
		Read:   resourceRouterMulticastRead,
		Update: resourceRouterMulticastUpdate,
		Delete: resourceRouterMulticastDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hello_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"pim_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dr_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"multicast_flow": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"igmp": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"query_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"query_max_response_time": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"multicast_routing": &schema.Schema{
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

func resourceRouterMulticastUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterMulticast(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterMulticast(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterMulticast")
	}

	return resourceRouterMulticastRead(d, m)
}

func resourceRouterMulticastDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterMulticast(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMulticast(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing RouterMulticast resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticastRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterMulticast(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticast(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticastInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {

			tmp["hello_interval"] = flattenRouterMulticastInterfaceHelloInterval(i["hello-interval"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenRouterMulticastInterfaceName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pim_mode"
		if _, ok := i["pim-mode"]; ok {

			tmp["pim_mode"] = flattenRouterMulticastInterfacePimMode(i["pim-mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dr_priority"
		if _, ok := i["dr-priority"]; ok {

			tmp["dr_priority"] = flattenRouterMulticastInterfaceDrPriority(i["dr-priority"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_flow"
		if _, ok := i["multicast-flow"]; ok {

			tmp["multicast_flow"] = flattenRouterMulticastInterfaceMulticastFlow(i["multicast-flow"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp"
		if _, ok := i["igmp"]; ok {

			tmp["igmp"] = flattenRouterMulticastInterfaceIgmp(i["igmp"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterMulticastInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfacePimMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceDrPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceMulticastFlow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "query_interval"
	if _, ok := i["query-interval"]; ok {

		result["query_interval"] = flattenRouterMulticastInterfaceIgmpQueryInterval(i["query-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "query_max_response_time"
	if _, ok := i["query-max-response-time"]; ok {

		result["query_max_response_time"] = flattenRouterMulticastInterfaceIgmpQueryMaxResponseTime(i["query-max-response-time"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenRouterMulticastInterfaceIgmpQueryInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpQueryMaxResponseTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastMulticastRouting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterMulticast(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if isImportTable() {
		if err = d.Set("interface", flattenRouterMulticastInterface(o["interface"], d, "interface", sv)); err != nil {
			if !fortiAPIPatch(o["interface"]) {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenRouterMulticastInterface(o["interface"], d, "interface", sv)); err != nil {
				if !fortiAPIPatch(o["interface"]) {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("multicast_routing", flattenRouterMulticastMulticastRouting(o["multicast-routing"], d, "multicast_routing", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-routing"]) {
			return fmt.Errorf("Error reading multicast_routing: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticastFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandRouterMulticastInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["hello-interval"], _ = expandRouterMulticastInterfaceHelloInterval(d, i["hello_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandRouterMulticastInterfaceName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pim_mode"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["pim-mode"], _ = expandRouterMulticastInterfacePimMode(d, i["pim_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dr_priority"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dr-priority"], _ = expandRouterMulticastInterfaceDrPriority(d, i["dr_priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_flow"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["multicast-flow"], _ = expandRouterMulticastInterfaceMulticastFlow(d, i["multicast_flow"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["igmp"], _ = expandRouterMulticastInterfaceIgmp(d, i["igmp"], pre_append, sv)
		} else {
			tmp["igmp"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterMulticastInterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfacePimMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceDrPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceMulticastFlow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "query_interval"
	if _, ok := d.GetOk(pre_append); ok {

		result["query-interval"], _ = expandRouterMulticastInterfaceIgmpQueryInterval(d, i["query_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "query_max_response_time"
	if _, ok := d.GetOk(pre_append); ok {

		result["query-max-response-time"], _ = expandRouterMulticastInterfaceIgmpQueryMaxResponseTime(d, i["query_max_response_time"], pre_append, sv)
	}

	return result, nil
}

func expandRouterMulticastInterfaceIgmpQueryInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpQueryMaxResponseTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastMulticastRouting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticast(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		if setArgNil {
			obj["interface"] = make([]struct{}, 0)
		} else {

			t, err := expandRouterMulticastInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("multicast_routing"); ok {
		if setArgNil {
			obj["multicast-routing"] = nil
		} else {

			t, err := expandRouterMulticastMulticastRouting(d, v, "multicast_routing", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["multicast-routing"] = t
			}
		}
	}

	return &obj, nil
}
