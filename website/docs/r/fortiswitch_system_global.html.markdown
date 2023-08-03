---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_global"
description: |-
  Configure global range attributes.
---

# fortiswitch_system_global
Configure global range attributes.

## Example Usage

```hcl
resource "fortiswitch_system_global" "name" {
        admin_concurrent = "enable"
        admin_https_pki_required = "enable"
        admin_https_ssl_versions = "tlsv1-0"
        admin_lockout_duration = "8"
        admin_lockout_threshold = "9"
        admin_port = "10"
        admin_scp = "enable"
        admin_sport = "13"
        admin_ssh_grace_time = "14"
        admin_ssh_port = "15"
        admin_ssh_v1 = "enable"
        admin_telnet_port = "17"
        admintimeout = "18"
        alert_interval = "19"
        alertd_relog = "enable"
        allow_subnet_overlap = "enable"
        arp_timeout = "22"
        auto_isl = "enable"
        cfg_revert_timeout = "25"
        cfg_save = "automatic"
        clt_cert_req = "enable"
        csr_ca_attribute = "enable"
        daily_restart = "enable"
        detect_ip_conflict = "enable"
        dh_params = "31"
        dhcp_client_location = "intfname"
        dhcp_option_format = "legacy"
        dhcp_remote_id = "mac"
        dhcp_server_access_list = "enable"
        dhcp_snoop_client_req = "forward-untrusted"
        dhcps_db_exp = "37"
        dhcps_db_per_port_learn_limit = "38"
        dst = "enable"
        failtime = "40"
        fortilink_auto_discovery = "enable"
        hostname = "myhostname"
        image_rotation = "disable"
        interval = "44"
        ip_conflict_ignore_default = "enable"
        ipv6_accept_dad = "46"
        ipv6_all_forwarding = "enable"
        kernel_crashlog = "enable"
        kernel_devicelog = "enable"
        l3_host_expiry = "enable"
        language = "browser"
        ldapconntimeout = "52"
        private_data_encryption = "disable"
        radius_coa_port = "56"
        radius_port = "57"
        remoteauthtimeout = "58"
        revision_backup_on_logout = "enable"
        revision_backup_on_upgrade = "enable"
        strong_crypto = "enable"
        switch_mgmt_mode = "local"
        tcp_mss_min = "64"
        tcp6_mss_min = "65"
        tftp = "enable"
        timezone = "01"
}
```

## Argument Reference

The following arguments are supported:

