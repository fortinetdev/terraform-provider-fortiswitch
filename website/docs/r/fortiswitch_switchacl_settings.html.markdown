---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchacl_settings"
description: |-
  Configure access-control lists global settings on Switch.
---

# fortiswitch_switchacl_settings
Configure access-control lists global settings on Switch.

## Example Usage

```hcl
resource "fortiswitch_switchacl_settings" "name" {
        density_mode = "enable"
        trunk_load_balance = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `density_mode` - Density mode enable/disable. Valid values: `enable`, `disable`.
* `trunk_load_balance` - Enable/disable trunk-load-balancing for ACL actions. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchAcl Settings can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchacl_settings.labelname SwitchAclSettings

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchacl_settings.labelname SwitchAclSettings
$ unset "FORTISWITCH_IMPORT_TABLE"
```
