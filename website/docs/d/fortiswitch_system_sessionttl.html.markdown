---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_sessionttl"
description: |-
  Get information on fortiswitch system sessionttl.
---

# Data Source: fortiswitch_system_sessionttl
Use this data source to get information on fortiswitch system sessionttl

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `default` - Default timeout.
* `port` - Configure session ttl port. The structure of `port` block is documented below.

The `port` block contains:

* `end_port` - End port number.
* `timeout` - Configure timeout.
* `protocol` - Protocol (0-255).
* `id` - Table entry ID.
* `start_port` - Start port number.

