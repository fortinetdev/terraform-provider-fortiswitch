---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemcertificate_ocsp"
description: |-
  Ocsp configuration.
---

# fortiswitch_systemcertificate_ocsp
Ocsp configuration.

## Argument Reference

The following arguments are supported:

* `url` - URL to OCSP server.
* `cert` - OCSP server certificate.
* `unavail_action` - Action when server is unavailable. Valid values: `revoke`, `ignore`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SystemCertificate Ocsp can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemcertificate_ocsp.labelname SystemCertificateOcsp

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemcertificate_ocsp.labelname SystemCertificateOcsp
$ unset "FORTISWITCH_IMPORT_TABLE"
```
