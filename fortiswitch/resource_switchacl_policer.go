// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Policer configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchAclPolicer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchAclPolicerCreate,
		Read:   resourceSwitchAclPolicerRead,
		Update: resourceSwitchAclPolicerUpdate,
		Delete: resourceSwitchAclPolicerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"maximum_burst": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"guaranteed_burst": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"guaranteed_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fswid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchAclPolicerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAclPolicer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchAclPolicer resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchAclPolicer(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchAclPolicer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchAclPolicer")
	}

	return resourceSwitchAclPolicerRead(d, m)
}

func resourceSwitchAclPolicerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAclPolicer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAclPolicer resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchAclPolicer(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAclPolicer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchAclPolicer")
	}

	return resourceSwitchAclPolicerRead(d, m)
}

func resourceSwitchAclPolicerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchAclPolicer(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchAclPolicer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchAclPolicerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchAclPolicer(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAclPolicer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchAclPolicer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAclPolicer resource from API: %v", err)
	}
	return nil
}

func flattenSwitchAclPolicerMaximumBurst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPolicerGuaranteedBurst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPolicerDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPolicerGuaranteedBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPolicerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclPolicerId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchAclPolicer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("maximum_burst", flattenSwitchAclPolicerMaximumBurst(o["maximum-burst"], d, "maximum_burst", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-burst"]) {
			return fmt.Errorf("Error reading maximum_burst: %v", err)
		}
	}

	if err = d.Set("guaranteed_burst", flattenSwitchAclPolicerGuaranteedBurst(o["guaranteed-burst"], d, "guaranteed_burst", sv)); err != nil {
		if !fortiAPIPatch(o["guaranteed-burst"]) {
			return fmt.Errorf("Error reading guaranteed_burst: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchAclPolicerDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("guaranteed_bandwidth", flattenSwitchAclPolicerGuaranteedBandwidth(o["guaranteed-bandwidth"], d, "guaranteed_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["guaranteed-bandwidth"]) {
			return fmt.Errorf("Error reading guaranteed_bandwidth: %v", err)
		}
	}

	if err = d.Set("type", flattenSwitchAclPolicerType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("fswid", flattenSwitchAclPolicerId(o["id"], d, "fswid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fswid: %v", err)
		}
	}

	return nil
}

func flattenSwitchAclPolicerFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchAclPolicerMaximumBurst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPolicerGuaranteedBurst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPolicerDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPolicerGuaranteedBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPolicerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclPolicerId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchAclPolicer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("maximum_burst"); ok {

		t, err := expandSwitchAclPolicerMaximumBurst(d, v, "maximum_burst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-burst"] = t
		}
	}

	if v, ok := d.GetOk("guaranteed_burst"); ok {

		t, err := expandSwitchAclPolicerGuaranteedBurst(d, v, "guaranteed_burst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guaranteed-burst"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchAclPolicerDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("guaranteed_bandwidth"); ok {

		t, err := expandSwitchAclPolicerGuaranteedBandwidth(d, v, "guaranteed_bandwidth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guaranteed-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandSwitchAclPolicerType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("fswid"); ok {

		t, err := expandSwitchAclPolicerId(d, v, "fswid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
