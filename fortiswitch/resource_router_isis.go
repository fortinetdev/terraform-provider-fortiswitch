// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: ISIS configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterIsis() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterIsisUpdate,
		Read:   resourceRouterIsisRead,
		Update: resourceRouterIsisUpdate,
		Delete: resourceRouterIsisDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"default_information_metric6": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16777215),
				Optional:     true,
				Computed:     true,
			},
			"auth_sendonly_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_information_originate6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"summary_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"metric_style": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redistribute6_l1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lsp_refresh_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"max_lsp_lifetime": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(350, 65535),
				Optional:     true,
				Computed:     true,
			},
			"ignore_attached_bit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redistribute6_l1_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"auth_keychain_area": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"auth_mode_area": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_information_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redistribute_l1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"net": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"net": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"summary_address6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"auth_mode_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lsp_gen_interval_l1": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"lsp_gen_interval_l2": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"redistribute_l1_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"overload_bit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_password_area": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 59),
				Optional:     true,
			},
			"default_information_metric": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16777215),
				Optional:     true,
				Computed:     true,
			},
			"default_information_originate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_password_hello": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 59),
							Optional:     true,
						},
						"priority_l2": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"priority_l1": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"hello_multiplier_l2": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),
							Optional:     true,
							Computed:     true,
						},
						"hello_multiplier_l1": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),
							Optional:     true,
							Computed:     true,
						},
						"auth_mode_hello": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"passive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"circuit_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bfd6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wide_metric_l1": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16777214),
							Optional:     true,
							Computed:     true,
						},
						"wide_metric_l2": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16777214),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metric_l1": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 63),
							Optional:     true,
							Computed:     true,
						},
						"metric_l2": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 63),
							Optional:     true,
							Computed:     true,
						},
						"status6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"auth_keychain_hello": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"hello_interval_l2": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"csnp_interval_l2": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"csnp_interval_l1": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"hello_padding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hello_interval_l1": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"auth_keychain_domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"spf_interval_exp_l2": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"spf_interval_exp_l1": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"auth_password_domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 59),
				Optional:     true,
			},
			"redistribute": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metric": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"routemap": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_sendonly_area": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_information_level6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redistribute6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metric": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"routemap": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"protocol": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"log_neighbour_changes": &schema.Schema{
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

func resourceRouterIsisUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterIsis(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterIsis resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterIsis(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterIsis resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterIsis")
	}

	return resourceRouterIsisRead(d, m)
}

func resourceRouterIsisDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterIsis(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating RouterIsis resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterIsis(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing RouterIsis resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterIsisRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterIsis(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterIsis resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterIsis(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterIsis resource from API: %v", err)
	}
	return nil
}

func flattenRouterIsisDefaultInformationMetric6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisAuthSendonlyDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisDefaultInformationOriginate6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisSummaryAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = flattenRouterIsisSummaryAddressPrefix(i["prefix"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterIsisSummaryAddressId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {

			tmp["level"] = flattenRouterIsisSummaryAddressLevel(i["level"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterIsisSummaryAddressPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterIsisSummaryAddressId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisSummaryAddressLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisMetricStyle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6L1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisLspRefreshInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisMaxLspLifetime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisIgnoreAttachedBit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6L1List(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisAuthKeychainArea(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisAuthModeArea(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisDefaultInformationLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistributeL1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisNet(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "net"
		if _, ok := i["net"]; ok {

			tmp["net"] = flattenRouterIsisNetNet(i["net"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "net", d)
	return result
}

func flattenRouterIsisNetNet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisSummaryAddress6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {

			tmp["level"] = flattenRouterIsisSummaryAddress6Level(i["level"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {

			tmp["prefix6"] = flattenRouterIsisSummaryAddress6Prefix6(i["prefix6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterIsisSummaryAddress6Id(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterIsisSummaryAddress6Level(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisSummaryAddress6Prefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisSummaryAddress6Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisAuthModeDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisLspGenIntervalL1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisLspGenIntervalL2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistributeL1List(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisOverloadBit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisAuthPasswordArea(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisDefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_password_hello"
		if _, ok := i["auth-password-hello"]; ok {

			tmp["auth_password_hello"] = flattenRouterIsisInterfaceAuthPasswordHello(i["auth-password-hello"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_l2"
		if _, ok := i["priority-l2"]; ok {

			tmp["priority_l2"] = flattenRouterIsisInterfacePriorityL2(i["priority-l2"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_l1"
		if _, ok := i["priority-l1"]; ok {

			tmp["priority_l1"] = flattenRouterIsisInterfacePriorityL1(i["priority-l1"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier_l2"
		if _, ok := i["hello-multiplier-l2"]; ok {

			tmp["hello_multiplier_l2"] = flattenRouterIsisInterfaceHelloMultiplierL2(i["hello-multiplier-l2"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier_l1"
		if _, ok := i["hello-multiplier-l1"]; ok {

			tmp["hello_multiplier_l1"] = flattenRouterIsisInterfaceHelloMultiplierL1(i["hello-multiplier-l1"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode_hello"
		if _, ok := i["auth-mode-hello"]; ok {

			tmp["auth_mode_hello"] = flattenRouterIsisInterfaceAuthModeHello(i["auth-mode-hello"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {

			tmp["bfd"] = flattenRouterIsisInterfaceBfd(i["bfd"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := i["passive"]; ok {

			tmp["passive"] = flattenRouterIsisInterfacePassive(i["passive"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_type"
		if _, ok := i["circuit-type"]; ok {

			tmp["circuit_type"] = flattenRouterIsisInterfaceCircuitType(i["circuit-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd6"
		if _, ok := i["bfd6"]; ok {

			tmp["bfd6"] = flattenRouterIsisInterfaceBfd6(i["bfd6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wide_metric_l1"
		if _, ok := i["wide-metric-l1"]; ok {

			tmp["wide_metric_l1"] = flattenRouterIsisInterfaceWideMetricL1(i["wide-metric-l1"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wide_metric_l2"
		if _, ok := i["wide-metric-l2"]; ok {

			tmp["wide_metric_l2"] = flattenRouterIsisInterfaceWideMetricL2(i["wide-metric-l2"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenRouterIsisInterfaceStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_l1"
		if _, ok := i["metric-l1"]; ok {

			tmp["metric_l1"] = flattenRouterIsisInterfaceMetricL1(i["metric-l1"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_l2"
		if _, ok := i["metric-l2"]; ok {

			tmp["metric_l2"] = flattenRouterIsisInterfaceMetricL2(i["metric-l2"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status6"
		if _, ok := i["status6"]; ok {

			tmp["status6"] = flattenRouterIsisInterfaceStatus6(i["status6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenRouterIsisInterfaceName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain_hello"
		if _, ok := i["auth-keychain-hello"]; ok {

			tmp["auth_keychain_hello"] = flattenRouterIsisInterfaceAuthKeychainHello(i["auth-keychain-hello"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval_l2"
		if _, ok := i["hello-interval-l2"]; ok {

			tmp["hello_interval_l2"] = flattenRouterIsisInterfaceHelloIntervalL2(i["hello-interval-l2"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csnp_interval_l2"
		if _, ok := i["csnp-interval-l2"]; ok {

			tmp["csnp_interval_l2"] = flattenRouterIsisInterfaceCsnpIntervalL2(i["csnp-interval-l2"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csnp_interval_l1"
		if _, ok := i["csnp-interval-l1"]; ok {

			tmp["csnp_interval_l1"] = flattenRouterIsisInterfaceCsnpIntervalL1(i["csnp-interval-l1"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_padding"
		if _, ok := i["hello-padding"]; ok {

			tmp["hello_padding"] = flattenRouterIsisInterfaceHelloPadding(i["hello-padding"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval_l1"
		if _, ok := i["hello-interval-l1"]; ok {

			tmp["hello_interval_l1"] = flattenRouterIsisInterfaceHelloIntervalL1(i["hello-interval-l1"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterIsisInterfaceAuthPasswordHello(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfacePriorityL2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfacePriorityL1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceHelloMultiplierL2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceHelloMultiplierL1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceAuthModeHello(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfacePassive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceCircuitType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceBfd6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceWideMetricL1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceWideMetricL2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceMetricL1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceMetricL2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceStatus6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceAuthKeychainHello(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceHelloIntervalL2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceCsnpIntervalL2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceCsnpIntervalL1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceHelloPadding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisInterfaceHelloIntervalL1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisAuthKeychainDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisSpfIntervalExpL2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisSpfIntervalExpL1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisAuthPasswordDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistribute(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["status"] = flattenRouterIsisRedistributeStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {

			tmp["protocol"] = flattenRouterIsisRedistributeProtocol(i["protocol"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {

			tmp["level"] = flattenRouterIsisRedistributeLevel(i["level"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {

			tmp["metric"] = flattenRouterIsisRedistributeMetric(i["metric"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {

			tmp["routemap"] = flattenRouterIsisRedistributeRoutemap(i["routemap"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := i["metric-type"]; ok {

			tmp["metric_type"] = flattenRouterIsisRedistributeMetricType(i["metric-type"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "protocol", d)
	return result
}

func flattenRouterIsisRedistributeStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistributeProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistributeLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistributeMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistributeMetricType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRouterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisAuthSendonlyArea(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisIsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisDefaultInformationLevel6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["status"] = flattenRouterIsisRedistribute6Status(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {

			tmp["metric"] = flattenRouterIsisRedistribute6Metric(i["metric"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {

			tmp["routemap"] = flattenRouterIsisRedistribute6Routemap(i["routemap"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {

			tmp["protocol"] = flattenRouterIsisRedistribute6Protocol(i["protocol"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {

			tmp["level"] = flattenRouterIsisRedistribute6Level(i["level"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "protocol", d)
	return result
}

func flattenRouterIsisRedistribute6Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6Metric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6Routemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6Protocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6Level(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterIsisLogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterIsis(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("default_information_metric6", flattenRouterIsisDefaultInformationMetric6(o["default-information-metric6"], d, "default_information_metric6", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-metric6"]) {
			return fmt.Errorf("Error reading default_information_metric6: %v", err)
		}
	}

	if err = d.Set("auth_sendonly_domain", flattenRouterIsisAuthSendonlyDomain(o["auth-sendonly-domain"], d, "auth_sendonly_domain", sv)); err != nil {
		if !fortiAPIPatch(o["auth-sendonly-domain"]) {
			return fmt.Errorf("Error reading auth_sendonly_domain: %v", err)
		}
	}

	if err = d.Set("default_information_originate6", flattenRouterIsisDefaultInformationOriginate6(o["default-information-originate6"], d, "default_information_originate6", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-originate6"]) {
			return fmt.Errorf("Error reading default_information_originate6: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("summary_address", flattenRouterIsisSummaryAddress(o["summary-address"], d, "summary_address", sv)); err != nil {
			if !fortiAPIPatch(o["summary-address"]) {
				return fmt.Errorf("Error reading summary_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("summary_address"); ok {
			if err = d.Set("summary_address", flattenRouterIsisSummaryAddress(o["summary-address"], d, "summary_address", sv)); err != nil {
				if !fortiAPIPatch(o["summary-address"]) {
					return fmt.Errorf("Error reading summary_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("metric_style", flattenRouterIsisMetricStyle(o["metric-style"], d, "metric_style", sv)); err != nil {
		if !fortiAPIPatch(o["metric-style"]) {
			return fmt.Errorf("Error reading metric_style: %v", err)
		}
	}

	if err = d.Set("redistribute6_l1", flattenRouterIsisRedistribute6L1(o["redistribute6-l1"], d, "redistribute6_l1", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute6-l1"]) {
			return fmt.Errorf("Error reading redistribute6_l1: %v", err)
		}
	}

	if err = d.Set("lsp_refresh_interval", flattenRouterIsisLspRefreshInterval(o["lsp-refresh-interval"], d, "lsp_refresh_interval", sv)); err != nil {
		if !fortiAPIPatch(o["lsp-refresh-interval"]) {
			return fmt.Errorf("Error reading lsp_refresh_interval: %v", err)
		}
	}

	if err = d.Set("max_lsp_lifetime", flattenRouterIsisMaxLspLifetime(o["max-lsp-lifetime"], d, "max_lsp_lifetime", sv)); err != nil {
		if !fortiAPIPatch(o["max-lsp-lifetime"]) {
			return fmt.Errorf("Error reading max_lsp_lifetime: %v", err)
		}
	}

	if err = d.Set("ignore_attached_bit", flattenRouterIsisIgnoreAttachedBit(o["ignore-attached-bit"], d, "ignore_attached_bit", sv)); err != nil {
		if !fortiAPIPatch(o["ignore-attached-bit"]) {
			return fmt.Errorf("Error reading ignore_attached_bit: %v", err)
		}
	}

	if err = d.Set("redistribute6_l1_list", flattenRouterIsisRedistribute6L1List(o["redistribute6-l1-list"], d, "redistribute6_l1_list", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute6-l1-list"]) {
			return fmt.Errorf("Error reading redistribute6_l1_list: %v", err)
		}
	}

	if err = d.Set("auth_keychain_area", flattenRouterIsisAuthKeychainArea(o["auth-keychain-area"], d, "auth_keychain_area", sv)); err != nil {
		if !fortiAPIPatch(o["auth-keychain-area"]) {
			return fmt.Errorf("Error reading auth_keychain_area: %v", err)
		}
	}

	if err = d.Set("auth_mode_area", flattenRouterIsisAuthModeArea(o["auth-mode-area"], d, "auth_mode_area", sv)); err != nil {
		if !fortiAPIPatch(o["auth-mode-area"]) {
			return fmt.Errorf("Error reading auth_mode_area: %v", err)
		}
	}

	if err = d.Set("default_information_level", flattenRouterIsisDefaultInformationLevel(o["default-information-level"], d, "default_information_level", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-level"]) {
			return fmt.Errorf("Error reading default_information_level: %v", err)
		}
	}

	if err = d.Set("redistribute_l1", flattenRouterIsisRedistributeL1(o["redistribute-l1"], d, "redistribute_l1", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute-l1"]) {
			return fmt.Errorf("Error reading redistribute_l1: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("net", flattenRouterIsisNet(o["net"], d, "net", sv)); err != nil {
			if !fortiAPIPatch(o["net"]) {
				return fmt.Errorf("Error reading net: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("net"); ok {
			if err = d.Set("net", flattenRouterIsisNet(o["net"], d, "net", sv)); err != nil {
				if !fortiAPIPatch(o["net"]) {
					return fmt.Errorf("Error reading net: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("summary_address6", flattenRouterIsisSummaryAddress6(o["summary-address6"], d, "summary_address6", sv)); err != nil {
			if !fortiAPIPatch(o["summary-address6"]) {
				return fmt.Errorf("Error reading summary_address6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("summary_address6"); ok {
			if err = d.Set("summary_address6", flattenRouterIsisSummaryAddress6(o["summary-address6"], d, "summary_address6", sv)); err != nil {
				if !fortiAPIPatch(o["summary-address6"]) {
					return fmt.Errorf("Error reading summary_address6: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth_mode_domain", flattenRouterIsisAuthModeDomain(o["auth-mode-domain"], d, "auth_mode_domain", sv)); err != nil {
		if !fortiAPIPatch(o["auth-mode-domain"]) {
			return fmt.Errorf("Error reading auth_mode_domain: %v", err)
		}
	}

	if err = d.Set("lsp_gen_interval_l1", flattenRouterIsisLspGenIntervalL1(o["lsp-gen-interval-l1"], d, "lsp_gen_interval_l1", sv)); err != nil {
		if !fortiAPIPatch(o["lsp-gen-interval-l1"]) {
			return fmt.Errorf("Error reading lsp_gen_interval_l1: %v", err)
		}
	}

	if err = d.Set("lsp_gen_interval_l2", flattenRouterIsisLspGenIntervalL2(o["lsp-gen-interval-l2"], d, "lsp_gen_interval_l2", sv)); err != nil {
		if !fortiAPIPatch(o["lsp-gen-interval-l2"]) {
			return fmt.Errorf("Error reading lsp_gen_interval_l2: %v", err)
		}
	}

	if err = d.Set("redistribute_l1_list", flattenRouterIsisRedistributeL1List(o["redistribute-l1-list"], d, "redistribute_l1_list", sv)); err != nil {
		if !fortiAPIPatch(o["redistribute-l1-list"]) {
			return fmt.Errorf("Error reading redistribute_l1_list: %v", err)
		}
	}

	if err = d.Set("overload_bit", flattenRouterIsisOverloadBit(o["overload-bit"], d, "overload_bit", sv)); err != nil {
		if !fortiAPIPatch(o["overload-bit"]) {
			return fmt.Errorf("Error reading overload_bit: %v", err)
		}
	}

	if err = d.Set("auth_password_area", flattenRouterIsisAuthPasswordArea(o["auth-password-area"], d, "auth_password_area", sv)); err != nil {
		if !fortiAPIPatch(o["auth-password-area"]) {
			return fmt.Errorf("Error reading auth_password_area: %v", err)
		}
	}

	if err = d.Set("default_information_metric", flattenRouterIsisDefaultInformationMetric(o["default-information-metric"], d, "default_information_metric", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-metric"]) {
			return fmt.Errorf("Error reading default_information_metric: %v", err)
		}
	}

	if err = d.Set("default_information_originate", flattenRouterIsisDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("interface", flattenRouterIsisInterface(o["interface"], d, "interface", sv)); err != nil {
			if !fortiAPIPatch(o["interface"]) {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenRouterIsisInterface(o["interface"], d, "interface", sv)); err != nil {
				if !fortiAPIPatch(o["interface"]) {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth_keychain_domain", flattenRouterIsisAuthKeychainDomain(o["auth-keychain-domain"], d, "auth_keychain_domain", sv)); err != nil {
		if !fortiAPIPatch(o["auth-keychain-domain"]) {
			return fmt.Errorf("Error reading auth_keychain_domain: %v", err)
		}
	}

	if err = d.Set("spf_interval_exp_l2", flattenRouterIsisSpfIntervalExpL2(o["spf-interval-exp-l2"], d, "spf_interval_exp_l2", sv)); err != nil {
		if !fortiAPIPatch(o["spf-interval-exp-l2"]) {
			return fmt.Errorf("Error reading spf_interval_exp_l2: %v", err)
		}
	}

	if err = d.Set("spf_interval_exp_l1", flattenRouterIsisSpfIntervalExpL1(o["spf-interval-exp-l1"], d, "spf_interval_exp_l1", sv)); err != nil {
		if !fortiAPIPatch(o["spf-interval-exp-l1"]) {
			return fmt.Errorf("Error reading spf_interval_exp_l1: %v", err)
		}
	}

	if err = d.Set("auth_password_domain", flattenRouterIsisAuthPasswordDomain(o["auth-password-domain"], d, "auth_password_domain", sv)); err != nil {
		if !fortiAPIPatch(o["auth-password-domain"]) {
			return fmt.Errorf("Error reading auth_password_domain: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("redistribute", flattenRouterIsisRedistribute(o["redistribute"], d, "redistribute", sv)); err != nil {
			if !fortiAPIPatch(o["redistribute"]) {
				return fmt.Errorf("Error reading redistribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute"); ok {
			if err = d.Set("redistribute", flattenRouterIsisRedistribute(o["redistribute"], d, "redistribute", sv)); err != nil {
				if !fortiAPIPatch(o["redistribute"]) {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			}
		}
	}

	if err = d.Set("router_id", flattenRouterIsisRouterId(o["router-id"], d, "router_id", sv)); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("auth_sendonly_area", flattenRouterIsisAuthSendonlyArea(o["auth-sendonly-area"], d, "auth_sendonly_area", sv)); err != nil {
		if !fortiAPIPatch(o["auth-sendonly-area"]) {
			return fmt.Errorf("Error reading auth_sendonly_area: %v", err)
		}
	}

	if err = d.Set("is_type", flattenRouterIsisIsType(o["is-type"], d, "is_type", sv)); err != nil {
		if !fortiAPIPatch(o["is-type"]) {
			return fmt.Errorf("Error reading is_type: %v", err)
		}
	}

	if err = d.Set("default_information_level6", flattenRouterIsisDefaultInformationLevel6(o["default-information-level6"], d, "default_information_level6", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-level6"]) {
			return fmt.Errorf("Error reading default_information_level6: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("redistribute6", flattenRouterIsisRedistribute6(o["redistribute6"], d, "redistribute6", sv)); err != nil {
			if !fortiAPIPatch(o["redistribute6"]) {
				return fmt.Errorf("Error reading redistribute6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute6"); ok {
			if err = d.Set("redistribute6", flattenRouterIsisRedistribute6(o["redistribute6"], d, "redistribute6", sv)); err != nil {
				if !fortiAPIPatch(o["redistribute6"]) {
					return fmt.Errorf("Error reading redistribute6: %v", err)
				}
			}
		}
	}

	if err = d.Set("log_neighbour_changes", flattenRouterIsisLogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes", sv)); err != nil {
		if !fortiAPIPatch(o["log-neighbour-changes"]) {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	return nil
}

func flattenRouterIsisFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandRouterIsisDefaultInformationMetric6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAuthSendonlyDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisDefaultInformationOriginate6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisSummaryAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix"], _ = expandRouterIsisSummaryAddressPrefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterIsisSummaryAddressId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["level"], _ = expandRouterIsisSummaryAddressLevel(d, i["level"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterIsisSummaryAddressPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisSummaryAddressId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisSummaryAddressLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisMetricStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6L1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisLspRefreshInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisMaxLspLifetime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIgnoreAttachedBit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6L1List(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAuthKeychainArea(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAuthModeArea(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisDefaultInformationLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeL1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisNet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "net"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["net"], _ = expandRouterIsisNetNet(d, i["net"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterIsisNetNet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisSummaryAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["level"], _ = expandRouterIsisSummaryAddress6Level(d, i["level"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix6"], _ = expandRouterIsisSummaryAddress6Prefix6(d, i["prefix6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterIsisSummaryAddress6Id(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterIsisSummaryAddress6Level(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisSummaryAddress6Prefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisSummaryAddress6Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAuthModeDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisLspGenIntervalL1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisLspGenIntervalL2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeL1List(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisOverloadBit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAuthPasswordArea(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisDefaultInformationMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_password_hello"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["auth-password-hello"], _ = expandRouterIsisInterfaceAuthPasswordHello(d, i["auth_password_hello"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_l2"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["priority-l2"], _ = expandRouterIsisInterfacePriorityL2(d, i["priority_l2"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_l1"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["priority-l1"], _ = expandRouterIsisInterfacePriorityL1(d, i["priority_l1"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier_l2"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["hello-multiplier-l2"], _ = expandRouterIsisInterfaceHelloMultiplierL2(d, i["hello_multiplier_l2"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier_l1"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["hello-multiplier-l1"], _ = expandRouterIsisInterfaceHelloMultiplierL1(d, i["hello_multiplier_l1"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode_hello"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["auth-mode-hello"], _ = expandRouterIsisInterfaceAuthModeHello(d, i["auth_mode_hello"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["bfd"], _ = expandRouterIsisInterfaceBfd(d, i["bfd"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["passive"], _ = expandRouterIsisInterfacePassive(d, i["passive"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["circuit-type"], _ = expandRouterIsisInterfaceCircuitType(d, i["circuit_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["bfd6"], _ = expandRouterIsisInterfaceBfd6(d, i["bfd6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wide_metric_l1"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["wide-metric-l1"], _ = expandRouterIsisInterfaceWideMetricL1(d, i["wide_metric_l1"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wide_metric_l2"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["wide-metric-l2"], _ = expandRouterIsisInterfaceWideMetricL2(d, i["wide_metric_l2"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandRouterIsisInterfaceStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_l1"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["metric-l1"], _ = expandRouterIsisInterfaceMetricL1(d, i["metric_l1"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_l2"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["metric-l2"], _ = expandRouterIsisInterfaceMetricL2(d, i["metric_l2"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status6"], _ = expandRouterIsisInterfaceStatus6(d, i["status6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandRouterIsisInterfaceName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain_hello"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["auth-keychain-hello"], _ = expandRouterIsisInterfaceAuthKeychainHello(d, i["auth_keychain_hello"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval_l2"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["hello-interval-l2"], _ = expandRouterIsisInterfaceHelloIntervalL2(d, i["hello_interval_l2"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csnp_interval_l2"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["csnp-interval-l2"], _ = expandRouterIsisInterfaceCsnpIntervalL2(d, i["csnp_interval_l2"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csnp_interval_l1"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["csnp-interval-l1"], _ = expandRouterIsisInterfaceCsnpIntervalL1(d, i["csnp_interval_l1"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_padding"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["hello-padding"], _ = expandRouterIsisInterfaceHelloPadding(d, i["hello_padding"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval_l1"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["hello-interval-l1"], _ = expandRouterIsisInterfaceHelloIntervalL1(d, i["hello_interval_l1"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterIsisInterfaceAuthPasswordHello(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfacePriorityL2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfacePriorityL1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceHelloMultiplierL2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceHelloMultiplierL1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceAuthModeHello(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfacePassive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceCircuitType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceBfd6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceWideMetricL1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceWideMetricL2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceMetricL1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceMetricL2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceStatus6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceAuthKeychainHello(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceHelloIntervalL2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceCsnpIntervalL2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceCsnpIntervalL1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceHelloPadding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisInterfaceHelloIntervalL1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAuthKeychainDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisSpfIntervalExpL2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisSpfIntervalExpL1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAuthPasswordDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["status"], _ = expandRouterIsisRedistributeStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["protocol"], _ = expandRouterIsisRedistributeProtocol(d, i["protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["level"], _ = expandRouterIsisRedistributeLevel(d, i["level"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["metric"], _ = expandRouterIsisRedistributeMetric(d, i["metric"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["routemap"], _ = expandRouterIsisRedistributeRoutemap(d, i["routemap"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["metric-type"], _ = expandRouterIsisRedistributeMetricType(d, i["metric_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterIsisRedistributeStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeMetricType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRouterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAuthSendonlyArea(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisDefaultInformationLevel6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["status"], _ = expandRouterIsisRedistribute6Status(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["metric"], _ = expandRouterIsisRedistribute6Metric(d, i["metric"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["routemap"], _ = expandRouterIsisRedistribute6Routemap(d, i["routemap"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["protocol"], _ = expandRouterIsisRedistribute6Protocol(d, i["protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["level"], _ = expandRouterIsisRedistribute6Level(d, i["level"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterIsisRedistribute6Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6Metric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6Routemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6Protocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6Level(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisLogNeighbourChanges(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterIsis(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_information_metric6"); ok {
		if setArgNil {
			obj["default-information-metric6"] = nil
		} else {

			t, err := expandRouterIsisDefaultInformationMetric6(d, v, "default_information_metric6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-information-metric6"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_sendonly_domain"); ok {
		if setArgNil {
			obj["auth-sendonly-domain"] = nil
		} else {

			t, err := expandRouterIsisAuthSendonlyDomain(d, v, "auth_sendonly_domain", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-sendonly-domain"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_information_originate6"); ok {
		if setArgNil {
			obj["default-information-originate6"] = nil
		} else {

			t, err := expandRouterIsisDefaultInformationOriginate6(d, v, "default_information_originate6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-information-originate6"] = t
			}
		}
	}

	if v, ok := d.GetOk("summary_address"); ok || d.HasChange("summary_address") {
		if setArgNil {
			obj["summary-address"] = make([]struct{}, 0)
		} else {

			t, err := expandRouterIsisSummaryAddress(d, v, "summary_address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["summary-address"] = t
			}
		}
	}

	if v, ok := d.GetOk("metric_style"); ok {
		if setArgNil {
			obj["metric-style"] = nil
		} else {

			t, err := expandRouterIsisMetricStyle(d, v, "metric_style", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["metric-style"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute6_l1"); ok {
		if setArgNil {
			obj["redistribute6-l1"] = nil
		} else {

			t, err := expandRouterIsisRedistribute6L1(d, v, "redistribute6_l1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute6-l1"] = t
			}
		}
	}

	if v, ok := d.GetOk("lsp_refresh_interval"); ok {
		if setArgNil {
			obj["lsp-refresh-interval"] = nil
		} else {

			t, err := expandRouterIsisLspRefreshInterval(d, v, "lsp_refresh_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["lsp-refresh-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_lsp_lifetime"); ok {
		if setArgNil {
			obj["max-lsp-lifetime"] = nil
		} else {

			t, err := expandRouterIsisMaxLspLifetime(d, v, "max_lsp_lifetime", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-lsp-lifetime"] = t
			}
		}
	}

	if v, ok := d.GetOk("ignore_attached_bit"); ok {
		if setArgNil {
			obj["ignore-attached-bit"] = nil
		} else {

			t, err := expandRouterIsisIgnoreAttachedBit(d, v, "ignore_attached_bit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ignore-attached-bit"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute6_l1_list"); ok {
		if setArgNil {
			obj["redistribute6-l1-list"] = nil
		} else {

			t, err := expandRouterIsisRedistribute6L1List(d, v, "redistribute6_l1_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute6-l1-list"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_keychain_area"); ok {
		if setArgNil {
			obj["auth-keychain-area"] = nil
		} else {

			t, err := expandRouterIsisAuthKeychainArea(d, v, "auth_keychain_area", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-keychain-area"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_mode_area"); ok {
		if setArgNil {
			obj["auth-mode-area"] = nil
		} else {

			t, err := expandRouterIsisAuthModeArea(d, v, "auth_mode_area", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-mode-area"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_information_level"); ok {
		if setArgNil {
			obj["default-information-level"] = nil
		} else {

			t, err := expandRouterIsisDefaultInformationLevel(d, v, "default_information_level", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-information-level"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute_l1"); ok {
		if setArgNil {
			obj["redistribute-l1"] = nil
		} else {

			t, err := expandRouterIsisRedistributeL1(d, v, "redistribute_l1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute-l1"] = t
			}
		}
	}

	if v, ok := d.GetOk("net"); ok || d.HasChange("net") {
		if setArgNil {
			obj["net"] = make([]struct{}, 0)
		} else {

			t, err := expandRouterIsisNet(d, v, "net", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["net"] = t
			}
		}
	}

	if v, ok := d.GetOk("summary_address6"); ok || d.HasChange("summary_address6") {
		if setArgNil {
			obj["summary-address6"] = make([]struct{}, 0)
		} else {

			t, err := expandRouterIsisSummaryAddress6(d, v, "summary_address6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["summary-address6"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_mode_domain"); ok {
		if setArgNil {
			obj["auth-mode-domain"] = nil
		} else {

			t, err := expandRouterIsisAuthModeDomain(d, v, "auth_mode_domain", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-mode-domain"] = t
			}
		}
	}

	if v, ok := d.GetOk("lsp_gen_interval_l1"); ok {
		if setArgNil {
			obj["lsp-gen-interval-l1"] = nil
		} else {

			t, err := expandRouterIsisLspGenIntervalL1(d, v, "lsp_gen_interval_l1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["lsp-gen-interval-l1"] = t
			}
		}
	}

	if v, ok := d.GetOk("lsp_gen_interval_l2"); ok {
		if setArgNil {
			obj["lsp-gen-interval-l2"] = nil
		} else {

			t, err := expandRouterIsisLspGenIntervalL2(d, v, "lsp_gen_interval_l2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["lsp-gen-interval-l2"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute_l1_list"); ok {
		if setArgNil {
			obj["redistribute-l1-list"] = nil
		} else {

			t, err := expandRouterIsisRedistributeL1List(d, v, "redistribute_l1_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute-l1-list"] = t
			}
		}
	}

	if v, ok := d.GetOk("overload_bit"); ok {
		if setArgNil {
			obj["overload-bit"] = nil
		} else {

			t, err := expandRouterIsisOverloadBit(d, v, "overload_bit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["overload-bit"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_password_area"); ok {
		if setArgNil {
			obj["auth-password-area"] = nil
		} else {

			t, err := expandRouterIsisAuthPasswordArea(d, v, "auth_password_area", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-password-area"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_information_metric"); ok {
		if setArgNil {
			obj["default-information-metric"] = nil
		} else {

			t, err := expandRouterIsisDefaultInformationMetric(d, v, "default_information_metric", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-information-metric"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_information_originate"); ok {
		if setArgNil {
			obj["default-information-originate"] = nil
		} else {

			t, err := expandRouterIsisDefaultInformationOriginate(d, v, "default_information_originate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-information-originate"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		if setArgNil {
			obj["interface"] = make([]struct{}, 0)
		} else {

			t, err := expandRouterIsisInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_keychain_domain"); ok {
		if setArgNil {
			obj["auth-keychain-domain"] = nil
		} else {

			t, err := expandRouterIsisAuthKeychainDomain(d, v, "auth_keychain_domain", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-keychain-domain"] = t
			}
		}
	}

	if v, ok := d.GetOk("spf_interval_exp_l2"); ok {
		if setArgNil {
			obj["spf-interval-exp-l2"] = nil
		} else {

			t, err := expandRouterIsisSpfIntervalExpL2(d, v, "spf_interval_exp_l2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["spf-interval-exp-l2"] = t
			}
		}
	}

	if v, ok := d.GetOk("spf_interval_exp_l1"); ok {
		if setArgNil {
			obj["spf-interval-exp-l1"] = nil
		} else {

			t, err := expandRouterIsisSpfIntervalExpL1(d, v, "spf_interval_exp_l1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["spf-interval-exp-l1"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_password_domain"); ok {
		if setArgNil {
			obj["auth-password-domain"] = nil
		} else {

			t, err := expandRouterIsisAuthPasswordDomain(d, v, "auth_password_domain", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-password-domain"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute"); ok || d.HasChange("redistribute") {
		if setArgNil {
			obj["redistribute"] = make([]struct{}, 0)
		} else {

			t, err := expandRouterIsisRedistribute(d, v, "redistribute", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute"] = t
			}
		}
	}

	if v, ok := d.GetOk("router_id"); ok {
		if setArgNil {
			obj["router-id"] = nil
		} else {

			t, err := expandRouterIsisRouterId(d, v, "router_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["router-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_sendonly_area"); ok {
		if setArgNil {
			obj["auth-sendonly-area"] = nil
		} else {

			t, err := expandRouterIsisAuthSendonlyArea(d, v, "auth_sendonly_area", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-sendonly-area"] = t
			}
		}
	}

	if v, ok := d.GetOk("is_type"); ok {
		if setArgNil {
			obj["is-type"] = nil
		} else {

			t, err := expandRouterIsisIsType(d, v, "is_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["is-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_information_level6"); ok {
		if setArgNil {
			obj["default-information-level6"] = nil
		} else {

			t, err := expandRouterIsisDefaultInformationLevel6(d, v, "default_information_level6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-information-level6"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute6"); ok || d.HasChange("redistribute6") {
		if setArgNil {
			obj["redistribute6"] = make([]struct{}, 0)
		} else {

			t, err := expandRouterIsisRedistribute6(d, v, "redistribute6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute6"] = t
			}
		}
	}

	if v, ok := d.GetOk("log_neighbour_changes"); ok {
		if setArgNil {
			obj["log-neighbour-changes"] = nil
		} else {

			t, err := expandRouterIsisLogNeighbourChanges(d, v, "log_neighbour_changes", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-neighbour-changes"] = t
			}
		}
	}

	return &obj, nil
}
