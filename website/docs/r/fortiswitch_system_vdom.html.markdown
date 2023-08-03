---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_vdom"
description: |-
  Virtual domain configuration.
---

# fortiswitch_system_vdom
Virtual domain configuration.

## Argument Reference

The following arguments are supported:

* `vcluster_id` - Vcluster-id.
* `name` - Vd name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Vdom can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_vdom.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_vdom.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
