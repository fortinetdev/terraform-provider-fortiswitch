// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: OSPF configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceRouterOspf() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterOspfRead,
		Schema: map[string]*schema.Schema{

			"summary_address": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"area": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nssa_translator_role": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dead_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"hello_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"transmit_delay": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"retransmit_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"authentication": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"peer": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"md5_keys": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"key": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"authentication_key": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"shortcut": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"stub_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"range": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"substitute": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"prefix": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"substitute_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"advertise": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"filter_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"list": &schema.Schema{
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
						"default_cost": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"distance_intra_area": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"distance_external": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"network": &schema.Schema{
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
						"area": &schema.Schema{
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
			"default_information_metric_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_overflow_time_to_recover": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"log_neighbour_changes": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rfc1583_compatible": &schema.Schema{
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
			"passive_interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"distribute_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol": &schema.Schema{
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
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"routemap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"metric": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"database_overflow": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_overflow_max_external_lsa": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"spf_timers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"distance_inter_area": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"abr_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"authentication_key": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dead_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_multiplier": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"transmit_delay": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ucast_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"retransmit_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cost": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"md5_keys": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"key": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"mtu_ignore": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"summary_address": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"tag": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"area": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nssa_translator_role": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"virtual_link": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dead_interval": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"hello_interval": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"transmit_delay": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"retransmit_interval": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"authentication": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"peer": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"authentication_key": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"shortcut": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"stub_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"range": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"substitute": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"prefix": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"substitute_status": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"advertise": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"filter_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"direction": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"list": &schema.Schema{
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
									"default_cost": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"distance_intra_area": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"distance_external": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"network": &schema.Schema{
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
									"area": &schema.Schema{
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
						"default_information_metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_overflow_time_to_recover": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"log_neighbour_changes": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"rfc1583_compatible": &schema.Schema{
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
						"passive_interface": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"distribute_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_list": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"protocol": &schema.Schema{
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
						"distance": &schema.Schema{
							Type:     schema.TypeInt,
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
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"metric_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"tag": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"routemap": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"metric": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"database_overflow": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_overflow_max_external_lsa": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"spf_timers": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distance_inter_area": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"abr_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"priority": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"authentication_key": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"dead_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"hello_multiplier": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"transmit_delay": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"ucast_ttl": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"mtu": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"retransmit_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"authentication": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"cost": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"hello_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"ttl": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"md5_keys": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"key": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"mtu_ignore": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterOspfRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "RouterOspf"

	o, err := c.ReadRouterOspf(mkey)
	if err != nil {
		return fmt.Errorf("Error describing RouterOspf: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterOspf(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterOspf from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterOspfSummaryAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["prefix"] = dataSourceFlattenRouterOspfSummaryAddressPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := i["tag"]; ok {
			tmp["tag"] = dataSourceFlattenRouterOspfSummaryAddressTag(i["tag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfSummaryAddressId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfSummaryAddressPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterOspfSummaryAddressTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfSummaryAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfArea(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := i["nssa-translator-role"]; ok {
			tmp["nssa_translator_role"] = dataSourceFlattenRouterOspfAreaNssaTranslatorRole(i["nssa-translator-role"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := i["virtual-link"]; ok {
			tmp["virtual_link"] = dataSourceFlattenRouterOspfAreaVirtualLink(i["virtual-link"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
		if _, ok := i["shortcut"]; ok {
			tmp["shortcut"] = dataSourceFlattenRouterOspfAreaShortcut(i["shortcut"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := i["stub-type"]; ok {
			tmp["stub_type"] = dataSourceFlattenRouterOspfAreaStubType(i["stub-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := i["range"]; ok {
			tmp["range"] = dataSourceFlattenRouterOspfAreaRange(i["range"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list"
		if _, ok := i["filter-list"]; ok {
			tmp["filter_list"] = dataSourceFlattenRouterOspfAreaFilterList(i["filter-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := i["default-cost"]; ok {
			tmp["default_cost"] = dataSourceFlattenRouterOspfAreaDefaultCost(i["default-cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = dataSourceFlattenRouterOspfAreaType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfAreaId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfAreaNssaTranslatorRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLink(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			tmp["dead_interval"] = dataSourceFlattenRouterOspfAreaVirtualLinkDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = dataSourceFlattenRouterOspfAreaVirtualLinkHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterOspfAreaVirtualLinkName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			tmp["transmit_delay"] = dataSourceFlattenRouterOspfAreaVirtualLinkTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			tmp["retransmit_interval"] = dataSourceFlattenRouterOspfAreaVirtualLinkRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = dataSourceFlattenRouterOspfAreaVirtualLinkAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {
			tmp["peer"] = dataSourceFlattenRouterOspfAreaVirtualLinkPeer(i["peer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := i["md5-keys"]; ok {
			tmp["md5_keys"] = dataSourceFlattenRouterOspfAreaVirtualLinkMd5Keys(i["md5-keys"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := i["authentication-key"]; ok {
			tmp["authentication_key"] = dataSourceFlattenRouterOspfAreaVirtualLinkAuthenticationKey(i["authentication-key"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfAreaVirtualLinkDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkMd5Keys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfAreaVirtualLinkMd5KeysId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {
			tmp["key"] = dataSourceFlattenRouterOspfAreaVirtualLinkMd5KeysKey(i["key"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfAreaVirtualLinkMd5KeysId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkMd5KeysKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkAuthenticationKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaShortcut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaStubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute"
		if _, ok := i["substitute"]; ok {
			tmp["substitute"] = dataSourceFlattenRouterOspfAreaRangeSubstitute(i["substitute"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterOspfAreaRangePrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute_status"
		if _, ok := i["substitute-status"]; ok {
			tmp["substitute_status"] = dataSourceFlattenRouterOspfAreaRangeSubstituteStatus(i["substitute-status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfAreaRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := i["advertise"]; ok {
			tmp["advertise"] = dataSourceFlattenRouterOspfAreaRangeAdvertise(i["advertise"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfAreaRangeSubstitute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterOspfAreaRangePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterOspfAreaRangeSubstituteStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaRangeAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaFilterList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			tmp["direction"] = dataSourceFlattenRouterOspfAreaFilterListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := i["list"]; ok {
			tmp["list"] = dataSourceFlattenRouterOspfAreaFilterListList(i["list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfAreaFilterListId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfAreaFilterListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaFilterListList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaFilterListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaDefaultCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistanceIntraArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistanceExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["prefix"] = dataSourceFlattenRouterOspfNetworkPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfNetworkId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area"
		if _, ok := i["area"]; ok {
			tmp["area"] = dataSourceFlattenRouterOspfNetworkArea(i["area"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterOspfNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfNetworkArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDefaultInformationMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDatabaseOverflowTimeToRecover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfLogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRfc1583Compatible(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfPassiveInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterOspfPassiveInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfPassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistributeList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := i["access-list"]; ok {
			tmp["access_list"] = dataSourceFlattenRouterOspfDistributeListAccessList(i["access-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			tmp["protocol"] = dataSourceFlattenRouterOspfDistributeListProtocol(i["protocol"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfDistributeListId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfDistributeListAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistributeListProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistributeListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["status"] = dataSourceFlattenRouterOspfRedistributeStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterOspfRedistributeName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := i["metric-type"]; ok {
			tmp["metric_type"] = dataSourceFlattenRouterOspfRedistributeMetricType(i["metric-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := i["tag"]; ok {
			tmp["tag"] = dataSourceFlattenRouterOspfRedistributeTag(i["tag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			tmp["routemap"] = dataSourceFlattenRouterOspfRedistributeRoutemap(i["routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			tmp["metric"] = dataSourceFlattenRouterOspfRedistributeMetric(i["metric"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDatabaseOverflow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDatabaseOverflowMaxExternalLsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfSpfTimers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistanceInterArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAbrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = dataSourceFlattenRouterOspfInterfacePriority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := i["authentication-key"]; ok {
			tmp["authentication_key"] = dataSourceFlattenRouterOspfInterfaceAuthenticationKey(i["authentication-key"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterOspfInterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			tmp["dead_interval"] = dataSourceFlattenRouterOspfInterfaceDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier"
		if _, ok := i["hello-multiplier"]; ok {
			tmp["hello_multiplier"] = dataSourceFlattenRouterOspfInterfaceHelloMultiplier(i["hello-multiplier"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			tmp["bfd"] = dataSourceFlattenRouterOspfInterfaceBfd(i["bfd"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			tmp["transmit_delay"] = dataSourceFlattenRouterOspfInterfaceTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ucast_ttl"
		if _, ok := i["ucast-ttl"]; ok {
			tmp["ucast_ttl"] = dataSourceFlattenRouterOspfInterfaceUcastTtl(i["ucast-ttl"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := i["mtu"]; ok {
			tmp["mtu"] = dataSourceFlattenRouterOspfInterfaceMtu(i["mtu"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			tmp["retransmit_interval"] = dataSourceFlattenRouterOspfInterfaceRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = dataSourceFlattenRouterOspfInterfaceAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			tmp["cost"] = dataSourceFlattenRouterOspfInterfaceCost(i["cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = dataSourceFlattenRouterOspfInterfaceHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttl"
		if _, ok := i["ttl"]; ok {
			tmp["ttl"] = dataSourceFlattenRouterOspfInterfaceTtl(i["ttl"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := i["md5-keys"]; ok {
			tmp["md5_keys"] = dataSourceFlattenRouterOspfInterfaceMd5Keys(i["md5-keys"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := i["mtu-ignore"]; ok {
			tmp["mtu_ignore"] = dataSourceFlattenRouterOspfInterfaceMtuIgnore(i["mtu-ignore"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfInterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceAuthenticationKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceHelloMultiplier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceUcastTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceMd5Keys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfInterfaceMd5KeysId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {
			tmp["key"] = dataSourceFlattenRouterOspfInterfaceMd5KeysKey(i["key"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfInterfaceMd5KeysId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceMd5KeysKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfInterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_address"
		if _, ok := i["summary-address"]; ok {
			tmp["summary_address"] = dataSourceFlattenRouterOspfVrfSummaryAddress(i["summary-address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area"
		if _, ok := i["area"]; ok {
			tmp["area"] = dataSourceFlattenRouterOspfVrfArea(i["area"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance_intra_area"
		if _, ok := i["distance-intra-area"]; ok {
			tmp["distance_intra_area"] = dataSourceFlattenRouterOspfVrfDistanceIntraArea(i["distance-intra-area"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance_external"
		if _, ok := i["distance-external"]; ok {
			tmp["distance_external"] = dataSourceFlattenRouterOspfVrfDistanceExternal(i["distance-external"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network"
		if _, ok := i["network"]; ok {
			tmp["network"] = dataSourceFlattenRouterOspfVrfNetwork(i["network"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "router_id"
		if _, ok := i["router-id"]; ok {
			tmp["router_id"] = dataSourceFlattenRouterOspfVrfRouterId(i["router-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_information_metric_type"
		if _, ok := i["default-information-metric-type"]; ok {
			tmp["default_information_metric_type"] = dataSourceFlattenRouterOspfVrfDefaultInformationMetricType(i["default-information-metric-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "database_overflow_time_to_recover"
		if _, ok := i["database-overflow-time-to-recover"]; ok {
			tmp["database_overflow_time_to_recover"] = dataSourceFlattenRouterOspfVrfDatabaseOverflowTimeToRecover(i["database-overflow-time-to-recover"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_neighbour_changes"
		if _, ok := i["log-neighbour-changes"]; ok {
			tmp["log_neighbour_changes"] = dataSourceFlattenRouterOspfVrfLogNeighbourChanges(i["log-neighbour-changes"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rfc1583_compatible"
		if _, ok := i["rfc1583-compatible"]; ok {
			tmp["rfc1583_compatible"] = dataSourceFlattenRouterOspfVrfRfc1583Compatible(i["rfc1583-compatible"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_information_metric"
		if _, ok := i["default-information-metric"]; ok {
			tmp["default_information_metric"] = dataSourceFlattenRouterOspfVrfDefaultInformationMetric(i["default-information-metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_information_originate"
		if _, ok := i["default-information-originate"]; ok {
			tmp["default_information_originate"] = dataSourceFlattenRouterOspfVrfDefaultInformationOriginate(i["default-information-originate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive_interface"
		if _, ok := i["passive-interface"]; ok {
			tmp["passive_interface"] = dataSourceFlattenRouterOspfVrfPassiveInterface(i["passive-interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list"
		if _, ok := i["distribute-list"]; ok {
			tmp["distribute_list"] = dataSourceFlattenRouterOspfVrfDistributeList(i["distribute-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {
			tmp["distance"] = dataSourceFlattenRouterOspfVrfDistance(i["distance"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redistribute"
		if _, ok := i["redistribute"]; ok {
			tmp["redistribute"] = dataSourceFlattenRouterOspfVrfRedistribute(i["redistribute"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "database_overflow"
		if _, ok := i["database-overflow"]; ok {
			tmp["database_overflow"] = dataSourceFlattenRouterOspfVrfDatabaseOverflow(i["database-overflow"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "database_overflow_max_external_lsa"
		if _, ok := i["database-overflow-max-external-lsa"]; ok {
			tmp["database_overflow_max_external_lsa"] = dataSourceFlattenRouterOspfVrfDatabaseOverflowMaxExternalLsa(i["database-overflow-max-external-lsa"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterOspfVrfName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spf_timers"
		if _, ok := i["spf-timers"]; ok {
			tmp["spf_timers"] = dataSourceFlattenRouterOspfVrfSpfTimers(i["spf-timers"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance_inter_area"
		if _, ok := i["distance-inter-area"]; ok {
			tmp["distance_inter_area"] = dataSourceFlattenRouterOspfVrfDistanceInterArea(i["distance-inter-area"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "abr_type"
		if _, ok := i["abr-type"]; ok {
			tmp["abr_type"] = dataSourceFlattenRouterOspfVrfAbrType(i["abr-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterOspfVrfInterface(i["interface"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfVrfSummaryAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["prefix"] = dataSourceFlattenRouterOspfVrfSummaryAddressPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := i["tag"]; ok {
			tmp["tag"] = dataSourceFlattenRouterOspfVrfSummaryAddressTag(i["tag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfVrfSummaryAddressId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfVrfSummaryAddressPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterOspfVrfSummaryAddressTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfSummaryAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfArea(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := i["nssa-translator-role"]; ok {
			tmp["nssa_translator_role"] = dataSourceFlattenRouterOspfVrfAreaNssaTranslatorRole(i["nssa-translator-role"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := i["virtual-link"]; ok {
			tmp["virtual_link"] = dataSourceFlattenRouterOspfVrfAreaVirtualLink(i["virtual-link"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
		if _, ok := i["shortcut"]; ok {
			tmp["shortcut"] = dataSourceFlattenRouterOspfVrfAreaShortcut(i["shortcut"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := i["stub-type"]; ok {
			tmp["stub_type"] = dataSourceFlattenRouterOspfVrfAreaStubType(i["stub-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := i["range"]; ok {
			tmp["range"] = dataSourceFlattenRouterOspfVrfAreaRange(i["range"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list"
		if _, ok := i["filter-list"]; ok {
			tmp["filter_list"] = dataSourceFlattenRouterOspfVrfAreaFilterList(i["filter-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := i["default-cost"]; ok {
			tmp["default_cost"] = dataSourceFlattenRouterOspfVrfAreaDefaultCost(i["default-cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = dataSourceFlattenRouterOspfVrfAreaType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfVrfAreaId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfVrfAreaNssaTranslatorRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaVirtualLink(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			tmp["dead_interval"] = dataSourceFlattenRouterOspfVrfAreaVirtualLinkDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = dataSourceFlattenRouterOspfVrfAreaVirtualLinkHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterOspfVrfAreaVirtualLinkName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			tmp["transmit_delay"] = dataSourceFlattenRouterOspfVrfAreaVirtualLinkTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			tmp["retransmit_interval"] = dataSourceFlattenRouterOspfVrfAreaVirtualLinkRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = dataSourceFlattenRouterOspfVrfAreaVirtualLinkAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {
			tmp["peer"] = dataSourceFlattenRouterOspfVrfAreaVirtualLinkPeer(i["peer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := i["authentication-key"]; ok {
			tmp["authentication_key"] = dataSourceFlattenRouterOspfVrfAreaVirtualLinkAuthenticationKey(i["authentication-key"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfVrfAreaVirtualLinkDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaVirtualLinkHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaVirtualLinkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaVirtualLinkTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaVirtualLinkRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaVirtualLinkAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaVirtualLinkPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaVirtualLinkAuthenticationKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaShortcut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaStubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute"
		if _, ok := i["substitute"]; ok {
			tmp["substitute"] = dataSourceFlattenRouterOspfVrfAreaRangeSubstitute(i["substitute"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterOspfVrfAreaRangePrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute_status"
		if _, ok := i["substitute-status"]; ok {
			tmp["substitute_status"] = dataSourceFlattenRouterOspfVrfAreaRangeSubstituteStatus(i["substitute-status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfVrfAreaRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := i["advertise"]; ok {
			tmp["advertise"] = dataSourceFlattenRouterOspfVrfAreaRangeAdvertise(i["advertise"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfVrfAreaRangeSubstitute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterOspfVrfAreaRangePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterOspfVrfAreaRangeSubstituteStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaRangeAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaFilterList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			tmp["direction"] = dataSourceFlattenRouterOspfVrfAreaFilterListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := i["list"]; ok {
			tmp["list"] = dataSourceFlattenRouterOspfVrfAreaFilterListList(i["list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfVrfAreaFilterListId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfVrfAreaFilterListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaFilterListList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaFilterListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaDefaultCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfDistanceIntraArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfDistanceExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["prefix"] = dataSourceFlattenRouterOspfVrfNetworkPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfVrfNetworkId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area"
		if _, ok := i["area"]; ok {
			tmp["area"] = dataSourceFlattenRouterOspfVrfNetworkArea(i["area"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfVrfNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterOspfVrfNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfNetworkArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfRouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfDefaultInformationMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfDatabaseOverflowTimeToRecover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfLogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfRfc1583Compatible(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfDefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfPassiveInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterOspfVrfPassiveInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfVrfPassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfDistributeList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := i["access-list"]; ok {
			tmp["access_list"] = dataSourceFlattenRouterOspfVrfDistributeListAccessList(i["access-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			tmp["protocol"] = dataSourceFlattenRouterOspfVrfDistributeListProtocol(i["protocol"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfVrfDistributeListId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfVrfDistributeListAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfDistributeListProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfDistributeListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["status"] = dataSourceFlattenRouterOspfVrfRedistributeStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterOspfVrfRedistributeName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := i["metric-type"]; ok {
			tmp["metric_type"] = dataSourceFlattenRouterOspfVrfRedistributeMetricType(i["metric-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := i["tag"]; ok {
			tmp["tag"] = dataSourceFlattenRouterOspfVrfRedistributeTag(i["tag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			tmp["routemap"] = dataSourceFlattenRouterOspfVrfRedistributeRoutemap(i["routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			tmp["metric"] = dataSourceFlattenRouterOspfVrfRedistributeMetric(i["metric"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfVrfRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfRedistributeMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfRedistributeTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfDatabaseOverflow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfDatabaseOverflowMaxExternalLsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfSpfTimers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfDistanceInterArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfAbrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = dataSourceFlattenRouterOspfVrfInterfacePriority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := i["authentication-key"]; ok {
			tmp["authentication_key"] = dataSourceFlattenRouterOspfVrfInterfaceAuthenticationKey(i["authentication-key"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterOspfVrfInterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			tmp["dead_interval"] = dataSourceFlattenRouterOspfVrfInterfaceDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier"
		if _, ok := i["hello-multiplier"]; ok {
			tmp["hello_multiplier"] = dataSourceFlattenRouterOspfVrfInterfaceHelloMultiplier(i["hello-multiplier"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			tmp["transmit_delay"] = dataSourceFlattenRouterOspfVrfInterfaceTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ucast_ttl"
		if _, ok := i["ucast-ttl"]; ok {
			tmp["ucast_ttl"] = dataSourceFlattenRouterOspfVrfInterfaceUcastTtl(i["ucast-ttl"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := i["mtu"]; ok {
			tmp["mtu"] = dataSourceFlattenRouterOspfVrfInterfaceMtu(i["mtu"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			tmp["retransmit_interval"] = dataSourceFlattenRouterOspfVrfInterfaceRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = dataSourceFlattenRouterOspfVrfInterfaceAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			tmp["cost"] = dataSourceFlattenRouterOspfVrfInterfaceCost(i["cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = dataSourceFlattenRouterOspfVrfInterfaceHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttl"
		if _, ok := i["ttl"]; ok {
			tmp["ttl"] = dataSourceFlattenRouterOspfVrfInterfaceTtl(i["ttl"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := i["md5-keys"]; ok {
			tmp["md5_keys"] = dataSourceFlattenRouterOspfVrfInterfaceMd5Keys(i["md5-keys"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := i["mtu-ignore"]; ok {
			tmp["mtu_ignore"] = dataSourceFlattenRouterOspfVrfInterfaceMtuIgnore(i["mtu-ignore"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfVrfInterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceAuthenticationKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceHelloMultiplier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceUcastTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceMd5Keys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterOspfVrfInterfaceMd5KeysId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {
			tmp["key"] = dataSourceFlattenRouterOspfVrfInterfaceMd5KeysKey(i["key"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfVrfInterfaceMd5KeysId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceMd5KeysKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfVrfInterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterOspf(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("summary_address", dataSourceFlattenRouterOspfSummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
		if !fortiAPIPatch(o["summary-address"]) {
			return fmt.Errorf("Error reading summary_address: %v", err)
		}
	}

	if err = d.Set("area", dataSourceFlattenRouterOspfArea(o["area"], d, "area")); err != nil {
		if !fortiAPIPatch(o["area"]) {
			return fmt.Errorf("Error reading area: %v", err)
		}
	}

	if err = d.Set("distance_intra_area", dataSourceFlattenRouterOspfDistanceIntraArea(o["distance-intra-area"], d, "distance_intra_area")); err != nil {
		if !fortiAPIPatch(o["distance-intra-area"]) {
			return fmt.Errorf("Error reading distance_intra_area: %v", err)
		}
	}

	if err = d.Set("distance_external", dataSourceFlattenRouterOspfDistanceExternal(o["distance-external"], d, "distance_external")); err != nil {
		if !fortiAPIPatch(o["distance-external"]) {
			return fmt.Errorf("Error reading distance_external: %v", err)
		}
	}

	if err = d.Set("network", dataSourceFlattenRouterOspfNetwork(o["network"], d, "network")); err != nil {
		if !fortiAPIPatch(o["network"]) {
			return fmt.Errorf("Error reading network: %v", err)
		}
	}

	if err = d.Set("router_id", dataSourceFlattenRouterOspfRouterId(o["router-id"], d, "router_id")); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("default_information_metric_type", dataSourceFlattenRouterOspfDefaultInformationMetricType(o["default-information-metric-type"], d, "default_information_metric_type")); err != nil {
		if !fortiAPIPatch(o["default-information-metric-type"]) {
			return fmt.Errorf("Error reading default_information_metric_type: %v", err)
		}
	}

	if err = d.Set("database_overflow_time_to_recover", dataSourceFlattenRouterOspfDatabaseOverflowTimeToRecover(o["database-overflow-time-to-recover"], d, "database_overflow_time_to_recover")); err != nil {
		if !fortiAPIPatch(o["database-overflow-time-to-recover"]) {
			return fmt.Errorf("Error reading database_overflow_time_to_recover: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", dataSourceFlattenRouterOspfLogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes")); err != nil {
		if !fortiAPIPatch(o["log-neighbour-changes"]) {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	if err = d.Set("rfc1583_compatible", dataSourceFlattenRouterOspfRfc1583Compatible(o["rfc1583-compatible"], d, "rfc1583_compatible")); err != nil {
		if !fortiAPIPatch(o["rfc1583-compatible"]) {
			return fmt.Errorf("Error reading rfc1583_compatible: %v", err)
		}
	}

	if err = d.Set("default_information_metric", dataSourceFlattenRouterOspfDefaultInformationMetric(o["default-information-metric"], d, "default_information_metric")); err != nil {
		if !fortiAPIPatch(o["default-information-metric"]) {
			return fmt.Errorf("Error reading default_information_metric: %v", err)
		}
	}

	if err = d.Set("default_information_originate", dataSourceFlattenRouterOspfDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("passive_interface", dataSourceFlattenRouterOspfPassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
		if !fortiAPIPatch(o["passive-interface"]) {
			return fmt.Errorf("Error reading passive_interface: %v", err)
		}
	}

	if err = d.Set("distribute_list", dataSourceFlattenRouterOspfDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
		if !fortiAPIPatch(o["distribute-list"]) {
			return fmt.Errorf("Error reading distribute_list: %v", err)
		}
	}

	if err = d.Set("distance", dataSourceFlattenRouterOspfDistance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("redistribute", dataSourceFlattenRouterOspfRedistribute(o["redistribute"], d, "redistribute")); err != nil {
		if !fortiAPIPatch(o["redistribute"]) {
			return fmt.Errorf("Error reading redistribute: %v", err)
		}
	}

	if err = d.Set("database_overflow", dataSourceFlattenRouterOspfDatabaseOverflow(o["database-overflow"], d, "database_overflow")); err != nil {
		if !fortiAPIPatch(o["database-overflow"]) {
			return fmt.Errorf("Error reading database_overflow: %v", err)
		}
	}

	if err = d.Set("database_overflow_max_external_lsa", dataSourceFlattenRouterOspfDatabaseOverflowMaxExternalLsa(o["database-overflow-max-external-lsa"], d, "database_overflow_max_external_lsa")); err != nil {
		if !fortiAPIPatch(o["database-overflow-max-external-lsa"]) {
			return fmt.Errorf("Error reading database_overflow_max_external_lsa: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenRouterOspfName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("spf_timers", dataSourceFlattenRouterOspfSpfTimers(o["spf-timers"], d, "spf_timers")); err != nil {
		if !fortiAPIPatch(o["spf-timers"]) {
			return fmt.Errorf("Error reading spf_timers: %v", err)
		}
	}

	if err = d.Set("distance_inter_area", dataSourceFlattenRouterOspfDistanceInterArea(o["distance-inter-area"], d, "distance_inter_area")); err != nil {
		if !fortiAPIPatch(o["distance-inter-area"]) {
			return fmt.Errorf("Error reading distance_inter_area: %v", err)
		}
	}

	if err = d.Set("abr_type", dataSourceFlattenRouterOspfAbrType(o["abr-type"], d, "abr_type")); err != nil {
		if !fortiAPIPatch(o["abr-type"]) {
			return fmt.Errorf("Error reading abr_type: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenRouterOspfInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vrf", dataSourceFlattenRouterOspfVrf(o["vrf"], d, "vrf")); err != nil {
		if !fortiAPIPatch(o["vrf"]) {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterOspfFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
