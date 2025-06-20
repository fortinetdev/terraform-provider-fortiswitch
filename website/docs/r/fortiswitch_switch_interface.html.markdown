---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_interface"
description: |-
  Usable interfaces (trunks and ports).
---

# fortiswitch_switch_interface
Usable interfaces (trunks and ports).

## Example Usage

```hcl
resource "fortiswitch_switch_interface" "name" {
        name = "port2"
        native_vlan = "30"
        packet_sample_rate = "31"
        packet_sampler = "enabled"
}
```

## Argument Reference

The following arguments are supported:

* `allow_arp_monitor` - Enable/Disable ARP monitoring. Valid values: `disable`, `enable`.
* `fortilink_l3_mode` - FortiLink L3 uplink port. Valid values: `enable`, `disable`.
* `igmp_snooping_flood_reports` - Enable/disable flooding of IGMP snooping reports to this interface. Valid values: `enable`, `disable`.
* `interface_mode` - Set interface mode - L2 or L3. Valid values: `L2`, `L3`.
* `auto_discovery_fortilink` - Enable/disable automatic FortiLink discovery mode. Valid values: `disable`, `enable`.
* `trust_dot1p_map` - QOS trust 802.1p map.
* `discard_mode` - Configure discard mode for interface. Valid values: `none`, `all-tagged`, `all-untagged`.
* `qos_policy` - QOS egress COS queue policy.
* `arp_inspection_trust` - Dynamic ARP Inspection (trusted or untrusted). Valid values: `trusted`, `untrusted`.
* `security_groups` - Group name. The structure of `security_groups` block is documented below.
* `private_vlan` - Configure private VLAN. Valid values: `disable`, `promiscuous`, `sub-vlan`.
* `sub_vlan` - Private VLAN sub-VLAN to host.
* `allowed_sub_vlans` - Sub-VLANs allowed to egress this interface.
* `sample_direction` - SFlow sample direction. Valid values: `tx`, `rx`, `both`.
* `dhcp_snoop_option82_trust` - Enable/Disable (allow/disallow) dhcp pkt with option82 on untrusted interface. Valid values: `enable`, `disable`.
* `edge_port` - Enable/disable interface as edge port. Valid values: `enabled`, `disabled`.
* `stp_bpdu_guard` - Enable/disable STP BPDU guard protection (stp-state and edge-port must be enabled). Valid values: `enabled`, `disabled`.
* `dhcp_snoop_option82_override` - Configure per vlan option82 override. The structure of `dhcp_snoop_option82_override` block is documented below.
* `ptp_status` - PTP Admin. Status. Valid values: `enable`, `disable`.
* `loop_guard` - Enable/disable loop guard protection. Valid values: `enabled`, `disabled`.
* `qnq` - Configure QinQ. The structure of `qnq` block is documented below.
* `allowed_vlans` - Allowed VLANs.
* `nac` - Enable/disable NAC in Fortilink mode. Valid values: `enable`, `disable`.
* `auto_discovery_fortilink_packet_interval` - FortiLink packet interval for automatic discovery (3 - 300 sec).
* `primary_vlan` - Private VLAN to host.
* `vlan_mapping` - Configure vlan mapping table. The structure of `vlan_mapping` block is documented below.
* `raguard` - IPV6 RA guard configuration. The structure of `raguard` block is documented below.
* `stp_bpdu_guard_timeout` - BPDU Guard disabling protection (min).
* `stp_loop_protection` - Enable/disable spanning tree protocol loop guard protection (stp-state must be enabled). Valid values: `enabled`, `disabled`.
* `learning_limit` - Limit the number of dynamic MAC addresses on this port.
* `learning_limit_action` - Enable/disable turning off this interface on learn limit violation. Valid values: `none`, `shutdown`.
* `loop_guard_mac_move_threshold` - Trigger loop guard if MAC move per second of this interface reaches this threshold.
* `log_mac_event` - Enable/disable logging for dynamic MAC address events. Valid values: `enable`, `disable`.
* `dhcp_snoop_learning_limit` - Maximum DHCP IP learned on the interface.
* `type` - Interface type. Valid values: `physical`, `trunk`.
* `snmp_index` - SNMP index.
* `dhcp_snoop_learning_limit_check` - Enable/Disable DHCP learning limit check on the interface. Valid values: `disable`, `enable`.
* `switch_port_mode` - Enable/disable port as L2 switch port (enable) or as pure routed port (disable). Valid values: `disable`, `enable`.
* `description` - Description.
* `stp_state` - Enable/disable spanning tree protocol. Valid values: `enabled`, `disabled`.
* `vlan_traffic_type` - Configure traffic tagging. Valid values: `untagged`, `tagged`.
* `sticky_mac` - Enable/disable Sticky MAC for this interface. Valid values: `enable`, `disable`.
* `packet_sampler` - Enable/disable packet sampling. Valid values: `enabled`, `disabled`.
* `loop_guard_timeout` - Loop guard disabling protection (min).
* `rpvst_port` - Enable/disable interface to inter-op with pvst Valid values: `enabled`, `disabled`.
* `ip_mac_binding` - Enable/disable ip-mac-binding on this interaface. Valid values: `global`, `enable`, `disable`.
* `mld_snooping_flood_reports` - Enable/disable flooding of MLD reports to this interface. Valid values: `enable`, `disable`.
* `sflow_counter_interval` - SFlow sampler counter polling interval (0:disable - 255).
* `name` - Interface name.
* `native_vlan` - Native (untagged) VLAN.
* `vlan_mapping_miss_drop` - Enabled or disabled drop if mapping missed. Valid values: `disable`, `enable`.
* `untagged_vlans` - Configure VLANs permitted to be transmitted without VLAN tags.
* `private_vlan_port_type` - Private VLAN or sub-VLAN based port type.
* `dhcp_snooping` - DHCP snooping interface (trusted or untrusted). Valid values: `trusted`, `untrusted`.
* `filter_sub_vlans` - Private VLAN or sub-VLAN based port type. Valid values: `disable`, `enable`.
* `stp_root_guard` - Enable/disable STP root guard protection (stp-state must be enabled). Valid values: `enabled`, `disabled`.
* `trust_ip_dscp_map` - QOS trust IP-DSCP map.
* `port_security` - Configure port security. The structure of `port_security` block is documented below.
* `default_cos` - Set default COS for untagged packets.
* `mcast_snooping_flood_traffic` - Enable/disable flooding of multicast snooping traffic to this interface. Valid values: `enable`, `disable`.
* `vlan_tpid` - Configure ether-type.
* `packet_sample_rate` - Packet sample rate (0 - 99999).
* `ptp_policy` - PTP policy.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `security_groups` block supports:

