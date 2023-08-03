---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_dnsserver"
description: |-
  Dns-server.
---

# fortiswitch_system_dnsserver
Dns-server.

## Example Usage

```hcl
resource "fortiswitch_system_dnsserver" "name" {
    name = "internal"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Dns server name.
* `mode` - Dns server mode. Valid values: `recursive`, `non-recursive`, `forward-only`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System DnsServer can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_dnsserver.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_dnsserver.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
