// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Ocsp configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCertificateOcsp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCertificateOcspUpdate,
		Read:   resourceSystemCertificateOcspRead,
		Update: resourceSystemCertificateOcspUpdate,
		Delete: resourceSystemCertificateOcspDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"unavail_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemCertificateOcspUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCertificateOcsp(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateOcsp resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemCertificateOcsp(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateOcsp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCertificateOcsp")
	}

	return resourceSystemCertificateOcspRead(d, m)
}

func resourceSystemCertificateOcspDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCertificateOcsp(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateOcsp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateOcsp(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemCertificateOcsp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateOcspRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemCertificateOcsp(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateOcsp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateOcsp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateOcsp resource from API: %v", err)
	}
	return nil
}

func flattenSystemCertificateOcspUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateOcspUnavailAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateOcsp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("url", flattenSystemCertificateOcspUrl(o["url"], d, "url", sv)); err != nil {
		if !fortiAPIPatch(o["url"]) {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("cert", flattenSystemCertificateOcspCert(o["cert"], d, "cert", sv)); err != nil {
		if !fortiAPIPatch(o["cert"]) {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("unavail_action", flattenSystemCertificateOcspUnavailAction(o["unavail-action"], d, "unavail_action", sv)); err != nil {
		if !fortiAPIPatch(o["unavail-action"]) {
			return fmt.Errorf("Error reading unavail_action: %v", err)
		}
	}

	return nil
}

func flattenSystemCertificateOcspFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemCertificateOcspUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateOcspUnavailAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateOcsp(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("url"); ok {
		if setArgNil {
			obj["url"] = nil
		} else {

			t, err := expandSystemCertificateOcspUrl(d, v, "url", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["url"] = t
			}
		}
	}

	if v, ok := d.GetOk("cert"); ok {
		if setArgNil {
			obj["cert"] = nil
		} else {

			t, err := expandSystemCertificateOcspCert(d, v, "cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("unavail_action"); ok {
		if setArgNil {
			obj["unavail-action"] = nil
		} else {

			t, err := expandSystemCertificateOcspUnavailAction(d, v, "unavail_action", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["unavail-action"] = t
			}
		}
	}

	return &obj, nil
}
