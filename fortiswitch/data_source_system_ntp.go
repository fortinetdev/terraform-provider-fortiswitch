// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Ntp system info configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemNtp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemNtpRead,
		Schema: map[string]*schema.Schema{

			"key_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_time_adjustments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_unsync_source": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"ntpsync": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"syncinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ntpserver": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ntpv3": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"key_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"key_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemNtpRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemNtp"

	o, err := c.ReadSystemNtp(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemNtp: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemNtp(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemNtp from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemNtpKeyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpLogTimeAdjustments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpServerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpAllowUnsyncSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpsync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpSyncinterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserver(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ntpv3"
		if _, ok := i["ntpv3"]; ok {
			tmp["ntpv3"] = dataSourceFlattenSystemNtpNtpserverNtpv3(i["ntpv3"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			tmp["server"] = dataSourceFlattenSystemNtpNtpserverServer(i["server"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = dataSourceFlattenSystemNtpNtpserverAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {
			tmp["key"] = dataSourceFlattenSystemNtpNtpserverKey(i["key"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenSystemNtpNtpserverId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_id"
		if _, ok := i["key-id"]; ok {
			tmp["key_id"] = dataSourceFlattenSystemNtpNtpserverKeyId(i["key-id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemNtpNtpserverNtpv3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserverServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserverAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserverKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserverId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserverKeyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpKeyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := i["interface-name"]; ok {
			tmp["interface_name"] = dataSourceFlattenSystemNtpInterfaceInterfaceName(i["interface-name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemNtpInterfaceInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemNtp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("key_type", dataSourceFlattenSystemNtpKeyType(o["key-type"], d, "key_type")); err != nil {
		if !fortiAPIPatch(o["key-type"]) {
			return fmt.Errorf("Error reading key_type: %v", err)
		}
	}

	if err = d.Set("log_time_adjustments", dataSourceFlattenSystemNtpLogTimeAdjustments(o["log-time-adjustments"], d, "log_time_adjustments")); err != nil {
		if !fortiAPIPatch(o["log-time-adjustments"]) {
			return fmt.Errorf("Error reading log_time_adjustments: %v", err)
		}
	}

	if err = d.Set("server_mode", dataSourceFlattenSystemNtpServerMode(o["server-mode"], d, "server_mode")); err != nil {
		if !fortiAPIPatch(o["server-mode"]) {
			return fmt.Errorf("Error reading server_mode: %v", err)
		}
	}

	if err = d.Set("source_ip", dataSourceFlattenSystemNtpSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("allow_unsync_source", dataSourceFlattenSystemNtpAllowUnsyncSource(o["allow-unsync-source"], d, "allow_unsync_source")); err != nil {
		if !fortiAPIPatch(o["allow-unsync-source"]) {
			return fmt.Errorf("Error reading allow_unsync_source: %v", err)
		}
	}

	if err = d.Set("authentication", dataSourceFlattenSystemNtpAuthentication(o["authentication"], d, "authentication")); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("source_ip6", dataSourceFlattenSystemNtpSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("ntpsync", dataSourceFlattenSystemNtpNtpsync(o["ntpsync"], d, "ntpsync")); err != nil {
		if !fortiAPIPatch(o["ntpsync"]) {
			return fmt.Errorf("Error reading ntpsync: %v", err)
		}
	}

	if err = d.Set("syncinterval", dataSourceFlattenSystemNtpSyncinterval(o["syncinterval"], d, "syncinterval")); err != nil {
		if !fortiAPIPatch(o["syncinterval"]) {
			return fmt.Errorf("Error reading syncinterval: %v", err)
		}
	}

	if err = d.Set("ntpserver", dataSourceFlattenSystemNtpNtpserver(o["ntpserver"], d, "ntpserver")); err != nil {
		if !fortiAPIPatch(o["ntpserver"]) {
			return fmt.Errorf("Error reading ntpserver: %v", err)
		}
	}

	if err = d.Set("key_id", dataSourceFlattenSystemNtpKeyId(o["key-id"], d, "key_id")); err != nil {
		if !fortiAPIPatch(o["key-id"]) {
			return fmt.Errorf("Error reading key_id: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemNtpInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemNtpFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
