// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure global range attributes.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemGlobal() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemGlobalRead,
		Schema: map[string]*schema.Schema{

			"n8021x_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_all_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcps_db_per_port_learn_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"private_data_encryption": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_circuit_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remoteauthtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"alert_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"pre_login_banner": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_mgmt_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"strong_crypto": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"detect_ip_conflict": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"asset_tag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_conflict_ignore_default": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_sport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"image_rotation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_telnet_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"kernel_crashlog": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortilink_auto_discovery": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"kernel_devicelog": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"revision_backup_on_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_https_ssl_versions": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"revision_backup_on_logout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tcp6_mss_min": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cfg_revert_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_ssh_v1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_subnet_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_password_hash": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dh_params": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ldapconntimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"reset_button": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tcp_mss_min": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_restrict_local": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_concurrent": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admintimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"arp_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_lockout_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dhcp_server_access_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"l3_host_expiry": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tcp_options": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"post_login_banner": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"failtime": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_lockout_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dhcps_db_exp": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"n8021x_ca_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_remote_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_snoop_client_req": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"arp_inspection_monitor_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dhcp_client_location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"csr_ca_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"delaycli_timeout_cleanup": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipv6_accept_dad": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_ssh_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_isl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"language": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"radius_coa_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cfg_save": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"restart_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_option_format": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_scp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alertd_relog": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"clt_cert_req": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"daily_restart": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tftp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_ssh_grace_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_https_pki_required": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"radius_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemGlobalRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemGlobal"

	o, err := c.ReadSystemGlobal(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemGlobal: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemGlobal from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemGlobal8021XCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpv6AllForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDhcpsDbPerPortLearnLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalPrivateDataEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDhcpCircuitId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRemoteauthtimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAlertInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalPreLoginBanner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSwitchMgmtMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalStrongCrypto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDetectIpConflict(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAssetTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpConflictIgnoreDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTimezone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminSport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalImageRotation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminTelnetPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalKernelCrashlog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFortilinkAutoDiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalKernelDevicelog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRevisionBackupOnUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminHttpsSslVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRevisionBackupOnLogout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTcp6MssMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCfgRevertTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminSshV1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAllowSubnetOverlap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminPasswordHash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDhParams(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalLdapconntimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalResetButton(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTcpMssMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminRestrictLocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminConcurrent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdmintimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalArpTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminLockoutDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDhcpServerAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalL3HostExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTcpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalPostLoginBanner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFailtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminLockoutThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDhcpsDbExp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobal8021XCaCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDhcpRemoteId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDhcpSnoopClientReq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalArpInspectionMonitorTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDhcpClientLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCsrCaAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDelaycliTimeoutCleanup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpv6AcceptDad(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminSshPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAutoIsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalLanguage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRadiusCoaPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCfgSave(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRestartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDhcpOptionFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminScp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAlertdRelog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCltCertReq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDailyRestart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTftp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminSshGraceTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminHttpsPkiRequired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRadiusPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("n8021x_certificate", dataSourceFlattenSystemGlobal8021XCertificate(o["802.1x-certificate"], d, "n8021x_certificate")); err != nil {
		if !fortiAPIPatch(o["802.1x-certificate"]) {
			return fmt.Errorf("Error reading n8021x_certificate: %v", err)
		}
	}

	if err = d.Set("ipv6_all_forwarding", dataSourceFlattenSystemGlobalIpv6AllForwarding(o["ipv6-all-forwarding"], d, "ipv6_all_forwarding")); err != nil {
		if !fortiAPIPatch(o["ipv6-all-forwarding"]) {
			return fmt.Errorf("Error reading ipv6_all_forwarding: %v", err)
		}
	}

	if err = d.Set("dhcps_db_per_port_learn_limit", dataSourceFlattenSystemGlobalDhcpsDbPerPortLearnLimit(o["dhcps-db-per-port-learn-limit"], d, "dhcps_db_per_port_learn_limit")); err != nil {
		if !fortiAPIPatch(o["dhcps-db-per-port-learn-limit"]) {
			return fmt.Errorf("Error reading dhcps_db_per_port_learn_limit: %v", err)
		}
	}

	if err = d.Set("private_data_encryption", dataSourceFlattenSystemGlobalPrivateDataEncryption(o["private-data-encryption"], d, "private_data_encryption")); err != nil {
		if !fortiAPIPatch(o["private-data-encryption"]) {
			return fmt.Errorf("Error reading private_data_encryption: %v", err)
		}
	}

	if err = d.Set("dhcp_circuit_id", dataSourceFlattenSystemGlobalDhcpCircuitId(o["dhcp-circuit-id"], d, "dhcp_circuit_id")); err != nil {
		if !fortiAPIPatch(o["dhcp-circuit-id"]) {
			return fmt.Errorf("Error reading dhcp_circuit_id: %v", err)
		}
	}

	if err = d.Set("remoteauthtimeout", dataSourceFlattenSystemGlobalRemoteauthtimeout(o["remoteauthtimeout"], d, "remoteauthtimeout")); err != nil {
		if !fortiAPIPatch(o["remoteauthtimeout"]) {
			return fmt.Errorf("Error reading remoteauthtimeout: %v", err)
		}
	}

	if err = d.Set("alert_interval", dataSourceFlattenSystemGlobalAlertInterval(o["alert-interval"], d, "alert_interval")); err != nil {
		if !fortiAPIPatch(o["alert-interval"]) {
			return fmt.Errorf("Error reading alert_interval: %v", err)
		}
	}

	if err = d.Set("pre_login_banner", dataSourceFlattenSystemGlobalPreLoginBanner(o["pre-login-banner"], d, "pre_login_banner")); err != nil {
		if !fortiAPIPatch(o["pre-login-banner"]) {
			return fmt.Errorf("Error reading pre_login_banner: %v", err)
		}
	}

	if err = d.Set("switch_mgmt_mode", dataSourceFlattenSystemGlobalSwitchMgmtMode(o["switch-mgmt-mode"], d, "switch_mgmt_mode")); err != nil {
		if !fortiAPIPatch(o["switch-mgmt-mode"]) {
			return fmt.Errorf("Error reading switch_mgmt_mode: %v", err)
		}
	}

	if err = d.Set("strong_crypto", dataSourceFlattenSystemGlobalStrongCrypto(o["strong-crypto"], d, "strong_crypto")); err != nil {
		if !fortiAPIPatch(o["strong-crypto"]) {
			return fmt.Errorf("Error reading strong_crypto: %v", err)
		}
	}

	if err = d.Set("detect_ip_conflict", dataSourceFlattenSystemGlobalDetectIpConflict(o["detect-ip-conflict"], d, "detect_ip_conflict")); err != nil {
		if !fortiAPIPatch(o["detect-ip-conflict"]) {
			return fmt.Errorf("Error reading detect_ip_conflict: %v", err)
		}
	}

	if err = d.Set("asset_tag", dataSourceFlattenSystemGlobalAssetTag(o["asset-tag"], d, "asset_tag")); err != nil {
		if !fortiAPIPatch(o["asset-tag"]) {
			return fmt.Errorf("Error reading asset_tag: %v", err)
		}
	}

	if err = d.Set("ip_conflict_ignore_default", dataSourceFlattenSystemGlobalIpConflictIgnoreDefault(o["ip-conflict-ignore-default"], d, "ip_conflict_ignore_default")); err != nil {
		if !fortiAPIPatch(o["ip-conflict-ignore-default"]) {
			return fmt.Errorf("Error reading ip_conflict_ignore_default: %v", err)
		}
	}

	if err = d.Set("timezone", dataSourceFlattenSystemGlobalTimezone(o["timezone"], d, "timezone")); err != nil {
		if !fortiAPIPatch(o["timezone"]) {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

	if err = d.Set("admin_sport", dataSourceFlattenSystemGlobalAdminSport(o["admin-sport"], d, "admin_sport")); err != nil {
		if !fortiAPIPatch(o["admin-sport"]) {
			return fmt.Errorf("Error reading admin_sport: %v", err)
		}
	}

	if err = d.Set("image_rotation", dataSourceFlattenSystemGlobalImageRotation(o["image-rotation"], d, "image_rotation")); err != nil {
		if !fortiAPIPatch(o["image-rotation"]) {
			return fmt.Errorf("Error reading image_rotation: %v", err)
		}
	}

	if err = d.Set("admin_telnet_port", dataSourceFlattenSystemGlobalAdminTelnetPort(o["admin-telnet-port"], d, "admin_telnet_port")); err != nil {
		if !fortiAPIPatch(o["admin-telnet-port"]) {
			return fmt.Errorf("Error reading admin_telnet_port: %v", err)
		}
	}

	if err = d.Set("kernel_crashlog", dataSourceFlattenSystemGlobalKernelCrashlog(o["kernel-crashlog"], d, "kernel_crashlog")); err != nil {
		if !fortiAPIPatch(o["kernel-crashlog"]) {
			return fmt.Errorf("Error reading kernel_crashlog: %v", err)
		}
	}

	if err = d.Set("fortilink_auto_discovery", dataSourceFlattenSystemGlobalFortilinkAutoDiscovery(o["fortilink-auto-discovery"], d, "fortilink_auto_discovery")); err != nil {
		if !fortiAPIPatch(o["fortilink-auto-discovery"]) {
			return fmt.Errorf("Error reading fortilink_auto_discovery: %v", err)
		}
	}

	if err = d.Set("kernel_devicelog", dataSourceFlattenSystemGlobalKernelDevicelog(o["kernel-devicelog"], d, "kernel_devicelog")); err != nil {
		if !fortiAPIPatch(o["kernel-devicelog"]) {
			return fmt.Errorf("Error reading kernel_devicelog: %v", err)
		}
	}

	if err = d.Set("revision_backup_on_upgrade", dataSourceFlattenSystemGlobalRevisionBackupOnUpgrade(o["revision-backup-on-upgrade"], d, "revision_backup_on_upgrade")); err != nil {
		if !fortiAPIPatch(o["revision-backup-on-upgrade"]) {
			return fmt.Errorf("Error reading revision_backup_on_upgrade: %v", err)
		}
	}

	if err = d.Set("admin_https_ssl_versions", dataSourceFlattenSystemGlobalAdminHttpsSslVersions(o["admin-https-ssl-versions"], d, "admin_https_ssl_versions")); err != nil {
		if !fortiAPIPatch(o["admin-https-ssl-versions"]) {
			return fmt.Errorf("Error reading admin_https_ssl_versions: %v", err)
		}
	}

	if err = d.Set("hostname", dataSourceFlattenSystemGlobalHostname(o["hostname"], d, "hostname")); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("revision_backup_on_logout", dataSourceFlattenSystemGlobalRevisionBackupOnLogout(o["revision-backup-on-logout"], d, "revision_backup_on_logout")); err != nil {
		if !fortiAPIPatch(o["revision-backup-on-logout"]) {
			return fmt.Errorf("Error reading revision_backup_on_logout: %v", err)
		}
	}

	if err = d.Set("tcp6_mss_min", dataSourceFlattenSystemGlobalTcp6MssMin(o["tcp6-mss-min"], d, "tcp6_mss_min")); err != nil {
		if !fortiAPIPatch(o["tcp6-mss-min"]) {
			return fmt.Errorf("Error reading tcp6_mss_min: %v", err)
		}
	}

	if err = d.Set("cfg_revert_timeout", dataSourceFlattenSystemGlobalCfgRevertTimeout(o["cfg-revert-timeout"], d, "cfg_revert_timeout")); err != nil {
		if !fortiAPIPatch(o["cfg-revert-timeout"]) {
			return fmt.Errorf("Error reading cfg_revert_timeout: %v", err)
		}
	}

	if err = d.Set("admin_ssh_v1", dataSourceFlattenSystemGlobalAdminSshV1(o["admin-ssh-v1"], d, "admin_ssh_v1")); err != nil {
		if !fortiAPIPatch(o["admin-ssh-v1"]) {
			return fmt.Errorf("Error reading admin_ssh_v1: %v", err)
		}
	}

	if err = d.Set("allow_subnet_overlap", dataSourceFlattenSystemGlobalAllowSubnetOverlap(o["allow-subnet-overlap"], d, "allow_subnet_overlap")); err != nil {
		if !fortiAPIPatch(o["allow-subnet-overlap"]) {
			return fmt.Errorf("Error reading allow_subnet_overlap: %v", err)
		}
	}

	if err = d.Set("admin_password_hash", dataSourceFlattenSystemGlobalAdminPasswordHash(o["admin-password-hash"], d, "admin_password_hash")); err != nil {
		if !fortiAPIPatch(o["admin-password-hash"]) {
			return fmt.Errorf("Error reading admin_password_hash: %v", err)
		}
	}

	if err = d.Set("dh_params", dataSourceFlattenSystemGlobalDhParams(o["dh-params"], d, "dh_params")); err != nil {
		if !fortiAPIPatch(o["dh-params"]) {
			return fmt.Errorf("Error reading dh_params: %v", err)
		}
	}

	if err = d.Set("ldapconntimeout", dataSourceFlattenSystemGlobalLdapconntimeout(o["ldapconntimeout"], d, "ldapconntimeout")); err != nil {
		if !fortiAPIPatch(o["ldapconntimeout"]) {
			return fmt.Errorf("Error reading ldapconntimeout: %v", err)
		}
	}

	if err = d.Set("reset_button", dataSourceFlattenSystemGlobalResetButton(o["reset-button"], d, "reset_button")); err != nil {
		if !fortiAPIPatch(o["reset-button"]) {
			return fmt.Errorf("Error reading reset_button: %v", err)
		}
	}

	if err = d.Set("tcp_mss_min", dataSourceFlattenSystemGlobalTcpMssMin(o["tcp-mss-min"], d, "tcp_mss_min")); err != nil {
		if !fortiAPIPatch(o["tcp-mss-min"]) {
			return fmt.Errorf("Error reading tcp_mss_min: %v", err)
		}
	}

	if err = d.Set("admin_restrict_local", dataSourceFlattenSystemGlobalAdminRestrictLocal(o["admin-restrict-local"], d, "admin_restrict_local")); err != nil {
		if !fortiAPIPatch(o["admin-restrict-local"]) {
			return fmt.Errorf("Error reading admin_restrict_local: %v", err)
		}
	}

	if err = d.Set("admin_concurrent", dataSourceFlattenSystemGlobalAdminConcurrent(o["admin-concurrent"], d, "admin_concurrent")); err != nil {
		if !fortiAPIPatch(o["admin-concurrent"]) {
			return fmt.Errorf("Error reading admin_concurrent: %v", err)
		}
	}

	if err = d.Set("admintimeout", dataSourceFlattenSystemGlobalAdmintimeout(o["admintimeout"], d, "admintimeout")); err != nil {
		if !fortiAPIPatch(o["admintimeout"]) {
			return fmt.Errorf("Error reading admintimeout: %v", err)
		}
	}

	if err = d.Set("arp_timeout", dataSourceFlattenSystemGlobalArpTimeout(o["arp-timeout"], d, "arp_timeout")); err != nil {
		if !fortiAPIPatch(o["arp-timeout"]) {
			return fmt.Errorf("Error reading arp_timeout: %v", err)
		}
	}

	if err = d.Set("admin_lockout_duration", dataSourceFlattenSystemGlobalAdminLockoutDuration(o["admin-lockout-duration"], d, "admin_lockout_duration")); err != nil {
		if !fortiAPIPatch(o["admin-lockout-duration"]) {
			return fmt.Errorf("Error reading admin_lockout_duration: %v", err)
		}
	}

	if err = d.Set("dhcp_server_access_list", dataSourceFlattenSystemGlobalDhcpServerAccessList(o["dhcp-server-access-list"], d, "dhcp_server_access_list")); err != nil {
		if !fortiAPIPatch(o["dhcp-server-access-list"]) {
			return fmt.Errorf("Error reading dhcp_server_access_list: %v", err)
		}
	}

	if err = d.Set("admin_port", dataSourceFlattenSystemGlobalAdminPort(o["admin-port"], d, "admin_port")); err != nil {
		if !fortiAPIPatch(o["admin-port"]) {
			return fmt.Errorf("Error reading admin_port: %v", err)
		}
	}

	if err = d.Set("l3_host_expiry", dataSourceFlattenSystemGlobalL3HostExpiry(o["l3-host-expiry"], d, "l3_host_expiry")); err != nil {
		if !fortiAPIPatch(o["l3-host-expiry"]) {
			return fmt.Errorf("Error reading l3_host_expiry: %v", err)
		}
	}

	if err = d.Set("tcp_options", dataSourceFlattenSystemGlobalTcpOptions(o["tcp-options"], d, "tcp_options")); err != nil {
		if !fortiAPIPatch(o["tcp-options"]) {
			return fmt.Errorf("Error reading tcp_options: %v", err)
		}
	}

	if err = d.Set("post_login_banner", dataSourceFlattenSystemGlobalPostLoginBanner(o["post-login-banner"], d, "post_login_banner")); err != nil {
		if !fortiAPIPatch(o["post-login-banner"]) {
			return fmt.Errorf("Error reading post_login_banner: %v", err)
		}
	}

	if err = d.Set("failtime", dataSourceFlattenSystemGlobalFailtime(o["failtime"], d, "failtime")); err != nil {
		if !fortiAPIPatch(o["failtime"]) {
			return fmt.Errorf("Error reading failtime: %v", err)
		}
	}

	if err = d.Set("admin_lockout_threshold", dataSourceFlattenSystemGlobalAdminLockoutThreshold(o["admin-lockout-threshold"], d, "admin_lockout_threshold")); err != nil {
		if !fortiAPIPatch(o["admin-lockout-threshold"]) {
			return fmt.Errorf("Error reading admin_lockout_threshold: %v", err)
		}
	}

	if err = d.Set("dhcps_db_exp", dataSourceFlattenSystemGlobalDhcpsDbExp(o["dhcps-db-exp"], d, "dhcps_db_exp")); err != nil {
		if !fortiAPIPatch(o["dhcps-db-exp"]) {
			return fmt.Errorf("Error reading dhcps_db_exp: %v", err)
		}
	}

	if err = d.Set("n8021x_ca_certificate", dataSourceFlattenSystemGlobal8021XCaCertificate(o["802.1x-ca-certificate"], d, "n8021x_ca_certificate")); err != nil {
		if !fortiAPIPatch(o["802.1x-ca-certificate"]) {
			return fmt.Errorf("Error reading n8021x_ca_certificate: %v", err)
		}
	}

	if err = d.Set("dhcp_remote_id", dataSourceFlattenSystemGlobalDhcpRemoteId(o["dhcp-remote-id"], d, "dhcp_remote_id")); err != nil {
		if !fortiAPIPatch(o["dhcp-remote-id"]) {
			return fmt.Errorf("Error reading dhcp_remote_id: %v", err)
		}
	}

	if err = d.Set("dhcp_snoop_client_req", dataSourceFlattenSystemGlobalDhcpSnoopClientReq(o["dhcp-snoop-client-req"], d, "dhcp_snoop_client_req")); err != nil {
		if !fortiAPIPatch(o["dhcp-snoop-client-req"]) {
			return fmt.Errorf("Error reading dhcp_snoop_client_req: %v", err)
		}
	}

	if err = d.Set("arp_inspection_monitor_timeout", dataSourceFlattenSystemGlobalArpInspectionMonitorTimeout(o["arp-inspection-monitor-timeout"], d, "arp_inspection_monitor_timeout")); err != nil {
		if !fortiAPIPatch(o["arp-inspection-monitor-timeout"]) {
			return fmt.Errorf("Error reading arp_inspection_monitor_timeout: %v", err)
		}
	}

	if err = d.Set("dhcp_client_location", dataSourceFlattenSystemGlobalDhcpClientLocation(o["dhcp-client-location"], d, "dhcp_client_location")); err != nil {
		if !fortiAPIPatch(o["dhcp-client-location"]) {
			return fmt.Errorf("Error reading dhcp_client_location: %v", err)
		}
	}

	if err = d.Set("csr_ca_attribute", dataSourceFlattenSystemGlobalCsrCaAttribute(o["csr-ca-attribute"], d, "csr_ca_attribute")); err != nil {
		if !fortiAPIPatch(o["csr-ca-attribute"]) {
			return fmt.Errorf("Error reading csr_ca_attribute: %v", err)
		}
	}

	if err = d.Set("delaycli_timeout_cleanup", dataSourceFlattenSystemGlobalDelaycliTimeoutCleanup(o["delaycli-timeout-cleanup"], d, "delaycli_timeout_cleanup")); err != nil {
		if !fortiAPIPatch(o["delaycli-timeout-cleanup"]) {
			return fmt.Errorf("Error reading delaycli_timeout_cleanup: %v", err)
		}
	}

	if err = d.Set("ipv6_accept_dad", dataSourceFlattenSystemGlobalIpv6AcceptDad(o["ipv6-accept-dad"], d, "ipv6_accept_dad")); err != nil {
		if !fortiAPIPatch(o["ipv6-accept-dad"]) {
			return fmt.Errorf("Error reading ipv6_accept_dad: %v", err)
		}
	}

	if err = d.Set("admin_ssh_port", dataSourceFlattenSystemGlobalAdminSshPort(o["admin-ssh-port"], d, "admin_ssh_port")); err != nil {
		if !fortiAPIPatch(o["admin-ssh-port"]) {
			return fmt.Errorf("Error reading admin_ssh_port: %v", err)
		}
	}

	if err = d.Set("admin_server_cert", dataSourceFlattenSystemGlobalAdminServerCert(o["admin-server-cert"], d, "admin_server_cert")); err != nil {
		if !fortiAPIPatch(o["admin-server-cert"]) {
			return fmt.Errorf("Error reading admin_server_cert: %v", err)
		}
	}

	if err = d.Set("auto_isl", dataSourceFlattenSystemGlobalAutoIsl(o["auto-isl"], d, "auto_isl")); err != nil {
		if !fortiAPIPatch(o["auto-isl"]) {
			return fmt.Errorf("Error reading auto_isl: %v", err)
		}
	}

	if err = d.Set("language", dataSourceFlattenSystemGlobalLanguage(o["language"], d, "language")); err != nil {
		if !fortiAPIPatch(o["language"]) {
			return fmt.Errorf("Error reading language: %v", err)
		}
	}

	if err = d.Set("radius_coa_port", dataSourceFlattenSystemGlobalRadiusCoaPort(o["radius-coa-port"], d, "radius_coa_port")); err != nil {
		if !fortiAPIPatch(o["radius-coa-port"]) {
			return fmt.Errorf("Error reading radius_coa_port: %v", err)
		}
	}

	if err = d.Set("dst", dataSourceFlattenSystemGlobalDst(o["dst"], d, "dst")); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("interval", dataSourceFlattenSystemGlobalInterval(o["interval"], d, "interval")); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("cfg_save", dataSourceFlattenSystemGlobalCfgSave(o["cfg-save"], d, "cfg_save")); err != nil {
		if !fortiAPIPatch(o["cfg-save"]) {
			return fmt.Errorf("Error reading cfg_save: %v", err)
		}
	}

	if err = d.Set("restart_time", dataSourceFlattenSystemGlobalRestartTime(o["restart-time"], d, "restart_time")); err != nil {
		if !fortiAPIPatch(o["restart-time"]) {
			return fmt.Errorf("Error reading restart_time: %v", err)
		}
	}

	if err = d.Set("dhcp_option_format", dataSourceFlattenSystemGlobalDhcpOptionFormat(o["dhcp-option-format"], d, "dhcp_option_format")); err != nil {
		if !fortiAPIPatch(o["dhcp-option-format"]) {
			return fmt.Errorf("Error reading dhcp_option_format: %v", err)
		}
	}

	if err = d.Set("admin_scp", dataSourceFlattenSystemGlobalAdminScp(o["admin-scp"], d, "admin_scp")); err != nil {
		if !fortiAPIPatch(o["admin-scp"]) {
			return fmt.Errorf("Error reading admin_scp: %v", err)
		}
	}

	if err = d.Set("alertd_relog", dataSourceFlattenSystemGlobalAlertdRelog(o["alertd-relog"], d, "alertd_relog")); err != nil {
		if !fortiAPIPatch(o["alertd-relog"]) {
			return fmt.Errorf("Error reading alertd_relog: %v", err)
		}
	}

	if err = d.Set("clt_cert_req", dataSourceFlattenSystemGlobalCltCertReq(o["clt-cert-req"], d, "clt_cert_req")); err != nil {
		if !fortiAPIPatch(o["clt-cert-req"]) {
			return fmt.Errorf("Error reading clt_cert_req: %v", err)
		}
	}

	if err = d.Set("daily_restart", dataSourceFlattenSystemGlobalDailyRestart(o["daily-restart"], d, "daily_restart")); err != nil {
		if !fortiAPIPatch(o["daily-restart"]) {
			return fmt.Errorf("Error reading daily_restart: %v", err)
		}
	}

	if err = d.Set("tftp", dataSourceFlattenSystemGlobalTftp(o["tftp"], d, "tftp")); err != nil {
		if !fortiAPIPatch(o["tftp"]) {
			return fmt.Errorf("Error reading tftp: %v", err)
		}
	}

	if err = d.Set("admin_ssh_grace_time", dataSourceFlattenSystemGlobalAdminSshGraceTime(o["admin-ssh-grace-time"], d, "admin_ssh_grace_time")); err != nil {
		if !fortiAPIPatch(o["admin-ssh-grace-time"]) {
			return fmt.Errorf("Error reading admin_ssh_grace_time: %v", err)
		}
	}

	if err = d.Set("admin_https_pki_required", dataSourceFlattenSystemGlobalAdminHttpsPkiRequired(o["admin-https-pki-required"], d, "admin_https_pki_required")); err != nil {
		if !fortiAPIPatch(o["admin-https-pki-required"]) {
			return fmt.Errorf("Error reading admin_https_pki_required: %v", err)
		}
	}

	if err = d.Set("radius_port", dataSourceFlattenSystemGlobalRadiusPort(o["radius-port"], d, "radius_port")); err != nil {
		if !fortiAPIPatch(o["radius-port"]) {
			return fmt.Errorf("Error reading radius_port: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemGlobalFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
