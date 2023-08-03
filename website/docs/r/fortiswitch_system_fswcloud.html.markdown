---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_fswcloud"
description: |-
  FortiSwitch cloud manager configuration.
---

# fortiswitch_system_fswcloud
FortiSwitch cloud manager configuration. Applies to FortiSwitch Version `<= 7.0.2`.

## Example Usage

```hcl
resource "fortiswitch_system_fswcloud" "name" {
        interval = "3"
        name = "default_name_4"
        port = "5"
        status = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable fsw-cloud service. Valid values: `enable`, `disable`.
* `interval` - Service name resolution time interval (3-300sec, default=45).
* `port` - Port Number.
* `name` - Fully qualified domain name or IP address of fsw-cloud service.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FswCloud can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_fswcloud.labelname SystemFswCloud

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_fswcloud.labelname SystemFswCloud
$ unset "FORTISWITCH_IMPORT_TABLE"
```
