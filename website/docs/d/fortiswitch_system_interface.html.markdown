---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_interface"
description: |-
  Get information on an fortiswitch system interface.
---

# Data Source: fortiswitch_system_interface
Use this data source to get information on an fortiswitch system interface

## Argument Reference

* `name` - (Required) Specify the name of the desired system interface.

## Attribute Reference

The following attributes are exported:

* `defaultgw` - Enable/disable default gateway.
* `gwdetect` - Enable/disable gateway detection.
* `dhcp_client_status` - DHCP client connection status.
* `weight` - Default weight for static routes if route has no weight configured (0 - 255).
* `ip` - Interface IPv4 address.
* `vrrp_virtual_mac` - enable to use virtual MAC for VRRP
* `bfd_detect_mult` - BFD detection multiplier.
* `bfd_required_min_rx` - BFD required minimal receive interval.
* `l2_interface` - L2 interface name.
* `src_check_allow_default` - Enable/disable.When src ip lookup hits default route,enable means allow pkt else drop.
* `dhcp_relay_ip` - DHCP relay IP address.
* `forward_domain` - TP mode forward domain.
* `speed` - Speed (copper mode port only).
* `vlanforward` - Enable/disable VLAN forwarding.
* `priority` - Priority of learned routes.
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `macaddr` - MAC address.
* `bfd_desired_min_tx` - BFD desired minimal transmit interval.
* `switch` - Contained in switch.
* `vlanid` - VLAN ID.
* `cli_conn_status` - CLI connection status.
* `detectserver` - IP address to PING for gateway detection.
* `vrrp` - VRRP configuration The structure of `vrrp` block is documented below.
* `allowaccess` - Interface management access.
* `dhcp_client_identifier` - DHCP client identifier.
* `ping_serv_status` - PING server status.
* `type` - Interface type.
* `snmp_index` - SNMP index.
* `icmp_redirect` - Enable/disable ICMP rediect.
* `dhcp_relay_option82` - Enable / Disable DHCP relay option-82 insertion.
* `description` - Description.
* `remote_ip` - Remote IP address of tunnel.
* `dns_server_override` - Enable/disable use of DNS server aquired by DHCP or PPPoE.
* `secondary_ip` - Enable/disable use of secondary IP address.
* `vrf` - VRF.
* `dhcp_expire` - DHCP client expiration.
* `interface` - Interface name.
* `dhcp_vendor_specific_option` - DHCP Vendor specific option 43.
* `vdom` - Virtual domain name.
* `dhcp_relay_service` - Enable/disable use DHCP relay service.
* `distance` - Distance of learned routes.
* `name` - Name.
* `detectprotocol` - Protocol to use for gateway detection.
* `src_check` - Enable/disable source IP check.
* `ipv6` - IPv6 address. The structure of `ipv6` block is documented below.
* `dynamicgw` - Dynamic gateway.
* `mtu_override` - Enable/disable override of default MTU.
* `mtu` - Maximum transportation unit (MTU).
* `dynamic_dns1` - Primary dynamic DNS server.
* `alias` - Alias.
* `auth_type` - PPP authentication type.
* `dynamic_dns2` - Secondary dynamic DNS server.
* `status` - Interface status.
* `mode` - Interface addressing mode.
* `ha_priority` - PING server HA election priority (1 - 50).
* `secondaryip` - Second IP address of interface. The structure of `secondaryip` block is documented below.
* `switch_members` - Switch interfaces. The structure of `switch_members` block is documented below.

The `vrrp` block contains:

* `status` - Enable/disable status.
* `priority` - Priority of the virtual router (1 - 255).
* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `start_time` - Startup time (1 - 255 seconds).
* `backup_vmac_fwd` - Enable/disable backup-vmac-fwd.
* `vrid` - Virtual router identifier (1 - 255).
* `vrgrp` - VRRP group ID (1 - 65535).
* `preempt` - Enable/disable preempt mode.
* `version` - VRRP version.
* `vrdst` - Monitor the route to this destination.
* `vrip` - IP address of the virtual router.
* `netmask` - Netmask of the virtual router.

The `ipv6` block contains:

* `ip6_allowaccess` - Allow management access to the interface.
* `ip6_retrans_time` - IPv6 retransmit time (milliseconds).
* `vrrp6` - IPv6 VRRP configuration. The structure of `vrrp6` block is documented below.
* `ip6_other_flag` - Enable/disable sending of IPv6 other flag.
* `ip6_dns_server_override` - Enable/disable using the DNS server acquired by DHCP.
* `vrip6_link_local` - Link-local IPv6 address of virtual router.
* `ip6_address` - Primary IPv6 address prefix of interface.
* `ip6_prefix_list` - IPv6 advertised prefix list. The structure of `ip6_prefix_list` block is documented below.
* `ip6_link_mtu` - IPv6 link MTU.
* `ip6_manage_flag` - Enable/disable sending of IPv6 managed flag.
* `ip6_min_interval` - IPv6 minimum interval (sec) after which RA will be sent.
* `ip6_mode` - Addressing mode (static, DHCP).
* `ip6_hop_limit` - IPv6 hop limit.
* `ip6_reachable_time` - IPv6 reachable time (milliseconds).
* `ip6_extra_addr` - Extra IPv6 address prefixes of interface. The structure of `ip6_extra_addr` block is documented below.
* `ip6_default_life` - IPv6 default life (sec).
* `ip6_max_interval` - IPv6 maximum interval (sec) after which RA will be sent.
* `vrrp_virtual_mac6` - Enable/disable virtual MAC for VRRP.
* `ip6_unknown_mcast_to_cpu` - Enable/disable unknown mcast to cpu.
* `ip6_send_adv` - Enable/disable sending of IPv6 Router advertisement.
* `autoconf` - Enable/disable address automatic config.
* `dhcp6_information_request` - Enable/disable DHCPv6 information request.

The `vrrp6` block contains:

* `status` - Enable/disable VRRP.
* `priority` - Priority of the virtual router (1 - 255).
* `vrip6` - IPv6 address of the virtual router.
* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `start_time` - Startup time (1 - 255 seconds).
* `vrid` - Virtual router identifier (1 - 255).
* `vrgrp` - VRRP group ID (1 - 65535).
* `preempt` - Enable/disable preempt mode.
* `accept_mode` - Enable/disable accept mode.
* `vrdst6` - Monitor the route to this destination.

The `ip6_prefix_list` block contains:

* `valid_life_time` - Valid life time (sec).
* `prefix` - IPv6 prefix.
* `autonomous_flag` - Enable/disable autonomous flag.
* `onlink_flag` - Enable/disable onlink flag.
* `preferred_life_time` - Preferred life time (sec).

The `ip6_extra_addr` block contains:

* `prefix` - IPv6 address prefix.

The `secondaryip` block contains:

* `gwdetect` - Enable/disable gateway detection.
* `detectprotocol` - Protocol to use for gateway detection.
* `ip` - Interface IPv4 address.
* `detectserver` - IP address to PING for gateway detection.
* `allowaccess` - Interface management access.
* `ping_serv_status` - PING server status.
* `ha_priority` - PING server HA election priority (1 - 50).
* `id` - Id.

The `switch_members` block contains:

* `member_name` - Interface name.

