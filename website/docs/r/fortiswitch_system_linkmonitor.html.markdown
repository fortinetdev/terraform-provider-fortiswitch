---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_linkmonitor"
description: |-
  Configure Link Health Monitor.
---

# fortiswitch_system_linkmonitor
Configure Link Health Monitor.

## Example Usage

```hcl
resource "fortiswitch_system_linkmonitor" "name" {
        addr_mode = "ipv4"
        failtime = "4"
        gateway_ip = "3.4.5.6"
        interval = "9"
        name = "test4"
        protocol = "arp"
}
```

## Argument Reference

The following arguments are supported:

* `update_cascade_interface` - Enable/disable update cascade interface. Valid values: `enable`, `disable`.
* `status` - Enable/disable link monitor administrative status. Valid values: `enable`, `disable`.
* `timeout` - Detect request timeout.
* `protocol` - Protocols used to detect the server. Valid values: `arp`, `ping`, `ping6`.
* `name` - Link monitor name.
* `http_match` - Response value from detected server in http-get.
* `source_ip` - Source IP used in packet to the server.
* `interval` - Detection interval.
* `gateway_ip6` - Gateway IPv6 address used to PING the server.
* `failtime` - Number of retry attempts before bringing server down.
* `update_static_route` - Enable/disable update static route. Valid values: `enable`, `disable`.
* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
* `http_get` - HTTP GET URL string.
* `source_ip6` - Source IPv6 address used in packet to the server.
* `srcintf` - Interface where the monitor traffic is sent.
* `security_mode` - Twamp controller security mode. Valid values: `none`, `authentication`.
* `packet_size` - Packet size of a twamp test session,.
* `gateway_ip` - Gateway IP used to PING the server.
* `password` - Twamp controller password in authentication mode.
* `port` - Port number to poll.
* `recoverytime` - Number of retry attempts before bringing server up.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System LinkMonitor can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_linkmonitor.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_linkmonitor.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
