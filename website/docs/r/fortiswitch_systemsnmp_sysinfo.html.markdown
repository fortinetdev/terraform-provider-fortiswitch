---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemsnmp_sysinfo"
description: |-
  SNMP system info configuration.
---

# fortiswitch_systemsnmp_sysinfo
SNMP system info configuration.

## Example Usage

```hcl
resource "fortiswitch_systemsnmp_sysinfo" "name" {
        status = "enable"
        trap_high_cpu_threshold = "8"
        trap_log_full_threshold = "9"
        trap_low_memory_threshold = "10"
        trap_temp_alarm_threshold = "11"
        trap_temp_warning_threshold = "10"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable snmp. Valid values: `enable`, `disable`.
* `trap_temp_alarm_threshold` - Alarm threshold degree in C.
* `description` - System description.
* `contact_info` - Contact information.
* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_log_full_threshold` - Log disk usage when trap is sent.
* `trap_low_memory_threshold` - Memory usage when trap is sent.
* `engine_id` - Local snmp engineID string (max 24 char).
* `location` - System location.
* `trap_temp_warning_threshold` - Warning threshold degree in C.
* `trap_high_cpu_interval` - Time period over which the CPU usage is calculated. Valid values: `1min`, `10min`, `30min`, `1hr`, `12hr`, `24hr`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SystemSnmp Sysinfo can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemsnmp_sysinfo.labelname SystemSnmpSysinfo

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemsnmp_sysinfo.labelname SystemSnmpSysinfo
$ unset "FORTISWITCH_IMPORT_TABLE"
```
