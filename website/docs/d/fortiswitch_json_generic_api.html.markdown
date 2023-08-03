---
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_json_generic_api"
subcategory: "FortiSwitch Generic"
description: |-
  Provides a FortiAPI Generic Interface data source.
---

# Data Source: fortiswitch_json_generic_api
Provides a FortiAPI Generic Interface data source.

## Example Usage
```hcl
resource "fortiswitch_json_generic_api" "test" {
  path   = "/api/v2/cmdb/system/admin"
  method = "GET"
  json   = ""
}

output "response" {
  value = "${fortiswitch_json_generic_api.test.response}"
}
```

## Argument Reference
The following arguments are supported:

* `path` - (required) FortiAPI URL path.
* `specialparams` - URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `path` - FortiAPI URL path.
* `specialparams` - URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
* `response` - FortiAPI returns results.

-> API parameter reference: https://fndn.fortinet.net/index.php?/fortiapi/1-fortiswitch/ , if you do not have an account, please contact the Fortinet SE serving you.