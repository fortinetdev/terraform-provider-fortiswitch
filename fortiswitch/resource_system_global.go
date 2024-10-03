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

func resourceSystemGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemGlobalUpdate,
		Read:   resourceSystemGlobalRead,
		Update: resourceSystemGlobalUpdate,
		Delete: resourceSystemGlobalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"n8021x_certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ipv6_all_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcps_db_per_port_learn_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1024),
				Optional:     true,
				Computed:     true,
			},
			"private_data_encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_circuit_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remoteauthtimeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 300),
				Optional:     true,
				Computed:     true,
			},
			"alert_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pre_login_banner": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional:     true,
				Computed:     true,
			},
			"switch_mgmt_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strong_crypto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"detect_ip_conflict": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"asset_tag": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"ip_conflict_ignore_default": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_sport": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"image_rotation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_telnet_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"kernel_crashlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortilink_auto_discovery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kernel_devicelog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"revision_backup_on_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_https_ssl_versions": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"revision_backup_on_logout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp6_mss_min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cfg_revert_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_ssh_v1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_subnet_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_password_hash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dh_params": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ldapconntimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reset_button": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_mss_min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_restrict_local": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_concurrent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admintimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"arp_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_lockout_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcp_server_access_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"l3_host_expiry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"post_login_banner": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional:     true,
				Computed:     true,
			},
			"failtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_lockout_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcps_db_exp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 259200),
				Optional:     true,
				Computed:     true,
			},
			"n8021x_ca_certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"dhcp_remote_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_snoop_client_req": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp_inspection_monitor_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10080),
				Optional:     true,
				Computed:     true,
			},
			"dhcp_client_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"csr_ca_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delaycli_timeout_cleanup": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipv6_accept_dad": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2),
				Optional:     true,
				Computed:     true,
			},
			"admin_ssh_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"admin_server_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"auto_isl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_coa_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cfg_save": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restart_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_option_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_scp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alertd_relog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clt_cert_req": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"daily_restart": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tftp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_ssh_grace_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),
				Optional:     true,
				Computed:     true,
			},
			"admin_https_pki_required": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemGlobal(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemGlobal resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemGlobal(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemGlobal")
	}

	return resourceSystemGlobalRead(d, m)
}

func resourceSystemGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemGlobal(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemGlobal(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGlobalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemGlobal(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGlobal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemGlobal resource from API: %v", err)
	}
	return nil
}

func flattenSystemGlobal8021XCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIpv6AllForwarding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDhcpsDbPerPortLearnLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalPrivateDataEncryption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDhcpCircuitId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalRemoteauthtimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAlertInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalPreLoginBanner(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSwitchMgmtMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalStrongCrypto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDetectIpConflict(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAssetTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIpConflictIgnoreDefault(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTimezone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminSport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalImageRotation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminTelnetPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalKernelCrashlog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFortilinkAutoDiscovery(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalKernelDevicelog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalRevisionBackupOnUpgrade(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminHttpsSslVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalRevisionBackupOnLogout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTcp6MssMin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCfgRevertTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminSshV1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAllowSubnetOverlap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminPasswordHash(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDhParams(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalLdapconntimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalResetButton(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTcpMssMin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminRestrictLocal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminConcurrent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdmintimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalArpTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminLockoutDuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDhcpServerAccessList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalL3HostExpiry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTcpOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalPostLoginBanner(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFailtime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminLockoutThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDhcpsDbExp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobal8021XCaCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDhcpRemoteId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDhcpSnoopClientReq(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalArpInspectionMonitorTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDhcpClientLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCsrCaAttribute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDelaycliTimeoutCleanup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIpv6AcceptDad(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminSshPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAutoIsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalLanguage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalRadiusCoaPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCfgSave(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalRestartTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDhcpOptionFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminScp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAlertdRelog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCltCertReq(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDailyRestart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTftp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminSshGraceTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminHttpsPkiRequired(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalRadiusPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemGlobal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("n8021x_certificate", flattenSystemGlobal8021XCertificate(o["802.1x-certificate"], d, "n8021x_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["802.1x-certificate"]) {
			return fmt.Errorf("Error reading n8021x_certificate: %v", err)
		}
	}

	if err = d.Set("ipv6_all_forwarding", flattenSystemGlobalIpv6AllForwarding(o["ipv6-all-forwarding"], d, "ipv6_all_forwarding", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-all-forwarding"]) {
			return fmt.Errorf("Error reading ipv6_all_forwarding: %v", err)
		}
	}

	if err = d.Set("dhcps_db_per_port_learn_limit", flattenSystemGlobalDhcpsDbPerPortLearnLimit(o["dhcps-db-per-port-learn-limit"], d, "dhcps_db_per_port_learn_limit", sv)); err != nil {
		if !fortiAPIPatch(o["dhcps-db-per-port-learn-limit"]) {
			return fmt.Errorf("Error reading dhcps_db_per_port_learn_limit: %v", err)
		}
	}

	if err = d.Set("private_data_encryption", flattenSystemGlobalPrivateDataEncryption(o["private-data-encryption"], d, "private_data_encryption", sv)); err != nil {
		if !fortiAPIPatch(o["private-data-encryption"]) {
			return fmt.Errorf("Error reading private_data_encryption: %v", err)
		}
	}

	if err = d.Set("dhcp_circuit_id", flattenSystemGlobalDhcpCircuitId(o["dhcp-circuit-id"], d, "dhcp_circuit_id", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-circuit-id"]) {
			return fmt.Errorf("Error reading dhcp_circuit_id: %v", err)
		}
	}

	if err = d.Set("remoteauthtimeout", flattenSystemGlobalRemoteauthtimeout(o["remoteauthtimeout"], d, "remoteauthtimeout", sv)); err != nil {
		if !fortiAPIPatch(o["remoteauthtimeout"]) {
			return fmt.Errorf("Error reading remoteauthtimeout: %v", err)
		}
	}

	if err = d.Set("alert_interval", flattenSystemGlobalAlertInterval(o["alert-interval"], d, "alert_interval", sv)); err != nil {
		if !fortiAPIPatch(o["alert-interval"]) {
			return fmt.Errorf("Error reading alert_interval: %v", err)
		}
	}

	if err = d.Set("pre_login_banner", flattenSystemGlobalPreLoginBanner(o["pre-login-banner"], d, "pre_login_banner", sv)); err != nil {
		if !fortiAPIPatch(o["pre-login-banner"]) {
			return fmt.Errorf("Error reading pre_login_banner: %v", err)
		}
	}

	if err = d.Set("switch_mgmt_mode", flattenSystemGlobalSwitchMgmtMode(o["switch-mgmt-mode"], d, "switch_mgmt_mode", sv)); err != nil {
		if !fortiAPIPatch(o["switch-mgmt-mode"]) {
			return fmt.Errorf("Error reading switch_mgmt_mode: %v", err)
		}
	}

	if err = d.Set("strong_crypto", flattenSystemGlobalStrongCrypto(o["strong-crypto"], d, "strong_crypto", sv)); err != nil {
		if !fortiAPIPatch(o["strong-crypto"]) {
			return fmt.Errorf("Error reading strong_crypto: %v", err)
		}
	}

	if err = d.Set("detect_ip_conflict", flattenSystemGlobalDetectIpConflict(o["detect-ip-conflict"], d, "detect_ip_conflict", sv)); err != nil {
		if !fortiAPIPatch(o["detect-ip-conflict"]) {
			return fmt.Errorf("Error reading detect_ip_conflict: %v", err)
		}
	}

	if err = d.Set("asset_tag", flattenSystemGlobalAssetTag(o["asset-tag"], d, "asset_tag", sv)); err != nil {
		if !fortiAPIPatch(o["asset-tag"]) {
			return fmt.Errorf("Error reading asset_tag: %v", err)
		}
	}

	if err = d.Set("ip_conflict_ignore_default", flattenSystemGlobalIpConflictIgnoreDefault(o["ip-conflict-ignore-default"], d, "ip_conflict_ignore_default", sv)); err != nil {
		if !fortiAPIPatch(o["ip-conflict-ignore-default"]) {
			return fmt.Errorf("Error reading ip_conflict_ignore_default: %v", err)
		}
	}

	if err = d.Set("timezone", flattenSystemGlobalTimezone(o["timezone"], d, "timezone", sv)); err != nil {
		if !fortiAPIPatch(o["timezone"]) {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

	if err = d.Set("admin_sport", flattenSystemGlobalAdminSport(o["admin-sport"], d, "admin_sport", sv)); err != nil {
		if !fortiAPIPatch(o["admin-sport"]) {
			return fmt.Errorf("Error reading admin_sport: %v", err)
		}
	}

	if err = d.Set("image_rotation", flattenSystemGlobalImageRotation(o["image-rotation"], d, "image_rotation", sv)); err != nil {
		if !fortiAPIPatch(o["image-rotation"]) {
			return fmt.Errorf("Error reading image_rotation: %v", err)
		}
	}

	if err = d.Set("admin_telnet_port", flattenSystemGlobalAdminTelnetPort(o["admin-telnet-port"], d, "admin_telnet_port", sv)); err != nil {
		if !fortiAPIPatch(o["admin-telnet-port"]) {
			return fmt.Errorf("Error reading admin_telnet_port: %v", err)
		}
	}

	if err = d.Set("kernel_crashlog", flattenSystemGlobalKernelCrashlog(o["kernel-crashlog"], d, "kernel_crashlog", sv)); err != nil {
		if !fortiAPIPatch(o["kernel-crashlog"]) {
			return fmt.Errorf("Error reading kernel_crashlog: %v", err)
		}
	}

	if err = d.Set("fortilink_auto_discovery", flattenSystemGlobalFortilinkAutoDiscovery(o["fortilink-auto-discovery"], d, "fortilink_auto_discovery", sv)); err != nil {
		if !fortiAPIPatch(o["fortilink-auto-discovery"]) {
			return fmt.Errorf("Error reading fortilink_auto_discovery: %v", err)
		}
	}

	if err = d.Set("kernel_devicelog", flattenSystemGlobalKernelDevicelog(o["kernel-devicelog"], d, "kernel_devicelog", sv)); err != nil {
		if !fortiAPIPatch(o["kernel-devicelog"]) {
			return fmt.Errorf("Error reading kernel_devicelog: %v", err)
		}
	}

	if err = d.Set("revision_backup_on_upgrade", flattenSystemGlobalRevisionBackupOnUpgrade(o["revision-backup-on-upgrade"], d, "revision_backup_on_upgrade", sv)); err != nil {
		if !fortiAPIPatch(o["revision-backup-on-upgrade"]) {
			return fmt.Errorf("Error reading revision_backup_on_upgrade: %v", err)
		}
	}

	if err = d.Set("admin_https_ssl_versions", flattenSystemGlobalAdminHttpsSslVersions(o["admin-https-ssl-versions"], d, "admin_https_ssl_versions", sv)); err != nil {
		if !fortiAPIPatch(o["admin-https-ssl-versions"]) {
			return fmt.Errorf("Error reading admin_https_ssl_versions: %v", err)
		}
	}

	if err = d.Set("hostname", flattenSystemGlobalHostname(o["hostname"], d, "hostname", sv)); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("revision_backup_on_logout", flattenSystemGlobalRevisionBackupOnLogout(o["revision-backup-on-logout"], d, "revision_backup_on_logout", sv)); err != nil {
		if !fortiAPIPatch(o["revision-backup-on-logout"]) {
			return fmt.Errorf("Error reading revision_backup_on_logout: %v", err)
		}
	}

	if err = d.Set("tcp6_mss_min", flattenSystemGlobalTcp6MssMin(o["tcp6-mss-min"], d, "tcp6_mss_min", sv)); err != nil {
		if !fortiAPIPatch(o["tcp6-mss-min"]) {
			return fmt.Errorf("Error reading tcp6_mss_min: %v", err)
		}
	}

	if err = d.Set("cfg_revert_timeout", flattenSystemGlobalCfgRevertTimeout(o["cfg-revert-timeout"], d, "cfg_revert_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["cfg-revert-timeout"]) {
			return fmt.Errorf("Error reading cfg_revert_timeout: %v", err)
		}
	}

	if err = d.Set("admin_ssh_v1", flattenSystemGlobalAdminSshV1(o["admin-ssh-v1"], d, "admin_ssh_v1", sv)); err != nil {
		if !fortiAPIPatch(o["admin-ssh-v1"]) {
			return fmt.Errorf("Error reading admin_ssh_v1: %v", err)
		}
	}

	if err = d.Set("allow_subnet_overlap", flattenSystemGlobalAllowSubnetOverlap(o["allow-subnet-overlap"], d, "allow_subnet_overlap", sv)); err != nil {
		if !fortiAPIPatch(o["allow-subnet-overlap"]) {
			return fmt.Errorf("Error reading allow_subnet_overlap: %v", err)
		}
	}

	if err = d.Set("admin_password_hash", flattenSystemGlobalAdminPasswordHash(o["admin-password-hash"], d, "admin_password_hash", sv)); err != nil {
		if !fortiAPIPatch(o["admin-password-hash"]) {
			return fmt.Errorf("Error reading admin_password_hash: %v", err)
		}
	}

	if err = d.Set("dh_params", flattenSystemGlobalDhParams(o["dh-params"], d, "dh_params", sv)); err != nil {
		if !fortiAPIPatch(o["dh-params"]) {
			return fmt.Errorf("Error reading dh_params: %v", err)
		}
	}

	if err = d.Set("ldapconntimeout", flattenSystemGlobalLdapconntimeout(o["ldapconntimeout"], d, "ldapconntimeout", sv)); err != nil {
		if !fortiAPIPatch(o["ldapconntimeout"]) {
			return fmt.Errorf("Error reading ldapconntimeout: %v", err)
		}
	}

	if err = d.Set("reset_button", flattenSystemGlobalResetButton(o["reset-button"], d, "reset_button", sv)); err != nil {
		if !fortiAPIPatch(o["reset-button"]) {
			return fmt.Errorf("Error reading reset_button: %v", err)
		}
	}

	if err = d.Set("tcp_mss_min", flattenSystemGlobalTcpMssMin(o["tcp-mss-min"], d, "tcp_mss_min", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-mss-min"]) {
			return fmt.Errorf("Error reading tcp_mss_min: %v", err)
		}
	}

	if err = d.Set("admin_restrict_local", flattenSystemGlobalAdminRestrictLocal(o["admin-restrict-local"], d, "admin_restrict_local", sv)); err != nil {
		if !fortiAPIPatch(o["admin-restrict-local"]) {
			return fmt.Errorf("Error reading admin_restrict_local: %v", err)
		}
	}

	if err = d.Set("admin_concurrent", flattenSystemGlobalAdminConcurrent(o["admin-concurrent"], d, "admin_concurrent", sv)); err != nil {
		if !fortiAPIPatch(o["admin-concurrent"]) {
			return fmt.Errorf("Error reading admin_concurrent: %v", err)
		}
	}

	if err = d.Set("admintimeout", flattenSystemGlobalAdmintimeout(o["admintimeout"], d, "admintimeout", sv)); err != nil {
		if !fortiAPIPatch(o["admintimeout"]) {
			return fmt.Errorf("Error reading admintimeout: %v", err)
		}
	}

	if err = d.Set("arp_timeout", flattenSystemGlobalArpTimeout(o["arp-timeout"], d, "arp_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["arp-timeout"]) {
			return fmt.Errorf("Error reading arp_timeout: %v", err)
		}
	}

	if err = d.Set("admin_lockout_duration", flattenSystemGlobalAdminLockoutDuration(o["admin-lockout-duration"], d, "admin_lockout_duration", sv)); err != nil {
		if !fortiAPIPatch(o["admin-lockout-duration"]) {
			return fmt.Errorf("Error reading admin_lockout_duration: %v", err)
		}
	}

	if err = d.Set("dhcp_server_access_list", flattenSystemGlobalDhcpServerAccessList(o["dhcp-server-access-list"], d, "dhcp_server_access_list", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-server-access-list"]) {
			return fmt.Errorf("Error reading dhcp_server_access_list: %v", err)
		}
	}

	if err = d.Set("admin_port", flattenSystemGlobalAdminPort(o["admin-port"], d, "admin_port", sv)); err != nil {
		if !fortiAPIPatch(o["admin-port"]) {
			return fmt.Errorf("Error reading admin_port: %v", err)
		}
	}

	if err = d.Set("l3_host_expiry", flattenSystemGlobalL3HostExpiry(o["l3-host-expiry"], d, "l3_host_expiry", sv)); err != nil {
		if !fortiAPIPatch(o["l3-host-expiry"]) {
			return fmt.Errorf("Error reading l3_host_expiry: %v", err)
		}
	}

	if err = d.Set("tcp_options", flattenSystemGlobalTcpOptions(o["tcp-options"], d, "tcp_options", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-options"]) {
			return fmt.Errorf("Error reading tcp_options: %v", err)
		}
	}

	if err = d.Set("post_login_banner", flattenSystemGlobalPostLoginBanner(o["post-login-banner"], d, "post_login_banner", sv)); err != nil {
		if !fortiAPIPatch(o["post-login-banner"]) {
			return fmt.Errorf("Error reading post_login_banner: %v", err)
		}
	}

	if err = d.Set("failtime", flattenSystemGlobalFailtime(o["failtime"], d, "failtime", sv)); err != nil {
		if !fortiAPIPatch(o["failtime"]) {
			return fmt.Errorf("Error reading failtime: %v", err)
		}
	}

	if err = d.Set("admin_lockout_threshold", flattenSystemGlobalAdminLockoutThreshold(o["admin-lockout-threshold"], d, "admin_lockout_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["admin-lockout-threshold"]) {
			return fmt.Errorf("Error reading admin_lockout_threshold: %v", err)
		}
	}

	if err = d.Set("dhcps_db_exp", flattenSystemGlobalDhcpsDbExp(o["dhcps-db-exp"], d, "dhcps_db_exp", sv)); err != nil {
		if !fortiAPIPatch(o["dhcps-db-exp"]) {
			return fmt.Errorf("Error reading dhcps_db_exp: %v", err)
		}
	}

	if err = d.Set("n8021x_ca_certificate", flattenSystemGlobal8021XCaCertificate(o["802.1x-ca-certificate"], d, "n8021x_ca_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["802.1x-ca-certificate"]) {
			return fmt.Errorf("Error reading n8021x_ca_certificate: %v", err)
		}
	}

	if err = d.Set("dhcp_remote_id", flattenSystemGlobalDhcpRemoteId(o["dhcp-remote-id"], d, "dhcp_remote_id", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-remote-id"]) {
			return fmt.Errorf("Error reading dhcp_remote_id: %v", err)
		}
	}

	if err = d.Set("dhcp_snoop_client_req", flattenSystemGlobalDhcpSnoopClientReq(o["dhcp-snoop-client-req"], d, "dhcp_snoop_client_req", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-snoop-client-req"]) {
			return fmt.Errorf("Error reading dhcp_snoop_client_req: %v", err)
		}
	}

	if err = d.Set("arp_inspection_monitor_timeout", flattenSystemGlobalArpInspectionMonitorTimeout(o["arp-inspection-monitor-timeout"], d, "arp_inspection_monitor_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["arp-inspection-monitor-timeout"]) {
			return fmt.Errorf("Error reading arp_inspection_monitor_timeout: %v", err)
		}
	}

	if err = d.Set("dhcp_client_location", flattenSystemGlobalDhcpClientLocation(o["dhcp-client-location"], d, "dhcp_client_location", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-client-location"]) {
			return fmt.Errorf("Error reading dhcp_client_location: %v", err)
		}
	}

	if err = d.Set("csr_ca_attribute", flattenSystemGlobalCsrCaAttribute(o["csr-ca-attribute"], d, "csr_ca_attribute", sv)); err != nil {
		if !fortiAPIPatch(o["csr-ca-attribute"]) {
			return fmt.Errorf("Error reading csr_ca_attribute: %v", err)
		}
	}

	if err = d.Set("delaycli_timeout_cleanup", flattenSystemGlobalDelaycliTimeoutCleanup(o["delaycli-timeout-cleanup"], d, "delaycli_timeout_cleanup", sv)); err != nil {
		if !fortiAPIPatch(o["delaycli-timeout-cleanup"]) {
			return fmt.Errorf("Error reading delaycli_timeout_cleanup: %v", err)
		}
	}

	if err = d.Set("ipv6_accept_dad", flattenSystemGlobalIpv6AcceptDad(o["ipv6-accept-dad"], d, "ipv6_accept_dad", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-accept-dad"]) {
			return fmt.Errorf("Error reading ipv6_accept_dad: %v", err)
		}
	}

	if err = d.Set("admin_ssh_port", flattenSystemGlobalAdminSshPort(o["admin-ssh-port"], d, "admin_ssh_port", sv)); err != nil {
		if !fortiAPIPatch(o["admin-ssh-port"]) {
			return fmt.Errorf("Error reading admin_ssh_port: %v", err)
		}
	}

	if err = d.Set("admin_server_cert", flattenSystemGlobalAdminServerCert(o["admin-server-cert"], d, "admin_server_cert", sv)); err != nil {
		if !fortiAPIPatch(o["admin-server-cert"]) {
			return fmt.Errorf("Error reading admin_server_cert: %v", err)
		}
	}

	if err = d.Set("auto_isl", flattenSystemGlobalAutoIsl(o["auto-isl"], d, "auto_isl", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl"]) {
			return fmt.Errorf("Error reading auto_isl: %v", err)
		}
	}

	if err = d.Set("language", flattenSystemGlobalLanguage(o["language"], d, "language", sv)); err != nil {
		if !fortiAPIPatch(o["language"]) {
			return fmt.Errorf("Error reading language: %v", err)
		}
	}

	if err = d.Set("radius_coa_port", flattenSystemGlobalRadiusCoaPort(o["radius-coa-port"], d, "radius_coa_port", sv)); err != nil {
		if !fortiAPIPatch(o["radius-coa-port"]) {
			return fmt.Errorf("Error reading radius_coa_port: %v", err)
		}
	}

	if err = d.Set("dst", flattenSystemGlobalDst(o["dst"], d, "dst", sv)); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("interval", flattenSystemGlobalInterval(o["interval"], d, "interval", sv)); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("cfg_save", flattenSystemGlobalCfgSave(o["cfg-save"], d, "cfg_save", sv)); err != nil {
		if !fortiAPIPatch(o["cfg-save"]) {
			return fmt.Errorf("Error reading cfg_save: %v", err)
		}
	}

	if err = d.Set("restart_time", flattenSystemGlobalRestartTime(o["restart-time"], d, "restart_time", sv)); err != nil {
		if !fortiAPIPatch(o["restart-time"]) {
			return fmt.Errorf("Error reading restart_time: %v", err)
		}
	}

	if err = d.Set("dhcp_option_format", flattenSystemGlobalDhcpOptionFormat(o["dhcp-option-format"], d, "dhcp_option_format", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-option-format"]) {
			return fmt.Errorf("Error reading dhcp_option_format: %v", err)
		}
	}

	if err = d.Set("admin_scp", flattenSystemGlobalAdminScp(o["admin-scp"], d, "admin_scp", sv)); err != nil {
		if !fortiAPIPatch(o["admin-scp"]) {
			return fmt.Errorf("Error reading admin_scp: %v", err)
		}
	}

	if err = d.Set("alertd_relog", flattenSystemGlobalAlertdRelog(o["alertd-relog"], d, "alertd_relog", sv)); err != nil {
		if !fortiAPIPatch(o["alertd-relog"]) {
			return fmt.Errorf("Error reading alertd_relog: %v", err)
		}
	}

	if err = d.Set("clt_cert_req", flattenSystemGlobalCltCertReq(o["clt-cert-req"], d, "clt_cert_req", sv)); err != nil {
		if !fortiAPIPatch(o["clt-cert-req"]) {
			return fmt.Errorf("Error reading clt_cert_req: %v", err)
		}
	}

	if err = d.Set("daily_restart", flattenSystemGlobalDailyRestart(o["daily-restart"], d, "daily_restart", sv)); err != nil {
		if !fortiAPIPatch(o["daily-restart"]) {
			return fmt.Errorf("Error reading daily_restart: %v", err)
		}
	}

	if err = d.Set("tftp", flattenSystemGlobalTftp(o["tftp"], d, "tftp", sv)); err != nil {
		if !fortiAPIPatch(o["tftp"]) {
			return fmt.Errorf("Error reading tftp: %v", err)
		}
	}

	if err = d.Set("admin_ssh_grace_time", flattenSystemGlobalAdminSshGraceTime(o["admin-ssh-grace-time"], d, "admin_ssh_grace_time", sv)); err != nil {
		if !fortiAPIPatch(o["admin-ssh-grace-time"]) {
			return fmt.Errorf("Error reading admin_ssh_grace_time: %v", err)
		}
	}

	if err = d.Set("admin_https_pki_required", flattenSystemGlobalAdminHttpsPkiRequired(o["admin-https-pki-required"], d, "admin_https_pki_required", sv)); err != nil {
		if !fortiAPIPatch(o["admin-https-pki-required"]) {
			return fmt.Errorf("Error reading admin_https_pki_required: %v", err)
		}
	}

	if err = d.Set("radius_port", flattenSystemGlobalRadiusPort(o["radius-port"], d, "radius_port", sv)); err != nil {
		if !fortiAPIPatch(o["radius-port"]) {
			return fmt.Errorf("Error reading radius_port: %v", err)
		}
	}

	return nil
}

func flattenSystemGlobalFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemGlobal8021XCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpv6AllForwarding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDhcpsDbPerPortLearnLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPrivateDataEncryption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDhcpCircuitId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRemoteauthtimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAlertInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPreLoginBanner(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSwitchMgmtMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalStrongCrypto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDetectIpConflict(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAssetTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpConflictIgnoreDefault(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTimezone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminSport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalImageRotation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminTelnetPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalKernelCrashlog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortilinkAutoDiscovery(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalKernelDevicelog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRevisionBackupOnUpgrade(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminHttpsSslVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRevisionBackupOnLogout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTcp6MssMin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCfgRevertTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminSshV1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAllowSubnetOverlap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminPasswordHash(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDhParams(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLdapconntimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalResetButton(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTcpMssMin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminRestrictLocal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminConcurrent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdmintimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalArpTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminLockoutDuration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDhcpServerAccessList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalL3HostExpiry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTcpOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPostLoginBanner(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFailtime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminLockoutThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDhcpsDbExp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobal8021XCaCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDhcpRemoteId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDhcpSnoopClientReq(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalArpInspectionMonitorTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDhcpClientLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCsrCaAttribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDelaycliTimeoutCleanup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpv6AcceptDad(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminSshPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAutoIsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLanguage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRadiusCoaPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCfgSave(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRestartTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDhcpOptionFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminScp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAlertdRelog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCltCertReq(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDailyRestart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTftp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminSshGraceTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminHttpsPkiRequired(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRadiusPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGlobal(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n8021x_certificate"); ok {
		if setArgNil {
			obj["802.1x-certificate"] = nil
		} else {

			t, err := expandSystemGlobal8021XCertificate(d, v, "n8021x_certificate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["802.1x-certificate"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipv6_all_forwarding"); ok {
		if setArgNil {
			obj["ipv6-all-forwarding"] = nil
		} else {

			t, err := expandSystemGlobalIpv6AllForwarding(d, v, "ipv6_all_forwarding", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipv6-all-forwarding"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("dhcps_db_per_port_learn_limit"); ok {
		if setArgNil {
			obj["dhcps-db-per-port-learn-limit"] = nil
		} else {

			t, err := expandSystemGlobalDhcpsDbPerPortLearnLimit(d, v, "dhcps_db_per_port_learn_limit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcps-db-per-port-learn-limit"] = t
			}
		}
	}

	if v, ok := d.GetOk("private_data_encryption"); ok {
		if setArgNil {
			obj["private-data-encryption"] = nil
		} else {

			t, err := expandSystemGlobalPrivateDataEncryption(d, v, "private_data_encryption", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["private-data-encryption"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_circuit_id"); ok {
		if setArgNil {
			obj["dhcp-circuit-id"] = nil
		} else {

			t, err := expandSystemGlobalDhcpCircuitId(d, v, "dhcp_circuit_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-circuit-id"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("remoteauthtimeout"); ok {
		if setArgNil {
			obj["remoteauthtimeout"] = nil
		} else {

			t, err := expandSystemGlobalRemoteauthtimeout(d, v, "remoteauthtimeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["remoteauthtimeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("alert_interval"); ok {
		if setArgNil {
			obj["alert-interval"] = nil
		} else {

			t, err := expandSystemGlobalAlertInterval(d, v, "alert_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["alert-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("pre_login_banner"); ok {
		if setArgNil {
			obj["pre-login-banner"] = nil
		} else {

			t, err := expandSystemGlobalPreLoginBanner(d, v, "pre_login_banner", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pre-login-banner"] = t
			}
		}
	}

	if v, ok := d.GetOk("switch_mgmt_mode"); ok {
		if setArgNil {
			obj["switch-mgmt-mode"] = nil
		} else {

			t, err := expandSystemGlobalSwitchMgmtMode(d, v, "switch_mgmt_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["switch-mgmt-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("strong_crypto"); ok {
		if setArgNil {
			obj["strong-crypto"] = nil
		} else {

			t, err := expandSystemGlobalStrongCrypto(d, v, "strong_crypto", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["strong-crypto"] = t
			}
		}
	}

	if v, ok := d.GetOk("detect_ip_conflict"); ok {
		if setArgNil {
			obj["detect-ip-conflict"] = nil
		} else {

			t, err := expandSystemGlobalDetectIpConflict(d, v, "detect_ip_conflict", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["detect-ip-conflict"] = t
			}
		}
	}

	if v, ok := d.GetOk("asset_tag"); ok {
		if setArgNil {
			obj["asset-tag"] = nil
		} else {

			t, err := expandSystemGlobalAssetTag(d, v, "asset_tag", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["asset-tag"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip_conflict_ignore_default"); ok {
		if setArgNil {
			obj["ip-conflict-ignore-default"] = nil
		} else {

			t, err := expandSystemGlobalIpConflictIgnoreDefault(d, v, "ip_conflict_ignore_default", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip-conflict-ignore-default"] = t
			}
		}
	}

	if v, ok := d.GetOk("timezone"); ok {
		if setArgNil {
			obj["timezone"] = nil
		} else {

			t, err := expandSystemGlobalTimezone(d, v, "timezone", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timezone"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_sport"); ok {
		if setArgNil {
			obj["admin-sport"] = nil
		} else {

			t, err := expandSystemGlobalAdminSport(d, v, "admin_sport", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-sport"] = t
			}
		}
	}

	if v, ok := d.GetOk("image_rotation"); ok {
		if setArgNil {
			obj["image-rotation"] = nil
		} else {

			t, err := expandSystemGlobalImageRotation(d, v, "image_rotation", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["image-rotation"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_telnet_port"); ok {
		if setArgNil {
			obj["admin-telnet-port"] = nil
		} else {

			t, err := expandSystemGlobalAdminTelnetPort(d, v, "admin_telnet_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-telnet-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("kernel_crashlog"); ok {
		if setArgNil {
			obj["kernel-crashlog"] = nil
		} else {

			t, err := expandSystemGlobalKernelCrashlog(d, v, "kernel_crashlog", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["kernel-crashlog"] = t
			}
		}
	}

	if v, ok := d.GetOk("fortilink_auto_discovery"); ok {
		if setArgNil {
			obj["fortilink-auto-discovery"] = nil
		} else {

			t, err := expandSystemGlobalFortilinkAutoDiscovery(d, v, "fortilink_auto_discovery", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fortilink-auto-discovery"] = t
			}
		}
	}

	if v, ok := d.GetOk("kernel_devicelog"); ok {
		if setArgNil {
			obj["kernel-devicelog"] = nil
		} else {

			t, err := expandSystemGlobalKernelDevicelog(d, v, "kernel_devicelog", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["kernel-devicelog"] = t
			}
		}
	}

	if v, ok := d.GetOk("revision_backup_on_upgrade"); ok {
		if setArgNil {
			obj["revision-backup-on-upgrade"] = nil
		} else {

			t, err := expandSystemGlobalRevisionBackupOnUpgrade(d, v, "revision_backup_on_upgrade", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["revision-backup-on-upgrade"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_https_ssl_versions"); ok {
		if setArgNil {
			obj["admin-https-ssl-versions"] = nil
		} else {

			t, err := expandSystemGlobalAdminHttpsSslVersions(d, v, "admin_https_ssl_versions", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-https-ssl-versions"] = t
			}
		}
	}

	if v, ok := d.GetOk("hostname"); ok {
		if setArgNil {
			obj["hostname"] = nil
		} else {

			t, err := expandSystemGlobalHostname(d, v, "hostname", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostname"] = t
			}
		}
	}

	if v, ok := d.GetOk("revision_backup_on_logout"); ok {
		if setArgNil {
			obj["revision-backup-on-logout"] = nil
		} else {

			t, err := expandSystemGlobalRevisionBackupOnLogout(d, v, "revision_backup_on_logout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["revision-backup-on-logout"] = t
			}
		}
	}

	if v, ok := d.GetOk("tcp6_mss_min"); ok {
		if setArgNil {
			obj["tcp6-mss-min"] = nil
		} else {

			t, err := expandSystemGlobalTcp6MssMin(d, v, "tcp6_mss_min", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tcp6-mss-min"] = t
			}
		}
	}

	if v, ok := d.GetOk("cfg_revert_timeout"); ok {
		if setArgNil {
			obj["cfg-revert-timeout"] = nil
		} else {

			t, err := expandSystemGlobalCfgRevertTimeout(d, v, "cfg_revert_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cfg-revert-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_ssh_v1"); ok {
		if setArgNil {
			obj["admin-ssh-v1"] = nil
		} else {

			t, err := expandSystemGlobalAdminSshV1(d, v, "admin_ssh_v1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-ssh-v1"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_subnet_overlap"); ok {
		if setArgNil {
			obj["allow-subnet-overlap"] = nil
		} else {

			t, err := expandSystemGlobalAllowSubnetOverlap(d, v, "allow_subnet_overlap", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-subnet-overlap"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_password_hash"); ok {
		if setArgNil {
			obj["admin-password-hash"] = nil
		} else {

			t, err := expandSystemGlobalAdminPasswordHash(d, v, "admin_password_hash", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-password-hash"] = t
			}
		}
	}

	if v, ok := d.GetOk("dh_params"); ok {
		if setArgNil {
			obj["dh-params"] = nil
		} else {

			t, err := expandSystemGlobalDhParams(d, v, "dh_params", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dh-params"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("ldapconntimeout"); ok {
		if setArgNil {
			obj["ldapconntimeout"] = nil
		} else {

			t, err := expandSystemGlobalLdapconntimeout(d, v, "ldapconntimeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ldapconntimeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("reset_button"); ok {
		if setArgNil {
			obj["reset-button"] = nil
		} else {

			t, err := expandSystemGlobalResetButton(d, v, "reset_button", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["reset-button"] = t
			}
		}
	}

	if v, ok := d.GetOk("tcp_mss_min"); ok {
		if setArgNil {
			obj["tcp-mss-min"] = nil
		} else {

			t, err := expandSystemGlobalTcpMssMin(d, v, "tcp_mss_min", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tcp-mss-min"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_restrict_local"); ok {
		if setArgNil {
			obj["admin-restrict-local"] = nil
		} else {

			t, err := expandSystemGlobalAdminRestrictLocal(d, v, "admin_restrict_local", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-restrict-local"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_concurrent"); ok {
		if setArgNil {
			obj["admin-concurrent"] = nil
		} else {

			t, err := expandSystemGlobalAdminConcurrent(d, v, "admin_concurrent", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-concurrent"] = t
			}
		}
	}

	if v, ok := d.GetOk("admintimeout"); ok {
		if setArgNil {
			obj["admintimeout"] = nil
		} else {

			t, err := expandSystemGlobalAdmintimeout(d, v, "admintimeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admintimeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("arp_timeout"); ok {
		if setArgNil {
			obj["arp-timeout"] = nil
		} else {

			t, err := expandSystemGlobalArpTimeout(d, v, "arp_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["arp-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_lockout_duration"); ok {
		if setArgNil {
			obj["admin-lockout-duration"] = nil
		} else {

			t, err := expandSystemGlobalAdminLockoutDuration(d, v, "admin_lockout_duration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-lockout-duration"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_server_access_list"); ok {
		if setArgNil {
			obj["dhcp-server-access-list"] = nil
		} else {

			t, err := expandSystemGlobalDhcpServerAccessList(d, v, "dhcp_server_access_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-server-access-list"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_port"); ok {
		if setArgNil {
			obj["admin-port"] = nil
		} else {

			t, err := expandSystemGlobalAdminPort(d, v, "admin_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("l3_host_expiry"); ok {
		if setArgNil {
			obj["l3-host-expiry"] = nil
		} else {

			t, err := expandSystemGlobalL3HostExpiry(d, v, "l3_host_expiry", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["l3-host-expiry"] = t
			}
		}
	}

	if v, ok := d.GetOk("tcp_options"); ok {
		if setArgNil {
			obj["tcp-options"] = nil
		} else {

			t, err := expandSystemGlobalTcpOptions(d, v, "tcp_options", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tcp-options"] = t
			}
		}
	}

	if v, ok := d.GetOk("post_login_banner"); ok {
		if setArgNil {
			obj["post-login-banner"] = nil
		} else {

			t, err := expandSystemGlobalPostLoginBanner(d, v, "post_login_banner", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["post-login-banner"] = t
			}
		}
	}

	if v, ok := d.GetOk("failtime"); ok {
		if setArgNil {
			obj["failtime"] = nil
		} else {

			t, err := expandSystemGlobalFailtime(d, v, "failtime", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["failtime"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_lockout_threshold"); ok {
		if setArgNil {
			obj["admin-lockout-threshold"] = nil
		} else {

			t, err := expandSystemGlobalAdminLockoutThreshold(d, v, "admin_lockout_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-lockout-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcps_db_exp"); ok {
		if setArgNil {
			obj["dhcps-db-exp"] = nil
		} else {

			t, err := expandSystemGlobalDhcpsDbExp(d, v, "dhcps_db_exp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcps-db-exp"] = t
			}
		}
	}

	if v, ok := d.GetOk("n8021x_ca_certificate"); ok {
		if setArgNil {
			obj["802.1x-ca-certificate"] = nil
		} else {

			t, err := expandSystemGlobal8021XCaCertificate(d, v, "n8021x_ca_certificate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["802.1x-ca-certificate"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_remote_id"); ok {
		if setArgNil {
			obj["dhcp-remote-id"] = nil
		} else {

			t, err := expandSystemGlobalDhcpRemoteId(d, v, "dhcp_remote_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-remote-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_snoop_client_req"); ok {
		if setArgNil {
			obj["dhcp-snoop-client-req"] = nil
		} else {

			t, err := expandSystemGlobalDhcpSnoopClientReq(d, v, "dhcp_snoop_client_req", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-snoop-client-req"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("arp_inspection_monitor_timeout"); ok {
		if setArgNil {
			obj["arp-inspection-monitor-timeout"] = nil
		} else {

			t, err := expandSystemGlobalArpInspectionMonitorTimeout(d, v, "arp_inspection_monitor_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["arp-inspection-monitor-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_client_location"); ok {
		if setArgNil {
			obj["dhcp-client-location"] = nil
		} else {

			t, err := expandSystemGlobalDhcpClientLocation(d, v, "dhcp_client_location", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-client-location"] = t
			}
		}
	}

	if v, ok := d.GetOk("csr_ca_attribute"); ok {
		if setArgNil {
			obj["csr-ca-attribute"] = nil
		} else {

			t, err := expandSystemGlobalCsrCaAttribute(d, v, "csr_ca_attribute", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["csr-ca-attribute"] = t
			}
		}
	}

	if v, ok := d.GetOk("delaycli_timeout_cleanup"); ok {
		if setArgNil {
			obj["delaycli-timeout-cleanup"] = nil
		} else {

			t, err := expandSystemGlobalDelaycliTimeoutCleanup(d, v, "delaycli_timeout_cleanup", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["delaycli-timeout-cleanup"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("ipv6_accept_dad"); ok {
		if setArgNil {
			obj["ipv6-accept-dad"] = nil
		} else {

			t, err := expandSystemGlobalIpv6AcceptDad(d, v, "ipv6_accept_dad", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipv6-accept-dad"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_ssh_port"); ok {
		if setArgNil {
			obj["admin-ssh-port"] = nil
		} else {

			t, err := expandSystemGlobalAdminSshPort(d, v, "admin_ssh_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-ssh-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_server_cert"); ok {
		if setArgNil {
			obj["admin-server-cert"] = nil
		} else {

			t, err := expandSystemGlobalAdminServerCert(d, v, "admin_server_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-server-cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_isl"); ok {
		if setArgNil {
			obj["auto-isl"] = nil
		} else {

			t, err := expandSystemGlobalAutoIsl(d, v, "auto_isl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-isl"] = t
			}
		}
	}

	if v, ok := d.GetOk("language"); ok {
		if setArgNil {
			obj["language"] = nil
		} else {

			t, err := expandSystemGlobalLanguage(d, v, "language", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["language"] = t
			}
		}
	}

	if v, ok := d.GetOk("radius_coa_port"); ok {
		if setArgNil {
			obj["radius-coa-port"] = nil
		} else {

			t, err := expandSystemGlobalRadiusCoaPort(d, v, "radius_coa_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["radius-coa-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("dst"); ok {
		if setArgNil {
			obj["dst"] = nil
		} else {

			t, err := expandSystemGlobalDst(d, v, "dst", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dst"] = t
			}
		}
	}

	if v, ok := d.GetOk("interval"); ok {
		if setArgNil {
			obj["interval"] = nil
		} else {

			t, err := expandSystemGlobalInterval(d, v, "interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("cfg_save"); ok {
		if setArgNil {
			obj["cfg-save"] = nil
		} else {

			t, err := expandSystemGlobalCfgSave(d, v, "cfg_save", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cfg-save"] = t
			}
		}
	}

	if v, ok := d.GetOk("restart_time"); ok {
		if setArgNil {
			obj["restart-time"] = nil
		} else {

			t, err := expandSystemGlobalRestartTime(d, v, "restart_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["restart-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_option_format"); ok {
		if setArgNil {
			obj["dhcp-option-format"] = nil
		} else {

			t, err := expandSystemGlobalDhcpOptionFormat(d, v, "dhcp_option_format", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-option-format"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_scp"); ok {
		if setArgNil {
			obj["admin-scp"] = nil
		} else {

			t, err := expandSystemGlobalAdminScp(d, v, "admin_scp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-scp"] = t
			}
		}
	}

	if v, ok := d.GetOk("alertd_relog"); ok {
		if setArgNil {
			obj["alertd-relog"] = nil
		} else {

			t, err := expandSystemGlobalAlertdRelog(d, v, "alertd_relog", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["alertd-relog"] = t
			}
		}
	}

	if v, ok := d.GetOk("clt_cert_req"); ok {
		if setArgNil {
			obj["clt-cert-req"] = nil
		} else {

			t, err := expandSystemGlobalCltCertReq(d, v, "clt_cert_req", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["clt-cert-req"] = t
			}
		}
	}

	if v, ok := d.GetOk("daily_restart"); ok {
		if setArgNil {
			obj["daily-restart"] = nil
		} else {

			t, err := expandSystemGlobalDailyRestart(d, v, "daily_restart", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["daily-restart"] = t
			}
		}
	}

	if v, ok := d.GetOk("tftp"); ok {
		if setArgNil {
			obj["tftp"] = nil
		} else {

			t, err := expandSystemGlobalTftp(d, v, "tftp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tftp"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_ssh_grace_time"); ok {
		if setArgNil {
			obj["admin-ssh-grace-time"] = nil
		} else {

			t, err := expandSystemGlobalAdminSshGraceTime(d, v, "admin_ssh_grace_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-ssh-grace-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_https_pki_required"); ok {
		if setArgNil {
			obj["admin-https-pki-required"] = nil
		} else {

			t, err := expandSystemGlobalAdminHttpsPkiRequired(d, v, "admin_https_pki_required", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-https-pki-required"] = t
			}
		}
	}

	if v, ok := d.GetOk("radius_port"); ok {
		if setArgNil {
			obj["radius-port"] = nil
		} else {

			t, err := expandSystemGlobalRadiusPort(d, v, "radius_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["radius-port"] = t
			}
		}
	}

	return &obj, nil
}
