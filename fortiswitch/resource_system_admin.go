// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Administrative user configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminCreate,
		Read:   resourceSystemAdminRead,
		Update: resourceSystemAdminUpdate,
		Delete: resourceSystemAdminDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ip6_trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_admin": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"accprofile_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"accprofile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"force_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_remove_admin_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wildcard_fallback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"peer_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hidden": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 204),
				Optional:     true,
			},
			"vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 11),
				Optional:     true,
				Computed:     true,
			},
			"remote_auth": &schema.Schema{
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
			"password_expire": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"wildcard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_public_key1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_public_key3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_public_key2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAdminCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdmin resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAdmin(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemAdmin resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAdmin")
	}

	return resourceSystemAdminRead(d, m)
}

func resourceSystemAdminUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdmin resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAdmin(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdmin resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAdmin")
	}

	return resourceSystemAdminRead(d, m)
}

func resourceSystemAdminDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAdmin(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAdmin(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdmin(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdmin resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminIp6Trusthost2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminIsAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost7(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost5(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminAccprofileOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost8(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost9(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminAccprofile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminForcePasswordChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminAllowRemoveAdminSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminWildcardFallback(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminSchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminPeerGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminTrusthost10(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminHidden(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminTrusthost8(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost9(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost7(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost5(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminIp6Trusthost10(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminRemoteAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminPasswordExpire(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminRemoteGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminWildcard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminSshPublicKey1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminSshPublicKey3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminSshPublicKey2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAdminPeerAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAdmin(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ip6_trusthost2", flattenSystemAdminIp6Trusthost2(o["ip6-trusthost2"], d, "ip6_trusthost2", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost2"]) {
			return fmt.Errorf("Error reading ip6_trusthost2: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost3", flattenSystemAdminIp6Trusthost3(o["ip6-trusthost3"], d, "ip6_trusthost3", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost3"]) {
			return fmt.Errorf("Error reading ip6_trusthost3: %v", err)
		}
	}

	if err = d.Set("is_admin", flattenSystemAdminIsAdmin(o["is-admin"], d, "is_admin", sv)); err != nil {
		if !fortiAPIPatch(o["is-admin"]) {
			return fmt.Errorf("Error reading is_admin: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost1", flattenSystemAdminIp6Trusthost1(o["ip6-trusthost1"], d, "ip6_trusthost1", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost1"]) {
			return fmt.Errorf("Error reading ip6_trusthost1: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost6", flattenSystemAdminIp6Trusthost6(o["ip6-trusthost6"], d, "ip6_trusthost6", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost6"]) {
			return fmt.Errorf("Error reading ip6_trusthost6: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost7", flattenSystemAdminIp6Trusthost7(o["ip6-trusthost7"], d, "ip6_trusthost7", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost7"]) {
			return fmt.Errorf("Error reading ip6_trusthost7: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost4", flattenSystemAdminIp6Trusthost4(o["ip6-trusthost4"], d, "ip6_trusthost4", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost4"]) {
			return fmt.Errorf("Error reading ip6_trusthost4: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost5", flattenSystemAdminIp6Trusthost5(o["ip6-trusthost5"], d, "ip6_trusthost5", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost5"]) {
			return fmt.Errorf("Error reading ip6_trusthost5: %v", err)
		}
	}

	if err = d.Set("accprofile_override", flattenSystemAdminAccprofileOverride(o["accprofile-override"], d, "accprofile_override", sv)); err != nil {
		if !fortiAPIPatch(o["accprofile-override"]) {
			return fmt.Errorf("Error reading accprofile_override: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost8", flattenSystemAdminIp6Trusthost8(o["ip6-trusthost8"], d, "ip6_trusthost8", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost8"]) {
			return fmt.Errorf("Error reading ip6_trusthost8: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost9", flattenSystemAdminIp6Trusthost9(o["ip6-trusthost9"], d, "ip6_trusthost9", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost9"]) {
			return fmt.Errorf("Error reading ip6_trusthost9: %v", err)
		}
	}

	if err = d.Set("accprofile", flattenSystemAdminAccprofile(o["accprofile"], d, "accprofile", sv)); err != nil {
		if !fortiAPIPatch(o["accprofile"]) {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("force_password_change", flattenSystemAdminForcePasswordChange(o["force-password-change"], d, "force_password_change", sv)); err != nil {
		if !fortiAPIPatch(o["force-password-change"]) {
			return fmt.Errorf("Error reading force_password_change: %v", err)
		}
	}

	if err = d.Set("allow_remove_admin_session", flattenSystemAdminAllowRemoveAdminSession(o["allow-remove-admin-session"], d, "allow_remove_admin_session", sv)); err != nil {
		if !fortiAPIPatch(o["allow-remove-admin-session"]) {
			return fmt.Errorf("Error reading allow_remove_admin_session: %v", err)
		}
	}

	if err = d.Set("wildcard_fallback", flattenSystemAdminWildcardFallback(o["wildcard-fallback"], d, "wildcard_fallback", sv)); err != nil {
		if !fortiAPIPatch(o["wildcard-fallback"]) {
			return fmt.Errorf("Error reading wildcard_fallback: %v", err)
		}
	}

	if err = d.Set("schedule", flattenSystemAdminSchedule(o["schedule"], d, "schedule", sv)); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("peer_group", flattenSystemAdminPeerGroup(o["peer-group"], d, "peer_group", sv)); err != nil {
		if !fortiAPIPatch(o["peer-group"]) {
			return fmt.Errorf("Error reading peer_group: %v", err)
		}
	}

	if err = d.Set("comments", flattenSystemAdminComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("trusthost10", flattenSystemAdminTrusthost10(o["trusthost10"], d, "trusthost10", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost10"]) {
			return fmt.Errorf("Error reading trusthost10: %v", err)
		}
	}

	if err = d.Set("hidden", flattenSystemAdminHidden(o["hidden"], d, "hidden", sv)); err != nil {
		if !fortiAPIPatch(o["hidden"]) {
			return fmt.Errorf("Error reading hidden: %v", err)
		}
	}

	if err = d.Set("trusthost8", flattenSystemAdminTrusthost8(o["trusthost8"], d, "trusthost8", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost8"]) {
			return fmt.Errorf("Error reading trusthost8: %v", err)
		}
	}

	if err = d.Set("trusthost9", flattenSystemAdminTrusthost9(o["trusthost9"], d, "trusthost9", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost9"]) {
			return fmt.Errorf("Error reading trusthost9: %v", err)
		}
	}

	if err = d.Set("trusthost6", flattenSystemAdminTrusthost6(o["trusthost6"], d, "trusthost6", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost6"]) {
			return fmt.Errorf("Error reading trusthost6: %v", err)
		}
	}

	if err = d.Set("trusthost7", flattenSystemAdminTrusthost7(o["trusthost7"], d, "trusthost7", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost7"]) {
			return fmt.Errorf("Error reading trusthost7: %v", err)
		}
	}

	if err = d.Set("trusthost4", flattenSystemAdminTrusthost4(o["trusthost4"], d, "trusthost4", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost4"]) {
			return fmt.Errorf("Error reading trusthost4: %v", err)
		}
	}

	if err = d.Set("trusthost5", flattenSystemAdminTrusthost5(o["trusthost5"], d, "trusthost5", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost5"]) {
			return fmt.Errorf("Error reading trusthost5: %v", err)
		}
	}

	if err = d.Set("trusthost2", flattenSystemAdminTrusthost2(o["trusthost2"], d, "trusthost2", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost2"]) {
			return fmt.Errorf("Error reading trusthost2: %v", err)
		}
	}

	if err = d.Set("trusthost3", flattenSystemAdminTrusthost3(o["trusthost3"], d, "trusthost3", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost3"]) {
			return fmt.Errorf("Error reading trusthost3: %v", err)
		}
	}

	if err = d.Set("trusthost1", flattenSystemAdminTrusthost1(o["trusthost1"], d, "trusthost1", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost1"]) {
			return fmt.Errorf("Error reading trusthost1: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost10", flattenSystemAdminIp6Trusthost10(o["ip6-trusthost10"], d, "ip6_trusthost10", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost10"]) {
			return fmt.Errorf("Error reading ip6_trusthost10: %v", err)
		}
	}

	if err = d.Set("password", flattenSystemAdminPassword(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemAdminVdom(o["vdom"], d, "vdom", sv)); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("remote_auth", flattenSystemAdminRemoteAuth(o["remote-auth"], d, "remote_auth", sv)); err != nil {
		if !fortiAPIPatch(o["remote-auth"]) {
			return fmt.Errorf("Error reading remote_auth: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAdminName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("password_expire", flattenSystemAdminPasswordExpire(o["password-expire"], d, "password_expire", sv)); err != nil {
		if !fortiAPIPatch(o["password-expire"]) {
			return fmt.Errorf("Error reading password_expire: %v", err)
		}
	}

	if err = d.Set("remote_group", flattenSystemAdminRemoteGroup(o["remote-group"], d, "remote_group", sv)); err != nil {
		if !fortiAPIPatch(o["remote-group"]) {
			return fmt.Errorf("Error reading remote_group: %v", err)
		}
	}

	if err = d.Set("wildcard", flattenSystemAdminWildcard(o["wildcard"], d, "wildcard", sv)); err != nil {
		if !fortiAPIPatch(o["wildcard"]) {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	if err = d.Set("ssh_public_key1", flattenSystemAdminSshPublicKey1(o["ssh-public-key1"], d, "ssh_public_key1", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-public-key1"]) {
			return fmt.Errorf("Error reading ssh_public_key1: %v", err)
		}
	}

	if err = d.Set("ssh_public_key3", flattenSystemAdminSshPublicKey3(o["ssh-public-key3"], d, "ssh_public_key3", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-public-key3"]) {
			return fmt.Errorf("Error reading ssh_public_key3: %v", err)
		}
	}

	if err = d.Set("ssh_public_key2", flattenSystemAdminSshPublicKey2(o["ssh-public-key2"], d, "ssh_public_key2", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-public-key2"]) {
			return fmt.Errorf("Error reading ssh_public_key2: %v", err)
		}
	}

	if err = d.Set("peer_auth", flattenSystemAdminPeerAuth(o["peer-auth"], d, "peer_auth", sv)); err != nil {
		if !fortiAPIPatch(o["peer-auth"]) {
			return fmt.Errorf("Error reading peer_auth: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemAdminIp6Trusthost2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIsAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost7(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminAccprofileOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost8(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost9(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminAccprofile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminForcePasswordChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminAllowRemoveAdminSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminWildcardFallback(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminPeerGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost10(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminHidden(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost8(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost9(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost7(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost10(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminRemoteAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminPasswordExpire(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminRemoteGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminWildcard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSshPublicKey1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return toCertFormat(v), nil
}

func expandSystemAdminSshPublicKey3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return toCertFormat(v), nil
}

func expandSystemAdminSshPublicKey2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return toCertFormat(v), nil
}

func expandSystemAdminPeerAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdmin(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip6_trusthost2"); ok {

		t, err := expandSystemAdminIp6Trusthost2(d, v, "ip6_trusthost2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost3"); ok {

		t, err := expandSystemAdminIp6Trusthost3(d, v, "ip6_trusthost3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("is_admin"); ok {

		t, err := expandSystemAdminIsAdmin(d, v, "is_admin", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["is-admin"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost1"); ok {

		t, err := expandSystemAdminIp6Trusthost1(d, v, "ip6_trusthost1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost6"); ok {

		t, err := expandSystemAdminIp6Trusthost6(d, v, "ip6_trusthost6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost7"); ok {

		t, err := expandSystemAdminIp6Trusthost7(d, v, "ip6_trusthost7", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost4"); ok {

		t, err := expandSystemAdminIp6Trusthost4(d, v, "ip6_trusthost4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost5"); ok {

		t, err := expandSystemAdminIp6Trusthost5(d, v, "ip6_trusthost5", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("accprofile_override"); ok {

		t, err := expandSystemAdminAccprofileOverride(d, v, "accprofile_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile-override"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost8"); ok {

		t, err := expandSystemAdminIp6Trusthost8(d, v, "ip6_trusthost8", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost9"); ok {

		t, err := expandSystemAdminIp6Trusthost9(d, v, "ip6_trusthost9", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost9"] = t
		}
	}

	if v, ok := d.GetOk("accprofile"); ok {

		t, err := expandSystemAdminAccprofile(d, v, "accprofile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	}

	if v, ok := d.GetOk("force_password_change"); ok {

		t, err := expandSystemAdminForcePasswordChange(d, v, "force_password_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["force-password-change"] = t
		}
	}

	if v, ok := d.GetOk("allow_remove_admin_session"); ok {

		t, err := expandSystemAdminAllowRemoveAdminSession(d, v, "allow_remove_admin_session", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-remove-admin-session"] = t
		}
	}

	if v, ok := d.GetOk("wildcard_fallback"); ok {

		t, err := expandSystemAdminWildcardFallback(d, v, "wildcard_fallback", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard-fallback"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {

		t, err := expandSystemAdminSchedule(d, v, "schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("peer_group"); ok {

		t, err := expandSystemAdminPeerGroup(d, v, "peer_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-group"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandSystemAdminComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("trusthost10"); ok {

		t, err := expandSystemAdminTrusthost10(d, v, "trusthost10", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("hidden"); ok {

		t, err := expandSystemAdminHidden(d, v, "hidden", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hidden"] = t
		}
	}

	if v, ok := d.GetOk("trusthost8"); ok {

		t, err := expandSystemAdminTrusthost8(d, v, "trusthost8", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("trusthost9"); ok {

		t, err := expandSystemAdminTrusthost9(d, v, "trusthost9", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost9"] = t
		}
	}

	if v, ok := d.GetOk("trusthost6"); ok {

		t, err := expandSystemAdminTrusthost6(d, v, "trusthost6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("trusthost7"); ok {

		t, err := expandSystemAdminTrusthost7(d, v, "trusthost7", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("trusthost4"); ok {

		t, err := expandSystemAdminTrusthost4(d, v, "trusthost4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("trusthost5"); ok {

		t, err := expandSystemAdminTrusthost5(d, v, "trusthost5", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("trusthost2"); ok {

		t, err := expandSystemAdminTrusthost2(d, v, "trusthost2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("trusthost3"); ok {

		t, err := expandSystemAdminTrusthost3(d, v, "trusthost3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("trusthost1"); ok {

		t, err := expandSystemAdminTrusthost1(d, v, "trusthost1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost10"); ok {

		t, err := expandSystemAdminIp6Trusthost10(d, v, "ip6_trusthost10", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {

		t, err := expandSystemAdminPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {

		t, err := expandSystemAdminVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("remote_auth"); ok {

		t, err := expandSystemAdminRemoteAuth(d, v, "remote_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-auth"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemAdminName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password_expire"); ok {

		t, err := expandSystemAdminPasswordExpire(d, v, "password_expire", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-expire"] = t
		}
	}

	if v, ok := d.GetOk("remote_group"); ok {

		t, err := expandSystemAdminRemoteGroup(d, v, "remote_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-group"] = t
		}
	}

	if v, ok := d.GetOk("wildcard"); ok {

		t, err := expandSystemAdminWildcard(d, v, "wildcard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard"] = t
		}
	}

	if v, ok := d.GetOk("ssh_public_key1"); ok {

		t, err := expandSystemAdminSshPublicKey1(d, v, "ssh_public_key1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-public-key1"] = t
		}
	}

	if v, ok := d.GetOk("ssh_public_key3"); ok {

		t, err := expandSystemAdminSshPublicKey3(d, v, "ssh_public_key3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-public-key3"] = t
		}
	}

	if v, ok := d.GetOk("ssh_public_key2"); ok {

		t, err := expandSystemAdminSshPublicKey2(d, v, "ssh_public_key2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-public-key2"] = t
		}
	}

	if v, ok := d.GetOk("peer_auth"); ok {

		t, err := expandSystemAdminPeerAuth(d, v, "peer_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-auth"] = t
		}
	}

	return &obj, nil
}
