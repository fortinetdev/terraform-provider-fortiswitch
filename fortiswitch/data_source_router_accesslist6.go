// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: IPv6 access list configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceRouterAccessList6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterAccessList6Read,
		Schema: map[string]*schema.Schema{

			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"exact_match": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"flags": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix6": &schema.Schema{
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

func dataSourceRouterAccessList6Read(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterAccessList6: type error")
	}

	o, err := c.ReadRouterAccessList6(mkey)
	if err != nil {
		return fmt.Errorf("Error describing RouterAccessList6: %v", err)
	}

	if o == nil {
		d.SetId("")
		return fmt.Errorf("Error describing RouterAccessList6:  resource not found")
	}

	err = dataSourceRefreshObjectRouterAccessList6(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterAccessList6 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterAccessList6Rule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = dataSourceFlattenRouterAccessList6RuleAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exact_match"
		if _, ok := i["exact-match"]; ok {
			tmp["exact_match"] = dataSourceFlattenRouterAccessList6RuleExactMatch(i["exact-match"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {
			tmp["flags"] = dataSourceFlattenRouterAccessList6RuleFlags(i["flags"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterAccessList6RuleId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			tmp["prefix6"] = dataSourceFlattenRouterAccessList6RulePrefix6(i["prefix6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterAccessList6RuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterAccessList6RuleExactMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterAccessList6RuleFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterAccessList6RuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterAccessList6RulePrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterAccessList6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterAccessList6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterAccessList6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("rule", dataSourceFlattenRouterAccessList6Rule(o["rule"], d, "rule")); err != nil {
		if !fortiAPIPatch(o["rule"]) {
			return fmt.Errorf("Error reading rule: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenRouterAccessList6Name(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenRouterAccessList6Comments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterAccessList6FortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
