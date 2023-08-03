---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_proxyarp"
description: |-
  Configure proxy-arp.
---

# fortiswitch_system_proxyarp
Configure proxy-arp.

## Example Usage

```hcl
resource "fortiswitch_system_proxyarp" "name" {
        fswid =  "3"
        interface = "port9"
        ip = "6.4.3.65"
}
```

## Argument Reference

The following arguments are supported:

* `interface` - Interface acting proxy-arp.
* `ip` - IP address to be proxied.
* `fswid` - Unique integer ID of the entry.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fswid}}.

## Import

System ProxyArp can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_proxyarp.labelname {{fswid}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_proxyarp.labelname {{fswid}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
