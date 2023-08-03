// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Remote certificate.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCertificateRemote() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCertificateRemoteCreate,
		Read:   resourceSystemCertificateRemoteRead,
		Update: resourceSystemCertificateRemoteUpdate,
		Delete: resourceSystemCertificateRemoteDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"remote": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceSystemCertificateRemoteCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCertificateRemote(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateRemote resource while getting object: %v", err)
	}

	o, err := c.CreateSystemCertificateRemote(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateRemote resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCertificateRemote")
	}

	return resourceSystemCertificateRemoteRead(d, m)
}

func resourceSystemCertificateRemoteUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCertificateRemote(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateRemote resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemCertificateRemote(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateRemote resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCertificateRemote")
	}

	return resourceSystemCertificateRemoteRead(d, m)
}

func resourceSystemCertificateRemoteDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemCertificateRemote(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateRemote resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateRemoteRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemCertificateRemote(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateRemote resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateRemote(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateRemote resource from API: %v", err)
	}
	return nil
}

func flattenSystemCertificateRemoteRemote(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateRemoteName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateRemote(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("remote", flattenSystemCertificateRemoteRemote(o["remote"], d, "remote", sv)); err != nil {
		if !fortiAPIPatch(o["remote"]) {
			return fmt.Errorf("Error reading remote: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemCertificateRemoteName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemCertificateRemoteFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemCertificateRemoteRemote(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateRemoteName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateRemote(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("remote"); ok {

		t, err := expandSystemCertificateRemoteRemote(d, v, "remote", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemCertificateRemoteName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
