---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchacl_8021X"
description: |-
  802-1X Radius Dynamic Ingress Policy configuration.
---

# fortiswitch_switchacl_8021X
802-1X Radius Dynamic Ingress Policy configuration. Applies to FortiSwitch Version `>= 7.0.2`.

## Example Usage

```hcl
resource "fortiswitch_switchacl_8021X" "name" {
        access_list_entry {
            action {
                count = "enable"
                drop = "enable"
            }
            classifier {
                dst_ip_prefix = "1.1.1.1/24"
                dst_mac = "00:00:5e:00:53:af"
                ether_type = "11"
                service = "ALL"
                src_ip_prefix = "2.2.2/24"
                src_mac = "00:00:5e:00:53:af"
                vlan_id = "14"
            }
            description = "test"
            group = "3"
            id =  "17"
        }
        description = "test"
        filter_id = "2"
        fswid =  "20"
}
```

## Argument Reference

The following arguments are supported:

* `access_list_entry` - Access Control List Entry configuration. The structure of `access_list_entry` block is documented below.
* `filter_id` - filter-id of the policy.
* `fswid` - 802-1X Dynamic Ingress policy ID.
* `description` - Description of the policy.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `access_list_entry` block supports:

* `action` - Actions for the policy. The structure of `action` block is documented below.
* `group` - Group ID of the policy.
* `id` - Ingress policy ID.
* `classifier` - Match-conditions for the policy. The structure of `classifier` block is documented below.
* `description` - Description of the policy.

The `action` block supports:

* `count` - Count enable/disable action. Valid values: `enable`, `disable`.
* `drop` - Drop enable/disable action. Valid values: `enable`, `disable`.

The `classifier` block supports:

* `dst_mac` - Destination mac address to be matched.
* `service` - Service name.
* `dst_ip_prefix` - Destination-ip address to be matched.
* `src_mac` - Source mac address to be matched.
* `ether_type` - Ether type to be matched.
* `vlan_id` - Vlan id to be matched.
* `src_ip_prefix` - Source-ip address to be matched.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

SwitchAcl 8021X can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchacl_8021X.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchacl_8021X.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
