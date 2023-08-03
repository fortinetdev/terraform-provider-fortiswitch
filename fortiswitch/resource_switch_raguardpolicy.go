// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: IPV6 RA Guard policy.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchRaguardPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchRaguardPolicyCreate,
		Read:   resourceSwitchRaguardPolicyRead,
		Update: resourceSwitchRaguardPolicyUpdate,
		Delete: resourceSwitchRaguardPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"max_router_preference": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_hop_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"max_hop_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"managed_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_src_addr": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"device_role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"other_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_prefix": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchRaguardPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchRaguardPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchRaguardPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchRaguardPolicy(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchRaguardPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchRaguardPolicy")
	}

	return resourceSwitchRaguardPolicyRead(d, m)
}

func resourceSwitchRaguardPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchRaguardPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchRaguardPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchRaguardPolicy(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchRaguardPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchRaguardPolicy")
	}

	return resourceSwitchRaguardPolicyRead(d, m)
}

func resourceSwitchRaguardPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchRaguardPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchRaguardPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchRaguardPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchRaguardPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchRaguardPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchRaguardPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchRaguardPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchRaguardPolicyMaxRouterPreference(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchRaguardPolicyMinHopLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchRaguardPolicyMaxHopLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchRaguardPolicyManagedFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchRaguardPolicyMatchSrcAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchRaguardPolicyDeviceRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchRaguardPolicyOtherFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchRaguardPolicyMatchPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchRaguardPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchRaguardPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("max_router_preference", flattenSwitchRaguardPolicyMaxRouterPreference(o["max-router-preference"], d, "max_router_preference", sv)); err != nil {
		if !fortiAPIPatch(o["max-router-preference"]) {
			return fmt.Errorf("Error reading max_router_preference: %v", err)
		}
	}

	if err = d.Set("min_hop_limit", flattenSwitchRaguardPolicyMinHopLimit(o["min-hop-limit"], d, "min_hop_limit", sv)); err != nil {
		if !fortiAPIPatch(o["min-hop-limit"]) {
			return fmt.Errorf("Error reading min_hop_limit: %v", err)
		}
	}

	if err = d.Set("max_hop_limit", flattenSwitchRaguardPolicyMaxHopLimit(o["max-hop-limit"], d, "max_hop_limit", sv)); err != nil {
		if !fortiAPIPatch(o["max-hop-limit"]) {
			return fmt.Errorf("Error reading max_hop_limit: %v", err)
		}
	}

	if err = d.Set("managed_flag", flattenSwitchRaguardPolicyManagedFlag(o["managed-flag"], d, "managed_flag", sv)); err != nil {
		if !fortiAPIPatch(o["managed-flag"]) {
			return fmt.Errorf("Error reading managed_flag: %v", err)
		}
	}

	if err = d.Set("match_src_addr", flattenSwitchRaguardPolicyMatchSrcAddr(o["match-src-addr"], d, "match_src_addr", sv)); err != nil {
		if !fortiAPIPatch(o["match-src-addr"]) {
			return fmt.Errorf("Error reading match_src_addr: %v", err)
		}
	}

	if err = d.Set("device_role", flattenSwitchRaguardPolicyDeviceRole(o["device-role"], d, "device_role", sv)); err != nil {
		if !fortiAPIPatch(o["device-role"]) {
			return fmt.Errorf("Error reading device_role: %v", err)
		}
	}

	if err = d.Set("other_flag", flattenSwitchRaguardPolicyOtherFlag(o["other-flag"], d, "other_flag", sv)); err != nil {
		if !fortiAPIPatch(o["other-flag"]) {
			return fmt.Errorf("Error reading other_flag: %v", err)
		}
	}

	if err = d.Set("match_prefix", flattenSwitchRaguardPolicyMatchPrefix(o["match-prefix"], d, "match_prefix", sv)); err != nil {
		if !fortiAPIPatch(o["match-prefix"]) {
			return fmt.Errorf("Error reading match_prefix: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchRaguardPolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSwitchRaguardPolicyFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchRaguardPolicyMaxRouterPreference(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchRaguardPolicyMinHopLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchRaguardPolicyMaxHopLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchRaguardPolicyManagedFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchRaguardPolicyMatchSrcAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchRaguardPolicyDeviceRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchRaguardPolicyOtherFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchRaguardPolicyMatchPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchRaguardPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchRaguardPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_router_preference"); ok {

		t, err := expandSwitchRaguardPolicyMaxRouterPreference(d, v, "max_router_preference", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-router-preference"] = t
		}
	}

	if v, ok := d.GetOkExists("min_hop_limit"); ok {

		t, err := expandSwitchRaguardPolicyMinHopLimit(d, v, "min_hop_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-hop-limit"] = t
		}
	}

	if v, ok := d.GetOkExists("max_hop_limit"); ok {

		t, err := expandSwitchRaguardPolicyMaxHopLimit(d, v, "max_hop_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-hop-limit"] = t
		}
	}

	if v, ok := d.GetOk("managed_flag"); ok {

		t, err := expandSwitchRaguardPolicyManagedFlag(d, v, "managed_flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["managed-flag"] = t
		}
	}

	if v, ok := d.GetOk("match_src_addr"); ok {

		t, err := expandSwitchRaguardPolicyMatchSrcAddr(d, v, "match_src_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-src-addr"] = t
		}
	}

	if v, ok := d.GetOk("device_role"); ok {

		t, err := expandSwitchRaguardPolicyDeviceRole(d, v, "device_role", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-role"] = t
		}
	}

	if v, ok := d.GetOk("other_flag"); ok {

		t, err := expandSwitchRaguardPolicyOtherFlag(d, v, "other_flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["other-flag"] = t
		}
	}

	if v, ok := d.GetOk("match_prefix"); ok {

		t, err := expandSwitchRaguardPolicyMatchPrefix(d, v, "match_prefix", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-prefix"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchRaguardPolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