* `n8021x_certificate` - Certificate for Port Security (802.1x).
* `ipv6_all_forwarding` - Enable/disable ipv6 all forwarding. Valid values: `enable`, `disable`.
* `dhcps_db_per_port_learn_limit` - Per Interface dhcp-server entries learn limit (default=64).
* `private_data_encryption` - Enable/disable private data encryption using an AES 128-bit key. Valid values: `disable`, `enable`.
* `dhcp_circuit_id` - List the parameters to be included to inform about client identification. Valid values: `intfname`, `vlan`, `hostname`, `mode`, `description`.
* `remoteauthtimeout` - Remote authentication (RADIUS/LDAP) time-out (0 - 300).
* `alert_interval` - Interval between each syslog entry when a sensor is out-of-range with respect to its threshold (in mins).
* `pre_login_banner` - System pre-login banner message.
* `switch_mgmt_mode` - Switch mode setting. Valid values: `local`, `fortilink`.
* `strong_crypto` - Enable/disable strong cryptography for HTTPS/SSH access. Valid values: `enable`, `disable`.
* `detect_ip_conflict` - Enable/disable detection of IP address conflicts. Valid values: `enable`, `disable`.
* `asset_tag` - Asset tag.
* `ip_conflict_ignore_default` - Enable/disable IP conflict detection for default IP address. Valid values: `enable`, `disable`.
* `timezone` - Time zone. Valid values: `01`, `02`, `03`, `04`, `05`, `81`, `06`, `07`, `08`, `09`, `10`, `11`, `12`, `13`, `74`, `14`, `77`, `15`, `87`, `16`, `17`, `18`, `19`, `20`, `75`, `21`, `22`, `23`, `24`, `80`, `79`, `25`, `26`, `27`, `28`, `78`, `29`, `30`, `31`, `32`, `33`, `34`, `35`, `36`, `37`, `38`, `83`, `84`, `40`, `85`, `41`, `42`, `43`, `39`, `44`, `46`, `47`, `51`, `48`, `45`, `49`, `50`, `52`, `53`, `54`, `55`, `56`, `57`, `58`, `59`, `60`, `62`, `63`, `61`, `64`, `65`, `66`, `67`, `68`, `69`, `70`, `71`, `72`, `00`, `82`, `73`, `86`, `76`.
* `admin_sport` - Administrative access HTTPS port (1 - 65535).
* `image_rotation` - Enable/disable image upgrade partition rotation. Valid values: `disable`, `enable`.
* `admin_telnet_port` - Administrative access TELNET port (1 - 65535).
* `kernel_crashlog` - Enable/disable capture of kernel error messages to crash log. Valid values: `enable`, `disable`.
* `fortilink_auto_discovery` - Enable/disable automatic discovery of FortiLink. Valid values: `enable`, `disable`.
* `kernel_devicelog` - Enable/disable capture of kernel device messages to log. Valid values: `enable`, `disable`.
* `revision_backup_on_upgrade` - Enable/disable automatic revision backup upon upgrade of system image. Valid values: `enable`, `disable`.
* `admin_https_ssl_versions` - Allowed SSL/TLS versions for web administration. Valid values: `tlsv1-0`, `tlsv1-1`, `tlsv1-2`, `tlsv1-3`.
* `hostname` - FortiSwitch hostname.
* `revision_backup_on_logout` - Enable/disable automatic revision backup upon logout. Valid values: `enable`, `disable`.
* `tcp6_mss_min` - Minimum allowed TCP MSS value in bytes.
* `cfg_revert_timeout` - Time-out for reverting to the last saved configuration (10 - 2147483647).
* `admin_ssh_v1` - Enable/disable SSH v1 compatibility. Valid values: `enable`, `disable`.
* `allow_subnet_overlap` - Enable/disable subnet overlap. Valid values: `enable`, `disable`.
* `dh_params` - Minimum size of Diffie-Hellman prime for HTTPS/SSH (bits).
* `ldapconntimeout` - LDAP connection time-out (0 - 2147483647 milliseconds).
* `tcp_mss_min` - Minimum allowed TCP MSS value in bytes.
* `admin_concurrent` - Enable/disable concurrent login of adminstrative users. Valid values: `enable`, `disable`.
* `admintimeout` - Idle time-out for firewall administration.
* `arp_timeout` - ARP timeout value in seconds.
* `admin_lockout_duration` - Lockout duration for FortiSwitch administration (1 - 2147483647 sec).
* `dhcp_server_access_list` - Enable/Disable trusted DHCP Server list. Valid values: `enable`, `disable`.
* `admin_port` - Administrative access HTTP port (1 - 65535).
* `l3_host_expiry` - Enable/disable l3 host expiry. Valid values: `enable`, `disable`.
* `post_login_banner` - System post-login banner message.
* `failtime` - Fail-time for PING server lost.
* `admin_lockout_threshold` - Lockout threshold for FortiSwitch administration.
* `dhcps_db_exp` - Expiry time for dhcp-snoop server-db entry (300-259200 sec, default=86400sec).
* `n8021x_ca_certificate` - CA certificate for Port Security (802.1x).
* `dhcp_remote_id` - List the parameters to be included in remote-id field. Valid values: `mac`, `hostname`, `ip`.
* `dhcp_snoop_client_req` - Client DHCP packet broadcast mode. Valid values: `forward-untrusted`, `drop-untrusted`.
* `dhcp_client_location` - List the parameters to be included to inform about client location. Valid values: `intfname`, `vlan`, `hostname`, `mode`, `description`.
* `csr_ca_attribute` - Enable/disable CA attribute in CSR. Valid values: `enable`, `disable`.
* `ipv6_accept_dad` - Whether to accept ipv6 DAD (Duplicate Address Detection).
	0: Disable DAD;
	1: Enable DAD (default);
	2: Enable DAD, and disable IPv6 operation if MAC-based duplicate link-local address has been found.
* `admin_ssh_port` - Administrative access SSH port (1 - 65535).
* `admin_server_cert` - Administrative HTTPS server certificate.
* `auto_isl` - Enable/disable automatic inter-switch LAG. Valid values: `enable`, `disable`.
* `language` - GUI display language. Valid values: `browser`, `english`, `simch`, `japanese`, `korean`, `spanish`, `trach`, `french`, `portuguese`, `german`.
* `radius_coa_port` - RADIUS CoA port number.
* `dst` - Enable/disable daylight saving time. Valid values: `enable`, `disable`.
* `interval` - Dead gateway detection interval.
* `cfg_save` - Configure configuration saving mode (valid only for changes made in the CLI). Valid values: `automatic`, `manual`, `revert`.
* `restart_time` - Daily restart time <hh:mm>.
* `dhcp_option_format` - DHCP Option format string. Valid values: `legacy`, `ascii`.
* `admin_scp` - Enable/disable downloading of system configuraiton using SCP. Valid values: `enable`, `disable`.
* `alertd_relog` - Enable/disable re-logs when a sensor exceeds it's threshold. Valid values: `enable`, `disable`.
* `clt_cert_req` - Enable the requirement of client certificate for GUI login. Valid values: `enable`, `disable`.
* `daily_restart` - Enable/disable FortiSwitch daily reboot. Valid values: `enable`, `disable`.
* `tftp` - Enable/disable TFTP. Valid values: `enable`, `disable`.
* `admin_ssh_grace_time` - Administrative access login grace time (10 - 3600 sec).
* `admin_https_pki_required` - Enable/disable HTTPS login page when PKI is enabled. Valid values: `enable`, `disable`.
* `radius_port` - RADIUS server port number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Global can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_global.labelname SystemGlobal

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_global.labelname SystemGlobal
$ unset "FORTISWITCH_IMPORT_TABLE"
```
