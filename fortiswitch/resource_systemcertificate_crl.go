// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Certificate Revokation List.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCertificateCrl() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCertificateCrlCreate,
		Read:   resourceSystemCertificateCrlRead,
		Update: resourceSystemCertificateCrlUpdate,
		Delete: resourceSystemCertificateCrlDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"http_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"ldap_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"update_vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 11),
				Optional:     true,
				Computed:     true,
			},
			"ldap_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ldap_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Sensitive:    true,
			},
			"scep_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"scep_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"crl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemCertificateCrlCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCertificateCrl(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCrl resource while getting object: %v", err)
	}

	o, err := c.CreateSystemCertificateCrl(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCrl resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCertificateCrl")
	}

	return resourceSystemCertificateCrlRead(d, m)
}

func resourceSystemCertificateCrlUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCertificateCrl(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCrl resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemCertificateCrl(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCrl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCertificateCrl")
	}

	return resourceSystemCertificateCrlRead(d, m)
}

func resourceSystemCertificateCrlDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemCertificateCrl(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateCrl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateCrlRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemCertificateCrl(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCrl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateCrl(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCrl resource from API: %v", err)
	}
	return nil
}

func flattenSystemCertificateCrlName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlHttpUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlLdapUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlUpdateVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlLdapServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlUpdateInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlLdapPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlScepCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlScepUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCertificateCrlCrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateCrl(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemCertificateCrlName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("http_url", flattenSystemCertificateCrlHttpUrl(o["http-url"], d, "http_url", sv)); err != nil {
		if !fortiAPIPatch(o["http-url"]) {
			return fmt.Errorf("Error reading http_url: %v", err)
		}
	}

	if err = d.Set("ldap_username", flattenSystemCertificateCrlLdapUsername(o["ldap-username"], d, "ldap_username", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-username"]) {
			return fmt.Errorf("Error reading ldap_username: %v", err)
		}
	}

	if err = d.Set("update_vdom", flattenSystemCertificateCrlUpdateVdom(o["update-vdom"], d, "update_vdom", sv)); err != nil {
		if !fortiAPIPatch(o["update-vdom"]) {
			return fmt.Errorf("Error reading update_vdom: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenSystemCertificateCrlLdapServer(o["ldap-server"], d, "ldap_server", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenSystemCertificateCrlUpdateInterval(o["update-interval"], d, "update_interval", sv)); err != nil {
		if !fortiAPIPatch(o["update-interval"]) {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("scep_cert", flattenSystemCertificateCrlScepCert(o["scep-cert"], d, "scep_cert", sv)); err != nil {
		if !fortiAPIPatch(o["scep-cert"]) {
			return fmt.Errorf("Error reading scep_cert: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenSystemCertificateCrlScepUrl(o["scep-url"], d, "scep_url", sv)); err != nil {
		if !fortiAPIPatch(o["scep-url"]) {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("crl", flattenSystemCertificateCrlCrl(o["crl"], d, "crl", sv)); err != nil {
		if !fortiAPIPatch(o["crl"]) {
			return fmt.Errorf("Error reading crl: %v", err)
		}
	}

	return nil
}

func flattenSystemCertificateCrlFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemCertificateCrlName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlHttpUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlLdapUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlUpdateVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlUpdateInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlLdapPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlScepCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlScepUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCertificateCrlCrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateCrl(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemCertificateCrlName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("http_url"); ok {

		t, err := expandSystemCertificateCrlHttpUrl(d, v, "http_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-url"] = t
		}
	}

	if v, ok := d.GetOk("ldap_username"); ok {

		t, err := expandSystemCertificateCrlLdapUsername(d, v, "ldap_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-username"] = t
		}
	}

	if v, ok := d.GetOk("update_vdom"); ok {

		t, err := expandSystemCertificateCrlUpdateVdom(d, v, "update_vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-vdom"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok {

		t, err := expandSystemCertificateCrlLdapServer(d, v, "ldap_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("update_interval"); ok {

		t, err := expandSystemCertificateCrlUpdateInterval(d, v, "update_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-interval"] = t
		}
	}

	if v, ok := d.GetOk("ldap_password"); ok {

		t, err := expandSystemCertificateCrlLdapPassword(d, v, "ldap_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-password"] = t
		}
	}

	if v, ok := d.GetOk("scep_cert"); ok {

		t, err := expandSystemCertificateCrlScepCert(d, v, "scep_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-cert"] = t
		}
	}

	if v, ok := d.GetOk("scep_url"); ok {

		t, err := expandSystemCertificateCrlScepUrl(d, v, "scep_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	}

	if v, ok := d.GetOk("crl"); ok {

		t, err := expandSystemCertificateCrlCrl(d, v, "crl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crl"] = t
		}
	}

	return &obj, nil
}
