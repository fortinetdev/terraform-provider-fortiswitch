// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Settings for local disk logging.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogDiskSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogDiskSettingUpdate,
		Read:   resourceLogDiskSettingRead,
		Update: resourceLogDiskSettingUpdate,
		Delete: resourceLogDiskSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"uploaduser": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"uploadip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uploadtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"diskfull": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"roll_schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"report_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"roll_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_log_file_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"uploadtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"drive_standby_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"upload_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"full_final_warning_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 100),
				Optional:     true,
				Computed:     true,
			},
			"uploadpass": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"upload_ssl_conn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_delete_files": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"uploaddir": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"full_first_warning_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 98),
				Optional:     true,
				Computed:     true,
			},
			"uploadport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"upload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_destination": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uploadsched": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uploadzip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"roll_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"full_second_warning_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 99),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceLogDiskSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogDiskSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogDiskSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogDiskSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogDiskSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogDiskSetting")
	}

	return resourceLogDiskSettingRead(d, m)
}

func resourceLogDiskSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogDiskSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogDiskSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogDiskSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing LogDiskSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogDiskSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogDiskSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogDiskSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogDiskSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogDiskSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogDiskSettingUploaduser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadtime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingDiskfull(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingRollSchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingReportQuota(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingRollTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingMaxLogFileSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadtype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingDriveStandbyTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingFullFinalWarningThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadpass(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadSslConn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadDeleteFiles(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingLogQuota(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploaddir(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingFullFirstWarningThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUpload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadDestination(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadsched(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadzip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingRollDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingFullSecondWarningThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogDiskSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("uploaduser", flattenLogDiskSettingUploaduser(o["uploaduser"], d, "uploaduser", sv)); err != nil {
		if !fortiAPIPatch(o["uploaduser"]) {
			return fmt.Errorf("Error reading uploaduser: %v", err)
		}
	}

	if err = d.Set("uploadip", flattenLogDiskSettingUploadip(o["uploadip"], d, "uploadip", sv)); err != nil {
		if !fortiAPIPatch(o["uploadip"]) {
			return fmt.Errorf("Error reading uploadip: %v", err)
		}
	}

	if err = d.Set("uploadtime", flattenLogDiskSettingUploadtime(o["uploadtime"], d, "uploadtime", sv)); err != nil {
		if !fortiAPIPatch(o["uploadtime"]) {
			return fmt.Errorf("Error reading uploadtime: %v", err)
		}
	}

	if err = d.Set("diskfull", flattenLogDiskSettingDiskfull(o["diskfull"], d, "diskfull", sv)); err != nil {
		if !fortiAPIPatch(o["diskfull"]) {
			return fmt.Errorf("Error reading diskfull: %v", err)
		}
	}

	if err = d.Set("roll_schedule", flattenLogDiskSettingRollSchedule(o["roll-schedule"], d, "roll_schedule", sv)); err != nil {
		if !fortiAPIPatch(o["roll-schedule"]) {
			return fmt.Errorf("Error reading roll_schedule: %v", err)
		}
	}

	if err = d.Set("report_quota", flattenLogDiskSettingReportQuota(o["report-quota"], d, "report_quota", sv)); err != nil {
		if !fortiAPIPatch(o["report-quota"]) {
			return fmt.Errorf("Error reading report_quota: %v", err)
		}
	}

	if err = d.Set("roll_time", flattenLogDiskSettingRollTime(o["roll-time"], d, "roll_time", sv)); err != nil {
		if !fortiAPIPatch(o["roll-time"]) {
			return fmt.Errorf("Error reading roll_time: %v", err)
		}
	}

	if err = d.Set("max_log_file_size", flattenLogDiskSettingMaxLogFileSize(o["max-log-file-size"], d, "max_log_file_size", sv)); err != nil {
		if !fortiAPIPatch(o["max-log-file-size"]) {
			return fmt.Errorf("Error reading max_log_file_size: %v", err)
		}
	}

	if err = d.Set("uploadtype", flattenLogDiskSettingUploadtype(o["uploadtype"], d, "uploadtype", sv)); err != nil {
		if !fortiAPIPatch(o["uploadtype"]) {
			return fmt.Errorf("Error reading uploadtype: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogDiskSettingSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("status", flattenLogDiskSettingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("drive_standby_time", flattenLogDiskSettingDriveStandbyTime(o["drive-standby-time"], d, "drive_standby_time", sv)); err != nil {
		if !fortiAPIPatch(o["drive-standby-time"]) {
			return fmt.Errorf("Error reading drive_standby_time: %v", err)
		}
	}

	if err = d.Set("upload_format", flattenLogDiskSettingUploadFormat(o["upload-format"], d, "upload_format", sv)); err != nil {
		if !fortiAPIPatch(o["upload-format"]) {
			return fmt.Errorf("Error reading upload_format: %v", err)
		}
	}

	if err = d.Set("full_final_warning_threshold", flattenLogDiskSettingFullFinalWarningThreshold(o["full-final-warning-threshold"], d, "full_final_warning_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["full-final-warning-threshold"]) {
			return fmt.Errorf("Error reading full_final_warning_threshold: %v", err)
		}
	}

	if err = d.Set("uploadpass", flattenLogDiskSettingUploadpass(o["uploadpass"], d, "uploadpass", sv)); err != nil {
		if !fortiAPIPatch(o["uploadpass"]) {
			return fmt.Errorf("Error reading uploadpass: %v", err)
		}
	}

	if err = d.Set("upload_ssl_conn", flattenLogDiskSettingUploadSslConn(o["upload-ssl-conn"], d, "upload_ssl_conn", sv)); err != nil {
		if !fortiAPIPatch(o["upload-ssl-conn"]) {
			return fmt.Errorf("Error reading upload_ssl_conn: %v", err)
		}
	}

	if err = d.Set("upload_delete_files", flattenLogDiskSettingUploadDeleteFiles(o["upload-delete-files"], d, "upload_delete_files", sv)); err != nil {
		if !fortiAPIPatch(o["upload-delete-files"]) {
			return fmt.Errorf("Error reading upload_delete_files: %v", err)
		}
	}

	if err = d.Set("log_quota", flattenLogDiskSettingLogQuota(o["log-quota"], d, "log_quota", sv)); err != nil {
		if !fortiAPIPatch(o["log-quota"]) {
			return fmt.Errorf("Error reading log_quota: %v", err)
		}
	}

	if err = d.Set("uploaddir", flattenLogDiskSettingUploaddir(o["uploaddir"], d, "uploaddir", sv)); err != nil {
		if !fortiAPIPatch(o["uploaddir"]) {
			return fmt.Errorf("Error reading uploaddir: %v", err)
		}
	}

	if err = d.Set("full_first_warning_threshold", flattenLogDiskSettingFullFirstWarningThreshold(o["full-first-warning-threshold"], d, "full_first_warning_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["full-first-warning-threshold"]) {
			return fmt.Errorf("Error reading full_first_warning_threshold: %v", err)
		}
	}

	if err = d.Set("uploadport", flattenLogDiskSettingUploadport(o["uploadport"], d, "uploadport", sv)); err != nil {
		if !fortiAPIPatch(o["uploadport"]) {
			return fmt.Errorf("Error reading uploadport: %v", err)
		}
	}

	if err = d.Set("upload", flattenLogDiskSettingUpload(o["upload"], d, "upload", sv)); err != nil {
		if !fortiAPIPatch(o["upload"]) {
			return fmt.Errorf("Error reading upload: %v", err)
		}
	}

	if err = d.Set("upload_destination", flattenLogDiskSettingUploadDestination(o["upload-destination"], d, "upload_destination", sv)); err != nil {
		if !fortiAPIPatch(o["upload-destination"]) {
			return fmt.Errorf("Error reading upload_destination: %v", err)
		}
	}

	if err = d.Set("uploadsched", flattenLogDiskSettingUploadsched(o["uploadsched"], d, "uploadsched", sv)); err != nil {
		if !fortiAPIPatch(o["uploadsched"]) {
			return fmt.Errorf("Error reading uploadsched: %v", err)
		}
	}

	if err = d.Set("uploadzip", flattenLogDiskSettingUploadzip(o["uploadzip"], d, "uploadzip", sv)); err != nil {
		if !fortiAPIPatch(o["uploadzip"]) {
			return fmt.Errorf("Error reading uploadzip: %v", err)
		}
	}

	if err = d.Set("roll_day", flattenLogDiskSettingRollDay(o["roll-day"], d, "roll_day", sv)); err != nil {
		if !fortiAPIPatch(o["roll-day"]) {
			return fmt.Errorf("Error reading roll_day: %v", err)
		}
	}

	if err = d.Set("full_second_warning_threshold", flattenLogDiskSettingFullSecondWarningThreshold(o["full-second-warning-threshold"], d, "full_second_warning_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["full-second-warning-threshold"]) {
			return fmt.Errorf("Error reading full_second_warning_threshold: %v", err)
		}
	}

	return nil
}

func flattenLogDiskSettingFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandLogDiskSettingUploaduser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadtime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingDiskfull(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingRollSchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingReportQuota(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingRollTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingMaxLogFileSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadtype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingDriveStandbyTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingFullFinalWarningThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadpass(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadSslConn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadDeleteFiles(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingLogQuota(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploaddir(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingFullFirstWarningThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUpload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadDestination(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadsched(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadzip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingRollDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingFullSecondWarningThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogDiskSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("uploaduser"); ok {
		if setArgNil {
			obj["uploaduser"] = nil
		} else {

			t, err := expandLogDiskSettingUploaduser(d, v, "uploaduser", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploaduser"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploadip"); ok {
		if setArgNil {
			obj["uploadip"] = nil
		} else {

			t, err := expandLogDiskSettingUploadip(d, v, "uploadip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploadip"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploadtime"); ok {
		if setArgNil {
			obj["uploadtime"] = nil
		} else {

			t, err := expandLogDiskSettingUploadtime(d, v, "uploadtime", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploadtime"] = t
			}
		}
	}

	if v, ok := d.GetOk("diskfull"); ok {
		if setArgNil {
			obj["diskfull"] = nil
		} else {

			t, err := expandLogDiskSettingDiskfull(d, v, "diskfull", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["diskfull"] = t
			}
		}
	}

	if v, ok := d.GetOk("roll_schedule"); ok {
		if setArgNil {
			obj["roll-schedule"] = nil
		} else {

			t, err := expandLogDiskSettingRollSchedule(d, v, "roll_schedule", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["roll-schedule"] = t
			}
		}
	}

	if v, ok := d.GetOk("report_quota"); ok {
		if setArgNil {
			obj["report-quota"] = nil
		} else {

			t, err := expandLogDiskSettingReportQuota(d, v, "report_quota", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["report-quota"] = t
			}
		}
	}

	if v, ok := d.GetOk("roll_time"); ok {
		if setArgNil {
			obj["roll-time"] = nil
		} else {

			t, err := expandLogDiskSettingRollTime(d, v, "roll_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["roll-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_log_file_size"); ok {
		if setArgNil {
			obj["max-log-file-size"] = nil
		} else {

			t, err := expandLogDiskSettingMaxLogFileSize(d, v, "max_log_file_size", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-log-file-size"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploadtype"); ok {
		if setArgNil {
			obj["uploadtype"] = nil
		} else {

			t, err := expandLogDiskSettingUploadtype(d, v, "uploadtype", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploadtype"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {

			t, err := expandLogDiskSettingSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandLogDiskSettingStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("drive_standby_time"); ok {
		if setArgNil {
			obj["drive-standby-time"] = nil
		} else {

			t, err := expandLogDiskSettingDriveStandbyTime(d, v, "drive_standby_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["drive-standby-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_format"); ok {
		if setArgNil {
			obj["upload-format"] = nil
		} else {

			t, err := expandLogDiskSettingUploadFormat(d, v, "upload_format", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-format"] = t
			}
		}
	}

	if v, ok := d.GetOk("full_final_warning_threshold"); ok {
		if setArgNil {
			obj["full-final-warning-threshold"] = nil
		} else {

			t, err := expandLogDiskSettingFullFinalWarningThreshold(d, v, "full_final_warning_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["full-final-warning-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploadpass"); ok {
		if setArgNil {
			obj["uploadpass"] = nil
		} else {

			t, err := expandLogDiskSettingUploadpass(d, v, "uploadpass", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploadpass"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_ssl_conn"); ok {
		if setArgNil {
			obj["upload-ssl-conn"] = nil
		} else {

			t, err := expandLogDiskSettingUploadSslConn(d, v, "upload_ssl_conn", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-ssl-conn"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_delete_files"); ok {
		if setArgNil {
			obj["upload-delete-files"] = nil
		} else {

			t, err := expandLogDiskSettingUploadDeleteFiles(d, v, "upload_delete_files", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-delete-files"] = t
			}
		}
	}

	if v, ok := d.GetOk("log_quota"); ok {
		if setArgNil {
			obj["log-quota"] = nil
		} else {

			t, err := expandLogDiskSettingLogQuota(d, v, "log_quota", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-quota"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploaddir"); ok {
		if setArgNil {
			obj["uploaddir"] = nil
		} else {

			t, err := expandLogDiskSettingUploaddir(d, v, "uploaddir", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploaddir"] = t
			}
		}
	}

	if v, ok := d.GetOk("full_first_warning_threshold"); ok {
		if setArgNil {
			obj["full-first-warning-threshold"] = nil
		} else {

			t, err := expandLogDiskSettingFullFirstWarningThreshold(d, v, "full_first_warning_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["full-first-warning-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploadport"); ok {
		if setArgNil {
			obj["uploadport"] = nil
		} else {

			t, err := expandLogDiskSettingUploadport(d, v, "uploadport", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploadport"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload"); ok {
		if setArgNil {
			obj["upload"] = nil
		} else {

			t, err := expandLogDiskSettingUpload(d, v, "upload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_destination"); ok {
		if setArgNil {
			obj["upload-destination"] = nil
		} else {

			t, err := expandLogDiskSettingUploadDestination(d, v, "upload_destination", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-destination"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploadsched"); ok {
		if setArgNil {
			obj["uploadsched"] = nil
		} else {

			t, err := expandLogDiskSettingUploadsched(d, v, "uploadsched", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploadsched"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploadzip"); ok {
		if setArgNil {
			obj["uploadzip"] = nil
		} else {

			t, err := expandLogDiskSettingUploadzip(d, v, "uploadzip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploadzip"] = t
			}
		}
	}

	if v, ok := d.GetOk("roll_day"); ok {
		if setArgNil {
			obj["roll-day"] = nil
		} else {

			t, err := expandLogDiskSettingRollDay(d, v, "roll_day", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["roll-day"] = t
			}
		}
	}

	if v, ok := d.GetOk("full_second_warning_threshold"); ok {
		if setArgNil {
			obj["full-second-warning-threshold"] = nil
		} else {

			t, err := expandLogDiskSettingFullSecondWarningThreshold(d, v, "full_second_warning_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["full-second-warning-threshold"] = t
			}
		}
	}

	return &obj, nil
}
