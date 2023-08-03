---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_vdomdns"
description: |-
  Vdom dns configuration.
---

# fortiswitch_system_vdomdns
Vdom dns configuration.

## Example Usage

```hcl
resource "fortiswitch_system_vdomdns" "name" {
        vdom_dns = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `vdom_dns` - Enable/Disable dns per vdom. Valid values: `enable`, `disable`.
* `source_ip` - Source IP for communications to DNS server.
* `primary` - Vdom primary dns ip.
* `ip6_secondary` - Vdom IPv6 seondary dns ip.
* `ip6_primary` - Vdom IPv6 primary dns ip.
* `secondary` - Vdom secondary dns ip.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System VdomDns can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_vdomdns.labelname SystemVdomDns

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_vdomdns.labelname SystemVdomDns
$ unset "FORTISWITCH_IMPORT_TABLE"
```
