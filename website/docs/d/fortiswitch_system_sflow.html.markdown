---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_sflow"
description: |-
  Get information on fortiswitch system sflow.
---

# Data Source: fortiswitch_system_sflow
Use this data source to get information on fortiswitch system sflow

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `dummy` - SFlow dummy.
* `collectors` - Collectors. The structure of `collectors` block is documented below.

The `collectors` block contains:

* `ip` - Collector IP.
* `name` - Collector name.
* `port` - SFlow collector port (0 - 65535).