* `name` - Group name.

The `dhcp_snoop_option82_override` block supports:

* `circuit_id` - Text String of Circuit Id.
* `id` - Vlan Id.
* `remote_id` - Text String of Remote Id.

The `qnq` block supports:

* `status` - Enable/Disable QinQ mode. Valid values: `disable`, `enable`.
* `vlan_mapping` - Configure Vlan Mapping. The structure of `vlan_mapping` block is documented below.
* `stp_qnq_admin` - Enable/Disable QnQ to manage STP admin status. Valid values: `disable`, `enable`.
* `native_c_vlan` - Native c vlan for untagged packets.
* `remove_inner` - Remove inner-tag upon egress. Valid values: `disable`, `enable`.
* `priority` - Follow S-Tag or C-Tag's priority. Valid values: `follow-c-tag`, `follow-s-tag`.
* `vlan_mapping_miss_drop` - Enabled or disabled drop if mapping missed. Valid values: `disable`, `enable`.
* `untagged_s_vlan` - Add s-vlan to untagged packet.
* `add_inner` - Add inner-tag for untagged packets upon ingress.
* `allowed_c_vlan` - Allowed c vlans.
* `edge_type` - Choose edge type. Valid values: `customer`.
* `s_tag_priority` - Set priority value if packets follow S-Tag's priority.

