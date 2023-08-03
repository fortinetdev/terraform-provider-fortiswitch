// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: System Flow Export settings.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFlowExport() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFlowExportUpdate,
		Read:   resourceSystemFlowExportRead,
		Update: resourceSystemFlowExportUpdate,
		Delete: resourceSystemFlowExportDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"timeout_tcp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 604800),
				Optional:     true,
				Computed:     true,
			},
			"timeout_udp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 604800),
				Optional:     true,
				Computed:     true,
			},
			"collectors": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
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
						"transport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout_tcp_rst": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 604800),
				Optional:     true,
				Computed:     true,
			},
			"aggregates": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
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
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout_tcp_fin": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 604800),
				Optional:     true,
				Computed:     true,
			},
			"template_export_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),
				Optional:     true,
				Computed:     true,
			},
			"filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"timeout_icmp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 604800),
				Optional:     true,
				Computed:     true,
			},
			"max_export_pkt_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(512, 9216),
				Optional:     true,
				Computed:     true,
			},
			"timeout_general": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 604800),
				Optional:     true,
				Computed:     true,
			},
			"identity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout_max": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 604800),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemFlowExportUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemFlowExport(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemFlowExport resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFlowExport(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemFlowExport resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFlowExport")
	}

	return resourceSystemFlowExportRead(d, m)
}

func resourceSystemFlowExportDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemFlowExport(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemFlowExport resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFlowExport(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing SystemFlowExport resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFlowExportRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemFlowExport(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemFlowExport resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFlowExport(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemFlowExport resource from API: %v", err)
	}
	return nil
}

func flattenSystemFlowExportTimeoutTcp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportTimeoutUdp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportCollectors(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenSystemFlowExportCollectorsIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemFlowExportCollectorsName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transport"
		if _, ok := i["transport"]; ok {

			tmp["transport"] = flattenSystemFlowExportCollectorsTransport(i["transport"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenSystemFlowExportCollectorsPort(i["port"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemFlowExportCollectorsIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportCollectorsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportCollectorsTransport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportCollectorsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportTimeoutTcpRst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportAggregates(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenSystemFlowExportAggregatesIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSystemFlowExportAggregatesId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemFlowExportAggregatesIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemFlowExportAggregatesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportTimeoutTcpFin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportTemplateExportPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportTimeoutIcmp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportMaxExportPktSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportTimeoutGeneral(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportIdentity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFlowExportTimeoutMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemFlowExport(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("timeout_tcp", flattenSystemFlowExportTimeoutTcp(o["timeout-tcp"], d, "timeout_tcp", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-tcp"]) {
			return fmt.Errorf("Error reading timeout_tcp: %v", err)
		}
	}

	if err = d.Set("timeout_udp", flattenSystemFlowExportTimeoutUdp(o["timeout-udp"], d, "timeout_udp", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-udp"]) {
			return fmt.Errorf("Error reading timeout_udp: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("collectors", flattenSystemFlowExportCollectors(o["collectors"], d, "collectors", sv)); err != nil {
			if !fortiAPIPatch(o["collectors"]) {
				return fmt.Errorf("Error reading collectors: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("collectors"); ok {
			if err = d.Set("collectors", flattenSystemFlowExportCollectors(o["collectors"], d, "collectors", sv)); err != nil {
				if !fortiAPIPatch(o["collectors"]) {
					return fmt.Errorf("Error reading collectors: %v", err)
				}
			}
		}
	}

	if err = d.Set("level", flattenSystemFlowExportLevel(o["level"], d, "level", sv)); err != nil {
		if !fortiAPIPatch(o["level"]) {
			return fmt.Errorf("Error reading level: %v", err)
		}
	}

	if err = d.Set("timeout_tcp_rst", flattenSystemFlowExportTimeoutTcpRst(o["timeout-tcp-rst"], d, "timeout_tcp_rst", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-tcp-rst"]) {
			return fmt.Errorf("Error reading timeout_tcp_rst: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("aggregates", flattenSystemFlowExportAggregates(o["aggregates"], d, "aggregates", sv)); err != nil {
			if !fortiAPIPatch(o["aggregates"]) {
				return fmt.Errorf("Error reading aggregates: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("aggregates"); ok {
			if err = d.Set("aggregates", flattenSystemFlowExportAggregates(o["aggregates"], d, "aggregates", sv)); err != nil {
				if !fortiAPIPatch(o["aggregates"]) {
					return fmt.Errorf("Error reading aggregates: %v", err)
				}
			}
		}
	}

	if err = d.Set("format", flattenSystemFlowExportFormat(o["format"], d, "format", sv)); err != nil {
		if !fortiAPIPatch(o["format"]) {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	if err = d.Set("timeout_tcp_fin", flattenSystemFlowExportTimeoutTcpFin(o["timeout-tcp-fin"], d, "timeout_tcp_fin", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-tcp-fin"]) {
			return fmt.Errorf("Error reading timeout_tcp_fin: %v", err)
		}
	}

	if err = d.Set("template_export_period", flattenSystemFlowExportTemplateExportPeriod(o["template-export-period"], d, "template_export_period", sv)); err != nil {
		if !fortiAPIPatch(o["template-export-period"]) {
			return fmt.Errorf("Error reading template_export_period: %v", err)
		}
	}

	if err = d.Set("filter", flattenSystemFlowExportFilter(o["filter"], d, "filter", sv)); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("timeout_icmp", flattenSystemFlowExportTimeoutIcmp(o["timeout-icmp"], d, "timeout_icmp", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-icmp"]) {
			return fmt.Errorf("Error reading timeout_icmp: %v", err)
		}
	}

	if err = d.Set("max_export_pkt_size", flattenSystemFlowExportMaxExportPktSize(o["max-export-pkt-size"], d, "max_export_pkt_size", sv)); err != nil {
		if !fortiAPIPatch(o["max-export-pkt-size"]) {
			return fmt.Errorf("Error reading max_export_pkt_size: %v", err)
		}
	}

	if err = d.Set("timeout_general", flattenSystemFlowExportTimeoutGeneral(o["timeout-general"], d, "timeout_general", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-general"]) {
			return fmt.Errorf("Error reading timeout_general: %v", err)
		}
	}

	if err = d.Set("identity", flattenSystemFlowExportIdentity(o["identity"], d, "identity", sv)); err != nil {
		if !fortiAPIPatch(o["identity"]) {
			return fmt.Errorf("Error reading identity: %v", err)
		}
	}

	if err = d.Set("timeout_max", flattenSystemFlowExportTimeoutMax(o["timeout-max"], d, "timeout_max", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-max"]) {
			return fmt.Errorf("Error reading timeout_max: %v", err)
		}
	}

	return nil
}

func flattenSystemFlowExportFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemFlowExportTimeoutTcp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportTimeoutUdp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportCollectors(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandSystemFlowExportCollectorsIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemFlowExportCollectorsName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transport"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["transport"], _ = expandSystemFlowExportCollectorsTransport(d, i["transport"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandSystemFlowExportCollectorsPort(d, i["port"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemFlowExportCollectorsIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportCollectorsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportCollectorsTransport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportCollectorsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportTimeoutTcpRst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportAggregates(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandSystemFlowExportAggregatesIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSystemFlowExportAggregatesId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemFlowExportAggregatesIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportAggregatesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportTimeoutTcpFin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportTemplateExportPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportTimeoutIcmp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportMaxExportPktSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportTimeoutGeneral(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportIdentity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFlowExportTimeoutMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFlowExport(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("timeout_tcp"); ok {
		if setArgNil {
			obj["timeout-tcp"] = nil
		} else {

			t, err := expandSystemFlowExportTimeoutTcp(d, v, "timeout_tcp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout-tcp"] = t
			}
		}
	}

	if v, ok := d.GetOk("timeout_udp"); ok {
		if setArgNil {
			obj["timeout-udp"] = nil
		} else {

			t, err := expandSystemFlowExportTimeoutUdp(d, v, "timeout_udp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout-udp"] = t
			}
		}
	}

	if v, ok := d.GetOk("collectors"); ok || d.HasChange("collectors") {
		if setArgNil {
			obj["collectors"] = make([]struct{}, 0)
		} else {

			t, err := expandSystemFlowExportCollectors(d, v, "collectors", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["collectors"] = t
			}
		}
	}

	if v, ok := d.GetOk("level"); ok {
		if setArgNil {
			obj["level"] = nil
		} else {

			t, err := expandSystemFlowExportLevel(d, v, "level", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["level"] = t
			}
		}
	}

	if v, ok := d.GetOk("timeout_tcp_rst"); ok {
		if setArgNil {
			obj["timeout-tcp-rst"] = nil
		} else {

			t, err := expandSystemFlowExportTimeoutTcpRst(d, v, "timeout_tcp_rst", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout-tcp-rst"] = t
			}
		}
	}

	if v, ok := d.GetOk("aggregates"); ok || d.HasChange("aggregates") {
		if setArgNil {
			obj["aggregates"] = make([]struct{}, 0)
		} else {

			t, err := expandSystemFlowExportAggregates(d, v, "aggregates", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["aggregates"] = t
			}
		}
	}

	if v, ok := d.GetOk("format"); ok {
		if setArgNil {
			obj["format"] = nil
		} else {

			t, err := expandSystemFlowExportFormat(d, v, "format", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["format"] = t
			}
		}
	}

	if v, ok := d.GetOk("timeout_tcp_fin"); ok {
		if setArgNil {
			obj["timeout-tcp-fin"] = nil
		} else {

			t, err := expandSystemFlowExportTimeoutTcpFin(d, v, "timeout_tcp_fin", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout-tcp-fin"] = t
			}
		}
	}

	if v, ok := d.GetOk("template_export_period"); ok {
		if setArgNil {
			obj["template-export-period"] = nil
		} else {

			t, err := expandSystemFlowExportTemplateExportPeriod(d, v, "template_export_period", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["template-export-period"] = t
			}
		}
	}

	if v, ok := d.GetOk("filter"); ok {
		if setArgNil {
			obj["filter"] = nil
		} else {

			t, err := expandSystemFlowExportFilter(d, v, "filter", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["filter"] = t
			}
		}
	}

	if v, ok := d.GetOk("timeout_icmp"); ok {
		if setArgNil {
			obj["timeout-icmp"] = nil
		} else {

			t, err := expandSystemFlowExportTimeoutIcmp(d, v, "timeout_icmp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout-icmp"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_export_pkt_size"); ok {
		if setArgNil {
			obj["max-export-pkt-size"] = nil
		} else {

			t, err := expandSystemFlowExportMaxExportPktSize(d, v, "max_export_pkt_size", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-export-pkt-size"] = t
			}
		}
	}

	if v, ok := d.GetOk("timeout_general"); ok {
		if setArgNil {
			obj["timeout-general"] = nil
		} else {

			t, err := expandSystemFlowExportTimeoutGeneral(d, v, "timeout_general", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout-general"] = t
			}
		}
	}

	if v, ok := d.GetOk("identity"); ok {
		if setArgNil {
			obj["identity"] = nil
		} else {

			t, err := expandSystemFlowExportIdentity(d, v, "identity", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["identity"] = t
			}
		}
	}

	if v, ok := d.GetOk("timeout_max"); ok {
		if setArgNil {
			obj["timeout-max"] = nil
		} else {

			t, err := expandSystemFlowExportTimeoutMax(d, v, "timeout_max", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout-max"] = t
			}
		}
	}

	return &obj, nil
}
