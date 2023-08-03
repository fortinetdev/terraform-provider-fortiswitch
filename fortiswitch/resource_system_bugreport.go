// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure bug report.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemBugReport() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemBugReportUpdate,
		Read:   resourceSystemBugReportRead,
		Update: resourceSystemBugReportUpdate,
		Delete: resourceSystemBugReportDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mailto": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"username_smtp": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemBugReportUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemBugReport(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemBugReport resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemBugReport(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemBugReport resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemBugReport")
	}

	return resourceSystemBugReportRead(d, m)
}

func resourceSystemBugReportDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemBugReport(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemBugReport resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemBugReport(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemBugReport resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemBugReportRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemBugReport(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemBugReport resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemBugReport(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemBugReport resource from API: %v", err)
	}
	return nil
}

func flattenSystemBugReportMailto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemBugReportUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemBugReportAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemBugReportServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemBugReportUsernameSmtp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemBugReportPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemBugReport(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mailto", flattenSystemBugReportMailto(o["mailto"], d, "mailto", sv)); err != nil {
		if !fortiAPIPatch(o["mailto"]) {
			return fmt.Errorf("Error reading mailto: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemBugReportUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("auth", flattenSystemBugReportAuth(o["auth"], d, "auth", sv)); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemBugReportServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("username_smtp", flattenSystemBugReportUsernameSmtp(o["username-smtp"], d, "username_smtp", sv)); err != nil {
		if !fortiAPIPatch(o["username-smtp"]) {
			return fmt.Errorf("Error reading username_smtp: %v", err)
		}
	}

	if err = d.Set("password", flattenSystemBugReportPassword(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	return nil
}

func flattenSystemBugReportFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemBugReportMailto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemBugReportUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemBugReportAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemBugReportServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemBugReportUsernameSmtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemBugReportPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemBugReport(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mailto"); ok {
		if setArgNil {
			obj["mailto"] = nil
		} else {

			t, err := expandSystemBugReportMailto(d, v, "mailto", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mailto"] = t
			}
		}
	}

	if v, ok := d.GetOk("username"); ok {
		if setArgNil {
			obj["username"] = nil
		} else {

			t, err := expandSystemBugReportUsername(d, v, "username", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["username"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth"); ok {
		if setArgNil {
			obj["auth"] = nil
		} else {

			t, err := expandSystemBugReportAuth(d, v, "auth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth"] = t
			}
		}
	}

	if v, ok := d.GetOk("server"); ok {
		if setArgNil {
			obj["server"] = nil
		} else {

			t, err := expandSystemBugReportServer(d, v, "server", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server"] = t
			}
		}
	}

	if v, ok := d.GetOk("username_smtp"); ok {
		if setArgNil {
			obj["username-smtp"] = nil
		} else {

			t, err := expandSystemBugReportUsernameSmtp(d, v, "username_smtp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["username-smtp"] = t
			}
		}
	}

	if v, ok := d.GetOk("password"); ok {
		if setArgNil {
			obj["password"] = nil
		} else {

			t, err := expandSystemBugReportPassword(d, v, "password", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["password"] = t
			}
		}
	}

	return &obj, nil
}
