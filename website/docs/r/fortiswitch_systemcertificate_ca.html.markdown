---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemcertificate_ca"
description: |-
  CA certificate.
---

# fortiswitch_systemcertificate_ca
CA certificate.

## Example Usage

```hcl
resource "fortiswitch_systemcertificate_ca" "scc" {
        auto_update_days = "9"
        auto_update_days_warning = "9"
        name = "Fortinet_802.1x_CA"
        scep_url = "http =//www.sldf.com"
}
```

## Argument Reference

The following arguments are supported:

* `info` - Info about name.
* `scep_url` - URL of SCEP server.
* `ca` - CA certificate.
* `auto_update_days_warning` - Days to send update before auto-update, 0=disabled.
* `name` - Name.
* `auto_update_days` - Days to auto-update before expired, 0=disabled.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemCertificate Ca can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemcertificate_ca.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemcertificate_ca.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
