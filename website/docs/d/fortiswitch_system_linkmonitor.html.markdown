---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_linkmonitor"
description: |-
  Get information on an fortiswitch system linkmonitor.
---

# Data Source: fortiswitch_system_linkmonitor
Use this data source to get information on an fortiswitch system linkmonitor

## Argument Reference

* `name` - (Required) Specify the name of the desired system linkmonitor.

## Attribute Reference

The following attributes are exported:

* `update_cascade_interface` - Enable/disable update cascade interface.
* `status` - Enable/disable link monitor administrative status.
* `timeout` - Detect request timeout.
* `protocol` - Protocols used to detect the server.
* `name` - Link monitor name.
* `http_match` - Response value from detected server in http-get.
* `source_ip` - Source IP used in packet to the server.
* `interval` - Detection interval.
* `gateway_ip6` - Gateway IPv6 address used to PING the server.
* `server` - Server address(es). The structure of `server` block is documented below.
* `failtime` - Number of retry attempts before bringing server down.
* `update_static_route` - Enable/disable update static route.
* `addr_mode` - Address mode (IPv4 or IPv6).
* `http_get` - HTTP GET URL string.
* `source_ip6` - Source IPv6 address used in packet to the server.
* `srcintf` - Interface where the monitor traffic is sent.
* `security_mode` - Twamp controller security mode.
* `packet_size` - Packet size of a twamp test session,.
* `gateway_ip` - Gateway IP used to PING the server.
* `password` - Twamp controller password in authentication mode.
* `port` - Port number to poll.
* `recoverytime` - Number of retry attempts before bringing server up.

The `server` block contains:

* `address` - Server address.

