---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_accesslist"
description: |-
  Access list configuration.
---

# fortiswitch_router_accesslist
Access list configuration.


## Example Usage

```hcl
resource "fortiswitch_router_accesslist" "trname" {
  comments = "test accesslist"
  name     = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `comments` - Comment.
* `rule` - Rule. The structure of `rule` block is documented below.

The `rule` block supports:

* `id` - Rule ID.
* `action` - Permit or deny this IP address and netmask prefix.
* `prefix` - IPv4 prefix to define regular filter criteria, such as "any" or subnets.
* `wildcard` - Wildcard to define Cisco-style wildcard filter criteria.
* `exact_match` - Enable/disable exact match.
* `flags` - Flags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router AccessList can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_router_accesslist.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="true"
$ terraform import fortiswitch_router_accesslist.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```

