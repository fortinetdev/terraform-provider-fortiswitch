---
subcategory: "FortiSwitch User"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_user_setting"
description: |-
  User authentication setting.
---

# fortiswitch_user_setting
User authentication setting.

## Example Usage

```hcl
resource "fortiswitch_user_setting" "usetting" {
  auth_http_basic = "enable"
  auth_invalid_max = "6"
  auth_ports {
      id =  "9"
      port = "10"
      type = "http"
  }
  auth_secure_http = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `auth_multi_group` - Enable/disable retrival of all groups a user belongs to. Valid values: `enable`, `disable`.
* `auth_timeout` - Firewall user authentication timeout (1 - 480).
* `auth_ports` - Authentication port table. The structure of `auth_ports` block is documented below.
* `auth_cert` - HTTPS server cert for policy authentication.
* `auth_http_basic` - Enable/disable use of HTTP BASIC for HTTP authentication. Valid values: `enable`, `disable`.
* `auth_invalid_max` - Maximum number of invalid authentication attempts allowed before blackout (1 - 100).
* `auth_blackout_time` - Authentication blackout time (0 - 3600 sec).
* `auth_timeout_type` - Authenticated policy expiration behavior. Valid values: `idle-timeout`, `hard-timeout`, `new-session`.
* `auth_secure_http` - Enable/disable use of HTTPS for HTTP authentication. Valid values: `enable`, `disable`.
* `auth_type` - Allowed firewall policy authentication methods. Valid values: `http`, `https`, `ftp`, `telnet`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `auth_ports` block supports:

* `type` - Service type. Valid values: `http`, `https`, `ftp`, `telnet`.
* `id` - ID.
* `port` - Non-standard port for firewall user authentication (1 - 65535).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

User Setting can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_user_setting.labelname UserSetting

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_user_setting.labelname UserSetting
$ unset "FORTISWITCH_IMPORT_TABLE"
```
