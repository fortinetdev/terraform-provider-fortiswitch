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

func dataSourceSystemConsole() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemConsoleRead,
		Schema: map[string]*schema.Schema{

			"hostname_display_length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"output": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"baudrate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"login": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemConsoleRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemConsole"

	o, err := c.ReadSystemConsole(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemConsole: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemConsole(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemConsole from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemConsoleHostnameDisplayLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemConsoleOutput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemConsoleBaudrate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemConsoleMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemConsoleLogin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemConsole(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("hostname_display_length", dataSourceFlattenSystemConsoleHostnameDisplayLength(o["hostname-display-length"], d, "hostname_display_length")); err != nil {
		if !fortiAPIPatch(o["hostname-display-length"]) {
			return fmt.Errorf("Error reading hostname_display_length: %v", err)
		}
	}

	if err = d.Set("output", dataSourceFlattenSystemConsoleOutput(o["output"], d, "output")); err != nil {
		if !fortiAPIPatch(o["output"]) {
			return fmt.Errorf("Error reading output: %v", err)
		}
	}

	if err = d.Set("baudrate", dataSourceFlattenSystemConsoleBaudrate(o["baudrate"], d, "baudrate")); err != nil {
		if !fortiAPIPatch(o["baudrate"]) {
			return fmt.Errorf("Error reading baudrate: %v", err)
		}
	}

	if err = d.Set("mode", dataSourceFlattenSystemConsoleMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("login", dataSourceFlattenSystemConsoleLogin(o["login"], d, "login")); err != nil {
		if !fortiAPIPatch(o["login"]) {
			return fmt.Errorf("Error reading login: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemConsoleFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
