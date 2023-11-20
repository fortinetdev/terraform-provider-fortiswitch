---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchlldp_profile"
description: |-
  LLDP configuration profiles.
---

# fortiswitch_switchlldp_profile
LLDP configuration profiles.

## Example Usage

```hcl
resource "fortiswitch_switchlldp_profile" "name" {
        auto_isl = "enable"
        auto_isl_hello_timer = "6"
        auto_isl_port_group = "7"
        auto_isl_receive_timeout = "8"
        auto_mclag_icl = "enable"
        med_network_policy {
            assign_vlan = "disable"
            dscp = "21"
            name = "default_name_22"
            priority = "5"
            status = "disable"
            vlan = "25"
        }
        med_tlvs = "inventory_management"
        name = "default_name_27"
}
```

## Argument Reference

The following arguments are supported:

* `auto_isl_auth_encrypt` - Automatic inter-switch LAG authentication encryption.
* `vlan_name_map` - VLANs that advertise Vlan Names
* `auto_mclag_icl` - Enable/disable MCLAG inter chassis link. Valid values: `enable`, `disable`.
* `name` - Profile name.
* `auto_isl` - Enable/disable automatic inter-switch LAG. Valid values: `enable`, `disable`.
* `n8023_tlvs` - Transmitted IEEE 802.3 TLVs. Valid values: `max-frame-size`, `power-negotiation`, `eee-config`.
* `auto_isl_auth` - Automatic inter-switch LAG authentication. Valid values: `legacy`, `strict`, `relax`.
* `auto_isl_receive_timeout` - Automatic ISL timeout (0 - 300 sec).
* `auto_isl_auth_macsec_profile` - Fortilink LLDP macsec port profile.
* `auto_isl_port_group` - Automatic inter-switch LAG port group.
* `custom_tlvs` - Organizationally Specific TLV configuration. The structure of `custom_tlvs` block is documented below.
* `auto_isl_auth_user` - Automatic authentication User certificate.
* `auto_isl_auth_reauth` - Automatic authentication reauth period (10 - 3600 mins).
* `med_location_service` - LLDP MED location service configuration. The structure of `med_location_service` block is documented below.
* `n8021_tlvs` - Transmitted IEEE 802.1 TLVs.
* `med_tlvs` - Transmitted LLDP-MED TLVs. Valid values: `inventory-management`, `network-policy`, `location-identification`, `power-management`.
* `med_network_policy` - LLDP MED network policy configuration. The structure of `med_network_policy` block is documented below.
* `auto_isl_auth_identity` - Automatic authentication identity.
* `auto_isl_hello_timer` - Automatic ISL hello timer (1 - 30 sec).
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `custom_tlvs` block supports:

* `subtype` - Organizationally defined subtype.
* `oui` - Organizationally unique identifier.
* `name` - TLV name (not sent).
* `information_string` - Organizationally defined information string.

The `med_location_service` block supports:

* `status` - Enable/disable Location Service TLV. Valid values: `disable`, `enable`.
* `sys_location_id` - LLDP System Location Id.
* `name` - Policy type name.

The `med_network_policy` block supports:

* `status` - Enable/disable this TLV. Valid values: `disable`, `enable`.
* `name` - Policy type name.
* `vlan` - VLAN to advertise (if configured on port).
* `dscp` - Advertised DSCP value.
* `priority` - Advertised L2 priority.
* `assign_vlan` - Enable/disable automatically adding this VLAN to ports with this profile (does not affect trunks). Valid values: `disable`, `enable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchLldp Profile can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchlldp_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchlldp_profile.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
