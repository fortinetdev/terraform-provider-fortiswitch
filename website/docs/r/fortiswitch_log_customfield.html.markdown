---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_log_customfield"
description: |-
  Custom field configuation.
---

# fortiswitch_log_customfield
Custom field configuation.

## Example Usage

```hcl
resource "fortiswitch_log_customfield" "name" {
        fswid =  "3"
        name = "maxx"
        value = "2"
}
```

## Argument Reference

The following arguments are supported:

* `fswid` - ID.
* `value` - Value.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

Log CustomField can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_log_customfield.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_log_customfield.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
