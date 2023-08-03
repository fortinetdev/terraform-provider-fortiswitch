---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_logfortianalyzer_overridesetting"
description: |-
  Setting for FortiAnalyzer.
---

# fortiswitch_logfortianalyzer_overridesetting
Setting for FortiAnalyzer.

## Example Usage

```hcl
resource "fortiswitch_logfortianalyzer_overridesetting" "name" {
        address_mode = "static"
        buffer_max_send = "21"
        conn_timeout = "6"
        enc_algorithm = "high"
        encrypt = "disable"
        fdp_device = "test"
        fdp_interface = "port4"
        hmac_algorithm = "sha256"
        ips_archive = "enable"
        localid = "2"
        max_buffer_size = "14"
        mgmt_name = "admin"
        override = "enable"
        psksecret = "test"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable FortiAnalyzer. Valid values: `enable`, `disable`.
* `enc_algorithm` - Whether to send FortiAnalyzer log data with SSL encryption. Valid values: `default`, `high`, `low`, `disable`.
* `hmac_algorithm` - FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.
* `encrypt` - Whether to send FortiAnalyzer log data in IPsec tunnel. Valid values: `disable`, `enable`.
* `localid` - Local id for IPsec tunnel to FortiAnalyzer.
* `__change_ip` - Hidden attribute.
* `ips_archive` - Whether to enable IPS packet archive. Valid values: `enable`, `disable`.
* `source_ip` - Source IP address of the FortiAnalyzer.
* `server` - IP address of the remote FortiAnalyzer.
* `conn_timeout` - FortiAnalyzer connection time-out in seconds (for status and log buffer).
* `max_buffer_size` - Maximum buffer size, in MBytes, 0--1024, 0=disabled.
* `psksecret` - Pre-shared key for IPsec tunnel to FortiAnalyzer.
* `mgmt_name` - Hidden management name of FortiAnalyzer.
* `fdp_interface` - Interface for FortiAnalyzer auto-discovery.
* `override` - Override FortiAnalyzer settings or use the global settings. Valid values: `enable`, `disable`.
* `buffer_max_send` - Maximum log transmission rate for buffered logs.
* `fdp_device` - Serial number of FortiAnalyzer to connect to.
* `address_mode` - FortiAnalyzer IP addressing mode. Valid values: `static`, `auto-discovery`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogFortianalyzer OverrideSetting can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_logfortianalyzer_overridesetting.labelname LogFortianalyzerOverrideSetting

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_logfortianalyzer_overridesetting.labelname LogFortianalyzerOverrideSetting
$ unset "FORTISWITCH_IMPORT_TABLE"
```
