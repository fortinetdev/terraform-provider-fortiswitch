// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Virtual domain configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVdom() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomCreate,
		Read:   resourceSystemVdomRead,
		Update: resourceSystemVdomUpdate,
		Delete: resourceSystemVdomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vcluster_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 11),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemVdomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdom resource while getting object: %v", err)
	}

	o, err := c.CreateSystemVdom(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemVdom resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdom")
	}

	return resourceSystemVdomRead(d, m)
}

func resourceSystemVdomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdom resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVdom(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdom")
	}

	return resourceSystemVdomRead(d, m)
}

func resourceSystemVdomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemVdom(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemVdom(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdom(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdom resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomVclusterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemVdom(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("vcluster_id", flattenSystemVdomVclusterId(o["vcluster-id"], d, "vcluster_id", sv)); err != nil {
		if !fortiAPIPatch(o["vcluster-id"]) {
			return fmt.Errorf("Error reading vcluster_id: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemVdomName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemVdomVclusterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdom(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("vcluster_id"); ok {

		t, err := expandSystemVdomVclusterId(d, v, "vcluster_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcluster-id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemVdomName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
