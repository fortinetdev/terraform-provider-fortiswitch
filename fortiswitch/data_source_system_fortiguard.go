// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure FortiGuard services.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemFortiguard() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemFortiguardRead,
		Schema: map[string]*schema.Schema{

			"webfilter_license": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"webfilter_cache": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"avquery_license": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"antispam_cache_mpercent": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"webfilter_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"antispam_cache": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"webfilter_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"load_balance_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"antispam_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"avquery_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"webfilter_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"avquery_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"client_override_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"antispam_score_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"avquery_cache": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"avquery_cache_mpercent": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"avquery_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"avquery_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"antispam_license": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"antispam_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"analysis_service": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_override_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"srv_ovrd": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"webfilter_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"service_account_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"antispam_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"antispam_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"srv_ovrd_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemFortiguardRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemFortiguard"

	o, err := c.ReadSystemFortiguard(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemFortiguard: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemFortiguard(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemFortiguard from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemFortiguardWebfilterLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardWebfilterCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAvqueryLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamCacheMpercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardWebfilterTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardWebfilterExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardLoadBalanceServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAvqueryForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardWebfilterForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAvqueryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardClientOverrideStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamScoreThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAvqueryCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAvqueryCacheMpercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAvqueryCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAvqueryExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAnalysisService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardClientOverrideIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardSrvOvrd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardWebfilterCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardServiceAccountId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardSrvOvrdList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenSystemFortiguardSrvOvrdListIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			tmp["ip6"] = dataSourceFlattenSystemFortiguardSrvOvrdListIp6(i["ip6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {
			tmp["addr_type"] = dataSourceFlattenSystemFortiguardSrvOvrdListAddrType(i["addr-type"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemFortiguardSrvOvrdListIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardSrvOvrdListIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardSrvOvrdListAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemFortiguard(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("webfilter_license", dataSourceFlattenSystemFortiguardWebfilterLicense(o["webfilter-license"], d, "webfilter_license")); err != nil {
		if !fortiAPIPatch(o["webfilter-license"]) {
			return fmt.Errorf("Error reading webfilter_license: %v", err)
		}
	}

	if err = d.Set("webfilter_cache", dataSourceFlattenSystemFortiguardWebfilterCache(o["webfilter-cache"], d, "webfilter_cache")); err != nil {
		if !fortiAPIPatch(o["webfilter-cache"]) {
			return fmt.Errorf("Error reading webfilter_cache: %v", err)
		}
	}

	if err = d.Set("avquery_license", dataSourceFlattenSystemFortiguardAvqueryLicense(o["avquery-license"], d, "avquery_license")); err != nil {
		if !fortiAPIPatch(o["avquery-license"]) {
			return fmt.Errorf("Error reading avquery_license: %v", err)
		}
	}

	if err = d.Set("antispam_cache_mpercent", dataSourceFlattenSystemFortiguardAntispamCacheMpercent(o["antispam-cache-mpercent"], d, "antispam_cache_mpercent")); err != nil {
		if !fortiAPIPatch(o["antispam-cache-mpercent"]) {
			return fmt.Errorf("Error reading antispam_cache_mpercent: %v", err)
		}
	}

	if err = d.Set("webfilter_timeout", dataSourceFlattenSystemFortiguardWebfilterTimeout(o["webfilter-timeout"], d, "webfilter_timeout")); err != nil {
		if !fortiAPIPatch(o["webfilter-timeout"]) {
			return fmt.Errorf("Error reading webfilter_timeout: %v", err)
		}
	}

	if err = d.Set("antispam_cache", dataSourceFlattenSystemFortiguardAntispamCache(o["antispam-cache"], d, "antispam_cache")); err != nil {
		if !fortiAPIPatch(o["antispam-cache"]) {
			return fmt.Errorf("Error reading antispam_cache: %v", err)
		}
	}

	if err = d.Set("webfilter_expiration", dataSourceFlattenSystemFortiguardWebfilterExpiration(o["webfilter-expiration"], d, "webfilter_expiration")); err != nil {
		if !fortiAPIPatch(o["webfilter-expiration"]) {
			return fmt.Errorf("Error reading webfilter_expiration: %v", err)
		}
	}

	if err = d.Set("load_balance_servers", dataSourceFlattenSystemFortiguardLoadBalanceServers(o["load-balance-servers"], d, "load_balance_servers")); err != nil {
		if !fortiAPIPatch(o["load-balance-servers"]) {
			return fmt.Errorf("Error reading load_balance_servers: %v", err)
		}
	}

	if err = d.Set("antispam_force_off", dataSourceFlattenSystemFortiguardAntispamForceOff(o["antispam-force-off"], d, "antispam_force_off")); err != nil {
		if !fortiAPIPatch(o["antispam-force-off"]) {
			return fmt.Errorf("Error reading antispam_force_off: %v", err)
		}
	}

	if err = d.Set("avquery_force_off", dataSourceFlattenSystemFortiguardAvqueryForceOff(o["avquery-force-off"], d, "avquery_force_off")); err != nil {
		if !fortiAPIPatch(o["avquery-force-off"]) {
			return fmt.Errorf("Error reading avquery_force_off: %v", err)
		}
	}

	if err = d.Set("webfilter_force_off", dataSourceFlattenSystemFortiguardWebfilterForceOff(o["webfilter-force-off"], d, "webfilter_force_off")); err != nil {
		if !fortiAPIPatch(o["webfilter-force-off"]) {
			return fmt.Errorf("Error reading webfilter_force_off: %v", err)
		}
	}

	if err = d.Set("avquery_timeout", dataSourceFlattenSystemFortiguardAvqueryTimeout(o["avquery-timeout"], d, "avquery_timeout")); err != nil {
		if !fortiAPIPatch(o["avquery-timeout"]) {
			return fmt.Errorf("Error reading avquery_timeout: %v", err)
		}
	}

	if err = d.Set("client_override_status", dataSourceFlattenSystemFortiguardClientOverrideStatus(o["client-override-status"], d, "client_override_status")); err != nil {
		if !fortiAPIPatch(o["client-override-status"]) {
			return fmt.Errorf("Error reading client_override_status: %v", err)
		}
	}

	if err = d.Set("antispam_score_threshold", dataSourceFlattenSystemFortiguardAntispamScoreThreshold(o["antispam-score-threshold"], d, "antispam_score_threshold")); err != nil {
		if !fortiAPIPatch(o["antispam-score-threshold"]) {
			return fmt.Errorf("Error reading antispam_score_threshold: %v", err)
		}
	}

	if err = d.Set("avquery_cache", dataSourceFlattenSystemFortiguardAvqueryCache(o["avquery-cache"], d, "avquery_cache")); err != nil {
		if !fortiAPIPatch(o["avquery-cache"]) {
			return fmt.Errorf("Error reading avquery_cache: %v", err)
		}
	}

	if err = d.Set("avquery_cache_mpercent", dataSourceFlattenSystemFortiguardAvqueryCacheMpercent(o["avquery-cache-mpercent"], d, "avquery_cache_mpercent")); err != nil {
		if !fortiAPIPatch(o["avquery-cache-mpercent"]) {
			return fmt.Errorf("Error reading avquery_cache_mpercent: %v", err)
		}
	}

	if err = d.Set("avquery_cache_ttl", dataSourceFlattenSystemFortiguardAvqueryCacheTtl(o["avquery-cache-ttl"], d, "avquery_cache_ttl")); err != nil {
		if !fortiAPIPatch(o["avquery-cache-ttl"]) {
			return fmt.Errorf("Error reading avquery_cache_ttl: %v", err)
		}
	}

	if err = d.Set("avquery_expiration", dataSourceFlattenSystemFortiguardAvqueryExpiration(o["avquery-expiration"], d, "avquery_expiration")); err != nil {
		if !fortiAPIPatch(o["avquery-expiration"]) {
			return fmt.Errorf("Error reading avquery_expiration: %v", err)
		}
	}

	if err = d.Set("hostname", dataSourceFlattenSystemFortiguardHostname(o["hostname"], d, "hostname")); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("antispam_license", dataSourceFlattenSystemFortiguardAntispamLicense(o["antispam-license"], d, "antispam_license")); err != nil {
		if !fortiAPIPatch(o["antispam-license"]) {
			return fmt.Errorf("Error reading antispam_license: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenSystemFortiguardPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("antispam_cache_ttl", dataSourceFlattenSystemFortiguardAntispamCacheTtl(o["antispam-cache-ttl"], d, "antispam_cache_ttl")); err != nil {
		if !fortiAPIPatch(o["antispam-cache-ttl"]) {
			return fmt.Errorf("Error reading antispam_cache_ttl: %v", err)
		}
	}

	if err = d.Set("analysis_service", dataSourceFlattenSystemFortiguardAnalysisService(o["analysis-service"], d, "analysis_service")); err != nil {
		if !fortiAPIPatch(o["analysis-service"]) {
			return fmt.Errorf("Error reading analysis_service: %v", err)
		}
	}

	if err = d.Set("client_override_ip", dataSourceFlattenSystemFortiguardClientOverrideIp(o["client-override-ip"], d, "client_override_ip")); err != nil {
		if !fortiAPIPatch(o["client-override-ip"]) {
			return fmt.Errorf("Error reading client_override_ip: %v", err)
		}
	}

	if err = d.Set("srv_ovrd", dataSourceFlattenSystemFortiguardSrvOvrd(o["srv-ovrd"], d, "srv_ovrd")); err != nil {
		if !fortiAPIPatch(o["srv-ovrd"]) {
			return fmt.Errorf("Error reading srv_ovrd: %v", err)
		}
	}

	if err = d.Set("webfilter_cache_ttl", dataSourceFlattenSystemFortiguardWebfilterCacheTtl(o["webfilter-cache-ttl"], d, "webfilter_cache_ttl")); err != nil {
		if !fortiAPIPatch(o["webfilter-cache-ttl"]) {
			return fmt.Errorf("Error reading webfilter_cache_ttl: %v", err)
		}
	}

	if err = d.Set("service_account_id", dataSourceFlattenSystemFortiguardServiceAccountId(o["service-account-id"], d, "service_account_id")); err != nil {
		if !fortiAPIPatch(o["service-account-id"]) {
			return fmt.Errorf("Error reading service_account_id: %v", err)
		}
	}

	if err = d.Set("antispam_expiration", dataSourceFlattenSystemFortiguardAntispamExpiration(o["antispam-expiration"], d, "antispam_expiration")); err != nil {
		if !fortiAPIPatch(o["antispam-expiration"]) {
			return fmt.Errorf("Error reading antispam_expiration: %v", err)
		}
	}

	if err = d.Set("antispam_timeout", dataSourceFlattenSystemFortiguardAntispamTimeout(o["antispam-timeout"], d, "antispam_timeout")); err != nil {
		if !fortiAPIPatch(o["antispam-timeout"]) {
			return fmt.Errorf("Error reading antispam_timeout: %v", err)
		}
	}

	if err = d.Set("srv_ovrd_list", dataSourceFlattenSystemFortiguardSrvOvrdList(o["srv-ovrd-list"], d, "srv_ovrd_list")); err != nil {
		if !fortiAPIPatch(o["srv-ovrd-list"]) {
			return fmt.Errorf("Error reading srv_ovrd_list: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemFortiguardFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
