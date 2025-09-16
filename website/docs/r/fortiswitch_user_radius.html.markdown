---
subcategory: "FortiSwitch User"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_user_radius"
description: |-
  RADIUS server entry configuration.
---

# fortiswitch_user_radius
RADIUS server entry configuration.

## Example Usage

```hcl
resource "fortiswitch_user_radius" "uradius" {
          acct_server {
              id = "6"
              port = "7"
              secret = "123"
              server = "192.168.100.40"
              status = "enable"
          }
          addr_mode = "ipv4"
          all_usergroup = "disable"
          auth_type = "auto"
          frame_mtu_size = "600"
          link_monitor = "disable"
          link_monitor_interval = "16"
          name = "radius"
          nas_ip = "1.2.3.4"
          server = "192.168.100.40"
          service_type = "login"
}
```

## Argument Reference

The following arguments are supported:

* `radius_coa` - Enable/disable RADIUS CoA services from this server. Valid values: `disable`, `enable`.
* `link_monitor_interval` - Time in seconds ( default 600 ) for the link-monitor interval
* `secondary_secret` - Secret key to access the secondary server.
* `source_ip6` - Source IPv6 for communications to RADIUS server.
* `nas_ip6` - NAS IPv6 for the RADIUS request.
* `acct_interim_interval` - Time in seconds ( default 600 ) between each accounting interim update message.
* `source_ip` - Source IPv4 for communications to RADIUS server.
* `auth_type` - Authentication protocol. Valid values: `auto`, `ms_chap_v2`, `ms_chap`, `chap`, `pap`.
* `all_usergroup` - Enable/disable automatic inclusion of this RADIUS server to all user groups. Valid values: `disable`, `enable`.
* `secret` - Secret key to access the primary server.
* `nas_ip` - NAS IPv4 for the RADIUS request.
* `link_monitor` - Enable/disable RADIUS link-monitor service from this server. Valid values: `disable`, `enable`.
* `acct_fast_framedip_detect` - Time in seconds ( default 4 ) for Accounting message Framed-IP detection from DHCP Snooping.
* `frame_mtu_size` - Frame MTU Size.
* `secondary_server` - Secondary RADIUS domain name or IP address.
* `service_type` - Radius Service Type. Valid values: `login`, `framed`, `callback-login`, `callback-framed`, `outbound`, `administrative`, `nas-prompt`, `authenticate-only`, `callback-nas-prompt`, `call-check`, `callback-administrative`.
* `name` - RADIUS server entry name.
* `server` - Primary server domain name or IP address.
* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
* `radius_coa_secret` - Secret key to access the local Radius CoA server.
* `transport_type` - Enable/disable Radsec ( Radius Security ) services with different protocol from this server. Valid values: `UDP`, `TLS`, `DTLS`.
* `radsec_port` - Server Radsec service port number.
* `radsec_oper_mode` - Radsec ( Radius Security ) TLS operation mode. Valid values: `TLS-X.509`.
* `radsec_tls_min_ver` - Radsec ( Radius Security ) TLS min version. Valid values: `TLSv1`, `TLSv1-1`, `TLSv1-2`, `TLSv1-3`.
* `radsec_dtls_min_ver` - Radsec ( Radius Security ) DTLS min version. Valid values: `DTLSv1`, `DTLSv1-2`.
* `radsec_server_ca_cert` - CA certificate for Radsec Server.
* `radsec_client_cert` - Client certificate for Radsec Client.
* `radsec_cert_validate` - Enable/disable RADSEC TLS peer certificate check. Valid values: `disable`, `enable`.
* `radsec_cert_cn_dns` - Radsec certificate CN or DNS string.
* `radsec_idle_timeout` - Server Radsec Idle timout 60 to 3600 seconds.
* `radsec_connect_timeout` - Server Radsec connection timeout 1 to 5 seconds.
* `acct_server` - Additional accounting servers. The structure of `acct_server` block is documented below.
* `radius_port` - Local RADIUS service port number.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].

The `acct_server` block supports:

* `status` -  Enable/disable Status. Valid values: `enable`, `disable`.
* `secret` - Secret key.
* `port` -  RADIUS accounting port number.
* `id` - ID (0 - 4294967295).
* `server` - Server IP address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Radius can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_user_radius.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_user_radius.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
