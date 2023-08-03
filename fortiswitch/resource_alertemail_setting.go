// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Alertemail setting configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAlertemailSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlertemailSettingUpdate,
		Read:   resourceAlertemailSettingRead,
		Update: resourceAlertemailSettingUpdate,
		Delete: resourceAlertemailSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"email_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99999),
				Optional:     true,
				Computed:     true,
			},
			"critical_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"debug_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"error_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sslvpn_authentication_errors_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_errors_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antivirus_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"warning_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"firewall_authentication_failure_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"emergency_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fds_license_expiring_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"configuration_changes_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"notification_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"information_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fds_license_expiring_days": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"admin_login_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"alert_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mailto1": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"mailto3": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"mailto2": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"webfilter_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiguard_log_quota_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_disk_usage": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99),
				Optional:     true,
				Computed:     true,
			},
			"ppp_errors_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"violation_traffic_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_disk_usage_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fds_update_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"amc_interface_bypass_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAlertemailSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectAlertemailSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating AlertemailSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateAlertemailSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating AlertemailSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("AlertemailSetting")
	}

	return resourceAlertemailSettingRead(d, m)
}

func resourceAlertemailSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectAlertemailSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating AlertemailSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateAlertemailSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing AlertemailSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAlertemailSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadAlertemailSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading AlertemailSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAlertemailSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading AlertemailSetting resource from API: %v", err)
	}
	return nil
}

func flattenAlertemailSettingEmailInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingCriticalInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingDebugInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingErrorInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingSslvpnAuthenticationErrorsLogs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingIpsecErrorsLogs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingAntivirusLogs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingWarningInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingFirewallAuthenticationFailureLogs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingEmergencyInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingFdsLicenseExpiringWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingConfigurationChangesLogs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingNotificationInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingInformationInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingFdsLicenseExpiringDays(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingAdminLoginLogs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingAlertInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingMailto1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingMailto3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingMailto2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingWebfilterLogs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingFortiguardLogQuotaWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingFilterMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingIpsLogs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingHaLogs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingLocalDiskUsage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingPppErrorsLogs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingViolationTrafficLogs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingLogDiskUsageWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingFdsUpdateLogs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAlertemailSettingAmcInterfaceBypassMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectAlertemailSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("email_interval", flattenAlertemailSettingEmailInterval(o["email-interval"], d, "email_interval", sv)); err != nil {
		if !fortiAPIPatch(o["email-interval"]) {
			return fmt.Errorf("Error reading email_interval: %v", err)
		}
	}

	if err = d.Set("critical_interval", flattenAlertemailSettingCriticalInterval(o["critical-interval"], d, "critical_interval", sv)); err != nil {
		if !fortiAPIPatch(o["critical-interval"]) {
			return fmt.Errorf("Error reading critical_interval: %v", err)
		}
	}

	if err = d.Set("debug_interval", flattenAlertemailSettingDebugInterval(o["debug-interval"], d, "debug_interval", sv)); err != nil {
		if !fortiAPIPatch(o["debug-interval"]) {
			return fmt.Errorf("Error reading debug_interval: %v", err)
		}
	}

	if err = d.Set("error_interval", flattenAlertemailSettingErrorInterval(o["error-interval"], d, "error_interval", sv)); err != nil {
		if !fortiAPIPatch(o["error-interval"]) {
			return fmt.Errorf("Error reading error_interval: %v", err)
		}
	}

	if err = d.Set("sslvpn_authentication_errors_logs", flattenAlertemailSettingSslvpnAuthenticationErrorsLogs(o["sslvpn-authentication-errors-logs"], d, "sslvpn_authentication_errors_logs", sv)); err != nil {
		if !fortiAPIPatch(o["sslvpn-authentication-errors-logs"]) {
			return fmt.Errorf("Error reading sslvpn_authentication_errors_logs: %v", err)
		}
	}

	if err = d.Set("ipsec_errors_logs", flattenAlertemailSettingIpsecErrorsLogs(o["IPsec-errors-logs"], d, "ipsec_errors_logs", sv)); err != nil {
		if !fortiAPIPatch(o["IPsec-errors-logs"]) {
			return fmt.Errorf("Error reading ipsec_errors_logs: %v", err)
		}
	}

	if err = d.Set("antivirus_logs", flattenAlertemailSettingAntivirusLogs(o["antivirus-logs"], d, "antivirus_logs", sv)); err != nil {
		if !fortiAPIPatch(o["antivirus-logs"]) {
			return fmt.Errorf("Error reading antivirus_logs: %v", err)
		}
	}

	if err = d.Set("warning_interval", flattenAlertemailSettingWarningInterval(o["warning-interval"], d, "warning_interval", sv)); err != nil {
		if !fortiAPIPatch(o["warning-interval"]) {
			return fmt.Errorf("Error reading warning_interval: %v", err)
		}
	}

	if err = d.Set("firewall_authentication_failure_logs", flattenAlertemailSettingFirewallAuthenticationFailureLogs(o["firewall-authentication-failure-logs"], d, "firewall_authentication_failure_logs", sv)); err != nil {
		if !fortiAPIPatch(o["firewall-authentication-failure-logs"]) {
			return fmt.Errorf("Error reading firewall_authentication_failure_logs: %v", err)
		}
	}

	if err = d.Set("severity", flattenAlertemailSettingSeverity(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("emergency_interval", flattenAlertemailSettingEmergencyInterval(o["emergency-interval"], d, "emergency_interval", sv)); err != nil {
		if !fortiAPIPatch(o["emergency-interval"]) {
			return fmt.Errorf("Error reading emergency_interval: %v", err)
		}
	}

	if err = d.Set("fds_license_expiring_warning", flattenAlertemailSettingFdsLicenseExpiringWarning(o["FDS-license-expiring-warning"], d, "fds_license_expiring_warning", sv)); err != nil {
		if !fortiAPIPatch(o["FDS-license-expiring-warning"]) {
			return fmt.Errorf("Error reading fds_license_expiring_warning: %v", err)
		}
	}

	if err = d.Set("configuration_changes_logs", flattenAlertemailSettingConfigurationChangesLogs(o["configuration-changes-logs"], d, "configuration_changes_logs", sv)); err != nil {
		if !fortiAPIPatch(o["configuration-changes-logs"]) {
			return fmt.Errorf("Error reading configuration_changes_logs: %v", err)
		}
	}

	if err = d.Set("notification_interval", flattenAlertemailSettingNotificationInterval(o["notification-interval"], d, "notification_interval", sv)); err != nil {
		if !fortiAPIPatch(o["notification-interval"]) {
			return fmt.Errorf("Error reading notification_interval: %v", err)
		}
	}

	if err = d.Set("information_interval", flattenAlertemailSettingInformationInterval(o["information-interval"], d, "information_interval", sv)); err != nil {
		if !fortiAPIPatch(o["information-interval"]) {
			return fmt.Errorf("Error reading information_interval: %v", err)
		}
	}

	if err = d.Set("fds_license_expiring_days", flattenAlertemailSettingFdsLicenseExpiringDays(o["FDS-license-expiring-days"], d, "fds_license_expiring_days", sv)); err != nil {
		if !fortiAPIPatch(o["FDS-license-expiring-days"]) {
			return fmt.Errorf("Error reading fds_license_expiring_days: %v", err)
		}
	}

	if err = d.Set("admin_login_logs", flattenAlertemailSettingAdminLoginLogs(o["admin-login-logs"], d, "admin_login_logs", sv)); err != nil {
		if !fortiAPIPatch(o["admin-login-logs"]) {
			return fmt.Errorf("Error reading admin_login_logs: %v", err)
		}
	}

	if err = d.Set("username", flattenAlertemailSettingUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("alert_interval", flattenAlertemailSettingAlertInterval(o["alert-interval"], d, "alert_interval", sv)); err != nil {
		if !fortiAPIPatch(o["alert-interval"]) {
			return fmt.Errorf("Error reading alert_interval: %v", err)
		}
	}

	if err = d.Set("mailto1", flattenAlertemailSettingMailto1(o["mailto1"], d, "mailto1", sv)); err != nil {
		if !fortiAPIPatch(o["mailto1"]) {
			return fmt.Errorf("Error reading mailto1: %v", err)
		}
	}

	if err = d.Set("mailto3", flattenAlertemailSettingMailto3(o["mailto3"], d, "mailto3", sv)); err != nil {
		if !fortiAPIPatch(o["mailto3"]) {
			return fmt.Errorf("Error reading mailto3: %v", err)
		}
	}

	if err = d.Set("mailto2", flattenAlertemailSettingMailto2(o["mailto2"], d, "mailto2", sv)); err != nil {
		if !fortiAPIPatch(o["mailto2"]) {
			return fmt.Errorf("Error reading mailto2: %v", err)
		}
	}

	if err = d.Set("webfilter_logs", flattenAlertemailSettingWebfilterLogs(o["webfilter-logs"], d, "webfilter_logs", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-logs"]) {
			return fmt.Errorf("Error reading webfilter_logs: %v", err)
		}
	}

	if err = d.Set("fortiguard_log_quota_warning", flattenAlertemailSettingFortiguardLogQuotaWarning(o["fortiguard-log-quota-warning"], d, "fortiguard_log_quota_warning", sv)); err != nil {
		if !fortiAPIPatch(o["fortiguard-log-quota-warning"]) {
			return fmt.Errorf("Error reading fortiguard_log_quota_warning: %v", err)
		}
	}

	if err = d.Set("filter_mode", flattenAlertemailSettingFilterMode(o["filter-mode"], d, "filter_mode", sv)); err != nil {
		if !fortiAPIPatch(o["filter-mode"]) {
			return fmt.Errorf("Error reading filter_mode: %v", err)
		}
	}

	if err = d.Set("ips_logs", flattenAlertemailSettingIpsLogs(o["IPS-logs"], d, "ips_logs", sv)); err != nil {
		if !fortiAPIPatch(o["IPS-logs"]) {
			return fmt.Errorf("Error reading ips_logs: %v", err)
		}
	}

	if err = d.Set("ha_logs", flattenAlertemailSettingHaLogs(o["HA-logs"], d, "ha_logs", sv)); err != nil {
		if !fortiAPIPatch(o["HA-logs"]) {
			return fmt.Errorf("Error reading ha_logs: %v", err)
		}
	}

	if err = d.Set("local_disk_usage", flattenAlertemailSettingLocalDiskUsage(o["local-disk-usage"], d, "local_disk_usage", sv)); err != nil {
		if !fortiAPIPatch(o["local-disk-usage"]) {
			return fmt.Errorf("Error reading local_disk_usage: %v", err)
		}
	}

	if err = d.Set("ppp_errors_logs", flattenAlertemailSettingPppErrorsLogs(o["PPP-errors-logs"], d, "ppp_errors_logs", sv)); err != nil {
		if !fortiAPIPatch(o["PPP-errors-logs"]) {
			return fmt.Errorf("Error reading ppp_errors_logs: %v", err)
		}
	}

	if err = d.Set("violation_traffic_logs", flattenAlertemailSettingViolationTrafficLogs(o["violation-traffic-logs"], d, "violation_traffic_logs", sv)); err != nil {
		if !fortiAPIPatch(o["violation-traffic-logs"]) {
			return fmt.Errorf("Error reading violation_traffic_logs: %v", err)
		}
	}

	if err = d.Set("log_disk_usage_warning", flattenAlertemailSettingLogDiskUsageWarning(o["log-disk-usage-warning"], d, "log_disk_usage_warning", sv)); err != nil {
		if !fortiAPIPatch(o["log-disk-usage-warning"]) {
			return fmt.Errorf("Error reading log_disk_usage_warning: %v", err)
		}
	}

	if err = d.Set("fds_update_logs", flattenAlertemailSettingFdsUpdateLogs(o["FDS-update-logs"], d, "fds_update_logs", sv)); err != nil {
		if !fortiAPIPatch(o["FDS-update-logs"]) {
			return fmt.Errorf("Error reading fds_update_logs: %v", err)
		}
	}

	if err = d.Set("amc_interface_bypass_mode", flattenAlertemailSettingAmcInterfaceBypassMode(o["amc-interface-bypass-mode"], d, "amc_interface_bypass_mode", sv)); err != nil {
		if !fortiAPIPatch(o["amc-interface-bypass-mode"]) {
			return fmt.Errorf("Error reading amc_interface_bypass_mode: %v", err)
		}
	}

	return nil
}

func flattenAlertemailSettingFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandAlertemailSettingEmailInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingCriticalInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingDebugInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingErrorInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingSslvpnAuthenticationErrorsLogs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingIpsecErrorsLogs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingAntivirusLogs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingWarningInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingFirewallAuthenticationFailureLogs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingEmergencyInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingFdsLicenseExpiringWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingConfigurationChangesLogs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingNotificationInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingInformationInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingFdsLicenseExpiringDays(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingAdminLoginLogs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingAlertInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingMailto1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingMailto3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingMailto2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingWebfilterLogs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingFortiguardLogQuotaWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingFilterMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingIpsLogs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingHaLogs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingLocalDiskUsage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingPppErrorsLogs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingViolationTrafficLogs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingLogDiskUsageWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingFdsUpdateLogs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAlertemailSettingAmcInterfaceBypassMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectAlertemailSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("email_interval"); ok {
		if setArgNil {
			obj["email-interval"] = nil
		} else {

			t, err := expandAlertemailSettingEmailInterval(d, v, "email_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["email-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("critical_interval"); ok {
		if setArgNil {
			obj["critical-interval"] = nil
		} else {

			t, err := expandAlertemailSettingCriticalInterval(d, v, "critical_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["critical-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("debug_interval"); ok {
		if setArgNil {
			obj["debug-interval"] = nil
		} else {

			t, err := expandAlertemailSettingDebugInterval(d, v, "debug_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["debug-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("error_interval"); ok {
		if setArgNil {
			obj["error-interval"] = nil
		} else {

			t, err := expandAlertemailSettingErrorInterval(d, v, "error_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["error-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("sslvpn_authentication_errors_logs"); ok {
		if setArgNil {
			obj["sslvpn-authentication-errors-logs"] = nil
		} else {

			t, err := expandAlertemailSettingSslvpnAuthenticationErrorsLogs(d, v, "sslvpn_authentication_errors_logs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sslvpn-authentication-errors-logs"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec_errors_logs"); ok {
		if setArgNil {
			obj["IPsec-errors-logs"] = nil
		} else {

			t, err := expandAlertemailSettingIpsecErrorsLogs(d, v, "ipsec_errors_logs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["IPsec-errors-logs"] = t
			}
		}
	}

	if v, ok := d.GetOk("antivirus_logs"); ok {
		if setArgNil {
			obj["antivirus-logs"] = nil
		} else {

			t, err := expandAlertemailSettingAntivirusLogs(d, v, "antivirus_logs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antivirus-logs"] = t
			}
		}
	}

	if v, ok := d.GetOk("warning_interval"); ok {
		if setArgNil {
			obj["warning-interval"] = nil
		} else {

			t, err := expandAlertemailSettingWarningInterval(d, v, "warning_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["warning-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("firewall_authentication_failure_logs"); ok {
		if setArgNil {
			obj["firewall-authentication-failure-logs"] = nil
		} else {

			t, err := expandAlertemailSettingFirewallAuthenticationFailureLogs(d, v, "firewall_authentication_failure_logs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["firewall-authentication-failure-logs"] = t
			}
		}
	}

	if v, ok := d.GetOk("severity"); ok {
		if setArgNil {
			obj["severity"] = nil
		} else {

			t, err := expandAlertemailSettingSeverity(d, v, "severity", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["severity"] = t
			}
		}
	}

	if v, ok := d.GetOk("emergency_interval"); ok {
		if setArgNil {
			obj["emergency-interval"] = nil
		} else {

			t, err := expandAlertemailSettingEmergencyInterval(d, v, "emergency_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["emergency-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("fds_license_expiring_warning"); ok {
		if setArgNil {
			obj["FDS-license-expiring-warning"] = nil
		} else {

			t, err := expandAlertemailSettingFdsLicenseExpiringWarning(d, v, "fds_license_expiring_warning", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["FDS-license-expiring-warning"] = t
			}
		}
	}

	if v, ok := d.GetOk("configuration_changes_logs"); ok {
		if setArgNil {
			obj["configuration-changes-logs"] = nil
		} else {

			t, err := expandAlertemailSettingConfigurationChangesLogs(d, v, "configuration_changes_logs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["configuration-changes-logs"] = t
			}
		}
	}

	if v, ok := d.GetOk("notification_interval"); ok {
		if setArgNil {
			obj["notification-interval"] = nil
		} else {

			t, err := expandAlertemailSettingNotificationInterval(d, v, "notification_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["notification-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("information_interval"); ok {
		if setArgNil {
			obj["information-interval"] = nil
		} else {

			t, err := expandAlertemailSettingInformationInterval(d, v, "information_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["information-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("fds_license_expiring_days"); ok {
		if setArgNil {
			obj["FDS-license-expiring-days"] = nil
		} else {

			t, err := expandAlertemailSettingFdsLicenseExpiringDays(d, v, "fds_license_expiring_days", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["FDS-license-expiring-days"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_login_logs"); ok {
		if setArgNil {
			obj["admin-login-logs"] = nil
		} else {

			t, err := expandAlertemailSettingAdminLoginLogs(d, v, "admin_login_logs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-login-logs"] = t
			}
		}
	}

	if v, ok := d.GetOk("username"); ok {
		if setArgNil {
			obj["username"] = nil
		} else {

			t, err := expandAlertemailSettingUsername(d, v, "username", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["username"] = t
			}
		}
	}

	if v, ok := d.GetOk("alert_interval"); ok {
		if setArgNil {
			obj["alert-interval"] = nil
		} else {

			t, err := expandAlertemailSettingAlertInterval(d, v, "alert_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["alert-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("mailto1"); ok {
		if setArgNil {
			obj["mailto1"] = nil
		} else {

			t, err := expandAlertemailSettingMailto1(d, v, "mailto1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mailto1"] = t
			}
		}
	}

	if v, ok := d.GetOk("mailto3"); ok {
		if setArgNil {
			obj["mailto3"] = nil
		} else {

			t, err := expandAlertemailSettingMailto3(d, v, "mailto3", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mailto3"] = t
			}
		}
	}

	if v, ok := d.GetOk("mailto2"); ok {
		if setArgNil {
			obj["mailto2"] = nil
		} else {

			t, err := expandAlertemailSettingMailto2(d, v, "mailto2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mailto2"] = t
			}
		}
	}

	if v, ok := d.GetOk("webfilter_logs"); ok {
		if setArgNil {
			obj["webfilter-logs"] = nil
		} else {

			t, err := expandAlertemailSettingWebfilterLogs(d, v, "webfilter_logs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["webfilter-logs"] = t
			}
		}
	}

	if v, ok := d.GetOk("fortiguard_log_quota_warning"); ok {
		if setArgNil {
			obj["fortiguard-log-quota-warning"] = nil
		} else {

			t, err := expandAlertemailSettingFortiguardLogQuotaWarning(d, v, "fortiguard_log_quota_warning", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fortiguard-log-quota-warning"] = t
			}
		}
	}

	if v, ok := d.GetOk("filter_mode"); ok {
		if setArgNil {
			obj["filter-mode"] = nil
		} else {

			t, err := expandAlertemailSettingFilterMode(d, v, "filter_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["filter-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("ips_logs"); ok {
		if setArgNil {
			obj["IPS-logs"] = nil
		} else {

			t, err := expandAlertemailSettingIpsLogs(d, v, "ips_logs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["IPS-logs"] = t
			}
		}
	}

	if v, ok := d.GetOk("ha_logs"); ok {
		if setArgNil {
			obj["HA-logs"] = nil
		} else {

			t, err := expandAlertemailSettingHaLogs(d, v, "ha_logs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["HA-logs"] = t
			}
		}
	}

	if v, ok := d.GetOk("local_disk_usage"); ok {
		if setArgNil {
			obj["local-disk-usage"] = nil
		} else {

			t, err := expandAlertemailSettingLocalDiskUsage(d, v, "local_disk_usage", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["local-disk-usage"] = t
			}
		}
	}

	if v, ok := d.GetOk("ppp_errors_logs"); ok {
		if setArgNil {
			obj["PPP-errors-logs"] = nil
		} else {

			t, err := expandAlertemailSettingPppErrorsLogs(d, v, "ppp_errors_logs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["PPP-errors-logs"] = t
			}
		}
	}

	if v, ok := d.GetOk("violation_traffic_logs"); ok {
		if setArgNil {
			obj["violation-traffic-logs"] = nil
		} else {

			t, err := expandAlertemailSettingViolationTrafficLogs(d, v, "violation_traffic_logs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["violation-traffic-logs"] = t
			}
		}
	}

	if v, ok := d.GetOk("log_disk_usage_warning"); ok {
		if setArgNil {
			obj["log-disk-usage-warning"] = nil
		} else {

			t, err := expandAlertemailSettingLogDiskUsageWarning(d, v, "log_disk_usage_warning", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-disk-usage-warning"] = t
			}
		}
	}

	if v, ok := d.GetOk("fds_update_logs"); ok {
		if setArgNil {
			obj["FDS-update-logs"] = nil
		} else {

			t, err := expandAlertemailSettingFdsUpdateLogs(d, v, "fds_update_logs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["FDS-update-logs"] = t
			}
		}
	}

	if v, ok := d.GetOk("amc_interface_bypass_mode"); ok {
		if setArgNil {
			obj["amc-interface-bypass-mode"] = nil
		} else {

			t, err := expandAlertemailSettingAmcInterfaceBypassMode(d, v, "amc_interface_bypass_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["amc-interface-bypass-mode"] = t
			}
		}
	}

	return &obj, nil
}
