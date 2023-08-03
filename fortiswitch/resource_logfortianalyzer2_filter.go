// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Filters for FortiAnalyzer.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogFortianalyzer2Filter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortianalyzer2FilterUpdate,
		Read:   resourceLogFortianalyzer2FilterRead,
		Update: resourceLogFortianalyzer2FilterUpdate,
		Delete: resourceLogFortianalyzer2FilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogFortianalyzer2FilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogFortianalyzer2Filter(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer2Filter resource while getting object: %v", err)
	}

	o, err := c.UpdateLogFortianalyzer2Filter(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer2Filter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogFortianalyzer2Filter")
	}

	return resourceLogFortianalyzer2FilterRead(d, m)
}

func resourceLogFortianalyzer2FilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogFortianalyzer2Filter(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer2Filter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogFortianalyzer2Filter(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing LogFortianalyzer2Filter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortianalyzer2FilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogFortianalyzer2Filter(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer2Filter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortianalyzer2Filter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer2Filter resource from API: %v", err)
	}
	return nil
}

func flattenLogFortianalyzer2FilterOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2FilterSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogFortianalyzer2Filter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("override", flattenLogFortianalyzer2FilterOverride(o["override"], d, "override", sv)); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("severity", flattenLogFortianalyzer2FilterSeverity(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	return nil
}

func flattenLogFortianalyzer2FilterFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandLogFortianalyzer2FilterOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2FilterSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortianalyzer2Filter(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("override"); ok {
		if setArgNil {
			obj["override"] = nil
		} else {

			t, err := expandLogFortianalyzer2FilterOverride(d, v, "override", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["override"] = t
			}
		}
	}

	if v, ok := d.GetOk("severity"); ok {
		if setArgNil {
			obj["severity"] = nil
		} else {

			t, err := expandLogFortianalyzer2FilterSeverity(d, v, "severity", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["severity"] = t
			}
		}
	}

	return &obj, nil
}
