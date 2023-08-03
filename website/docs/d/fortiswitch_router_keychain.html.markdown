---
subcategory: "FortiSwitch Router"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_router_keychain"
description: |-
  Get information on an fortiswitch router keychain.
---

# Data Source: fortiswitch_router_keychain
Use this data source to get information on an fortiswitch router keychain

## Argument Reference

* `name` - (Required) Specify the name of the desired router keychain.

## Attribute Reference

The following attributes are exported:

* `name` - Key-chain name.
* `key` - Key. The structure of `key` block is documented below.

The `key` block contains:

* `accept_lifetime` - Lifetime of received authentication key.
* `key_string` - Password for the key.
* `id` - Key id <0-2147483647>.
* `send_lifetime` - Lifetime of sent authentication key.

