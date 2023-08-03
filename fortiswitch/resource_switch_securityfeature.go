// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Switch security feature control nobs.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchSecurityFeature() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchSecurityFeatureUpdate,
		Read:   resourceSwitchSecurityFeatureRead,
		Update: resourceSwitchSecurityFeatureUpdate,
		Delete: resourceSwitchSecurityFeatureDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"tcp_port_eq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"macsa_eq_macda": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_flag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip_eq_dip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_mcast_sa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_flag_sf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"v4_first_frag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_sa_mac_all_zero": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_hdr_partial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_port_eq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_flag_fup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchSecurityFeatureUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchSecurityFeature(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchSecurityFeature resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchSecurityFeature(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchSecurityFeature resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchSecurityFeature")
	}

	return resourceSwitchSecurityFeatureRead(d, m)
}

func resourceSwitchSecurityFeatureDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchSecurityFeature(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchSecurityFeature resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchSecurityFeature(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchSecurityFeature resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchSecurityFeatureRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchSecurityFeature(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchSecurityFeature resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchSecurityFeature(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchSecurityFeature resource from API: %v", err)
	}
	return nil
}

func flattenSwitchSecurityFeatureTcpPortEq(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchSecurityFeatureMacsaEqMacda(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchSecurityFeatureTcpFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchSecurityFeatureSipEqDip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchSecurityFeatureAllowMcastSa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchSecurityFeatureTcpFlagSf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchSecurityFeatureV4FirstFrag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchSecurityFeatureAllowSaMacAllZero(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchSecurityFeatureTcpHdrPartial(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchSecurityFeatureUdpPortEq(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchSecurityFeatureTcpFlagFup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchSecurityFeature(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("tcp_port_eq", flattenSwitchSecurityFeatureTcpPortEq(o["tcp-port-eq"], d, "tcp_port_eq", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-port-eq"]) {
			return fmt.Errorf("Error reading tcp_port_eq: %v", err)
		}
	}

	if err = d.Set("macsa_eq_macda", flattenSwitchSecurityFeatureMacsaEqMacda(o["macsa-eq-macda"], d, "macsa_eq_macda", sv)); err != nil {
		if !fortiAPIPatch(o["macsa-eq-macda"]) {
			return fmt.Errorf("Error reading macsa_eq_macda: %v", err)
		}
	}

	if err = d.Set("tcp_flag", flattenSwitchSecurityFeatureTcpFlag(o["tcp-flag"], d, "tcp_flag", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-flag"]) {
			return fmt.Errorf("Error reading tcp_flag: %v", err)
		}
	}

	if err = d.Set("sip_eq_dip", flattenSwitchSecurityFeatureSipEqDip(o["sip-eq-dip"], d, "sip_eq_dip", sv)); err != nil {
		if !fortiAPIPatch(o["sip-eq-dip"]) {
			return fmt.Errorf("Error reading sip_eq_dip: %v", err)
		}
	}

	if err = d.Set("allow_mcast_sa", flattenSwitchSecurityFeatureAllowMcastSa(o["allow-mcast-sa"], d, "allow_mcast_sa", sv)); err != nil {
		if !fortiAPIPatch(o["allow-mcast-sa"]) {
			return fmt.Errorf("Error reading allow_mcast_sa: %v", err)
		}
	}

	if err = d.Set("tcp_flag_sf", flattenSwitchSecurityFeatureTcpFlagSf(o["tcp-flag-SF"], d, "tcp_flag_sf", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-flag-SF"]) {
			return fmt.Errorf("Error reading tcp_flag_sf: %v", err)
		}
	}

	if err = d.Set("v4_first_frag", flattenSwitchSecurityFeatureV4FirstFrag(o["v4-first-frag"], d, "v4_first_frag", sv)); err != nil {
		if !fortiAPIPatch(o["v4-first-frag"]) {
			return fmt.Errorf("Error reading v4_first_frag: %v", err)
		}
	}

	if err = d.Set("allow_sa_mac_all_zero", flattenSwitchSecurityFeatureAllowSaMacAllZero(o["allow-sa-mac-all-zero"], d, "allow_sa_mac_all_zero", sv)); err != nil {
		if !fortiAPIPatch(o["allow-sa-mac-all-zero"]) {
			return fmt.Errorf("Error reading allow_sa_mac_all_zero: %v", err)
		}
	}

	if err = d.Set("tcp_hdr_partial", flattenSwitchSecurityFeatureTcpHdrPartial(o["tcp-hdr-partial"], d, "tcp_hdr_partial", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-hdr-partial"]) {
			return fmt.Errorf("Error reading tcp_hdr_partial: %v", err)
		}
	}

	if err = d.Set("udp_port_eq", flattenSwitchSecurityFeatureUdpPortEq(o["udp-port-eq"], d, "udp_port_eq", sv)); err != nil {
		if !fortiAPIPatch(o["udp-port-eq"]) {
			return fmt.Errorf("Error reading udp_port_eq: %v", err)
		}
	}

	if err = d.Set("tcp_flag_fup", flattenSwitchSecurityFeatureTcpFlagFup(o["tcp-flag-FUP"], d, "tcp_flag_fup", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-flag-FUP"]) {
			return fmt.Errorf("Error reading tcp_flag_fup: %v", err)
		}
	}

	return nil
}

func flattenSwitchSecurityFeatureFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchSecurityFeatureTcpPortEq(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchSecurityFeatureMacsaEqMacda(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchSecurityFeatureTcpFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchSecurityFeatureSipEqDip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchSecurityFeatureAllowMcastSa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchSecurityFeatureTcpFlagSf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchSecurityFeatureV4FirstFrag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchSecurityFeatureAllowSaMacAllZero(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchSecurityFeatureTcpHdrPartial(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchSecurityFeatureUdpPortEq(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchSecurityFeatureTcpFlagFup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchSecurityFeature(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("tcp_port_eq"); ok {
		if setArgNil {
			obj["tcp-port-eq"] = nil
		} else {

			t, err := expandSwitchSecurityFeatureTcpPortEq(d, v, "tcp_port_eq", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tcp-port-eq"] = t
			}
		}
	}

	if v, ok := d.GetOk("macsa_eq_macda"); ok {
		if setArgNil {
			obj["macsa-eq-macda"] = nil
		} else {

			t, err := expandSwitchSecurityFeatureMacsaEqMacda(d, v, "macsa_eq_macda", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["macsa-eq-macda"] = t
			}
		}
	}

	if v, ok := d.GetOk("tcp_flag"); ok {
		if setArgNil {
			obj["tcp-flag"] = nil
		} else {

			t, err := expandSwitchSecurityFeatureTcpFlag(d, v, "tcp_flag", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tcp-flag"] = t
			}
		}
	}

	if v, ok := d.GetOk("sip_eq_dip"); ok {
		if setArgNil {
			obj["sip-eq-dip"] = nil
		} else {

			t, err := expandSwitchSecurityFeatureSipEqDip(d, v, "sip_eq_dip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sip-eq-dip"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_mcast_sa"); ok {
		if setArgNil {
			obj["allow-mcast-sa"] = nil
		} else {

			t, err := expandSwitchSecurityFeatureAllowMcastSa(d, v, "allow_mcast_sa", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-mcast-sa"] = t
			}
		}
	}

	if v, ok := d.GetOk("tcp_flag_sf"); ok {
		if setArgNil {
			obj["tcp-flag-SF"] = nil
		} else {

			t, err := expandSwitchSecurityFeatureTcpFlagSf(d, v, "tcp_flag_sf", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tcp-flag-SF"] = t
			}
		}
	}

	if v, ok := d.GetOk("v4_first_frag"); ok {
		if setArgNil {
			obj["v4-first-frag"] = nil
		} else {

			t, err := expandSwitchSecurityFeatureV4FirstFrag(d, v, "v4_first_frag", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["v4-first-frag"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_sa_mac_all_zero"); ok {
		if setArgNil {
			obj["allow-sa-mac-all-zero"] = nil
		} else {

			t, err := expandSwitchSecurityFeatureAllowSaMacAllZero(d, v, "allow_sa_mac_all_zero", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-sa-mac-all-zero"] = t
			}
		}
	}

	if v, ok := d.GetOk("tcp_hdr_partial"); ok {
		if setArgNil {
			obj["tcp-hdr-partial"] = nil
		} else {

			t, err := expandSwitchSecurityFeatureTcpHdrPartial(d, v, "tcp_hdr_partial", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tcp-hdr-partial"] = t
			}
		}
	}

	if v, ok := d.GetOk("udp_port_eq"); ok {
		if setArgNil {
			obj["udp-port-eq"] = nil
		} else {

			t, err := expandSwitchSecurityFeatureUdpPortEq(d, v, "udp_port_eq", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["udp-port-eq"] = t
			}
		}
	}

	if v, ok := d.GetOk("tcp_flag_fup"); ok {
		if setArgNil {
			obj["tcp-flag-FUP"] = nil
		} else {

			t, err := expandSwitchSecurityFeatureTcpFlagFup(d, v, "tcp_flag_fup", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tcp-flag-FUP"] = t
			}
		}
	}

	return &obj, nil
}
