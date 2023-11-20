// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: MACsec configuration profiles.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchMacsecProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchMacsecProfileCreate,
		Read:   resourceSwitchMacsecProfileRead,
		Update: resourceSwitchMacsecProfileUpdate,
		Delete: resourceSwitchMacsecProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_tls_ca_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"replay_protect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_tls_identity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"replay_window": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16777215),
				Optional:     true,
				Computed:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"macsec_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"include_mka_icv_ind": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_tls_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"cipher_suite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_policy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"security_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"exclude_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"eap_tls_radius_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"macsec_validate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mka_priority": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"encrypt_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mka_psk": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"crypto_alg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"mka_cak": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"mka_ckn": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"confident_offset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"include_macsec_sci": &schema.Schema{
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

func resourceSwitchMacsecProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchMacsecProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchMacsecProfile resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchMacsecProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchMacsecProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchMacsecProfile")
	}

	return resourceSwitchMacsecProfileRead(d, m)
}

func resourceSwitchMacsecProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchMacsecProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchMacsecProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchMacsecProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchMacsecProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchMacsecProfile")
	}

	return resourceSwitchMacsecProfileRead(d, m)
}

func resourceSwitchMacsecProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchMacsecProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchMacsecProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchMacsecProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchMacsecProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchMacsecProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchMacsecProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchMacsecProfile resource from API: %v", err)
	}
	return nil
}

func flattenSwitchMacsecProfileStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileEapTlsCaCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileReplayProtect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileEapTlsIdentity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileReplayWindow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileMacsecMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileIncludeMkaIcvInd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileEapTlsCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileCipherSuite(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileTrafficPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenSwitchMacsecProfileTrafficPolicyStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchMacsecProfileTrafficPolicyName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_policy"
		if _, ok := i["security-policy"]; ok {

			tmp["security_policy"] = flattenSwitchMacsecProfileTrafficPolicySecurityPolicy(i["security-policy"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclude_protocol"
		if _, ok := i["exclude-protocol"]; ok {

			tmp["exclude_protocol"] = flattenSwitchMacsecProfileTrafficPolicyExcludeProtocol(i["exclude-protocol"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchMacsecProfileTrafficPolicyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileTrafficPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileTrafficPolicySecurityPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileTrafficPolicyExcludeProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileEapTlsRadiusServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileCipher_Suite(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileMacsecValidate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileMkaPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileEncryptTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileMkaPsk(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenSwitchMacsecProfileMkaPskStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "crypto_alg"
		if _, ok := i["crypto-alg"]; ok {

			tmp["crypto_alg"] = flattenSwitchMacsecProfileMkaPskCryptoAlg(i["crypto-alg"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchMacsecProfileMkaPskName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mka_cak"
		if _, ok := i["mka-cak"]; ok {

			tmp["mka_cak"] = flattenSwitchMacsecProfileMkaPskMkaCak(i["mka-cak"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mka_ckn"
		if _, ok := i["mka-ckn"]; ok {

			tmp["mka_ckn"] = flattenSwitchMacsecProfileMkaPskMkaCkn(i["mka-ckn"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchMacsecProfileMkaPskStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileMkaPskCryptoAlg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileMkaPskName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileMkaPskMkaCak(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileMkaPskMkaCkn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileConfidentOffset(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchMacsecProfileIncludeMacsecSci(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchMacsecProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSwitchMacsecProfileStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("eap_tls_ca_cert", flattenSwitchMacsecProfileEapTlsCaCert(o["eap-tls-ca-cert"], d, "eap_tls_ca_cert", sv)); err != nil {
		if !fortiAPIPatch(o["eap-tls-ca-cert"]) {
			return fmt.Errorf("Error reading eap_tls_ca_cert: %v", err)
		}
	}

	if err = d.Set("replay_protect", flattenSwitchMacsecProfileReplayProtect(o["replay-protect"], d, "replay_protect", sv)); err != nil {
		if !fortiAPIPatch(o["replay-protect"]) {
			return fmt.Errorf("Error reading replay_protect: %v", err)
		}
	}

	if err = d.Set("eap_tls_identity", flattenSwitchMacsecProfileEapTlsIdentity(o["eap-tls-identity"], d, "eap_tls_identity", sv)); err != nil {
		if !fortiAPIPatch(o["eap-tls-identity"]) {
			return fmt.Errorf("Error reading eap_tls_identity: %v", err)
		}
	}

	if err = d.Set("replay_window", flattenSwitchMacsecProfileReplayWindow(o["replay-window"], d, "replay_window", sv)); err != nil {
		if !fortiAPIPatch(o["replay-window"]) {
			return fmt.Errorf("Error reading replay_window: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchMacsecProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("macsec_mode", flattenSwitchMacsecProfileMacsecMode(o["macsec-mode"], d, "macsec_mode", sv)); err != nil {
		if !fortiAPIPatch(o["macsec-mode"]) {
			return fmt.Errorf("Error reading macsec_mode: %v", err)
		}
	}

	if err = d.Set("include_mka_icv_ind", flattenSwitchMacsecProfileIncludeMkaIcvInd(o["include-mka-icv-ind"], d, "include_mka_icv_ind", sv)); err != nil {
		if !fortiAPIPatch(o["include-mka-icv-ind"]) {
			return fmt.Errorf("Error reading include_mka_icv_ind: %v", err)
		}
	}

	if err = d.Set("eap_tls_cert", flattenSwitchMacsecProfileEapTlsCert(o["eap-tls-cert"], d, "eap_tls_cert", sv)); err != nil {
		if !fortiAPIPatch(o["eap-tls-cert"]) {
			return fmt.Errorf("Error reading eap_tls_cert: %v", err)
		}
	}

	if err = d.Set("cipher_suite", flattenSwitchMacsecProfileCipherSuite(o["cipher-suite"], d, "cipher_suite", sv)); err != nil {
		if !fortiAPIPatch(o["cipher-suite"]) {
			return fmt.Errorf("Error reading cipher_suite: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("traffic_policy", flattenSwitchMacsecProfileTrafficPolicy(o["traffic-policy"], d, "traffic_policy", sv)); err != nil {
			if !fortiAPIPatch(o["traffic-policy"]) {
				return fmt.Errorf("Error reading traffic_policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("traffic_policy"); ok {
			if err = d.Set("traffic_policy", flattenSwitchMacsecProfileTrafficPolicy(o["traffic-policy"], d, "traffic_policy", sv)); err != nil {
				if !fortiAPIPatch(o["traffic-policy"]) {
					return fmt.Errorf("Error reading traffic_policy: %v", err)
				}
			}
		}
	}

	if err = d.Set("eap_tls_radius_server", flattenSwitchMacsecProfileEapTlsRadiusServer(o["eap-tls-radius-server"], d, "eap_tls_radius_server", sv)); err != nil {
		if !fortiAPIPatch(o["eap-tls-radius-server"]) {
			return fmt.Errorf("Error reading eap_tls_radius_server: %v", err)
		}
	}

	if err = d.Set("cipher_suite", flattenSwitchMacsecProfileCipher_Suite(o["cipher_suite"], d, "cipher_suite", sv)); err != nil {
		if !fortiAPIPatch(o["cipher_suite"]) {
			return fmt.Errorf("Error reading cipher_suite: %v", err)
		}
	}

	if err = d.Set("macsec_validate", flattenSwitchMacsecProfileMacsecValidate(o["macsec-validate"], d, "macsec_validate", sv)); err != nil {
		if !fortiAPIPatch(o["macsec-validate"]) {
			return fmt.Errorf("Error reading macsec_validate: %v", err)
		}
	}

	if err = d.Set("mka_priority", flattenSwitchMacsecProfileMkaPriority(o["mka-priority"], d, "mka_priority", sv)); err != nil {
		if !fortiAPIPatch(o["mka-priority"]) {
			return fmt.Errorf("Error reading mka_priority: %v", err)
		}
	}

	if err = d.Set("encrypt_traffic", flattenSwitchMacsecProfileEncryptTraffic(o["encrypt-traffic"], d, "encrypt_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["encrypt-traffic"]) {
			return fmt.Errorf("Error reading encrypt_traffic: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mka_psk", flattenSwitchMacsecProfileMkaPsk(o["mka-psk"], d, "mka_psk", sv)); err != nil {
			if !fortiAPIPatch(o["mka-psk"]) {
				return fmt.Errorf("Error reading mka_psk: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mka_psk"); ok {
			if err = d.Set("mka_psk", flattenSwitchMacsecProfileMkaPsk(o["mka-psk"], d, "mka_psk", sv)); err != nil {
				if !fortiAPIPatch(o["mka-psk"]) {
					return fmt.Errorf("Error reading mka_psk: %v", err)
				}
			}
		}
	}

	if err = d.Set("confident_offset", flattenSwitchMacsecProfileConfidentOffset(o["confident-offset"], d, "confident_offset", sv)); err != nil {
		if !fortiAPIPatch(o["confident-offset"]) {
			return fmt.Errorf("Error reading confident_offset: %v", err)
		}
	}

	if err = d.Set("include_macsec_sci", flattenSwitchMacsecProfileIncludeMacsecSci(o["include-macsec-sci"], d, "include_macsec_sci", sv)); err != nil {
		if !fortiAPIPatch(o["include-macsec-sci"]) {
			return fmt.Errorf("Error reading include_macsec_sci: %v", err)
		}
	}

	return nil
}

func flattenSwitchMacsecProfileFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSwitchMacsecProfileStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileEapTlsCaCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileReplayProtect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileEapTlsIdentity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileReplayWindow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileMacsecMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileIncludeMkaIcvInd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileEapTlsCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileCipherSuite(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileTrafficPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandSwitchMacsecProfileTrafficPolicyStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSwitchMacsecProfileTrafficPolicyName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_policy"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["security-policy"], _ = expandSwitchMacsecProfileTrafficPolicySecurityPolicy(d, i["security_policy"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclude_protocol"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["exclude-protocol"], _ = expandSwitchMacsecProfileTrafficPolicyExcludeProtocol(d, i["exclude_protocol"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchMacsecProfileTrafficPolicyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileTrafficPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileTrafficPolicySecurityPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileTrafficPolicyExcludeProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileEapTlsRadiusServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileCipher_Suite(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileMacsecValidate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileMkaPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileEncryptTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileMkaPsk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandSwitchMacsecProfileMkaPskStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "crypto_alg"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["crypto-alg"], _ = expandSwitchMacsecProfileMkaPskCryptoAlg(d, i["crypto_alg"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSwitchMacsecProfileMkaPskName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mka_cak"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mka-cak"], _ = expandSwitchMacsecProfileMkaPskMkaCak(d, i["mka_cak"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mka_ckn"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mka-ckn"], _ = expandSwitchMacsecProfileMkaPskMkaCkn(d, i["mka_ckn"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchMacsecProfileMkaPskStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileMkaPskCryptoAlg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileMkaPskName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileMkaPskMkaCak(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileMkaPskMkaCkn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileConfidentOffset(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchMacsecProfileIncludeMacsecSci(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchMacsecProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {

		t, err := expandSwitchMacsecProfileStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("eap_tls_ca_cert"); ok {

		t, err := expandSwitchMacsecProfileEapTlsCaCert(d, v, "eap_tls_ca_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-tls-ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("replay_protect"); ok {

		t, err := expandSwitchMacsecProfileReplayProtect(d, v, "replay_protect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replay-protect"] = t
		}
	}

	if v, ok := d.GetOk("eap_tls_identity"); ok {

		t, err := expandSwitchMacsecProfileEapTlsIdentity(d, v, "eap_tls_identity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-tls-identity"] = t
		}
	}

	if v, ok := d.GetOkExists("replay_window"); ok {

		t, err := expandSwitchMacsecProfileReplayWindow(d, v, "replay_window", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replay-window"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchMacsecProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("macsec_mode"); ok {

		t, err := expandSwitchMacsecProfileMacsecMode(d, v, "macsec_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macsec-mode"] = t
		}
	}

	if v, ok := d.GetOk("include_mka_icv_ind"); ok {

		t, err := expandSwitchMacsecProfileIncludeMkaIcvInd(d, v, "include_mka_icv_ind", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-mka-icv-ind"] = t
		}
	}

	if v, ok := d.GetOk("eap_tls_cert"); ok {

		t, err := expandSwitchMacsecProfileEapTlsCert(d, v, "eap_tls_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-tls-cert"] = t
		}
	}

	if v, ok := d.GetOk("cipher_suite"); ok {

		t, err := expandSwitchMacsecProfileCipherSuite(d, v, "cipher_suite", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cipher-suite"] = t
		}
	}

	if v, ok := d.GetOk("traffic_policy"); ok || d.HasChange("traffic_policy") {

		t, err := expandSwitchMacsecProfileTrafficPolicy(d, v, "traffic_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-policy"] = t
		}
	}

	if v, ok := d.GetOk("eap_tls_radius_server"); ok {

		t, err := expandSwitchMacsecProfileEapTlsRadiusServer(d, v, "eap_tls_radius_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-tls-radius-server"] = t
		}
	}

	if v, ok := d.GetOk("cipher_suite"); ok {

		t, err := expandSwitchMacsecProfileCipher_Suite(d, v, "cipher_suite", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cipher_suite"] = t
		}
	}

	if v, ok := d.GetOk("macsec_validate"); ok {

		t, err := expandSwitchMacsecProfileMacsecValidate(d, v, "macsec_validate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macsec-validate"] = t
		}
	}

	if v, ok := d.GetOkExists("mka_priority"); ok {

		t, err := expandSwitchMacsecProfileMkaPriority(d, v, "mka_priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mka-priority"] = t
		}
	}

	if v, ok := d.GetOk("encrypt_traffic"); ok {

		t, err := expandSwitchMacsecProfileEncryptTraffic(d, v, "encrypt_traffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encrypt-traffic"] = t
		}
	}

	if v, ok := d.GetOk("mka_psk"); ok || d.HasChange("mka_psk") {

		t, err := expandSwitchMacsecProfileMkaPsk(d, v, "mka_psk", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mka-psk"] = t
		}
	}

	if v, ok := d.GetOk("confident_offset"); ok {

		t, err := expandSwitchMacsecProfileConfidentOffset(d, v, "confident_offset", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["confident-offset"] = t
		}
	}

	if v, ok := d.GetOk("include_macsec_sci"); ok {

		t, err := expandSwitchMacsecProfileIncludeMacsecSci(d, v, "include_macsec_sci", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-macsec-sci"] = t
		}
	}

	return &obj, nil
}
