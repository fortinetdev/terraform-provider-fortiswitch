---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_logdisk_setting"
description: |-
  Settings for local disk logging.
---

# fortiswitch_logdisk_setting
Settings for local disk logging.

## Example Usage

```hcl
resource "fortiswitch_logdisk_setting" "name" {
        diskfull = "overwrite"
        drive_standby_time = "4"
        roll_schedule = "daily"
        roll_time = "11:11"
        source_ip = "84.230.14.43"
        status = "enable"
        upload = "disable"
        upload_delete_files = "enable"
        upload_destination = "ftp_server"
        upload_format = "compact"
        upload_ssl_conn = "default"
        uploadport = "24"
        uploadsched = "disable"
        uploadtime = "23"
        uploadtype = "traffic"
        uploadzip = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `uploaduser` - User account in the uploading server.
* `uploadip` - IP address of the log uploading server.
* `uploadtime` - Time of scheduled upload.
* `diskfull` - Policy to apply when disk is full. Valid values: `overwrite`, `nolog`.
* `roll_schedule` - Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
* `report_quota` - Report quota.
* `roll_time` - Time to roll logs [hh:mm].
* `max_log_file_size` - Max log file size in MB before rolling (may not be accurate all the time).
* `uploadtype` - Types of log files that need to be uploaded. Valid values: `traffic`, `event`, `virus`, `webfilter`, `attack`, `spamfilter`, `dlp-archive`, `dlp`, `app-ctrl`.
* `source_ip` - Source IP address of the disk log uploading.
* `status` - Enable/disable local disk log. Valid values: `enable`, `disable`.
* `drive_standby_time` - Power management timeout(0-19800 sec)(0 disable).
* `upload_format` - Upload compact/text logs. Valid values: `compact`, `text`.
* `full_final_warning_threshold` - Log full final warning threshold(3-100), the default is 95.
* `uploadpass` - Password of the user account in the uploading server.
* `upload_ssl_conn` - Enable/disable SSL communication when uploading. Valid values: `default`, `high`, `low`, `disable`.
* `upload_delete_files` - Delete log files after uploading (default=enable). Valid values: `enable`, `disable`.
* `log_quota` - Disk log quota.
* `uploaddir` - Log file uploading remote directory.
* `full_first_warning_threshold` - Log full first warning threshold(1-98), the default is 75.
* `uploadport` - Port of the log uploading server.
* `upload` - Whether to upload the log file when rolling. Valid values: `enable`, `disable`.
* `upload_destination` - Server type. Valid values: `ftp-server`.
* `uploadsched` - Scheduled upload (disable=upload when rolling). Valid values: `disable`, `enable`.
* `uploadzip` - Compress upload logs. Valid values: `disable`, `enable`.
* `roll_day` - Days of week to roll logs. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
* `full_second_warning_threshold` - Log full second warning threshold(2-99), the default is 90.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogDisk Setting can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_logdisk_setting.labelname LogDiskSetting

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_logdisk_setting.labelname LogDiskSetting
$ unset "FORTISWITCH_IMPORT_TABLE"
```
