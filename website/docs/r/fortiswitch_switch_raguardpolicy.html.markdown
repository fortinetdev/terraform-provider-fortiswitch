---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_raguardpolicy"
description: |-
  IPV6 RA Guard policy.
---

# fortiswitch_switch_raguardpolicy
IPV6 RA Guard policy.

## Argument Reference

The following arguments are supported:

* `max_router_preference` - Max-router-preference. Valid values: `high`, `medium`, `low`.
* `min_hop_limit` - Min hop limit.
* `max_hop_limit` - Max hop limit.
* `managed_flag` - Managed flag. Valid values: `on`, `off`.
* `match_src_addr` - Match src ip address permitted by access-list.
* `device_role` - Device-role. Valid values: `host`, `router`.
* `other_flag` - Other flag. Valid values: `on`, `off`.
* `match_prefix` - Match prefix permitted by Prefix-list.
* `name` - RA Guard policy name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Switch RaguardPolicy can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_raguardpolicy.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_raguardpolicy.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
