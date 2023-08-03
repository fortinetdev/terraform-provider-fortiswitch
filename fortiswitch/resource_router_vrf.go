// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: VRF configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterVrf() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterVrfCreate,
		Read:   resourceRouterVrfRead,
		Update: resourceRouterVrfUpdate,
		Delete: resourceRouterVrfDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vrfid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceRouterVrfCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterVrf(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterVrf resource while getting object: %v", err)
	}

	o, err := c.CreateRouterVrf(obj)

	if err != nil {
		return fmt.Errorf("Error creating RouterVrf resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterVrf")
	}

	return resourceRouterVrfRead(d, m)
}

func resourceRouterVrfUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterVrf(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterVrf resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterVrf(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterVrf resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterVrf")
	}

	return resourceRouterVrfRead(d, m)
}

func resourceRouterVrfDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterVrf(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting RouterVrf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterVrfRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterVrf(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterVrf resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterVrf(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterVrf resource from API: %v", err)
	}
	return nil
}

func flattenRouterVrfVrfid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterVrfName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterVrf(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("vrfid", flattenRouterVrfVrfid(o["vrfid"], d, "vrfid", sv)); err != nil {
		if !fortiAPIPatch(o["vrfid"]) {
			return fmt.Errorf("Error reading vrfid: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterVrfName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenRouterVrfFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandRouterVrfVrfid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterVrfName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterVrf(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("vrfid"); ok {

		t, err := expandRouterVrfVrfid(d, v, "vrfid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrfid"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandRouterVrfName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
