---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_admin"
description: |-
  Administrative user configuration.
---

# fortiswitch_system_admin
Administrative user configuration.

## Example Usage

```hcl
resource "fortiswitch_system_admin" "admin" {
       accprofile = "prof_admin"
       accprofile_override = "disable"
       comments = "test_new"
       name = "api_admin_4"
       password = "1234mn90812"
}
```

## Argument Reference

The following arguments are supported:

* `ip6_trusthost2` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `ip6_trusthost3` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `is_admin` - User has administrative privileges.
* `ip6_trusthost1` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `ip6_trusthost6` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `ip6_trusthost7` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `ip6_trusthost4` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `ip6_trusthost5` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `accprofile_override` - Enable/disable remote authentication server to override access profile. Valid values: `enable`, `disable`.
* `ip6_trusthost8` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `ip6_trusthost9` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `accprofile` - Administrative user access profile.
* `force_password_change` - Enable/disable forcing of password change on next login. Valid values: `enable`, `disable`.
* `allow_remove_admin_session` - Enable/disable privileged administrative users to remove administrative sessions. Valid values: `enable`, `disable`.
* `wildcard_fallback` - Enable/disable attempting authentication against wildcard accounts if authenticating this account fails. Valid values: `enable`, `disable`.
* `schedule` - Schedule name.
* `peer_group` - Peer group name.
* `comments` - Comment.
* `trusthost10` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `hidden` - Administrative user hidden attribute.
* `trusthost8` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `trusthost9` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `trusthost6` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `trusthost7` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `trusthost4` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `trusthost5` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `trusthost2` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `trusthost3` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `trusthost1` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `ip6_trusthost10` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `password` - Remote authentication password.
* `vdom` - Virtual domain name.
* `remote_auth` - Enable/disable remote authentication. Valid values: `enable`, `disable`.
* `name` - Adminstrative user name.
* `password_expire` - Password expire time.
* `remote_group` - Remote authentication group name.
* `wildcard` - Enable/disable wildcard RADIUS authentication. Valid values: `enable`, `disable`.
* `ssh_public_key1` - SSH public key1.
* `ssh_public_key3` - SSH public key3.
* `ssh_public_key2` - SSH public key2.
* `peer_auth` - Enable/disable peer authentication. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Admin can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_admin.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_admin.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
