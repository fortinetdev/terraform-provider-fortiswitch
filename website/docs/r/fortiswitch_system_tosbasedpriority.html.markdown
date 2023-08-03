---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_tosbasedpriority"
description: |-
  Configure tos based priority table.
---

# fortiswitch_system_tosbasedpriority
Configure tos based priority table.

## Example Usage

```hcl
resource "fortiswitch_system_tosbasedpriority" "name" {
        fswid =  "3"
        priority = "low"
        tos = "5"
}
```

## Argument Reference

The following arguments are supported:

* `priority` - Tos based priority level. Valid values: `low`, `medium`, `high`.
* `tos` - IP tos value [0-15].
* `fswid` - Item id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

System TosBasedPriority can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_tosbasedpriority.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_tosbasedpriority.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
