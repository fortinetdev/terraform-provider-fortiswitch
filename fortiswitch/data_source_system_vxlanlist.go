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

func dataSourceSystemVxlanList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemVxlanListRead,

		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"namelist": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceSystemVxlanListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	filter := d.Get("filter").(string)
	if filter != "" {
		filter = escapeFilter(filter)
	}

	o, err := c.GenericGroupRead("/api/v2/cmdb/system/vxlan", filter)
	if err != nil {
		return fmt.Errorf("Error describing SystemVxlan: %v", err)
	}

	var tmps []string
	if o != nil {
		for _, r := range o {
			if r == nil {
				continue
			}
			i := r.(map[string]interface{})

			if _, ok := i["name"]; ok {
				tmps = append(tmps, fortiStringValue(i["name"]))
			}
		}
	}
	d.Set("namelist", tmps)

	d.SetId("DataSourceSystemVxlanList")

	return nil
}
