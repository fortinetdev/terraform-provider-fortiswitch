---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_fm"
description: |-
  Fm.
---

# fortiswitch_system_fm
Fm.

## Example Usage

```hcl
resource "fortiswitch_system_fm" "name" {
        auto_backup = "enable"
        fswid =  "4"
        ip = "3.4.5.6"
        ipsec = "enable"
        scheduled_config_restore = "enable"
        status = "enable"
        vdom = "root"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Status. Valid values: `enable`, `disable`.
* `auto_backup` - Auto-backup. Valid values: `enable`, `disable`.
* `ip` - IP.
* `scheduled_config_restore` - Scheduled-config-restore. Valid values: `enable`, `disable`.
* `fswid` - Id.
* `vdom` - Vdom.
* `ipsec` - IPSec. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fm can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_fm.labelname SystemFm

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_fm.labelname SystemFm
$ unset "FORTISWITCH_IMPORT_TABLE"
```
