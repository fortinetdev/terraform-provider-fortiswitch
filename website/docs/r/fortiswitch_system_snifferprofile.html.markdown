---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_snifferprofile"
description: |-
  Show packet sniffer configuration.
---

# fortiswitch_system_snifferprofile
Show packet sniffer configuration.

## Example Usage

```hcl
resource "fortiswitch_system_snifferprofile" "name" {
        filter = "test"
        max_pkt_count = "4"
        max_pkt_len = "64"
        profile_name = "2"
        system_interface = "internal"
}
```

## Argument Reference

The following arguments are supported:

* `switch_interface` - Select switch-interface name on which packets are to be captured.
* `max_pkt_count` - Maximum number of packet to be captured on the interface  (1-INT_MAX, default=4000).
* `system_interface` - Select system-interface name on which packets are to be captured.
* `filter` - Flexible logical filters for sniffer (or "none")
                    For example:  To print udp 1812 traffic between forti1 and either forti2 or forti3
                    'udp and port 1812 and host forti1 and \( forti2 or forti3 \).'
* `profile_name` - Give name to the sniffer profile.
* `max_pkt_len` - Maximum packet length to be captured on the interface  (64-1534, default=128).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{profile_name}}.

## Import

System SnifferProfile can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_snifferprofile.labelname {{profile_name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_snifferprofile.labelname {{profile_name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
