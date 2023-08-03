---
subcategory: "FortiSwitch User"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_user_peer"
description: |-
  config peer user
---

# fortiswitch_user_peer
config peer user

## Example Usage

```hcl
resource "fortiswitch_user_peer" "upeer" {
        name = "peer"
        passwd = "123"
}
```

## Argument Reference

The following arguments are supported:

* `cn` - peer certificate common name
* `cn_type` - peer certificate common name type Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.
* `two_factor` - Enable 2-factor authentication (certificate + password) Valid values: `enable`, `disable`.
* `ldap_username` - username for LDAP server bind
* `ca` - peer certificate CA (CA name in local)
* `mandatory_ca_verify` - mandatory CA verify Valid values: `enable`, `disable`.
* `ldap_server` - LDAP server for access rights check
* `name` - peer name
* `passwd` - Config user password
* `ldap_password` - password for LDAP server bind
* `ldap_mode` - peer ldap mode Valid values: `password`, `principal-name`.
* `subject` - peer certificate name constraints


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Peer can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_user_peer.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_user_peer.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
