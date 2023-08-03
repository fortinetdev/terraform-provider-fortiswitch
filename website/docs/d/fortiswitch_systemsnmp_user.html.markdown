---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemsnmp_user"
description: |-
  Get information on an fortiswitch systemsnmp user.
---

# Data Source: fortiswitch_systemsnmp_user
Use this data source to get information on an fortiswitch systemsnmp user

## Argument Reference

* `name` - (Required) Specify the name of the desired systemsnmp user.

## Attribute Reference

The following attributes are exported:

* `notify_hosts` - Send notifications (traps) to these hosts.
* `priv_proto` - Privacy (encryption) protocol.
* `name` - SNMP user name.
* `query_port` - SNMPv3 query port.
* `auth_pwd` - Password for authentication protocol.
* `priv_pwd` - Password for privacy (encryption) protocol.
* `security_level` - Security level for message authentication and encryption.
* `auth_proto` - Authentication protocol.
* `queries` - Enable/disable queries for this user.
* `events` - SNMP notifications (traps) to send.

