// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure DNS.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemDns() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemDnsRead,
		Schema: map[string]*schema.Schema{

			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"primary": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_cache_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cache_notfound_responses": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_secondary": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_primary": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemDnsRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemDns"

	o, err := c.ReadSystemDns(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemDns: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemDns(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemDns from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemDnsDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDnsCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDnsCacheLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsCacheNotfoundResponses(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsIp6Secondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsIp6Primary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsSecondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemDns(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("domain", dataSourceFlattenSystemDnsDomain(o["domain"], d, "domain")); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("dns_cache_ttl", dataSourceFlattenSystemDnsDnsCacheTtl(o["dns-cache-ttl"], d, "dns_cache_ttl")); err != nil {
		if !fortiAPIPatch(o["dns-cache-ttl"]) {
			return fmt.Errorf("Error reading dns_cache_ttl: %v", err)
		}
	}

	if err = d.Set("primary", dataSourceFlattenSystemDnsPrimary(o["primary"], d, "primary")); err != nil {
		if !fortiAPIPatch(o["primary"]) {
			return fmt.Errorf("Error reading primary: %v", err)
		}
	}

	if err = d.Set("dns_cache_limit", dataSourceFlattenSystemDnsDnsCacheLimit(o["dns-cache-limit"], d, "dns_cache_limit")); err != nil {
		if !fortiAPIPatch(o["dns-cache-limit"]) {
			return fmt.Errorf("Error reading dns_cache_limit: %v", err)
		}
	}

	if err = d.Set("cache_notfound_responses", dataSourceFlattenSystemDnsCacheNotfoundResponses(o["cache-notfound-responses"], d, "cache_notfound_responses")); err != nil {
		if !fortiAPIPatch(o["cache-notfound-responses"]) {
			return fmt.Errorf("Error reading cache_notfound_responses: %v", err)
		}
	}

	if err = d.Set("ip6_secondary", dataSourceFlattenSystemDnsIp6Secondary(o["ip6-secondary"], d, "ip6_secondary")); err != nil {
		if !fortiAPIPatch(o["ip6-secondary"]) {
			return fmt.Errorf("Error reading ip6_secondary: %v", err)
		}
	}

	if err = d.Set("ip6_primary", dataSourceFlattenSystemDnsIp6Primary(o["ip6-primary"], d, "ip6_primary")); err != nil {
		if !fortiAPIPatch(o["ip6-primary"]) {
			return fmt.Errorf("Error reading ip6_primary: %v", err)
		}
	}

	if err = d.Set("source_ip", dataSourceFlattenSystemDnsSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("secondary", dataSourceFlattenSystemDnsSecondary(o["secondary"], d, "secondary")); err != nil {
		if !fortiAPIPatch(o["secondary"]) {
			return fmt.Errorf("Error reading secondary: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemDnsFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
