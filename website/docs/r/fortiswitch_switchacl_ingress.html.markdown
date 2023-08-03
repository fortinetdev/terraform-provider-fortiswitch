---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchacl_ingress"
description: |-
  Ingress Policy configuration.
---

# fortiswitch_switchacl_ingress
Ingress Policy configuration.

## Example Usage

```hcl
resource "fortiswitch_switchacl_ingress" "name" {
        action {
            count = "enable"
        }
        fswid =  "98"
        ingress_interface_all = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Set policy status. Valid values: `active`, `inactive`.
* `ingress_interface_all` - Select all interface. Valid values: `enable`, `disable`.
* `description` - Description of the policy.
* `schedule` - schedule list. The structure of `schedule` block is documented below.
* `ingress_interface` - Interface list to which policy is bound on the ingress. The structure of `ingress_interface` block is documented below.
* `classifier` - Match-conditions for the policy. The structure of `classifier` block is documented below.
* `action` - Actions for the policy. The structure of `action` block is documented below.
* `group` - Group ID of the policy.
* `fswid` - Ingress policy ID.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `schedule` block supports:

* `schedule_name` - Schedule name.

The `ingress_interface` block supports:

* `member_name` - Interface name.

The `classifier` block supports:

* `dst_mac` - Destination mac address to be matched.
* `cos` - 802.1Q CoS value to be matched.
* `service` - Service name.
* `dst_ip_prefix` - Destination-ip address to be matched.
* `src_mac` - Source mac address to be matched.
* `dscp` - DSCP value to be matched.
* `dst_ip6_prefix` - Destination-ip6 address to be matched.
* `ether_type` - Ether type to be matched.
* `vlan_id` - Vlan id to be matched.
* `src_ip_prefix` - Source-ip address to be matched.
* `src_ip6_prefix` - Source-ip6 address to be matched.

The `action` block supports:

* `count` - Count enable/disable action. Valid values: `enable`, `disable`.
* `redirect` - Redirect interface name.
* `redirect_bcast_cpu` - Redirect broadcast to all ports including CPU. Valid values: `enable`, `disable`.
* `redirect_physical_port` - List of physical ports to redirect. The structure of `redirect_physical_port` block is documented below.
* `cos_queue` - COS queue number (0 - 7), or unset to disable.
* `remark_dscp` - Remark DSCP value (0 - 63), or unset to disable.
* `egress_mask` - List of physical ports to be configured in egress mask. The structure of `egress_mask` block is documented below.
* `drop` - Drop enable/disable action. Valid values: `enable`, `disable`.
* `redirect_bcast_no_cpu` - Redirect broadcast to all ports excluding CPU. Valid values: `enable`, `disable`.
* `policer` - Policer id.
* `remark_cos` - Remark CoS value (0 - 7), or unset to disable.
* `mirror` - Mirror session name.
* `count_type` - Count-type(two colors): all, green, yellow, or red. Valid values: `all`, `green`, `yellow`, `red`.
* `cpu_cos_queue` - CPU COS queue number(17 - 25). Only if packets reach to CPU.
* `outer_vlan_tag` - Outer vlan tag.

The `redirect_physical_port` block supports:

* `member_name` - Physical port name.

The `egress_mask` block supports:

* `member_name` - Physical port name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

SwitchAcl Ingress can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchacl_ingress.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchacl_ingress.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
