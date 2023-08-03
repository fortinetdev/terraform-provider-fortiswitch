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

func dataSourceRouterIsis() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterIsisRead,
		Schema: map[string]*schema.Schema{

			"default_information_metric6": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auth_sendonly_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_information_originate6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"summary_address": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"metric_style": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute6_l1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lsp_refresh_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_lsp_lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ignore_attached_bit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute6_l1_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_keychain_area": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_mode_area": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_information_level": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_l1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"net": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"net": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"summary_address6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"auth_mode_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lsp_gen_interval_l1": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lsp_gen_interval_l2": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"redistribute_l1_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"overload_bit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_password_area": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_information_metric": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"default_information_originate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_password_hello": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"priority_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_multiplier_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_multiplier_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"auth_mode_hello": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"passive": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"circuit_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bfd6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"wide_metric_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"wide_metric_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"metric_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"metric_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth_keychain_hello": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hello_interval_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"csnp_interval_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"csnp_interval_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_padding": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hello_interval_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"auth_keychain_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"spf_interval_exp_l2": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"spf_interval_exp_l1": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auth_password_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"metric": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"routemap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_sendonly_area": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_information_level6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"metric": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"routemap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"log_neighbour_changes": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRouterIsisRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "RouterIsis"

	o, err := c.ReadRouterIsis(mkey)
	if err != nil {
		return fmt.Errorf("Error describing RouterIsis: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterIsis(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterIsis from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterIsisDefaultInformationMetric6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthSendonlyDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisDefaultInformationOriginate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisSummaryAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
			tmp["prefix"] = dataSourceFlattenRouterIsisSummaryAddressPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterIsisSummaryAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {
			tmp["level"] = dataSourceFlattenRouterIsisSummaryAddressLevel(i["level"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterIsisSummaryAddressPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterIsisSummaryAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisSummaryAddressLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisMetricStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6L1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisLspRefreshInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisMaxLspLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIgnoreAttachedBit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6L1List(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthKeychainArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthModeArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisDefaultInformationLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisNet(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
			tmp["net"] = dataSourceFlattenRouterIsisNetNet(i["net"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterIsisNetNet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisSummaryAddress6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
			tmp["level"] = dataSourceFlattenRouterIsisSummaryAddress6Level(i["level"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			tmp["prefix6"] = dataSourceFlattenRouterIsisSummaryAddress6Prefix6(i["prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterIsisSummaryAddress6Id(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterIsisSummaryAddress6Level(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisSummaryAddress6Prefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisSummaryAddress6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthModeDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisLspGenIntervalL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisLspGenIntervalL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeL1List(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisOverloadBit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthPasswordArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisDefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
			tmp["auth_password_hello"] = dataSourceFlattenRouterIsisInterfaceAuthPasswordHello(i["auth-password-hello"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_l2"
		if _, ok := i["priority-l2"]; ok {
			tmp["priority_l2"] = dataSourceFlattenRouterIsisInterfacePriorityL2(i["priority-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_l1"
		if _, ok := i["priority-l1"]; ok {
			tmp["priority_l1"] = dataSourceFlattenRouterIsisInterfacePriorityL1(i["priority-l1"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier_l2"
		if _, ok := i["hello-multiplier-l2"]; ok {
			tmp["hello_multiplier_l2"] = dataSourceFlattenRouterIsisInterfaceHelloMultiplierL2(i["hello-multiplier-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier_l1"
		if _, ok := i["hello-multiplier-l1"]; ok {
			tmp["hello_multiplier_l1"] = dataSourceFlattenRouterIsisInterfaceHelloMultiplierL1(i["hello-multiplier-l1"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode_hello"
		if _, ok := i["auth-mode-hello"]; ok {
			tmp["auth_mode_hello"] = dataSourceFlattenRouterIsisInterfaceAuthModeHello(i["auth-mode-hello"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			tmp["bfd"] = dataSourceFlattenRouterIsisInterfaceBfd(i["bfd"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := i["passive"]; ok {
			tmp["passive"] = dataSourceFlattenRouterIsisInterfacePassive(i["passive"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_type"
		if _, ok := i["circuit-type"]; ok {
			tmp["circuit_type"] = dataSourceFlattenRouterIsisInterfaceCircuitType(i["circuit-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd6"
		if _, ok := i["bfd6"]; ok {
			tmp["bfd6"] = dataSourceFlattenRouterIsisInterfaceBfd6(i["bfd6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wide_metric_l1"
		if _, ok := i["wide-metric-l1"]; ok {
			tmp["wide_metric_l1"] = dataSourceFlattenRouterIsisInterfaceWideMetricL1(i["wide-metric-l1"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wide_metric_l2"
		if _, ok := i["wide-metric-l2"]; ok {
			tmp["wide_metric_l2"] = dataSourceFlattenRouterIsisInterfaceWideMetricL2(i["wide-metric-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterIsisInterfaceStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_l1"
		if _, ok := i["metric-l1"]; ok {
			tmp["metric_l1"] = dataSourceFlattenRouterIsisInterfaceMetricL1(i["metric-l1"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_l2"
		if _, ok := i["metric-l2"]; ok {
			tmp["metric_l2"] = dataSourceFlattenRouterIsisInterfaceMetricL2(i["metric-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status6"
		if _, ok := i["status6"]; ok {
			tmp["status6"] = dataSourceFlattenRouterIsisInterfaceStatus6(i["status6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterIsisInterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain_hello"
		if _, ok := i["auth-keychain-hello"]; ok {
			tmp["auth_keychain_hello"] = dataSourceFlattenRouterIsisInterfaceAuthKeychainHello(i["auth-keychain-hello"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval_l2"
		if _, ok := i["hello-interval-l2"]; ok {
			tmp["hello_interval_l2"] = dataSourceFlattenRouterIsisInterfaceHelloIntervalL2(i["hello-interval-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csnp_interval_l2"
		if _, ok := i["csnp-interval-l2"]; ok {
			tmp["csnp_interval_l2"] = dataSourceFlattenRouterIsisInterfaceCsnpIntervalL2(i["csnp-interval-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csnp_interval_l1"
		if _, ok := i["csnp-interval-l1"]; ok {
			tmp["csnp_interval_l1"] = dataSourceFlattenRouterIsisInterfaceCsnpIntervalL1(i["csnp-interval-l1"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_padding"
		if _, ok := i["hello-padding"]; ok {
			tmp["hello_padding"] = dataSourceFlattenRouterIsisInterfaceHelloPadding(i["hello-padding"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval_l1"
		if _, ok := i["hello-interval-l1"]; ok {
			tmp["hello_interval_l1"] = dataSourceFlattenRouterIsisInterfaceHelloIntervalL1(i["hello-interval-l1"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterIsisInterfaceAuthPasswordHello(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfacePriorityL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfacePriorityL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceHelloMultiplierL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceHelloMultiplierL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceAuthModeHello(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfacePassive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceCircuitType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceBfd6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceWideMetricL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceWideMetricL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceMetricL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceMetricL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceStatus6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceAuthKeychainHello(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceHelloIntervalL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceCsnpIntervalL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceCsnpIntervalL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceHelloPadding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisInterfaceHelloIntervalL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthKeychainDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisSpfIntervalExpL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisSpfIntervalExpL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthPasswordDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
			tmp["status"] = dataSourceFlattenRouterIsisRedistributeStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			tmp["protocol"] = dataSourceFlattenRouterIsisRedistributeProtocol(i["protocol"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {
			tmp["level"] = dataSourceFlattenRouterIsisRedistributeLevel(i["level"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			tmp["metric"] = dataSourceFlattenRouterIsisRedistributeMetric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			tmp["routemap"] = dataSourceFlattenRouterIsisRedistributeRoutemap(i["routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := i["metric-type"]; ok {
			tmp["metric_type"] = dataSourceFlattenRouterIsisRedistributeMetricType(i["metric-type"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterIsisRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthSendonlyArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisDefaultInformationLevel6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
			tmp["status"] = dataSourceFlattenRouterIsisRedistribute6Status(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			tmp["metric"] = dataSourceFlattenRouterIsisRedistribute6Metric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			tmp["routemap"] = dataSourceFlattenRouterIsisRedistribute6Routemap(i["routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			tmp["protocol"] = dataSourceFlattenRouterIsisRedistribute6Protocol(i["protocol"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {
			tmp["level"] = dataSourceFlattenRouterIsisRedistribute6Level(i["level"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterIsisRedistribute6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6Metric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6Routemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6Protocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6Level(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisLogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterIsis(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("default_information_metric6", dataSourceFlattenRouterIsisDefaultInformationMetric6(o["default-information-metric6"], d, "default_information_metric6")); err != nil {
		if !fortiAPIPatch(o["default-information-metric6"]) {
			return fmt.Errorf("Error reading default_information_metric6: %v", err)
		}
	}

	if err = d.Set("auth_sendonly_domain", dataSourceFlattenRouterIsisAuthSendonlyDomain(o["auth-sendonly-domain"], d, "auth_sendonly_domain")); err != nil {
		if !fortiAPIPatch(o["auth-sendonly-domain"]) {
			return fmt.Errorf("Error reading auth_sendonly_domain: %v", err)
		}
	}

	if err = d.Set("default_information_originate6", dataSourceFlattenRouterIsisDefaultInformationOriginate6(o["default-information-originate6"], d, "default_information_originate6")); err != nil {
		if !fortiAPIPatch(o["default-information-originate6"]) {
			return fmt.Errorf("Error reading default_information_originate6: %v", err)
		}
	}

	if err = d.Set("summary_address", dataSourceFlattenRouterIsisSummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
		if !fortiAPIPatch(o["summary-address"]) {
			return fmt.Errorf("Error reading summary_address: %v", err)
		}
	}

	if err = d.Set("metric_style", dataSourceFlattenRouterIsisMetricStyle(o["metric-style"], d, "metric_style")); err != nil {
		if !fortiAPIPatch(o["metric-style"]) {
			return fmt.Errorf("Error reading metric_style: %v", err)
		}
	}

	if err = d.Set("redistribute6_l1", dataSourceFlattenRouterIsisRedistribute6L1(o["redistribute6-l1"], d, "redistribute6_l1")); err != nil {
		if !fortiAPIPatch(o["redistribute6-l1"]) {
			return fmt.Errorf("Error reading redistribute6_l1: %v", err)
		}
	}

	if err = d.Set("lsp_refresh_interval", dataSourceFlattenRouterIsisLspRefreshInterval(o["lsp-refresh-interval"], d, "lsp_refresh_interval")); err != nil {
		if !fortiAPIPatch(o["lsp-refresh-interval"]) {
			return fmt.Errorf("Error reading lsp_refresh_interval: %v", err)
		}
	}

	if err = d.Set("max_lsp_lifetime", dataSourceFlattenRouterIsisMaxLspLifetime(o["max-lsp-lifetime"], d, "max_lsp_lifetime")); err != nil {
		if !fortiAPIPatch(o["max-lsp-lifetime"]) {
			return fmt.Errorf("Error reading max_lsp_lifetime: %v", err)
		}
	}

	if err = d.Set("ignore_attached_bit", dataSourceFlattenRouterIsisIgnoreAttachedBit(o["ignore-attached-bit"], d, "ignore_attached_bit")); err != nil {
		if !fortiAPIPatch(o["ignore-attached-bit"]) {
			return fmt.Errorf("Error reading ignore_attached_bit: %v", err)
		}
	}

	if err = d.Set("redistribute6_l1_list", dataSourceFlattenRouterIsisRedistribute6L1List(o["redistribute6-l1-list"], d, "redistribute6_l1_list")); err != nil {
		if !fortiAPIPatch(o["redistribute6-l1-list"]) {
			return fmt.Errorf("Error reading redistribute6_l1_list: %v", err)
		}
	}

	if err = d.Set("auth_keychain_area", dataSourceFlattenRouterIsisAuthKeychainArea(o["auth-keychain-area"], d, "auth_keychain_area")); err != nil {
		if !fortiAPIPatch(o["auth-keychain-area"]) {
			return fmt.Errorf("Error reading auth_keychain_area: %v", err)
		}
	}

	if err = d.Set("auth_mode_area", dataSourceFlattenRouterIsisAuthModeArea(o["auth-mode-area"], d, "auth_mode_area")); err != nil {
		if !fortiAPIPatch(o["auth-mode-area"]) {
			return fmt.Errorf("Error reading auth_mode_area: %v", err)
		}
	}

	if err = d.Set("default_information_level", dataSourceFlattenRouterIsisDefaultInformationLevel(o["default-information-level"], d, "default_information_level")); err != nil {
		if !fortiAPIPatch(o["default-information-level"]) {
			return fmt.Errorf("Error reading default_information_level: %v", err)
		}
	}

	if err = d.Set("redistribute_l1", dataSourceFlattenRouterIsisRedistributeL1(o["redistribute-l1"], d, "redistribute_l1")); err != nil {
		if !fortiAPIPatch(o["redistribute-l1"]) {
			return fmt.Errorf("Error reading redistribute_l1: %v", err)
		}
	}

	if err = d.Set("net", dataSourceFlattenRouterIsisNet(o["net"], d, "net")); err != nil {
		if !fortiAPIPatch(o["net"]) {
			return fmt.Errorf("Error reading net: %v", err)
		}
	}

	if err = d.Set("summary_address6", dataSourceFlattenRouterIsisSummaryAddress6(o["summary-address6"], d, "summary_address6")); err != nil {
		if !fortiAPIPatch(o["summary-address6"]) {
			return fmt.Errorf("Error reading summary_address6: %v", err)
		}
	}

	if err = d.Set("auth_mode_domain", dataSourceFlattenRouterIsisAuthModeDomain(o["auth-mode-domain"], d, "auth_mode_domain")); err != nil {
		if !fortiAPIPatch(o["auth-mode-domain"]) {
			return fmt.Errorf("Error reading auth_mode_domain: %v", err)
		}
	}

	if err = d.Set("lsp_gen_interval_l1", dataSourceFlattenRouterIsisLspGenIntervalL1(o["lsp-gen-interval-l1"], d, "lsp_gen_interval_l1")); err != nil {
		if !fortiAPIPatch(o["lsp-gen-interval-l1"]) {
			return fmt.Errorf("Error reading lsp_gen_interval_l1: %v", err)
		}
	}

	if err = d.Set("lsp_gen_interval_l2", dataSourceFlattenRouterIsisLspGenIntervalL2(o["lsp-gen-interval-l2"], d, "lsp_gen_interval_l2")); err != nil {
		if !fortiAPIPatch(o["lsp-gen-interval-l2"]) {
			return fmt.Errorf("Error reading lsp_gen_interval_l2: %v", err)
		}
	}

	if err = d.Set("redistribute_l1_list", dataSourceFlattenRouterIsisRedistributeL1List(o["redistribute-l1-list"], d, "redistribute_l1_list")); err != nil {
		if !fortiAPIPatch(o["redistribute-l1-list"]) {
			return fmt.Errorf("Error reading redistribute_l1_list: %v", err)
		}
	}

	if err = d.Set("overload_bit", dataSourceFlattenRouterIsisOverloadBit(o["overload-bit"], d, "overload_bit")); err != nil {
		if !fortiAPIPatch(o["overload-bit"]) {
			return fmt.Errorf("Error reading overload_bit: %v", err)
		}
	}

	if err = d.Set("auth_password_area", dataSourceFlattenRouterIsisAuthPasswordArea(o["auth-password-area"], d, "auth_password_area")); err != nil {
		if !fortiAPIPatch(o["auth-password-area"]) {
			return fmt.Errorf("Error reading auth_password_area: %v", err)
		}
	}

	if err = d.Set("default_information_metric", dataSourceFlattenRouterIsisDefaultInformationMetric(o["default-information-metric"], d, "default_information_metric")); err != nil {
		if !fortiAPIPatch(o["default-information-metric"]) {
			return fmt.Errorf("Error reading default_information_metric: %v", err)
		}
	}

	if err = d.Set("default_information_originate", dataSourceFlattenRouterIsisDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenRouterIsisInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("auth_keychain_domain", dataSourceFlattenRouterIsisAuthKeychainDomain(o["auth-keychain-domain"], d, "auth_keychain_domain")); err != nil {
		if !fortiAPIPatch(o["auth-keychain-domain"]) {
			return fmt.Errorf("Error reading auth_keychain_domain: %v", err)
		}
	}

	if err = d.Set("spf_interval_exp_l2", dataSourceFlattenRouterIsisSpfIntervalExpL2(o["spf-interval-exp-l2"], d, "spf_interval_exp_l2")); err != nil {
		if !fortiAPIPatch(o["spf-interval-exp-l2"]) {
			return fmt.Errorf("Error reading spf_interval_exp_l2: %v", err)
		}
	}

	if err = d.Set("spf_interval_exp_l1", dataSourceFlattenRouterIsisSpfIntervalExpL1(o["spf-interval-exp-l1"], d, "spf_interval_exp_l1")); err != nil {
		if !fortiAPIPatch(o["spf-interval-exp-l1"]) {
			return fmt.Errorf("Error reading spf_interval_exp_l1: %v", err)
		}
	}

	if err = d.Set("auth_password_domain", dataSourceFlattenRouterIsisAuthPasswordDomain(o["auth-password-domain"], d, "auth_password_domain")); err != nil {
		if !fortiAPIPatch(o["auth-password-domain"]) {
			return fmt.Errorf("Error reading auth_password_domain: %v", err)
		}
	}

	if err = d.Set("redistribute", dataSourceFlattenRouterIsisRedistribute(o["redistribute"], d, "redistribute")); err != nil {
		if !fortiAPIPatch(o["redistribute"]) {
			return fmt.Errorf("Error reading redistribute: %v", err)
		}
	}

	if err = d.Set("router_id", dataSourceFlattenRouterIsisRouterId(o["router-id"], d, "router_id")); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("auth_sendonly_area", dataSourceFlattenRouterIsisAuthSendonlyArea(o["auth-sendonly-area"], d, "auth_sendonly_area")); err != nil {
		if !fortiAPIPatch(o["auth-sendonly-area"]) {
			return fmt.Errorf("Error reading auth_sendonly_area: %v", err)
		}
	}

	if err = d.Set("is_type", dataSourceFlattenRouterIsisIsType(o["is-type"], d, "is_type")); err != nil {
		if !fortiAPIPatch(o["is-type"]) {
			return fmt.Errorf("Error reading is_type: %v", err)
		}
	}

	if err = d.Set("default_information_level6", dataSourceFlattenRouterIsisDefaultInformationLevel6(o["default-information-level6"], d, "default_information_level6")); err != nil {
		if !fortiAPIPatch(o["default-information-level6"]) {
			return fmt.Errorf("Error reading default_information_level6: %v", err)
		}
	}

	if err = d.Set("redistribute6", dataSourceFlattenRouterIsisRedistribute6(o["redistribute6"], d, "redistribute6")); err != nil {
		if !fortiAPIPatch(o["redistribute6"]) {
			return fmt.Errorf("Error reading redistribute6: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", dataSourceFlattenRouterIsisLogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes")); err != nil {
		if !fortiAPIPatch(o["log-neighbour-changes"]) {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterIsisFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
