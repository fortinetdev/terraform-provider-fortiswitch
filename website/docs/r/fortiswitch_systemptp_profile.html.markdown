---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_systemptp_profile"
description: |-
  PTP policy configuration.
---

# fortiswitch_systemptp_profile
PTP policy configuration. Applies to FortiSwitch Version `>= 7.2.5`.

## Argument Reference

The following arguments are supported:

* `priority1` - PTP priority1 (0-255)
* `domain` - PTP domain (0-255)
* `announce_interval` - Announce interval. Valid values: `0.25sec`, `0.5sec`, `1sec`, `2sec`, `4sec`.
* `name` - Profile name.
* `min_delay_req_interval` - Min Delay Request interval. Valid values: `0.25sec`, `0.5sec`, `1sec`, `2sec`, `4sec`.
* `pdelay_req_interval` - PDelay Request interval. Valid values: `0.25sec`, `0.5sec`, `1sec`, `2sec`, `4sec`.
* `sync_interval` - Sync interval. Valid values: `0.25sec`, `0.5sec`, `1sec`, `2sec`, `4sec`.
* `priority2` - PTP priority2 (0-255)
* `mode` - Select PTP mode. Valid values: `transparent-e2e`.
* `ptp_profile` - Select PTP profile.
* `announce_timeout` - PTP Announce timeout (2-10)
* `transport` - Select PTP transport. Valid values: `l2-mcast`.
* `description` - Description.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemPtp Profile can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_systemptp_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_systemptp_profile.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
