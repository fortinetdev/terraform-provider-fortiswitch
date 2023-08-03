---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_bugreport"
description: |-
  Configure bug report.
---

# fortiswitch_system_bugreport
Configure bug report.

## Example Usage

```hcl
resource "fortiswitch_system_bugreport" "sbr" {
        auth = "yes"
        mailto = "asdf@for.com"
        password = "123"
        server = "192.168.100.40"
        username = "maxx"
}
```

## Argument Reference

The following arguments are supported:

* `mailto` - Recipient.
* `username` - Username on this SMTP server.
* `auth` - SMTP server authentication. Valid values: `yes`, `no`.
* `server` - SMTP server name.
* `username_smtp` - Username used for authentication on SMTP server.
* `password` - Password on this SMTP server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System BugReport can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_bugreport.labelname SystemBugReport

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_bugreport.labelname SystemBugReport
$ unset "FORTISWITCH_IMPORT_TABLE"
```
