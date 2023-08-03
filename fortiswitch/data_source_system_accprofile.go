// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure system administrative access group.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemAccprofile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAccprofileRead,
		Schema: map[string]*schema.Schema{

			"exec_alias_grp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"swmonguardgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sysgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pktmongrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"loggrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mntgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"netgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admingrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"routegrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"swcoregrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"utilgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alias_commands": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemAccprofileRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemAccprofile: type error")
	}

	o, err := c.ReadSystemAccprofile(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemAccprofile: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemAccprofile(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemAccprofile from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemAccprofileExecAliasGrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSwmonguardgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSysgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofilePktmongrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileLoggrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileMntgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileNetgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileAdmingrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileRoutegrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSwcoregrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtilgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileAliasCommands(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_name"
		if _, ok := i["command-name"]; ok {
			tmp["command_name"] = dataSourceFlattenSystemAccprofileAliasCommandsCommandName(i["command-name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAccprofileAliasCommandsCommandName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAccprofile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("exec_alias_grp", dataSourceFlattenSystemAccprofileExecAliasGrp(o["exec-alias-grp"], d, "exec_alias_grp")); err != nil {
		if !fortiAPIPatch(o["exec-alias-grp"]) {
			return fmt.Errorf("Error reading exec_alias_grp: %v", err)
		}
	}

	if err = d.Set("swmonguardgrp", dataSourceFlattenSystemAccprofileSwmonguardgrp(o["swmonguardgrp"], d, "swmonguardgrp")); err != nil {
		if !fortiAPIPatch(o["swmonguardgrp"]) {
			return fmt.Errorf("Error reading swmonguardgrp: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenSystemAccprofileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("sysgrp", dataSourceFlattenSystemAccprofileSysgrp(o["sysgrp"], d, "sysgrp")); err != nil {
		if !fortiAPIPatch(o["sysgrp"]) {
			return fmt.Errorf("Error reading sysgrp: %v", err)
		}
	}

	if err = d.Set("pktmongrp", dataSourceFlattenSystemAccprofilePktmongrp(o["pktmongrp"], d, "pktmongrp")); err != nil {
		if !fortiAPIPatch(o["pktmongrp"]) {
			return fmt.Errorf("Error reading pktmongrp: %v", err)
		}
	}

	if err = d.Set("loggrp", dataSourceFlattenSystemAccprofileLoggrp(o["loggrp"], d, "loggrp")); err != nil {
		if !fortiAPIPatch(o["loggrp"]) {
			return fmt.Errorf("Error reading loggrp: %v", err)
		}
	}

	if err = d.Set("mntgrp", dataSourceFlattenSystemAccprofileMntgrp(o["mntgrp"], d, "mntgrp")); err != nil {
		if !fortiAPIPatch(o["mntgrp"]) {
			return fmt.Errorf("Error reading mntgrp: %v", err)
		}
	}

	if err = d.Set("netgrp", dataSourceFlattenSystemAccprofileNetgrp(o["netgrp"], d, "netgrp")); err != nil {
		if !fortiAPIPatch(o["netgrp"]) {
			return fmt.Errorf("Error reading netgrp: %v", err)
		}
	}

	if err = d.Set("admingrp", dataSourceFlattenSystemAccprofileAdmingrp(o["admingrp"], d, "admingrp")); err != nil {
		if !fortiAPIPatch(o["admingrp"]) {
			return fmt.Errorf("Error reading admingrp: %v", err)
		}
	}

	if err = d.Set("routegrp", dataSourceFlattenSystemAccprofileRoutegrp(o["routegrp"], d, "routegrp")); err != nil {
		if !fortiAPIPatch(o["routegrp"]) {
			return fmt.Errorf("Error reading routegrp: %v", err)
		}
	}

	if err = d.Set("swcoregrp", dataSourceFlattenSystemAccprofileSwcoregrp(o["swcoregrp"], d, "swcoregrp")); err != nil {
		if !fortiAPIPatch(o["swcoregrp"]) {
			return fmt.Errorf("Error reading swcoregrp: %v", err)
		}
	}

	if err = d.Set("utilgrp", dataSourceFlattenSystemAccprofileUtilgrp(o["utilgrp"], d, "utilgrp")); err != nil {
		if !fortiAPIPatch(o["utilgrp"]) {
			return fmt.Errorf("Error reading utilgrp: %v", err)
		}
	}

	if err = d.Set("alias_commands", dataSourceFlattenSystemAccprofileAliasCommands(o["alias-commands"], d, "alias_commands")); err != nil {
		if !fortiAPIPatch(o["alias-commands"]) {
			return fmt.Errorf("Error reading alias_commands: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemAccprofileFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}
