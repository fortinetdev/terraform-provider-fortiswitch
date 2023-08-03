// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Settings.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSettingsUpdate,
		Read:   resourceSystemSettingsRead,
		Update: resourceSystemSettingsUpdate,
		Delete: resourceSystemSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"utf8_spam_tagging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ecmp_max_paths": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sip_udp_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_ip_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpn_stats_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd_detect_mult": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 50),
				Optional:     true,
				Computed:     true,
			},
			"bfd_required_min_rx": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100000),
				Optional:     true,
				Computed:     true,
			},
			"wccp_cache_engine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_ecmp_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_ttl_notchange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"vpn_stats_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"opmode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd_desired_min_tx": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100000),
				Optional:     true,
				Computed:     true,
			},
			"sccp_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"asymroute6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"asymroute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_subnet_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"manageip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"sip_helper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip_tcp_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"multicast_skip_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip_nat_trace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd_dont_enforce_src_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strict_src_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSettings")
	}

	return resourceSystemSettingsRead(d, m)
}

func resourceSystemSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSettings resource from API: %v", err)
	}
	return nil
}

func flattenSystemSettingsUtf8SpamTagging(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsEcmpMaxPaths(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSipUdpPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsPerIpBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsVpnStatsLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsBfdDetectMult(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsBfdRequiredMinRx(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsWccpCacheEngine(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsIpEcmpMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsMulticastTtlNotchange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsMulticastForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsVpnStatsPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsOpmode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsBfdDesiredMinTx(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSccpPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsAsymroute6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsAsymroute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsAllowSubnetOverlap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsManageip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSipHelper(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSipTcpPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsMulticastSkipPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSipNatTrace(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsBfdDontEnforceSrcPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsStrictSrcCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("utf8_spam_tagging", flattenSystemSettingsUtf8SpamTagging(o["utf8-spam-tagging"], d, "utf8_spam_tagging", sv)); err != nil {
		if !fortiAPIPatch(o["utf8-spam-tagging"]) {
			return fmt.Errorf("Error reading utf8_spam_tagging: %v", err)
		}
	}

	if err = d.Set("ecmp_max_paths", flattenSystemSettingsEcmpMaxPaths(o["ecmp-max-paths"], d, "ecmp_max_paths", sv)); err != nil {
		if !fortiAPIPatch(o["ecmp-max-paths"]) {
			return fmt.Errorf("Error reading ecmp_max_paths: %v", err)
		}
	}

	if err = d.Set("sip_udp_port", flattenSystemSettingsSipUdpPort(o["sip-udp-port"], d, "sip_udp_port", sv)); err != nil {
		if !fortiAPIPatch(o["sip-udp-port"]) {
			return fmt.Errorf("Error reading sip_udp_port: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemSettingsIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("per_ip_bandwidth", flattenSystemSettingsPerIpBandwidth(o["per-ip-bandwidth"], d, "per_ip_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["per-ip-bandwidth"]) {
			return fmt.Errorf("Error reading per_ip_bandwidth: %v", err)
		}
	}

	if err = d.Set("vpn_stats_log", flattenSystemSettingsVpnStatsLog(o["vpn-stats-log"], d, "vpn_stats_log", sv)); err != nil {
		if !fortiAPIPatch(o["vpn-stats-log"]) {
			return fmt.Errorf("Error reading vpn_stats_log: %v", err)
		}
	}

	if err = d.Set("bfd_detect_mult", flattenSystemSettingsBfdDetectMult(o["bfd-detect-mult"], d, "bfd_detect_mult", sv)); err != nil {
		if !fortiAPIPatch(o["bfd-detect-mult"]) {
			return fmt.Errorf("Error reading bfd_detect_mult: %v", err)
		}
	}

	if err = d.Set("bfd_required_min_rx", flattenSystemSettingsBfdRequiredMinRx(o["bfd-required-min-rx"], d, "bfd_required_min_rx", sv)); err != nil {
		if !fortiAPIPatch(o["bfd-required-min-rx"]) {
			return fmt.Errorf("Error reading bfd_required_min_rx: %v", err)
		}
	}

	if err = d.Set("wccp_cache_engine", flattenSystemSettingsWccpCacheEngine(o["wccp-cache-engine"], d, "wccp_cache_engine", sv)); err != nil {
		if !fortiAPIPatch(o["wccp-cache-engine"]) {
			return fmt.Errorf("Error reading wccp_cache_engine: %v", err)
		}
	}

	if err = d.Set("ip_ecmp_mode", flattenSystemSettingsIpEcmpMode(o["ip-ecmp-mode"], d, "ip_ecmp_mode", sv)); err != nil {
		if !fortiAPIPatch(o["ip-ecmp-mode"]) {
			return fmt.Errorf("Error reading ip_ecmp_mode: %v", err)
		}
	}

	if err = d.Set("multicast_ttl_notchange", flattenSystemSettingsMulticastTtlNotchange(o["multicast-ttl-notchange"], d, "multicast_ttl_notchange", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-ttl-notchange"]) {
			return fmt.Errorf("Error reading multicast_ttl_notchange: %v", err)
		}
	}

	if err = d.Set("gateway", flattenSystemSettingsGateway(o["gateway"], d, "gateway", sv)); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("multicast_forward", flattenSystemSettingsMulticastForward(o["multicast-forward"], d, "multicast_forward", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-forward"]) {
			return fmt.Errorf("Error reading multicast_forward: %v", err)
		}
	}

	if err = d.Set("bfd", flattenSystemSettingsBfd(o["bfd"], d, "bfd", sv)); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("comments", flattenSystemSettingsComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("vpn_stats_period", flattenSystemSettingsVpnStatsPeriod(o["vpn-stats-period"], d, "vpn_stats_period", sv)); err != nil {
		if !fortiAPIPatch(o["vpn-stats-period"]) {
			return fmt.Errorf("Error reading vpn_stats_period: %v", err)
		}
	}

	if err = d.Set("opmode", flattenSystemSettingsOpmode(o["opmode"], d, "opmode", sv)); err != nil {
		if !fortiAPIPatch(o["opmode"]) {
			return fmt.Errorf("Error reading opmode: %v", err)
		}
	}

	if err = d.Set("bfd_desired_min_tx", flattenSystemSettingsBfdDesiredMinTx(o["bfd-desired-min-tx"], d, "bfd_desired_min_tx", sv)); err != nil {
		if !fortiAPIPatch(o["bfd-desired-min-tx"]) {
			return fmt.Errorf("Error reading bfd_desired_min_tx: %v", err)
		}
	}

	if err = d.Set("sccp_port", flattenSystemSettingsSccpPort(o["sccp-port"], d, "sccp_port", sv)); err != nil {
		if !fortiAPIPatch(o["sccp-port"]) {
			return fmt.Errorf("Error reading sccp_port: %v", err)
		}
	}

	if err = d.Set("asymroute6", flattenSystemSettingsAsymroute6(o["asymroute6"], d, "asymroute6", sv)); err != nil {
		if !fortiAPIPatch(o["asymroute6"]) {
			return fmt.Errorf("Error reading asymroute6: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSettingsStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("asymroute", flattenSystemSettingsAsymroute(o["asymroute"], d, "asymroute", sv)); err != nil {
		if !fortiAPIPatch(o["asymroute"]) {
			return fmt.Errorf("Error reading asymroute: %v", err)
		}
	}

	if err = d.Set("allow_subnet_overlap", flattenSystemSettingsAllowSubnetOverlap(o["allow-subnet-overlap"], d, "allow_subnet_overlap", sv)); err != nil {
		if !fortiAPIPatch(o["allow-subnet-overlap"]) {
			return fmt.Errorf("Error reading allow_subnet_overlap: %v", err)
		}
	}

	if err = d.Set("manageip", flattenSystemSettingsManageip(o["manageip"], d, "manageip", sv)); err != nil {
		if !fortiAPIPatch(o["manageip"]) {
			return fmt.Errorf("Error reading manageip: %v", err)
		}
	}

	if err = d.Set("device", flattenSystemSettingsDevice(o["device"], d, "device", sv)); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("sip_helper", flattenSystemSettingsSipHelper(o["sip-helper"], d, "sip_helper", sv)); err != nil {
		if !fortiAPIPatch(o["sip-helper"]) {
			return fmt.Errorf("Error reading sip_helper: %v", err)
		}
	}

	if err = d.Set("sip_tcp_port", flattenSystemSettingsSipTcpPort(o["sip-tcp-port"], d, "sip_tcp_port", sv)); err != nil {
		if !fortiAPIPatch(o["sip-tcp-port"]) {
			return fmt.Errorf("Error reading sip_tcp_port: %v", err)
		}
	}

	if err = d.Set("multicast_skip_policy", flattenSystemSettingsMulticastSkipPolicy(o["multicast-skip-policy"], d, "multicast_skip_policy", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-skip-policy"]) {
			return fmt.Errorf("Error reading multicast_skip_policy: %v", err)
		}
	}

	if err = d.Set("sip_nat_trace", flattenSystemSettingsSipNatTrace(o["sip-nat-trace"], d, "sip_nat_trace", sv)); err != nil {
		if !fortiAPIPatch(o["sip-nat-trace"]) {
			return fmt.Errorf("Error reading sip_nat_trace: %v", err)
		}
	}

	if err = d.Set("bfd_dont_enforce_src_port", flattenSystemSettingsBfdDontEnforceSrcPort(o["bfd-dont-enforce-src-port"], d, "bfd_dont_enforce_src_port", sv)); err != nil {
		if !fortiAPIPatch(o["bfd-dont-enforce-src-port"]) {
			return fmt.Errorf("Error reading bfd_dont_enforce_src_port: %v", err)
		}
	}

	if err = d.Set("strict_src_check", flattenSystemSettingsStrictSrcCheck(o["strict-src-check"], d, "strict_src_check", sv)); err != nil {
		if !fortiAPIPatch(o["strict-src-check"]) {
			return fmt.Errorf("Error reading strict_src_check: %v", err)
		}
	}

	return nil
}

func flattenSystemSettingsFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemSettingsUtf8SpamTagging(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsEcmpMaxPaths(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSipUdpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsPerIpBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsVpnStatsLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfdDetectMult(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfdRequiredMinRx(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsWccpCacheEngine(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIpEcmpMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsMulticastTtlNotchange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsMulticastForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsVpnStatsPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsOpmode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfdDesiredMinTx(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSccpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAsymroute6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAsymroute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAllowSubnetOverlap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsManageip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSipHelper(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSipTcpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsMulticastSkipPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSipNatTrace(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfdDontEnforceSrcPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsStrictSrcCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("utf8_spam_tagging"); ok {
		if setArgNil {
			obj["utf8-spam-tagging"] = nil
		} else {

			t, err := expandSystemSettingsUtf8SpamTagging(d, v, "utf8_spam_tagging", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["utf8-spam-tagging"] = t
			}
		}
	}

	if v, ok := d.GetOk("ecmp_max_paths"); ok {
		if setArgNil {
			obj["ecmp-max-paths"] = nil
		} else {

			t, err := expandSystemSettingsEcmpMaxPaths(d, v, "ecmp_max_paths", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ecmp-max-paths"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("sip_udp_port"); ok {
		if setArgNil {
			obj["sip-udp-port"] = nil
		} else {

			t, err := expandSystemSettingsSipUdpPort(d, v, "sip_udp_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sip-udp-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		if setArgNil {
			obj["ip"] = nil
		} else {

			t, err := expandSystemSettingsIp(d, v, "ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("per_ip_bandwidth"); ok {
		if setArgNil {
			obj["per-ip-bandwidth"] = nil
		} else {

			t, err := expandSystemSettingsPerIpBandwidth(d, v, "per_ip_bandwidth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["per-ip-bandwidth"] = t
			}
		}
	}

	if v, ok := d.GetOk("vpn_stats_log"); ok {
		if setArgNil {
			obj["vpn-stats-log"] = nil
		} else {

			t, err := expandSystemSettingsVpnStatsLog(d, v, "vpn_stats_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vpn-stats-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("bfd_detect_mult"); ok {
		if setArgNil {
			obj["bfd-detect-mult"] = nil
		} else {

			t, err := expandSystemSettingsBfdDetectMult(d, v, "bfd_detect_mult", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bfd-detect-mult"] = t
			}
		}
	}

	if v, ok := d.GetOk("bfd_required_min_rx"); ok {
		if setArgNil {
			obj["bfd-required-min-rx"] = nil
		} else {

			t, err := expandSystemSettingsBfdRequiredMinRx(d, v, "bfd_required_min_rx", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bfd-required-min-rx"] = t
			}
		}
	}

	if v, ok := d.GetOk("wccp_cache_engine"); ok {
		if setArgNil {
			obj["wccp-cache-engine"] = nil
		} else {

			t, err := expandSystemSettingsWccpCacheEngine(d, v, "wccp_cache_engine", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wccp-cache-engine"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip_ecmp_mode"); ok {
		if setArgNil {
			obj["ip-ecmp-mode"] = nil
		} else {

			t, err := expandSystemSettingsIpEcmpMode(d, v, "ip_ecmp_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip-ecmp-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("multicast_ttl_notchange"); ok {
		if setArgNil {
			obj["multicast-ttl-notchange"] = nil
		} else {

			t, err := expandSystemSettingsMulticastTtlNotchange(d, v, "multicast_ttl_notchange", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["multicast-ttl-notchange"] = t
			}
		}
	}

	if v, ok := d.GetOk("gateway"); ok {
		if setArgNil {
			obj["gateway"] = nil
		} else {

			t, err := expandSystemSettingsGateway(d, v, "gateway", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gateway"] = t
			}
		}
	}

	if v, ok := d.GetOk("multicast_forward"); ok {
		if setArgNil {
			obj["multicast-forward"] = nil
		} else {

			t, err := expandSystemSettingsMulticastForward(d, v, "multicast_forward", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["multicast-forward"] = t
			}
		}
	}

	if v, ok := d.GetOk("bfd"); ok {
		if setArgNil {
			obj["bfd"] = nil
		} else {

			t, err := expandSystemSettingsBfd(d, v, "bfd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bfd"] = t
			}
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		if setArgNil {
			obj["comments"] = nil
		} else {

			t, err := expandSystemSettingsComments(d, v, "comments", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["comments"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("vpn_stats_period"); ok {
		if setArgNil {
			obj["vpn-stats-period"] = nil
		} else {

			t, err := expandSystemSettingsVpnStatsPeriod(d, v, "vpn_stats_period", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vpn-stats-period"] = t
			}
		}
	}

	if v, ok := d.GetOk("opmode"); ok {
		if setArgNil {
			obj["opmode"] = nil
		} else {

			t, err := expandSystemSettingsOpmode(d, v, "opmode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["opmode"] = t
			}
		}
	}

	if v, ok := d.GetOk("bfd_desired_min_tx"); ok {
		if setArgNil {
			obj["bfd-desired-min-tx"] = nil
		} else {

			t, err := expandSystemSettingsBfdDesiredMinTx(d, v, "bfd_desired_min_tx", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bfd-desired-min-tx"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("sccp_port"); ok {
		if setArgNil {
			obj["sccp-port"] = nil
		} else {

			t, err := expandSystemSettingsSccpPort(d, v, "sccp_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sccp-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("asymroute6"); ok {
		if setArgNil {
			obj["asymroute6"] = nil
		} else {

			t, err := expandSystemSettingsAsymroute6(d, v, "asymroute6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["asymroute6"] = t
			}
		}
	}

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandSystemSettingsStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("asymroute"); ok {
		if setArgNil {
			obj["asymroute"] = nil
		} else {

			t, err := expandSystemSettingsAsymroute(d, v, "asymroute", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["asymroute"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_subnet_overlap"); ok {
		if setArgNil {
			obj["allow-subnet-overlap"] = nil
		} else {

			t, err := expandSystemSettingsAllowSubnetOverlap(d, v, "allow_subnet_overlap", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-subnet-overlap"] = t
			}
		}
	}

	if v, ok := d.GetOk("manageip"); ok {
		if setArgNil {
			obj["manageip"] = nil
		} else {

			t, err := expandSystemSettingsManageip(d, v, "manageip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["manageip"] = t
			}
		}
	}

	if v, ok := d.GetOk("device"); ok {
		if setArgNil {
			obj["device"] = nil
		} else {

			t, err := expandSystemSettingsDevice(d, v, "device", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["device"] = t
			}
		}
	}

	if v, ok := d.GetOk("sip_helper"); ok {
		if setArgNil {
			obj["sip-helper"] = nil
		} else {

			t, err := expandSystemSettingsSipHelper(d, v, "sip_helper", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sip-helper"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("sip_tcp_port"); ok {
		if setArgNil {
			obj["sip-tcp-port"] = nil
		} else {

			t, err := expandSystemSettingsSipTcpPort(d, v, "sip_tcp_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sip-tcp-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("multicast_skip_policy"); ok {
		if setArgNil {
			obj["multicast-skip-policy"] = nil
		} else {

			t, err := expandSystemSettingsMulticastSkipPolicy(d, v, "multicast_skip_policy", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["multicast-skip-policy"] = t
			}
		}
	}

	if v, ok := d.GetOk("sip_nat_trace"); ok {
		if setArgNil {
			obj["sip-nat-trace"] = nil
		} else {

			t, err := expandSystemSettingsSipNatTrace(d, v, "sip_nat_trace", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sip-nat-trace"] = t
			}
		}
	}

	if v, ok := d.GetOk("bfd_dont_enforce_src_port"); ok {
		if setArgNil {
			obj["bfd-dont-enforce-src-port"] = nil
		} else {

			t, err := expandSystemSettingsBfdDontEnforceSrcPort(d, v, "bfd_dont_enforce_src_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bfd-dont-enforce-src-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("strict_src_check"); ok {
		if setArgNil {
			obj["strict-src-check"] = nil
		} else {

			t, err := expandSystemSettingsStrictSrcCheck(d, v, "strict_src_check", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["strict-src-check"] = t
			}
		}
	}

	return &obj, nil
}
