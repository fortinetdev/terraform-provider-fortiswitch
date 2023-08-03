---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_logmemory_filter"
description: |-
  Filters for memory buffer.
---

# fortiswitch_logmemory_filter
Filters for memory buffer.

## Example Usage

```hcl
resource "fortiswitch_logmemory_filter" "lmf" {
        admin = "enable"
        auth = "enable"
        cpu_memory_usage = "enable"
        dhcp = "enable"
        event = "enable"
        ha = "enable"
        pattern = "enable"
        severity = "emergency"
        system = "enable"
        wireless_activity = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `severity` - The least severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `admin` - Whether to log admin login/logout messages. Valid values: `enable`, `disable`.
* `pattern` - Whether to log pattern update messages. Valid values: `enable`, `disable`.
* `wireless_activity` - Whether to log wireless activity. Valid values: `enable`, `disable`.
* `system` - Enable/disable logging of system activity messages. Valid values: `enable`, `disable`.
* `auth` - Whether to log firewall authentication messages. Valid values: `enable`, `disable`.
* `cpu_memory_usage` - Whether to log CPU & memory usage every 5 minutes. Valid values: `enable`, `disable`.
* `override` - Override FortiAnalyzer setting or use the global setting.
* `dhcp` - Whether to log DHCP service messages. Valid values: `enable`, `disable`.
* `ha` - Whether to log HA activity messages. Valid values: `enable`, `disable`.
* `event` - Enable/disable logging of event messages. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogMemory Filter can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_logmemory_filter.labelname LogMemoryFilter

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_logmemory_filter.labelname LogMemoryFilter
$ unset "FORTISWITCH_IMPORT_TABLE"
```
