---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemdhcp_server"
description: |-
  Configure DHCP servers.
---

# fortiswitch_systemdhcp_server
Configure DHCP servers.

## Example Usage

```hcl
resource "fortiswitch_systemdhcp_server" "sdhcp" {
        auto_configuration = "enable"
        conflicted_ip_timeout = "8640000"
        exclude_range {
            end_ip = "1.1.1.3"
            id =  "1"
            start_ip = "1.1.1.2"
        }
        fswid =  "1"
        interface = "mgmt"
        ip_mode = "range"
        ip_range {
            end_ip = "1.1.1.4"
            id =  "1"
            start_ip = "1.1.1.1"
        }
        lease_time = "8640000"
        netmask = "255.255.255.0"
        ntp_server1 = "1.1.1.1"
        options {
            code = "31"
            id =  "32"
            ip = "1.1.1.1"
            type = "hex"
            value = "23"
        }
        wifi_ac1 = "2.2.2.2"
        wins_server1 = "3.3.3.3"
}
```

## Argument Reference

The following arguments are supported:

* `domain` - Domain name suffix for the IP addresses that the DHCP server assigns to clients.
* `lease_time` - Lease time in seconds, 0 means unlimited.
* `exclude_range` - Exclude one or more ranges of IP addresses from being assigned to clients. The structure of `exclude_range` block is documented below.
* `server_type` - DHCP server can be a normal DHCP server or an IPsec DHCP server. Valid values: `regular`.
* `conflicted_ip_timeout` - Time in seconds to wait after a conflicted IP address is removed from the DHCP range before it can be reused.
* `timezone_option` - Options for the DHCP server to set the client's time zone. Valid values: `disable`, `default`, `specify`.
* `timezone` - Select the time zone to be assigned to DHCP clients. Valid values: `01`, `02`, `03`, `04`, `05`, `81`, `06`, `07`, `08`, `09`, `10`, `11`, `12`, `13`, `74`, `14`, `77`, `15`, `87`, `16`, `17`, `18`, `19`, `20`, `75`, `21`, `22`, `23`, `24`, `80`, `79`, `25`, `26`, `27`, `28`, `78`, `29`, `30`, `31`, `32`, `33`, `34`, `35`, `36`, `37`, `38`, `83`, `84`, `40`, `85`, `41`, `42`, `43`, `39`, `44`, `46`, `47`, `51`, `48`, `45`, `49`, `50`, `52`, `53`, `54`, `55`, `56`, `57`, `58`, `59`, `60`, `62`, `63`, `61`, `64`, `65`, `66`, `67`, `68`, `69`, `70`, `71`, `72`, `00`, `82`, `73`, `86`, `76`.
* `fswid` - ID.
* `filename` - Name of the boot file on the TFTP server.
* `ntp_server1` - NTP server 1.
* `default_gateway` - Default gateway IP address assigned by the DHCP server.
* `next_server` - IP address of a server (for example, a TFTP sever) that DHCP clients can download a boot file from.
* `ntp_server2` - NTP server 2.
* `dns_service` - Options for assigning DNS servers to DHCP clients. Valid values: `local`, `default`, `specify`.
* `ip_range` - DHCP IP range configuration. The structure of `ip_range` block is documented below.
* `ip_mode` - Method used to assign client IP. Valid values: `range`, `usrgrp`.
* `auto_configuration` - Enable/disable auto configuration. Valid values: `disable`, `enable`.
* `status` - Enable/disable this DHCP configuration. Valid values: `disable`, `enable`.
* `reserved_address` - Options for the DHCP server to assign IP settings to specific MAC addresses. The structure of `reserved_address` block is documented below.
* `ntp_server3` - NTP server 3.
* `netmask` - Netmask assigned by the DHCP server.
* `ntp_service` - Options for assigning Network Time Protocol (NTP) servers to DHCP clients. Valid values: `local`, `default`, `specify`.
* `wins_server1` - WINS server 1.
* `wins_server2` - WINS server 2.
* `interface` - DHCP server can assign IP configurations to clients connected to this interface.
* `wifi_ac1` - WiFi Access Controller 1 IP address (DHCP option 138, RFC 5417).
* `wifi_ac2` - WiFi Access Controller 2 IP address (DHCP option 138, RFC 5417).
* `wifi_ac3` - WiFi Access Controller 3 IP address (DHCP option 138, RFC 5417).
* `options` - DHCP options. The structure of `options` block is documented below.
* `vci_string` - One or more VCI strings in quotes separated by spaces. The structure of `vci_string` block is documented below.
* `dns_server2` - DNS server 2.
* `dns_server3` - DNS server 3.
* `dns_server1` - DNS server 1.
* `tftp_server` - One or more hostnames or IP addresses of the TFTP servers in quotes separated by spaces. The structure of `tftp_server` block is documented below.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served. Valid values: `disable`, `enable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `exclude_range` block supports:

* `id` - ID.
* `start_ip` - Start of IP range.
* `end_ip` - End of IP range.

The `ip_range` block supports:

* `id` - ID.
* `start_ip` - Start of IP range.
* `end_ip` - End of IP range.

The `reserved_address` block supports:

* `description` - Description.
* `ip` - IP address to be reserved for the MAC address.
* `mac` - MAC address of the client that will get the reserved IP address.
* `circuit_id_type` - DHCP option type. Valid values: `hex`, `string`.
* `circuit_id` - Option 82 circuit-ID of the client that will get the reserved IP address.
* `action` - Options for the DHCP server to configure the client with the reserved MAC address. Valid values: `assign`, `block`, `reserved`.
* `remote_id` - Option 82 remote-ID of the client that will get the reserved IP address.
* `type` - DHCP reserved-address type. Valid values: `mac`, `option82`.
* `id` - ID.
* `remote_id_type` - DHCP option type. Valid values: `hex`, `string`.

The `options` block supports:

* `ip` - DHCP option IPs.
* `code` - DHCP option code.
* `type` - DHCP option type. Valid values: `hex`, `string`, `ip`, `fqdn`.
* `id` - ID.
* `value` - DHCP option value.

The `vci_string` block supports:

* `vci_string` - VCI strings.

The `tftp_server` block supports:

* `tftp_server` - TFTP server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

SystemDhcp Server can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemdhcp_server.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemdhcp_server.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
