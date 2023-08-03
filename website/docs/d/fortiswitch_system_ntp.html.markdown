---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_ntp"
description: |-
  Get information on fortiswitch system ntp.
---

# Data Source: fortiswitch_system_ntp
Use this data source to get information on fortiswitch system ntp

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `key_type` - Key type for authentication (MD5, SHA1).
* `log_time_adjustments` - Enable/disable logging of NTP time adjustments.
* `server_mode` - On Enabling, your FortiSwitch relays NTP requests to its configured NTP server.
* `source_ip` - Source IP for communications to NTP server.
* `allow_unsync_source` - Enable/disable allowance of unsynchronized NTP server source.
* `authentication` - Enable/disable authentication.
* `source_ip6` - Source IPv6 address for communication to the NTP server.
* `key` - Key for authentication.
* `ntpsync` - Enable/disable synchronization with NTP server.
* `syncinterval` - NTP synchronization interval (1 - 1440) min.
* `ntpserver` - NTP server. The structure of `ntpserver` block is documented below.
* `key_id` - Key ID for authentication.

The `ntpserver` block contains:

* `ntpv3` - Enable to use NTPv3 instead of NTPv4.
* `server` - IP address/hostname of NTP server.
* `authentication` - Enable/disable MD5(NTPv3)/SHA1(NTPv4) authentication.
* `key` - Key for MD5(NTPv3)/SHA1(NTPv4) authentication.
* `id` - Time server ID.
* `key_id` - Key ID for authentication.

