---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchacl_policer"
description: |-
  Policer configuration.
---

# fortiswitch_switchacl_policer
Policer configuration.

## Example Usage

```hcl
resource "fortiswitch_switchacl_policer" "name" {
        description = "test"
        guaranteed_bandwidth = "4"
        guaranteed_burst = "5"
        fswid =  "6"
        maximum_burst = "7"
        type = "ingress"
}
```

## Argument Reference

The following arguments are supported:

* `maximum_burst` - Maximum burst size in bytes (max value = 4294967295).
* `guaranteed_burst` - Guaranteed burst size in bytes (max value = 4294967295).
* `description` - Description of the policer.
* `guaranteed_bandwidth` - Guaranteed bandwidth in kbps (max value = 524287000).
* `type` - Configure type of policer(ingress/egress). Valid values: `ingress`, `egress`.
* `fswid` - Policer id valid range 1-2048.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

SwitchAcl Policer can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchacl_policer.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchacl_policer.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
