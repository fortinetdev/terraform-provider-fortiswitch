// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Object tags.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemObjectTag() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemObjectTagCreate,
		Read:   resourceSystemObjectTagRead,
		Update: resourceSystemObjectTagUpdate,
		Delete: resourceSystemObjectTagDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
			},
		},
	}
}

func resourceSystemObjectTagCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemObjectTag(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemObjectTag resource while getting object: %v", err)
	}

	o, err := c.CreateSystemObjectTag(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemObjectTag resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemObjectTag")
	}

	return resourceSystemObjectTagRead(d, m)
}

func resourceSystemObjectTagUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemObjectTag(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemObjectTag resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemObjectTag(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemObjectTag resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemObjectTag")
	}

	return resourceSystemObjectTagRead(d, m)
}

func resourceSystemObjectTagDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemObjectTag(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemObjectTag resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemObjectTagRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemObjectTag(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemObjectTag resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemObjectTag(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemObjectTag resource from API: %v", err)
	}
	return nil
}

func flattenSystemObjectTagName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemObjectTag(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemObjectTagName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemObjectTagFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemObjectTagName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemObjectTag(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemObjectTagName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
