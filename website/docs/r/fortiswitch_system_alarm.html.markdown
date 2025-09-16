---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_alarm"
description: |-
  Alarm configuration.
---

# fortiswitch_system_alarm
Alarm configuration. Applies to FortiSwitch Version `<= 7.6.2`.

## Example Usage

```hcl
resource "fortiswitch_system_alarm" "router" {
   audible = "enable"
   status = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable alarm. Valid values: `enable`, `disable`.
* `audible` - Enable/disable audible alarm. Valid values: `enable`, `disable`.
* `groups` - Alarm groups. The structure of `groups` block is documented below.
* `sequence` - Sequence id of alarms.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `groups` block supports:

* `user_auth_failure_threshold` - User authentication failure threshold.
* `encryption_failure_threshold` - Encryption failure threshold.
* `admin_auth_lockout_threshold` - Admin authentication lockout threshold.
* `fw_policy_id_threshold` - Firewall policy id threshold.
* `self_test_failure_threshold` - Self-test failure threshold.
* `fw_policy_violations` - Firewall policy violations. The structure of `fw_policy_violations` block is documented below.
* `fw_policy_id` - Firewall policy id.
* `period` - Time period in seconds (0=from start up).
* `replay_attempt_threshold` - Replay attempt threshold.
* `admin_auth_failure_threshold` - Admin authentication failure threshold.
* `decryption_failure_threshold` - Decryption failure threshold.
* `user_auth_lockout_threshold` - User authentication lockout threshold.
* `log_full_warning_threshold` - Log full warning threshold.
* `id` - Group id.

The `fw_policy_violations` block supports:

* `threshold` - Firewall policy violation threshold.
* `dst_port` - Destination port (0=all).
* `src_port` - Source port (0=all).
* `src_ip` - Source ip (0=all).
* `dst_ip` - Destination ip (0=all).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Alarm can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_alarm.labelname SystemAlarm

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_alarm.labelname SystemAlarm
$ unset "FORTISWITCH_IMPORT_TABLE"
```
