---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_flancloud"
description: |-
  FortiLAN cloud manager configuration.
---

# fortiswitch_system_flancloud
FortiLAN cloud manager configuration. Applies to FortiSwitch Version `>= 7.0.3`.

## Example Usage

```hcl
resource "fortiswitch_system_flancloud" "name" {
        interval = "3"
        name = "default_name_4"
        port = "5"
        status = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable FortiLAN-cloud service. Valid values: `enable`, `disable`.
* `interval` - Service name resolution time interval (3-300sec, default=3).
* `port` - Port Number.
* `name` - Fully qualified domain name or IP address of FortiLAN-cloud service.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FlanCloud can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_flancloud.labelname SystemFlanCloud

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_flancloud.labelname SystemFlanCloud
$ unset "FORTISWITCH_IMPORT_TABLE"
```
