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

func resourceSystemFortianalyzer3() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFortianalyzer3Update,
		Read:   resourceSystemFortianalyzer3Read,
		Update: resourceSystemFortianalyzer3Update,
		Delete: resourceSystemFortianalyzer3Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"__change_ip": &schema.Schema{
				Type:     schema.TypeInt,
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
			"fdp_device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Optional:     true,
				Computed:     true,
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
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemFortianalyzer3Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemFortianalyzer3(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortianalyzer3 resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFortianalyzer3(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortianalyzer3 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFortianalyzer3")
	}

	return resourceSystemFortianalyzer3Read(d, m)
}

func resourceSystemFortianalyzer3Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemFortianalyzer3(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemFortianalyzer3 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFortianalyzer3(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemFortianalyzer3 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFortianalyzer3Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemFortianalyzer3(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortianalyzer3 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFortianalyzer3(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortianalyzer3 resource from API: %v", err)
	}
	return nil
}

func flattenSystemFortianalyzer3Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortianalyzer3__Change_Ip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortianalyzer3Encrypt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortianalyzer3Localid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortianalyzer3FdpDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortianalyzer3ConnTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortianalyzer3Psksecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortianalyzer3MgmtName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortianalyzer3FdpInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortianalyzer3Server(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortianalyzer3AddressMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemFortianalyzer3(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemFortianalyzer3Status(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("__change_ip", flattenSystemFortianalyzer3__Change_Ip(o["__change_ip"], d, "__change_ip", sv)); err != nil {
		if !fortiAPIPatch(o["__change_ip"]) {
			return fmt.Errorf("Error reading __change_ip: %v", err)
		}
	}

	if err = d.Set("encrypt", flattenSystemFortianalyzer3Encrypt(o["encrypt"], d, "encrypt", sv)); err != nil {
		if !fortiAPIPatch(o["encrypt"]) {
			return fmt.Errorf("Error reading encrypt: %v", err)
		}
	}

	if err = d.Set("localid", flattenSystemFortianalyzer3Localid(o["localid"], d, "localid", sv)); err != nil {
		if !fortiAPIPatch(o["localid"]) {
			return fmt.Errorf("Error reading localid: %v", err)
		}
	}

	if err = d.Set("fdp_device", flattenSystemFortianalyzer3FdpDevice(o["fdp-device"], d, "fdp_device", sv)); err != nil {
		if !fortiAPIPatch(o["fdp-device"]) {
			return fmt.Errorf("Error reading fdp_device: %v", err)
		}
	}

	if err = d.Set("conn_timeout", flattenSystemFortianalyzer3ConnTimeout(o["conn-timeout"], d, "conn_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["conn-timeout"]) {
			return fmt.Errorf("Error reading conn_timeout: %v", err)
		}
	}

	if err = d.Set("mgmt_name", flattenSystemFortianalyzer3MgmtName(o["mgmt-name"], d, "mgmt_name", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-name"]) {
			return fmt.Errorf("Error reading mgmt_name: %v", err)
		}
	}

	if err = d.Set("fdp_interface", flattenSystemFortianalyzer3FdpInterface(o["fdp-interface"], d, "fdp_interface", sv)); err != nil {
		if !fortiAPIPatch(o["fdp-interface"]) {
			return fmt.Errorf("Error reading fdp_interface: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemFortianalyzer3Server(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("address_mode", flattenSystemFortianalyzer3AddressMode(o["address-mode"], d, "address_mode", sv)); err != nil {
		if !fortiAPIPatch(o["address-mode"]) {
			return fmt.Errorf("Error reading address_mode: %v", err)
		}
	}

	return nil
}

func flattenSystemFortianalyzer3FortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemFortianalyzer3Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortianalyzer3__Change_Ip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortianalyzer3Encrypt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortianalyzer3Localid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortianalyzer3FdpDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortianalyzer3ConnTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortianalyzer3Psksecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortianalyzer3MgmtName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortianalyzer3FdpInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortianalyzer3Server(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortianalyzer3AddressMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFortianalyzer3(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandSystemFortianalyzer3Status(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("__change_ip"); ok {
		if setArgNil {
			obj["__change_ip"] = nil
		} else {

			t, err := expandSystemFortianalyzer3__Change_Ip(d, v, "__change_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["__change_ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("encrypt"); ok {
		if setArgNil {
			obj["encrypt"] = nil
		} else {

			t, err := expandSystemFortianalyzer3Encrypt(d, v, "encrypt", sv)
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

			t, err := expandSystemFortianalyzer3Localid(d, v, "localid", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["localid"] = t
			}
		}
	}

	if v, ok := d.GetOk("fdp_device"); ok {
		if setArgNil {
			obj["fdp-device"] = nil
		} else {

			t, err := expandSystemFortianalyzer3FdpDevice(d, v, "fdp_device", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fdp-device"] = t
			}
		}
	}

	if v, ok := d.GetOk("conn_timeout"); ok {
		if setArgNil {
			obj["conn-timeout"] = nil
		} else {

			t, err := expandSystemFortianalyzer3ConnTimeout(d, v, "conn_timeout", sv)
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

			t, err := expandSystemFortianalyzer3Psksecret(d, v, "psksecret", sv)
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

			t, err := expandSystemFortianalyzer3MgmtName(d, v, "mgmt_name", sv)
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

			t, err := expandSystemFortianalyzer3FdpInterface(d, v, "fdp_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fdp-interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("server"); ok {
		if setArgNil {
			obj["server"] = nil
		} else {

			t, err := expandSystemFortianalyzer3Server(d, v, "server", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server"] = t
			}
		}
	}

	if v, ok := d.GetOk("address_mode"); ok {
		if setArgNil {
			obj["address-mode"] = nil
		} else {

			t, err := expandSystemFortianalyzer3AddressMode(d, v, "address_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["address-mode"] = t
			}
		}
	}

	return &obj, nil
}