The `vlan_mapping` block supports:

* `new_s_vlan` - Set new service vlan.
* `match_c_vlan` - Matching customer(inner) vlan.
* `id` - Entry Id.
* `description` - Description of Mapping entry.

The `vlan_mapping` block supports:

* `direction` - Ingress or Egress direction. Valid values: `ingress`, `egress`.
* `description` - Description of Mapping entry.
* `match_s_vlan` - Matching service(outer) vlan.
* `new_s_vlan` - Set new service(outer) vlan.
* `action` - Vlan action if packet is matched. Valid values: `add`, `replace`, `delete`.
* `match_c_vlan` - Matching customer(inner) vlan.
* `id` - Entry Id.

The `raguard` block supports:

* `vlan_list` - Vlan list.
* `id` - ID.
* `raguard_policy` - RA Guard policy name.

The `port_security` block supports:

* `macsec_pae_mode` - Assign PAE mode to a MACSEC interface. Valid values: `none`, `supp`, `auth`.
* `auth_fail_vlan` - Enable/disable auth_fail vlan. Valid values: `disable`, `enable`.
* `macsec_profile` - macsec port profile.
* `allow_mac_move_to` - Enable/disable allow mac move mode to this port. Valid values: `disable`, `enable`.
* `auth_fail_vlanid` - Set auth_fail vlanid.
* `port_security_mode` - Security mode. Valid values: `none`, `802.1X`, `802.1X-mac-based`, `macsec`.
* `authserver_timeout_tagged` - Set authserver_timeout tagged vlan mode. Valid values: `disable`, `lldp-voice`, `static`.
* `mab_eapol_request` - Set MAB EAPOL Request.
* `dacl` - Enable/disable dynamic access control list mode. Valid values: `disable`, `enable`.
* `authserver_timeout_vlan` - Enable/disable authserver_timeout vlan. Valid values: `disable`, `enable`.
* `allow_mac_move` - Enable/disable allow mac move mode. Valid values: `disable`, `enable`.
* `guest_vlan` - Enable/disable guest vlan. Valid values: `disable`, `enable`.
* `guest_auth_delay` - Set guest auth delay.
* `client_limit` - Set MAX number of devices can accept in MAC mode.
* `auth_order` - set authentication auth order.
* `framevid_apply` - Enable/disable the capbility to apply the EAP/MAB frame vlan to the port native vlan. Valid values: `disable`, `enable`.
* `eap_auto_untagged_vlans` - Enable/disable EAP auto-untagged-vlans mode. Valid values: `disable`, `enable`.
* `authserver_timeout_tagged_lldp_voice_vlanid` - authserver_timeout tagged lldp voice vlanid.
* `authserver_timeout_tagged_vlanid` - Set authserver_timeout tagged vlanid.
* `mac_auth_bypass` - Enable/disable mac-authentication-bypass on this interaface. Valid values: `disable`, `enable`.
* `auth_priority` - set authentication auth priority. Valid values: `legacy`, `dot1x-MAB`, `MAB-dot1x`.
* `radius_timeout_overwrite` - Enable/disable radius server session timeout to overwrite local timeout. Valid values: `disable`, `enable`.
* `authserver_timeout_period` - Set authserver_timeout period.
* `open_auth` - Enable/disable open authentication on this interaface. Valid values: `disable`, `enable`.
* `authserver_timeout_vlanid` - Set authserver_timeout vlanid.
* `quarantine_vlan` - Enable/disable Quarantine VLAN detection. Valid values: `disable`, `enable`.
* `guest_vlanid` - Set guest vlanid.
* `eap_egress_tagged` - Enable/disable Egress frame tag. Valid values: `disable`, `enable`.
* `eap_passthru` - Enable/disable EAP pass-through mode. Valid values: `disable`, `enable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Switch Interface can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_interface.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_interface.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
