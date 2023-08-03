---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemsnmp_user"
description: |-
  SNMP user configuration.
---

# fortiswitch_systemsnmp_user
SNMP user configuration.

## Example Usage

```hcl
resource "fortiswitch_systemsnmp_user" "name" {
        auth_proto = "md5"
        auth_pwd = "123"
        events = "cpu_high"
        name = "default_name_6"
        priv_proto = "aes128"
        priv_pwd = "123"
        queries = "enable"
        query_port = "11"
        security_level = "no_auth_no_priv"
}
```

## Argument Reference

The following arguments are supported:

* `notify_hosts` - Send notifications (traps) to these hosts.
* `priv_proto` - Privacy (encryption) protocol. Valid values: `aes128`, `aes192`, `aes192c`, `aes256`, `aes256c`, `des`.
* `name` - SNMP user name.
* `query_port` - SNMPv3 query port.
* `auth_pwd` - Password for authentication protocol.
* `priv_pwd` - Password for privacy (encryption) protocol.
* `security_level` - Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
* `auth_proto` - Authentication protocol. Valid values: `md5`, `sha1`, `sha224`, `sha256`, `sha384`, `sha512`.
* `queries` - Enable/disable queries for this user. Valid values: `enable`, `disable`.
* `events` - SNMP notifications (traps) to send.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemSnmp User can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemsnmp_user.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemsnmp_user.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
