---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemautoupdate_override"
description: |-
  Configure override FDS server.
---

# fortiswitch_systemautoupdate_override
Configure override FDS server.

## Example Usage

```hcl
resource "fortiswitch_systemautoupdate_override" "saod" {
        address = "5.6.7.8"
        fail_over = "enable"
        status = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable FDS server override. Valid values: `enable`, `disable`.
* `fail_over` - Enable/disable FDS server fail over. Valid values: `enable`, `disable`.
* `address` - Set the FDS override server address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SystemAutoupdate Override can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemautoupdate_override.labelname SystemAutoupdateOverride

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemautoupdate_override.labelname SystemAutoupdateOverride
$ unset "FORTISWITCH_IMPORT_TABLE"
```
