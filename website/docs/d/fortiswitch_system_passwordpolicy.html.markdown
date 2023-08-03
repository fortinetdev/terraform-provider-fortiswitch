---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_passwordpolicy"
description: |-
  Get information on fortiswitch system passwordpolicy.
---

# Data Source: fortiswitch_system_passwordpolicy
Use this data source to get information on fortiswitch system passwordpolicy

## Argument Reference



## Attribute Reference

The following attributes are exported:

* `status` - Password policy status.
* `minimum_length` - Set password's minimum length.
* `change_4_characters` - Enable/disable changing at least 4 characters for new password.
* `min_number` - Min number of numeric characters in password.
* `min_non_alphanumeric` - Min number of non-alpha characters in password.
* `min_lower_case_letter` - Min number of lowercase characters in password.
* `min_upper_case_letter` - Min number of uppercase characters in password.
* `expire_status` - Enable/disable the password expiration.
* `expire_day` - Set number of days after which admin users' password will expire.
* `apply_to` - Apply password policy to.

