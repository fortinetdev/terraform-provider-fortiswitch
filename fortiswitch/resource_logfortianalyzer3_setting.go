// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Setting for FortiAnalyzer.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogFortianalyzer3Setting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortianalyzer3SettingUpdate,
		Read:   resourceLogFortianalyzer3SettingRead,
		Update: resourceLogFortianalyzer3SettingUpdate,
		Delete: resourceLogFortianalyzer3SettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"upload_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fdp_device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Optional:     true,
				Computed:     true,
			},
			"upload_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"buffer_max_send": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(20, 20000),
				Optional:     true,
				Computed:     true,
			},
			"address_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encrypt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hmac_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_archive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mgmt_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"fdp_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"__change_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"localid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"max_buffer_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1024),
				Optional:     true,
				Computed:     true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conn_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"psksecret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Sensitive:    true,
			},
			"upload_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogFortianalyzer3SettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogFortianalyzer3Setting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer3Setting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogFortianalyzer3Setting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer3Setting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogFortianalyzer3Setting")
	}

	return resourceLogFortianalyzer3SettingRead(d, m)
}

func resourceLogFortianalyzer3SettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogFortianalyzer3Setting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer3Setting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogFortianalyzer3Setting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing LogFortianalyzer3Setting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortianalyzer3SettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogFortianalyzer3Setting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer3Setting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortianalyzer3Setting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer3Setting resource from API: %v", err)
	}
	return nil
}

