// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure system administrative access group.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAccprofile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAccprofileCreate,
		Read:   resourceSystemAccprofileRead,
		Update: resourceSystemAccprofileUpdate,
		Delete: resourceSystemAccprofileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"exec_alias_grp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"swmonguardgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"sysgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pktmongrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loggrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mntgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admingrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routegrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"swcoregrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utilgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alias_commands": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemAccprofileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAccprofile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAccprofile resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAccprofile(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemAccprofile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAccprofile")
	}

	return resourceSystemAccprofileRead(d, m)
}

func resourceSystemAccprofileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAccprofile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAccprofile resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAccprofile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAccprofile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAccprofile")
	}

	return resourceSystemAccprofileRead(d, m)
}

func resourceSystemAccprofileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAccprofile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAccprofile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAccprofileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAccprofile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAccprofile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAccprofile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAccprofile resource from API: %v", err)
	}
	return nil
}

func flattenSystemAccprofileExecAliasGrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileSwmonguardgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileSysgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofilePktmongrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileLoggrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileMntgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileNetgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileAdmingrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileRoutegrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileSwcoregrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtilgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileAliasCommands(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_name"
		if _, ok := i["command-name"]; ok {

			tmp["command_name"] = flattenSystemAccprofileAliasCommandsCommandName(i["command-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "command_name", d)
	return result
}

func flattenSystemAccprofileAliasCommandsCommandName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAccprofile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("exec_alias_grp", flattenSystemAccprofileExecAliasGrp(o["exec-alias-grp"], d, "exec_alias_grp", sv)); err != nil {
		if !fortiAPIPatch(o["exec-alias-grp"]) {
			return fmt.Errorf("Error reading exec_alias_grp: %v", err)
		}
	}

	if err = d.Set("swmonguardgrp", flattenSystemAccprofileSwmonguardgrp(o["swmonguardgrp"], d, "swmonguardgrp", sv)); err != nil {
		if !fortiAPIPatch(o["swmonguardgrp"]) {
			return fmt.Errorf("Error reading swmonguardgrp: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAccprofileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("sysgrp", flattenSystemAccprofileSysgrp(o["sysgrp"], d, "sysgrp", sv)); err != nil {
		if !fortiAPIPatch(o["sysgrp"]) {
			return fmt.Errorf("Error reading sysgrp: %v", err)
		}
	}

	if err = d.Set("pktmongrp", flattenSystemAccprofilePktmongrp(o["pktmongrp"], d, "pktmongrp", sv)); err != nil {
		if !fortiAPIPatch(o["pktmongrp"]) {
			return fmt.Errorf("Error reading pktmongrp: %v", err)
		}
	}

	if err = d.Set("loggrp", flattenSystemAccprofileLoggrp(o["loggrp"], d, "loggrp", sv)); err != nil {
		if !fortiAPIPatch(o["loggrp"]) {
			return fmt.Errorf("Error reading loggrp: %v", err)
		}
	}

	if err = d.Set("mntgrp", flattenSystemAccprofileMntgrp(o["mntgrp"], d, "mntgrp", sv)); err != nil {
		if !fortiAPIPatch(o["mntgrp"]) {
			return fmt.Errorf("Error reading mntgrp: %v", err)
		}
	}

	if err = d.Set("netgrp", flattenSystemAccprofileNetgrp(o["netgrp"], d, "netgrp", sv)); err != nil {
		if !fortiAPIPatch(o["netgrp"]) {
			return fmt.Errorf("Error reading netgrp: %v", err)
		}
	}

	if err = d.Set("admingrp", flattenSystemAccprofileAdmingrp(o["admingrp"], d, "admingrp", sv)); err != nil {
		if !fortiAPIPatch(o["admingrp"]) {
			return fmt.Errorf("Error reading admingrp: %v", err)
		}
	}

	if err = d.Set("routegrp", flattenSystemAccprofileRoutegrp(o["routegrp"], d, "routegrp", sv)); err != nil {
		if !fortiAPIPatch(o["routegrp"]) {
			return fmt.Errorf("Error reading routegrp: %v", err)
		}
	}

	if err = d.Set("swcoregrp", flattenSystemAccprofileSwcoregrp(o["swcoregrp"], d, "swcoregrp", sv)); err != nil {
		if !fortiAPIPatch(o["swcoregrp"]) {
			return fmt.Errorf("Error reading swcoregrp: %v", err)
		}
	}

	if err = d.Set("utilgrp", flattenSystemAccprofileUtilgrp(o["utilgrp"], d, "utilgrp", sv)); err != nil {
		if !fortiAPIPatch(o["utilgrp"]) {
			return fmt.Errorf("Error reading utilgrp: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("alias_commands", flattenSystemAccprofileAliasCommands(o["alias-commands"], d, "alias_commands", sv)); err != nil {
			if !fortiAPIPatch(o["alias-commands"]) {
				return fmt.Errorf("Error reading alias_commands: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("alias_commands"); ok {
			if err = d.Set("alias_commands", flattenSystemAccprofileAliasCommands(o["alias-commands"], d, "alias_commands", sv)); err != nil {
				if !fortiAPIPatch(o["alias-commands"]) {
					return fmt.Errorf("Error reading alias_commands: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemAccprofileFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemAccprofileExecAliasGrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSwmonguardgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSysgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofilePktmongrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileLoggrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileMntgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileNetgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileAdmingrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileRoutegrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSwcoregrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtilgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileAliasCommands(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["command-name"], _ = expandSystemAccprofileAliasCommandsCommandName(d, i["command_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAccprofileAliasCommandsCommandName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAccprofile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("exec_alias_grp"); ok {

		t, err := expandSystemAccprofileExecAliasGrp(d, v, "exec_alias_grp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exec-alias-grp"] = t
		}
	}

	if v, ok := d.GetOk("swmonguardgrp"); ok {

		t, err := expandSystemAccprofileSwmonguardgrp(d, v, "swmonguardgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["swmonguardgrp"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemAccprofileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("sysgrp"); ok {

		t, err := expandSystemAccprofileSysgrp(d, v, "sysgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sysgrp"] = t
		}
	}

	if v, ok := d.GetOk("pktmongrp"); ok {

		t, err := expandSystemAccprofilePktmongrp(d, v, "pktmongrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pktmongrp"] = t
		}
	}

	if v, ok := d.GetOk("loggrp"); ok {

		t, err := expandSystemAccprofileLoggrp(d, v, "loggrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loggrp"] = t
		}
	}

	if v, ok := d.GetOk("mntgrp"); ok {

		t, err := expandSystemAccprofileMntgrp(d, v, "mntgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mntgrp"] = t
		}
	}

	if v, ok := d.GetOk("netgrp"); ok {

		t, err := expandSystemAccprofileNetgrp(d, v, "netgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netgrp"] = t
		}
	}

	if v, ok := d.GetOk("admingrp"); ok {

		t, err := expandSystemAccprofileAdmingrp(d, v, "admingrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admingrp"] = t
		}
	}

	if v, ok := d.GetOk("routegrp"); ok {

		t, err := expandSystemAccprofileRoutegrp(d, v, "routegrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["routegrp"] = t
		}
	}

	if v, ok := d.GetOk("swcoregrp"); ok {

		t, err := expandSystemAccprofileSwcoregrp(d, v, "swcoregrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["swcoregrp"] = t
		}
	}

	if v, ok := d.GetOk("utilgrp"); ok {

		t, err := expandSystemAccprofileUtilgrp(d, v, "utilgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utilgrp"] = t
		}
	}

	if v, ok := d.GetOk("alias_commands"); ok || d.HasChange("alias_commands") {

		t, err := expandSystemAccprofileAliasCommands(d, v, "alias_commands", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias-commands"] = t
		}
	}

	return &obj, nil
}
