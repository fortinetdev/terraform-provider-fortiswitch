---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_vxlan"
description: |-
  Get information on an fortiswitch system vxlan.
---

# Data Source: fortiswitch_system_vxlan
Use this data source to get information on an fortiswitch system vxlan

## Argument Reference

* `name` - (Required) Specify the name of the desired system vxlan.

## Attribute Reference

The following attributes are exported:

* `remote_ip` - IPv4 address of the VXLAN interface. The structure of `remote_ip` block is documented below.
* `name` - VXLAN interface name.
* `multicast_ttl` - VXLAN multicast TTL (1-255, default = 0).
* `tagged_vlans` - Tagged VLANs.
* `vni` - VXLAN network ID.
* `ip_version` - IP version to use for the VXLAN interface.
* `vlanid` - VXLAN VNI mapped VLAN ID.
* `interface` - Interface binding for tunnel initiator (source IP) / termination.

The `remote_ip` block contains:

* `ip` - IPv4 address.

