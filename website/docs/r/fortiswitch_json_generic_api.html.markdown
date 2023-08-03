---
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_json_generic_api"
sidebar_current: "docs-fortiswitch-resource-json-generic-api"
subcategory: "FortiSwitch Generic"
description: |-
  FortiAPI Generic Interface.
---

# fortiswitch_json_generic_api
FortiAPI Generic Interface.

## Example Usage
```hcl
resource "fortiswitch_json_generic_api" "test" {
  path   = "/api/v2/cmdb/system/admin"
  method = "POST"
  json   = <<EOF
{
  "accprofile": "prof_admin",
  "accprofile_override": "disable",
  "comments": "test_new",
  "name": "api_admin_6",
  "password": "1234mn90812"
}
EOF
}

output "response" {
  value = "${fortiswitch_json_generic_api.test.response}"
}
```

## Argument Reference
The following arguments are supported:

* `path` - (required) FortiAPI URL path.
* `method` - (required) HTTP method.
* `specialparams` - URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path.
* `force_recreate` - The argument is optional, if it is set, when its value changes, the resource will be re-created. It is usually used when the return value needs to be forced to update.
* `json` - Body data in JSON format.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `path` - FortiAPI URL path.
* `method` - HTTP method.
* `specialparams` - URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path.
* `force_recreate` - The argument is optional, if it is set, when its value changes, the resource will be re-created.
* `json` - Body data in JSON format.
* `response` - FortiAPI returns results.
