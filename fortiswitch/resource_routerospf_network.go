// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Enable OSPF on an IP network.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterospfNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterospfNetworkCreate,
		Read:   resourceRouterospfNetworkRead,
		Update: resourceRouterospfNetworkUpdate,
		Delete: resourceRouterospfNetworkDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fswid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"area": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterospfNetworkCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterospfNetwork(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterospfNetwork resource while getting object: %v", err)
	}

	o, err := c.CreateRouterospfNetwork(obj)

	if err != nil {
		return fmt.Errorf("Error creating RouterospfNetwork resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterospfNetwork")
	}

	return resourceRouterospfNetworkRead(d, m)
}

func resourceRouterospfNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterospfNetwork(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterospfNetwork resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterospfNetwork(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterospfNetwork resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterospfNetwork")
	}

	return resourceRouterospfNetworkRead(d, m)
}

func resourceRouterospfNetworkDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterospfNetwork(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting RouterospfNetwork resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterospfNetworkRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterospfNetwork(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterospfNetwork resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterospfNetwork(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterospfNetwork resource from API: %v", err)
	}
	return nil
}

func flattenRouterospfNetworkPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterospfNetworkId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfNetworkArea(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterospfNetwork(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("prefix", flattenRouterospfNetworkPrefix(o["prefix"], d, "prefix", sv)); err != nil {
		if !fortiAPIPatch(o["prefix"]) {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("fswid", flattenRouterospfNetworkId(o["id"], d, "fswid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fswid: %v", err)
		}
	}

	if err = d.Set("area", flattenRouterospfNetworkArea(o["area"], d, "area", sv)); err != nil {
		if !fortiAPIPatch(o["area"]) {
			return fmt.Errorf("Error reading area: %v", err)
		}
	}

	return nil
}

func flattenRouterospfNetworkFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandRouterospfNetworkPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfNetworkId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfNetworkArea(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterospfNetwork(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("prefix"); ok {

		t, err := expandRouterospfNetworkPrefix(d, v, "prefix", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("fswid"); ok {

		t, err := expandRouterospfNetworkId(d, v, "fswid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("area"); ok {

		t, err := expandRouterospfNetworkArea(d, v, "area", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["area"] = t
		}
	}

	return &obj, nil
}
