// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Show packet sniffer configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSnifferProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnifferProfileCreate,
		Read:   resourceSystemSnifferProfileRead,
		Update: resourceSystemSnifferProfileUpdate,
		Delete: resourceSystemSnifferProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"switch_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"max_pkt_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"system_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"profile_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 36),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"max_pkt_len": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(64, 1534),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemSnifferProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSnifferProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnifferProfile resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSnifferProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemSnifferProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSnifferProfile")
	}

	return resourceSystemSnifferProfileRead(d, m)
}

func resourceSystemSnifferProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSnifferProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnifferProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSnifferProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnifferProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSnifferProfile")
	}

	return resourceSystemSnifferProfileRead(d, m)
}

func resourceSystemSnifferProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemSnifferProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnifferProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnifferProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemSnifferProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnifferProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnifferProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnifferProfile resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnifferProfileSwitchInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnifferProfileMaxPktCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnifferProfileSystemInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnifferProfileFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnifferProfileProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnifferProfileMaxPktLen(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSnifferProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("switch_interface", flattenSystemSnifferProfileSwitchInterface(o["switch-interface"], d, "switch_interface", sv)); err != nil {
		if !fortiAPIPatch(o["switch-interface"]) {
			return fmt.Errorf("Error reading switch_interface: %v", err)
		}
	}

	if err = d.Set("max_pkt_count", flattenSystemSnifferProfileMaxPktCount(o["max-pkt-count"], d, "max_pkt_count", sv)); err != nil {
		if !fortiAPIPatch(o["max-pkt-count"]) {
			return fmt.Errorf("Error reading max_pkt_count: %v", err)
		}
	}

	if err = d.Set("system_interface", flattenSystemSnifferProfileSystemInterface(o["system-interface"], d, "system_interface", sv)); err != nil {
		if !fortiAPIPatch(o["system-interface"]) {
			return fmt.Errorf("Error reading system_interface: %v", err)
		}
	}

	if err = d.Set("filter", flattenSystemSnifferProfileFilter(o["filter"], d, "filter", sv)); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("profile_name", flattenSystemSnifferProfileProfileName(o["profile-name"], d, "profile_name", sv)); err != nil {
		if !fortiAPIPatch(o["profile-name"]) {
			return fmt.Errorf("Error reading profile_name: %v", err)
		}
	}

	if err = d.Set("max_pkt_len", flattenSystemSnifferProfileMaxPktLen(o["max-pkt-len"], d, "max_pkt_len", sv)); err != nil {
		if !fortiAPIPatch(o["max-pkt-len"]) {
			return fmt.Errorf("Error reading max_pkt_len: %v", err)
		}
	}

	return nil
}

func flattenSystemSnifferProfileFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemSnifferProfileSwitchInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnifferProfileMaxPktCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnifferProfileSystemInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnifferProfileFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnifferProfileProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnifferProfileMaxPktLen(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnifferProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("switch_interface"); ok {

		t, err := expandSystemSnifferProfileSwitchInterface(d, v, "switch_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-interface"] = t
		}
	}

	if v, ok := d.GetOk("max_pkt_count"); ok {

		t, err := expandSystemSnifferProfileMaxPktCount(d, v, "max_pkt_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-pkt-count"] = t
		}
	}

	if v, ok := d.GetOk("system_interface"); ok {

		t, err := expandSystemSnifferProfileSystemInterface(d, v, "system_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-interface"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok {

		t, err := expandSystemSnifferProfileFilter(d, v, "filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("profile_name"); ok {

		t, err := expandSystemSnifferProfileProfileName(d, v, "profile_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-name"] = t
		}
	}

	if v, ok := d.GetOk("max_pkt_len"); ok {

		t, err := expandSystemSnifferProfileMaxPktLen(d, v, "max_pkt_len", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-pkt-len"] = t
		}
	}

	return &obj, nil
}
