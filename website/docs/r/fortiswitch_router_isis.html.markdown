---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_isis"
description: |-
  ISIS configuration.
---

# fortiswitch_router_isis
ISIS configuration.

## Example Usage

```hcl
resource "fortiswitch_router_isis" "name" {
        auth_mode_area = "password"
        auth_mode_domain = "password"
        auth_password_area = "china"
        auth_password_domain = "admin"
        auth_sendonly_area = "enable"
        auth_sendonly_domain = "enable"
        default_information_level = "level_1_2"
        default_information_level6 = "level_1_2"
        default_information_metric = "13"
        default_information_metric6 = "14"
        default_information_originate = "enable"
        default_information_originate6 = "enable"
        ignore_attached_bit = "enable"
        is_type = "level_1_2"
        log_neighbour_changes = "enable"
        lsp_gen_interval_l1 = "44"
        lsp_gen_interval_l2 = "45"
        lsp_refresh_interval = "46"
        max_lsp_lifetime = "500"
        metric_style = "narrow"
}
```

## Argument Reference

The following arguments are supported:

* `default_information_metric6` - Default ipv6 route metric.
* `auth_sendonly_domain` - Enable authentication send-only for level 2 SNP PDUs. Valid values: `enable`, `disable`.
* `default_information_originate6` - Enable/disable generation of default ipv6 route. Valid values: `enable`, `always`, `disable`.
* `summary_address` - IS-IS summary addresses. The structure of `summary_address` block is documented below.
* `metric_style` - Use old-style (ISO 10589) or new-style packet formats. Valid values: `narrow`, `wide`, `transition`.
* `redistribute6_l1` - Redistribute level 1 v6 routes into level 2. Valid values: `enable`, `disable`.
* `lsp_refresh_interval` - LSP refresh time in seconds.
* `max_lsp_lifetime` - Maximum LSP lifetime in seconds.
* `ignore_attached_bit` - Ignore Attached bit on incoming L1 LSP. Valid values: `enable`, `disable`.
* `redistribute6_l1_list` - Access-list for redistribute v6 routes from l1 to l2.
* `auth_keychain_area` - IS-IS area authentication key-chain. Applicable when area's auth mode is md5.
* `auth_mode_area` - IS-IS area(level-1) authentication mode. Valid values: `password`, `md5`.
* `default_information_level` - Distribute default route into level's LSP. Valid values: `level-1-2`, `level-1`, `level-2`.
* `redistribute_l1` - Redistribute level 1 routes into level 2. Valid values: `enable`, `disable`.
* `net` - IS-IS net configuration. The structure of `net` block is documented below.
* `summary_address6` - IS-IS summary ipv6 addresses. The structure of `summary_address6` block is documented below.
* `auth_mode_domain` - ISIS domain(level-2) authentication mode. Valid values: `password`, `md5`.
* `lsp_gen_interval_l1` - Minimum interval for level 1 LSP regenerating.
* `lsp_gen_interval_l2` - Minimum interval for level 2 LSP regenerating.
* `redistribute_l1_list` - Access-list for redistribute l1 to l2.
* `overload_bit` - Signal other routers not to use us in SPF. Valid values: `enable`, `disable`.
* `auth_password_area` - IS-IS area(level-1) authentication password. Applicable when area's auth mode is password.
* `default_information_metric` - Default information metric.
* `default_information_originate` - Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
* `interface` - IS-IS interface configuration. The structure of `interface` block is documented below.
* `auth_keychain_domain` - IS-IS domain authentication key-chain. Applicable when domain's auth mode is md5.
* `spf_interval_exp_l2` - Level 2 SPF minimum calculation delay in secs.
* `spf_interval_exp_l1` - Level 1 SPF minimum calculation delay in secs.
* `auth_password_domain` - IS-IS domain(level-2) authentication password. Applicable when domain's auth mode is password.
* `redistribute` - IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
* `router_id` - Router ID.
* `auth_sendonly_area` - Enable authentication send-only for level 1 SNP PDUs. Valid values: `enable`, `disable`.
* `is_type` - IS-type. Valid values: `level-1-2`, `level-1`, `level-2-only`.
* `default_information_level6` - Distribute ipv6 default route into level's LSP. Valid values: `level-1-2`, `level-1`, `level-2`.
* `redistribute6` - IS-IS redistribute v6 protocols. The structure of `redistribute6` block is documented below.
* `log_neighbour_changes` - Enable logging of ISIS neighbour's changes Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `summary_address` block supports:

* `prefix` - prefix.
* `id` - Summary address entry id.
* `level` - Level. Valid values: `level-1-2`, `level-1`, `level-2`.

The `net` block supports:

* `net` - isis net xx.xxxx. ... .xxxx.xx

The `summary_address6` block supports:

* `level` - Level. Valid values: `level-1-2`, `level-1`, `level-2`.
* `prefix6` - IPv6 prefix
* `id` - Summary address entry id.

The `interface` block supports:

* `auth_password_hello` - Hello PDU authentication password. Applicable when hello's auth mode is password.
* `priority_l2` - Level 2 priority.
* `priority_l1` - Level 1 priority.
* `hello_multiplier_l2` - Level 2 multiplier for Hello holding time.
* `hello_multiplier_l1` - Level 1 multiplier for Hello holding time.
* `auth_mode_hello` - Hello PDU authentication mode. Valid values: `md5`, `password`.
* `bfd` - Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
* `passive` - Set this interface as passive. Valid values: `enable`, `disable`.
* `circuit_type` - IS-IS interface's circuit type. Valid values: `level-1-2`, `level-1`, `level-2`.
* `bfd6` - Ipv6 Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
* `wide_metric_l1` - Level 1 wide metric for interface.
* `wide_metric_l2` - Level 2 wide metric for interface.
* `status` - Enable the interface for IS-IS. Valid values: `enable`, `disable`.
* `metric_l1` - Level 1 metric for interface.
* `metric_l2` - Level 2 metric for interface.
* `status6` - Enable/disable interface for ipv6 IS-IS. Valid values: `enable`, `disable`.
* `name` - IS-IS interface name
* `auth_keychain_hello` - Hello PDU authentication key-chain. Applicable when hello's auth mode is md5.
* `hello_interval_l2` - Level 2 hello interval.
* `csnp_interval_l2` - Level 2 CSNP interval.
* `csnp_interval_l1` - Level 1 CSNP interval.
* `hello_padding` - Enable padding to IS-IS hello packets. Valid values: `enable`, `disable`.
* `hello_interval_l1` - Level 1 hello interval.

The `redistribute` block supports:

* `status` - status. Valid values: `enable`, `disable`.
* `protocol` - protocol name.
* `level` - level. Valid values: `level-1-2`, `level-1`, `level-2`.
* `metric` - metric.
* `routemap` - routemap name.
* `metric_type` - metric type. Valid values: `external`, `internal`.

The `redistribute6` block supports:

* `status` - status. Valid values: `enable`, `disable`.
* `metric` - metric.
* `routemap` - routemap name.
* `protocol` - protocol name.
* `level` - level. Valid values: `level-1-2`, `level-1`, `level-2`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Isis can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_isis.labelname RouterIsis

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_router_isis.labelname RouterIsis
$ unset "FORTISWITCH_IMPORT_TABLE"
```
