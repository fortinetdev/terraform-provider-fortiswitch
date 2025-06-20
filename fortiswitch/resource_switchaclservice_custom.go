// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Custom service configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchAclServiceCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchAclServiceCustomCreate,
		Read:   resourceSwitchAclServiceCustomRead,
		Update: resourceSwitchAclServiceCustomUpdate,
		Delete: resourceSwitchAclServiceCustomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"protocol_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
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
			"color": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"icmptype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sctp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_portmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_portmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sctp_portmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmpcode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchAclServiceCustomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAclServiceCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchAclServiceCustom resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchAclServiceCustom(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchAclServiceCustom resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchAclServiceCustom")
	}

	return resourceSwitchAclServiceCustomRead(d, m)
}

func resourceSwitchAclServiceCustomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchAclServiceCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAclServiceCustom resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchAclServiceCustom(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchAclServiceCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchAclServiceCustom")
	}

	return resourceSwitchAclServiceCustomRead(d, m)
}

func resourceSwitchAclServiceCustomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchAclServiceCustom(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchAclServiceCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchAclServiceCustomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchAclServiceCustom(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAclServiceCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchAclServiceCustom(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchAclServiceCustom resource from API: %v", err)
	}
	return nil
}

func flattenSwitchAclServiceCustomComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclServiceCustomProtocolNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclServiceCustomProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclServiceCustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclServiceCustomColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclServiceCustomIcmptype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclServiceCustomSctpPortrange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclServiceCustomTcpPortmask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclServiceCustomUdpPortmask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclServiceCustomSctpPortmask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclServiceCustomUdpPortrange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclServiceCustomIcmpcode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchAclServiceCustomTcpPortrange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchAclServiceCustom(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("comment", flattenSwitchAclServiceCustomComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("protocol_number", flattenSwitchAclServiceCustomProtocolNumber(o["protocol-number"], d, "protocol_number", sv)); err != nil {
		if !fortiAPIPatch(o["protocol-number"]) {
			return fmt.Errorf("Error reading protocol_number: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSwitchAclServiceCustomProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchAclServiceCustomName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("color", flattenSwitchAclServiceCustomColor(o["color"], d, "color", sv)); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("icmptype", flattenSwitchAclServiceCustomIcmptype(o["icmptype"], d, "icmptype", sv)); err != nil {
		if !fortiAPIPatch(o["icmptype"]) {
			return fmt.Errorf("Error reading icmptype: %v", err)
		}
	}

	if err = d.Set("sctp_portrange", flattenSwitchAclServiceCustomSctpPortrange(o["sctp-portrange"], d, "sctp_portrange", sv)); err != nil {
		if !fortiAPIPatch(o["sctp-portrange"]) {
			return fmt.Errorf("Error reading sctp_portrange: %v", err)
		}
	}

	if err = d.Set("tcp_portmask", flattenSwitchAclServiceCustomTcpPortmask(o["tcp-portmask"], d, "tcp_portmask", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-portmask"]) {
			return fmt.Errorf("Error reading tcp_portmask: %v", err)
		}
	}

	if err = d.Set("udp_portmask", flattenSwitchAclServiceCustomUdpPortmask(o["udp-portmask"], d, "udp_portmask", sv)); err != nil {
		if !fortiAPIPatch(o["udp-portmask"]) {
			return fmt.Errorf("Error reading udp_portmask: %v", err)
		}
	}

	if err = d.Set("sctp_portmask", flattenSwitchAclServiceCustomSctpPortmask(o["sctp-portmask"], d, "sctp_portmask", sv)); err != nil {
		if !fortiAPIPatch(o["sctp-portmask"]) {
			return fmt.Errorf("Error reading sctp_portmask: %v", err)
		}
	}

	if err = d.Set("udp_portrange", flattenSwitchAclServiceCustomUdpPortrange(o["udp-portrange"], d, "udp_portrange", sv)); err != nil {
		if !fortiAPIPatch(o["udp-portrange"]) {
			return fmt.Errorf("Error reading udp_portrange: %v", err)
		}
	}

	if err = d.Set("icmpcode", flattenSwitchAclServiceCustomIcmpcode(o["icmpcode"], d, "icmpcode", sv)); err != nil {
		if !fortiAPIPatch(o["icmpcode"]) {
			return fmt.Errorf("Error reading icmpcode: %v", err)
		}
	}

	if err = d.Set("tcp_portrange", flattenSwitchAclServiceCustomTcpPortrange(o["tcp-portrange"], d, "tcp_portrange", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-portrange"]) {
			return fmt.Errorf("Error reading tcp_portrange: %v", err)
		}
	}

	return nil
}

func flattenSwitchAclServiceCustomFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchAclServiceCustomComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclServiceCustomProtocolNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclServiceCustomProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclServiceCustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclServiceCustomColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclServiceCustomIcmptype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclServiceCustomSctpPortrange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclServiceCustomTcpPortmask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclServiceCustomUdpPortmask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclServiceCustomSctpPortmask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclServiceCustomUdpPortrange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclServiceCustomIcmpcode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchAclServiceCustomTcpPortrange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchAclServiceCustom(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandSwitchAclServiceCustomComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("protocol_number"); ok {

		t, err := expandSwitchAclServiceCustomProtocolNumber(d, v, "protocol_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol-number"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {

		t, err := expandSwitchAclServiceCustomProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchAclServiceCustomName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("color"); ok {

		t, err := expandSwitchAclServiceCustomColor(d, v, "color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("icmptype"); ok {

		t, err := expandSwitchAclServiceCustomIcmptype(d, v, "icmptype", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmptype"] = t
		}
	}

	if v, ok := d.GetOk("sctp_portrange"); ok {

		t, err := expandSwitchAclServiceCustomSctpPortrange(d, v, "sctp_portrange", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-portrange"] = t
		}
	}

	if v, ok := d.GetOk("tcp_portmask"); ok {

		t, err := expandSwitchAclServiceCustomTcpPortmask(d, v, "tcp_portmask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-portmask"] = t
		}
	}

	if v, ok := d.GetOk("udp_portmask"); ok {

		t, err := expandSwitchAclServiceCustomUdpPortmask(d, v, "udp_portmask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-portmask"] = t
		}
	}

	if v, ok := d.GetOk("sctp_portmask"); ok {

		t, err := expandSwitchAclServiceCustomSctpPortmask(d, v, "sctp_portmask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-portmask"] = t
		}
	}

	if v, ok := d.GetOk("udp_portrange"); ok {

		t, err := expandSwitchAclServiceCustomUdpPortrange(d, v, "udp_portrange", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-portrange"] = t
		}
	}

	if v, ok := d.GetOk("icmpcode"); ok {

		t, err := expandSwitchAclServiceCustomIcmpcode(d, v, "icmpcode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmpcode"] = t
		}
	}

	if v, ok := d.GetOk("tcp_portrange"); ok {

		t, err := expandSwitchAclServiceCustomTcpPortrange(d, v, "tcp_portrange", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-portrange"] = t
		}
	}

	return &obj, nil
}
