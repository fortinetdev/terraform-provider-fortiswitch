---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemsnmp_community"
description: |-
  Get information on an fortiswitch systemsnmp community.
---

# Data Source: fortiswitch_systemsnmp_community
Use this data source to get information on an fortiswitch systemsnmp community

## Argument Reference

* `fswid` - (Required) Specify the fswid of the desired systemsnmp community.

## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable this commuity.
* `trap_v2c_lport` - SNMP v2c trap local port.
* `hosts6` - Allow hosts configuration for IPv6. The structure of `hosts6` block is documented below.
* `name` - Community name.
* `query_v1_port` - SNMP v1 query port.
* `query_v2c_port` - SNMP v2c query port.
* `query_v2c_status` - Enable/disable snmp v2c query.
* `trap_v1_rport` - SNMP v1 trap remote port.
* `query_v1_status` - Enable/disable snmp v1 query.
* `trap_v2c_status` - Enable/disable snmp v2c trap.
* `hosts` - Allow hosts configuration. The structure of `hosts` block is documented below.
* `trap_v1_status` - Enable/disable snmp v1 trap.
* `events` - Trap snmp events.
* `trap_v2c_rport` - SNMP v2c trap remote port.
* `fswid` - Community id.
* `trap_v1_lport` - SNMP v1 trap local port.

The `hosts6` block contains:

* `source_ipv6` - Source ipv6 for snmp trap.
* `interface` - Allow interface name.
* `id` - Host6 entry id.
* `ipv6` - Allow host ipv6 address.

The `hosts` block contains:

* `interface` - Allow interface name.
* `ip` - Allow host ip address and netmask.
* `source_ip` - Source ip for snmp trap.
* `id` - Host entry id.

