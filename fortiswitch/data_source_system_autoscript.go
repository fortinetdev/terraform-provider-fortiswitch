// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure auto script.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemAutoScript() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAutoScriptRead,
		Schema: map[string]*schema.Schema{

			"output_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"repeat": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemAutoScriptRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemAutoScript: type error")
	}

	o, err := c.ReadSystemAutoScript(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemAutoScript: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemAutoScript(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemAutoScript from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemAutoScriptOutputSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoScriptRepeat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoScriptName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoScriptScript(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoScriptInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoScriptStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAutoScript(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("output_size", dataSourceFlattenSystemAutoScriptOutputSize(o["output-size"], d, "output_size")); err != nil {
		if !fortiAPIPatch(o["output-size"]) {
			return fmt.Errorf("Error reading output_size: %v", err)
		}
	}

	if err = d.Set("repeat", dataSourceFlattenSystemAutoScriptRepeat(o["repeat"], d, "repeat")); err != nil {
		if !fortiAPIPatch(o["repeat"]) {
			return fmt.Errorf("Error reading repeat: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenSystemAutoScriptName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("script", dataSourceFlattenSystemAutoScriptScript(o["script"], d, "script")); err != nil {
		if !fortiAPIPatch(o["script"]) {
			return fmt.Errorf("Error reading script: %v", err)
		}
	}

	if err = d.Set("interval", dataSourceFlattenSystemAutoScriptInterval(o["interval"], d, "interval")); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("start", dataSourceFlattenSystemAutoScriptStart(o["start"], d, "start")); err != nil {
		if !fortiAPIPatch(o["start"]) {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemAutoScriptFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
