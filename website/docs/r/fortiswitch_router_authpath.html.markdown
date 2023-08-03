---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_authpath"
description: |-
  Auth-based routing configuration.
---

# fortiswitch_router_authpath
Auth-based routing configuration.

## Example Usage

```hcl
resource "fortiswitch_router_authpath" "name" {
        device = "port5"
        gateway = "192.160.22.1"
        name = "default_name_5"
}
```

## Argument Reference

The following arguments are supported:

* `device` - Output interface.
* `name` - Name of the entry.
* `gateway` - Gateway IP address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router AuthPath can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_authpath.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_authpath.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
