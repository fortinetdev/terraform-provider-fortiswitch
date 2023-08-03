---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemcertificate_local"
description: |-
  Local keys and certificates.
---

# fortiswitch_systemcertificate_local
Local keys and certificates.

## Example Usage

```hcl
resource "fortiswitch_systemcertificate_local" "scrl" {
        auto_regenerate_days = "9"
        auto_regenerate_days_warning = "9"
        comments = "test"
        name = "Fortinet_802.1x"
        scep_url = "http =//www.test.com"
}
```

## Argument Reference

The following arguments are supported:

* `scep_password` - SCEP server challenge password for auto-regeneration.
* `name` - Name.
* `certificate` - Certificate.
* `private_key` - Private-key.
* `comments` - Comments.
* `auto_regenerate_days` - Days to auto-regenerate before expired, 0=disabled.
* `auto_regenerate_days_warning` - Days to send warning before auto-regenerateion, 0=disabled.
* `name_encoding` - Name encoding for auto-regeneration. Valid values: `printable`, `utf8`.
* `scep_url` - URL of SCEP server.
* `password` - Password.
* `csr` - Certificate Signing Request.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemCertificate Local can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemcertificate_local.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemcertificate_local.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
