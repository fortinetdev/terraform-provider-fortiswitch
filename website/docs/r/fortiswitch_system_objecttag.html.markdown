---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_objecttag"
description: |-
  Object tags.
---

# fortiswitch_system_objecttag
Object tags.

## Example Usage

```hcl
resource "fortiswitch_system_objecttag" "name" {
         name = "default_name_3"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Tag name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System ObjectTag can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_objecttag.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_objecttag.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
