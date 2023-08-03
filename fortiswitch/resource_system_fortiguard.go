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

func resourceSystemFortiguard() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFortiguardUpdate,
		Read:   resourceSystemFortiguardRead,
		Update: resourceSystemFortiguardUpdate,
		Delete: resourceSystemFortiguardDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"webfilter_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"webfilter_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"avquery_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"antispam_cache_mpercent": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 15),
				Optional:     true,
				Computed:     true,
			},
			"webfilter_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Optional:     true,
				Computed:     true,
			},
			"antispam_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"load_balance_servers": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 50),
				Optional:     true,
				Computed:     true,
			},
			"antispam_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"avquery_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"avquery_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Optional:     true,
				Computed:     true,
			},
			"client_override_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antispam_score_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"avquery_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"avquery_cache_mpercent": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 15),
				Optional:     true,
				Computed:     true,
			},
			"avquery_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"avquery_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 24),
				Optional:     true,
				Computed:     true,
			},
			"antispam_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antispam_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"analysis_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_override_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srv_ovrd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"service_account_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 20),
				Optional:     true,
				Computed:     true,
			},
			"antispam_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"antispam_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Optional:     true,
				Computed:     true,
			},
			"srv_ovrd_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemFortiguardUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemFortiguard(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortiguard resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFortiguard(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortiguard resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFortiguard")
	}

	return resourceSystemFortiguardRead(d, m)
}

func resourceSystemFortiguardDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemFortiguard(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemFortiguard resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFortiguard(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemFortiguard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFortiguardRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemFortiguard(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortiguard resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFortiguard(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortiguard resource from API: %v", err)
	}
	return nil
}

func flattenSystemFortiguardWebfilterLicense(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAvqueryLicense(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamCacheMpercent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterExpiration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardLoadBalanceServers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamForceOff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAvqueryForceOff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterForceOff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAvqueryTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardClientOverrideStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamScoreThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAvqueryCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAvqueryCacheMpercent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAvqueryCacheTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAvqueryExpiration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamLicense(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamCacheTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAnalysisService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardClientOverrideIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardSrvOvrd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterCacheTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardServiceAccountId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamExpiration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardSrvOvrdList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

			tmp["ip"] = flattenSystemFortiguardSrvOvrdListIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {

			tmp["ip6"] = flattenSystemFortiguardSrvOvrdListIp6(i["ip6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {

			tmp["addr_type"] = flattenSystemFortiguardSrvOvrdListAddrType(i["addr-type"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "ip", d)
	return result
}

func flattenSystemFortiguardSrvOvrdListIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardSrvOvrdListIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardSrvOvrdListAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemFortiguard(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("webfilter_license", flattenSystemFortiguardWebfilterLicense(o["webfilter-license"], d, "webfilter_license", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-license"]) {
			return fmt.Errorf("Error reading webfilter_license: %v", err)
		}
	}

	if err = d.Set("webfilter_cache", flattenSystemFortiguardWebfilterCache(o["webfilter-cache"], d, "webfilter_cache", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-cache"]) {
			return fmt.Errorf("Error reading webfilter_cache: %v", err)
		}
	}

	if err = d.Set("avquery_license", flattenSystemFortiguardAvqueryLicense(o["avquery-license"], d, "avquery_license", sv)); err != nil {
		if !fortiAPIPatch(o["avquery-license"]) {
			return fmt.Errorf("Error reading avquery_license: %v", err)
		}
	}

	if err = d.Set("antispam_cache_mpercent", flattenSystemFortiguardAntispamCacheMpercent(o["antispam-cache-mpercent"], d, "antispam_cache_mpercent", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-cache-mpercent"]) {
			return fmt.Errorf("Error reading antispam_cache_mpercent: %v", err)
		}
	}

	if err = d.Set("webfilter_timeout", flattenSystemFortiguardWebfilterTimeout(o["webfilter-timeout"], d, "webfilter_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-timeout"]) {
			return fmt.Errorf("Error reading webfilter_timeout: %v", err)
		}
	}

	if err = d.Set("antispam_cache", flattenSystemFortiguardAntispamCache(o["antispam-cache"], d, "antispam_cache", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-cache"]) {
			return fmt.Errorf("Error reading antispam_cache: %v", err)
		}
	}

	if err = d.Set("webfilter_expiration", flattenSystemFortiguardWebfilterExpiration(o["webfilter-expiration"], d, "webfilter_expiration", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-expiration"]) {
			return fmt.Errorf("Error reading webfilter_expiration: %v", err)
		}
	}

	if err = d.Set("load_balance_servers", flattenSystemFortiguardLoadBalanceServers(o["load-balance-servers"], d, "load_balance_servers", sv)); err != nil {
		if !fortiAPIPatch(o["load-balance-servers"]) {
			return fmt.Errorf("Error reading load_balance_servers: %v", err)
		}
	}

	if err = d.Set("antispam_force_off", flattenSystemFortiguardAntispamForceOff(o["antispam-force-off"], d, "antispam_force_off", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-force-off"]) {
			return fmt.Errorf("Error reading antispam_force_off: %v", err)
		}
	}

	if err = d.Set("avquery_force_off", flattenSystemFortiguardAvqueryForceOff(o["avquery-force-off"], d, "avquery_force_off", sv)); err != nil {
		if !fortiAPIPatch(o["avquery-force-off"]) {
			return fmt.Errorf("Error reading avquery_force_off: %v", err)
		}
	}

	if err = d.Set("webfilter_force_off", flattenSystemFortiguardWebfilterForceOff(o["webfilter-force-off"], d, "webfilter_force_off", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-force-off"]) {
			return fmt.Errorf("Error reading webfilter_force_off: %v", err)
		}
	}

	if err = d.Set("avquery_timeout", flattenSystemFortiguardAvqueryTimeout(o["avquery-timeout"], d, "avquery_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["avquery-timeout"]) {
			return fmt.Errorf("Error reading avquery_timeout: %v", err)
		}
	}

	if err = d.Set("client_override_status", flattenSystemFortiguardClientOverrideStatus(o["client-override-status"], d, "client_override_status", sv)); err != nil {
		if !fortiAPIPatch(o["client-override-status"]) {
			return fmt.Errorf("Error reading client_override_status: %v", err)
		}
	}

	if err = d.Set("antispam_score_threshold", flattenSystemFortiguardAntispamScoreThreshold(o["antispam-score-threshold"], d, "antispam_score_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-score-threshold"]) {
			return fmt.Errorf("Error reading antispam_score_threshold: %v", err)
		}
	}

	if err = d.Set("avquery_cache", flattenSystemFortiguardAvqueryCache(o["avquery-cache"], d, "avquery_cache", sv)); err != nil {
		if !fortiAPIPatch(o["avquery-cache"]) {
			return fmt.Errorf("Error reading avquery_cache: %v", err)
		}
	}

	if err = d.Set("avquery_cache_mpercent", flattenSystemFortiguardAvqueryCacheMpercent(o["avquery-cache-mpercent"], d, "avquery_cache_mpercent", sv)); err != nil {
		if !fortiAPIPatch(o["avquery-cache-mpercent"]) {
			return fmt.Errorf("Error reading avquery_cache_mpercent: %v", err)
		}
	}

	if err = d.Set("avquery_cache_ttl", flattenSystemFortiguardAvqueryCacheTtl(o["avquery-cache-ttl"], d, "avquery_cache_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["avquery-cache-ttl"]) {
			return fmt.Errorf("Error reading avquery_cache_ttl: %v", err)
		}
	}

	if err = d.Set("avquery_expiration", flattenSystemFortiguardAvqueryExpiration(o["avquery-expiration"], d, "avquery_expiration", sv)); err != nil {
		if !fortiAPIPatch(o["avquery-expiration"]) {
			return fmt.Errorf("Error reading avquery_expiration: %v", err)
		}
	}

	if err = d.Set("hostname", flattenSystemFortiguardHostname(o["hostname"], d, "hostname", sv)); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("antispam_license", flattenSystemFortiguardAntispamLicense(o["antispam-license"], d, "antispam_license", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-license"]) {
			return fmt.Errorf("Error reading antispam_license: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemFortiguardPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("antispam_cache_ttl", flattenSystemFortiguardAntispamCacheTtl(o["antispam-cache-ttl"], d, "antispam_cache_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-cache-ttl"]) {
			return fmt.Errorf("Error reading antispam_cache_ttl: %v", err)
		}
	}

	if err = d.Set("analysis_service", flattenSystemFortiguardAnalysisService(o["analysis-service"], d, "analysis_service", sv)); err != nil {
		if !fortiAPIPatch(o["analysis-service"]) {
			return fmt.Errorf("Error reading analysis_service: %v", err)
		}
	}

	if err = d.Set("client_override_ip", flattenSystemFortiguardClientOverrideIp(o["client-override-ip"], d, "client_override_ip", sv)); err != nil {
		if !fortiAPIPatch(o["client-override-ip"]) {
			return fmt.Errorf("Error reading client_override_ip: %v", err)
		}
	}

	if err = d.Set("srv_ovrd", flattenSystemFortiguardSrvOvrd(o["srv-ovrd"], d, "srv_ovrd", sv)); err != nil {
		if !fortiAPIPatch(o["srv-ovrd"]) {
			return fmt.Errorf("Error reading srv_ovrd: %v", err)
		}
	}

	if err = d.Set("webfilter_cache_ttl", flattenSystemFortiguardWebfilterCacheTtl(o["webfilter-cache-ttl"], d, "webfilter_cache_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-cache-ttl"]) {
			return fmt.Errorf("Error reading webfilter_cache_ttl: %v", err)
		}
	}

	if err = d.Set("service_account_id", flattenSystemFortiguardServiceAccountId(o["service-account-id"], d, "service_account_id", sv)); err != nil {
		if !fortiAPIPatch(o["service-account-id"]) {
			return fmt.Errorf("Error reading service_account_id: %v", err)
		}
	}

	if err = d.Set("antispam_expiration", flattenSystemFortiguardAntispamExpiration(o["antispam-expiration"], d, "antispam_expiration", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-expiration"]) {
			return fmt.Errorf("Error reading antispam_expiration: %v", err)
		}
	}

	if err = d.Set("antispam_timeout", flattenSystemFortiguardAntispamTimeout(o["antispam-timeout"], d, "antispam_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-timeout"]) {
			return fmt.Errorf("Error reading antispam_timeout: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("srv_ovrd_list", flattenSystemFortiguardSrvOvrdList(o["srv-ovrd-list"], d, "srv_ovrd_list", sv)); err != nil {
			if !fortiAPIPatch(o["srv-ovrd-list"]) {
				return fmt.Errorf("Error reading srv_ovrd_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srv_ovrd_list"); ok {
			if err = d.Set("srv_ovrd_list", flattenSystemFortiguardSrvOvrdList(o["srv-ovrd-list"], d, "srv_ovrd_list", sv)); err != nil {
				if !fortiAPIPatch(o["srv-ovrd-list"]) {
					return fmt.Errorf("Error reading srv_ovrd_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemFortiguardFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemFortiguardWebfilterLicense(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAvqueryLicense(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamCacheMpercent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterExpiration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardLoadBalanceServers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamForceOff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAvqueryForceOff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterForceOff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAvqueryTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardClientOverrideStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamScoreThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAvqueryCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAvqueryCacheMpercent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAvqueryCacheTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAvqueryExpiration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamLicense(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamCacheTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAnalysisService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardClientOverrideIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSrvOvrd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterCacheTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardServiceAccountId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamExpiration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSrvOvrdList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandSystemFortiguardSrvOvrdListIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip6"], _ = expandSystemFortiguardSrvOvrdListIp6(d, i["ip6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["addr-type"], _ = expandSystemFortiguardSrvOvrdListAddrType(d, i["addr_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemFortiguardSrvOvrdListIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSrvOvrdListIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSrvOvrdListAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFortiguard(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("webfilter_license"); ok {
		if setArgNil {
			obj["webfilter-license"] = nil
		} else {

			t, err := expandSystemFortiguardWebfilterLicense(d, v, "webfilter_license", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["webfilter-license"] = t
			}
		}
	}

	if v, ok := d.GetOk("webfilter_cache"); ok {
		if setArgNil {
			obj["webfilter-cache"] = nil
		} else {

			t, err := expandSystemFortiguardWebfilterCache(d, v, "webfilter_cache", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["webfilter-cache"] = t
			}
		}
	}

	if v, ok := d.GetOk("avquery_license"); ok {
		if setArgNil {
			obj["avquery-license"] = nil
		} else {

			t, err := expandSystemFortiguardAvqueryLicense(d, v, "avquery_license", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["avquery-license"] = t
			}
		}
	}

	if v, ok := d.GetOk("antispam_cache_mpercent"); ok {
		if setArgNil {
			obj["antispam-cache-mpercent"] = nil
		} else {

			t, err := expandSystemFortiguardAntispamCacheMpercent(d, v, "antispam_cache_mpercent", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-cache-mpercent"] = t
			}
		}
	}

	if v, ok := d.GetOk("webfilter_timeout"); ok {
		if setArgNil {
			obj["webfilter-timeout"] = nil
		} else {

			t, err := expandSystemFortiguardWebfilterTimeout(d, v, "webfilter_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["webfilter-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("antispam_cache"); ok {
		if setArgNil {
			obj["antispam-cache"] = nil
		} else {

			t, err := expandSystemFortiguardAntispamCache(d, v, "antispam_cache", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-cache"] = t
			}
		}
	}

	if v, ok := d.GetOk("webfilter_expiration"); ok {
		if setArgNil {
			obj["webfilter-expiration"] = nil
		} else {

			t, err := expandSystemFortiguardWebfilterExpiration(d, v, "webfilter_expiration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["webfilter-expiration"] = t
			}
		}
	}

	if v, ok := d.GetOk("load_balance_servers"); ok {
		if setArgNil {
			obj["load-balance-servers"] = nil
		} else {

			t, err := expandSystemFortiguardLoadBalanceServers(d, v, "load_balance_servers", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["load-balance-servers"] = t
			}
		}
	}

	if v, ok := d.GetOk("antispam_force_off"); ok {
		if setArgNil {
			obj["antispam-force-off"] = nil
		} else {

			t, err := expandSystemFortiguardAntispamForceOff(d, v, "antispam_force_off", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-force-off"] = t
			}
		}
	}

	if v, ok := d.GetOk("avquery_force_off"); ok {
		if setArgNil {
			obj["avquery-force-off"] = nil
		} else {

			t, err := expandSystemFortiguardAvqueryForceOff(d, v, "avquery_force_off", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["avquery-force-off"] = t
			}
		}
	}

	if v, ok := d.GetOk("webfilter_force_off"); ok {
		if setArgNil {
			obj["webfilter-force-off"] = nil
		} else {

			t, err := expandSystemFortiguardWebfilterForceOff(d, v, "webfilter_force_off", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["webfilter-force-off"] = t
			}
		}
	}

	if v, ok := d.GetOk("avquery_timeout"); ok {
		if setArgNil {
			obj["avquery-timeout"] = nil
		} else {

			t, err := expandSystemFortiguardAvqueryTimeout(d, v, "avquery_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["avquery-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("client_override_status"); ok {
		if setArgNil {
			obj["client-override-status"] = nil
		} else {

			t, err := expandSystemFortiguardClientOverrideStatus(d, v, "client_override_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["client-override-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("antispam_score_threshold"); ok {
		if setArgNil {
			obj["antispam-score-threshold"] = nil
		} else {

			t, err := expandSystemFortiguardAntispamScoreThreshold(d, v, "antispam_score_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-score-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("avquery_cache"); ok {
		if setArgNil {
			obj["avquery-cache"] = nil
		} else {

			t, err := expandSystemFortiguardAvqueryCache(d, v, "avquery_cache", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["avquery-cache"] = t
			}
		}
	}

	if v, ok := d.GetOk("avquery_cache_mpercent"); ok {
		if setArgNil {
			obj["avquery-cache-mpercent"] = nil
		} else {

			t, err := expandSystemFortiguardAvqueryCacheMpercent(d, v, "avquery_cache_mpercent", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["avquery-cache-mpercent"] = t
			}
		}
	}

	if v, ok := d.GetOk("avquery_cache_ttl"); ok {
		if setArgNil {
			obj["avquery-cache-ttl"] = nil
		} else {

			t, err := expandSystemFortiguardAvqueryCacheTtl(d, v, "avquery_cache_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["avquery-cache-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOk("avquery_expiration"); ok {
		if setArgNil {
			obj["avquery-expiration"] = nil
		} else {

			t, err := expandSystemFortiguardAvqueryExpiration(d, v, "avquery_expiration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["avquery-expiration"] = t
			}
		}
	}

	if v, ok := d.GetOk("hostname"); ok {
		if setArgNil {
			obj["hostname"] = nil
		} else {

			t, err := expandSystemFortiguardHostname(d, v, "hostname", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostname"] = t
			}
		}
	}

	if v, ok := d.GetOk("antispam_license"); ok {
		if setArgNil {
			obj["antispam-license"] = nil
		} else {

			t, err := expandSystemFortiguardAntispamLicense(d, v, "antispam_license", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-license"] = t
			}
		}
	}

	if v, ok := d.GetOk("port"); ok {
		if setArgNil {
			obj["port"] = nil
		} else {

			t, err := expandSystemFortiguardPort(d, v, "port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["port"] = t
			}
		}
	}

	if v, ok := d.GetOk("antispam_cache_ttl"); ok {
		if setArgNil {
			obj["antispam-cache-ttl"] = nil
		} else {

			t, err := expandSystemFortiguardAntispamCacheTtl(d, v, "antispam_cache_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-cache-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOk("analysis_service"); ok {
		if setArgNil {
			obj["analysis-service"] = nil
		} else {

			t, err := expandSystemFortiguardAnalysisService(d, v, "analysis_service", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["analysis-service"] = t
			}
		}
	}

	if v, ok := d.GetOk("client_override_ip"); ok {
		if setArgNil {
			obj["client-override-ip"] = nil
		} else {

			t, err := expandSystemFortiguardClientOverrideIp(d, v, "client_override_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["client-override-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("srv_ovrd"); ok {
		if setArgNil {
			obj["srv-ovrd"] = nil
		} else {

			t, err := expandSystemFortiguardSrvOvrd(d, v, "srv_ovrd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["srv-ovrd"] = t
			}
		}
	}

	if v, ok := d.GetOk("webfilter_cache_ttl"); ok {
		if setArgNil {
			obj["webfilter-cache-ttl"] = nil
		} else {

			t, err := expandSystemFortiguardWebfilterCacheTtl(d, v, "webfilter_cache_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["webfilter-cache-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOk("service_account_id"); ok {
		if setArgNil {
			obj["service-account-id"] = nil
		} else {

			t, err := expandSystemFortiguardServiceAccountId(d, v, "service_account_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["service-account-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("antispam_expiration"); ok {
		if setArgNil {
			obj["antispam-expiration"] = nil
		} else {

			t, err := expandSystemFortiguardAntispamExpiration(d, v, "antispam_expiration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-expiration"] = t
			}
		}
	}

	if v, ok := d.GetOk("antispam_timeout"); ok {
		if setArgNil {
			obj["antispam-timeout"] = nil
		} else {

			t, err := expandSystemFortiguardAntispamTimeout(d, v, "antispam_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("srv_ovrd_list"); ok || d.HasChange("srv_ovrd_list") {
		if setArgNil {
			obj["srv-ovrd-list"] = make([]struct{}, 0)
		} else {

			t, err := expandSystemFortiguardSrvOvrdList(d, v, "srv_ovrd_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["srv-ovrd-list"] = t
			}
		}
	}

	return &obj, nil
}
