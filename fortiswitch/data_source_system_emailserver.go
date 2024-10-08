// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Email server configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemEmailServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemEmailServerRead,
		Schema: map[string]*schema.Schema{

			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authenticate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"validate_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"reply_to": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemEmailServerRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemEmailServer"

	o, err := c.ReadSystemEmailServer(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemEmailServer: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemEmailServer(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemEmailServer from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemEmailServerUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemEmailServerAuthenticate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemEmailServerSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemEmailServerSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemEmailServerValidateServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemEmailServerServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemEmailServerSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemEmailServerReplyTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemEmailServerPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemEmailServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemEmailServerSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemEmailServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("username", dataSourceFlattenSystemEmailServerUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("authenticate", dataSourceFlattenSystemEmailServerAuthenticate(o["authenticate"], d, "authenticate")); err != nil {
		if !fortiAPIPatch(o["authenticate"]) {
			return fmt.Errorf("Error reading authenticate: %v", err)
		}
	}

	if err = d.Set("security", dataSourceFlattenSystemEmailServerSecurity(o["security"], d, "security")); err != nil {
		if !fortiAPIPatch(o["security"]) {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("source_ip", dataSourceFlattenSystemEmailServerSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("validate_server", dataSourceFlattenSystemEmailServerValidateServer(o["validate-server"], d, "validate_server")); err != nil {
		if !fortiAPIPatch(o["validate-server"]) {
			return fmt.Errorf("Error reading validate_server: %v", err)
		}
	}

	if err = d.Set("server", dataSourceFlattenSystemEmailServerServer(o["server"], d, "server")); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("source_ip6", dataSourceFlattenSystemEmailServerSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("reply_to", dataSourceFlattenSystemEmailServerReplyTo(o["reply-to"], d, "reply_to")); err != nil {
		if !fortiAPIPatch(o["reply-to"]) {
			return fmt.Errorf("Error reading reply_to: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenSystemEmailServerPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", dataSourceFlattenSystemEmailServerSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemEmailServerFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
