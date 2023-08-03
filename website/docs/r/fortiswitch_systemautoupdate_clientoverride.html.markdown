---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemautoupdate_clientoverride"
description: |-
  Configure client override for the FDN.
---

# fortiswitch_systemautoupdate_clientoverride
Configure client override for the FDN.

## Example Usage

```hcl
resource "fortiswitch_systemautoupdate_clientoverride" "fsat" {
        address = "2.3.4.5"
        status = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable client override. Valid values: `enable`, `disable`.
* `address` - Source ip for FGD AV/IPS updates.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SystemAutoupdate Clientoverride can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemautoupdate_clientoverride.labelname SystemAutoupdateClientoverride

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemautoupdate_clientoverride.labelname SystemAutoupdateClientoverride
$ unset "FORTISWITCH_IMPORT_TABLE"
```
