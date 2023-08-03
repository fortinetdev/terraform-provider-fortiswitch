---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchlldp_settings"
description: |-
  Global LLDP configuration.
---

# fortiswitch_switchlldp_settings
Global LLDP configuration.

## Example Usage

```hcl
resource "fortiswitch_switchlldp_settings" "name" {
        device_detection = "disable"
        fast_start_interval = "4"
        management_interface = "port4"
        status = "disable"
        tx_hold = "7"
        tx_interval = "8"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable LLDP. Valid values: `disable`, `enable`.
* `device_detection` - Enable/disable dynamic updates of LLDP neighbor devices to fortilink. Valid values: `disable`, `enable`.
* `management_address` - Advertise IPv4 and/or IPv6 address. Valid values: `ipv4`, `ipv6`, `none`.
* `tx_hold` - Number of TX intervals before local LLDP data expires(tx-hold x tx-interval = packet TTL).
* `tx_interval` - Frequency of LLDP PDU transmit (seconds).
* `management_interface` - Primary management interface to be advertised.
* `fast_start_interval` - Frequency of LLDP PDU transmit for the first 4 packets on link up(seconds).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchLldp Settings can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchlldp_settings.labelname SwitchLldpSettings

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchlldp_settings.labelname SwitchLldpSettings
$ unset "FORTISWITCH_IMPORT_TABLE"
```
