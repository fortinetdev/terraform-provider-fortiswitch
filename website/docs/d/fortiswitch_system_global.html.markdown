---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_global"
description: |-
  Get information on fortiswitch system global.
---

# Data Source: fortiswitch_system_global
Use this data source to get information on fortiswitch system global

## Example Usage

```hcl
data "fortiswitch_system_global" sample1 {
}

output output1 {
  value = data.fortiswitch_system_global.sample1
}
```

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `n8021x_certificate` - Certificate for Port Security (802.1x).
* `ipv6_all_forwarding` - Enable/disable ipv6 all forwarding.
* `dhcps_db_per_port_learn_limit` - Per Interface dhcp-server entries learn limit (default=64).
* `private_data_encryption` - Enable/disable private data encryption using an AES 128-bit key.
* `dhcp_circuit_id` - List the parameters to be included to inform about client identification.
* `remoteauthtimeout` - Remote authentication (RADIUS/LDAP) time-out (0 - 300).
* `alert_interval` - Interval between each syslog entry when a sensor is out-of-range with respect to its threshold (in mins).
* `pre_login_banner` - System pre-login banner message.
* `switch_mgmt_mode` - Switch mode setting.
* `strong_crypto` - Enable/disable strong cryptography for HTTPS/SSH access.
* `detect_ip_conflict` - Enable/disable detection of IP address conflicts.
* `asset_tag` - Asset tag.
* `ip_conflict_ignore_default` - Enable/disable IP conflict detection for default IP address.
* `timezone` - Time zone.
* `admin_sport` - Administrative access HTTPS port (1 - 65535).
* `image_rotation` - Enable/disable image upgrade partition rotation.
* `admin_telnet_port` - Administrative access TELNET port (1 - 65535).
* `kernel_crashlog` - Enable/disable capture of kernel error messages to crash log.
* `fortilink_auto_discovery` - Enable/disable automatic discovery of FortiLink.
* `kernel_devicelog` - Enable/disable capture of kernel device messages to log.
* `revision_backup_on_upgrade` - Enable/disable automatic revision backup upon upgrade of system image.
* `admin_https_ssl_versions` - Allowed SSL/TLS versions for web administration.
* `hostname` - FortiSwitch hostname.
* `revision_backup_on_logout` - Enable/disable automatic revision backup upon logout.
* `tcp6_mss_min` - Minimum allowed TCP MSS value in bytes.
* `cfg_revert_timeout` - Time-out for reverting to the last saved configuration (10 - 2147483647).
* `admin_ssh_v1` - Enable/disable SSH v1 compatibility.
* `allow_subnet_overlap` - Enable/disable subnet overlap.
* `dh_params` - Minimum size of Diffie-Hellman prime for HTTPS/SSH (bits).
* `ldapconntimeout` - LDAP connection time-out (0 - 2147483647 milliseconds).
* `tcp_mss_min` - Minimum allowed TCP MSS value in bytes.
* `admin_concurrent` - Enable/disable concurrent login of adminstrative users.
* `admintimeout` - Idle time-out for firewall administration.
* `arp_timeout` - ARP timeout value in seconds.
* `admin_lockout_duration` - Lockout duration for FortiSwitch administration (1 - 2147483647 sec).
* `dhcp_server_access_list` - Enable/Disable trusted DHCP Server list.
* `admin_port` - Administrative access HTTP port (1 - 65535).
* `l3_host_expiry` - Enable/disable l3 host expiry.
* `post_login_banner` - System post-login banner message.
* `failtime` - Fail-time for PING server lost.
* `admin_lockout_threshold` - Lockout threshold for FortiSwitch administration.
* `dhcps_db_exp` - Expiry time for dhcp-snoop server-db entry (300-259200 sec, default=86400sec).
* `n8021x_ca_certificate` - CA certificate for Port Security (802.1x).
* `dhcp_remote_id` - List the parameters to be included in remote-id field.
* `dhcp_snoop_client_req` - Client DHCP packet broadcast mode.
* `dhcp_client_location` - List the parameters to be included to inform about client location.
* `csr_ca_attribute` - Enable/disable CA attribute in CSR.
* `ipv6_accept_dad` - Whether to accept ipv6 DAD (Duplicate Address Detection).
	0: Disable DAD;
	1: Enable DAD (default);
	2: Enable DAD, and disable IPv6 operation if MAC-based duplicate link-local address has been found.
* `admin_ssh_port` - Administrative access SSH port (1 - 65535).
* `admin_server_cert` - Administrative HTTPS server certificate.
* `auto_isl` - Enable/disable automatic inter-switch LAG.
* `language` - GUI display language.
* `radius_coa_port` - RADIUS CoA port number.
* `dst` - Enable/disable daylight saving time.
* `interval` - Dead gateway detection interval.
* `cfg_save` - Configure configuration saving mode (valid only for changes made in the CLI).
* `restart_time` - Daily restart time <hh:mm>.
* `dhcp_option_format` - DHCP Option format string.
* `admin_scp` - Enable/disable downloading of system configuraiton using SCP.
* `alertd_relog` - Enable/disable re-logs when a sensor exceeds it's threshold.
* `clt_cert_req` - Enable the requirement of client certificate for GUI login.
* `daily_restart` - Enable/disable FortiSwitch daily reboot.
* `tftp` - Enable/disable TFTP.
* `admin_ssh_grace_time` - Administrative access login grace time (10 - 3600 sec).
* `admin_https_pki_required` - Enable/disable HTTPS login page when PKI is enabled.
* `radius_port` - RADIUS server port number.

