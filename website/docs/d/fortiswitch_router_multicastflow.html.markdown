---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_multicastflow"
description: |-
  Get information on an fortiswitch router multicastflow.
---

# Data Source: fortiswitch_router_multicastflow
Use this data source to get information on an fortiswitch router multicastflow

## Argument Reference

* `name` - (Required) Specify the name of the desired router multicastflow.

## Attribute Reference

The following attributes are exported:

* `flows` - Config multicast-flow entries. The structure of `flows` block is documented below.
* `name` - Name.
* `comments` - Description/comments.

The `flows` block contains:

* `group_addr` - Multicast group address.
* `source_addr` - Multicast source address.
* `id` - ID.
* `group_addr_end` - Multicast group end address.

