---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_accprofile"
description: |-
  Get information on an fortiswitch system accprofile.
---

# Data Source: fortiswitch_system_accprofile
Use this data source to get information on an fortiswitch system accprofile

## Argument Reference

* `name` - (Required) Specify the name of the desired system accprofile.

## Attribute Reference

The following attributes are exported:

* `exec_alias_grp` - Permission to execute alias commands.
* `swmonguardgrp` - Switch monitor and guard feature permission.
* `name` - Profile name.
* `sysgrp` - System access permission.
* `pktmongrp` - Packet and flow capture functionality.
* `loggrp` - Logging access permission.
* `mntgrp` - Critical system maintenance access permission.
* `netgrp` - Network access permission.
* `admingrp` - Administrative access permission.
* `routegrp` - Routing access permission.
* `swcoregrp` - Switch core access permission.
* `utilgrp` - Utilities access permission.
* `alias_commands` - Alias commands (or groups of commands) that can be used by this profile. The structure of `alias_commands` block is documented below.

The `alias_commands` block contains:

* `command_name` - Alias command or group name.

