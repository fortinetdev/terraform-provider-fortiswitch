---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_global"
description: |-
  Configure global settings.
---

# fortiswitch_switch_global
Configure global settings.

## Example Usage

```hcl
resource "fortiswitch_switch_global" "name" {
        name = "test"
        poe_alarm_threshold = "35"
        poe_guard_band = "36"
        poe_power_budget = "37"
        poe_power_mode = "priority"
        poe_pre_standard_detect = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `fortilink_heartbeat_timeout` - Max fortilinkd echo replies that can be missed before fortilink is considered down.
* `mclag_split_brain_priority` - Set MCLAG split brain priority
* `log_mac_limit_violations` - Enable/disable logs for Learning Limit Violations globally. Valid values: `enable`, `disable`.
* `trunk_hash_unkunicast_src_dst` - Enable/disable trunk hash for unknown unicast src-dst. Valid values: `enable`, `disable`.
* `poe_power_mode` - Set poe power mode to priority based or first come first served. Valid values: `priority`, `first-come-first-served`.
* `ip_mac_binding` - Configure ip-mac-binding status. Valid values: `enable`, `disable`.
* `vxlan_stp_virtual_root` - Enable/disable automatically making local switch the STP root for STP instances containing configured VXLAN’s access vlan. Valid values: `enable`, `disable`.
* `bpdu_learn` - Enable/disable BPDU learn. Valid values: `enable`, `disable`.
* `vxlan_dport` - VXLAN destination UDP port.
* `mac_address_algorithm` - Method to configure the fifth byte of the MAC address 
 (12:34:56:78:XX:XX, sixth byte automatically generated from managmenet MAC, channel, and port information). Valid values: `auto`, `manual`.
* `virtual_wire_tpid` - TPID value used by virtual-wires.
* `trunk_hash_unicast_src_port` - Enable/disable source port in Unicast trunk hashing. Valid values: `enable`, `disable`.
* `auto_fortilink_discovery` - Enable/disable automatic FortiLink discovery. Valid values: `enable`, `disable`.
* `max_path_in_ecmp_group` - Set max path in one ecmp group.
* `trunk_hash_mode` - Trunk hash mode. Valid values: `default`, `enhanced`.
* `fortilink_p2p_tpid` - FortiLink point-to-point TPID.
* `flood_vtp` - Enable/disable Cisco VTP flood in the vlan. Valid values: `enable`, `disable`.
* `mac_violation_timer` - Set a global timeout for Learning Limit Violations (0 = disabled).
* `mirror_qos` - QOS value for locally mirrored traffic.
* `mclag_port_base` - MCLAG port base.
* `flapguard_retain_trigger` - Enable/disable retention of triggered state upon reboot. Valid values: `enable`, `disable`.
* `mclag_split_brain_all_ports_down` - Enable/disable MCLAG split brain all ports down Valid values: `disable`, `enable`.
* `mac_address` - Manually configured MAC address when mac-address-algorithm is set to manual.
* `mclag_stp_aware` - MCLAG STP aware. Valid values: `enable`, `disable`.
* `fortilink_p2p_native_vlan` - FortiLink point to point native VLAN.
* `fortilink_vlan_optimization` - Controls VLAN assignment on ISL ports (assigns all 4k vlans when disabled). Valid values: `enable`, `disable`.
* `vxlan_stp_virtual_mac` - Virtual STP root MAC address
* `loop_guard_tx_interval` - Loop guard packet Tx interval (sec).
* `auto_isl_port_group` - Configure global automatic inter-switch link port groups (overrides port level port groups).
* `vxlan_port` - VXLAN destination UDP port.
* `poe_alarm_threshold` - Threshold (% of total power budget) above which an alarm event is generated.
* `poe_guard_band` - Reserves power (W) in case of a spike in PoE consumption.
* `auto_stp_priority` - Automatic assignment of STP priority for tier1 and tier2 switches. Valid values: `enable`, `disable`.
* `poe_power_budget` - Set/override maximum power budget.
* `vxlan_sport` - VXLAN source UDP port (0 - 65535).
* `mclag_peer_info_timeout` - MCLAG peer info timeout.
* `mclag_igmpsnooping_aware` - MCLAG IGMP-snooping aware. Valid values: `enable`, `disable`.
* `access_vlan_mode` - Intra VLAN traffic behavior with loss of connection to the FortiGate. Valid values: `legacy`, `fail-open`, `fail-close`.
* `poe_pre_standard_detect` - set poe-pre-standard-detect Valid values: `enable`, `disable`.
* `dhcp_snooping_database_export` - Enable/disable DHCP snoop database export to file. Valid values: `enable`, `disable`.
* `name` - Name.
* `auto_isl` - Enable/Disable automatic inter switch LAG. Valid values: `enable`, `disable`.
* `dmi_global_all` - Enable/disable DMI global status. Valid values: `enable`, `disable`.
* `l2_memory_check` - Enable/disable L2 memory check, default interval is 120 seconds. Valid values: `enable`, `disable`.
* `mclag_split_brain_detect` - Enable/disable MCLAG split brain detect. Valid values: `enable`, `disable`.
* `forti_trunk_dmac` - Destination MAC address to be used for FortiTrunk heartbeat packets.
* `mac_aging_interval` - MAC address aging interval (sec; remove any MAC addresses unused since the the last check.
* `port_security` - Global parameters for port-security. The structure of `port_security` block is documented below.
* `flood_unknown_multicast` - Enable/disable unknown mcast flood in the vlan. Valid values: `enable`, `disable`.
* `l2_memory_check_interval` - User defined interval to check L2 memory(second). 

The `port_security` block supports:

* `mac_called_station_delimiter` - MAC called station delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.
* `mab_reauth` - Enable or disable MAB reauthentication settings. Valid values: `disable`, `enable`.
* `max_reauth_attempt` - 802.1X/MAB maximum reauthorization attempt.
* `mac_case` - MAC case (default = lowercase). Valid values: `uppercase`, `lowercase`.
* `mac_password_delimiter` - MAC authentication password delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.
* `link_down_auth` - If link down detected, 'set-unauth' reverts to un-authorized state. Valid values: `set-unauth`, `no-action`.
* `mab_entry_as` - Confgure MAB MAC entry as static or dynamic. Valid values: `static`, `dynamic`.
* `mac_username_delimiter` - MAC authentication username delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.
* `reauth_period` - 802.1X/MAB reauthentication period ( minute ).
* `mac_calling_station_delimiter` - MAC calling station delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.
* `quarantine_vlan` - Enable or disable Quarantine VLAN detection. Valid values: `disable`, `enable`.
* `tx_period` - 802.1X tx period ( second ).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Switch Global can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_global.labelname SwitchGlobal

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_global.labelname SwitchGlobal
$ unset "FORTISWITCH_IMPORT_TABLE"
```
