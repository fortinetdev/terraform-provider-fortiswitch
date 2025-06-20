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

func dataSourceSystemTosBasedPriorityList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemTosBasedPriorityListRead,

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

func dataSourceSystemTosBasedPriorityListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	filter := d.Get("filter").(string)
	if filter != "" {
		filter = escapeFilter(filter)
	}

	o, err := c.GenericGroupRead("/api/v2/cmdb/system/tos-based-priority", filter)
	if err != nil {
		return fmt.Errorf("Error describing SystemTosBasedPriority: %v", err)
	}

	var tmps []int
	if o != nil {
		for _, r := range o {
			if r == nil {
				continue
			}
			i := r.(map[string]interface{})

			if _, ok := i["id"]; ok {
				tmps = append(tmps, fortiIntValue(i["id"]))
			}
		}
	}
	d.Set("fswidlist", tmps)

	d.SetId("DataSourceSystemTosBasedPriorityList")

	return nil
}
