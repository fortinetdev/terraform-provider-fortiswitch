---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_web"
description: |-
  Configure web attributes.
---

# fortiswitch_system_web
Configure web attributes. Applies to FortiSwitch Version `>= 7.2.0`.

## Argument Reference

The following arguments are supported:

* `https_server_cert` - Administrative HTTPS server certificate.
* `https_port` - Administrative access HTTPS port (1 - 65535).
* `https_pki_required` - Enable/disable HTTPS login page when PKI is enabled. Valid values: `enable`, `disable`.
* `http_port` - Administrative access HTTP port (1 - 65535).
* `gui_language` - Web display language. Valid values: `browser`, `english`, `simch`, `japanese`, `korean`, `spanish`, `trach`, `french`, `portuguese`, `german`.
* `https_ssl_versions` - Allowed SSL/TLS versions for web administration. Valid values: `tlsv1-0`, `tlsv1-1`, `tlsv1-2`, `tlsv1-3`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Web can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_web.labelname SystemWeb

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_web.labelname SystemWeb
$ unset "FORTISWITCH_IMPORT_TABLE"
```
