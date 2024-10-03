// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: RADIUS server entry configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserRadius() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserRadiusCreate,
		Read:   resourceUserRadiusRead,
		Update: resourceUserRadiusUpdate,
		Delete: resourceUserRadiusDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"radius_coa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link_monitor_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 120),
				Optional:     true,
				Computed:     true,
			},
			"secondary_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Sensitive:    true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nas_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acct_interim_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),
				Optional:     true,
				Computed:     true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"all_usergroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Sensitive:    true,
			},
			"nas_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acct_fast_framedip_detect": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 600),
				Optional:     true,
				Computed:     true,
			},
			"frame_mtu_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(600, 1500),
				Optional:     true,
				Computed:     true,
			},
			"secondary_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"service_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"addr_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_coa_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Sensitive:    true,
			},
			"acct_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secret": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Sensitive:    true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"server": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"radius_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceUserRadiusCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserRadius(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserRadius resource while getting object: %v", err)
	}

	o, err := c.CreateUserRadius(obj)

	if err != nil {
		return fmt.Errorf("Error creating UserRadius resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserRadius")
	}

	return resourceUserRadiusRead(d, m)
}

func resourceUserRadiusUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserRadius(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserRadius resource while getting object: %v", err)
	}

	o, err := c.UpdateUserRadius(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating UserRadius resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserRadius")
	}

	return resourceUserRadiusRead(d, m)
}

func resourceUserRadiusDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteUserRadius(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting UserRadius resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserRadiusRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadUserRadius(mkey)
	if err != nil {
		return fmt.Errorf("Error reading UserRadius resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserRadius(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserRadius resource from API: %v", err)
	}
	return nil
}

