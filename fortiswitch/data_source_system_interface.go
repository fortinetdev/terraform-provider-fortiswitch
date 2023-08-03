// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure interfaces.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemInterface() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemInterfaceRead,
		Schema: map[string]*schema.Schema{

			"defaultgw": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gwdetect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrrp_virtual_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bfd_detect_mult": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bfd_required_min_rx": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"l2_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"src_check_allow_default": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_relay_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"forward_domain": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"speed": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlanforward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"macaddr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bfd_desired_min_tx": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"switch": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlanid": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cli_conn_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"detectserver": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrrp": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"adv_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"start_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"backup_vmac_fwd": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vrid": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vrgrp": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"preempt": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vrdst": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vrip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_client_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ping_serv_status": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"snmp_index": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"icmp_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_relay_option82": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_server_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_expire": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_vendor_specific_option": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_relay_service": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"detectprotocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"src_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6_allowaccess": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_retrans_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vrrp6": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"priority": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"vrip6": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"adv_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"vrid": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"vrgrp": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"preempt": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"accept_mode": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"vrdst6": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"ip6_other_flag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_dns_server_override": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vrip6_link_local": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_prefix_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"valid_life_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"prefix": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"autonomous_flag": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"onlink_flag": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"preferred_life_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"ip6_link_mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip6_manage_flag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_min_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip6_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_hop_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip6_reachable_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip6_extra_addr": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"ip6_default_life": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip6_max_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vrrp_virtual_mac6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_unknown_mcast_to_cpu": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_send_adv": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"autoconf": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dhcp6_information_request": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dynamicgw": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mtu_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dynamic_dns1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_dns2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ha_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"secondaryip": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gwdetect": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"detectprotocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"detectserver": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allowaccess": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ping_serv_status": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ha_priority": &schema.Schema{
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
			"switch_members": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemInterfaceRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemInterface: type error")
	}

	o, err := c.ReadSystemInterface(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemInterface: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemInterface from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemInterfaceDefaultgw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceGwdetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemInterfaceVrrpVirtualMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceBfdDetectMult(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceBfdRequiredMinRx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceL2Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSrcCheckAllowDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpRelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	var temp_v = v.(string)
	temp_v = strings.TrimRight(temp_v, " ")
	temp_v = strings.ReplaceAll(temp_v, "\"", "")
	var rst_v interface{} = temp_v
	return rst_v
}

func dataSourceFlattenSystemInterfaceForwardDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVlanforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMacaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceBfdDesiredMinTx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceCliConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDetectserver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["status"] = dataSourceFlattenSystemInterfaceVrrpStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = dataSourceFlattenSystemInterfaceVrrpPriority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			tmp["adv_interval"] = dataSourceFlattenSystemInterfaceVrrpAdvInterval(i["adv-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			tmp["start_time"] = dataSourceFlattenSystemInterfaceVrrpStartTime(i["start-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "backup_vmac_fwd"
		if _, ok := i["backup-vmac-fwd"]; ok {
			tmp["backup_vmac_fwd"] = dataSourceFlattenSystemInterfaceVrrpBackupVmacFwd(i["backup-vmac-fwd"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			tmp["vrid"] = dataSourceFlattenSystemInterfaceVrrpVrid(i["vrid"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			tmp["vrgrp"] = dataSourceFlattenSystemInterfaceVrrpVrgrp(i["vrgrp"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			tmp["preempt"] = dataSourceFlattenSystemInterfaceVrrpPreempt(i["preempt"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			tmp["version"] = dataSourceFlattenSystemInterfaceVrrpVersion(i["version"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst"
		if _, ok := i["vrdst"]; ok {
			tmp["vrdst"] = dataSourceFlattenSystemInterfaceVrrpVrdst(i["vrdst"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip"
		if _, ok := i["vrip"]; ok {
			tmp["vrip"] = dataSourceFlattenSystemInterfaceVrrpVrip(i["vrip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceVrrpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpAdvInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpStartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpBackupVmacFwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpVrid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpVrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpPreempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpVrdst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpVrip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpClientIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePingServStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSnmpIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIcmpRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpRelayOption82(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceRemoteIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDnsServerOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcp_Expire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpVendorSpecificOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpRelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDetectprotocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSrcCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := i["ip6-allowaccess"]; ok {
		result["ip6_allowaccess"] = dataSourceFlattenSystemInterfaceIpv6Ip6Allowaccess(i["ip6-allowaccess"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_retrans_time"
	if _, ok := i["ip6-retrans-time"]; ok {
		result["ip6_retrans_time"] = dataSourceFlattenSystemInterfaceIpv6Ip6RetransTime(i["ip6-retrans-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp6"
	if _, ok := i["vrrp6"]; ok {
		result["vrrp6"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6(i["vrrp6"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_other_flag"
	if _, ok := i["ip6-other-flag"]; ok {
		result["ip6_other_flag"] = dataSourceFlattenSystemInterfaceIpv6Ip6OtherFlag(i["ip6-other-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_dns_server_override"
	if _, ok := i["ip6-dns-server-override"]; ok {
		result["ip6_dns_server_override"] = dataSourceFlattenSystemInterfaceIpv6Ip6DnsServerOverride(i["ip6-dns-server-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrip6_link_local"
	if _, ok := i["vrip6_link_local"]; ok {
		result["vrip6_link_local"] = dataSourceFlattenSystemInterfaceIpv6Vrip6_Link_Local(i["vrip6_link_local"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_address"
	if _, ok := i["ip6-address"]; ok {
		result["ip6_address"] = dataSourceFlattenSystemInterfaceIpv6Ip6Address(i["ip6-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_prefix_list"
	if _, ok := i["ip6-prefix-list"]; ok {
		result["ip6_prefix_list"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixList(i["ip6-prefix-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_link_mtu"
	if _, ok := i["ip6-link-mtu"]; ok {
		result["ip6_link_mtu"] = dataSourceFlattenSystemInterfaceIpv6Ip6LinkMtu(i["ip6-link-mtu"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_manage_flag"
	if _, ok := i["ip6-manage-flag"]; ok {
		result["ip6_manage_flag"] = dataSourceFlattenSystemInterfaceIpv6Ip6ManageFlag(i["ip6-manage-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_min_interval"
	if _, ok := i["ip6-min-interval"]; ok {
		result["ip6_min_interval"] = dataSourceFlattenSystemInterfaceIpv6Ip6MinInterval(i["ip6-min-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_mode"
	if _, ok := i["ip6-mode"]; ok {
		result["ip6_mode"] = dataSourceFlattenSystemInterfaceIpv6Ip6Mode(i["ip6-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_hop_limit"
	if _, ok := i["ip6-hop-limit"]; ok {
		result["ip6_hop_limit"] = dataSourceFlattenSystemInterfaceIpv6Ip6HopLimit(i["ip6-hop-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_reachable_time"
	if _, ok := i["ip6-reachable-time"]; ok {
		result["ip6_reachable_time"] = dataSourceFlattenSystemInterfaceIpv6Ip6ReachableTime(i["ip6-reachable-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_extra_addr"
	if _, ok := i["ip6-extra-addr"]; ok {
		result["ip6_extra_addr"] = dataSourceFlattenSystemInterfaceIpv6Ip6ExtraAddr(i["ip6-extra-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_default_life"
	if _, ok := i["ip6-default-life"]; ok {
		result["ip6_default_life"] = dataSourceFlattenSystemInterfaceIpv6Ip6DefaultLife(i["ip6-default-life"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_max_interval"
	if _, ok := i["ip6-max-interval"]; ok {
		result["ip6_max_interval"] = dataSourceFlattenSystemInterfaceIpv6Ip6MaxInterval(i["ip6-max-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp_virtual_mac6"
	if _, ok := i["vrrp-virtual-mac6"]; ok {
		result["vrrp_virtual_mac6"] = dataSourceFlattenSystemInterfaceIpv6VrrpVirtualMac6(i["vrrp-virtual-mac6"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_unknown_mcast_to_cpu"
	if _, ok := i["ip6-unknown-mcast-to-cpu"]; ok {
		result["ip6_unknown_mcast_to_cpu"] = dataSourceFlattenSystemInterfaceIpv6Ip6UnknownMcastToCpu(i["ip6-unknown-mcast-to-cpu"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_send_adv"
	if _, ok := i["ip6-send-adv"]; ok {
		result["ip6_send_adv"] = dataSourceFlattenSystemInterfaceIpv6Ip6SendAdv(i["ip6-send-adv"], d, pre_append)
	}

	pre_append = pre + ".0." + "autoconf"
	if _, ok := i["autoconf"]; ok {
		result["autoconf"] = dataSourceFlattenSystemInterfaceIpv6Autoconf(i["autoconf"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_information_request"
	if _, ok := i["dhcp6-information-request"]; ok {
		result["dhcp6_information_request"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6InformationRequest(i["dhcp6-information-request"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemInterfaceIpv6Ip6Allowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6RetransTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["status"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6Status(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6Priority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := i["vrip6"]; ok {
			tmp["vrip6"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrip6(i["vrip6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			tmp["adv_interval"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6AdvInterval(i["adv-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			tmp["start_time"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6StartTime(i["start-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			tmp["vrid"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrid(i["vrid"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			tmp["vrgrp"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrgrp(i["vrgrp"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			tmp["preempt"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6Preempt(i["preempt"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := i["accept-mode"]; ok {
			tmp["accept_mode"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6AcceptMode(i["accept-mode"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := i["vrdst6"]; ok {
			tmp["vrdst6"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrdst6(i["vrdst6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrip6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6AdvInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6StartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6Preempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6AcceptMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrdst6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6OtherFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6DnsServerOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrip6_Link_Local(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := i["valid-life-time"]; ok {
			tmp["valid_life_time"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListValidLifeTime(i["valid-life-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := i["autonomous-flag"]; ok {
			tmp["autonomous_flag"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListAutonomousFlag(i["autonomous-flag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			tmp["onlink_flag"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListOnlinkFlag(i["onlink-flag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := i["preferred-life-time"]; ok {
			tmp["preferred_life_time"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListPreferredLifeTime(i["preferred-life-time"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListValidLifeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListAutonomousFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListOnlinkFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListPreferredLifeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6LinkMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6ManageFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6MinInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6HopLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6ReachableTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6ExtraAddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["prefix"] = dataSourceFlattenSystemInterfaceIpv6Ip6ExtraAddrPrefix(i["prefix"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceIpv6Ip6ExtraAddrPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6DefaultLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6MaxInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6VrrpVirtualMac6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6UnknownMcastToCpu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6SendAdv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Autoconf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6InformationRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDynamicgw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMtuOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDynamic_Dns1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDynamic_Dns2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryip(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gwdetect"
		if _, ok := i["gwdetect"]; ok {
			tmp["gwdetect"] = dataSourceFlattenSystemInterfaceSecondaryipGwdetect(i["gwdetect"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectprotocol"
		if _, ok := i["detectprotocol"]; ok {
			tmp["detectprotocol"] = dataSourceFlattenSystemInterfaceSecondaryipDetectprotocol(i["detectprotocol"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenSystemInterfaceSecondaryipIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectserver"
		if _, ok := i["detectserver"]; ok {
			tmp["detectserver"] = dataSourceFlattenSystemInterfaceSecondaryipDetectserver(i["detectserver"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowaccess"
		if _, ok := i["allowaccess"]; ok {
			tmp["allowaccess"] = dataSourceFlattenSystemInterfaceSecondaryipAllowaccess(i["allowaccess"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ping_serv_status"
		if _, ok := i["ping-serv-status"]; ok {
			tmp["ping_serv_status"] = dataSourceFlattenSystemInterfaceSecondaryipPingServStatus(i["ping-serv-status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {
			tmp["ha_priority"] = dataSourceFlattenSystemInterfaceSecondaryipHaPriority(i["ha-priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenSystemInterfaceSecondaryipId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceSecondaryipGwdetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryipDetectprotocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryipIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryipDetectserver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryipAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryipPingServStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryipHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryipId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchMembers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_name"
		if _, ok := i["member-name"]; ok {
			tmp["member_name"] = dataSourceFlattenSystemInterfaceSwitchMembersMemberName(i["member-name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceSwitchMembersMemberName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("defaultgw", dataSourceFlattenSystemInterfaceDefaultgw(o["defaultgw"], d, "defaultgw")); err != nil {
		if !fortiAPIPatch(o["defaultgw"]) {
			return fmt.Errorf("Error reading defaultgw: %v", err)
		}
	}

	if err = d.Set("gwdetect", dataSourceFlattenSystemInterfaceGwdetect(o["gwdetect"], d, "gwdetect")); err != nil {
		if !fortiAPIPatch(o["gwdetect"]) {
			return fmt.Errorf("Error reading gwdetect: %v", err)
		}
	}

	if err = d.Set("weight", dataSourceFlattenSystemInterfaceWeight(o["weight"], d, "weight")); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("ip", dataSourceFlattenSystemInterfaceIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("vrrp_virtual_mac", dataSourceFlattenSystemInterfaceVrrpVirtualMac(o["vrrp-virtual-mac"], d, "vrrp_virtual_mac")); err != nil {
		if !fortiAPIPatch(o["vrrp-virtual-mac"]) {
			return fmt.Errorf("Error reading vrrp_virtual_mac: %v", err)
		}
	}

	if err = d.Set("bfd_detect_mult", dataSourceFlattenSystemInterfaceBfdDetectMult(o["bfd-detect-mult"], d, "bfd_detect_mult")); err != nil {
		if !fortiAPIPatch(o["bfd-detect-mult"]) {
			return fmt.Errorf("Error reading bfd_detect_mult: %v", err)
		}
	}

	if err = d.Set("bfd_required_min_rx", dataSourceFlattenSystemInterfaceBfdRequiredMinRx(o["bfd-required-min-rx"], d, "bfd_required_min_rx")); err != nil {
		if !fortiAPIPatch(o["bfd-required-min-rx"]) {
			return fmt.Errorf("Error reading bfd_required_min_rx: %v", err)
		}
	}

	if err = d.Set("l2_interface", dataSourceFlattenSystemInterfaceL2Interface(o["l2-interface"], d, "l2_interface")); err != nil {
		if !fortiAPIPatch(o["l2-interface"]) {
			return fmt.Errorf("Error reading l2_interface: %v", err)
		}
	}

	if err = d.Set("src_check_allow_default", dataSourceFlattenSystemInterfaceSrcCheckAllowDefault(o["src-check-allow-default"], d, "src_check_allow_default")); err != nil {
		if !fortiAPIPatch(o["src-check-allow-default"]) {
			return fmt.Errorf("Error reading src_check_allow_default: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_ip", dataSourceFlattenSystemInterfaceDhcpRelayIp(o["dhcp-relay-ip"], d, "dhcp_relay_ip")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-ip"]) {
			return fmt.Errorf("Error reading dhcp_relay_ip: %v", err)
		}
	}

	if err = d.Set("forward_domain", dataSourceFlattenSystemInterfaceForwardDomain(o["forward-domain"], d, "forward_domain")); err != nil {
		if !fortiAPIPatch(o["forward-domain"]) {
			return fmt.Errorf("Error reading forward_domain: %v", err)
		}
	}

	if err = d.Set("speed", dataSourceFlattenSystemInterfaceSpeed(o["speed"], d, "speed")); err != nil {
		if !fortiAPIPatch(o["speed"]) {
			return fmt.Errorf("Error reading speed: %v", err)
		}
	}

	if err = d.Set("vlanforward", dataSourceFlattenSystemInterfaceVlanforward(o["vlanforward"], d, "vlanforward")); err != nil {
		if !fortiAPIPatch(o["vlanforward"]) {
			return fmt.Errorf("Error reading vlanforward: %v", err)
		}
	}

	if err = d.Set("priority", dataSourceFlattenSystemInterfacePriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("bfd", dataSourceFlattenSystemInterfaceBfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("macaddr", dataSourceFlattenSystemInterfaceMacaddr(o["macaddr"], d, "macaddr")); err != nil {
		if !fortiAPIPatch(o["macaddr"]) {
			return fmt.Errorf("Error reading macaddr: %v", err)
		}
	}

	if err = d.Set("bfd_desired_min_tx", dataSourceFlattenSystemInterfaceBfdDesiredMinTx(o["bfd-desired-min-tx"], d, "bfd_desired_min_tx")); err != nil {
		if !fortiAPIPatch(o["bfd-desired-min-tx"]) {
			return fmt.Errorf("Error reading bfd_desired_min_tx: %v", err)
		}
	}

	if err = d.Set("switch", dataSourceFlattenSystemInterfaceSwitch(o["switch"], d, "switch")); err != nil {
		if !fortiAPIPatch(o["switch"]) {
			return fmt.Errorf("Error reading switch: %v", err)
		}
	}

	if err = d.Set("vlanid", dataSourceFlattenSystemInterfaceVlanid(o["vlanid"], d, "vlanid")); err != nil {
		if !fortiAPIPatch(o["vlanid"]) {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	if err = d.Set("cli_conn_status", dataSourceFlattenSystemInterfaceCliConnStatus(o["cli-conn-status"], d, "cli_conn_status")); err != nil {
		if !fortiAPIPatch(o["cli-conn-status"]) {
			return fmt.Errorf("Error reading cli_conn_status: %v", err)
		}
	}

	if err = d.Set("detectserver", dataSourceFlattenSystemInterfaceDetectserver(o["detectserver"], d, "detectserver")); err != nil {
		if !fortiAPIPatch(o["detectserver"]) {
			return fmt.Errorf("Error reading detectserver: %v", err)
		}
	}

	if err = d.Set("vrrp", dataSourceFlattenSystemInterfaceVrrp(o["vrrp"], d, "vrrp")); err != nil {
		if !fortiAPIPatch(o["vrrp"]) {
			return fmt.Errorf("Error reading vrrp: %v", err)
		}
	}

	if err = d.Set("allowaccess", dataSourceFlattenSystemInterfaceAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("dhcp_client_identifier", dataSourceFlattenSystemInterfaceDhcpClientIdentifier(o["dhcp-client-identifier"], d, "dhcp_client_identifier")); err != nil {
		if !fortiAPIPatch(o["dhcp-client-identifier"]) {
			return fmt.Errorf("Error reading dhcp_client_identifier: %v", err)
		}
	}

	if err = d.Set("ping_serv_status", dataSourceFlattenSystemInterfacePingServStatus(o["ping-serv-status"], d, "ping_serv_status")); err != nil {
		if !fortiAPIPatch(o["ping-serv-status"]) {
			return fmt.Errorf("Error reading ping_serv_status: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemInterfaceType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("snmp_index", dataSourceFlattenSystemInterfaceSnmpIndex(o["snmp-index"], d, "snmp_index")); err != nil {
		if !fortiAPIPatch(o["snmp-index"]) {
			return fmt.Errorf("Error reading snmp_index: %v", err)
		}
	}

	if err = d.Set("icmp_redirect", dataSourceFlattenSystemInterfaceIcmpRedirect(o["icmp-redirect"], d, "icmp_redirect")); err != nil {
		if !fortiAPIPatch(o["icmp-redirect"]) {
			return fmt.Errorf("Error reading icmp_redirect: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_option82", dataSourceFlattenSystemInterfaceDhcpRelayOption82(o["dhcp-relay-option82"], d, "dhcp_relay_option82")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-option82"]) {
			return fmt.Errorf("Error reading dhcp_relay_option82: %v", err)
		}
	}

	if err = d.Set("description", dataSourceFlattenSystemInterfaceDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("remote_ip", dataSourceFlattenSystemInterfaceRemoteIp(o["remote-ip"], d, "remote_ip")); err != nil {
		if !fortiAPIPatch(o["remote-ip"]) {
			return fmt.Errorf("Error reading remote_ip: %v", err)
		}
	}

	if err = d.Set("dns_server_override", dataSourceFlattenSystemInterfaceDnsServerOverride(o["dns-server-override"], d, "dns_server_override")); err != nil {
		if !fortiAPIPatch(o["dns-server-override"]) {
			return fmt.Errorf("Error reading dns_server_override: %v", err)
		}
	}

	if err = d.Set("secondary_ip", dataSourceFlattenSystemInterfaceSecondaryIp(o["secondary-IP"], d, "secondary_ip")); err != nil {
		if !fortiAPIPatch(o["secondary-IP"]) {
			return fmt.Errorf("Error reading secondary_ip: %v", err)
		}
	}

	if err = d.Set("vrf", dataSourceFlattenSystemInterfaceVrf(o["vrf"], d, "vrf")); err != nil {
		if !fortiAPIPatch(o["vrf"]) {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	if err = d.Set("dhcp_expire", dataSourceFlattenSystemInterfaceDhcp_Expire(o["dhcp_expire"], d, "dhcp_expire")); err != nil {
		if !fortiAPIPatch(o["dhcp_expire"]) {
			return fmt.Errorf("Error reading dhcp_expire: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemInterfaceInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("dhcp_vendor_specific_option", dataSourceFlattenSystemInterfaceDhcpVendorSpecificOption(o["dhcp-vendor-specific-option"], d, "dhcp_vendor_specific_option")); err != nil {
		if !fortiAPIPatch(o["dhcp-vendor-specific-option"]) {
			return fmt.Errorf("Error reading dhcp_vendor_specific_option: %v", err)
		}
	}

	if err = d.Set("vdom", dataSourceFlattenSystemInterfaceVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_service", dataSourceFlattenSystemInterfaceDhcpRelayService(o["dhcp-relay-service"], d, "dhcp_relay_service")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-service"]) {
			return fmt.Errorf("Error reading dhcp_relay_service: %v", err)
		}
	}

	if err = d.Set("distance", dataSourceFlattenSystemInterfaceDistance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenSystemInterfaceName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("detectprotocol", dataSourceFlattenSystemInterfaceDetectprotocol(o["detectprotocol"], d, "detectprotocol")); err != nil {
		if !fortiAPIPatch(o["detectprotocol"]) {
			return fmt.Errorf("Error reading detectprotocol: %v", err)
		}
	}

	if err = d.Set("src_check", dataSourceFlattenSystemInterfaceSrcCheck(o["src-check"], d, "src_check")); err != nil {
		if !fortiAPIPatch(o["src-check"]) {
			return fmt.Errorf("Error reading src_check: %v", err)
		}
	}

	if err = d.Set("ipv6", dataSourceFlattenSystemInterfaceIpv6(o["ipv6"], d, "ipv6")); err != nil {
		if !fortiAPIPatch(o["ipv6"]) {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	if err = d.Set("dynamicgw", dataSourceFlattenSystemInterfaceDynamicgw(o["dynamicgw"], d, "dynamicgw")); err != nil {
		if !fortiAPIPatch(o["dynamicgw"]) {
			return fmt.Errorf("Error reading dynamicgw: %v", err)
		}
	}

	if err = d.Set("mtu_override", dataSourceFlattenSystemInterfaceMtuOverride(o["mtu-override"], d, "mtu_override")); err != nil {
		if !fortiAPIPatch(o["mtu-override"]) {
			return fmt.Errorf("Error reading mtu_override: %v", err)
		}
	}

	if err = d.Set("mtu", dataSourceFlattenSystemInterfaceMtu(o["mtu"], d, "mtu")); err != nil {
		if !fortiAPIPatch(o["mtu"]) {
			return fmt.Errorf("Error reading mtu: %v", err)
		}
	}

	if err = d.Set("dynamic_dns1", dataSourceFlattenSystemInterfaceDynamic_Dns1(o["dynamic_dns1"], d, "dynamic_dns1")); err != nil {
		if !fortiAPIPatch(o["dynamic_dns1"]) {
			return fmt.Errorf("Error reading dynamic_dns1: %v", err)
		}
	}

	if err = d.Set("alias", dataSourceFlattenSystemInterfaceAlias(o["alias"], d, "alias")); err != nil {
		if !fortiAPIPatch(o["alias"]) {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("auth_type", dataSourceFlattenSystemInterfaceAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("dynamic_dns2", dataSourceFlattenSystemInterfaceDynamic_Dns2(o["dynamic_dns2"], d, "dynamic_dns2")); err != nil {
		if !fortiAPIPatch(o["dynamic_dns2"]) {
			return fmt.Errorf("Error reading dynamic_dns2: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemInterfaceStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("mode", dataSourceFlattenSystemInterfaceMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("ha_priority", dataSourceFlattenSystemInterfaceHaPriority(o["ha-priority"], d, "ha_priority")); err != nil {
		if !fortiAPIPatch(o["ha-priority"]) {
			return fmt.Errorf("Error reading ha_priority: %v", err)
		}
	}

	if err = d.Set("secondaryip", dataSourceFlattenSystemInterfaceSecondaryip(o["secondaryip"], d, "secondaryip")); err != nil {
		if !fortiAPIPatch(o["secondaryip"]) {
			return fmt.Errorf("Error reading secondaryip: %v", err)
		}
	}

	if err = d.Set("switch_members", dataSourceFlattenSystemInterfaceSwitchMembers(o["switch-members"], d, "switch_members")); err != nil {
		if !fortiAPIPatch(o["switch-members"]) {
			return fmt.Errorf("Error reading switch_members: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemInterfaceFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
