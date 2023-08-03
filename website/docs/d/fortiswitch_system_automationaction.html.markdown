---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_automationaction"
description: |-
  Get information on an fortiswitch system automationaction.
---

# Data Source: fortiswitch_system_automationaction
Use this data source to get information on an fortiswitch system automationaction

## Argument Reference

* `name` - (Required) Specify the name of the desired system automationaction.

## Attribute Reference

The following attributes are exported:

* `aws_api_key` - AWS API Gateway API key.
* `alicloud_account_id` - AliCloud account ID.
* `protocol` - Request protocol.
* `alicloud_function` - AliCloud function name.
* `aws_domain` - AWS domain.
* `aws_api_id` - AWS API Gateway ID.
* `email_to` - Email addresses. The structure of `email_to` block is documented below.
* `azure_api_key` - Azure function API key.
* `port` - Protocol port.
* `minimum_interval` - Limit execution to no more than once in this interval (in seconds).
* `gcp_function_domain` - Google Cloud function domain.
* `email_body` - Email body.
* `http_body` - Request body (if necessary). Should be serialized json string.
* `azure_domain` - Azure function domain.
* `alicloud_version` - AliCloud version.
* `accprofile` - Access profile for CLI script action to access FortiSwitch features.
* `headers` - Request headers. The structure of `headers` block is documented below.
* `script` - CLI script.
* `email_subject` - Email subject.
* `email_from` - Email sender name.
* `aws_api_path` - AWS API Gateway path.
* `method` - Request method (POST, PUT, GET, PATCH or DELETE).
* `action_type` - Action type.
* `alicloud_function_domain` - AliCloud function domain.
* `azure_app` - Azure function application name.
* `gcp_function` - Google Cloud function name.
* `alicloud_access_key_id` - AliCloud AccessKey ID.
* `gcp_project` - Google Cloud Platform project name.
* `alicloud_access_key_secret` - AliCloud AccessKey secret.
* `name` - Name.
* `aws_region` - AWS region.
* `azure_function` - Azure function name.
* `aws_api_stage` - AWS API Gateway deployment stage name.
* `alicloud_region` - AliCloud region.
* `uri` - Request API URI.
* `azure_function_authorization` - Azure function authorization level.
* `snmp_trap` - SNMP trap.
* `gcp_function_region` - Google Cloud function region.
* `alicloud_service` - AliCloud service name.
* `alicloud_function_authorization` - AliCloud function authorization type.

The `email_to` block contains:

* `name` - Email address.

The `headers` block contains:

* `header` - Request header.

