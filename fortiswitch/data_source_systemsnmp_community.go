// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: SNMP community configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemSnmpCommunity() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemSnmpCommunityRead,
		Schema: map[string]*schema.Schema{

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trap_v2c_lport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"hosts6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_ipv6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ipv6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"query_v1_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"query_v2c_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"query_v2c_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trap_v1_rport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"query_v1_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trap_v2c_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"trap_v1_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"events": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trap_v2c_rport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"fswid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"trap_v1_lport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemSnmpCommunityRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("fswid")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemSnmpCommunity: type error")
	}

	o, err := c.ReadSystemSnmpCommunity(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemSnmpCommunity: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemSnmpCommunity(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemSnmpCommunity from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemSnmpCommunityStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityTrapV2CLport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHosts6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ipv6"
		if _, ok := i["source-ipv6"]; ok {
			tmp["source_ipv6"] = dataSourceFlattenSystemSnmpCommunityHosts6SourceIpv6(i["source-ipv6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenSystemSnmpCommunityHosts6Interface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenSystemSnmpCommunityHosts6Id(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6"
		if _, ok := i["ipv6"]; ok {
			tmp["ipv6"] = dataSourceFlattenSystemSnmpCommunityHosts6Ipv6(i["ipv6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSnmpCommunityHosts6SourceIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHosts6Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHosts6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHosts6Ipv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityQueryV1Port(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityQueryV2CPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityQueryV2CStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityTrapV1Rport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityQueryV1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityTrapV2CStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHosts(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenSystemSnmpCommunityHostsInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenSystemSnmpCommunityHostsIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			tmp["source_ip"] = dataSourceFlattenSystemSnmpCommunityHostsSourceIp(i["source-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenSystemSnmpCommunityHostsId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSnmpCommunityHostsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHostsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHostsSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHostsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityTrapV1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityEvents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityTrapV2CRport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityTrapV1Lport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemSnmpCommunity(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenSystemSnmpCommunityStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trap_v2c_lport", dataSourceFlattenSystemSnmpCommunityTrapV2CLport(o["trap-v2c-lport"], d, "trap_v2c_lport")); err != nil {
		if !fortiAPIPatch(o["trap-v2c-lport"]) {
			return fmt.Errorf("Error reading trap_v2c_lport: %v", err)
		}
	}

	if err = d.Set("hosts6", dataSourceFlattenSystemSnmpCommunityHosts6(o["hosts6"], d, "hosts6")); err != nil {
		if !fortiAPIPatch(o["hosts6"]) {
			return fmt.Errorf("Error reading hosts6: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenSystemSnmpCommunityName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("query_v1_port", dataSourceFlattenSystemSnmpCommunityQueryV1Port(o["query-v1-port"], d, "query_v1_port")); err != nil {
		if !fortiAPIPatch(o["query-v1-port"]) {
			return fmt.Errorf("Error reading query_v1_port: %v", err)
		}
	}

	if err = d.Set("query_v2c_port", dataSourceFlattenSystemSnmpCommunityQueryV2CPort(o["query-v2c-port"], d, "query_v2c_port")); err != nil {
		if !fortiAPIPatch(o["query-v2c-port"]) {
			return fmt.Errorf("Error reading query_v2c_port: %v", err)
		}
	}

	if err = d.Set("query_v2c_status", dataSourceFlattenSystemSnmpCommunityQueryV2CStatus(o["query-v2c-status"], d, "query_v2c_status")); err != nil {
		if !fortiAPIPatch(o["query-v2c-status"]) {
			return fmt.Errorf("Error reading query_v2c_status: %v", err)
		}
	}

	if err = d.Set("trap_v1_rport", dataSourceFlattenSystemSnmpCommunityTrapV1Rport(o["trap-v1-rport"], d, "trap_v1_rport")); err != nil {
		if !fortiAPIPatch(o["trap-v1-rport"]) {
			return fmt.Errorf("Error reading trap_v1_rport: %v", err)
		}
	}

	if err = d.Set("query_v1_status", dataSourceFlattenSystemSnmpCommunityQueryV1Status(o["query-v1-status"], d, "query_v1_status")); err != nil {
		if !fortiAPIPatch(o["query-v1-status"]) {
			return fmt.Errorf("Error reading query_v1_status: %v", err)
		}
	}

	if err = d.Set("trap_v2c_status", dataSourceFlattenSystemSnmpCommunityTrapV2CStatus(o["trap-v2c-status"], d, "trap_v2c_status")); err != nil {
		if !fortiAPIPatch(o["trap-v2c-status"]) {
			return fmt.Errorf("Error reading trap_v2c_status: %v", err)
		}
	}

	if err = d.Set("hosts", dataSourceFlattenSystemSnmpCommunityHosts(o["hosts"], d, "hosts")); err != nil {
		if !fortiAPIPatch(o["hosts"]) {
			return fmt.Errorf("Error reading hosts: %v", err)
		}
	}

	if err = d.Set("trap_v1_status", dataSourceFlattenSystemSnmpCommunityTrapV1Status(o["trap-v1-status"], d, "trap_v1_status")); err != nil {
		if !fortiAPIPatch(o["trap-v1-status"]) {
			return fmt.Errorf("Error reading trap_v1_status: %v", err)
		}
	}

	if err = d.Set("events", dataSourceFlattenSystemSnmpCommunityEvents(o["events"], d, "events")); err != nil {
		if !fortiAPIPatch(o["events"]) {
			return fmt.Errorf("Error reading events: %v", err)
		}
	}

	if err = d.Set("trap_v2c_rport", dataSourceFlattenSystemSnmpCommunityTrapV2CRport(o["trap-v2c-rport"], d, "trap_v2c_rport")); err != nil {
		if !fortiAPIPatch(o["trap-v2c-rport"]) {
			return fmt.Errorf("Error reading trap_v2c_rport: %v", err)
		}
	}

	if err = d.Set("fswid", dataSourceFlattenSystemSnmpCommunityId(o["id"], d, "fswid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fswid: %v", err)
		}
	}

	if err = d.Set("trap_v1_lport", dataSourceFlattenSystemSnmpCommunityTrapV1Lport(o["trap-v1-lport"], d, "trap_v1_lport")); err != nil {
		if !fortiAPIPatch(o["trap-v1-lport"]) {
			return fmt.Errorf("Error reading trap_v1_lport: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemSnmpCommunityFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
