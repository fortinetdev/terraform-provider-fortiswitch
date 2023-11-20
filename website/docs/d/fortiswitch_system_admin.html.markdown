---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_admin"
description: |-
  Get information on an fortiswitch system admin.
---

# Data Source: fortiswitch_system_admin
Use this data source to get information on an fortiswitch system admin

## Example Usage

```hcl
data "fortiswitch_system_admin" sample1 {
  name = "api_admin"
}

output output1 {
  value = data.fortiswitch_system_admin.sample1
}
```

## Argument Reference

* `name` - (Required) Specify the name of the desired system admin.

## Attribute Reference

The following attributes are exported:

* `ip6_trusthost2` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `ip6_trusthost3` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `is_admin` - User has administrative privileges.
* `ip6_trusthost1` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `ip6_trusthost6` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `ip6_trusthost7` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `ip6_trusthost4` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `ip6_trusthost5` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `accprofile_override` - Enable/disable remote authentication server to override access profile.
* `ip6_trusthost8` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `ip6_trusthost9` - Trusted host one IP address (default = 0.0.0.0 0.0.0.0 to trust all IPs).
* `accprofile` - Administrative user access profile.
* `force_password_change` - Enable/disable forcing of password change on next login.
* `allow_remove_admin_session` - Enable/disable privileged administrative users to remove administrative sessions.
* `wildcard_fallback` - Enable/disable attempting authentication against wildcard accounts if authenticating this account fails.
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
* `remote_auth` - Enable/disable remote authentication.
* `name` - Adminstrative user name.
* `password_expire` - Password expire time.
* `remote_group` - Remote authentication group name.
* `wildcard` - Enable/disable wildcard RADIUS authentication.
* `ssh_public_key1` - SSH public key1.
* `ssh_public_key3` - SSH public key3.
* `ssh_public_key2` - SSH public key2.
* `peer_auth` - Enable/disable peer authentication.

