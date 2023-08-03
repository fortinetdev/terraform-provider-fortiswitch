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

func dataSourceRouterMulticast() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterMulticastRead,
		Schema: map[string]*schema.Schema{

			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hello_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"pim_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dr_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"multicast_flow": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"igmp": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"query_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"query_max_response_time": &schema.Schema{
										Type:     schema.TypeInt,
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
				Computed: true,
			},
		},
	}
}

func dataSourceRouterMulticastRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "RouterMulticast"

	o, err := c.ReadRouterMulticast(mkey)
	if err != nil {
		return fmt.Errorf("Error describing RouterMulticast: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterMulticast(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterMulticast from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterMulticastInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = dataSourceFlattenRouterMulticastInterfaceHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterMulticastInterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pim_mode"
		if _, ok := i["pim-mode"]; ok {
			tmp["pim_mode"] = dataSourceFlattenRouterMulticastInterfacePimMode(i["pim-mode"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dr_priority"
		if _, ok := i["dr-priority"]; ok {
			tmp["dr_priority"] = dataSourceFlattenRouterMulticastInterfaceDrPriority(i["dr-priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_flow"
		if _, ok := i["multicast-flow"]; ok {
			tmp["multicast_flow"] = dataSourceFlattenRouterMulticastInterfaceMulticastFlow(i["multicast-flow"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp"
		if _, ok := i["igmp"]; ok {
			tmp["igmp"] = dataSourceFlattenRouterMulticastInterfaceIgmp(i["igmp"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterMulticastInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfacePimMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceDrPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceMulticastFlow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceIgmp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "query_interval"
	if _, ok := i["query-interval"]; ok {
		result["query_interval"] = dataSourceFlattenRouterMulticastInterfaceIgmpQueryInterval(i["query-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "query_max_response_time"
	if _, ok := i["query-max-response-time"]; ok {
		result["query_max_response_time"] = dataSourceFlattenRouterMulticastInterfaceIgmpQueryMaxResponseTime(i["query-max-response-time"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenRouterMulticastInterfaceIgmpQueryInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceIgmpQueryMaxResponseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastMulticastRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterMulticast(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("interface", dataSourceFlattenRouterMulticastInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("multicast_routing", dataSourceFlattenRouterMulticastMulticastRouting(o["multicast-routing"], d, "multicast_routing")); err != nil {
		if !fortiAPIPatch(o["multicast-routing"]) {
			return fmt.Errorf("Error reading multicast_routing: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterMulticastFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
