// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Alias command definitions.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAliasCommand() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAliasCommandCreate,
		Read:   resourceSystemAliasCommandRead,
		Update: resourceSystemAliasCommandUpdate,
		Delete: resourceSystemAliasCommandDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"limit_shown_attributes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"table_entry_create": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permission": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"table_ids_allowed": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"script_arguments": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"help": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"range_delay": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"optional": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"table_path": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"allowed_values": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"range": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"command": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
				Computed:     true,
			},
			"path": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"table_listing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"read_only_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceSystemAliasCommandCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAliasCommand(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAliasCommand resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAliasCommand(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemAliasCommand resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAliasCommand")
	}

	return resourceSystemAliasCommandRead(d, m)
}

func resourceSystemAliasCommandUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAliasCommand(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAliasCommand resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAliasCommand(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAliasCommand resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAliasCommand")
	}

	return resourceSystemAliasCommandRead(d, m)
}

func resourceSystemAliasCommandDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAliasCommand(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAliasCommand resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAliasCommandRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAliasCommand(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAliasCommand resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAliasCommand(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAliasCommand resource from API: %v", err)
	}
	return nil
}

func flattenSystemAliasCommandLimitShownAttributes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandTableEntryCreate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandPermission(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandAttribute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandTableIdsAllowed(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_id"
		if _, ok := i["entry-id"]; ok {

			tmp["entry_id"] = flattenSystemAliasCommandTableIdsAllowedEntryId(i["entry-id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "entry_id", d)
	return result
}

func flattenSystemAliasCommandTableIdsAllowedEntryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandScriptArguments(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "help"
		if _, ok := i["help"]; ok {

			tmp["help"] = flattenSystemAliasCommandScriptArgumentsHelp(i["help"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range_delay"
		if _, ok := i["range-delay"]; ok {

			tmp["range_delay"] = flattenSystemAliasCommandScriptArgumentsRangeDelay(i["range-delay"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "optional"
		if _, ok := i["optional"]; ok {

			tmp["optional"] = flattenSystemAliasCommandScriptArgumentsOptional(i["optional"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_path"
		if _, ok := i["table-path"]; ok {

			tmp["table_path"] = flattenSystemAliasCommandScriptArgumentsTablePath(i["table-path"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_values"
		if _, ok := i["allowed-values"]; ok {

			tmp["allowed_values"] = flattenSystemAliasCommandScriptArgumentsAllowedValues(i["allowed-values"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := i["range"]; ok {

			tmp["range"] = flattenSystemAliasCommandScriptArgumentsRange(i["range"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenSystemAliasCommandScriptArgumentsType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSystemAliasCommandScriptArgumentsId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemAliasCommandScriptArgumentsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemAliasCommandScriptArgumentsHelp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandScriptArgumentsRangeDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandScriptArgumentsOptional(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandScriptArgumentsTablePath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandScriptArgumentsAllowedValues(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {

			tmp["value"] = flattenSystemAliasCommandScriptArgumentsAllowedValuesValue(i["value"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "value", d)
	return result
}

func flattenSystemAliasCommandScriptArgumentsAllowedValuesValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandScriptArgumentsRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandScriptArgumentsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandScriptArgumentsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandScriptArgumentsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandCommand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandTableListing(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandReadOnlyAttributes(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_name"
		if _, ok := i["attribute-name"]; ok {

			tmp["attribute_name"] = flattenSystemAliasCommandReadOnlyAttributesAttributeName(i["attribute-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "attribute_name", d)
	return result
}

func flattenSystemAliasCommandReadOnlyAttributesAttributeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAliasCommandDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAliasCommand(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("limit_shown_attributes", flattenSystemAliasCommandLimitShownAttributes(o["limit-shown-attributes"], d, "limit_shown_attributes", sv)); err != nil {
		if !fortiAPIPatch(o["limit-shown-attributes"]) {
			return fmt.Errorf("Error reading limit_shown_attributes: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAliasCommandName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("table_entry_create", flattenSystemAliasCommandTableEntryCreate(o["table-entry-create"], d, "table_entry_create", sv)); err != nil {
		if !fortiAPIPatch(o["table-entry-create"]) {
			return fmt.Errorf("Error reading table_entry_create: %v", err)
		}
	}

	if err = d.Set("permission", flattenSystemAliasCommandPermission(o["permission"], d, "permission", sv)); err != nil {
		if !fortiAPIPatch(o["permission"]) {
			return fmt.Errorf("Error reading permission: %v", err)
		}
	}

	if err = d.Set("attribute", flattenSystemAliasCommandAttribute(o["attribute"], d, "attribute", sv)); err != nil {
		if !fortiAPIPatch(o["attribute"]) {
			return fmt.Errorf("Error reading attribute: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("table_ids_allowed", flattenSystemAliasCommandTableIdsAllowed(o["table-ids-allowed"], d, "table_ids_allowed", sv)); err != nil {
			if !fortiAPIPatch(o["table-ids-allowed"]) {
				return fmt.Errorf("Error reading table_ids_allowed: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("table_ids_allowed"); ok {
			if err = d.Set("table_ids_allowed", flattenSystemAliasCommandTableIdsAllowed(o["table-ids-allowed"], d, "table_ids_allowed", sv)); err != nil {
				if !fortiAPIPatch(o["table-ids-allowed"]) {
					return fmt.Errorf("Error reading table_ids_allowed: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("script_arguments", flattenSystemAliasCommandScriptArguments(o["script-arguments"], d, "script_arguments", sv)); err != nil {
			if !fortiAPIPatch(o["script-arguments"]) {
				return fmt.Errorf("Error reading script_arguments: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("script_arguments"); ok {
			if err = d.Set("script_arguments", flattenSystemAliasCommandScriptArguments(o["script-arguments"], d, "script_arguments", sv)); err != nil {
				if !fortiAPIPatch(o["script-arguments"]) {
					return fmt.Errorf("Error reading script_arguments: %v", err)
				}
			}
		}
	}

	if err = d.Set("command", flattenSystemAliasCommandCommand(o["command"], d, "command", sv)); err != nil {
		if !fortiAPIPatch(o["command"]) {
			return fmt.Errorf("Error reading command: %v", err)
		}
	}

	if err = d.Set("path", flattenSystemAliasCommandPath(o["path"], d, "path", sv)); err != nil {
		if !fortiAPIPatch(o["path"]) {
			return fmt.Errorf("Error reading path: %v", err)
		}
	}

	if err = d.Set("table_listing", flattenSystemAliasCommandTableListing(o["table-listing"], d, "table_listing", sv)); err != nil {
		if !fortiAPIPatch(o["table-listing"]) {
			return fmt.Errorf("Error reading table_listing: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("read_only_attributes", flattenSystemAliasCommandReadOnlyAttributes(o["read-only-attributes"], d, "read_only_attributes", sv)); err != nil {
			if !fortiAPIPatch(o["read-only-attributes"]) {
				return fmt.Errorf("Error reading read_only_attributes: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("read_only_attributes"); ok {
			if err = d.Set("read_only_attributes", flattenSystemAliasCommandReadOnlyAttributes(o["read-only-attributes"], d, "read_only_attributes", sv)); err != nil {
				if !fortiAPIPatch(o["read-only-attributes"]) {
					return fmt.Errorf("Error reading read_only_attributes: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenSystemAliasCommandType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemAliasCommandDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	return nil
}

func flattenSystemAliasCommandFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemAliasCommandLimitShownAttributes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandTableEntryCreate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandPermission(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandAttribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandTableIdsAllowed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["entry-id"], _ = expandSystemAliasCommandTableIdsAllowedEntryId(d, i["entry_id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAliasCommandTableIdsAllowedEntryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandScriptArguments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "help"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["help"], _ = expandSystemAliasCommandScriptArgumentsHelp(d, i["help"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range_delay"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["range-delay"], _ = expandSystemAliasCommandScriptArgumentsRangeDelay(d, i["range_delay"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "optional"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["optional"], _ = expandSystemAliasCommandScriptArgumentsOptional(d, i["optional"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_path"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["table-path"], _ = expandSystemAliasCommandScriptArgumentsTablePath(d, i["table_path"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_values"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["allowed-values"], _ = expandSystemAliasCommandScriptArgumentsAllowedValues(d, i["allowed_values"], pre_append, sv)
		} else {
			tmp["allowed-values"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["range"], _ = expandSystemAliasCommandScriptArgumentsRange(d, i["range"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandSystemAliasCommandScriptArgumentsType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSystemAliasCommandScriptArgumentsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemAliasCommandScriptArgumentsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAliasCommandScriptArgumentsHelp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandScriptArgumentsRangeDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandScriptArgumentsOptional(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandScriptArgumentsTablePath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandScriptArgumentsAllowedValues(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["value"], _ = expandSystemAliasCommandScriptArgumentsAllowedValuesValue(d, i["value"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAliasCommandScriptArgumentsAllowedValuesValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandScriptArgumentsRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandScriptArgumentsType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandScriptArgumentsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandScriptArgumentsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandCommand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandTableListing(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandReadOnlyAttributes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["attribute-name"], _ = expandSystemAliasCommandReadOnlyAttributesAttributeName(d, i["attribute_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAliasCommandReadOnlyAttributesAttributeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommandDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAliasCommand(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("limit_shown_attributes"); ok {

		t, err := expandSystemAliasCommandLimitShownAttributes(d, v, "limit_shown_attributes", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["limit-shown-attributes"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemAliasCommandName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("table_entry_create"); ok {

		t, err := expandSystemAliasCommandTableEntryCreate(d, v, "table_entry_create", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-entry-create"] = t
		}
	}

	if v, ok := d.GetOk("permission"); ok {

		t, err := expandSystemAliasCommandPermission(d, v, "permission", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permission"] = t
		}
	}

	if v, ok := d.GetOk("attribute"); ok {

		t, err := expandSystemAliasCommandAttribute(d, v, "attribute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute"] = t
		}
	}

	if v, ok := d.GetOk("table_ids_allowed"); ok || d.HasChange("table_ids_allowed") {

		t, err := expandSystemAliasCommandTableIdsAllowed(d, v, "table_ids_allowed", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-ids-allowed"] = t
		}
	}

	if v, ok := d.GetOk("script_arguments"); ok || d.HasChange("script_arguments") {

		t, err := expandSystemAliasCommandScriptArguments(d, v, "script_arguments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script-arguments"] = t
		}
	}

	if v, ok := d.GetOk("command"); ok {

		t, err := expandSystemAliasCommandCommand(d, v, "command", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command"] = t
		}
	}

	if v, ok := d.GetOk("path"); ok {

		t, err := expandSystemAliasCommandPath(d, v, "path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["path"] = t
		}
	}

	if v, ok := d.GetOk("table_listing"); ok {

		t, err := expandSystemAliasCommandTableListing(d, v, "table_listing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-listing"] = t
		}
	}

	if v, ok := d.GetOk("read_only_attributes"); ok || d.HasChange("read_only_attributes") {

		t, err := expandSystemAliasCommandReadOnlyAttributes(d, v, "read_only_attributes", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["read-only-attributes"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandSystemAliasCommandType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSystemAliasCommandDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	return &obj, nil
}