func flattenLogFortianalyzer3SettingUploadInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingFdpDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingUploadOption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingBufferMaxSend(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingAddressMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingEncrypt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingHmacAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingUploadDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingIpsArchive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingMgmtName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingFdpInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3Setting__Change_Ip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingLocalid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingMaxBufferSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingConnTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingPsksecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3SettingUploadTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogFortianalyzer3Setting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("upload_interval", flattenLogFortianalyzer3SettingUploadInterval(o["upload-interval"], d, "upload_interval", sv)); err != nil {
		if !fortiAPIPatch(o["upload-interval"]) {
			return fmt.Errorf("Error reading upload_interval: %v", err)
		}
	}

	if err = d.Set("fdp_device", flattenLogFortianalyzer3SettingFdpDevice(o["fdp-device"], d, "fdp_device", sv)); err != nil {
		if !fortiAPIPatch(o["fdp-device"]) {
			return fmt.Errorf("Error reading fdp_device: %v", err)
		}
	}

	if err = d.Set("upload_option", flattenLogFortianalyzer3SettingUploadOption(o["upload-option"], d, "upload_option", sv)); err != nil {
		if !fortiAPIPatch(o["upload-option"]) {
			return fmt.Errorf("Error reading upload_option: %v", err)
		}
	}

	if err = d.Set("buffer_max_send", flattenLogFortianalyzer3SettingBufferMaxSend(o["buffer-max-send"], d, "buffer_max_send", sv)); err != nil {
		if !fortiAPIPatch(o["buffer-max-send"]) {
			return fmt.Errorf("Error reading buffer_max_send: %v", err)
		}
	}

	if err = d.Set("address_mode", flattenLogFortianalyzer3SettingAddressMode(o["address-mode"], d, "address_mode", sv)); err != nil {
		if !fortiAPIPatch(o["address-mode"]) {
			return fmt.Errorf("Error reading address_mode: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenLogFortianalyzer3SettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("encrypt", flattenLogFortianalyzer3SettingEncrypt(o["encrypt"], d, "encrypt", sv)); err != nil {
		if !fortiAPIPatch(o["encrypt"]) {
			return fmt.Errorf("Error reading encrypt: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogFortianalyzer3SettingSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("override", flattenLogFortianalyzer3SettingOverride(o["override"], d, "override", sv)); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("hmac_algorithm", flattenLogFortianalyzer3SettingHmacAlgorithm(o["hmac-algorithm"], d, "hmac_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["hmac-algorithm"]) {
			return fmt.Errorf("Error reading hmac_algorithm: %v", err)
		}
	}

	if err = d.Set("status", flattenLogFortianalyzer3SettingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upload_day", flattenLogFortianalyzer3SettingUploadDay(o["upload-day"], d, "upload_day", sv)); err != nil {
		if !fortiAPIPatch(o["upload-day"]) {
			return fmt.Errorf("Error reading upload_day: %v", err)
		}
	}

	if err = d.Set("ips_archive", flattenLogFortianalyzer3SettingIpsArchive(o["ips-archive"], d, "ips_archive", sv)); err != nil {
		if !fortiAPIPatch(o["ips-archive"]) {
			return fmt.Errorf("Error reading ips_archive: %v", err)
		}
	}

	if err = d.Set("mgmt_name", flattenLogFortianalyzer3SettingMgmtName(o["mgmt-name"], d, "mgmt_name", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-name"]) {
			return fmt.Errorf("Error reading mgmt_name: %v", err)
		}
	}

	if err = d.Set("fdp_interface", flattenLogFortianalyzer3SettingFdpInterface(o["fdp-interface"], d, "fdp_interface", sv)); err != nil {
		if !fortiAPIPatch(o["fdp-interface"]) {
			return fmt.Errorf("Error reading fdp_interface: %v", err)
		}
	}

	if err = d.Set("__change_ip", flattenLogFortianalyzer3Setting__Change_Ip(o["__change_ip"], d, "__change_ip", sv)); err != nil {
		if !fortiAPIPatch(o["__change_ip"]) {
			return fmt.Errorf("Error reading __change_ip: %v", err)
		}
	}

	if err = d.Set("localid", flattenLogFortianalyzer3SettingLocalid(o["localid"], d, "localid", sv)); err != nil {
		if !fortiAPIPatch(o["localid"]) {
			return fmt.Errorf("Error reading localid: %v", err)
		}
	}

	if err = d.Set("max_buffer_size", flattenLogFortianalyzer3SettingMaxBufferSize(o["max-buffer-size"], d, "max_buffer_size", sv)); err != nil {
		if !fortiAPIPatch(o["max-buffer-size"]) {
			return fmt.Errorf("Error reading max_buffer_size: %v", err)
		}
	}

	if err = d.Set("server", flattenLogFortianalyzer3SettingServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("conn_timeout", flattenLogFortianalyzer3SettingConnTimeout(o["conn-timeout"], d, "conn_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["conn-timeout"]) {
			return fmt.Errorf("Error reading conn_timeout: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenLogFortianalyzer3SettingUploadTime(o["upload-time"], d, "upload_time", sv)); err != nil {
		if !fortiAPIPatch(o["upload-time"]) {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	return nil
}

func flattenLogFortianalyzer3SettingFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandLogFortianalyzer3SettingUploadInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingFdpDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingUploadOption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingBufferMaxSend(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingAddressMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingEncrypt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingHmacAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingUploadDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingIpsArchive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingMgmtName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingFdpInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3Setting__Change_Ip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingLocalid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingMaxBufferSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingConnTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingPsksecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3SettingUploadTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortianalyzer3Setting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("upload_interval"); ok {
		if setArgNil {
			obj["upload-interval"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingUploadInterval(d, v, "upload_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("fdp_device"); ok {
		if setArgNil {
			obj["fdp-device"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingFdpDevice(d, v, "fdp_device", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fdp-device"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_option"); ok {
		if setArgNil {
			obj["upload-option"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingUploadOption(d, v, "upload_option", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-option"] = t
			}
		}
	}

	if v, ok := d.GetOk("buffer_max_send"); ok {
		if setArgNil {
			obj["buffer-max-send"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingBufferMaxSend(d, v, "buffer_max_send", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["buffer-max-send"] = t
			}
		}
	}

	if v, ok := d.GetOk("address_mode"); ok {
		if setArgNil {
			obj["address-mode"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingAddressMode(d, v, "address_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["address-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {
		if setArgNil {
			obj["enc-algorithm"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingEncAlgorithm(d, v, "enc_algorithm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["enc-algorithm"] = t
			}
		}
	}

	if v, ok := d.GetOk("encrypt"); ok {
		if setArgNil {
			obj["encrypt"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingEncrypt(d, v, "encrypt", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["encrypt"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("override"); ok {
		if setArgNil {
			obj["override"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingOverride(d, v, "override", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["override"] = t
			}
		}
	}

	if v, ok := d.GetOk("hmac_algorithm"); ok {
		if setArgNil {
			obj["hmac-algorithm"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingHmacAlgorithm(d, v, "hmac_algorithm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hmac-algorithm"] = t
			}
		}
	}

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_day"); ok {
		if setArgNil {
			obj["upload-day"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingUploadDay(d, v, "upload_day", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-day"] = t
			}
		}
	}

	if v, ok := d.GetOk("ips_archive"); ok {
		if setArgNil {
			obj["ips-archive"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingIpsArchive(d, v, "ips_archive", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ips-archive"] = t
			}
		}
	}

	if v, ok := d.GetOk("mgmt_name"); ok {
		if setArgNil {
			obj["mgmt-name"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingMgmtName(d, v, "mgmt_name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mgmt-name"] = t
			}
		}
	}

	if v, ok := d.GetOk("fdp_interface"); ok {
		if setArgNil {
			obj["fdp-interface"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingFdpInterface(d, v, "fdp_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fdp-interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("__change_ip"); ok {
		if setArgNil {
			obj["__change_ip"] = nil
		} else {

			t, err := expandLogFortianalyzer3Setting__Change_Ip(d, v, "__change_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["__change_ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("localid"); ok {
		if setArgNil {
			obj["localid"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingLocalid(d, v, "localid", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["localid"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("max_buffer_size"); ok {
		if setArgNil {
			obj["max-buffer-size"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingMaxBufferSize(d, v, "max_buffer_size", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-buffer-size"] = t
			}
		}
	}

	if v, ok := d.GetOk("server"); ok {
		if setArgNil {
			obj["server"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingServer(d, v, "server", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server"] = t
			}
		}
	}

	if v, ok := d.GetOk("conn_timeout"); ok {
		if setArgNil {
			obj["conn-timeout"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingConnTimeout(d, v, "conn_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["conn-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("psksecret"); ok {
		if setArgNil {
			obj["psksecret"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingPsksecret(d, v, "psksecret", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["psksecret"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_time"); ok {
		if setArgNil {
			obj["upload-time"] = nil
		} else {

			t, err := expandLogFortianalyzer3SettingUploadTime(d, v, "upload_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-time"] = t
			}
		}
	}

	return &obj, nil
}
