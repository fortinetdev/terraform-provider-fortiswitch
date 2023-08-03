// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure console.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemConsole() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemConsoleUpdate,
		Read:   resourceSystemConsoleRead,
		Update: resourceSystemConsoleUpdate,
		Delete: resourceSystemConsoleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"hostname_display_length": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(4, 35),
				Optional:     true,
				Computed:     true,
			},
			"output": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"baudrate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemConsoleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemConsole(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemConsole resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemConsole(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemConsole resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemConsole")
	}

	return resourceSystemConsoleRead(d, m)
}

func resourceSystemConsoleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemConsole(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemConsole resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemConsole(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemConsole resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemConsoleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemConsole(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemConsole resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemConsole(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemConsole resource from API: %v", err)
	}
	return nil
}

func flattenSystemConsoleHostnameDisplayLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemConsoleOutput(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemConsoleBaudrate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemConsoleMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemConsoleLogin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemConsole(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("hostname_display_length", flattenSystemConsoleHostnameDisplayLength(o["hostname-display-length"], d, "hostname_display_length", sv)); err != nil {
		if !fortiAPIPatch(o["hostname-display-length"]) {
			return fmt.Errorf("Error reading hostname_display_length: %v", err)
		}
	}

	if err = d.Set("output", flattenSystemConsoleOutput(o["output"], d, "output", sv)); err != nil {
		if !fortiAPIPatch(o["output"]) {
			return fmt.Errorf("Error reading output: %v", err)
		}
	}

	if err = d.Set("baudrate", flattenSystemConsoleBaudrate(o["baudrate"], d, "baudrate", sv)); err != nil {
		if !fortiAPIPatch(o["baudrate"]) {
			return fmt.Errorf("Error reading baudrate: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemConsoleMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("login", flattenSystemConsoleLogin(o["login"], d, "login", sv)); err != nil {
		if !fortiAPIPatch(o["login"]) {
			return fmt.Errorf("Error reading login: %v", err)
		}
	}

	return nil
}

func flattenSystemConsoleFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemConsoleHostnameDisplayLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleOutput(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleBaudrate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleLogin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemConsole(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("hostname_display_length"); ok {
		if setArgNil {
			obj["hostname-display-length"] = nil
		} else {

			t, err := expandSystemConsoleHostnameDisplayLength(d, v, "hostname_display_length", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostname-display-length"] = t
			}
		}
	}

	if v, ok := d.GetOk("output"); ok {
		if setArgNil {
			obj["output"] = nil
		} else {

			t, err := expandSystemConsoleOutput(d, v, "output", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["output"] = t
			}
		}
	}

	if v, ok := d.GetOk("baudrate"); ok {
		if setArgNil {
			obj["baudrate"] = nil
		} else {

			t, err := expandSystemConsoleBaudrate(d, v, "baudrate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["baudrate"] = t
			}
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		if setArgNil {
			obj["mode"] = nil
		} else {

			t, err := expandSystemConsoleMode(d, v, "mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("login"); ok {
		if setArgNil {
			obj["login"] = nil
		} else {

			t, err := expandSystemConsoleLogin(d, v, "login", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["login"] = t
			}
		}
	}

	return &obj, nil
}
