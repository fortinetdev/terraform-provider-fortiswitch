---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemsnmp_sysinfo"
description: |-
  Get information on fortiswitch systemsnmp sysinfo.
---

# Data Source: fortiswitch_systemsnmp_sysinfo
Use this data source to get information on fortiswitch systemsnmp sysinfo

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable snmp.
* `trap_temp_alarm_threshold` - Alarm threshold degree in C.
* `description` - System description.
* `contact_info` - Contact information.
* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_log_full_threshold` - Log disk usage when trap is sent.
* `trap_low_memory_threshold` - Memory usage when trap is sent.
* `engine_id` - Local snmp engineID string (max 24 char).
* `location` - System location.
* `trap_temp_warning_threshold` - Warning threshold degree in C.
* `trap_high_cpu_interval` - Time period over which the CPU usage is calculated.

