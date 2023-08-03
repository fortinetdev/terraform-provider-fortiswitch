---
subcategory: "FortiSwitch AlertEmail"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_alertemail_setting"
description: |-
  Alertemail setting configuration.
---

# fortiswitch_alertemail_setting
Alertemail setting configuration.

## Example Usage

```hcl
resource "fortiswitch_alertemail_setting" "aes" {
        admin_login_logs = "enable"
        alert_interval = "4"
        amc_interface_bypass_mode = "enable"
        antivirus_logs = "enable"
        configuration_changes_logs = "enable"
        critical_interval = "8"
        debug_interval = "9"
        email_interval = "10"
        emergency_interval = "11"
        error_interval = "12"
        filter_mode = "category"
        firewall_authentication_failure_logs = "enable"
        fortiguard_log_quota_warning = "enable"
        information_interval = "20"
        local_disk_usage = "23"
        log_disk_usage_warning = "enable"
        mailto1 = "abc@gmail.com"
        mailto2 = "def@gmail2.com"
        mailto3 = "qaz@gmail3.com"
        notification_interval = "28"
        severity = "emergency"
        sslvpn_authentication_errors_logs = "enable"
        violation_traffic_logs = "enable"
        warning_interval = "34"
        webfilter_logs = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `email_interval` - Interval between each email.
* `critical_interval` - Set Critical alert interval in minutes.
* `debug_interval` - Set Debug alert interval in minutes.
* `error_interval` - Set Error alert interval in minutes.
* `sslvpn_authentication_errors_logs` - Sslvpn-authentication-errors-logs. Valid values: `enable`, `disable`.
* `ipsec_errors_logs` - IPsec-errors-logs. Valid values: `enable`, `disable`.
* `antivirus_logs` - Antivirus-logs. Valid values: `enable`, `disable`.
* `warning_interval` - Set Warning alert interval in minutes.
* `firewall_authentication_failure_logs` - Firewall-authentication-failure-logs. Valid values: `enable`, `disable`.
* `severity` - The least severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `emergency_interval` - Set Emergency alert interval in minutes.
* `fds_license_expiring_warning` - FDS-license-expiring-warning. Valid values: `enable`, `disable`.
* `configuration_changes_logs` - Configuration-changes-logs. Valid values: `enable`, `disable`.
* `notification_interval` - Set Notification alert interval in minutes.
* `information_interval` - Set Information alert interval in minutes.
* `fds_license_expiring_days` - Send alertemail before these days FortiGuard license expire (1-100).
* `admin_login_logs` - Admin-login-logs. Valid values: `enable`, `disable`.
* `username` - Set email from address.
* `alert_interval` - Set Alert alert interval in minutes.
* `mailto1` - Set destination email address 1.
* `mailto3` - Set destination email address 3.
* `mailto2` - Set destination email address 2.
* `webfilter_logs` - Webfilter-logs. Valid values: `enable`, `disable`.
* `fortiguard_log_quota_warning` - Fortiguard-log-quota-warning. Valid values: `enable`, `disable`.
* `filter_mode` - Filter mode. Valid values: `category`, `threshold`.
* `ips_logs` - IPS-logs. Valid values: `enable`, `disable`.
* `ha_logs` - HA-logs. Valid values: `enable`, `disable`.
* `local_disk_usage` - Send alertemail when disk usage exceeds this threshold (1-99).
* `ppp_errors_logs` - PPP-errors-logs. Valid values: `enable`, `disable`.
* `violation_traffic_logs` - Violation-traffic-logs. Valid values: `enable`, `disable`.
* `log_disk_usage_warning` - Log-disk-usage-warning. Valid values: `enable`, `disable`.
* `fds_update_logs` - FDS-update-logs. Valid values: `enable`, `disable`.
* `amc_interface_bypass_mode` - Amc-interface-bypass-mode. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Alertemail Setting can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_alertemail_setting.labelname AlertemailSetting

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_alertemail_setting.labelname AlertemailSetting
$ unset "FORTISWITCH_IMPORT_TABLE"
```
