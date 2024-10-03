---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchmacsec_profile"
description: |-
  MACsec configuration profiles.
---

# fortiswitch_switchmacsec_profile
MACsec configuration profiles.

## Example Usage

```hcl
resource "fortiswitch_switchmacsec_profile" "name" {
        cipher_suite = "GCM_AES_128"
        confident_offset = "0"
        encrypt_traffic = "enable"
        include_macsec_sci = "enable"
        include_mka_icv_ind = "enable"
        macsec_mode = "static_cak"
        macsec_validate = "strict"
        mka_priority = "10"
        name = "default_name_17"
        replay_protect = "enable"
        replay_window = "19"
        status = "enable"
        traffic_policy {
            name = "default_name_22"
            security_policy = "must_secure"
            status = "enable"
        }
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable this Profile. Valid values: `enable`, `disable`.
* `eap_tls_ca_cert` - CA certificate for MACSEC CAK EAP-TLS.
* `replay_protect` - Enable/disable MACsec replay protection. Valid values: `enable`, `disable`.
* `eap_tls_identity` - Client identity for MACSEC CAK EAP-TLS.
* `replay_window` - MACsec replay window size.
* `name` - Profile name.
* `macsec_mode` - Set mode of the MACsec Profile.
* `include_mka_icv_ind` - Include MKA ICV indicator.
* `eap_tls_cert` - Client certificate for MACSEC CAK EAP-TLS.
* `cipher_suite` - MACsec cipher suite. Valid values: `GCM-AES-128`.
* `traffic_policy` - MACsec traffic policy configuration. The structure of `traffic_policy` block is documented below.
* `eap_tls_radius_server` - Radius Server for MACSEC CAK EAP-TLS.
* `cipher_suite` - MACsec cipher suite.
* `macsec_validate` - Choose different MACsec validate mode. Valid values: `strict`.
* `mka_priority` - MACsec MKA priority.
* `mka_sak_rekey_time` - MACsec MKA Session SAK rekey timer.
* `encrypt_traffic` - Enable/disable Encryption of MACsec traffic. Valid values: `enable`, `disable`.
* `mka_psk` - MACsec MKA pre-shared key configuration. The structure of `mka_psk` block is documented below.
* `confident_offset` - Choose different confident offset bytes. Valid values: `0`, `30`, `50`.
* `include_macsec_sci` - Include MACsec TX SCI. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `traffic_policy` block supports:

* `status` - Enable/disable this Traffic policy. Valid values: `enable`.
* `name` - Traffic policy type name.
* `security_policy` - Must/Should secure the traffic. Valid values: `must-secure`.
* `exclude_protocol` - Exclude protocols that should not be MACsec-secured. Valid values: `ipv4`, `ipv6`, `dot1q`, `qinq`, `fortilink`, `arp`, `stp`, `lldp`, `lacp`.

The `mka_psk` block supports:

* `status` - Status of this PSK. Valid values: `active`.
* `crypto_alg` - PSK crypto algorithm.
* `name` - pre-shared-key name.
* `mka_cak` - MKA CAK pre-shared key hex string.
* `mka_ckn` - MKA CKN pre-shared key hex string.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchMacsec Profile can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchmacsec_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchmacsec_profile.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
