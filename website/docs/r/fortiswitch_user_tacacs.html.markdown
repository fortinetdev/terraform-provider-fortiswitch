---
subcategory: "FortiSwitch User"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_user_tacacs"
description: |-
  TACACS+ server entry configuration.
---

# fortiswitch_user_tacacs
TACACS+ server entry configuration.

## Example Usage

```hcl
resource "fortiswitch_user_tacacs" "trname" {
  authen_type = "mschap"
  authorization = "enable"
  key = "342 wer "
  name = "tacacs"
  port = "8"
  server = "192.168.100.40"
  #source_ip = "20.20.20.10"
}
```

## Argument Reference

The following arguments are supported:

* `authen_type` - Authentication type. Valid values: `mschap`, `chap`, `pap`, `ascii`, `auto`.
* `name` - TACACS+ server entry name.
* `source_ip` - Source IP for communications to TACACS+ server.
* `server` - Server domain name or IP address.
* `key` - Key to access the server.
* `port` - TACACS+ server port number (1 - 65535).
* `authorization` - Enable/disable TACACS+ authorization. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Tacacs can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_user_tacacs.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_user_tacacs.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
