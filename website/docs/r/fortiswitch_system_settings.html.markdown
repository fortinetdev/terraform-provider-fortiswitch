---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_settings"
description: |-
  Settings.
---

# fortiswitch_system_settings
Settings.

## Example Usage

```hcl
resource "fortiswitch_system_settings" "name" {
        allow_subnet_overlap = "enable"
        bfd_dont_enforce_src_port = "enable"
        bfd_required_min_rx = "10"
        ecmp_max_paths = "13"
        multicast_forward = "enable"
        multicast_skip_policy = "enable"
        multicast_ttl_notchange = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `utf8_spam_tagging` - Convert spam tags to UTF8 for better non-ascii character support. Valid values: `enable`, `disable`.
* `ecmp_max_paths` - Maximum number of ECMP next-hops.
* `sip_udp_port` - UDP port the SIP proxy will monitor for SIP traffic.
* `ip` - IP address and netmask.
* `per_ip_bandwidth` - Per-ip-bandwidth disable. Valid values: `disable`, `enable`.
* `vpn_stats_log` - Enable periodic VPN log statistics. Valid values: `ipsec`, `pptp`, `l2tp`, `ssl`.
* `bfd_detect_mult` - BFD detection multiplier.
* `bfd_required_min_rx` - BFD required minimal rx interval.
* `wccp_cache_engine` - Enable wccp cache engine or not. Valid values: `enable`, `disable`.
* `ip_ecmp_mode` - IP ecmp mode. Valid values: `source-ip-based`, `dst-ip-based`, `port-based`.
* `multicast_ttl_notchange` - Multicast ttl not change. Valid values: `enable`, `disable`.
* `gateway` - Default gateway ip address.
* `multicast_forward` - Multicast forwarding. Valid values: `enable`, `disable`.
* `bfd` - Enable Bidirectional Forwarding Detection (BFD) on all interfaces. Valid values: `enable`, `disable`.
* `comments` - Vd comments.
* `vpn_stats_period` - Period to send VPN log statistics (seconds).
* `opmode` - Firewall operation mode. Valid values: `nat`.
* `bfd_desired_min_tx` - BFD desired minimal tx interval.
* `sccp_port` - TCP port the SCCP proxy will monitor for SCCP traffic.
* `asymroute6` - Asymmetric route6. Valid values: `enable`, `disable`.
* `status` - Enable/disable this VDOM. Valid values: `enable`, `disable`.
* `asymroute` - Asymmetric route. Valid values: `enable`, `disable`.
* `allow_subnet_overlap` - Allow one interface subnet overlap with other interfaces. Valid values: `enable`, `disable`.
* `manageip` - IP address and netmask.
* `device` - Interface.
* `sip_helper` - Helper to add dynamic sip firewall allow rule. Valid values: `enable`, `disable`.
* `sip_tcp_port` - TCP port the SIP proxy will monitor for SIP traffic.
* `multicast_skip_policy` - Skip policy check, and allow multicast through. Valid values: `enable`, `disable`.
* `sip_nat_trace` - Add original IP if NATed. Valid values: `enable`, `disable`.
* `bfd_dont_enforce_src_port` - Verify Source Port of BFD Packets. Valid values: `enable`, `disable`.
* `strict_src_check` - Strict source verification. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Settings can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_settings.labelname SystemSettings

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_settings.labelname SystemSettings
$ unset "FORTISWITCH_IMPORT_TABLE"
```
