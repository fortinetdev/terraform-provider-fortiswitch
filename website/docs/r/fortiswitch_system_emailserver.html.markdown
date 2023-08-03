---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_emailserver"
description: |-
  Email server configuration.
---

# fortiswitch_system_emailserver
Email server configuration. Applies to FortiSwitch Version `>= 7.2.0`.

## Argument Reference

The following arguments are supported:

* `username` - Set SMTP server user name for authentication.
* `authenticate` - Enable/disable authentication. Valid values: `enable`, `disable`.
* `security` - Set connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.
* `source_ip` - Set SMTP server source ip.
* `validate_server` - Enable/disable validation of server certificate. Valid values: `enable`, `disable`.
* `server` - Set SMTP server IP address or hostname.
* `source_ip6` - Set SMTP server IPv6 source IP.
* `reply_to` - Set Reply-To email address.
* `password` - Set SMTP server user password for authentication.
* `port` - Set SMTP server port.
* `ssl_min_proto_version` - Set minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `TLSv1-3`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System EmailServer can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_emailserver.labelname SystemEmailServer

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_emailserver.labelname SystemEmailServer
$ unset "FORTISWITCH_IMPORT_TABLE"
```
