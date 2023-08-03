---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemdhcp_server"
description: |-
  Get information on an fortiswitch systemdhcp server.
---

# Data Source: fortiswitch_systemdhcp_server
Use this data source to get information on an fortiswitch systemdhcp server

## Argument Reference

* `fswid` - (Required) Specify the fswid of the desired systemdhcp server.

## Attribute Reference

The following attributes are exported:

* `domain` - Domain name suffix for the IP addresses that the DHCP server assigns to clients.
* `lease_time` - Lease time in seconds, 0 means unlimited.
* `exclude_range` - Exclude one or more ranges of IP addresses from being assigned to clients. The structure of `exclude_range` block is documented below.
* `server_type` - DHCP server can be a normal DHCP server or an IPsec DHCP server.
* `conflicted_ip_timeout` - Time in seconds to wait after a conflicted IP address is removed from the DHCP range before it can be reused.
* `timezone_option` - Options for the DHCP server to set the client's time zone.
* `timezone` - Select the time zone to be assigned to DHCP clients.
* `fswid` - ID.
* `filename` - Name of the boot file on the TFTP server.
* `ntp_server1` - NTP server 1.
* `default_gateway` - Default gateway IP address assigned by the DHCP server.
* `next_server` - IP address of a server (for example, a TFTP sever) that DHCP clients can download a boot file from.
* `ntp_server2` - NTP server 2.
* `dns_service` - Options for assigning DNS servers to DHCP clients.
* `ip_range` - DHCP IP range configuration. The structure of `ip_range` block is documented below.
* `ip_mode` - Method used to assign client IP.
* `auto_configuration` - Enable/disable auto configuration.
* `status` - Enable/disable this DHCP configuration.
* `reserved_address` - Options for the DHCP server to assign IP settings to specific MAC addresses. The structure of `reserved_address` block is documented below.
* `ntp_server3` - NTP server 3.
* `netmask` - Netmask assigned by the DHCP server.
* `ntp_service` - Options for assigning Network Time Protocol (NTP) servers to DHCP clients.
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
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served.

The `exclude_range` block contains:

* `id` - ID.
* `start_ip` - Start of IP range.
* `end_ip` - End of IP range.

The `ip_range` block contains:

* `id` - ID.
* `start_ip` - Start of IP range.
* `end_ip` - End of IP range.

The `reserved_address` block contains:

* `description` - Description.
* `ip` - IP address to be reserved for the MAC address.
* `mac` - MAC address of the client that will get the reserved IP address.
* `circuit_id_type` - DHCP option type.
* `circuit_id` - Option 82 circuit-ID of the client that will get the reserved IP address.
* `action` - Options for the DHCP server to configure the client with the reserved MAC address.
* `remote_id` - Option 82 remote-ID of the client that will get the reserved IP address.
* `type` - DHCP reserved-address type.
* `id` - ID.
* `remote_id_type` - DHCP option type.

The `options` block contains:

* `ip` - DHCP option IPs.
* `code` - DHCP option code.
* `type` - DHCP option type.
* `id` - ID.
* `value` - DHCP option value.

The `vci_string` block contains:

* `vci_string` - VCI strings.

The `tftp_server` block contains:

* `tftp_server` - TFTP server.

