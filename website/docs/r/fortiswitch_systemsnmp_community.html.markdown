---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemsnmp_community"
description: |-
  SNMP community configuration.
---

# fortiswitch_systemsnmp_community
SNMP community configuration.

## Example Usage

```hcl
resource "fortiswitch_systemsnmp_community" "name" {
        events = "cpu_high"
        hosts {
            id =  "4"
            interface = "internal"
            ip = "20.21.21.20 255.255.0.0"
            source_ip = "20.20.20.10"
        }
        name = "default_name_19"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable this commuity. Valid values: `enable`, `disable`.
* `trap_v2c_lport` - SNMP v2c trap local port.
* `hosts6` - Allow hosts configuration for IPv6. The structure of `hosts6` block is documented below.
* `name` - Community name.
* `query_v1_port` - SNMP v1 query port.
* `query_v2c_port` - SNMP v2c query port.
* `query_v2c_status` - Enable/disable snmp v2c query. Valid values: `enable`, `disable`.
* `trap_v1_rport` - SNMP v1 trap remote port.
* `query_v1_status` - Enable/disable snmp v1 query. Valid values: `enable`, `disable`.
* `trap_v2c_status` - Enable/disable snmp v2c trap. Valid values: `enable`, `disable`.
* `hosts` - Allow hosts configuration. The structure of `hosts` block is documented below.
* `trap_v1_status` - Enable/disable snmp v1 trap. Valid values: `enable`, `disable`.
* `events` - Trap snmp events.
* `trap_v2c_rport` - SNMP v2c trap remote port.
* `fswid` - Community id.
* `trap_v1_lport` - SNMP v1 trap local port.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `hosts6` block supports:

* `source_ipv6` - Source ipv6 for snmp trap.
* `interface` - Allow interface name.
* `id` - Host6 entry id.
* `ipv6` - Allow host ipv6 address.

The `hosts` block supports:

* `interface` - Allow interface name.
* `ip` - Allow host ip address and netmask.
* `source_ip` - Source ip for snmp trap.
* `id` - Host entry id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

SystemSnmp Community can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemsnmp_community.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemsnmp_community.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
