// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Ntp system info configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNtpUpdate,
		Read:   resourceSystemNtpRead,
		Update: resourceSystemNtpUpdate,
		Delete: resourceSystemNtpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"key_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_time_adjustments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_unsync_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 59),
				Optional:     true,
			},
			"ntpsync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"syncinterval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),
				Optional:     true,
				Computed:     true,
			},
			"ntpserver": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ntpv3": &schema.Schema{
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
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 59),
							Optional:     true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"key_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"key_id": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSystemNtpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemNtp(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemNtp resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemNtp(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemNtp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemNtp")
	}

	return resourceSystemNtpRead(d, m)
}

func resourceSystemNtpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemNtp(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemNtp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNtp(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemNtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNtpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemNtp(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemNtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNtp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemNtp resource from API: %v", err)
	}
	return nil
}

func flattenSystemNtpKeyType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpLogTimeAdjustments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpServerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpAllowUnsyncSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpSourceIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpNtpsync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpSyncinterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpNtpserver(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ntpv3"
		if _, ok := i["ntpv3"]; ok {

			tmp["ntpv3"] = flattenSystemNtpNtpserverNtpv3(i["ntpv3"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {

			tmp["server"] = flattenSystemNtpNtpserverServer(i["server"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {

			tmp["authentication"] = flattenSystemNtpNtpserverAuthentication(i["authentication"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {

			tmp["key"] = flattenSystemNtpNtpserverKey(i["key"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSystemNtpNtpserverId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_id"
		if _, ok := i["key-id"]; ok {

			tmp["key_id"] = flattenSystemNtpNtpserverKeyId(i["key-id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemNtpNtpserverNtpv3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpNtpserverServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpNtpserverAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpNtpserverKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpNtpserverId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpNtpserverKeyId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpKeyId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemNtp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("key_type", flattenSystemNtpKeyType(o["key-type"], d, "key_type", sv)); err != nil {
		if !fortiAPIPatch(o["key-type"]) {
			return fmt.Errorf("Error reading key_type: %v", err)
		}
	}

	if err = d.Set("log_time_adjustments", flattenSystemNtpLogTimeAdjustments(o["log-time-adjustments"], d, "log_time_adjustments", sv)); err != nil {
		if !fortiAPIPatch(o["log-time-adjustments"]) {
			return fmt.Errorf("Error reading log_time_adjustments: %v", err)
		}
	}

	if err = d.Set("server_mode", flattenSystemNtpServerMode(o["server-mode"], d, "server_mode", sv)); err != nil {
		if !fortiAPIPatch(o["server-mode"]) {
			return fmt.Errorf("Error reading server_mode: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemNtpSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("allow_unsync_source", flattenSystemNtpAllowUnsyncSource(o["allow-unsync-source"], d, "allow_unsync_source", sv)); err != nil {
		if !fortiAPIPatch(o["allow-unsync-source"]) {
			return fmt.Errorf("Error reading allow_unsync_source: %v", err)
		}
	}

	if err = d.Set("authentication", flattenSystemNtpAuthentication(o["authentication"], d, "authentication", sv)); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenSystemNtpSourceIp6(o["source-ip6"], d, "source_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("key", flattenSystemNtpKey(o["key"], d, "key", sv)); err != nil {
		if !fortiAPIPatch(o["key"]) {
			return fmt.Errorf("Error reading key: %v", err)
		}
	}

	if err = d.Set("ntpsync", flattenSystemNtpNtpsync(o["ntpsync"], d, "ntpsync", sv)); err != nil {
		if !fortiAPIPatch(o["ntpsync"]) {
			return fmt.Errorf("Error reading ntpsync: %v", err)
		}
	}

	if err = d.Set("syncinterval", flattenSystemNtpSyncinterval(o["syncinterval"], d, "syncinterval", sv)); err != nil {
		if !fortiAPIPatch(o["syncinterval"]) {
			return fmt.Errorf("Error reading syncinterval: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ntpserver", flattenSystemNtpNtpserver(o["ntpserver"], d, "ntpserver", sv)); err != nil {
			if !fortiAPIPatch(o["ntpserver"]) {
				return fmt.Errorf("Error reading ntpserver: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ntpserver"); ok {
			if err = d.Set("ntpserver", flattenSystemNtpNtpserver(o["ntpserver"], d, "ntpserver", sv)); err != nil {
				if !fortiAPIPatch(o["ntpserver"]) {
					return fmt.Errorf("Error reading ntpserver: %v", err)
				}
			}
		}
	}

	if err = d.Set("key_id", flattenSystemNtpKeyId(o["key-id"], d, "key_id", sv)); err != nil {
		if !fortiAPIPatch(o["key-id"]) {
			return fmt.Errorf("Error reading key_id: %v", err)
		}
	}

	return nil
}

func flattenSystemNtpFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemNtpKeyType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpLogTimeAdjustments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpServerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpAllowUnsyncSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpSourceIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpsync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpSyncinterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserver(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ntpv3"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ntpv3"], _ = expandSystemNtpNtpserverNtpv3(d, i["ntpv3"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["server"], _ = expandSystemNtpNtpserverServer(d, i["server"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["authentication"], _ = expandSystemNtpNtpserverAuthentication(d, i["authentication"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["key"], _ = expandSystemNtpNtpserverKey(d, i["key"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSystemNtpNtpserverId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["key-id"], _ = expandSystemNtpNtpserverKeyId(d, i["key_id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNtpNtpserverNtpv3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverKeyId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpKeyId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNtp(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("key_type"); ok {
		if setArgNil {
			obj["key-type"] = nil
		} else {

			t, err := expandSystemNtpKeyType(d, v, "key_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["key-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("log_time_adjustments"); ok {
		if setArgNil {
			obj["log-time-adjustments"] = nil
		} else {

			t, err := expandSystemNtpLogTimeAdjustments(d, v, "log_time_adjustments", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-time-adjustments"] = t
			}
		}
	}

	if v, ok := d.GetOk("server_mode"); ok {
		if setArgNil {
			obj["server-mode"] = nil
		} else {

			t, err := expandSystemNtpServerMode(d, v, "server_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {

			t, err := expandSystemNtpSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_unsync_source"); ok {
		if setArgNil {
			obj["allow-unsync-source"] = nil
		} else {

			t, err := expandSystemNtpAllowUnsyncSource(d, v, "allow_unsync_source", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-unsync-source"] = t
			}
		}
	}

	if v, ok := d.GetOk("authentication"); ok {
		if setArgNil {
			obj["authentication"] = nil
		} else {

			t, err := expandSystemNtpAuthentication(d, v, "authentication", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["authentication"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {
		if setArgNil {
			obj["source-ip6"] = nil
		} else {

			t, err := expandSystemNtpSourceIp6(d, v, "source_ip6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip6"] = t
			}
		}
	}

	if v, ok := d.GetOk("key"); ok {
		if setArgNil {
			obj["key"] = nil
		} else {

			t, err := expandSystemNtpKey(d, v, "key", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["key"] = t
			}
		}
	}

	if v, ok := d.GetOk("ntpsync"); ok {
		if setArgNil {
			obj["ntpsync"] = nil
		} else {

			t, err := expandSystemNtpNtpsync(d, v, "ntpsync", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ntpsync"] = t
			}
		}
	}

	if v, ok := d.GetOk("syncinterval"); ok {
		if setArgNil {
			obj["syncinterval"] = nil
		} else {

			t, err := expandSystemNtpSyncinterval(d, v, "syncinterval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["syncinterval"] = t
			}
		}
	}

	if v, ok := d.GetOk("ntpserver"); ok || d.HasChange("ntpserver") {
		if setArgNil {
			obj["ntpserver"] = make([]struct{}, 0)
		} else {

			t, err := expandSystemNtpNtpserver(d, v, "ntpserver", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ntpserver"] = t
			}
		}
	}

	if v, ok := d.GetOk("key_id"); ok {
		if setArgNil {
			obj["key-id"] = nil
		} else {

			t, err := expandSystemNtpKeyId(d, v, "key_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["key-id"] = t
			}
		}
	}

	return &obj, nil
}
