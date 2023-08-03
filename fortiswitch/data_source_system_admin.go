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

func dataSourceSystemAdmin() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAdminRead,
		Schema: map[string]*schema.Schema{

			"ip6_trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_admin": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ip6_trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"accprofile_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"accprofile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"force_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_remove_admin_session": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hidden": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_auth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password_expire": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wildcard": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_public_key1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_public_key3": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_public_key2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_auth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemAdminRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemAdmin: type error")
	}

	o, err := c.ReadSystemAdmin(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemAdmin: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemAdmin(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemAdmin from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemAdminIp6Trusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIsAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminAccprofileOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminForcePasswordChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminAllowRemoveAdminSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminPeerGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminTrusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminHidden(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminTrusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminRemoteAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminPasswordExpire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminRemoteGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminWildcard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminSshPublicKey1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminSshPublicKey3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminSshPublicKey2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminPeerAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAdmin(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip6_trusthost2", dataSourceFlattenSystemAdminIp6Trusthost2(o["ip6-trusthost2"], d, "ip6_trusthost2")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost2"]) {
			return fmt.Errorf("Error reading ip6_trusthost2: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost3", dataSourceFlattenSystemAdminIp6Trusthost3(o["ip6-trusthost3"], d, "ip6_trusthost3")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost3"]) {
			return fmt.Errorf("Error reading ip6_trusthost3: %v", err)
		}
	}

	if err = d.Set("is_admin", dataSourceFlattenSystemAdminIsAdmin(o["is-admin"], d, "is_admin")); err != nil {
		if !fortiAPIPatch(o["is-admin"]) {
			return fmt.Errorf("Error reading is_admin: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost1", dataSourceFlattenSystemAdminIp6Trusthost1(o["ip6-trusthost1"], d, "ip6_trusthost1")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost1"]) {
			return fmt.Errorf("Error reading ip6_trusthost1: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost6", dataSourceFlattenSystemAdminIp6Trusthost6(o["ip6-trusthost6"], d, "ip6_trusthost6")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost6"]) {
			return fmt.Errorf("Error reading ip6_trusthost6: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost7", dataSourceFlattenSystemAdminIp6Trusthost7(o["ip6-trusthost7"], d, "ip6_trusthost7")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost7"]) {
			return fmt.Errorf("Error reading ip6_trusthost7: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost4", dataSourceFlattenSystemAdminIp6Trusthost4(o["ip6-trusthost4"], d, "ip6_trusthost4")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost4"]) {
			return fmt.Errorf("Error reading ip6_trusthost4: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost5", dataSourceFlattenSystemAdminIp6Trusthost5(o["ip6-trusthost5"], d, "ip6_trusthost5")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost5"]) {
			return fmt.Errorf("Error reading ip6_trusthost5: %v", err)
		}
	}

	if err = d.Set("accprofile_override", dataSourceFlattenSystemAdminAccprofileOverride(o["accprofile-override"], d, "accprofile_override")); err != nil {
		if !fortiAPIPatch(o["accprofile-override"]) {
			return fmt.Errorf("Error reading accprofile_override: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost8", dataSourceFlattenSystemAdminIp6Trusthost8(o["ip6-trusthost8"], d, "ip6_trusthost8")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost8"]) {
			return fmt.Errorf("Error reading ip6_trusthost8: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost9", dataSourceFlattenSystemAdminIp6Trusthost9(o["ip6-trusthost9"], d, "ip6_trusthost9")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost9"]) {
			return fmt.Errorf("Error reading ip6_trusthost9: %v", err)
		}
	}

	if err = d.Set("accprofile", dataSourceFlattenSystemAdminAccprofile(o["accprofile"], d, "accprofile")); err != nil {
		if !fortiAPIPatch(o["accprofile"]) {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("force_password_change", dataSourceFlattenSystemAdminForcePasswordChange(o["force-password-change"], d, "force_password_change")); err != nil {
		if !fortiAPIPatch(o["force-password-change"]) {
			return fmt.Errorf("Error reading force_password_change: %v", err)
		}
	}

	if err = d.Set("allow_remove_admin_session", dataSourceFlattenSystemAdminAllowRemoveAdminSession(o["allow-remove-admin-session"], d, "allow_remove_admin_session")); err != nil {
		if !fortiAPIPatch(o["allow-remove-admin-session"]) {
			return fmt.Errorf("Error reading allow_remove_admin_session: %v", err)
		}
	}

	if err = d.Set("schedule", dataSourceFlattenSystemAdminSchedule(o["schedule"], d, "schedule")); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("peer_group", dataSourceFlattenSystemAdminPeerGroup(o["peer-group"], d, "peer_group")); err != nil {
		if !fortiAPIPatch(o["peer-group"]) {
			return fmt.Errorf("Error reading peer_group: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenSystemAdminComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("trusthost10", dataSourceFlattenSystemAdminTrusthost10(o["trusthost10"], d, "trusthost10")); err != nil {
		if !fortiAPIPatch(o["trusthost10"]) {
			return fmt.Errorf("Error reading trusthost10: %v", err)
		}
	}

	if err = d.Set("hidden", dataSourceFlattenSystemAdminHidden(o["hidden"], d, "hidden")); err != nil {
		if !fortiAPIPatch(o["hidden"]) {
			return fmt.Errorf("Error reading hidden: %v", err)
		}
	}

	if err = d.Set("trusthost8", dataSourceFlattenSystemAdminTrusthost8(o["trusthost8"], d, "trusthost8")); err != nil {
		if !fortiAPIPatch(o["trusthost8"]) {
			return fmt.Errorf("Error reading trusthost8: %v", err)
		}
	}

	if err = d.Set("trusthost9", dataSourceFlattenSystemAdminTrusthost9(o["trusthost9"], d, "trusthost9")); err != nil {
		if !fortiAPIPatch(o["trusthost9"]) {
			return fmt.Errorf("Error reading trusthost9: %v", err)
		}
	}

	if err = d.Set("trusthost6", dataSourceFlattenSystemAdminTrusthost6(o["trusthost6"], d, "trusthost6")); err != nil {
		if !fortiAPIPatch(o["trusthost6"]) {
			return fmt.Errorf("Error reading trusthost6: %v", err)
		}
	}

	if err = d.Set("trusthost7", dataSourceFlattenSystemAdminTrusthost7(o["trusthost7"], d, "trusthost7")); err != nil {
		if !fortiAPIPatch(o["trusthost7"]) {
			return fmt.Errorf("Error reading trusthost7: %v", err)
		}
	}

	if err = d.Set("trusthost4", dataSourceFlattenSystemAdminTrusthost4(o["trusthost4"], d, "trusthost4")); err != nil {
		if !fortiAPIPatch(o["trusthost4"]) {
			return fmt.Errorf("Error reading trusthost4: %v", err)
		}
	}

	if err = d.Set("trusthost5", dataSourceFlattenSystemAdminTrusthost5(o["trusthost5"], d, "trusthost5")); err != nil {
		if !fortiAPIPatch(o["trusthost5"]) {
			return fmt.Errorf("Error reading trusthost5: %v", err)
		}
	}

	if err = d.Set("trusthost2", dataSourceFlattenSystemAdminTrusthost2(o["trusthost2"], d, "trusthost2")); err != nil {
		if !fortiAPIPatch(o["trusthost2"]) {
			return fmt.Errorf("Error reading trusthost2: %v", err)
		}
	}

	if err = d.Set("trusthost3", dataSourceFlattenSystemAdminTrusthost3(o["trusthost3"], d, "trusthost3")); err != nil {
		if !fortiAPIPatch(o["trusthost3"]) {
			return fmt.Errorf("Error reading trusthost3: %v", err)
		}
	}

	if err = d.Set("trusthost1", dataSourceFlattenSystemAdminTrusthost1(o["trusthost1"], d, "trusthost1")); err != nil {
		if !fortiAPIPatch(o["trusthost1"]) {
			return fmt.Errorf("Error reading trusthost1: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost10", dataSourceFlattenSystemAdminIp6Trusthost10(o["ip6-trusthost10"], d, "ip6_trusthost10")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost10"]) {
			return fmt.Errorf("Error reading ip6_trusthost10: %v", err)
		}
	}

	if err = d.Set("password", dataSourceFlattenSystemAdminPassword(o["password"], d, "password")); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("vdom", dataSourceFlattenSystemAdminVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("remote_auth", dataSourceFlattenSystemAdminRemoteAuth(o["remote-auth"], d, "remote_auth")); err != nil {
		if !fortiAPIPatch(o["remote-auth"]) {
			return fmt.Errorf("Error reading remote_auth: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenSystemAdminName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("password_expire", dataSourceFlattenSystemAdminPasswordExpire(o["password-expire"], d, "password_expire")); err != nil {
		if !fortiAPIPatch(o["password-expire"]) {
			return fmt.Errorf("Error reading password_expire: %v", err)
		}
	}

	if err = d.Set("remote_group", dataSourceFlattenSystemAdminRemoteGroup(o["remote-group"], d, "remote_group")); err != nil {
		if !fortiAPIPatch(o["remote-group"]) {
			return fmt.Errorf("Error reading remote_group: %v", err)
		}
	}

	if err = d.Set("wildcard", dataSourceFlattenSystemAdminWildcard(o["wildcard"], d, "wildcard")); err != nil {
		if !fortiAPIPatch(o["wildcard"]) {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	if err = d.Set("ssh_public_key1", dataSourceFlattenSystemAdminSshPublicKey1(o["ssh-public-key1"], d, "ssh_public_key1")); err != nil {
		if !fortiAPIPatch(o["ssh-public-key1"]) {
			return fmt.Errorf("Error reading ssh_public_key1: %v", err)
		}
	}

	if err = d.Set("ssh_public_key3", dataSourceFlattenSystemAdminSshPublicKey3(o["ssh-public-key3"], d, "ssh_public_key3")); err != nil {
		if !fortiAPIPatch(o["ssh-public-key3"]) {
			return fmt.Errorf("Error reading ssh_public_key3: %v", err)
		}
	}

	if err = d.Set("ssh_public_key2", dataSourceFlattenSystemAdminSshPublicKey2(o["ssh-public-key2"], d, "ssh_public_key2")); err != nil {
		if !fortiAPIPatch(o["ssh-public-key2"]) {
			return fmt.Errorf("Error reading ssh_public_key2: %v", err)
		}
	}

	if err = d.Set("peer_auth", dataSourceFlattenSystemAdminPeerAuth(o["peer-auth"], d, "peer_auth")); err != nil {
		if !fortiAPIPatch(o["peer-auth"]) {
			return fmt.Errorf("Error reading peer_auth: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemAdminFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
