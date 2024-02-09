// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Redistribute configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterospf6Redistribute() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterospf6RedistributeCreate,
		Read:   resourceRouterospf6RedistributeRead,
		Update: resourceRouterospf6RedistributeUpdate,
		Delete: resourceRouterospf6RedistributeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metric": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"routemap": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceRouterospf6RedistributeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterospf6Redistribute(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating Routerospf6Redistribute resource while getting object: %v", err)
	}

	o, err := c.CreateRouterospf6Redistribute(obj)

	if err != nil {
		return fmt.Errorf("Error creating Routerospf6Redistribute resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("Routerospf6Redistribute")
	}

	return resourceRouterospf6RedistributeRead(d, m)
}

func resourceRouterospf6RedistributeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterospf6Redistribute(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating Routerospf6Redistribute resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterospf6Redistribute(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Routerospf6Redistribute resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("Routerospf6Redistribute")
	}

	return resourceRouterospf6RedistributeRead(d, m)
}

func resourceRouterospf6RedistributeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterospf6Redistribute(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting Routerospf6Redistribute resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterospf6RedistributeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterospf6Redistribute(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Routerospf6Redistribute resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterospf6Redistribute(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading Routerospf6Redistribute resource from API: %v", err)
	}
	return nil
}

func flattenRouterospf6RedistributeStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospf6RedistributeMetricType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospf6RedistributeMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospf6RedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospf6RedistributeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterospf6Redistribute(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenRouterospf6RedistributeStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("metric_type", flattenRouterospf6RedistributeMetricType(o["metric-type"], d, "metric_type", sv)); err != nil {
		if !fortiAPIPatch(o["metric-type"]) {
			return fmt.Errorf("Error reading metric_type: %v", err)
		}
	}

	if err = d.Set("metric", flattenRouterospf6RedistributeMetric(o["metric"], d, "metric", sv)); err != nil {
		if !fortiAPIPatch(o["metric"]) {
			return fmt.Errorf("Error reading metric: %v", err)
		}
	}

	if err = d.Set("routemap", flattenRouterospf6RedistributeRoutemap(o["routemap"], d, "routemap", sv)); err != nil {
		if !fortiAPIPatch(o["routemap"]) {
			return fmt.Errorf("Error reading routemap: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterospf6RedistributeName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenRouterospf6RedistributeFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandRouterospf6RedistributeStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospf6RedistributeMetricType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospf6RedistributeMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospf6RedistributeRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospf6RedistributeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterospf6Redistribute(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {

		t, err := expandRouterospf6RedistributeStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("metric_type"); ok {

		t, err := expandRouterospf6RedistributeMetricType(d, v, "metric_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["metric-type"] = t
		}
	}

	if v, ok := d.GetOk("metric"); ok {

		t, err := expandRouterospf6RedistributeMetric(d, v, "metric", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["metric"] = t
		}
	}

	if v, ok := d.GetOk("routemap"); ok {

		t, err := expandRouterospf6RedistributeRoutemap(d, v, "routemap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["routemap"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandRouterospf6RedistributeName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
