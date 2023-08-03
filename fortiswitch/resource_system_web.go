// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure web attributes.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemWeb() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemWebUpdate,
		Read:   resourceSystemWebRead,
		Update: resourceSystemWebUpdate,
		Delete: resourceSystemWebDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"https_server_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"https_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"https_pki_required": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"gui_language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_ssl_versions": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemWebUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemWeb(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemWeb resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemWeb(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemWeb resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemWeb")
	}

	return resourceSystemWebRead(d, m)
}

func resourceSystemWebDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemWeb(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemWeb resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemWeb(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemWeb resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemWebRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemWeb(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemWeb resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemWeb(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemWeb resource from API: %v", err)
	}
	return nil
}

func flattenSystemWebHttpsServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemWebHttpsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemWebHttpsPkiRequired(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemWebHttpPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemWebGuiLanguage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemWebHttpsSslVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemWeb(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("https_server_cert", flattenSystemWebHttpsServerCert(o["https-server-cert"], d, "https_server_cert", sv)); err != nil {
		if !fortiAPIPatch(o["https-server-cert"]) {
			return fmt.Errorf("Error reading https_server_cert: %v", err)
		}
	}

	if err = d.Set("https_port", flattenSystemWebHttpsPort(o["https-port"], d, "https_port", sv)); err != nil {
		if !fortiAPIPatch(o["https-port"]) {
			return fmt.Errorf("Error reading https_port: %v", err)
		}
	}

	if err = d.Set("https_pki_required", flattenSystemWebHttpsPkiRequired(o["https-pki-required"], d, "https_pki_required", sv)); err != nil {
		if !fortiAPIPatch(o["https-pki-required"]) {
			return fmt.Errorf("Error reading https_pki_required: %v", err)
		}
	}

	if err = d.Set("http_port", flattenSystemWebHttpPort(o["http-port"], d, "http_port", sv)); err != nil {
		if !fortiAPIPatch(o["http-port"]) {
			return fmt.Errorf("Error reading http_port: %v", err)
		}
	}

	if err = d.Set("gui_language", flattenSystemWebGuiLanguage(o["gui-language"], d, "gui_language", sv)); err != nil {
		if !fortiAPIPatch(o["gui-language"]) {
			return fmt.Errorf("Error reading gui_language: %v", err)
		}
	}

	if err = d.Set("https_ssl_versions", flattenSystemWebHttpsSslVersions(o["https-ssl-versions"], d, "https_ssl_versions", sv)); err != nil {
		if !fortiAPIPatch(o["https-ssl-versions"]) {
			return fmt.Errorf("Error reading https_ssl_versions: %v", err)
		}
	}

	return nil
}

func flattenSystemWebFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemWebHttpsServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemWebHttpsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemWebHttpsPkiRequired(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemWebHttpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemWebGuiLanguage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemWebHttpsSslVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemWeb(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("https_server_cert"); ok {
		if setArgNil {
			obj["https-server-cert"] = nil
		} else {

			t, err := expandSystemWebHttpsServerCert(d, v, "https_server_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["https-server-cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("https_port"); ok {
		if setArgNil {
			obj["https-port"] = nil
		} else {

			t, err := expandSystemWebHttpsPort(d, v, "https_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["https-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("https_pki_required"); ok {
		if setArgNil {
			obj["https-pki-required"] = nil
		} else {

			t, err := expandSystemWebHttpsPkiRequired(d, v, "https_pki_required", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["https-pki-required"] = t
			}
		}
	}

	if v, ok := d.GetOk("http_port"); ok {
		if setArgNil {
			obj["http-port"] = nil
		} else {

			t, err := expandSystemWebHttpPort(d, v, "http_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["http-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("gui_language"); ok {
		if setArgNil {
			obj["gui-language"] = nil
		} else {

			t, err := expandSystemWebGuiLanguage(d, v, "gui_language", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gui-language"] = t
			}
		}
	}

	if v, ok := d.GetOk("https_ssl_versions"); ok {
		if setArgNil {
			obj["https-ssl-versions"] = nil
		} else {

			t, err := expandSystemWebHttpsSslVersions(d, v, "https_ssl_versions", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["https-ssl-versions"] = t
			}
		}
	}

	return &obj, nil
}
