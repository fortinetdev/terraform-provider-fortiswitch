---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_ntp"
description: |-
  Ntp system info configuration.
---

# fortiswitch_system_ntp
Ntp system info configuration.

## Example Usage

```hcl
resource "fortiswitch_system_ntp" "name" {
        allow_unsync_source = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `key_type` - Key type for authentication (MD5, SHA1). Valid values: `MD5`, `SHA1`.
* `log_time_adjustments` - Enable/disable logging of NTP time adjustments. Valid values: `enable`, `disable`.
* `server_mode` - On Enabling, your FortiSwitch relays NTP requests to its configured NTP server. Valid values: `enable`, `disable`.
* `source_ip` - Source IP for communications to NTP server.
* `allow_unsync_source` - Enable/disable allowance of unsynchronized NTP server source. Valid values: `enable`, `disable`.
* `authentication` - Enable/disable authentication. Valid values: `enable`, `disable`.
* `source_ip6` - Source IPv6 address for communication to the NTP server.
* `key` - Key for authentication.
* `ntpsync` - Enable/disable synchronization with NTP server. Valid values: `enable`, `disable`.
* `syncinterval` - NTP synchronization interval (1 - 1440) min.
* `ntpserver` - NTP server. The structure of `ntpserver` block is documented below.
* `key_id` - Key ID for authentication.
* `interface` - FortiSwitch interface(s) with NTP server mode enabled. Devices on your network can contact these interfaces for NTP services. The structure of `interface` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `ntpserver` block supports:

* `ntpv3` - Enable to use NTPv3 instead of NTPv4. Valid values: `enable`, `disable`.
* `server` - IP address/hostname of NTP server.
* `authentication` - Enable/disable MD5(NTPv3)/SHA1(NTPv4) authentication. Valid values: `enable`, `disable`.
* `key` - Key for MD5(NTPv3)/SHA1(NTPv4) authentication.
* `id` - Time server ID.
* `key_id` - Key ID for authentication.

The `interface` block supports:

* `interface_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ntp can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_ntp.labelname SystemNtp

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_ntp.labelname SystemNtp
$ unset "FORTISWITCH_IMPORT_TABLE"
```
