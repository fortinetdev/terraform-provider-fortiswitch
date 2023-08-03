---
subcategory: "FortiSwitch Switch"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_switchaclservice_custom"
description: |-
  Custom service configuration.
---

# fortiswitch_switchaclservice_custom
Custom service configuration.

## Example Usage

```hcl
resource "fortiswitch_switchaclservice_custom" "name" {
        color = "3"
        comment = "Comments."
        icmpcode = "5"
        icmptype = "6"
        name = "test4"
        protocol = "TCP/UDP/SCTP"
        protocol_number = "9"
        sctp_portrange = "2-3:4-6"
        tcp_portrange = "2-3:4-5"
        udp_portrange = "1-1:1-1"
}
```

## Argument Reference

The following arguments are supported:

* `comment` - Comments.
* `protocol_number` - IP protocol number.
* `protocol` - Protocol type. Valid values: `TCP/UDP/SCTP`, `ICMP`, `IP`.
* `name` - Custom service name.
* `color` - Set GUI icon color.
* `icmptype` - ICMP type.
* `sctp_portrange` - Multiple sctp port ranges.
* `udp_portrange` - Multiple udp port ranges.
* `icmpcode` - ICMP code.
* `tcp_portrange` - Multiple tcp port ranges.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchAclService Custom can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_switchaclservice_custom.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_switchaclservice_custom.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
