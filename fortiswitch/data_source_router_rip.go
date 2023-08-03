// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: RIP configuration.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceRouterRip() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterRipRead,
		Schema: map[string]*schema.Schema{

			"distance": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distance": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix": &schema.Schema{
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
			"default_metric": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"recv_buffer_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"timeout_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"offset_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"offset": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"access_list": &schema.Schema{
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
			"redistribute": &schema.Schema{
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
						"flags": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
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
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"garbage_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"distance": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_list": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"distance": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"prefix": &schema.Schema{
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
						"default_metric": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"recv_buffer_size": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"timeout_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"offset_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"interface": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"offset": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"access_list": &schema.Schema{
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
						"redistribute": &schema.Schema{
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
									"flags": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"neighbor": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": &schema.Schema{
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
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"garbage_timer": &schema.Schema{
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
						"update_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"split_horizon_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"auth_mode": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"send_version2_broadcast": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"send_version": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"auth_keychain": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"split_horizon": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"flags": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"auth_string": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"receive_version": &schema.Schema{
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
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"listname": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"interface": &schema.Schema{
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
								},
							},
						},
					},
				},
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
			"update_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"split_horizon_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"send_version2_broadcast": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"send_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth_keychain": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"split_horizon": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"flags": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"auth_string": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"receive_version": &schema.Schema{
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
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"listname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": &schema.Schema{
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
					},
				},
			},
		},
	}
}

func dataSourceRouterRipRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "RouterRip"

	o, err := c.ReadRouterRip(mkey)
	if err != nil {
		return fmt.Errorf("Error describing RouterRip: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterRip(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterRip from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterRipDistance(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["access_list"] = dataSourceFlattenRouterRipDistanceAccessList(i["access-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {
			tmp["distance"] = dataSourceFlattenRouterRipDistanceDistance(i["distance"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterRipDistancePrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterRipDistanceId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipDistanceAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistanceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistancePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterRipDistanceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRecvBufferSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipTimeoutTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["status"] = dataSourceFlattenRouterRipOffsetListStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			tmp["direction"] = dataSourceFlattenRouterRipOffsetListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterRipOffsetListInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "offset"
		if _, ok := i["offset"]; ok {
			tmp["offset"] = dataSourceFlattenRouterRipOffsetListOffset(i["offset"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := i["access-list"]; ok {
			tmp["access_list"] = dataSourceFlattenRouterRipOffsetListAccessList(i["access-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterRipOffsetListId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipOffsetListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListOffset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["status"] = dataSourceFlattenRouterRipRedistributeStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			tmp["metric"] = dataSourceFlattenRouterRipRedistributeMetric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			tmp["routemap"] = dataSourceFlattenRouterRipRedistributeRoutemap(i["routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {
			tmp["flags"] = dataSourceFlattenRouterRipRedistributeFlags(i["flags"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterRipRedistributeName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRedistributeFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenRouterRipNeighborIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterRipNeighborId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipNeighborId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipGarbageTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {
			tmp["distance"] = dataSourceFlattenRouterRipVrfDistance(i["distance"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_metric"
		if _, ok := i["default-metric"]; ok {
			tmp["default_metric"] = dataSourceFlattenRouterRipVrfDefaultMetric(i["default-metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recv_buffer_size"
		if _, ok := i["recv-buffer-size"]; ok {
			tmp["recv_buffer_size"] = dataSourceFlattenRouterRipVrfRecvBufferSize(i["recv-buffer-size"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timeout_timer"
		if _, ok := i["timeout-timer"]; ok {
			tmp["timeout_timer"] = dataSourceFlattenRouterRipVrfTimeoutTimer(i["timeout-timer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterRipVrfName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "offset_list"
		if _, ok := i["offset-list"]; ok {
			tmp["offset_list"] = dataSourceFlattenRouterRipVrfOffsetList(i["offset-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redistribute"
		if _, ok := i["redistribute"]; ok {
			tmp["redistribute"] = dataSourceFlattenRouterRipVrfRedistribute(i["redistribute"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor"
		if _, ok := i["neighbor"]; ok {
			tmp["neighbor"] = dataSourceFlattenRouterRipVrfNeighbor(i["neighbor"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			tmp["version"] = dataSourceFlattenRouterRipVrfVersion(i["version"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "garbage_timer"
		if _, ok := i["garbage-timer"]; ok {
			tmp["garbage_timer"] = dataSourceFlattenRouterRipVrfGarbageTimer(i["garbage-timer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_information_originate"
		if _, ok := i["default-information-originate"]; ok {
			tmp["default_information_originate"] = dataSourceFlattenRouterRipVrfDefaultInformationOriginate(i["default-information-originate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive_interface"
		if _, ok := i["passive-interface"]; ok {
			tmp["passive_interface"] = dataSourceFlattenRouterRipVrfPassiveInterface(i["passive-interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_timer"
		if _, ok := i["update-timer"]; ok {
			tmp["update_timer"] = dataSourceFlattenRouterRipVrfUpdateTimer(i["update-timer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterRipVrfInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list"
		if _, ok := i["distribute-list"]; ok {
			tmp["distribute_list"] = dataSourceFlattenRouterRipVrfDistributeList(i["distribute-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network"
		if _, ok := i["network"]; ok {
			tmp["network"] = dataSourceFlattenRouterRipVrfNetwork(i["network"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipVrfDistance(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["access_list"] = dataSourceFlattenRouterRipVrfDistanceAccessList(i["access-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {
			tmp["distance"] = dataSourceFlattenRouterRipVrfDistanceDistance(i["distance"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterRipVrfDistancePrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterRipVrfDistanceId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipVrfDistanceAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfDistanceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfDistancePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterRipVrfDistanceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfDefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfRecvBufferSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfTimeoutTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfOffsetList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["status"] = dataSourceFlattenRouterRipVrfOffsetListStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			tmp["direction"] = dataSourceFlattenRouterRipVrfOffsetListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterRipVrfOffsetListInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "offset"
		if _, ok := i["offset"]; ok {
			tmp["offset"] = dataSourceFlattenRouterRipVrfOffsetListOffset(i["offset"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := i["access-list"]; ok {
			tmp["access_list"] = dataSourceFlattenRouterRipVrfOffsetListAccessList(i["access-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterRipVrfOffsetListId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipVrfOffsetListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfOffsetListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfOffsetListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfOffsetListOffset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfOffsetListAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfOffsetListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["status"] = dataSourceFlattenRouterRipVrfRedistributeStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			tmp["metric"] = dataSourceFlattenRouterRipVrfRedistributeMetric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			tmp["routemap"] = dataSourceFlattenRouterRipVrfRedistributeRoutemap(i["routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {
			tmp["flags"] = dataSourceFlattenRouterRipVrfRedistributeFlags(i["flags"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterRipVrfRedistributeName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipVrfRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfRedistributeFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenRouterRipVrfNeighborIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterRipVrfNeighborId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipVrfNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfNeighborId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfGarbageTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfPassiveInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterRipVrfPassiveInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipVrfPassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfUpdateTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon_status"
		if _, ok := i["split-horizon-status"]; ok {
			tmp["split_horizon_status"] = dataSourceFlattenRouterRipVrfInterfaceSplitHorizonStatus(i["split-horizon-status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := i["auth-mode"]; ok {
			tmp["auth_mode"] = dataSourceFlattenRouterRipVrfInterfaceAuthMode(i["auth-mode"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterRipVrfInterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version2_broadcast"
		if _, ok := i["send-version2-broadcast"]; ok {
			tmp["send_version2_broadcast"] = dataSourceFlattenRouterRipVrfInterfaceSendVersion2Broadcast(i["send-version2-broadcast"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version"
		if _, ok := i["send-version"]; ok {
			tmp["send_version"] = dataSourceFlattenRouterRipVrfInterfaceSendVersion(i["send-version"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain"
		if _, ok := i["auth-keychain"]; ok {
			tmp["auth_keychain"] = dataSourceFlattenRouterRipVrfInterfaceAuthKeychain(i["auth-keychain"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon"
		if _, ok := i["split-horizon"]; ok {
			tmp["split_horizon"] = dataSourceFlattenRouterRipVrfInterfaceSplitHorizon(i["split-horizon"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {
			tmp["flags"] = dataSourceFlattenRouterRipVrfInterfaceFlags(i["flags"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_string"
		if _, ok := i["auth-string"]; ok {
			tmp["auth_string"] = dataSourceFlattenRouterRipVrfInterfaceAuthString(i["auth-string"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "receive_version"
		if _, ok := i["receive-version"]; ok {
			tmp["receive_version"] = dataSourceFlattenRouterRipVrfInterfaceReceiveVersion(i["receive-version"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipVrfInterfaceSplitHorizonStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfInterfaceAuthMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfInterfaceSendVersion2Broadcast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfInterfaceSendVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfInterfaceAuthKeychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfInterfaceSplitHorizon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfInterfaceFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfInterfaceAuthString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfInterfaceReceiveVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfDistributeList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["status"] = dataSourceFlattenRouterRipVrfDistributeListStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listname"
		if _, ok := i["listname"]; ok {
			tmp["listname"] = dataSourceFlattenRouterRipVrfDistributeListListname(i["listname"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			tmp["direction"] = dataSourceFlattenRouterRipVrfDistributeListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterRipVrfDistributeListInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterRipVrfDistributeListId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipVrfDistributeListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfDistributeListListname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfDistributeListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfDistributeListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfDistributeListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVrfNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["prefix"] = dataSourceFlattenRouterRipVrfNetworkPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterRipVrfNetworkId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipVrfNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterRipVrfNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipPassiveInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterRipPassiveInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipPassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipUpdateTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon_status"
		if _, ok := i["split-horizon-status"]; ok {
			tmp["split_horizon_status"] = dataSourceFlattenRouterRipInterfaceSplitHorizonStatus(i["split-horizon-status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := i["auth-mode"]; ok {
			tmp["auth_mode"] = dataSourceFlattenRouterRipInterfaceAuthMode(i["auth-mode"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterRipInterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version2_broadcast"
		if _, ok := i["send-version2-broadcast"]; ok {
			tmp["send_version2_broadcast"] = dataSourceFlattenRouterRipInterfaceSendVersion2Broadcast(i["send-version2-broadcast"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version"
		if _, ok := i["send-version"]; ok {
			tmp["send_version"] = dataSourceFlattenRouterRipInterfaceSendVersion(i["send-version"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain"
		if _, ok := i["auth-keychain"]; ok {
			tmp["auth_keychain"] = dataSourceFlattenRouterRipInterfaceAuthKeychain(i["auth-keychain"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon"
		if _, ok := i["split-horizon"]; ok {
			tmp["split_horizon"] = dataSourceFlattenRouterRipInterfaceSplitHorizon(i["split-horizon"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {
			tmp["flags"] = dataSourceFlattenRouterRipInterfaceFlags(i["flags"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_string"
		if _, ok := i["auth-string"]; ok {
			tmp["auth_string"] = dataSourceFlattenRouterRipInterfaceAuthString(i["auth-string"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "receive_version"
		if _, ok := i["receive-version"]; ok {
			tmp["receive_version"] = dataSourceFlattenRouterRipInterfaceReceiveVersion(i["receive-version"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipInterfaceSplitHorizonStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceAuthMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceSendVersion2Broadcast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceSendVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceAuthKeychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceSplitHorizon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceAuthString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceReceiveVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistributeList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["status"] = dataSourceFlattenRouterRipDistributeListStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listname"
		if _, ok := i["listname"]; ok {
			tmp["listname"] = dataSourceFlattenRouterRipDistributeListListname(i["listname"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			tmp["direction"] = dataSourceFlattenRouterRipDistributeListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterRipDistributeListInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterRipDistributeListId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipDistributeListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistributeListListname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistributeListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistributeListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistributeListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["prefix"] = dataSourceFlattenRouterRipNetworkPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterRipNetworkId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterRipNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterRip(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("distance", dataSourceFlattenRouterRipDistance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("default_metric", dataSourceFlattenRouterRipDefaultMetric(o["default-metric"], d, "default_metric")); err != nil {
		if !fortiAPIPatch(o["default-metric"]) {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if err = d.Set("recv_buffer_size", dataSourceFlattenRouterRipRecvBufferSize(o["recv-buffer-size"], d, "recv_buffer_size")); err != nil {
		if !fortiAPIPatch(o["recv-buffer-size"]) {
			return fmt.Errorf("Error reading recv_buffer_size: %v", err)
		}
	}

	if err = d.Set("timeout_timer", dataSourceFlattenRouterRipTimeoutTimer(o["timeout-timer"], d, "timeout_timer")); err != nil {
		if !fortiAPIPatch(o["timeout-timer"]) {
			return fmt.Errorf("Error reading timeout_timer: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenRouterRipName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("bfd", dataSourceFlattenRouterRipBfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("offset_list", dataSourceFlattenRouterRipOffsetList(o["offset-list"], d, "offset_list")); err != nil {
		if !fortiAPIPatch(o["offset-list"]) {
			return fmt.Errorf("Error reading offset_list: %v", err)
		}
	}

	if err = d.Set("redistribute", dataSourceFlattenRouterRipRedistribute(o["redistribute"], d, "redistribute")); err != nil {
		if !fortiAPIPatch(o["redistribute"]) {
			return fmt.Errorf("Error reading redistribute: %v", err)
		}
	}

	if err = d.Set("neighbor", dataSourceFlattenRouterRipNeighbor(o["neighbor"], d, "neighbor")); err != nil {
		if !fortiAPIPatch(o["neighbor"]) {
			return fmt.Errorf("Error reading neighbor: %v", err)
		}
	}

	if err = d.Set("version", dataSourceFlattenRouterRipVersion(o["version"], d, "version")); err != nil {
		if !fortiAPIPatch(o["version"]) {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	if err = d.Set("garbage_timer", dataSourceFlattenRouterRipGarbageTimer(o["garbage-timer"], d, "garbage_timer")); err != nil {
		if !fortiAPIPatch(o["garbage-timer"]) {
			return fmt.Errorf("Error reading garbage_timer: %v", err)
		}
	}

	if err = d.Set("vrf", dataSourceFlattenRouterRipVrf(o["vrf"], d, "vrf")); err != nil {
		if !fortiAPIPatch(o["vrf"]) {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	if err = d.Set("default_information_originate", dataSourceFlattenRouterRipDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("passive_interface", dataSourceFlattenRouterRipPassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
		if !fortiAPIPatch(o["passive-interface"]) {
			return fmt.Errorf("Error reading passive_interface: %v", err)
		}
	}

	if err = d.Set("update_timer", dataSourceFlattenRouterRipUpdateTimer(o["update-timer"], d, "update_timer")); err != nil {
		if !fortiAPIPatch(o["update-timer"]) {
			return fmt.Errorf("Error reading update_timer: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenRouterRipInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("distribute_list", dataSourceFlattenRouterRipDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
		if !fortiAPIPatch(o["distribute-list"]) {
			return fmt.Errorf("Error reading distribute_list: %v", err)
		}
	}

	if err = d.Set("network", dataSourceFlattenRouterRipNetwork(o["network"], d, "network")); err != nil {
		if !fortiAPIPatch(o["network"]) {
			return fmt.Errorf("Error reading network: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterRipFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
