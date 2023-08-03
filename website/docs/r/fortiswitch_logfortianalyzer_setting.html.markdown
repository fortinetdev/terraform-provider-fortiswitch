---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_logfortianalyzer_setting"
description: |-
  Setting for FortiAnalyzer.
---

# fortiswitch_logfortianalyzer_setting
Setting for FortiAnalyzer.

## Argument Reference

The following arguments are supported:

* `upload_interval` - Frequency to check log file for upload. Valid values: `daily`, `weekly`, `monthly`.
* `fdp_device` - Serial number of FortiAnalyzer to connect to.
* `upload_option` - Enable/disable logging to hard disk and then upload to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`.
* `buffer_max_send` - Maximum log transmission rate for buffered logs.
* `address_mode` - FortiAnalyzer IP addressing mode. Valid values: `static`, `auto-discovery`.
* `enc_algorithm` - Whether to send FortiAnalyzer log data with SSL encryption. Valid values: `default`, `high`, `low`, `disable`.
* `encrypt` - Whether to send FortiAnalyzer log data in IPsec tunnel. Valid values: `disable`, `enable`.
* `source_ip` - Source IP address of the FortiAnalyzer.
* `override` - Override FortiAnalyzer settings or use the global settings.
* `hmac_algorithm` - FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.
* `status` - Enable/disable FortiAnalyzer. Valid values: `enable`, `disable`.
* `upload_day` - Days of week(month) to upload logs.
* `ips_archive` - Whether to enable IPS packet archive. Valid values: `enable`, `disable`.
* `mgmt_name` - Hidden management name of FortiAnalyzer.
* `fdp_interface` - Interface for FortiAnalyzer auto-discovery.
* `__change_ip` - Hidden attribute.
* `localid` - Local id for IPsec tunnel to FortiAnalyzer.
* `max_buffer_size` - Maximum buffer size, in MBytes, 0--1024, 0=disabled.
* `server` - IP address of the remote FortiAnalyzer.
* `conn_timeout` - FortiAnalyzer connection time-out in seconds (for status and log buffer).
* `psksecret` - Pre-shared key for IPsec tunnel to FortiAnalyzer.
* `upload_time` - Time to upload logs [hh:mm].


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogFortianalyzer Setting can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_logfortianalyzer_setting.labelname LogFortianalyzerSetting

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_logfortianalyzer_setting.labelname LogFortianalyzerSetting
$ unset "FORTISWITCH_IMPORT_TABLE"
```
