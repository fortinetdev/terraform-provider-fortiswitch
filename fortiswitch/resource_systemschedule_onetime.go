// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: onetime schedule configuration

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemScheduleOnetime() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemScheduleOnetimeCreate,
		Read:   resourceSystemScheduleOnetimeRead,
		Update: resourceSystemScheduleOnetimeUpdate,
		Delete: resourceSystemScheduleOnetimeDelete,

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
		},
	}
}

func resourceSystemScheduleOnetimeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemScheduleOnetime(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemScheduleOnetime resource while getting object: %v", err)
	}

	o, err := c.CreateSystemScheduleOnetime(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemScheduleOnetime resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemScheduleOnetime")
	}

	return resourceSystemScheduleOnetimeRead(d, m)
}

func resourceSystemScheduleOnetimeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemScheduleOnetime(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemScheduleOnetime resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemScheduleOnetime(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemScheduleOnetime resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemScheduleOnetime")
	}

	return resourceSystemScheduleOnetimeRead(d, m)
}

func resourceSystemScheduleOnetimeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemScheduleOnetime(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemScheduleOnetime resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemScheduleOnetimeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemScheduleOnetime(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemScheduleOnetime resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemScheduleOnetime(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemScheduleOnetime resource from API: %v", err)
	}
	return nil
}

func flattenSystemScheduleOnetimeStart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemScheduleOnetimeEnd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemScheduleOnetimeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemScheduleOnetime(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("start", flattenSystemScheduleOnetimeStart(o["start"], d, "start", sv)); err != nil {
		if !fortiAPIPatch(o["start"]) {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	if err = d.Set("end", flattenSystemScheduleOnetimeEnd(o["end"], d, "end", sv)); err != nil {
		if !fortiAPIPatch(o["end"]) {
			return fmt.Errorf("Error reading end: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemScheduleOnetimeName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemScheduleOnetimeFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemScheduleOnetimeStart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemScheduleOnetimeEnd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemScheduleOnetimeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemScheduleOnetime(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("start"); ok {

		t, err := expandSystemScheduleOnetimeStart(d, v, "start", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start"] = t
		}
	}

	if v, ok := d.GetOk("end"); ok {

		t, err := expandSystemScheduleOnetimeEnd(d, v, "end", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemScheduleOnetimeName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
