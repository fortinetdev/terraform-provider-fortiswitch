// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Groups of alias commands.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAliasGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAliasGroupCreate,
		Read:   resourceSystemAliasGroupRead,
		Update: resourceSystemAliasGroupUpdate,
		Delete: resourceSystemAliasGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"commands": &schema.Schema{
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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemAliasGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAliasGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAliasGroup resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAliasGroup(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemAliasGroup resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAliasGroup")
	}

	return resourceSystemAliasGroupRead(d, m)
}

func resourceSystemAliasGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAliasGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAliasGroup resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAliasGroup(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAliasGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAliasGroup")
	}

	return resourceSystemAliasGroupRead(d, m)
}

func resourceSystemAliasGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAliasGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAliasGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAliasGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAliasGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAliasGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAliasGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAliasGroup resource from API: %v", err)
	}
	return nil
}

func flattenSystemAliasGroupCommands(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["command_name"] = flattenSystemAliasGroupCommandsCommandName(i["command-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "command_name", d)
	return result
}

func flattenSystemAliasGroupCommandsCommandName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasGroupDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAliasGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if isImportTable() {
		if err = d.Set("commands", flattenSystemAliasGroupCommands(o["commands"], d, "commands", sv)); err != nil {
			if !fortiAPIPatch(o["commands"]) {
				return fmt.Errorf("Error reading commands: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("commands"); ok {
			if err = d.Set("commands", flattenSystemAliasGroupCommands(o["commands"], d, "commands", sv)); err != nil {
				if !fortiAPIPatch(o["commands"]) {
					return fmt.Errorf("Error reading commands: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenSystemAliasGroupName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemAliasGroupDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	return nil
}

func flattenSystemAliasGroupFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemAliasGroupCommands(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["command-name"], _ = expandSystemAliasGroupCommandsCommandName(d, i["command_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAliasGroupCommandsCommandName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasGroupDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAliasGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("commands"); ok || d.HasChange("commands") {

		t, err := expandSystemAliasGroupCommands(d, v, "commands", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["commands"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemAliasGroupName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSystemAliasGroupDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	return &obj, nil
}
