---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_vdomproperty"
description: |-
  Vdom-property configuration.
---

# fortiswitch_system_vdomproperty
Vdom-property configuration.

## Example Usage

```hcl
resource "fortiswitch_system_vdomproperty" "name" {
        custom_service = "test"
        description = "test"
        dialup_tunnel = "test"
        firewall_address = "3.4.5.6"
        firewall_policy = "test"
        name = "test"
}
```

## Argument Reference

The following arguments are supported:

* `user_group` - Maximum number [guaranteed number] of user groups.
* `log_disk_quota` - Log disk quota in MB.
* `name` - Vdom name.
* `onetime_schedule` - Maximum number [guaranteed number] of firewall one-time schedules.
* `user` - Maximum number [guaranteed number] of local users.
* `recurring_schedule` - Maximum number [guaranteed number] of firewall recurring schedules.
* `firewall_policy` - Maximum number [guaranteed number] of firewall policies.
* `service_group` - Maximum number [guaranteed number] of firewall service groups.
* `custom_service` - Maximum number [guaranteed number] of firewall custom services.
* `ipsec_phase1` - Maximum number [guaranteed number] of vpn ipsec phase1 tunnels.
* `ipsec_phase2` - Maximum number [guaranteed number] of vpn ipsec phase2 tunnels.
* `session` - Maximum number [guaranteed number] of sessions.
* `firewall_addrgrp` - Maximum number [guaranteed number] of firewall address groups.
* `firewall_address` - Maximum number [guaranteed number] of firewall addresses.
* `proxy` - Maximum number [guaranteed number] of concurrent proxy users.
* `dialup_tunnel` - Maximum number [guaranteed number] of dial-up tunnels.
* `description` - Description.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System VdomProperty can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_vdomproperty.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_vdomproperty.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
