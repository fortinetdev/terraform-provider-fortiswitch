---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_vlan"
description: |-
  Configure optional per-VLAN settings.
---

# fortiswitch_switch_vlan
Configure optional per-VLAN settings.

## Argument Reference

The following arguments are supported:

* `community_vlans` - Communities within this private VLAN.
* `lan_subvlans` - LAN segment subvlans.
* `arp_inspection` - Enable/Disable Dynamic ARP Inspection.
* `igmp_snooping_proxy` - Enable/disable IGMP snooping proxy for the VLAN interface. Valid values: `enable`, `disable`.
* `policer` - Set policer on the VLAN traffic.
* `dhcp6_snooping` - Enable/Disable DHCPv6 snooping on this vlan. Valid values: `disable`, `enable`.
* `lan_segment_primary_vlan` - LAN Segment Primary VLAN ID.
* `learning` - Enable/disable L2 learning on this VLAN. Valid values: `disable`, `enable`.
* `fswid` - VLAN ID.
* `igmp_snooping_querier` - Enable/disable IGMP-snooping-querier for the VLAN interface. Valid values: `enable`, `disable`.
* `dhcp_snooping_static_client` - DHCP Snooping static clients. The structure of `dhcp_snooping_static_client` block is documented below.
* `member_by_mac` - Assign VLAN membership based on MAC address. The structure of `member_by_mac` block is documented below.
* `lan_segment` - Enable/disable LAN Segment. Valid values: `enable`, `disable`.
* `primary_vlan` - Primary VLAN ID.
* `igmp_snooping_querier_version` - IGMP-snooping-querier version.
* `cos_queue` - Set cos(0-7) on the VLAN traffic or unset to disable.
* `member_by_ipv6` - Assign VLAN membership based on IPv6 prefix. The structure of `member_by_ipv6` block is documented below.
* `mld_snooping_proxy` - Enable/disable MLD snooping proxy for the VLAN interface. Valid values: `enable`, `disable`.
* `mld_snooping_fast_leave` - Enable/disable MLD snooping fast leave. Valid values: `enable`, `disable`.
* `igmp_snooping` - Enable/disable IGMP-snooping for the VLAN interface. Valid values: `enable`, `disable`.
* `mld_snooping_querier_addr` - MLD-querier address.
* `dhcp_server_access_list` - Configure dhcp server access list. The structure of `dhcp_server_access_list` block is documented below.
* `learning_limit` - Limit the number of dynamic MAC addresses on this VLAN.
* `mrouter_ports` - Member interfaces. The structure of `mrouter_ports` block is documented below.
* `mld_snooping_static_group` - MLD static groups. The structure of `mld_snooping_static_group` block is documented below.
* `description` - Description.
* `igmp_snooping_querier_addr` - IGMP-snooping-querier address.
* `rspan_mode` - Stop L2 learning and interception of BPDUs and other packets on this VLAN. Valid values: `enable`, `disable`.
* `private_vlan_type` - Private VLAN type.
* `mld_snooping_querier` - Enable/disable MLD snooping querier for the VLAN interface. Valid values: `enable`, `disable`.
* `igmp_snooping_fast_leave` - Enable/disable IGMP snooping fast leave. Valid values: `enable`, `disable`.
* `dhcp_snooping` - Enable/Disable dhcp snooping on this vlan. Valid values: `disable`, `enable`.
* `lan_segment_type` - LAN segment type.
* `igmp_snooping_static_group` - IGMP static groups. The structure of `igmp_snooping_static_group` block is documented below.
* `dhcp_snooping_verify_mac` - Enable/Disable verify source mac. Valid values: `disable`, `enable`.
* `isolated_vlan` - Isolated VLAN.
* `dhcp_snooping_option82` - Enable/Disable inserting option82. Valid values: `disable`, `enable`.
* `member_by_ipv4` - Assign VLAN membership based on IPv4 address or subnet. The structure of `member_by_ipv4` block is documented below.
* `private_vlan` - Enable/disable private VLAN. Valid values: `enable`, `disable`.
* `member_by_proto` - Assign VLAN membership based on ethernet frametype and protocol. The structure of `member_by_proto` block is documented below.
* `access_vlan` - Block port-to-port traffic. Valid values: `disable`, `enable`.
* `assignment_priority` - 802.1x Radius (Tunnel-Private-Group-Id) vlanid assign-by-name priority (smaller is higher).
* `mld_snooping` - Enable/disable MLD snooping for the VLAN interface. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `dhcp_snooping_static_client` block supports:

* `ip_addr` - Client IPv4 address.
* `mac_addr` - Client MAC address.
* `name` - Client Name.
* `switch_interface` - Interface name.

The `member_by_mac` block supports:

* `mac` - MAC address.
* `id` - Entry ID.
* `description` - Description.

The `member_by_ipv6` block supports:

* `prefix` - IPv6 prefix (max = /64).
* `id` - Entry ID.
* `description` - Description.

The `dhcp_server_access_list` block supports:

* `server_ip` - IP address for DHCP Server.
* `name` - User given name for dhcp-server.
* `server_ip6` - IP address for DHCPv6 Server.

The `mrouter_ports` block supports:

* `member_name` - Interface name.

The `mld_snooping_static_group` block supports:

* `ignore_reports` - Enable/disable to ignore all MLD membership reports received for this group. Valid values: `enable`, `disable`.
* `mcast_addr` - IPv6 Multicast address for static-group.
* `name` - Group name.
* `members` - Member interfaces. The structure of `members` block is documented below.

The `members` block supports:

* `member_name` - Interface name.

The `igmp_snooping_static_group` block supports:

* `ignore_reports` - Enable/disable to ignore all IGMP membership reports received for this group. Valid values: `enable`, `disable`.
* `mcast_addr` - Multicast address for static-group.
* `name` - Group name.
* `members` - Member interfaces. The structure of `members` block is documented below.

The `members` block supports:

* `member_name` - Interface name.

The `member_by_ipv4` block supports:

* `description` - Description.
* `id` - Entry ID.
* `address` - Address(/32) or subnet.

The `member_by_proto` block supports:

* `protocol` - Ethernet protocols (0 - 65535).
* `frametypes` - Ethernet frame types to check. Valid values: `ethernet2`, `802.3d`, `llc`.
* `id` - Entry ID.
* `description` - Description.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

Switch Vlan can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_vlan.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_vlan.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
