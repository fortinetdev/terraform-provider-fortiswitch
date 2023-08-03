---
subcategory: "FortiSwitch User"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_user_ldap"
description: |-
  LDAP server entry configuration.
---

# fortiswitch_user_ldap
LDAP server entry configuration.

## Example Usage

```hcl
resource "fortiswitch_user_ldap" "uldap" {
        name = "test89"
        password = "123"
        password_expiry_warning = "enable"
        password_renewal = "enable"
        port = "13"
        secure = "disable"
        server = "192.168.100.40"
        type = "simple"
}
```

## Argument Reference

The following arguments are supported:

* `dn` - Distinguished name.
* `username` - Username (full DN) for initial binding.
* `name` - LDAP server entry name.
* `password_expiry_warning` - Enable/disable password expiry warnings. Valid values: `enable`, `disable`.
* `group_member_check` - Group member checking options. Valid values: `user-attr`, `group-object`.
* `server` - LDAP server domain name or IP address.
* `ca_cert` - CA certificate name.
* `group_object_filter` - Group searching filter.
* `cnid` - Common Name identifier (default = "cn").
* `member_attr` - Name of attribute from which to get group membership.
* `password_renewal` - Enable/disable online password renewal. Valid values: `enable`, `disable`.
* `password` - Password for initial binding.
* `type` - LDAP binding type. Valid values: `simple`, `anonymous`, `regular`.
* `port` - LDAP server port number (1 - 65535, default = 389).
* `secure` - SSL connection. Valid values: `disable`, `starttls`, `ldaps`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Ldap can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_user_ldap.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_user_ldap.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
