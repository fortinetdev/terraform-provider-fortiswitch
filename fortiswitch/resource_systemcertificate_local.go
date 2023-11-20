// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Local keys and certificates.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCertificateLocal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCertificateLocalCreate,
		Read:   resourceSystemCertificateLocalRead,
		Update: resourceSystemCertificateLocalUpdate,
		Delete: resourceSystemCertificateLocalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scep_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 123),
				Optional:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"auto_regenerate_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_regenerate_days_warning": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scep_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"csr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemCertificateLocalCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCertificateLocal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLocal resource while getting object: %v", err)
	}

	o, err := c.CreateSystemCertificateLocal(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLocal resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCertificateLocal")
	}

	return resourceSystemCertificateLocalRead(d, m)
}

func resourceSystemCertificateLocalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCertificateLocal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateLocal resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemCertificateLocal(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateLocal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCertificateLocal")
	}

	return resourceSystemCertificateLocalRead(d, m)
}

func resourceSystemCertificateLocalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemCertificateLocal(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateLocalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemCertificateLocal(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateLocal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLocal resource from API: %v", err)
	}
	return nil
}

func flattenSystemCertificateLocalInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalScepPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalPrivateKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalAutoRegenerateDays(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalAutoRegenerateDaysWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalNameEncoding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalScepUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateLocalCsr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateLocal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("info", flattenSystemCertificateLocalInfo(o["Info"], d, "info", sv)); err != nil {
		if !fortiAPIPatch(o["Info"]) {
			return fmt.Errorf("Error reading info: %v", err)
		}
	}

	if err = d.Set("scep_password", flattenSystemCertificateLocalScepPassword(o["scep-password"], d, "scep_password", sv)); err != nil {
		if !fortiAPIPatch(o["scep-password"]) {
			return fmt.Errorf("Error reading scep_password: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemCertificateLocalName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("certificate", flattenSystemCertificateLocalCertificate(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("private_key", flattenSystemCertificateLocalPrivateKey(o["private-key"], d, "private_key", sv)); err != nil {
		if !fortiAPIPatch(o["private-key"]) {
			return fmt.Errorf("Error reading private_key: %v", err)
		}
	}

	if err = d.Set("comments", flattenSystemCertificateLocalComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("auto_regenerate_days", flattenSystemCertificateLocalAutoRegenerateDays(o["auto-regenerate-days"], d, "auto_regenerate_days", sv)); err != nil {
		if !fortiAPIPatch(o["auto-regenerate-days"]) {
			return fmt.Errorf("Error reading auto_regenerate_days: %v", err)
		}
	}

	if err = d.Set("auto_regenerate_days_warning", flattenSystemCertificateLocalAutoRegenerateDaysWarning(o["auto-regenerate-days-warning"], d, "auto_regenerate_days_warning", sv)); err != nil {
		if !fortiAPIPatch(o["auto-regenerate-days-warning"]) {
			return fmt.Errorf("Error reading auto_regenerate_days_warning: %v", err)
		}
	}

	if err = d.Set("name_encoding", flattenSystemCertificateLocalNameEncoding(o["name-encoding"], d, "name_encoding", sv)); err != nil {
		if !fortiAPIPatch(o["name-encoding"]) {
			return fmt.Errorf("Error reading name_encoding: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenSystemCertificateLocalScepUrl(o["scep-url"], d, "scep_url", sv)); err != nil {
		if !fortiAPIPatch(o["scep-url"]) {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("password", flattenSystemCertificateLocalPassword(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("csr", flattenSystemCertificateLocalCsr(o["csr"], d, "csr", sv)); err != nil {
		if !fortiAPIPatch(o["csr"]) {
			return fmt.Errorf("Error reading csr: %v", err)
		}
	}

	return nil
}

func flattenSystemCertificateLocalFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemCertificateLocalInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalScepPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalPrivateKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalAutoRegenerateDays(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalAutoRegenerateDaysWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalNameEncoding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalScepUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateLocalCsr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateLocal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("info"); ok {

		t, err := expandSystemCertificateLocalInfo(d, v, "info", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Info"] = t
		}
	}

	if v, ok := d.GetOk("scep_password"); ok {

		t, err := expandSystemCertificateLocalScepPassword(d, v, "scep_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-password"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemCertificateLocalName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok {

		t, err := expandSystemCertificateLocalCertificate(d, v, "certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("private_key"); ok {

		t, err := expandSystemCertificateLocalPrivateKey(d, v, "private_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandSystemCertificateLocalComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("auto_regenerate_days"); ok {

		t, err := expandSystemCertificateLocalAutoRegenerateDays(d, v, "auto_regenerate_days", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-regenerate-days"] = t
		}
	}

	if v, ok := d.GetOk("auto_regenerate_days_warning"); ok {

		t, err := expandSystemCertificateLocalAutoRegenerateDaysWarning(d, v, "auto_regenerate_days_warning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-regenerate-days-warning"] = t
		}
	}

	if v, ok := d.GetOk("name_encoding"); ok {

		t, err := expandSystemCertificateLocalNameEncoding(d, v, "name_encoding", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name-encoding"] = t
		}
	}

	if v, ok := d.GetOk("scep_url"); ok {

		t, err := expandSystemCertificateLocalScepUrl(d, v, "scep_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {

		t, err := expandSystemCertificateLocalPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("csr"); ok {

		t, err := expandSystemCertificateLocalCsr(d, v, "csr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csr"] = t
		}
	}

	return &obj, nil
}
