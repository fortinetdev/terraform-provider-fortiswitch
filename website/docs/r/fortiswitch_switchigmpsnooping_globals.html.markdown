---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchigmpsnooping_globals"
description: |-
  Configure igmp-snooping on Switch.
---

# fortiswitch_switchigmpsnooping_globals
Configure igmp-snooping on Switch.

## Example Usage

```hcl
resource "fortiswitch_switchigmpsnooping_globals" "name" {
       aging_time = "200"
}
```

## Argument Reference

The following arguments are supported:

* `query_interval` - Max number of seconds after which IGMP query will be sent.
* `proxy_report_interval` - Unsolicited report interval in seconds.
* `query_max_response_timeout` - Max time a host waits before responses to general query message (in millisecond).
* `leave_response_timeout` - Switch waits after sending group specific query in response to leave message.
* `aging_time` - Max number of seconds to retain a multicast snooping entry for which no packets have been seen.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchIgmpSnooping Globals can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchigmpsnooping_globals.labelname SwitchIgmpSnoopingGlobals

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchigmpsnooping_globals.labelname SwitchIgmpSnoopingGlobals
$ unset "FORTISWITCH_IMPORT_TABLE"
```
