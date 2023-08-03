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

func resourceLogFortianalyzerOverrideSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortianalyzerOverrideSettingUpdate,
		Read:   resourceLogFortianalyzerOverrideSettingRead,
		Update: resourceLogFortianalyzerOverrideSettingUpdate,
		Delete: resourceLogFortianalyzerOverrideSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hmac_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encrypt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"localid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"__change_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ips_archive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"max_buffer_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1024),
				Optional:     true,
				Computed:     true,
			},
			"psksecret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
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
			"override": &schema.Schema{
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
			"fdp_device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Optional:     true,
				Computed:     true,
			},
			"address_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogFortianalyzerOverrideSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogFortianalyzerOverrideSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerOverrideSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogFortianalyzerOverrideSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerOverrideSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogFortianalyzerOverrideSetting")
	}

	return resourceLogFortianalyzerOverrideSettingRead(d, m)
}

func resourceLogFortianalyzerOverrideSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogFortianalyzerOverrideSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerOverrideSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogFortianalyzerOverrideSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing LogFortianalyzerOverrideSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortianalyzerOverrideSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogFortianalyzerOverrideSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzerOverrideSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortianalyzerOverrideSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzerOverrideSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogFortianalyzerOverrideSettingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingHmacAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingEncrypt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingLocalid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSetting__Change_Ip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingIpsArchive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingConnTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingMaxBufferSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingPsksecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingMgmtName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingFdpInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingBufferMaxSend(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingFdpDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingAddressMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogFortianalyzerOverrideSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenLogFortianalyzerOverrideSettingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenLogFortianalyzerOverrideSettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("hmac_algorithm", flattenLogFortianalyzerOverrideSettingHmacAlgorithm(o["hmac-algorithm"], d, "hmac_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["hmac-algorithm"]) {
			return fmt.Errorf("Error reading hmac_algorithm: %v", err)
		}
	}

	if err = d.Set("encrypt", flattenLogFortianalyzerOverrideSettingEncrypt(o["encrypt"], d, "encrypt", sv)); err != nil {
		if !fortiAPIPatch(o["encrypt"]) {
			return fmt.Errorf("Error reading encrypt: %v", err)
		}
	}

	if err = d.Set("localid", flattenLogFortianalyzerOverrideSettingLocalid(o["localid"], d, "localid", sv)); err != nil {
		if !fortiAPIPatch(o["localid"]) {
			return fmt.Errorf("Error reading localid: %v", err)
		}
	}

	if err = d.Set("__change_ip", flattenLogFortianalyzerOverrideSetting__Change_Ip(o["__change_ip"], d, "__change_ip", sv)); err != nil {
		if !fortiAPIPatch(o["__change_ip"]) {
			return fmt.Errorf("Error reading __change_ip: %v", err)
		}
	}

	if err = d.Set("ips_archive", flattenLogFortianalyzerOverrideSettingIpsArchive(o["ips-archive"], d, "ips_archive", sv)); err != nil {
		if !fortiAPIPatch(o["ips-archive"]) {
			return fmt.Errorf("Error reading ips_archive: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogFortianalyzerOverrideSettingSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("server", flattenLogFortianalyzerOverrideSettingServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("conn_timeout", flattenLogFortianalyzerOverrideSettingConnTimeout(o["conn-timeout"], d, "conn_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["conn-timeout"]) {
			return fmt.Errorf("Error reading conn_timeout: %v", err)
		}
	}

	if err = d.Set("max_buffer_size", flattenLogFortianalyzerOverrideSettingMaxBufferSize(o["max-buffer-size"], d, "max_buffer_size", sv)); err != nil {
		if !fortiAPIPatch(o["max-buffer-size"]) {
			return fmt.Errorf("Error reading max_buffer_size: %v", err)
		}
	}

	if err = d.Set("psksecret", flattenLogFortianalyzerOverrideSettingPsksecret(o["psksecret"], d, "psksecret", sv)); err != nil {
		if !fortiAPIPatch(o["psksecret"]) {
			return fmt.Errorf("Error reading psksecret: %v", err)
		}
	}

	if err = d.Set("mgmt_name", flattenLogFortianalyzerOverrideSettingMgmtName(o["mgmt-name"], d, "mgmt_name", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-name"]) {
			return fmt.Errorf("Error reading mgmt_name: %v", err)
		}
	}

	if err = d.Set("fdp_interface", flattenLogFortianalyzerOverrideSettingFdpInterface(o["fdp-interface"], d, "fdp_interface", sv)); err != nil {
		if !fortiAPIPatch(o["fdp-interface"]) {
			return fmt.Errorf("Error reading fdp_interface: %v", err)
		}
	}

	if err = d.Set("override", flattenLogFortianalyzerOverrideSettingOverride(o["override"], d, "override", sv)); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("buffer_max_send", flattenLogFortianalyzerOverrideSettingBufferMaxSend(o["buffer-max-send"], d, "buffer_max_send", sv)); err != nil {
		if !fortiAPIPatch(o["buffer-max-send"]) {
			return fmt.Errorf("Error reading buffer_max_send: %v", err)
		}
	}

	if err = d.Set("fdp_device", flattenLogFortianalyzerOverrideSettingFdpDevice(o["fdp-device"], d, "fdp_device", sv)); err != nil {
		if !fortiAPIPatch(o["fdp-device"]) {
			return fmt.Errorf("Error reading fdp_device: %v", err)
		}
	}

	if err = d.Set("address_mode", flattenLogFortianalyzerOverrideSettingAddressMode(o["address-mode"], d, "address_mode", sv)); err != nil {
		if !fortiAPIPatch(o["address-mode"]) {
			return fmt.Errorf("Error reading address_mode: %v", err)
		}
	}

	return nil
}

func flattenLogFortianalyzerOverrideSettingFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandLogFortianalyzerOverrideSettingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingHmacAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingEncrypt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingLocalid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSetting__Change_Ip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingIpsArchive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingConnTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingMaxBufferSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingPsksecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingMgmtName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingFdpInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingBufferMaxSend(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingFdpDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingAddressMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortianalyzerOverrideSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {
		if setArgNil {
			obj["enc-algorithm"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingEncAlgorithm(d, v, "enc_algorithm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["enc-algorithm"] = t
			}
		}
	}

	if v, ok := d.GetOk("hmac_algorithm"); ok {
		if setArgNil {
			obj["hmac-algorithm"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingHmacAlgorithm(d, v, "hmac_algorithm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hmac-algorithm"] = t
			}
		}
	}

	if v, ok := d.GetOk("encrypt"); ok {
		if setArgNil {
			obj["encrypt"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingEncrypt(d, v, "encrypt", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["encrypt"] = t
			}
		}
	}

	if v, ok := d.GetOk("localid"); ok {
		if setArgNil {
			obj["localid"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingLocalid(d, v, "localid", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["localid"] = t
			}
		}
	}

	if v, ok := d.GetOk("__change_ip"); ok {
		if setArgNil {
			obj["__change_ip"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSetting__Change_Ip(d, v, "__change_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["__change_ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("ips_archive"); ok {
		if setArgNil {
			obj["ips-archive"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingIpsArchive(d, v, "ips_archive", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ips-archive"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("server"); ok {
		if setArgNil {
			obj["server"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingServer(d, v, "server", sv)
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

			t, err := expandLogFortianalyzerOverrideSettingConnTimeout(d, v, "conn_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["conn-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("max_buffer_size"); ok {
		if setArgNil {
			obj["max-buffer-size"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingMaxBufferSize(d, v, "max_buffer_size", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-buffer-size"] = t
			}
		}
	}

	if v, ok := d.GetOk("psksecret"); ok {
		if setArgNil {
			obj["psksecret"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingPsksecret(d, v, "psksecret", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["psksecret"] = t
			}
		}
	}

	if v, ok := d.GetOk("mgmt_name"); ok {
		if setArgNil {
			obj["mgmt-name"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingMgmtName(d, v, "mgmt_name", sv)
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

			t, err := expandLogFortianalyzerOverrideSettingFdpInterface(d, v, "fdp_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fdp-interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("override"); ok {
		if setArgNil {
			obj["override"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingOverride(d, v, "override", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["override"] = t
			}
		}
	}

	if v, ok := d.GetOk("buffer_max_send"); ok {
		if setArgNil {
			obj["buffer-max-send"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingBufferMaxSend(d, v, "buffer_max_send", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["buffer-max-send"] = t
			}
		}
	}

	if v, ok := d.GetOk("fdp_device"); ok {
		if setArgNil {
			obj["fdp-device"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingFdpDevice(d, v, "fdp_device", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fdp-device"] = t
			}
		}
	}

	if v, ok := d.GetOk("address_mode"); ok {
		if setArgNil {
			obj["address-mode"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideSettingAddressMode(d, v, "address_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["address-mode"] = t
			}
		}
	}

	return &obj, nil
}
