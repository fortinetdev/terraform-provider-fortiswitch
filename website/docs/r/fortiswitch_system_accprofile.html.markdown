---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_accprofile"
description: |-
  Configure system administrative access group.
---

# fortiswitch_system_accprofile
Configure system administrative access group.

## Argument Reference

The following arguments are supported:

* `exec_alias_grp` - Permission to execute alias commands. Valid values: `none`, `read`, `read-write`.
* `swmonguardgrp` - Switch monitor and guard feature permission. Valid values: `none`, `read`, `read-write`.
* `name` - Profile name.
* `sysgrp` - System access permission. Valid values: `none`, `read`, `read-write`.
* `pktmongrp` - Packet and flow capture functionality. Valid values: `none`, `read`, `read-write`.
* `loggrp` - Logging access permission. Valid values: `none`, `read`, `read-write`, `w`, `r`, `rw`.
* `mntgrp` - Critical system maintenance access permission. Valid values: `none`, `read`, `read-write`.
* `netgrp` - Network access permission. Valid values: `none`, `read`, `read-write`.
* `admingrp` - Administrative access permission. Valid values: `none`, `read`, `read-write`.
* `routegrp` - Routing access permission. Valid values: `none`, `read`, `read-write`.
* `swcoregrp` - Switch core access permission. Valid values: `none`, `read`, `read-write`.
* `utilgrp` - Utilities access permission. Valid values: `none`, `read`, `read-write`.
* `alias_commands` - Alias commands (or groups of commands) that can be used by this profile. The structure of `alias_commands` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `alias_commands` block supports:

* `command_name` - Alias command or group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Accprofile can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_accprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_accprofile.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
