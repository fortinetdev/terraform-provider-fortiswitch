// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Dashboard CLI console configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGuiConsole() *schema.Resource {
	return &schema.Resource{
		Create: resourceGuiConsoleUpdate,
		Read:   resourceGuiConsoleRead,
		Update: resourceGuiConsoleUpdate,
		Delete: resourceGuiConsoleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"preferences": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceGuiConsoleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectGuiConsole(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating GuiConsole resource while getting object: %v", err)
	}

	o, err := c.UpdateGuiConsole(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating GuiConsole resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("GuiConsole")
	}

	return resourceGuiConsoleRead(d, m)
}

func resourceGuiConsoleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectGuiConsole(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating GuiConsole resource while getting object: %v", err)
	}

	_, err = c.UpdateGuiConsole(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing GuiConsole resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceGuiConsoleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadGuiConsole(mkey)
	if err != nil {
		return fmt.Errorf("Error reading GuiConsole resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectGuiConsole(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading GuiConsole resource from API: %v", err)
	}
	return nil
}

func flattenGuiConsolePreferences(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectGuiConsole(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("preferences", flattenGuiConsolePreferences(o["preferences"], d, "preferences", sv)); err != nil {
		if !fortiAPIPatch(o["preferences"]) {
			return fmt.Errorf("Error reading preferences: %v", err)
		}
	}

	return nil
}

func flattenGuiConsoleFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandGuiConsolePreferences(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectGuiConsole(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("preferences"); ok {
		if setArgNil {
			obj["preferences"] = nil
		} else {

			t, err := expandGuiConsolePreferences(d, v, "preferences", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["preferences"] = t
			}
		}
	}

	return &obj, nil
}
