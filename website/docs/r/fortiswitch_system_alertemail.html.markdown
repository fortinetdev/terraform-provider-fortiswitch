---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_alertemail"
description: |-
  Alert e-mail mail server configuration.
---

# fortiswitch_system_alertemail
Alert e-mail mail server configuration. Applies to FortiSwitch Version `<= 7.0.6`.

## Argument Reference

The following arguments are supported:

* `username` - Set SMTP server user name for authentication.
* `authenticate` - Enable/disable authentication. Valid values: `enable`, `disable`.
* `source_ip` - Set SMTP server source ip.
* `server` - Set SMTP server IP address or hostname.
* `password` - Set SMTP server user password for authentication.
* `port` - Set SMTP server port.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Alertemail can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_alertemail.labelname SystemAlertemail

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_alertemail.labelname SystemAlertemail
$ unset "FORTISWITCH_IMPORT_TABLE"
```
