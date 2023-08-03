---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_fortimanager"
description: |-
  FortiManagerconfiguration.
---

# fortiswitch_system_fortimanager
FortiManagerconfiguration.

By design considerations, the feature is using the fortiswitch_system_centralmanagement resource as documented below.

## Example
```hcl
resource "fortiswitch_system_centralmanagement" "trname" {
  allow_monitor                 = "enable"
  allow_push_configuration      = "enable"
  allow_push_firmware           = "enable"
  allow_remote_firmware_upgrade = "enable"
  enc_algorithm                 = "high"
  fmg                           = "\"192.168.52.177\""
  include_default_servers       = "enable"
  mode                          = "normal"
  type                          = "fortimanager"
  vdom                          = "root"
}
```

