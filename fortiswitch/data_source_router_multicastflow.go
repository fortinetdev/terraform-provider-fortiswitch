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

func dataSourceRouterMulticastFlow() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterMulticastFlowRead,
		Schema: map[string]*schema.Schema{

			"flows": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_addr": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_addr": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"group_addr_end": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRouterMulticastFlowRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterMulticastFlow: type error")
	}

	o, err := c.ReadRouterMulticastFlow(mkey)
	if err != nil {
		return fmt.Errorf("Error describing RouterMulticastFlow: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterMulticastFlow(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterMulticastFlow from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterMulticastFlowFlows(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
			tmp["group_addr"] = dataSourceFlattenRouterMulticastFlowFlowsGroupAddr(i["group-addr"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_addr"
		if _, ok := i["source-addr"]; ok {
			tmp["source_addr"] = dataSourceFlattenRouterMulticastFlowFlowsSourceAddr(i["source-addr"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterMulticastFlowFlowsId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_addr_end"
		if _, ok := i["group-addr-end"]; ok {
			tmp["group_addr_end"] = dataSourceFlattenRouterMulticastFlowFlowsGroupAddrEnd(i["group-addr-end"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterMulticastFlowFlowsGroupAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastFlowFlowsSourceAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastFlowFlowsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastFlowFlowsGroupAddrEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastFlowName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastFlowComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterMulticastFlow(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("flows", dataSourceFlattenRouterMulticastFlowFlows(o["flows"], d, "flows")); err != nil {
		if !fortiAPIPatch(o["flows"]) {
			return fmt.Errorf("Error reading flows: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenRouterMulticastFlowName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenRouterMulticastFlowComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterMulticastFlowFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
