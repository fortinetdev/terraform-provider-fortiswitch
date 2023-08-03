---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemautoupdate_tunneling"
description: |-
  Configure web proxy tunneling for the FDN.
---

# fortiswitch_systemautoupdate_tunneling
Configure web proxy tunneling for the FDN.

## Example Usage

```hcl
resource "fortiswitch_systemautoupdate_tunneling" "saot" {
        address = "5.6.7.8"
        password = "234"
        port = "5"
        status = "enable"
        username = "admin"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable web proxy tunneling. Valid values: `enable`, `disable`.
* `username` - Configure the web proxy username.
* `password` - Configure the web proxy password.
* `port` - Configure the web proxy port.
* `address` - Configure the web proxy ip address or fqdn.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SystemAutoupdate Tunneling can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemautoupdate_tunneling.labelname SystemAutoupdateTunneling

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemautoupdate_tunneling.labelname SystemAutoupdateTunneling
$ unset "FORTISWITCH_IMPORT_TABLE"
```
