---
subcategory: "FortiSwitch User"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_user_local"
description: |-
  Local user configuration.
---

# fortiswitch_user_local
Local user configuration.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable user status. Valid values: `enable`, `disable`.
* `name` - User name.
* `passwd` - User password.
* `ldap_server` - LDAP server name.
* `radius_server` - RADIUS server name.
* `tacacs_server` - TACACS+ server name.
* `type` - Authentication type. Valid values: `password`, `radius`, `tacacs+`, `ldap`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Local can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_user_local.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_user_local.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
