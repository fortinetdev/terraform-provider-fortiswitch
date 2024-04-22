// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: LDAP server entry configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserLdap() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserLdapCreate,
		Read:   resourceUserLdapRead,
		Update: resourceUserLdapUpdate,
		Delete: resourceUserLdapDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"dn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"password_expiry_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_member_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"ca_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"group_object_filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"cnid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 20),
				Optional:     true,
				Computed:     true,
			},
			"member_attr": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"password_renewal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Sensitive:    true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserLdapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserLdap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserLdap resource while getting object: %v", err)
	}

	o, err := c.CreateUserLdap(obj)

	if err != nil {
		return fmt.Errorf("Error creating UserLdap resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserLdap")
	}

	return resourceUserLdapRead(d, m)
}

func resourceUserLdapUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserLdap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserLdap resource while getting object: %v", err)
	}

	o, err := c.UpdateUserLdap(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating UserLdap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserLdap")
	}

	return resourceUserLdapRead(d, m)
}

func resourceUserLdapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteUserLdap(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting UserLdap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserLdapRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadUserLdap(mkey)
	if err != nil {
		return fmt.Errorf("Error reading UserLdap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserLdap(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserLdap resource from API: %v", err)
	}
	return nil
}

func flattenUserLdapDn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapPasswordExpiryWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapGroupMemberCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapCaCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapGroupObjectFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapCnid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapMemberAttr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapPasswordRenewal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapSecure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserLdap(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("dn", flattenUserLdapDn(o["dn"], d, "dn", sv)); err != nil {
		if !fortiAPIPatch(o["dn"]) {
			return fmt.Errorf("Error reading dn: %v", err)
		}
	}

	if err = d.Set("username", flattenUserLdapUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("name", flattenUserLdapName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("password_expiry_warning", flattenUserLdapPasswordExpiryWarning(o["password-expiry-warning"], d, "password_expiry_warning", sv)); err != nil {
		if !fortiAPIPatch(o["password-expiry-warning"]) {
			return fmt.Errorf("Error reading password_expiry_warning: %v", err)
		}
	}

	if err = d.Set("group_member_check", flattenUserLdapGroupMemberCheck(o["group-member-check"], d, "group_member_check", sv)); err != nil {
		if !fortiAPIPatch(o["group-member-check"]) {
			return fmt.Errorf("Error reading group_member_check: %v", err)
		}
	}

	if err = d.Set("server", flattenUserLdapServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("ca_cert", flattenUserLdapCaCert(o["ca-cert"], d, "ca_cert", sv)); err != nil {
		if !fortiAPIPatch(o["ca-cert"]) {
			return fmt.Errorf("Error reading ca_cert: %v", err)
		}
	}

	if err = d.Set("group_object_filter", flattenUserLdapGroupObjectFilter(o["group-object-filter"], d, "group_object_filter", sv)); err != nil {
		if !fortiAPIPatch(o["group-object-filter"]) {
			return fmt.Errorf("Error reading group_object_filter: %v", err)
		}
	}

	if err = d.Set("cnid", flattenUserLdapCnid(o["cnid"], d, "cnid", sv)); err != nil {
		if !fortiAPIPatch(o["cnid"]) {
			return fmt.Errorf("Error reading cnid: %v", err)
		}
	}

	if err = d.Set("member_attr", flattenUserLdapMemberAttr(o["member-attr"], d, "member_attr", sv)); err != nil {
		if !fortiAPIPatch(o["member-attr"]) {
			return fmt.Errorf("Error reading member_attr: %v", err)
		}
	}

	if err = d.Set("password_renewal", flattenUserLdapPasswordRenewal(o["password-renewal"], d, "password_renewal", sv)); err != nil {
		if !fortiAPIPatch(o["password-renewal"]) {
			return fmt.Errorf("Error reading password_renewal: %v", err)
		}
	}

	if err = d.Set("type", flattenUserLdapType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("port", flattenUserLdapPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("secure", flattenUserLdapSecure(o["secure"], d, "secure", sv)); err != nil {
		if !fortiAPIPatch(o["secure"]) {
			return fmt.Errorf("Error reading secure: %v", err)
		}
	}

	return nil
}

func flattenUserLdapFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandUserLdapDn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPasswordExpiryWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupMemberCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapCaCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupObjectFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapCnid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapMemberAttr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPasswordRenewal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSecure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserLdap(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dn"); ok {

		t, err := expandUserLdapDn(d, v, "dn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dn"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {

		t, err := expandUserLdapUsername(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandUserLdapName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password_expiry_warning"); ok {

		t, err := expandUserLdapPasswordExpiryWarning(d, v, "password_expiry_warning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-expiry-warning"] = t
		}
	}

	if v, ok := d.GetOk("group_member_check"); ok {

		t, err := expandUserLdapGroupMemberCheck(d, v, "group_member_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-member-check"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {

		t, err := expandUserLdapServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("ca_cert"); ok {

		t, err := expandUserLdapCaCert(d, v, "ca_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("group_object_filter"); ok {

		t, err := expandUserLdapGroupObjectFilter(d, v, "group_object_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-object-filter"] = t
		}
	}

	if v, ok := d.GetOk("cnid"); ok {

		t, err := expandUserLdapCnid(d, v, "cnid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cnid"] = t
		}
	}

	if v, ok := d.GetOk("member_attr"); ok {

		t, err := expandUserLdapMemberAttr(d, v, "member_attr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-attr"] = t
		}
	}

	if v, ok := d.GetOk("password_renewal"); ok {

		t, err := expandUserLdapPasswordRenewal(d, v, "password_renewal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-renewal"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {

		t, err := expandUserLdapPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandUserLdapType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {

		t, err := expandUserLdapPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("secure"); ok {

		t, err := expandUserLdapSecure(d, v, "secure", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure"] = t
		}
	}

	return &obj, nil
}
