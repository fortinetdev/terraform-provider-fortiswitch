// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: SNMP user configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemSnmpUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemSnmpUserRead,
		Schema: map[string]*schema.Schema{

			"notify_hosts": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"priv_proto": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"query_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auth_pwd": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"priv_pwd": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_level": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_proto": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"queries": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"events": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemSnmpUserRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemSnmpUser: type error")
	}

	o, err := c.ReadSystemSnmpUser(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemSnmpUser: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemSnmpUser(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemSnmpUser from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemSnmpUserNotifyHosts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserPrivProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserQueryPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserAuthPwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserPrivPwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserSecurityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserAuthProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserQueries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserEvents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemSnmpUser(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("notify_hosts", dataSourceFlattenSystemSnmpUserNotifyHosts(o["notify-hosts"], d, "notify_hosts")); err != nil {
		if !fortiAPIPatch(o["notify-hosts"]) {
			return fmt.Errorf("Error reading notify_hosts: %v", err)
		}
	}

	if err = d.Set("priv_proto", dataSourceFlattenSystemSnmpUserPrivProto(o["priv-proto"], d, "priv_proto")); err != nil {
		if !fortiAPIPatch(o["priv-proto"]) {
			return fmt.Errorf("Error reading priv_proto: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenSystemSnmpUserName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("query_port", dataSourceFlattenSystemSnmpUserQueryPort(o["query-port"], d, "query_port")); err != nil {
		if !fortiAPIPatch(o["query-port"]) {
			return fmt.Errorf("Error reading query_port: %v", err)
		}
	}

	if err = d.Set("auth_pwd", dataSourceFlattenSystemSnmpUserAuthPwd(o["auth-pwd"], d, "auth_pwd")); err != nil {
		if !fortiAPIPatch(o["auth-pwd"]) {
			return fmt.Errorf("Error reading auth_pwd: %v", err)
		}
	}

	if err = d.Set("priv_pwd", dataSourceFlattenSystemSnmpUserPrivPwd(o["priv-pwd"], d, "priv_pwd")); err != nil {
		if !fortiAPIPatch(o["priv-pwd"]) {
			return fmt.Errorf("Error reading priv_pwd: %v", err)
		}
	}

	if err = d.Set("security_level", dataSourceFlattenSystemSnmpUserSecurityLevel(o["security-level"], d, "security_level")); err != nil {
		if !fortiAPIPatch(o["security-level"]) {
			return fmt.Errorf("Error reading security_level: %v", err)
		}
	}

	if err = d.Set("auth_proto", dataSourceFlattenSystemSnmpUserAuthProto(o["auth-proto"], d, "auth_proto")); err != nil {
		if !fortiAPIPatch(o["auth-proto"]) {
			return fmt.Errorf("Error reading auth_proto: %v", err)
		}
	}

	if err = d.Set("queries", dataSourceFlattenSystemSnmpUserQueries(o["queries"], d, "queries")); err != nil {
		if !fortiAPIPatch(o["queries"]) {
			return fmt.Errorf("Error reading queries: %v", err)
		}
	}

	if err = d.Set("events", dataSourceFlattenSystemSnmpUserEvents(o["events"], d, "events")); err != nil {
		if !fortiAPIPatch(o["events"]) {
			return fmt.Errorf("Error reading events: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemSnmpUserFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
