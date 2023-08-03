---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_fortianalyzer"
description: |-
  Setting for FortiAnalyzer.
---

# fortiswitch_system_fortianalyzer
Setting for FortiAnalyzer.

## Example Usage

```hcl
resource "fortiswitch_system_fortianalyzer" "name" {
        address_mode = "static"
        conn_timeout = "5"
        encrypt = "disable"
        fdp_device = "port5"
        fdp_interface = "port5"
        localid = "3"
        mgmt_name = "test"
        psksecret = "123"
        server = "192.168.100.40"
        status = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable FortiAnalyzer. Valid values: `enable`, `disable`.
* `__change_ip` - Hidden attribute.
* `encrypt` - Whether to send FortiAnalyzer log data in IPsec tunnel. Valid values: `disable`, `enable`.
* `localid` - Local id for IPsec tunnel to FortiAnalyzer.
* `fdp_device` - Serial number of FortiAnalyzer to connect to.
* `conn_timeout` - FortiAnalyzer connection time-out in seconds (for status and log buffer).
* `psksecret` - Pre-shared key for IPsec tunnel to FortiAnalyzer.
* `mgmt_name` - Hidden management name of FortiAnalyzer.
* `fdp_interface` - Interface for FortiAnalyzer auto-discovery.
* `server` - IP address of the remote FortiAnalyzer.
* `address_mode` - FortiAnalyzer IP addressing mode. Valid values: `static`, `auto-discovery`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fortianalyzer can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_fortianalyzer.labelname SystemFortianalyzer

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_fortianalyzer.labelname SystemFortianalyzer
$ unset "FORTISWITCH_IMPORT_TABLE"
```
