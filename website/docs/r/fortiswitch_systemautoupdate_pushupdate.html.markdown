---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemautoupdate_pushupdate"
description: |-
  Configure push updates.
---

# fortiswitch_systemautoupdate_pushupdate
Configure push updates.

## Example Usage

```hcl
resource "fortiswitch_systemautoupdate_pushupdate" "saop" {
        address = "3.4.5.6"
        override = "enable"
        port = "5"
        status = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable push updates. Valid values: `enable`, `disable`.
* `override` - Enable the push update override server. Valid values: `enable`, `disable`.
* `port` - Configure the push update override port.
* `address` - Configure the push update override server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SystemAutoupdate PushUpdate can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemautoupdate_pushupdate.labelname SystemAutoupdatePushUpdate

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemautoupdate_pushupdate.labelname SystemAutoupdatePushUpdate
$ unset "FORTISWITCH_IMPORT_TABLE"
```
