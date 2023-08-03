// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: recurring schedule configuration

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemScheduleRecurring() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemScheduleRecurringCreate,
		Read:   resourceSystemScheduleRecurringRead,
		Update: resourceSystemScheduleRecurringUpdate,
		Delete: resourceSystemScheduleRecurringDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemScheduleRecurringCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemScheduleRecurring(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemScheduleRecurring resource while getting object: %v", err)
	}

	o, err := c.CreateSystemScheduleRecurring(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemScheduleRecurring resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemScheduleRecurring")
	}

	return resourceSystemScheduleRecurringRead(d, m)
}

func resourceSystemScheduleRecurringUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemScheduleRecurring(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemScheduleRecurring resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemScheduleRecurring(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemScheduleRecurring resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemScheduleRecurring")
	}

	return resourceSystemScheduleRecurringRead(d, m)
}

func resourceSystemScheduleRecurringDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemScheduleRecurring(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemScheduleRecurring resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemScheduleRecurringRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemScheduleRecurring(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemScheduleRecurring resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemScheduleRecurring(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemScheduleRecurring resource from API: %v", err)
	}
	return nil
}

func flattenSystemScheduleRecurringStart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemScheduleRecurringEnd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemScheduleRecurringName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemScheduleRecurringDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemScheduleRecurring(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("start", flattenSystemScheduleRecurringStart(o["start"], d, "start", sv)); err != nil {
		if !fortiAPIPatch(o["start"]) {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	if err = d.Set("end", flattenSystemScheduleRecurringEnd(o["end"], d, "end", sv)); err != nil {
		if !fortiAPIPatch(o["end"]) {
			return fmt.Errorf("Error reading end: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemScheduleRecurringName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("day", flattenSystemScheduleRecurringDay(o["day"], d, "day", sv)); err != nil {
		if !fortiAPIPatch(o["day"]) {
			return fmt.Errorf("Error reading day: %v", err)
		}
	}

	return nil
}

func flattenSystemScheduleRecurringFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemScheduleRecurringStart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemScheduleRecurringEnd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemScheduleRecurringName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemScheduleRecurringDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemScheduleRecurring(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("start"); ok {

		t, err := expandSystemScheduleRecurringStart(d, v, "start", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start"] = t
		}
	}

	if v, ok := d.GetOk("end"); ok {

		t, err := expandSystemScheduleRecurringEnd(d, v, "end", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemScheduleRecurringName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("day"); ok {

		t, err := expandSystemScheduleRecurringDay(d, v, "day", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["day"] = t
		}
	}

	return &obj, nil
}
