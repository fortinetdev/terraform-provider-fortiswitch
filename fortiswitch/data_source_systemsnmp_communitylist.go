// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

package fortiswitch

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemSnmpCommunityList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemSnmpCommunityListRead,

		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fswidlist": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func dataSourceSystemSnmpCommunityListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	filter := d.Get("filter").(string)
	if filter != "" {
		filter = escapeFilter(filter)
	}

	o, err := c.GenericGroupRead("/api/v2/cmdb/system.snmp/community", filter)
	if err != nil {
		return fmt.Errorf("Error describing SystemSnmpCommunity: %v", err)
	}

	var tmps []int
	if o != nil {
		if len(o) == 0 || o[0] == nil {
			return nil
		}

		for _, r := range o {
			i := r.(map[string]interface{})

			if _, ok := i["id"]; ok {
				tmps = append(tmps, fortiIntValue(i["id"]))
			}
		}
	}
	d.Set("fswidlist", tmps)

	d.SetId("DataSourceSystemSnmpCommunityList")

	return nil
}
