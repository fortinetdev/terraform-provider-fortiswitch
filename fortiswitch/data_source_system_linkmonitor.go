// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure Link Health Monitor.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemLinkMonitor() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemLinkMonitorRead,
		Schema: map[string]*schema.Schema{

			"update_cascade_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"http_match": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"gateway_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"failtime": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"update_static_route": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"addr_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_get": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"packet_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"gateway_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"recoverytime": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemLinkMonitorRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemLinkMonitor: type error")
	}

	o, err := c.ReadSystemLinkMonitor(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemLinkMonitor: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemLinkMonitor(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemLinkMonitor from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemLinkMonitorUpdateCascadeInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorHttpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorGatewayIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorFailtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorUpdateStaticRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorHttpGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorSrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorPacketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorGatewayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLinkMonitorRecoverytime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemLinkMonitor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("update_cascade_interface", dataSourceFlattenSystemLinkMonitorUpdateCascadeInterface(o["update-cascade-interface"], d, "update_cascade_interface")); err != nil {
		if !fortiAPIPatch(o["update-cascade-interface"]) {
			return fmt.Errorf("Error reading update_cascade_interface: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemLinkMonitorStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("timeout", dataSourceFlattenSystemLinkMonitorTimeout(o["timeout"], d, "timeout")); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("protocol", dataSourceFlattenSystemLinkMonitorProtocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenSystemLinkMonitorName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("http_match", dataSourceFlattenSystemLinkMonitorHttpMatch(o["http-match"], d, "http_match")); err != nil {
		if !fortiAPIPatch(o["http-match"]) {
			return fmt.Errorf("Error reading http_match: %v", err)
		}
	}

	if err = d.Set("source_ip", dataSourceFlattenSystemLinkMonitorSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("interval", dataSourceFlattenSystemLinkMonitorInterval(o["interval"], d, "interval")); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("gateway_ip6", dataSourceFlattenSystemLinkMonitorGatewayIp6(o["gateway-ip6"], d, "gateway_ip6")); err != nil {
		if !fortiAPIPatch(o["gateway-ip6"]) {
			return fmt.Errorf("Error reading gateway_ip6: %v", err)
		}
	}

	if err = d.Set("failtime", dataSourceFlattenSystemLinkMonitorFailtime(o["failtime"], d, "failtime")); err != nil {
		if !fortiAPIPatch(o["failtime"]) {
			return fmt.Errorf("Error reading failtime: %v", err)
		}
	}

	if err = d.Set("update_static_route", dataSourceFlattenSystemLinkMonitorUpdateStaticRoute(o["update-static-route"], d, "update_static_route")); err != nil {
		if !fortiAPIPatch(o["update-static-route"]) {
			return fmt.Errorf("Error reading update_static_route: %v", err)
		}
	}

	if err = d.Set("addr_mode", dataSourceFlattenSystemLinkMonitorAddrMode(o["addr-mode"], d, "addr_mode")); err != nil {
		if !fortiAPIPatch(o["addr-mode"]) {
			return fmt.Errorf("Error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("http_get", dataSourceFlattenSystemLinkMonitorHttpGet(o["http-get"], d, "http_get")); err != nil {
		if !fortiAPIPatch(o["http-get"]) {
			return fmt.Errorf("Error reading http_get: %v", err)
		}
	}

	if err = d.Set("source_ip6", dataSourceFlattenSystemLinkMonitorSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("srcintf", dataSourceFlattenSystemLinkMonitorSrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if !fortiAPIPatch(o["srcintf"]) {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("security_mode", dataSourceFlattenSystemLinkMonitorSecurityMode(o["security-mode"], d, "security_mode")); err != nil {
		if !fortiAPIPatch(o["security-mode"]) {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("packet_size", dataSourceFlattenSystemLinkMonitorPacketSize(o["packet-size"], d, "packet_size")); err != nil {
		if !fortiAPIPatch(o["packet-size"]) {
			return fmt.Errorf("Error reading packet_size: %v", err)
		}
	}

	if err = d.Set("gateway_ip", dataSourceFlattenSystemLinkMonitorGatewayIp(o["gateway-ip"], d, "gateway_ip")); err != nil {
		if !fortiAPIPatch(o["gateway-ip"]) {
			return fmt.Errorf("Error reading gateway_ip: %v", err)
		}
	}

	if err = d.Set("password", dataSourceFlattenSystemLinkMonitorPassword(o["password"], d, "password")); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenSystemLinkMonitorPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("recoverytime", dataSourceFlattenSystemLinkMonitorRecoverytime(o["recoverytime"], d, "recoverytime")); err != nil {
		if !fortiAPIPatch(o["recoverytime"]) {
			return fmt.Errorf("Error reading recoverytime: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemLinkMonitorFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
