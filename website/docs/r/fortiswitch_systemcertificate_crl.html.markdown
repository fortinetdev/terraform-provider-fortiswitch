---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemcertificate_crl"
description: |-
  Certificate Revokation List.
---

# fortiswitch_systemcertificate_crl
Certificate Revokation List.

## Example Usage

```hcl
resource "fortiswitch_systemcertificate_crl" "scrl" {
        scep_url = "http://www.test.com"
        ldap_username = "admin"
        name = "test"
        update_interval= "11"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `http_url` - URL of HTTP server for CRL update.
* `ldap_username` - Login name for LDAP server.
* `update_vdom` - Virtual domain for CRL update.
* `ldap_server` - LDAP server.
* `update_interval` - Second between updates, 0=disabled.
* `ldap_password` - Login password for LDAP server.
* `scep_cert` - Local certificate used for CRL update via SCEP.
* `scep_url` - URL of CA server for CRL update via SCEP.
* `crl` - Certificate Revocation List.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemCertificate Crl can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemcertificate_crl.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemcertificate_crl.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
