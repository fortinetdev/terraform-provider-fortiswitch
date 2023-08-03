// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: CA certificate.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCertificateCa() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCertificateCaCreate,
		Read:   resourceSystemCertificateCaRead,
		Update: resourceSystemCertificateCaUpdate,
		Delete: resourceSystemCertificateCaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"scep_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"ca": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_update_days_warning": &schema.Schema{
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
			"auto_update_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemCertificateCaCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCertificateCa(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCa resource while getting object: %v", err)
	}

	o, err := c.CreateSystemCertificateCa(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCa resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCertificateCa")
	}

	return resourceSystemCertificateCaRead(d, m)
}

func resourceSystemCertificateCaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCertificateCa(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCa resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemCertificateCa(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCa resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCertificateCa")
	}

	return resourceSystemCertificateCaRead(d, m)
}

func resourceSystemCertificateCaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemCertificateCa(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateCa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateCaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemCertificateCa(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateCa(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCa resource from API: %v", err)
	}
	return nil
}

func flattenSystemCertificateCaScepUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCaCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCaAutoUpdateDaysWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCaName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCaAutoUpdateDays(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateCa(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("scep_url", flattenSystemCertificateCaScepUrl(o["scep-url"], d, "scep_url", sv)); err != nil {
		if !fortiAPIPatch(o["scep-url"]) {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("ca", flattenSystemCertificateCaCa(o["ca"], d, "ca", sv)); err != nil {
		if !fortiAPIPatch(o["ca"]) {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}

	if err = d.Set("auto_update_days_warning", flattenSystemCertificateCaAutoUpdateDaysWarning(o["auto-update-days-warning"], d, "auto_update_days_warning", sv)); err != nil {
		if !fortiAPIPatch(o["auto-update-days-warning"]) {
			return fmt.Errorf("Error reading auto_update_days_warning: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemCertificateCaName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("auto_update_days", flattenSystemCertificateCaAutoUpdateDays(o["auto-update-days"], d, "auto_update_days", sv)); err != nil {
		if !fortiAPIPatch(o["auto-update-days"]) {
			return fmt.Errorf("Error reading auto_update_days: %v", err)
		}
	}

	return nil
}

func flattenSystemCertificateCaFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemCertificateCaScepUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCaCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCaAutoUpdateDaysWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCaName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCaAutoUpdateDays(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateCa(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("scep_url"); ok {

		t, err := expandSystemCertificateCaScepUrl(d, v, "scep_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	}

	if v, ok := d.GetOk("ca"); ok {

		t, err := expandSystemCertificateCaCa(d, v, "ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}

	if v, ok := d.GetOk("auto_update_days_warning"); ok {

		t, err := expandSystemCertificateCaAutoUpdateDaysWarning(d, v, "auto_update_days_warning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-update-days-warning"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemCertificateCaName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("auto_update_days"); ok {

		t, err := expandSystemCertificateCaAutoUpdateDays(d, v, "auto_update_days", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-update-days"] = t
		}
	}

	return &obj, nil
}
