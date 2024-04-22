// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Alert e-mail mail server configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAlertemail() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAlertemailUpdate,
		Read:   resourceSystemAlertemailRead,
		Update: resourceSystemAlertemailUpdate,
		Delete: resourceSystemAlertemailDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"authenticate": &schema.Schema{
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
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Sensitive:    true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAlertemailUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAlertemail(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAlertemail resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAlertemail(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAlertemail resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAlertemail")
	}

	return resourceSystemAlertemailRead(d, m)
}

func resourceSystemAlertemailDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAlertemail(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemAlertemail resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAlertemail(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemAlertemail resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAlertemailRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAlertemail(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAlertemail resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAlertemail(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAlertemail resource from API: %v", err)
	}
	return nil
}

func flattenSystemAlertemailUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAlertemailAuthenticate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAlertemailSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAlertemailServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAlertemailPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAlertemailPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAlertemail(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("username", flattenSystemAlertemailUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("authenticate", flattenSystemAlertemailAuthenticate(o["authenticate"], d, "authenticate", sv)); err != nil {
		if !fortiAPIPatch(o["authenticate"]) {
			return fmt.Errorf("Error reading authenticate: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemAlertemailSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemAlertemailServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemAlertemailPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	return nil
}

func flattenSystemAlertemailFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemAlertemailUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertemailAuthenticate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertemailSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertemailServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertemailPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAlertemailPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAlertemail(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("username"); ok {
		if setArgNil {
			obj["username"] = nil
		} else {

			t, err := expandSystemAlertemailUsername(d, v, "username", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["username"] = t
			}
		}
	}

	if v, ok := d.GetOk("authenticate"); ok {
		if setArgNil {
			obj["authenticate"] = nil
		} else {

			t, err := expandSystemAlertemailAuthenticate(d, v, "authenticate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["authenticate"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {

			t, err := expandSystemAlertemailSourceIp(d, v, "source_ip", sv)
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

			t, err := expandSystemAlertemailServer(d, v, "server", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server"] = t
			}
		}
	}

	if v, ok := d.GetOk("password"); ok {
		if setArgNil {
			obj["password"] = nil
		} else {

			t, err := expandSystemAlertemailPassword(d, v, "password", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["password"] = t
			}
		}
	}

	if v, ok := d.GetOk("port"); ok {
		if setArgNil {
			obj["port"] = nil
		} else {

			t, err := expandSystemAlertemailPort(d, v, "port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["port"] = t
			}
		}
	}

	return &obj, nil
}
