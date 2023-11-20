// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: PTP policy configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemPtpProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPtpProfileCreate,
		Read:   resourceSystemPtpProfileRead,
		Update: resourceSystemPtpProfileUpdate,
		Delete: resourceSystemPtpProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"pdelay_req_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ptp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemPtpProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemPtpProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemPtpProfile resource while getting object: %v", err)
	}

	o, err := c.CreateSystemPtpProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemPtpProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemPtpProfile")
	}

	return resourceSystemPtpProfileRead(d, m)
}

func resourceSystemPtpProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemPtpProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemPtpProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemPtpProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemPtpProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemPtpProfile")
	}

	return resourceSystemPtpProfileRead(d, m)
}

func resourceSystemPtpProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemPtpProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPtpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPtpProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemPtpProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemPtpProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPtpProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemPtpProfile resource from API: %v", err)
	}
	return nil
}

func flattenSystemPtpProfileDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpProfilePdelayReqInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpProfileMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpProfilePtpProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpProfileTransport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpProfileDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemPtpProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("domain", flattenSystemPtpProfileDomain(o["domain"], d, "domain", sv)); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemPtpProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pdelay_req_interval", flattenSystemPtpProfilePdelayReqInterval(o["pdelay-req-interval"], d, "pdelay_req_interval", sv)); err != nil {
		if !fortiAPIPatch(o["pdelay-req-interval"]) {
			return fmt.Errorf("Error reading pdelay_req_interval: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemPtpProfileMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("ptp_profile", flattenSystemPtpProfilePtpProfile(o["ptp-profile"], d, "ptp_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ptp-profile"]) {
			return fmt.Errorf("Error reading ptp_profile: %v", err)
		}
	}

	if err = d.Set("transport", flattenSystemPtpProfileTransport(o["transport"], d, "transport", sv)); err != nil {
		if !fortiAPIPatch(o["transport"]) {
			return fmt.Errorf("Error reading transport: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemPtpProfileDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	return nil
}

func flattenSystemPtpProfileFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemPtpProfileDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpProfilePdelayReqInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpProfileMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpProfilePtpProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpProfileTransport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpProfileDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPtpProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("domain"); ok {

		t, err := expandSystemPtpProfileDomain(d, v, "domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemPtpProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pdelay_req_interval"); ok {

		t, err := expandSystemPtpProfilePdelayReqInterval(d, v, "pdelay_req_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdelay-req-interval"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {

		t, err := expandSystemPtpProfileMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("ptp_profile"); ok {

		t, err := expandSystemPtpProfilePtpProfile(d, v, "ptp_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptp-profile"] = t
		}
	}

	if v, ok := d.GetOk("transport"); ok {

		t, err := expandSystemPtpProfileTransport(d, v, "transport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transport"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSystemPtpProfileDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	return &obj, nil
}
