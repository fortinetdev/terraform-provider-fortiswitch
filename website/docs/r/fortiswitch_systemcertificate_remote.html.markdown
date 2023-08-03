---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemcertificate_remote"
description: |-
  Remote certificate.
---

# fortiswitch_systemcertificate_remote
Remote certificate.

## Example Usage

```hcl
resource "fortiswitch_systemcertificate_remote" "scr" {
        name = "test"
        remote = "name"
}
```

## Argument Reference

The following arguments are supported:

* `remote` - Remote certificate.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemCertificate Remote can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemcertificate_remote.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemcertificate_remote.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
