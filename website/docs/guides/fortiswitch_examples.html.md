---
subcategory: ""
layout: "fortiswitch"
page_title: "Examples for fortiswitch modules"
description: |-
  Getting started with the terraform modules.
---

# Examples

```shell
terraform {
	required_providers {
		fortiswitch = {
			source = "fortinetdev/fortiswitch"
    }
  }
}

provider "fortiswitch" {
  hostname = "192.168.11.31"
  username = "admin"
  password = "fortinet"
  insecure = true
}

# Creat a system interface object
resource "fortiswitch_system_interface" "interface" {
        description = "terraform test"
        ip = "192.168.22.23/24"
        mode = "static"
        name = "test"
        # ...
}
```