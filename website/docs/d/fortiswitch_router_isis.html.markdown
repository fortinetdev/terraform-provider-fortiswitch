---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_isis"
description: |-
  Get information on fortiswitch router isis.
---

# Data Source: fortiswitch_router_isis
Use this data source to get information on fortiswitch router isis

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `default_information_metric6` - Default ipv6 route metric.
* `auth_sendonly_domain` - Enable authentication send-only for level 2 SNP PDUs.
* `default_information_originate6` - Enable/disable generation of default ipv6 route.
* `summary_address` - IS-IS summary addresses. The structure of `summary_address` block is documented below.
* `metric_style` - Use old-style (ISO 10589) or new-style packet formats.
* `redistribute6_l1` - Redistribute level 1 v6 routes into level 2.
* `lsp_refresh_interval` - LSP refresh time in seconds.
* `max_lsp_lifetime` - Maximum LSP lifetime in seconds.
* `ignore_attached_bit` - Ignore Attached bit on incoming L1 LSP.
* `redistribute6_l1_list` - Access-list for redistribute v6 routes from l1 to l2.
* `auth_keychain_area` - IS-IS area authentication key-chain. Applicable when area's auth mode is md5.
* `auth_mode_area` - IS-IS area(level-1) authentication mode.
* `default_information_level` - Distribute default route into level's LSP.
* `redistribute_l1` - Redistribute level 1 routes into level 2.
* `net` - IS-IS net configuration. The structure of `net` block is documented below.
* `summary_address6` - IS-IS summary ipv6 addresses. The structure of `summary_address6` block is documented below.
* `auth_mode_domain` - ISIS domain(level-2) authentication mode.
* `lsp_gen_interval_l1` - Minimum interval for level 1 LSP regenerating.
* `lsp_gen_interval_l2` - Minimum interval for level 2 LSP regenerating.
* `redistribute_l1_list` - Access-list for redistribute l1 to l2.
* `overload_bit` - Signal other routers not to use us in SPF.
* `auth_password_area` - IS-IS area(level-1) authentication password. Applicable when area's auth mode is password.
* `default_information_metric` - Default information metric.
* `default_information_originate` - Enable/disable generation of default route.
* `interface` - IS-IS interface configuration. The structure of `interface` block is documented below.
* `auth_keychain_domain` - IS-IS domain authentication key-chain. Applicable when domain's auth mode is md5.
* `spf_interval_exp_l2` - Level 2 SPF minimum calculation delay in secs.
* `spf_interval_exp_l1` - Level 1 SPF minimum calculation delay in secs.
* `auth_password_domain` - IS-IS domain(level-2) authentication password. Applicable when domain's auth mode is password.
* `redistribute` - IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
* `router_id` - Router ID.
* `auth_sendonly_area` - Enable authentication send-only for level 1 SNP PDUs.
* `is_type` - IS-type.
* `default_information_level6` - Distribute ipv6 default route into level's LSP.
* `redistribute6` - IS-IS redistribute v6 protocols. The structure of `redistribute6` block is documented below.
* `log_neighbour_changes` - Enable logging of ISIS neighbour's changes

The `summary_address` block contains:

* `prefix` - prefix.
* `id` - Summary address entry id.
* `level` - Level.

The `net` block contains:

* `net` - isis net xx.xxxx. ... .xxxx.xx

The `summary_address6` block contains:

* `level` - Level.
* `prefix6` - IPv6 prefix
* `id` - Summary address entry id.

The `interface` block contains:

* `auth_password_hello` - Hello PDU authentication password. Applicable when hello's auth mode is password.
* `priority_l2` - Level 2 priority.
* `priority_l1` - Level 1 priority.
* `hello_multiplier_l2` - Level 2 multiplier for Hello holding time.
* `hello_multiplier_l1` - Level 1 multiplier for Hello holding time.
* `auth_mode_hello` - Hello PDU authentication mode.
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `passive` - Set this interface as passive.
* `circuit_type` - IS-IS interface's circuit type.
* `bfd6` - Ipv6 Bidirectional Forwarding Detection (BFD).
* `wide_metric_l1` - Level 1 wide metric for interface.
* `wide_metric_l2` - Level 2 wide metric for interface.
* `status` - Enable the interface for IS-IS.
* `metric_l1` - Level 1 metric for interface.
* `metric_l2` - Level 2 metric for interface.
* `status6` - Enable/disable interface for ipv6 IS-IS.
* `name` - IS-IS interface name
* `auth_keychain_hello` - Hello PDU authentication key-chain. Applicable when hello's auth mode is md5.
* `hello_interval_l2` - Level 2 hello interval.
* `csnp_interval_l2` - Level 2 CSNP interval.
* `csnp_interval_l1` - Level 1 CSNP interval.
* `hello_padding` - Enable padding to IS-IS hello packets.
* `hello_interval_l1` - Level 1 hello interval.

The `redistribute` block contains:

* `status` - status.
* `protocol` - protocol name.
* `level` - level.
* `metric` - metric.
* `routemap` - routemap name.
* `metric_type` - metric type.

The `redistribute6` block contains:

* `status` - status.
* `metric` - metric.
* `routemap` - routemap name.
* `protocol` - protocol name.
* `level` - level.

