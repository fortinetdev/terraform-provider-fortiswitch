---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_managementtunnel"
description: |-
  Management tunnel configuration.
---

# fortiswitch_system_managementtunnel
Management tunnel configuration.

## Example Usage

```hcl
resource "fortiswitch_system_managementtunnel" "name" {
        allow_collect_statistics = "enable"
        allow_config_restore = "enable"
        allow_push_configuration = "enable"
        allow_push_firmware = "enable"
        authorized_manager_only = "enable"
        serial_number = "S548DN4K16000234"
        status = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable fgfm tunnel. Valid values: `enable`, `disable`.
* `allow_push_firmware` - Enable/disable push firmware. Valid values: `enable`, `disable`.
* `allow_collect_statistics` - Enable/disable collect run time statistics. Valid values: `enable`, `disable`.
* `allow_push_configuration` - Enable/disable push configuration. Valid values: `enable`, `disable`.
* `serial_number` - Serial number.
* `authorized_manager_only` - Enable/disable restricted to authorized manager only. Valid values: `enable`, `disable`.
* `allow_config_restore` - Enable/disable allow config restore. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System ManagementTunnel can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_managementtunnel.labelname SystemManagementTunnel

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_managementtunnel.labelname SystemManagementTunnel
$ unset "FORTISWITCH_IMPORT_TABLE"
```