func flattenUserRadiusRadiusCoa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusLinkMonitorInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusSecondarySecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusSourceIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusNasIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAcctInterimInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAuthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAllUsergroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusSecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusNasIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusLinkMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAcctFastFramedipDetect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusFrameMtuSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusSecondaryServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusServiceType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAddrMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRadiusCoaSecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAcctServer(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenUserRadiusAcctServerStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secret"
		if _, ok := i["secret"]; ok {

			tmp["secret"] = flattenUserRadiusAcctServerSecret(i["secret"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenUserRadiusAcctServerPort(i["port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenUserRadiusAcctServerId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {

			tmp["server"] = flattenUserRadiusAcctServerServer(i["server"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenUserRadiusAcctServerStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAcctServerSecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAcctServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAcctServerId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAcctServerServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRadiusPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserRadius(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("radius_coa", flattenUserRadiusRadiusCoa(o["radius-coa"], d, "radius_coa", sv)); err != nil {
		if !fortiAPIPatch(o["radius-coa"]) {
			return fmt.Errorf("Error reading radius_coa: %v", err)
		}
	}

	if err = d.Set("link_monitor_interval", flattenUserRadiusLinkMonitorInterval(o["link-monitor-interval"], d, "link_monitor_interval", sv)); err != nil {
		if !fortiAPIPatch(o["link-monitor-interval"]) {
			return fmt.Errorf("Error reading link_monitor_interval: %v", err)
		}
	}

	if err = d.Set("secondary_secret", flattenUserRadiusSecondarySecret(o["secondary-secret"], d, "secondary_secret", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-secret"]) {
			return fmt.Errorf("Error reading secondary_secret: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenUserRadiusSourceIp6(o["source-ip6"], d, "source_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("nas_ip6", flattenUserRadiusNasIp6(o["nas-ip6"], d, "nas_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["nas-ip6"]) {
			return fmt.Errorf("Error reading nas_ip6: %v", err)
		}
	}

	if err = d.Set("acct_interim_interval", flattenUserRadiusAcctInterimInterval(o["acct-interim-interval"], d, "acct_interim_interval", sv)); err != nil {
		if !fortiAPIPatch(o["acct-interim-interval"]) {
			return fmt.Errorf("Error reading acct_interim_interval: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenUserRadiusSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenUserRadiusAuthType(o["auth-type"], d, "auth_type", sv)); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("all_usergroup", flattenUserRadiusAllUsergroup(o["all-usergroup"], d, "all_usergroup", sv)); err != nil {
		if !fortiAPIPatch(o["all-usergroup"]) {
			return fmt.Errorf("Error reading all_usergroup: %v", err)
		}
	}

	if err = d.Set("nas_ip", flattenUserRadiusNasIp(o["nas-ip"], d, "nas_ip", sv)); err != nil {
		if !fortiAPIPatch(o["nas-ip"]) {
			return fmt.Errorf("Error reading nas_ip: %v", err)
		}
	}

	if err = d.Set("link_monitor", flattenUserRadiusLinkMonitor(o["link-monitor"], d, "link_monitor", sv)); err != nil {
		if !fortiAPIPatch(o["link-monitor"]) {
			return fmt.Errorf("Error reading link_monitor: %v", err)
		}
	}

	if err = d.Set("acct_fast_framedip_detect", flattenUserRadiusAcctFastFramedipDetect(o["acct-fast-framedip-detect"], d, "acct_fast_framedip_detect", sv)); err != nil {
		if !fortiAPIPatch(o["acct-fast-framedip-detect"]) {
			return fmt.Errorf("Error reading acct_fast_framedip_detect: %v", err)
		}
	}

	if err = d.Set("frame_mtu_size", flattenUserRadiusFrameMtuSize(o["frame-mtu-size"], d, "frame_mtu_size", sv)); err != nil {
		if !fortiAPIPatch(o["frame-mtu-size"]) {
			return fmt.Errorf("Error reading frame_mtu_size: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenUserRadiusSecondaryServer(o["secondary-server"], d, "secondary_server", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-server"]) {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("service_type", flattenUserRadiusServiceType(o["service-type"], d, "service_type", sv)); err != nil {
		if !fortiAPIPatch(o["service-type"]) {
			return fmt.Errorf("Error reading service_type: %v", err)
		}
	}

	if err = d.Set("name", flattenUserRadiusName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server", flattenUserRadiusServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("addr_mode", flattenUserRadiusAddrMode(o["addr-mode"], d, "addr_mode", sv)); err != nil {
		if !fortiAPIPatch(o["addr-mode"]) {
			return fmt.Errorf("Error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("radius_coa_secret", flattenUserRadiusRadiusCoaSecret(o["radius-coa-secret"], d, "radius_coa_secret", sv)); err != nil {
		if !fortiAPIPatch(o["radius-coa-secret"]) {
			return fmt.Errorf("Error reading radius_coa_secret: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("acct_server", flattenUserRadiusAcctServer(o["acct-server"], d, "acct_server", sv)); err != nil {
			if !fortiAPIPatch(o["acct-server"]) {
				return fmt.Errorf("Error reading acct_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("acct_server"); ok {
			if err = d.Set("acct_server", flattenUserRadiusAcctServer(o["acct-server"], d, "acct_server", sv)); err != nil {
				if !fortiAPIPatch(o["acct-server"]) {
					return fmt.Errorf("Error reading acct_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("radius_port", flattenUserRadiusRadiusPort(o["radius-port"], d, "radius_port", sv)); err != nil {
		if !fortiAPIPatch(o["radius-port"]) {
			return fmt.Errorf("Error reading radius_port: %v", err)
		}
	}

	return nil
}

func flattenUserRadiusFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandUserRadiusRadiusCoa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusLinkMonitorInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSecondarySecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSourceIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusNasIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAcctInterimInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAuthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAllUsergroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusNasIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusLinkMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAcctFastFramedipDetect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusFrameMtuSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSecondaryServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusServiceType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAddrMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRadiusCoaSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAcctServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandUserRadiusAcctServerStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secret"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["secret"], _ = expandUserRadiusAcctServerSecret(d, i["secret"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandUserRadiusAcctServerPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandUserRadiusAcctServerId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["server"], _ = expandUserRadiusAcctServerServer(d, i["server"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserRadiusAcctServerStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAcctServerSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAcctServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAcctServerId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAcctServerServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRadiusPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserRadius(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("radius_coa"); ok {

		t, err := expandUserRadiusRadiusCoa(d, v, "radius_coa", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-coa"] = t
		}
	}

	if v, ok := d.GetOk("link_monitor_interval"); ok {

		t, err := expandUserRadiusLinkMonitorInterval(d, v, "link_monitor_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-monitor-interval"] = t
		}
	}

	if v, ok := d.GetOk("secondary_secret"); ok {

		t, err := expandUserRadiusSecondarySecret(d, v, "secondary_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-secret"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {

		t, err := expandUserRadiusSourceIp6(d, v, "source_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("nas_ip6"); ok {

		t, err := expandUserRadiusNasIp6(d, v, "nas_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-ip6"] = t
		}
	}

	if v, ok := d.GetOk("acct_interim_interval"); ok {

		t, err := expandUserRadiusAcctInterimInterval(d, v, "acct_interim_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-interim-interval"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {

		t, err := expandUserRadiusSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok {

		t, err := expandUserRadiusAuthType(d, v, "auth_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("all_usergroup"); ok {

		t, err := expandUserRadiusAllUsergroup(d, v, "all_usergroup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["all-usergroup"] = t
		}
	}

	if v, ok := d.GetOk("secret"); ok {

		t, err := expandUserRadiusSecret(d, v, "secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret"] = t
		}
	}

	if v, ok := d.GetOk("nas_ip"); ok {

		t, err := expandUserRadiusNasIp(d, v, "nas_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-ip"] = t
		}
	}

	if v, ok := d.GetOk("link_monitor"); ok {

		t, err := expandUserRadiusLinkMonitor(d, v, "link_monitor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-monitor"] = t
		}
	}

	if v, ok := d.GetOk("acct_fast_framedip_detect"); ok {

		t, err := expandUserRadiusAcctFastFramedipDetect(d, v, "acct_fast_framedip_detect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-fast-framedip-detect"] = t
		}
	}

	if v, ok := d.GetOk("frame_mtu_size"); ok {

		t, err := expandUserRadiusFrameMtuSize(d, v, "frame_mtu_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["frame-mtu-size"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok {

		t, err := expandUserRadiusSecondaryServer(d, v, "secondary_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("service_type"); ok {

		t, err := expandUserRadiusServiceType(d, v, "service_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandUserRadiusName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {

		t, err := expandUserRadiusServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("addr_mode"); ok {

		t, err := expandUserRadiusAddrMode(d, v, "addr_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("radius_coa_secret"); ok {

		t, err := expandUserRadiusRadiusCoaSecret(d, v, "radius_coa_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-coa-secret"] = t
		}
	}

	if v, ok := d.GetOk("acct_server"); ok || d.HasChange("acct_server") {

		t, err := expandUserRadiusAcctServer(d, v, "acct_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-server"] = t
		}
	}

	if v, ok := d.GetOk("radius_port"); ok {

		t, err := expandUserRadiusRadiusPort(d, v, "radius_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-port"] = t
		}
	}

	return &obj, nil
}
