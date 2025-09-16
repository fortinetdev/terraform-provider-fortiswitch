---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchacl_egress"
description: |-
  Egress Policy configuration.
---

# fortiswitch_switchacl_egress
Egress Policy configuration.

## Example Usage

```hcl
resource "fortiswitch_switchacl_egress" "name" {
        fswid =  "33"
        interface = "internal"
        action {
            count = "enable"    # api issue
        }
        status = "active"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Set policy status. Valid values: `active`, `inactive`.
* `group` - Group ID of the policy.
* `description` - Description of the policy.
* `schedule` - schedule list. The structure of `schedule` block is documented below.
* `classifier` - Match-conditions for the policy. The structure of `classifier` block is documented below.
* `action` - Actions for the policy. The structure of `action` block is documented below.
* `interface` - Interface to which policy is bound on the egress.
* `fswid` - Egress policy ID.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `schedule` block supports:

* `schedule_name` - Schedule name.

The `classifier` block supports:

* `dst_mac` - Destination mac address to be matched.
* `cos` - 802.1Q CoS value to be matched.
* `service` - Service name.
* `dst_ip_prefix` - Destination-ip address to be matched.
* `src_mac` - Source mac address to be matched.
* `dscp` - DSCP value to be matched.
* `ether_type` - Ether type to be matched.
* `vlan_id` - Vlan id to be matched.
* `src_ip_prefix` - Source-ip address to be matched.

The `action` block supports:

* `count` - Count enable/disable action. Valid values: `enable`, `disable`.
* `redirect` - Redirect interface name.
* `remark_dscp` - Remark DSCP value (0 - 63), or unset to disable.
* `drop` - Drop enable/disable action. Valid values: `enable`, `disable`.
* `policer` - Policer id.
* `mirror` - Mirror session name.
* `count_type` - Count-type(two colors): all, green, yellow. Valid values: `all`, `green`, `yellow`.
* `outer_vlan_tag` - Outer vlan tag.
* `remark_cos` - Remark COS value (0 - 7), or unset to disable.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

SwitchAcl Egress can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchacl_egress.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchacl_egress.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
