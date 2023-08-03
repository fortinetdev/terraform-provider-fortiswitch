---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_securityfeature"
description: |-
  Switch security feature control nobs.
---

# fortiswitch_switch_securityfeature
Switch security feature control nobs.

## Argument Reference

The following arguments are supported:

* `tcp_port_eq` - TCP packet with Source and destination TCP port equal. Valid values: `disable`, `enable`.
* `macsa_eq_macda` - Packet with source MAC equal to Destination MAC. Valid values: `disable`, `enable`.
* `tcp_flag` - DoS attack checking for TCP flags. Valid values: `disable`, `enable`.
* `sip_eq_dip` - TCP packet with Source IP equal to Destination IP. Valid values: `disable`, `enable`.
* `allow_mcast_sa` - Ethernet packet whose source-mac is multicast. Valid values: `disable`, `enable`.
* `tcp_flag_sf` - TCP packet with SYN and FIN bit enable. Valid values: `disable`, `enable`.
* `v4_first_frag` - DoS attack checking for IPv4 first fragment. Valid values: `disable`, `enable`.
* `allow_sa_mac_all_zero` - Ethernet packet whose source-mac is all zero's. Valid values: `disable`, `enable`.
* `tcp_hdr_partial` - TCP packet with partial header. Valid values: `disable`, `enable`.
* `udp_port_eq` - IP packet with source and destination UDP port equal. Valid values: `disable`, `enable`.
* `tcp_flag_fup` - TCP packet with FIN, URG, PSH bit enable and sequence number is zero. Valid values: `disable`, `enable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Switch SecurityFeature can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_securityfeature.labelname SwitchSecurityFeature

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_securityfeature.labelname SwitchSecurityFeature
$ unset "FORTISWITCH_IMPORT_TABLE"
```
