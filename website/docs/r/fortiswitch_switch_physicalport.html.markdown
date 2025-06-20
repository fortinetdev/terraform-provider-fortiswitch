---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_physicalport"
description: |-
  Physical port specific configuration.
---

# fortiswitch_switch_physicalport
Physical port specific configuration.

## Argument Reference

The following arguments are supported:

* `macsec_pae_mode` - Assign PAE mode to a MACSEC interface. Valid values: `none`, `supp`, `auth`.
* `l2_learning` - Enable / disable dynamic MAC address learning. Valid values: `enabled`, `disabled`.
* `dmi_status` - DMI status. Valid values: `enable`, `disable`, `global`.
* `storm_control_mode` - Storm control mode. Valid values: `global`, `override`, `disabled`.
* `speed` - Configure interface speed and duplex.
* `fortilink_p2p` - FortiLink point-to-point. Valid values: `enable`, `disable`.
* `flap_rate` - Number of stage change events needed within flap-duration.
* `macsec_profile` - macsec port profile.
* `egress_drop_mode` - Enable/Disable egress drop. Valid values: `enabled`, `disabled`.
* `loopback` - Phy Port Loopback. Valid values: `local`, `remote`, `disable`.
* `priority_based_flow_control` - Enable / disable priority-based flow control. 802.3 flow control will be applied when disabled Valid values: `disable`, `enable`.
* `owning_interface` - Trunk interface.
* `energy_efficient_ethernet` - Enable / disable energy efficient ethernet. Valid values: `enable`, `disable`.
* `qsfp_low_power_mode` - Enable/Disable QSFP low power mode. Valid values: `enabled`, `disabled`.
* `poe_status` - Enable/disable PSE. Valid values: `enable`, `disable`.
* `status` - Administrative status. Valid values: `up`, `down`.
* `medium` - Configure port preference for shared ports. Valid values: `fiber-preferred`, `copper-preferred`, `fiber-forced`, `copper-forced`.
* `cdp_status` - CDP transmit and receive status (LLDP must be enabled in LLDP settings). Valid values: `disable`, `rx-only`, `tx-only`, `tx-rx`.
* `description` - Description.
* `flapguard` - Enable / disable FlapGuard. Valid values: `enabled`, `disabled`.
* `flap_duration` - Period over which flap events are calculated (seconds).
* `eee_tx_wake_time` - EEE Transmit wake time (microseconds)(0-2560).
* `storm_control` - Storm control. The structure of `storm_control` block is documented below.
* `security_mode` - Security mode. Valid values: `none`, `macsec`.
* `flapguard_state` - Timestamp of last triggered event (or 0 if none).
* `pause_resume` - Resume threshold for resuming reception on pause metering of an ingress port. Valid values: `75%`, `50%`, `25%`.
* `l2_sa_unknown` - Forward / drop unknown(SMAC) packets when dynamic MAC address learning is disabled. Valid values: `forward`, `drop`.
* `lldp_status` - LLDP transmit and receive status. Valid values: `disable`, `rx-only`, `tx-only`, `tx-rx`.
* `eee_tx_idle_time` - EEE Transmit idle time (microseconds)(0-2560).
* `name` - Port name.
* `lldp_profile` - LLDP port TLV profile.
* `flap_trig` - Flag is set if triggered on this port.
* `link_status` - Physical link status.
* `flap_timeout` - Flap guard disabling protection (min).
* `poe_port_mode` - IEEE802.3AF/IEEE802.3AT Valid values: `IEEE802_3AF`, `IEEE802_3AT`.
* `poe_max_power_mode` - poe: set max power mode Valid values: `class-based`, `30W`.
* `port_index` - Port index.
* `max_frame_size` - Maximum frame size.
* `flow_control` - Configure flow control pause frames. Valid values: `disable`, `tx`, `rx`, `both`.
* `pause_meter_rate` - Configure ingress metering rate. In kbits. 0 = disabled.
* `poe_port_priority` - Configure port priority Valid values: `low-priority`, `high-priority`, `critical-priority`.

The `storm_control` block supports:

* `broadcast` - Enable/disable broadcast storm control. Valid values: `enable`, `disable`.
* `unknown_unicast` - Enable/disable unknown unicast storm control. Valid values: `enable`, `disable`.
* `unknown_multicast` - Enable/disable unknown multicast storm control. Valid values: `enable`, `disable`.
* `rate` - Storm control traffic rate.
* `burst_size_level` - Storm control burst size level 0-4.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Switch PhysicalPort can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_physicalport.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_physicalport.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
