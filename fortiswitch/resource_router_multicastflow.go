// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Multicast-flow configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticastFlow() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticastFlowCreate,
		Read:   resourceRouterMulticastFlowRead,
		Update: resourceRouterMulticastFlowUpdate,
		Delete: resourceRouterMulticastFlowDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"flows": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"group_addr_end": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
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

func resourceRouterMulticastFlowCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterMulticastFlow(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticastFlow resource while getting object: %v", err)
	}

	o, err := c.CreateRouterMulticastFlow(obj)

	if err != nil {
		return fmt.Errorf("Error creating RouterMulticastFlow resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterMulticastFlow")
	}

	return resourceRouterMulticastFlowRead(d, m)
}

func resourceRouterMulticastFlowUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterMulticastFlow(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastFlow resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterMulticastFlow(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastFlow resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterMulticastFlow")
	}

	return resourceRouterMulticastFlowRead(d, m)
}

func resourceRouterMulticastFlowDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterMulticastFlow(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticastFlow resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticastFlowRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterMulticastFlow(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastFlow resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticastFlow(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastFlow resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticastFlowFlows(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_addr"
		if _, ok := i["group-addr"]; ok {

			tmp["group_addr"] = flattenRouterMulticastFlowFlowsGroupAddr(i["group-addr"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_addr"
		if _, ok := i["source-addr"]; ok {

			tmp["source_addr"] = flattenRouterMulticastFlowFlowsSourceAddr(i["source-addr"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterMulticastFlowFlowsId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_addr_end"
		if _, ok := i["group-addr-end"]; ok {

			tmp["group_addr_end"] = flattenRouterMulticastFlowFlowsGroupAddrEnd(i["group-addr-end"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterMulticastFlowFlowsGroupAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastFlowFlowsSourceAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastFlowFlowsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastFlowFlowsGroupAddrEnd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastFlowName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastFlowComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterMulticastFlow(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if isImportTable() {
		if err = d.Set("flows", flattenRouterMulticastFlowFlows(o["flows"], d, "flows", sv)); err != nil {
			if !fortiAPIPatch(o["flows"]) {
				return fmt.Errorf("Error reading flows: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("flows"); ok {
			if err = d.Set("flows", flattenRouterMulticastFlowFlows(o["flows"], d, "flows", sv)); err != nil {
				if !fortiAPIPatch(o["flows"]) {
					return fmt.Errorf("Error reading flows: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenRouterMulticastFlowName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenRouterMulticastFlowComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticastFlowFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandRouterMulticastFlowFlows(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_addr"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["group-addr"], _ = expandRouterMulticastFlowFlowsGroupAddr(d, i["group_addr"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_addr"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["source-addr"], _ = expandRouterMulticastFlowFlowsSourceAddr(d, i["source_addr"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterMulticastFlowFlowsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_addr_end"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["group-addr-end"], _ = expandRouterMulticastFlowFlowsGroupAddrEnd(d, i["group_addr_end"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterMulticastFlowFlowsGroupAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastFlowFlowsSourceAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastFlowFlowsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastFlowFlowsGroupAddrEnd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastFlowName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastFlowComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticastFlow(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("flows"); ok || d.HasChange("flows") {

		t, err := expandRouterMulticastFlowFlows(d, v, "flows", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flows"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandRouterMulticastFlowName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandRouterMulticastFlowComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	return &obj, nil
}
