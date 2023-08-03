---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemalias_command"
description: |-
  Alias command definitions.
---

# fortiswitch_systemalias_command
Alias command definitions.

## Argument Reference

The following arguments are supported:

* `limit_shown_attributes` - Enable/disable limiting of config displayed in show and get. Valid values: `disable`, `enable`.
* `name` - Alias name.
* `table_entry_create` - Allow/prevent this script from creating new entries in config tables. Valid values: `allow`, `deny`.
* `permission` - Allow read and write operations, or only read operations on this path. Valid values: `read`, `read-write`.
* `attribute` - Attribute which can be retreived or modified.
* `table_ids_allowed` - Table entries this command is limited to. The structure of `table_ids_allowed` block is documented below.
* `script_arguments` - Optional meta-data and control values for script arguments. The structure of `script_arguments` block is documented below.
* `command` - Script command template (use $1 for user argument).
* `path` - Path to configuration object or table.
* `table_listing` - Allow/prevent listing of all entries in the config table. Valid values: `allow`, `deny`.
* `read_only_attributes` - Additional attributes allowed in get/show output when limit-shown-attributes is enabled. The structure of `read_only_attributes` block is documented below.
* `type` - Command type to alias. Valid values: `configuration`, `script`.
* `description` - Command description and/or help message.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `table_ids_allowed` block supports:

* `entry_id` - Entry ID.

The `script_arguments` block supports:

* `help` - A help message for the argument.
* `range_delay` - When running against a range of values, delay this many seconds between values when executing.
* `optional` - Enable/disable the option to omit this argument. Valid values: `disable`, `enable`.
* `table_path` - Path to configuration object or table.
* `allowed_values` - Values to limit this argument to. The structure of `allowed_values` block is documented below.
* `range` - Enable/disable the option to pass a range, or list, of values for this argument. Valid values: `disable`, `enable`.
* `type` - Argument data type. Valid values: `string`, `integer`, `table-id`.
* `id` - Argument ID.
* `name` - Display name for the argument.

The `allowed_values` block supports:

* `value` - Allowed value.

The `read_only_attributes` block supports:

* `attribute_name` - Attribute name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemAlias Command can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemalias_command.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemalias_command.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
