// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: User authentication setting.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserSettingUpdate,
		Read:   resourceUserSettingRead,
		Update: resourceUserSettingUpdate,
		Delete: resourceUserSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"auth_multi_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 480),
				Optional:     true,
				Computed:     true,
			},
			"auth_ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"auth_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"auth_http_basic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_invalid_max": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"auth_blackout_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),
				Optional:     true,
				Computed:     true,
			},
			"auth_timeout_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_secure_http": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceUserSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateUserSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating UserSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserSetting")
	}

	return resourceUserSettingRead(d, m)
}

func resourceUserSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating UserSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateUserSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing UserSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadUserSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading UserSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserSetting resource from API: %v", err)
	}
	return nil
}

func flattenUserSettingAuthMultiGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthPorts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenUserSettingAuthPortsType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenUserSettingAuthPortsId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenUserSettingAuthPortsPort(i["port"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenUserSettingAuthPortsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthPortsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthPortsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthHttpBasic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthInvalidMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthBlackoutTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthTimeoutType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthSecureHttp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("auth_multi_group", flattenUserSettingAuthMultiGroup(o["auth-multi-group"], d, "auth_multi_group", sv)); err != nil {
		if !fortiAPIPatch(o["auth-multi-group"]) {
			return fmt.Errorf("Error reading auth_multi_group: %v", err)
		}
	}

	if err = d.Set("auth_timeout", flattenUserSettingAuthTimeout(o["auth-timeout"], d, "auth_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["auth-timeout"]) {
			return fmt.Errorf("Error reading auth_timeout: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("auth_ports", flattenUserSettingAuthPorts(o["auth-ports"], d, "auth_ports", sv)); err != nil {
			if !fortiAPIPatch(o["auth-ports"]) {
				return fmt.Errorf("Error reading auth_ports: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("auth_ports"); ok {
			if err = d.Set("auth_ports", flattenUserSettingAuthPorts(o["auth-ports"], d, "auth_ports", sv)); err != nil {
				if !fortiAPIPatch(o["auth-ports"]) {
					return fmt.Errorf("Error reading auth_ports: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth_cert", flattenUserSettingAuthCert(o["auth-cert"], d, "auth_cert", sv)); err != nil {
		if !fortiAPIPatch(o["auth-cert"]) {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_http_basic", flattenUserSettingAuthHttpBasic(o["auth-http-basic"], d, "auth_http_basic", sv)); err != nil {
		if !fortiAPIPatch(o["auth-http-basic"]) {
			return fmt.Errorf("Error reading auth_http_basic: %v", err)
		}
	}

	if err = d.Set("auth_invalid_max", flattenUserSettingAuthInvalidMax(o["auth-invalid-max"], d, "auth_invalid_max", sv)); err != nil {
		if !fortiAPIPatch(o["auth-invalid-max"]) {
			return fmt.Errorf("Error reading auth_invalid_max: %v", err)
		}
	}

	if err = d.Set("auth_blackout_time", flattenUserSettingAuthBlackoutTime(o["auth-blackout-time"], d, "auth_blackout_time", sv)); err != nil {
		if !fortiAPIPatch(o["auth-blackout-time"]) {
			return fmt.Errorf("Error reading auth_blackout_time: %v", err)
		}
	}

	if err = d.Set("auth_timeout_type", flattenUserSettingAuthTimeoutType(o["auth-timeout-type"], d, "auth_timeout_type", sv)); err != nil {
		if !fortiAPIPatch(o["auth-timeout-type"]) {
			return fmt.Errorf("Error reading auth_timeout_type: %v", err)
		}
	}

	if err = d.Set("auth_secure_http", flattenUserSettingAuthSecureHttp(o["auth-secure-http"], d, "auth_secure_http", sv)); err != nil {
		if !fortiAPIPatch(o["auth-secure-http"]) {
			return fmt.Errorf("Error reading auth_secure_http: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenUserSettingAuthType(o["auth-type"], d, "auth_type", sv)); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	return nil
}

func flattenUserSettingFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandUserSettingAuthMultiGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandUserSettingAuthPortsType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandUserSettingAuthPortsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandUserSettingAuthPortsPort(d, i["port"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserSettingAuthPortsType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthPortsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthPortsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthHttpBasic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthInvalidMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthBlackoutTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthTimeoutType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthSecureHttp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_multi_group"); ok {
		if setArgNil {
			obj["auth-multi-group"] = nil
		} else {

			t, err := expandUserSettingAuthMultiGroup(d, v, "auth_multi_group", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-multi-group"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_timeout"); ok {
		if setArgNil {
			obj["auth-timeout"] = nil
		} else {

			t, err := expandUserSettingAuthTimeout(d, v, "auth_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_ports"); ok || d.HasChange("auth_ports") {
		if setArgNil {
			obj["auth-ports"] = make([]struct{}, 0)
		} else {

			t, err := expandUserSettingAuthPorts(d, v, "auth_ports", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-ports"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_cert"); ok {
		if setArgNil {
			obj["auth-cert"] = nil
		} else {

			t, err := expandUserSettingAuthCert(d, v, "auth_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_http_basic"); ok {
		if setArgNil {
			obj["auth-http-basic"] = nil
		} else {

			t, err := expandUserSettingAuthHttpBasic(d, v, "auth_http_basic", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-http-basic"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_invalid_max"); ok {
		if setArgNil {
			obj["auth-invalid-max"] = nil
		} else {

			t, err := expandUserSettingAuthInvalidMax(d, v, "auth_invalid_max", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-invalid-max"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("auth_blackout_time"); ok {
		if setArgNil {
			obj["auth-blackout-time"] = nil
		} else {

			t, err := expandUserSettingAuthBlackoutTime(d, v, "auth_blackout_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-blackout-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_timeout_type"); ok {
		if setArgNil {
			obj["auth-timeout-type"] = nil
		} else {

			t, err := expandUserSettingAuthTimeoutType(d, v, "auth_timeout_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-timeout-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_secure_http"); ok {
		if setArgNil {
			obj["auth-secure-http"] = nil
		} else {

			t, err := expandUserSettingAuthSecureHttp(d, v, "auth_secure_http", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-secure-http"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_type"); ok {
		if setArgNil {
			obj["auth-type"] = nil
		} else {

			t, err := expandUserSettingAuthType(d, v, "auth_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-type"] = t
			}
		}
	}

	return &obj, nil
}
