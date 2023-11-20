// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: OSPF interface configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterospfInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterospfInterfaceCreate,
		Read:   resourceRouterospfInterfaceRead,
		Update: resourceRouterospfInterfaceUpdate,
		Delete: resourceRouterospfInterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"authentication_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 8),
				Optional:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"dead_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hello_multiplier": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transmit_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ucast_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retransmit_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cost": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hello_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"md5_keys": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 17),
							Optional:     true,
						},
					},
				},
			},
			"mtu_ignore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceRouterospfInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterospfInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterospfInterface resource while getting object: %v", err)
	}

	o, err := c.CreateRouterospfInterface(obj)

	if err != nil {
		return fmt.Errorf("Error creating RouterospfInterface resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterospfInterface")
	}

	return resourceRouterospfInterfaceRead(d, m)
}

func resourceRouterospfInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterospfInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterospfInterface resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterospfInterface(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterospfInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterospfInterface")
	}

	return resourceRouterospfInterfaceRead(d, m)
}

func resourceRouterospfInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterospfInterface(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting RouterospfInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterospfInterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterospfInterface(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterospfInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterospfInterface(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterospfInterface resource from API: %v", err)
	}
	return nil
}

func flattenRouterospfInterfacePriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceAuthenticationKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceHelloMultiplier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceUcastTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceMtu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceMd5Keys(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterospfInterfaceMd5KeysId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {

			tmp["key"] = flattenRouterospfInterfaceMd5KeysKey(i["key"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterospfInterfaceMd5KeysId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceMd5KeysKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfInterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterospfInterface(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("priority", flattenRouterospfInterfacePriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("authentication_key", flattenRouterospfInterfaceAuthenticationKey(o["authentication-key"], d, "authentication_key", sv)); err != nil {
		if !fortiAPIPatch(o["authentication-key"]) {
			return fmt.Errorf("Error reading authentication_key: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterospfInterfaceName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("dead_interval", flattenRouterospfInterfaceDeadInterval(o["dead-interval"], d, "dead_interval", sv)); err != nil {
		if !fortiAPIPatch(o["dead-interval"]) {
			return fmt.Errorf("Error reading dead_interval: %v", err)
		}
	}

	if err = d.Set("hello_multiplier", flattenRouterospfInterfaceHelloMultiplier(o["hello-multiplier"], d, "hello_multiplier", sv)); err != nil {
		if !fortiAPIPatch(o["hello-multiplier"]) {
			return fmt.Errorf("Error reading hello_multiplier: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterospfInterfaceBfd(o["bfd"], d, "bfd", sv)); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("transmit_delay", flattenRouterospfInterfaceTransmitDelay(o["transmit-delay"], d, "transmit_delay", sv)); err != nil {
		if !fortiAPIPatch(o["transmit-delay"]) {
			return fmt.Errorf("Error reading transmit_delay: %v", err)
		}
	}

	if err = d.Set("ucast_ttl", flattenRouterospfInterfaceUcastTtl(o["ucast-ttl"], d, "ucast_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["ucast-ttl"]) {
			return fmt.Errorf("Error reading ucast_ttl: %v", err)
		}
	}

	if err = d.Set("mtu", flattenRouterospfInterfaceMtu(o["mtu"], d, "mtu", sv)); err != nil {
		if !fortiAPIPatch(o["mtu"]) {
			return fmt.Errorf("Error reading mtu: %v", err)
		}
	}

	if err = d.Set("retransmit_interval", flattenRouterospfInterfaceRetransmitInterval(o["retransmit-interval"], d, "retransmit_interval", sv)); err != nil {
		if !fortiAPIPatch(o["retransmit-interval"]) {
			return fmt.Errorf("Error reading retransmit_interval: %v", err)
		}
	}

	if err = d.Set("authentication", flattenRouterospfInterfaceAuthentication(o["authentication"], d, "authentication", sv)); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("cost", flattenRouterospfInterfaceCost(o["cost"], d, "cost", sv)); err != nil {
		if !fortiAPIPatch(o["cost"]) {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	if err = d.Set("hello_interval", flattenRouterospfInterfaceHelloInterval(o["hello-interval"], d, "hello_interval", sv)); err != nil {
		if !fortiAPIPatch(o["hello-interval"]) {
			return fmt.Errorf("Error reading hello_interval: %v", err)
		}
	}

	if err = d.Set("ttl", flattenRouterospfInterfaceTtl(o["ttl"], d, "ttl", sv)); err != nil {
		if !fortiAPIPatch(o["ttl"]) {
			return fmt.Errorf("Error reading ttl: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("md5_keys", flattenRouterospfInterfaceMd5Keys(o["md5-keys"], d, "md5_keys", sv)); err != nil {
			if !fortiAPIPatch(o["md5-keys"]) {
				return fmt.Errorf("Error reading md5_keys: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("md5_keys"); ok {
			if err = d.Set("md5_keys", flattenRouterospfInterfaceMd5Keys(o["md5-keys"], d, "md5_keys", sv)); err != nil {
				if !fortiAPIPatch(o["md5-keys"]) {
					return fmt.Errorf("Error reading md5_keys: %v", err)
				}
			}
		}
	}

	if err = d.Set("mtu_ignore", flattenRouterospfInterfaceMtuIgnore(o["mtu-ignore"], d, "mtu_ignore", sv)); err != nil {
		if !fortiAPIPatch(o["mtu-ignore"]) {
			return fmt.Errorf("Error reading mtu_ignore: %v", err)
		}
	}

	return nil
}

func flattenRouterospfInterfaceFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandRouterospfInterfacePriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceAuthenticationKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceDeadInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceHelloMultiplier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceTransmitDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceUcastTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceMtu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceRetransmitInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceMd5Keys(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterospfInterfaceMd5KeysId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["key"], _ = expandRouterospfInterfaceMd5KeysKey(d, i["key"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterospfInterfaceMd5KeysId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceMd5KeysKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfInterfaceMtuIgnore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterospfInterface(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("priority"); ok {

		t, err := expandRouterospfInterfacePriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("authentication_key"); ok {

		t, err := expandRouterospfInterfaceAuthenticationKey(d, v, "authentication_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication-key"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandRouterospfInterfaceName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("dead_interval"); ok {

		t, err := expandRouterospfInterfaceDeadInterval(d, v, "dead_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dead-interval"] = t
		}
	}

	if v, ok := d.GetOk("hello_multiplier"); ok {

		t, err := expandRouterospfInterfaceHelloMultiplier(d, v, "hello_multiplier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-multiplier"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok {

		t, err := expandRouterospfInterfaceBfd(d, v, "bfd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("transmit_delay"); ok {

		t, err := expandRouterospfInterfaceTransmitDelay(d, v, "transmit_delay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transmit-delay"] = t
		}
	}

	if v, ok := d.GetOk("ucast_ttl"); ok {

		t, err := expandRouterospfInterfaceUcastTtl(d, v, "ucast_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ucast-ttl"] = t
		}
	}

	if v, ok := d.GetOk("mtu"); ok {

		t, err := expandRouterospfInterfaceMtu(d, v, "mtu", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu"] = t
		}
	}

	if v, ok := d.GetOk("retransmit_interval"); ok {

		t, err := expandRouterospfInterfaceRetransmitInterval(d, v, "retransmit_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retransmit-interval"] = t
		}
	}

	if v, ok := d.GetOk("authentication"); ok {

		t, err := expandRouterospfInterfaceAuthentication(d, v, "authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("cost"); ok {

		t, err := expandRouterospfInterfaceCost(d, v, "cost", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	if v, ok := d.GetOk("hello_interval"); ok {

		t, err := expandRouterospfInterfaceHelloInterval(d, v, "hello_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-interval"] = t
		}
	}

	if v, ok := d.GetOk("ttl"); ok {

		t, err := expandRouterospfInterfaceTtl(d, v, "ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttl"] = t
		}
	}

	if v, ok := d.GetOk("md5_keys"); ok || d.HasChange("md5_keys") {

		t, err := expandRouterospfInterfaceMd5Keys(d, v, "md5_keys", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["md5-keys"] = t
		}
	}

	if v, ok := d.GetOk("mtu_ignore"); ok {

		t, err := expandRouterospfInterfaceMtuIgnore(d, v, "mtu_ignore", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu-ignore"] = t
		}
	}

	return &obj, nil
}
