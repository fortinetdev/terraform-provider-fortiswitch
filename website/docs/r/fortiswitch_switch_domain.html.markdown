---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switch_domain"
description: |-
  Switch forwarding domains.
---

# fortiswitch_switch_domain
Switch forwarding domains.

## Example Usage

```hcl
resource "fortiswitch_switch_domain" "name" {
        ha_l2_clear_on_role_change = "monitor_ports"
        inter_front_panel_traffic = "enable"
        name = "default_name_6"
        priority = "7"
        vcluster_id = "8"
}
```

## Argument Reference

The following arguments are supported:

* `inter_front_panel_traffic` - Traffic flow between front panel ports (eg. f1<->f2). Valid values: `enable`, `disable`.
* `name` - Domain name.
* `priority` - Priority value(0-255).
* `ha_block` - Select port types to be blocked if domain becomes ha slave. Valid values: `monitor-ports`, `blade-ports`, `misc-ports`.
* `ha_l2_clear_on_role_change` - Select port types have their L2 tables cleared when changing HA roles. Valid values: `monitor-ports`, `blade-ports`, `misc-ports`.
* `vcluster_id` - HA virtual cluster id(1-255).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Switch Domain can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switch_domain.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switch_domain.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
