// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Vdom dns configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVdomDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomDnsUpdate,
		Read:   resourceSystemVdomDnsRead,
		Update: resourceSystemVdomDnsUpdate,
		Delete: resourceSystemVdomDnsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdom_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary": &schema.Schema{
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
			"secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemVdomDnsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdomDns(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomDns resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVdomDns(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomDns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdomDns")
	}

	return resourceSystemVdomDnsRead(d, m)
}

func resourceSystemVdomDnsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdomDns(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemVdomDns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVdomDns(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemVdomDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomDnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemVdomDns(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomDns(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomDns resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomDnsVdomDns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsPrimary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsIp6Secondary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsIp6Primary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsSecondary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemVdomDns(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("vdom_dns", flattenSystemVdomDnsVdomDns(o["vdom-dns"], d, "vdom_dns", sv)); err != nil {
		if !fortiAPIPatch(o["vdom-dns"]) {
			return fmt.Errorf("Error reading vdom_dns: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemVdomDnsSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("primary", flattenSystemVdomDnsPrimary(o["primary"], d, "primary", sv)); err != nil {
		if !fortiAPIPatch(o["primary"]) {
			return fmt.Errorf("Error reading primary: %v", err)
		}
	}

	if err = d.Set("ip6_secondary", flattenSystemVdomDnsIp6Secondary(o["ip6-secondary"], d, "ip6_secondary", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-secondary"]) {
			return fmt.Errorf("Error reading ip6_secondary: %v", err)
		}
	}

	if err = d.Set("ip6_primary", flattenSystemVdomDnsIp6Primary(o["ip6-primary"], d, "ip6_primary", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-primary"]) {
			return fmt.Errorf("Error reading ip6_primary: %v", err)
		}
	}

	if err = d.Set("secondary", flattenSystemVdomDnsSecondary(o["secondary"], d, "secondary", sv)); err != nil {
		if !fortiAPIPatch(o["secondary"]) {
			return fmt.Errorf("Error reading secondary: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomDnsFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemVdomDnsVdomDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsPrimary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsIp6Secondary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsIp6Primary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsSecondary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdomDns(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("vdom_dns"); ok {
		if setArgNil {
			obj["vdom-dns"] = nil
		} else {

			t, err := expandSystemVdomDnsVdomDns(d, v, "vdom_dns", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vdom-dns"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {

			t, err := expandSystemVdomDnsSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("primary"); ok {
		if setArgNil {
			obj["primary"] = nil
		} else {

			t, err := expandSystemVdomDnsPrimary(d, v, "primary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["primary"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip6_secondary"); ok {
		if setArgNil {
			obj["ip6-secondary"] = nil
		} else {

			t, err := expandSystemVdomDnsIp6Secondary(d, v, "ip6_secondary", sv)
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

			t, err := expandSystemVdomDnsIp6Primary(d, v, "ip6_primary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip6-primary"] = t
			}
		}
	}

	if v, ok := d.GetOk("secondary"); ok {
		if setArgNil {
			obj["secondary"] = nil
		} else {

			t, err := expandSystemVdomDnsSecondary(d, v, "secondary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["secondary"] = t
			}
		}
	}

	return &obj, nil
}
