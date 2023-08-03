// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure IPv6 neighbor cache table.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemIpv6NeighborCache() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemIpv6NeighborCacheRead,
		Schema: map[string]*schema.Schema{

			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fswid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemIpv6NeighborCacheRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("fswid")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemIpv6NeighborCache: type error")
	}

	o, err := c.ReadSystemIpv6NeighborCache(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemIpv6NeighborCache: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemIpv6NeighborCache(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemIpv6NeighborCache from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemIpv6NeighborCacheInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemIpv6NeighborCacheMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemIpv6NeighborCacheId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemIpv6NeighborCacheIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemIpv6NeighborCache(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("interface", dataSourceFlattenSystemIpv6NeighborCacheInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("mac", dataSourceFlattenSystemIpv6NeighborCacheMac(o["mac"], d, "mac")); err != nil {
		if !fortiAPIPatch(o["mac"]) {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("fswid", dataSourceFlattenSystemIpv6NeighborCacheId(o["id"], d, "fswid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fswid: %v", err)
		}
	}

	if err = d.Set("ipv6", dataSourceFlattenSystemIpv6NeighborCacheIpv6(o["ipv6"], d, "ipv6")); err != nil {
		if !fortiAPIPatch(o["ipv6"]) {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemIpv6NeighborCacheFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
