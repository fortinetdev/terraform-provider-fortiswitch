---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_resourcelimits"
description: |-
  Resource limits configuration.
---

# fortiswitch_system_resourcelimits
Resource limits configuration.

## Example Usage

```hcl
resource "fortiswitch_system_resourcelimits" "name" {
        custom_service = "3"
        dialup_tunnel = "4"
        firewall_address = "5"
        firewall_addrgrp = "6"
        firewall_policy = "7"
        ipsec_phase1 = "8"
        ipsec_phase2 = "9"
        log_disk_quota = "10"
        onetime_schedule = "11"
        proxy = "12"
        recurring_schedule = "13"
        service_group = "14"
        session = "15"
        user = "16"
        user_group = "17"
}
```

## Argument Reference

The following arguments are supported:

* `log_disk_quota` - Log disk quota in MB.
* `user_group` - Maximum number of user groups.
* `onetime_schedule` - Maximum number of firewall one-time schedules.
* `recurring_schedule` - Maximum number of firewall recurring schedules.
* `firewall_policy` - Maximum number of firewall policies.
* `service_group` - Maximum number of firewall service groups.
* `custom_service` - Maximum number of firewall custom services.
* `ipsec_phase1` - Maximum number of vpn ipsec phase1 tunnels.
* `ipsec_phase2` - Maximum number of vpn ipsec phase2 tunnels.
* `session` - Maximum number of sessions.
* `firewall_addrgrp` - Maximum number of firewall address groups.
* `firewall_address` - Maximum number of firewall addresses.
* `proxy` - Maximum number of concurrent explicit proxy users.
* `dialup_tunnel` - Maximum number of dial-up tunnels.
* `user` - Maximum number of local users.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System ResourceLimits can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_resourcelimits.labelname SystemResourceLimits

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_resourcelimits.labelname SystemResourceLimits
$ unset "FORTISWITCH_IMPORT_TABLE"
```
