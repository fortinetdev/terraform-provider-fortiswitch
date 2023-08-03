// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure VXLAN devices.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemVxlan() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemVxlanRead,
		Schema: map[string]*schema.Schema{

			"remote_ip": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"multicast_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tagged_vlans": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vni": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlanid": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemVxlanRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemVxlan: type error")
	}

	o, err := c.ReadSystemVxlan(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemVxlan: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemVxlan(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemVxlan from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemVxlanRemoteIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["ip"] = dataSourceFlattenSystemVxlanRemoteIpIp(i["ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemVxlanRemoteIpIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVxlanName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVxlanMulticastTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVxlanTaggedVlans(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVxlanVni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVxlanIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVxlanVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVxlanInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemVxlan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("remote_ip", dataSourceFlattenSystemVxlanRemoteIp(o["remote-ip"], d, "remote_ip")); err != nil {
		if !fortiAPIPatch(o["remote-ip"]) {
			return fmt.Errorf("Error reading remote_ip: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenSystemVxlanName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("multicast_ttl", dataSourceFlattenSystemVxlanMulticastTtl(o["multicast-ttl"], d, "multicast_ttl")); err != nil {
		if !fortiAPIPatch(o["multicast-ttl"]) {
			return fmt.Errorf("Error reading multicast_ttl: %v", err)
		}
	}

	if err = d.Set("tagged_vlans", dataSourceFlattenSystemVxlanTaggedVlans(o["tagged-vlans"], d, "tagged_vlans")); err != nil {
		if !fortiAPIPatch(o["tagged-vlans"]) {
			return fmt.Errorf("Error reading tagged_vlans: %v", err)
		}
	}

	if err = d.Set("vni", dataSourceFlattenSystemVxlanVni(o["vni"], d, "vni")); err != nil {
		if !fortiAPIPatch(o["vni"]) {
			return fmt.Errorf("Error reading vni: %v", err)
		}
	}

	if err = d.Set("ip_version", dataSourceFlattenSystemVxlanIpVersion(o["ip-version"], d, "ip_version")); err != nil {
		if !fortiAPIPatch(o["ip-version"]) {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("vlanid", dataSourceFlattenSystemVxlanVlanid(o["vlanid"], d, "vlanid")); err != nil {
		if !fortiAPIPatch(o["vlanid"]) {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemVxlanInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemVxlanFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
