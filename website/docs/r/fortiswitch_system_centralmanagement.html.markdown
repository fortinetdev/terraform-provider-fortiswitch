---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_centralmanagement"
description: |-
  Central management configuration.
---

# fortiswitch_system_centralmanagement
Central management configuration.

## Example Usage

```hcl
resource "fortiswitch_system_centralmanagement" "trname1" {
  allow_monitor                 = "enable"
  allow_push_configuration      = "enable"
  allow_push_firmware           = "enable"
  allow_remote_firmware_upgrade = "enable"
  enc_algorithm                 = "high"
  fmg                           = "0.0.0.0"
  fmg_source_ip6                = "::"
  include_default_servers       = "enable"
  mode                          = "normal"
  schedule_config_restore       = "enable"
  schedule_script_restore       = "enable"
  type                          = "fortimanager"
  vdom                          = "root"
}

resource "fortiswitch_system_centralmanagement" "trname2" {
  allow_monitor                 = "enable"
  allow_push_configuration      = "enable"
  allow_push_firmware           = "enable"
  allow_remote_firmware_upgrade = "enable"
  enc_algorithm                 = "high"
  fmg                           = "\"192.168.52.177\""
  include_default_servers       = "enable"
  mode                          = "normal"
  type                          = "fortimanager"
  vdom                          = "root"
}

```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable management. Valid values: `enable`, `disable`.
* `enc_algorithm` - Whether to use SSL encryption. Valid values: `default`, `high`, `low`.
* `fmg_source_ip` - Source address to use when connecting fortimanager.
* `schedule_config_restore` - Enable/disable schedule config restore. Valid values: `enable`, `disable`.
* `allow_pushd_firmware` - Enable/disable push firmware. Valid values: `enable`, `disable`.
* `allow_push_configuration` - Enable/disable push configuration. Valid values: `enable`, `disable`.
* `allow_remote_firmware_upgrade` - Enable/disable push firmware. Valid values: `enable`, `disable`.
* `serial_number` - Serial number.
* `mode` - Normal/backup management mode. Valid values: `normal`, `backup`.
* `schedule_script_restore` - Enable/disable schedule config restore. Valid values: `enable`, `disable`.
* `fmg` - Address of fortimanager(ip or fqdn name).
* `allow_monitor` - Enable/disable remote monitor of device. Valid values: `enable`, `disable`.
* `type` - Type of management server. Valid values: `fortimanager`, `fortiguard`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System CentralManagement can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_centralmanagement.labelname SystemCentralManagement

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_centralmanagement.labelname SystemCentralManagement
$ unset "FORTISWITCH_IMPORT_TABLE"
```
