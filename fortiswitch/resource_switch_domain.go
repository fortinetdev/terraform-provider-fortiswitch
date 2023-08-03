// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Switch forwarding domains.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchDomainCreate,
		Read:   resourceSwitchDomainRead,
		Update: resourceSwitchDomainUpdate,
		Delete: resourceSwitchDomainDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"inter_front_panel_traffic": &schema.Schema{
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
			"priority": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"ha_block": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_l2_clear_on_role_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcluster_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchDomainCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchDomain(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchDomain resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchDomain(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchDomain resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchDomain")
	}

	return resourceSwitchDomainRead(d, m)
}

func resourceSwitchDomainUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchDomain(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchDomain resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchDomain(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchDomain resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchDomain")
	}

	return resourceSwitchDomainRead(d, m)
}

func resourceSwitchDomainDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchDomain(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchDomain resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchDomainRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchDomain(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchDomain resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchDomain(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchDomain resource from API: %v", err)
	}
	return nil
}

func flattenSwitchDomainInterFrontPanelTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchDomainName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchDomainPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchDomainHaBlock(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchDomainHaL2ClearOnRoleChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchDomainVclusterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchDomain(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("inter_front_panel_traffic", flattenSwitchDomainInterFrontPanelTraffic(o["inter-front-panel-traffic"], d, "inter_front_panel_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["inter-front-panel-traffic"]) {
			return fmt.Errorf("Error reading inter_front_panel_traffic: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchDomainName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("priority", flattenSwitchDomainPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("ha_block", flattenSwitchDomainHaBlock(o["ha-block"], d, "ha_block", sv)); err != nil {
		if !fortiAPIPatch(o["ha-block"]) {
			return fmt.Errorf("Error reading ha_block: %v", err)
		}
	}

	if err = d.Set("ha_l2_clear_on_role_change", flattenSwitchDomainHaL2ClearOnRoleChange(o["ha-L2-clear-on-role-change"], d, "ha_l2_clear_on_role_change", sv)); err != nil {
		if !fortiAPIPatch(o["ha-L2-clear-on-role-change"]) {
			return fmt.Errorf("Error reading ha_l2_clear_on_role_change: %v", err)
		}
	}

	if err = d.Set("vcluster_id", flattenSwitchDomainVclusterId(o["vcluster-id"], d, "vcluster_id", sv)); err != nil {
		if !fortiAPIPatch(o["vcluster-id"]) {
			return fmt.Errorf("Error reading vcluster_id: %v", err)
		}
	}

	return nil
}

func flattenSwitchDomainFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchDomainInterFrontPanelTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchDomainName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchDomainPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchDomainHaBlock(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchDomainHaL2ClearOnRoleChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchDomainVclusterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchDomain(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("inter_front_panel_traffic"); ok {

		t, err := expandSwitchDomainInterFrontPanelTraffic(d, v, "inter_front_panel_traffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inter-front-panel-traffic"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchDomainName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("priority"); ok {

		t, err := expandSwitchDomainPriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("ha_block"); ok {

		t, err := expandSwitchDomainHaBlock(d, v, "ha_block", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-block"] = t
		}
	}

	if v, ok := d.GetOk("ha_l2_clear_on_role_change"); ok {

		t, err := expandSwitchDomainHaL2ClearOnRoleChange(d, v, "ha_l2_clear_on_role_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-L2-clear-on-role-change"] = t
		}
	}

	if v, ok := d.GetOk("vcluster_id"); ok {

		t, err := expandSwitchDomainVclusterId(d, v, "vcluster_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcluster-id"] = t
		}
	}

	return &obj, nil
}
