// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure tos based priority table.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemTosBasedPriority() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemTosBasedPriorityRead,
		Schema: map[string]*schema.Schema{

			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"fswid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func dataSourceSystemTosBasedPriorityRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("fswid")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemTosBasedPriority: type error")
	}

	o, err := c.ReadSystemTosBasedPriority(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemTosBasedPriority: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemTosBasedPriority(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemTosBasedPriority from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemTosBasedPriorityPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemTosBasedPriorityTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemTosBasedPriorityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemTosBasedPriority(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("priority", dataSourceFlattenSystemTosBasedPriorityPriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("tos", dataSourceFlattenSystemTosBasedPriorityTos(o["tos"], d, "tos")); err != nil {
		if !fortiAPIPatch(o["tos"]) {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("fswid", dataSourceFlattenSystemTosBasedPriorityId(o["id"], d, "fswid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fswid: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemTosBasedPriorityFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
