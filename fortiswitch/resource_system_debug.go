// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Application and CLI debug values to set at startup and retain over reboot.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDebug() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDebugUpdate,
		Read:   resourceSystemDebugRead,
		Update: resourceSystemDebugUpdate,
		Delete: resourceSystemDebugDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"cli": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 8),
				Optional:     true,
				Computed:     true,
			},
			"radvd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"raguard": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"miglogd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcp6c": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eap_proxy": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wpa_supp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"macsec_srv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcps": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fnbamd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcprelay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"snmpd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dnsproxy": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sflowd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcpc": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"router_launcher": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sshd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ctrld": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"stpd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trunkd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lacpd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lldpmedd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipconflictd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"httpsd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"link_monitor": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"libswitchd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"switch_launcher": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"alertd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"l2d": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"l2dbg": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nwmcfgd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nwmonitord": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"portspeedd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"l3": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mcast_snooping": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dmid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"scep": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cu_swtpd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fortilinkd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"flcmdd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gvrpd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"flan_mgr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rsyslogd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vrrpd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fpmd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ospfd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ospf6d": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pbrd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"isisd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ripd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ripngd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bgpd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"zebra": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bfdd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"staticd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pimd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ntpd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wiredap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip6addrd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gratarp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radius_das": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gui": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"statsd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"flow_export": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"erspan_auto_mgr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"autod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"email_server": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_script": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"apache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"delayclid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rvi_daemon": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"access_vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDebugUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemDebug(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDebug resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDebug(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemDebug resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemDebug")
	}

	return resourceSystemDebugRead(d, m)
}

func resourceSystemDebugDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemDebug(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemDebug resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDebug(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemDebug resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDebugRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemDebug(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemDebug resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDebug(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDebug resource from API: %v", err)
	}
	return nil
}

func flattenSystemDebugCli(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugRadvd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugRaguard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugMiglogd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugDhcp6C(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugEap_Proxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugWpa_Supp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugMacsec_Srv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugDhcps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugFnbamd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugDhcprelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugSnmpd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugDnsproxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugSflowd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugDhcpc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugRouterLauncher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugSshd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugCtrld(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugStpd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugTrunkd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugLacpd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugLldpmedd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugIpconflictd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugHttpsd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugLinkMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugLibswitchd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugSwitchLauncher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugAlertd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugL2D(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugL2Dbg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugNwmcfgd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugNwmonitord(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugPortspeedd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugL3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugMcastSnooping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugDmid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugScep(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugCu_Swtpd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugFortilinkd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugFlcmdd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugGvrpd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugFlanMgr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugRsyslogd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugVrrpd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugFpmd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugOspfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugOspf6D(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugPbrd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugIsisd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugRipd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugRipngd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugBgpd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugZebra(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugBfdd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugStaticd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugPimd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugNtpd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugWiredap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugIp6Addrd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugGratarp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugRadius_Das(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugGui(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugStatsd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugFlowExport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugErspanAutoMgr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugAutod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugEmailServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugAutoScript(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugApache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugDelayclid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugRviDaemon(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDebugAccessVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDebug(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("cli", flattenSystemDebugCli(o["cli"], d, "cli", sv)); err != nil {
		if !fortiAPIPatch(o["cli"]) {
			return fmt.Errorf("Error reading cli: %v", err)
		}
	}

	if err = d.Set("radvd", flattenSystemDebugRadvd(o["radvd"], d, "radvd", sv)); err != nil {
		if !fortiAPIPatch(o["radvd"]) {
			return fmt.Errorf("Error reading radvd: %v", err)
		}
	}

	if err = d.Set("raguard", flattenSystemDebugRaguard(o["raguard"], d, "raguard", sv)); err != nil {
		if !fortiAPIPatch(o["raguard"]) {
			return fmt.Errorf("Error reading raguard: %v", err)
		}
	}

	if err = d.Set("miglogd", flattenSystemDebugMiglogd(o["miglogd"], d, "miglogd", sv)); err != nil {
		if !fortiAPIPatch(o["miglogd"]) {
			return fmt.Errorf("Error reading miglogd: %v", err)
		}
	}

	if err = d.Set("dhcp6c", flattenSystemDebugDhcp6C(o["dhcp6c"], d, "dhcp6c", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp6c"]) {
			return fmt.Errorf("Error reading dhcp6c: %v", err)
		}
	}

	if err = d.Set("eap_proxy", flattenSystemDebugEap_Proxy(o["eap_proxy"], d, "eap_proxy", sv)); err != nil {
		if !fortiAPIPatch(o["eap_proxy"]) {
			return fmt.Errorf("Error reading eap_proxy: %v", err)
		}
	}

	if err = d.Set("wpa_supp", flattenSystemDebugWpa_Supp(o["wpa_supp"], d, "wpa_supp", sv)); err != nil {
		if !fortiAPIPatch(o["wpa_supp"]) {
			return fmt.Errorf("Error reading wpa_supp: %v", err)
		}
	}

	if err = d.Set("macsec_srv", flattenSystemDebugMacsec_Srv(o["macsec_srv"], d, "macsec_srv", sv)); err != nil {
		if !fortiAPIPatch(o["macsec_srv"]) {
			return fmt.Errorf("Error reading macsec_srv: %v", err)
		}
	}

	if err = d.Set("dhcps", flattenSystemDebugDhcps(o["dhcps"], d, "dhcps", sv)); err != nil {
		if !fortiAPIPatch(o["dhcps"]) {
			return fmt.Errorf("Error reading dhcps: %v", err)
		}
	}

	if err = d.Set("fnbamd", flattenSystemDebugFnbamd(o["fnbamd"], d, "fnbamd", sv)); err != nil {
		if !fortiAPIPatch(o["fnbamd"]) {
			return fmt.Errorf("Error reading fnbamd: %v", err)
		}
	}

	if err = d.Set("dhcprelay", flattenSystemDebugDhcprelay(o["dhcprelay"], d, "dhcprelay", sv)); err != nil {
		if !fortiAPIPatch(o["dhcprelay"]) {
			return fmt.Errorf("Error reading dhcprelay: %v", err)
		}
	}

	if err = d.Set("snmpd", flattenSystemDebugSnmpd(o["snmpd"], d, "snmpd", sv)); err != nil {
		if !fortiAPIPatch(o["snmpd"]) {
			return fmt.Errorf("Error reading snmpd: %v", err)
		}
	}

	if err = d.Set("dnsproxy", flattenSystemDebugDnsproxy(o["dnsproxy"], d, "dnsproxy", sv)); err != nil {
		if !fortiAPIPatch(o["dnsproxy"]) {
			return fmt.Errorf("Error reading dnsproxy: %v", err)
		}
	}

	if err = d.Set("sflowd", flattenSystemDebugSflowd(o["sflowd"], d, "sflowd", sv)); err != nil {
		if !fortiAPIPatch(o["sflowd"]) {
			return fmt.Errorf("Error reading sflowd: %v", err)
		}
	}

	if err = d.Set("dhcpc", flattenSystemDebugDhcpc(o["dhcpc"], d, "dhcpc", sv)); err != nil {
		if !fortiAPIPatch(o["dhcpc"]) {
			return fmt.Errorf("Error reading dhcpc: %v", err)
		}
	}

	if err = d.Set("router_launcher", flattenSystemDebugRouterLauncher(o["router-launcher"], d, "router_launcher", sv)); err != nil {
		if !fortiAPIPatch(o["router-launcher"]) {
			return fmt.Errorf("Error reading router_launcher: %v", err)
		}
	}

	if err = d.Set("sshd", flattenSystemDebugSshd(o["sshd"], d, "sshd", sv)); err != nil {
		if !fortiAPIPatch(o["sshd"]) {
			return fmt.Errorf("Error reading sshd: %v", err)
		}
	}

	if err = d.Set("ctrld", flattenSystemDebugCtrld(o["ctrld"], d, "ctrld", sv)); err != nil {
		if !fortiAPIPatch(o["ctrld"]) {
			return fmt.Errorf("Error reading ctrld: %v", err)
		}
	}

	if err = d.Set("stpd", flattenSystemDebugStpd(o["stpd"], d, "stpd", sv)); err != nil {
		if !fortiAPIPatch(o["stpd"]) {
			return fmt.Errorf("Error reading stpd: %v", err)
		}
	}

	if err = d.Set("trunkd", flattenSystemDebugTrunkd(o["trunkd"], d, "trunkd", sv)); err != nil {
		if !fortiAPIPatch(o["trunkd"]) {
			return fmt.Errorf("Error reading trunkd: %v", err)
		}
	}

	if err = d.Set("lacpd", flattenSystemDebugLacpd(o["lacpd"], d, "lacpd", sv)); err != nil {
		if !fortiAPIPatch(o["lacpd"]) {
			return fmt.Errorf("Error reading lacpd: %v", err)
		}
	}

	if err = d.Set("lldpmedd", flattenSystemDebugLldpmedd(o["lldpmedd"], d, "lldpmedd", sv)); err != nil {
		if !fortiAPIPatch(o["lldpmedd"]) {
			return fmt.Errorf("Error reading lldpmedd: %v", err)
		}
	}

	if err = d.Set("ipconflictd", flattenSystemDebugIpconflictd(o["ipconflictd"], d, "ipconflictd", sv)); err != nil {
		if !fortiAPIPatch(o["ipconflictd"]) {
			return fmt.Errorf("Error reading ipconflictd: %v", err)
		}
	}

	if err = d.Set("httpsd", flattenSystemDebugHttpsd(o["httpsd"], d, "httpsd", sv)); err != nil {
		if !fortiAPIPatch(o["httpsd"]) {
			return fmt.Errorf("Error reading httpsd: %v", err)
		}
	}

	if err = d.Set("link_monitor", flattenSystemDebugLinkMonitor(o["link-monitor"], d, "link_monitor", sv)); err != nil {
		if !fortiAPIPatch(o["link-monitor"]) {
			return fmt.Errorf("Error reading link_monitor: %v", err)
		}
	}

	if err = d.Set("libswitchd", flattenSystemDebugLibswitchd(o["libswitchd"], d, "libswitchd", sv)); err != nil {
		if !fortiAPIPatch(o["libswitchd"]) {
			return fmt.Errorf("Error reading libswitchd: %v", err)
		}
	}

	if err = d.Set("switch_launcher", flattenSystemDebugSwitchLauncher(o["switch-launcher"], d, "switch_launcher", sv)); err != nil {
		if !fortiAPIPatch(o["switch-launcher"]) {
			return fmt.Errorf("Error reading switch_launcher: %v", err)
		}
	}

	if err = d.Set("alertd", flattenSystemDebugAlertd(o["alertd"], d, "alertd", sv)); err != nil {
		if !fortiAPIPatch(o["alertd"]) {
			return fmt.Errorf("Error reading alertd: %v", err)
		}
	}

	if err = d.Set("l2d", flattenSystemDebugL2D(o["l2d"], d, "l2d", sv)); err != nil {
		if !fortiAPIPatch(o["l2d"]) {
			return fmt.Errorf("Error reading l2d: %v", err)
		}
	}

	if err = d.Set("l2dbg", flattenSystemDebugL2Dbg(o["l2dbg"], d, "l2dbg", sv)); err != nil {
		if !fortiAPIPatch(o["l2dbg"]) {
			return fmt.Errorf("Error reading l2dbg: %v", err)
		}
	}

	if err = d.Set("nwmcfgd", flattenSystemDebugNwmcfgd(o["nwmcfgd"], d, "nwmcfgd", sv)); err != nil {
		if !fortiAPIPatch(o["nwmcfgd"]) {
			return fmt.Errorf("Error reading nwmcfgd: %v", err)
		}
	}

	if err = d.Set("nwmonitord", flattenSystemDebugNwmonitord(o["nwmonitord"], d, "nwmonitord", sv)); err != nil {
		if !fortiAPIPatch(o["nwmonitord"]) {
			return fmt.Errorf("Error reading nwmonitord: %v", err)
		}
	}

	if err = d.Set("portspeedd", flattenSystemDebugPortspeedd(o["portspeedd"], d, "portspeedd", sv)); err != nil {
		if !fortiAPIPatch(o["portspeedd"]) {
			return fmt.Errorf("Error reading portspeedd: %v", err)
		}
	}

	if err = d.Set("l3", flattenSystemDebugL3(o["l3"], d, "l3", sv)); err != nil {
		if !fortiAPIPatch(o["l3"]) {
			return fmt.Errorf("Error reading l3: %v", err)
		}
	}

	if err = d.Set("mcast_snooping", flattenSystemDebugMcastSnooping(o["mcast-snooping"], d, "mcast_snooping", sv)); err != nil {
		if !fortiAPIPatch(o["mcast-snooping"]) {
			return fmt.Errorf("Error reading mcast_snooping: %v", err)
		}
	}

	if err = d.Set("dmid", flattenSystemDebugDmid(o["dmid"], d, "dmid", sv)); err != nil {
		if !fortiAPIPatch(o["dmid"]) {
			return fmt.Errorf("Error reading dmid: %v", err)
		}
	}

	if err = d.Set("scep", flattenSystemDebugScep(o["scep"], d, "scep", sv)); err != nil {
		if !fortiAPIPatch(o["scep"]) {
			return fmt.Errorf("Error reading scep: %v", err)
		}
	}

	if err = d.Set("cu_swtpd", flattenSystemDebugCu_Swtpd(o["cu_swtpd"], d, "cu_swtpd", sv)); err != nil {
		if !fortiAPIPatch(o["cu_swtpd"]) {
			return fmt.Errorf("Error reading cu_swtpd: %v", err)
		}
	}

	if err = d.Set("fortilinkd", flattenSystemDebugFortilinkd(o["fortilinkd"], d, "fortilinkd", sv)); err != nil {
		if !fortiAPIPatch(o["fortilinkd"]) {
			return fmt.Errorf("Error reading fortilinkd: %v", err)
		}
	}

	if err = d.Set("flcmdd", flattenSystemDebugFlcmdd(o["flcmdd"], d, "flcmdd", sv)); err != nil {
		if !fortiAPIPatch(o["flcmdd"]) {
			return fmt.Errorf("Error reading flcmdd: %v", err)
		}
	}

	if err = d.Set("gvrpd", flattenSystemDebugGvrpd(o["gvrpd"], d, "gvrpd", sv)); err != nil {
		if !fortiAPIPatch(o["gvrpd"]) {
			return fmt.Errorf("Error reading gvrpd: %v", err)
		}
	}

	if err = d.Set("flan_mgr", flattenSystemDebugFlanMgr(o["flan-mgr"], d, "flan_mgr", sv)); err != nil {
		if !fortiAPIPatch(o["flan-mgr"]) {
			return fmt.Errorf("Error reading flan_mgr: %v", err)
		}
	}

	if err = d.Set("rsyslogd", flattenSystemDebugRsyslogd(o["rsyslogd"], d, "rsyslogd", sv)); err != nil {
		if !fortiAPIPatch(o["rsyslogd"]) {
			return fmt.Errorf("Error reading rsyslogd: %v", err)
		}
	}

	if err = d.Set("vrrpd", flattenSystemDebugVrrpd(o["vrrpd"], d, "vrrpd", sv)); err != nil {
		if !fortiAPIPatch(o["vrrpd"]) {
			return fmt.Errorf("Error reading vrrpd: %v", err)
		}
	}

	if err = d.Set("fpmd", flattenSystemDebugFpmd(o["fpmd"], d, "fpmd", sv)); err != nil {
		if !fortiAPIPatch(o["fpmd"]) {
			return fmt.Errorf("Error reading fpmd: %v", err)
		}
	}

	if err = d.Set("ospfd", flattenSystemDebugOspfd(o["ospfd"], d, "ospfd", sv)); err != nil {
		if !fortiAPIPatch(o["ospfd"]) {
			return fmt.Errorf("Error reading ospfd: %v", err)
		}
	}

	if err = d.Set("ospf6d", flattenSystemDebugOspf6D(o["ospf6d"], d, "ospf6d", sv)); err != nil {
		if !fortiAPIPatch(o["ospf6d"]) {
			return fmt.Errorf("Error reading ospf6d: %v", err)
		}
	}

	if err = d.Set("pbrd", flattenSystemDebugPbrd(o["pbrd"], d, "pbrd", sv)); err != nil {
		if !fortiAPIPatch(o["pbrd"]) {
			return fmt.Errorf("Error reading pbrd: %v", err)
		}
	}

	if err = d.Set("isisd", flattenSystemDebugIsisd(o["isisd"], d, "isisd", sv)); err != nil {
		if !fortiAPIPatch(o["isisd"]) {
			return fmt.Errorf("Error reading isisd: %v", err)
		}
	}

	if err = d.Set("ripd", flattenSystemDebugRipd(o["ripd"], d, "ripd", sv)); err != nil {
		if !fortiAPIPatch(o["ripd"]) {
			return fmt.Errorf("Error reading ripd: %v", err)
		}
	}

	if err = d.Set("ripngd", flattenSystemDebugRipngd(o["ripngd"], d, "ripngd", sv)); err != nil {
		if !fortiAPIPatch(o["ripngd"]) {
			return fmt.Errorf("Error reading ripngd: %v", err)
		}
	}

	if err = d.Set("bgpd", flattenSystemDebugBgpd(o["bgpd"], d, "bgpd", sv)); err != nil {
		if !fortiAPIPatch(o["bgpd"]) {
			return fmt.Errorf("Error reading bgpd: %v", err)
		}
	}

	if err = d.Set("zebra", flattenSystemDebugZebra(o["zebra"], d, "zebra", sv)); err != nil {
		if !fortiAPIPatch(o["zebra"]) {
			return fmt.Errorf("Error reading zebra: %v", err)
		}
	}

	if err = d.Set("bfdd", flattenSystemDebugBfdd(o["bfdd"], d, "bfdd", sv)); err != nil {
		if !fortiAPIPatch(o["bfdd"]) {
			return fmt.Errorf("Error reading bfdd: %v", err)
		}
	}

	if err = d.Set("staticd", flattenSystemDebugStaticd(o["staticd"], d, "staticd", sv)); err != nil {
		if !fortiAPIPatch(o["staticd"]) {
			return fmt.Errorf("Error reading staticd: %v", err)
		}
	}

	if err = d.Set("pimd", flattenSystemDebugPimd(o["pimd"], d, "pimd", sv)); err != nil {
		if !fortiAPIPatch(o["pimd"]) {
			return fmt.Errorf("Error reading pimd: %v", err)
		}
	}

	if err = d.Set("ntpd", flattenSystemDebugNtpd(o["ntpd"], d, "ntpd", sv)); err != nil {
		if !fortiAPIPatch(o["ntpd"]) {
			return fmt.Errorf("Error reading ntpd: %v", err)
		}
	}

	if err = d.Set("wiredap", flattenSystemDebugWiredap(o["wiredap"], d, "wiredap", sv)); err != nil {
		if !fortiAPIPatch(o["wiredap"]) {
			return fmt.Errorf("Error reading wiredap: %v", err)
		}
	}

	if err = d.Set("ip6addrd", flattenSystemDebugIp6Addrd(o["ip6addrd"], d, "ip6addrd", sv)); err != nil {
		if !fortiAPIPatch(o["ip6addrd"]) {
			return fmt.Errorf("Error reading ip6addrd: %v", err)
		}
	}

	if err = d.Set("gratarp", flattenSystemDebugGratarp(o["gratarp"], d, "gratarp", sv)); err != nil {
		if !fortiAPIPatch(o["gratarp"]) {
			return fmt.Errorf("Error reading gratarp: %v", err)
		}
	}

	if err = d.Set("radius_das", flattenSystemDebugRadius_Das(o["radius_das"], d, "radius_das", sv)); err != nil {
		if !fortiAPIPatch(o["radius_das"]) {
			return fmt.Errorf("Error reading radius_das: %v", err)
		}
	}

	if err = d.Set("gui", flattenSystemDebugGui(o["gui"], d, "gui", sv)); err != nil {
		if !fortiAPIPatch(o["gui"]) {
			return fmt.Errorf("Error reading gui: %v", err)
		}
	}

	if err = d.Set("statsd", flattenSystemDebugStatsd(o["statsd"], d, "statsd", sv)); err != nil {
		if !fortiAPIPatch(o["statsd"]) {
			return fmt.Errorf("Error reading statsd: %v", err)
		}
	}

	if err = d.Set("flow_export", flattenSystemDebugFlowExport(o["flow-export"], d, "flow_export", sv)); err != nil {
		if !fortiAPIPatch(o["flow-export"]) {
			return fmt.Errorf("Error reading flow_export: %v", err)
		}
	}

	if err = d.Set("erspan_auto_mgr", flattenSystemDebugErspanAutoMgr(o["erspan-auto-mgr"], d, "erspan_auto_mgr", sv)); err != nil {
		if !fortiAPIPatch(o["erspan-auto-mgr"]) {
			return fmt.Errorf("Error reading erspan_auto_mgr: %v", err)
		}
	}

	if err = d.Set("autod", flattenSystemDebugAutod(o["autod"], d, "autod", sv)); err != nil {
		if !fortiAPIPatch(o["autod"]) {
			return fmt.Errorf("Error reading autod: %v", err)
		}
	}

	if err = d.Set("email_server", flattenSystemDebugEmailServer(o["email-server"], d, "email_server", sv)); err != nil {
		if !fortiAPIPatch(o["email-server"]) {
			return fmt.Errorf("Error reading email_server: %v", err)
		}
	}

	if err = d.Set("auto_script", flattenSystemDebugAutoScript(o["auto-script"], d, "auto_script", sv)); err != nil {
		if !fortiAPIPatch(o["auto-script"]) {
			return fmt.Errorf("Error reading auto_script: %v", err)
		}
	}

	if err = d.Set("apache", flattenSystemDebugApache(o["apache"], d, "apache", sv)); err != nil {
		if !fortiAPIPatch(o["apache"]) {
			return fmt.Errorf("Error reading apache: %v", err)
		}
	}

	if err = d.Set("delayclid", flattenSystemDebugDelayclid(o["delayclid"], d, "delayclid", sv)); err != nil {
		if !fortiAPIPatch(o["delayclid"]) {
			return fmt.Errorf("Error reading delayclid: %v", err)
		}
	}

	if err = d.Set("rvi_daemon", flattenSystemDebugRviDaemon(o["rvi-daemon"], d, "rvi_daemon", sv)); err != nil {
		if !fortiAPIPatch(o["rvi-daemon"]) {
			return fmt.Errorf("Error reading rvi_daemon: %v", err)
		}
	}

	if err = d.Set("access_vlan", flattenSystemDebugAccessVlan(o["access-vlan"], d, "access_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["access-vlan"]) {
			return fmt.Errorf("Error reading access_vlan: %v", err)
		}
	}

	return nil
}

func flattenSystemDebugFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemDebugCli(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugRadvd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugRaguard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugMiglogd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugDhcp6C(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugEap_Proxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugWpa_Supp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugMacsec_Srv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugDhcps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugFnbamd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugDhcprelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugSnmpd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugDnsproxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugSflowd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugDhcpc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugRouterLauncher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugSshd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugCtrld(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugStpd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugTrunkd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugLacpd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugLldpmedd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugIpconflictd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugHttpsd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugLinkMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugLibswitchd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugSwitchLauncher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugAlertd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugL2D(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugL2Dbg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugNwmcfgd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugNwmonitord(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugPortspeedd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugL3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugMcastSnooping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugDmid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugScep(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugCu_Swtpd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugFortilinkd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugFlcmdd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugGvrpd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugFlanMgr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugRsyslogd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugVrrpd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugFpmd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugOspfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugOspf6D(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugPbrd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugIsisd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugRipd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugRipngd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugBgpd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugZebra(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugBfdd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugStaticd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugPimd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugNtpd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugWiredap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugIp6Addrd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugGratarp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugRadius_Das(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugGui(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugStatsd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugFlowExport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugErspanAutoMgr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugAutod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugEmailServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugAutoScript(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugApache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugDelayclid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugRviDaemon(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDebugAccessVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDebug(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("cli"); ok {
		if setArgNil {
			obj["cli"] = nil
		} else {

			t, err := expandSystemDebugCli(d, v, "cli", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cli"] = t
			}
		}
	}

	if v, ok := d.GetOk("radvd"); ok {
		if setArgNil {
			obj["radvd"] = nil
		} else {

			t, err := expandSystemDebugRadvd(d, v, "radvd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["radvd"] = t
			}
		}
	}

	if v, ok := d.GetOk("raguard"); ok {
		if setArgNil {
			obj["raguard"] = nil
		} else {

			t, err := expandSystemDebugRaguard(d, v, "raguard", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["raguard"] = t
			}
		}
	}

	if v, ok := d.GetOk("miglogd"); ok {
		if setArgNil {
			obj["miglogd"] = nil
		} else {

			t, err := expandSystemDebugMiglogd(d, v, "miglogd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["miglogd"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp6c"); ok {
		if setArgNil {
			obj["dhcp6c"] = nil
		} else {

			t, err := expandSystemDebugDhcp6C(d, v, "dhcp6c", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp6c"] = t
			}
		}
	}

	if v, ok := d.GetOk("eap_proxy"); ok {
		if setArgNil {
			obj["eap_proxy"] = nil
		} else {

			t, err := expandSystemDebugEap_Proxy(d, v, "eap_proxy", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["eap_proxy"] = t
			}
		}
	}

	if v, ok := d.GetOk("wpa_supp"); ok {
		if setArgNil {
			obj["wpa_supp"] = nil
		} else {

			t, err := expandSystemDebugWpa_Supp(d, v, "wpa_supp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wpa_supp"] = t
			}
		}
	}

	if v, ok := d.GetOk("macsec_srv"); ok {
		if setArgNil {
			obj["macsec_srv"] = nil
		} else {

			t, err := expandSystemDebugMacsec_Srv(d, v, "macsec_srv", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["macsec_srv"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcps"); ok {
		if setArgNil {
			obj["dhcps"] = nil
		} else {

			t, err := expandSystemDebugDhcps(d, v, "dhcps", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcps"] = t
			}
		}
	}

	if v, ok := d.GetOk("fnbamd"); ok {
		if setArgNil {
			obj["fnbamd"] = nil
		} else {

			t, err := expandSystemDebugFnbamd(d, v, "fnbamd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fnbamd"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcprelay"); ok {
		if setArgNil {
			obj["dhcprelay"] = nil
		} else {

			t, err := expandSystemDebugDhcprelay(d, v, "dhcprelay", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcprelay"] = t
			}
		}
	}

	if v, ok := d.GetOk("snmpd"); ok {
		if setArgNil {
			obj["snmpd"] = nil
		} else {

			t, err := expandSystemDebugSnmpd(d, v, "snmpd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["snmpd"] = t
			}
		}
	}

	if v, ok := d.GetOk("dnsproxy"); ok {
		if setArgNil {
			obj["dnsproxy"] = nil
		} else {

			t, err := expandSystemDebugDnsproxy(d, v, "dnsproxy", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dnsproxy"] = t
			}
		}
	}

	if v, ok := d.GetOk("sflowd"); ok {
		if setArgNil {
			obj["sflowd"] = nil
		} else {

			t, err := expandSystemDebugSflowd(d, v, "sflowd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sflowd"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcpc"); ok {
		if setArgNil {
			obj["dhcpc"] = nil
		} else {

			t, err := expandSystemDebugDhcpc(d, v, "dhcpc", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcpc"] = t
			}
		}
	}

	if v, ok := d.GetOk("router_launcher"); ok {
		if setArgNil {
			obj["router-launcher"] = nil
		} else {

			t, err := expandSystemDebugRouterLauncher(d, v, "router_launcher", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["router-launcher"] = t
			}
		}
	}

	if v, ok := d.GetOk("sshd"); ok {
		if setArgNil {
			obj["sshd"] = nil
		} else {

			t, err := expandSystemDebugSshd(d, v, "sshd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sshd"] = t
			}
		}
	}

	if v, ok := d.GetOk("ctrld"); ok {
		if setArgNil {
			obj["ctrld"] = nil
		} else {

			t, err := expandSystemDebugCtrld(d, v, "ctrld", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ctrld"] = t
			}
		}
	}

	if v, ok := d.GetOk("stpd"); ok {
		if setArgNil {
			obj["stpd"] = nil
		} else {

			t, err := expandSystemDebugStpd(d, v, "stpd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["stpd"] = t
			}
		}
	}

	if v, ok := d.GetOk("trunkd"); ok {
		if setArgNil {
			obj["trunkd"] = nil
		} else {

			t, err := expandSystemDebugTrunkd(d, v, "trunkd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trunkd"] = t
			}
		}
	}

	if v, ok := d.GetOk("lacpd"); ok {
		if setArgNil {
			obj["lacpd"] = nil
		} else {

			t, err := expandSystemDebugLacpd(d, v, "lacpd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["lacpd"] = t
			}
		}
	}

	if v, ok := d.GetOk("lldpmedd"); ok {
		if setArgNil {
			obj["lldpmedd"] = nil
		} else {

			t, err := expandSystemDebugLldpmedd(d, v, "lldpmedd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["lldpmedd"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipconflictd"); ok {
		if setArgNil {
			obj["ipconflictd"] = nil
		} else {

			t, err := expandSystemDebugIpconflictd(d, v, "ipconflictd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipconflictd"] = t
			}
		}
	}

	if v, ok := d.GetOk("httpsd"); ok {
		if setArgNil {
			obj["httpsd"] = nil
		} else {

			t, err := expandSystemDebugHttpsd(d, v, "httpsd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["httpsd"] = t
			}
		}
	}

	if v, ok := d.GetOk("link_monitor"); ok {
		if setArgNil {
			obj["link-monitor"] = nil
		} else {

			t, err := expandSystemDebugLinkMonitor(d, v, "link_monitor", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["link-monitor"] = t
			}
		}
	}

	if v, ok := d.GetOk("libswitchd"); ok {
		if setArgNil {
			obj["libswitchd"] = nil
		} else {

			t, err := expandSystemDebugLibswitchd(d, v, "libswitchd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["libswitchd"] = t
			}
		}
	}

	if v, ok := d.GetOk("switch_launcher"); ok {
		if setArgNil {
			obj["switch-launcher"] = nil
		} else {

			t, err := expandSystemDebugSwitchLauncher(d, v, "switch_launcher", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["switch-launcher"] = t
			}
		}
	}

	if v, ok := d.GetOk("alertd"); ok {
		if setArgNil {
			obj["alertd"] = nil
		} else {

			t, err := expandSystemDebugAlertd(d, v, "alertd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["alertd"] = t
			}
		}
	}

	if v, ok := d.GetOk("l2d"); ok {
		if setArgNil {
			obj["l2d"] = nil
		} else {

			t, err := expandSystemDebugL2D(d, v, "l2d", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["l2d"] = t
			}
		}
	}

	if v, ok := d.GetOk("l2dbg"); ok {
		if setArgNil {
			obj["l2dbg"] = nil
		} else {

			t, err := expandSystemDebugL2Dbg(d, v, "l2dbg", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["l2dbg"] = t
			}
		}
	}

	if v, ok := d.GetOk("nwmcfgd"); ok {
		if setArgNil {
			obj["nwmcfgd"] = nil
		} else {

			t, err := expandSystemDebugNwmcfgd(d, v, "nwmcfgd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["nwmcfgd"] = t
			}
		}
	}

	if v, ok := d.GetOk("nwmonitord"); ok {
		if setArgNil {
			obj["nwmonitord"] = nil
		} else {

			t, err := expandSystemDebugNwmonitord(d, v, "nwmonitord", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["nwmonitord"] = t
			}
		}
	}

	if v, ok := d.GetOk("portspeedd"); ok {
		if setArgNil {
			obj["portspeedd"] = nil
		} else {

			t, err := expandSystemDebugPortspeedd(d, v, "portspeedd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["portspeedd"] = t
			}
		}
	}

	if v, ok := d.GetOk("l3"); ok {
		if setArgNil {
			obj["l3"] = nil
		} else {

			t, err := expandSystemDebugL3(d, v, "l3", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["l3"] = t
			}
		}
	}

	if v, ok := d.GetOk("mcast_snooping"); ok {
		if setArgNil {
			obj["mcast-snooping"] = nil
		} else {

			t, err := expandSystemDebugMcastSnooping(d, v, "mcast_snooping", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mcast-snooping"] = t
			}
		}
	}

	if v, ok := d.GetOk("dmid"); ok {
		if setArgNil {
			obj["dmid"] = nil
		} else {

			t, err := expandSystemDebugDmid(d, v, "dmid", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dmid"] = t
			}
		}
	}

	if v, ok := d.GetOk("scep"); ok {
		if setArgNil {
			obj["scep"] = nil
		} else {

			t, err := expandSystemDebugScep(d, v, "scep", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["scep"] = t
			}
		}
	}

	if v, ok := d.GetOk("cu_swtpd"); ok {
		if setArgNil {
			obj["cu_swtpd"] = nil
		} else {

			t, err := expandSystemDebugCu_Swtpd(d, v, "cu_swtpd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cu_swtpd"] = t
			}
		}
	}

	if v, ok := d.GetOk("fortilinkd"); ok {
		if setArgNil {
			obj["fortilinkd"] = nil
		} else {

			t, err := expandSystemDebugFortilinkd(d, v, "fortilinkd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fortilinkd"] = t
			}
		}
	}

	if v, ok := d.GetOk("flcmdd"); ok {
		if setArgNil {
			obj["flcmdd"] = nil
		} else {

			t, err := expandSystemDebugFlcmdd(d, v, "flcmdd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["flcmdd"] = t
			}
		}
	}

	if v, ok := d.GetOk("gvrpd"); ok {
		if setArgNil {
			obj["gvrpd"] = nil
		} else {

			t, err := expandSystemDebugGvrpd(d, v, "gvrpd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gvrpd"] = t
			}
		}
	}

	if v, ok := d.GetOk("flan_mgr"); ok {
		if setArgNil {
			obj["flan-mgr"] = nil
		} else {

			t, err := expandSystemDebugFlanMgr(d, v, "flan_mgr", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["flan-mgr"] = t
			}
		}
	}

	if v, ok := d.GetOk("rsyslogd"); ok {
		if setArgNil {
			obj["rsyslogd"] = nil
		} else {

			t, err := expandSystemDebugRsyslogd(d, v, "rsyslogd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rsyslogd"] = t
			}
		}
	}

	if v, ok := d.GetOk("vrrpd"); ok {
		if setArgNil {
			obj["vrrpd"] = nil
		} else {

			t, err := expandSystemDebugVrrpd(d, v, "vrrpd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vrrpd"] = t
			}
		}
	}

	if v, ok := d.GetOk("fpmd"); ok {
		if setArgNil {
			obj["fpmd"] = nil
		} else {

			t, err := expandSystemDebugFpmd(d, v, "fpmd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fpmd"] = t
			}
		}
	}

	if v, ok := d.GetOk("ospfd"); ok {
		if setArgNil {
			obj["ospfd"] = nil
		} else {

			t, err := expandSystemDebugOspfd(d, v, "ospfd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospfd"] = t
			}
		}
	}

	if v, ok := d.GetOk("ospf6d"); ok {
		if setArgNil {
			obj["ospf6d"] = nil
		} else {

			t, err := expandSystemDebugOspf6D(d, v, "ospf6d", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf6d"] = t
			}
		}
	}

	if v, ok := d.GetOk("pbrd"); ok {
		if setArgNil {
			obj["pbrd"] = nil
		} else {

			t, err := expandSystemDebugPbrd(d, v, "pbrd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pbrd"] = t
			}
		}
	}

	if v, ok := d.GetOk("isisd"); ok {
		if setArgNil {
			obj["isisd"] = nil
		} else {

			t, err := expandSystemDebugIsisd(d, v, "isisd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["isisd"] = t
			}
		}
	}

	if v, ok := d.GetOk("ripd"); ok {
		if setArgNil {
			obj["ripd"] = nil
		} else {

			t, err := expandSystemDebugRipd(d, v, "ripd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ripd"] = t
			}
		}
	}

	if v, ok := d.GetOk("ripngd"); ok {
		if setArgNil {
			obj["ripngd"] = nil
		} else {

			t, err := expandSystemDebugRipngd(d, v, "ripngd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ripngd"] = t
			}
		}
	}

	if v, ok := d.GetOk("bgpd"); ok {
		if setArgNil {
			obj["bgpd"] = nil
		} else {

			t, err := expandSystemDebugBgpd(d, v, "bgpd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bgpd"] = t
			}
		}
	}

	if v, ok := d.GetOk("zebra"); ok {
		if setArgNil {
			obj["zebra"] = nil
		} else {

			t, err := expandSystemDebugZebra(d, v, "zebra", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["zebra"] = t
			}
		}
	}

	if v, ok := d.GetOk("bfdd"); ok {
		if setArgNil {
			obj["bfdd"] = nil
		} else {

			t, err := expandSystemDebugBfdd(d, v, "bfdd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bfdd"] = t
			}
		}
	}

	if v, ok := d.GetOk("staticd"); ok {
		if setArgNil {
			obj["staticd"] = nil
		} else {

			t, err := expandSystemDebugStaticd(d, v, "staticd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["staticd"] = t
			}
		}
	}

	if v, ok := d.GetOk("pimd"); ok {
		if setArgNil {
			obj["pimd"] = nil
		} else {

			t, err := expandSystemDebugPimd(d, v, "pimd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pimd"] = t
			}
		}
	}

	if v, ok := d.GetOk("ntpd"); ok {
		if setArgNil {
			obj["ntpd"] = nil
		} else {

			t, err := expandSystemDebugNtpd(d, v, "ntpd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ntpd"] = t
			}
		}
	}

	if v, ok := d.GetOk("wiredap"); ok {
		if setArgNil {
			obj["wiredap"] = nil
		} else {

			t, err := expandSystemDebugWiredap(d, v, "wiredap", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wiredap"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip6addrd"); ok {
		if setArgNil {
			obj["ip6addrd"] = nil
		} else {

			t, err := expandSystemDebugIp6Addrd(d, v, "ip6addrd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip6addrd"] = t
			}
		}
	}

	if v, ok := d.GetOk("gratarp"); ok {
		if setArgNil {
			obj["gratarp"] = nil
		} else {

			t, err := expandSystemDebugGratarp(d, v, "gratarp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gratarp"] = t
			}
		}
	}

	if v, ok := d.GetOk("radius_das"); ok {
		if setArgNil {
			obj["radius_das"] = nil
		} else {

			t, err := expandSystemDebugRadius_Das(d, v, "radius_das", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["radius_das"] = t
			}
		}
	}

	if v, ok := d.GetOk("gui"); ok {
		if setArgNil {
			obj["gui"] = nil
		} else {

			t, err := expandSystemDebugGui(d, v, "gui", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gui"] = t
			}
		}
	}

	if v, ok := d.GetOk("statsd"); ok {
		if setArgNil {
			obj["statsd"] = nil
		} else {

			t, err := expandSystemDebugStatsd(d, v, "statsd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["statsd"] = t
			}
		}
	}

	if v, ok := d.GetOk("flow_export"); ok {
		if setArgNil {
			obj["flow-export"] = nil
		} else {

			t, err := expandSystemDebugFlowExport(d, v, "flow_export", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["flow-export"] = t
			}
		}
	}

	if v, ok := d.GetOk("erspan_auto_mgr"); ok {
		if setArgNil {
			obj["erspan-auto-mgr"] = nil
		} else {

			t, err := expandSystemDebugErspanAutoMgr(d, v, "erspan_auto_mgr", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["erspan-auto-mgr"] = t
			}
		}
	}

	if v, ok := d.GetOk("autod"); ok {
		if setArgNil {
			obj["autod"] = nil
		} else {

			t, err := expandSystemDebugAutod(d, v, "autod", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["autod"] = t
			}
		}
	}

	if v, ok := d.GetOk("email_server"); ok {
		if setArgNil {
			obj["email-server"] = nil
		} else {

			t, err := expandSystemDebugEmailServer(d, v, "email_server", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["email-server"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_script"); ok {
		if setArgNil {
			obj["auto-script"] = nil
		} else {

			t, err := expandSystemDebugAutoScript(d, v, "auto_script", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-script"] = t
			}
		}
	}

	if v, ok := d.GetOk("apache"); ok {
		if setArgNil {
			obj["apache"] = nil
		} else {

			t, err := expandSystemDebugApache(d, v, "apache", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["apache"] = t
			}
		}
	}

	if v, ok := d.GetOk("delayclid"); ok {
		if setArgNil {
			obj["delayclid"] = nil
		} else {

			t, err := expandSystemDebugDelayclid(d, v, "delayclid", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["delayclid"] = t
			}
		}
	}

	if v, ok := d.GetOk("rvi_daemon"); ok {
		if setArgNil {
			obj["rvi-daemon"] = nil
		} else {

			t, err := expandSystemDebugRviDaemon(d, v, "rvi_daemon", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rvi-daemon"] = t
			}
		}
	}

	if v, ok := d.GetOk("access_vlan"); ok {
		if setArgNil {
			obj["access-vlan"] = nil
		} else {

			t, err := expandSystemDebugAccessVlan(d, v, "access_vlan", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["access-vlan"] = t
			}
		}
	}

	return &obj, nil
}
