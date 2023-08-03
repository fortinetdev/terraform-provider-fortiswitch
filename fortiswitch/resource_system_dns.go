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

func resourceSystemDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDnsUpdate,
		Read:   resourceSystemDnsRead,
		Update: resourceSystemDnsUpdate,
		Delete: resourceSystemDnsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"dns_cache_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),
				Optional:     true,
				Computed:     true,
			},
			"primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_cache_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cache_notfound_responses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDnsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemDns(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDns(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemDns")
	}

	return resourceSystemDnsRead(d, m)
}

func resourceSystemDnsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemDns(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDns(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemDns(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDns(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource from API: %v", err)
	}
	return nil
}

func flattenSystemDnsDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDnsCacheTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsPrimary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDnsCacheLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsCacheNotfoundResponses(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsIp6Secondary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsIp6Primary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsSecondary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDns(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("domain", flattenSystemDnsDomain(o["domain"], d, "domain", sv)); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("dns_cache_ttl", flattenSystemDnsDnsCacheTtl(o["dns-cache-ttl"], d, "dns_cache_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["dns-cache-ttl"]) {
			return fmt.Errorf("Error reading dns_cache_ttl: %v", err)
		}
	}

	if err = d.Set("primary", flattenSystemDnsPrimary(o["primary"], d, "primary", sv)); err != nil {
		if !fortiAPIPatch(o["primary"]) {
			return fmt.Errorf("Error reading primary: %v", err)
		}
	}

	if err = d.Set("dns_cache_limit", flattenSystemDnsDnsCacheLimit(o["dns-cache-limit"], d, "dns_cache_limit", sv)); err != nil {
		if !fortiAPIPatch(o["dns-cache-limit"]) {
			return fmt.Errorf("Error reading dns_cache_limit: %v", err)
		}
	}

	if err = d.Set("cache_notfound_responses", flattenSystemDnsCacheNotfoundResponses(o["cache-notfound-responses"], d, "cache_notfound_responses", sv)); err != nil {
		if !fortiAPIPatch(o["cache-notfound-responses"]) {
			return fmt.Errorf("Error reading cache_notfound_responses: %v", err)
		}
	}

	if err = d.Set("ip6_secondary", flattenSystemDnsIp6Secondary(o["ip6-secondary"], d, "ip6_secondary", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-secondary"]) {
			return fmt.Errorf("Error reading ip6_secondary: %v", err)
		}
	}

	if err = d.Set("ip6_primary", flattenSystemDnsIp6Primary(o["ip6-primary"], d, "ip6_primary", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-primary"]) {
			return fmt.Errorf("Error reading ip6_primary: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemDnsSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("secondary", flattenSystemDnsSecondary(o["secondary"], d, "secondary", sv)); err != nil {
		if !fortiAPIPatch(o["secondary"]) {
			return fmt.Errorf("Error reading secondary: %v", err)
		}
	}

	return nil
}

func flattenSystemDnsFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemDnsDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDnsCacheTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsPrimary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDnsCacheLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsCacheNotfoundResponses(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsIp6Secondary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsIp6Primary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsSecondary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDns(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("domain"); ok {
		if setArgNil {
			obj["domain"] = nil
		} else {

			t, err := expandSystemDnsDomain(d, v, "domain", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["domain"] = t
			}
		}
	}

	if v, ok := d.GetOk("dns_cache_ttl"); ok {
		if setArgNil {
			obj["dns-cache-ttl"] = nil
		} else {

			t, err := expandSystemDnsDnsCacheTtl(d, v, "dns_cache_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dns-cache-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOk("primary"); ok {
		if setArgNil {
			obj["primary"] = nil
		} else {

			t, err := expandSystemDnsPrimary(d, v, "primary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["primary"] = t
			}
		}
	}

	if v, ok := d.GetOk("dns_cache_limit"); ok {
		if setArgNil {
			obj["dns-cache-limit"] = nil
		} else {

			t, err := expandSystemDnsDnsCacheLimit(d, v, "dns_cache_limit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dns-cache-limit"] = t
			}
		}
	}

	if v, ok := d.GetOk("cache_notfound_responses"); ok {
		if setArgNil {
			obj["cache-notfound-responses"] = nil
		} else {

			t, err := expandSystemDnsCacheNotfoundResponses(d, v, "cache_notfound_responses", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cache-notfound-responses"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip6_secondary"); ok {
		if setArgNil {
			obj["ip6-secondary"] = nil
		} else {

			t, err := expandSystemDnsIp6Secondary(d, v, "ip6_secondary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip6-secondary"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip6_primary"); ok {
		if setArgNil {
			obj["ip6-primary"] = nil
		} else {

			t, err := expandSystemDnsIp6Primary(d, v, "ip6_primary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip6-primary"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {

			t, err := expandSystemDnsSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("secondary"); ok {
		if setArgNil {
			obj["secondary"] = nil
		} else {

			t, err := expandSystemDnsSecondary(d, v, "secondary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["secondary"] = t
			}
		}
	}

	return &obj, nil
}
