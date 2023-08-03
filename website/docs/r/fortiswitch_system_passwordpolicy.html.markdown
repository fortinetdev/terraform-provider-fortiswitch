---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_passwordpolicy"
description: |-
  Config password policy.
---

# fortiswitch_system_passwordpolicy
Config password policy.

## Example Usage

```hcl
resource "fortiswitch_system_passwordpolicy" "psw" {
  apply_to              = "admin-password"
  change_4_characters   = "disable"
  expire_day            = 90
  expire_status         = "disable"
  min_lower_case_letter = 0
  min_non_alphanumeric  = 0
  min_number            = 0
  min_upper_case_letter = 0
  minimum_length        = 8
  status                = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Password policy status. Valid values: `enable`.
* `minimum_length` - Set password's minimum length.
* `change_4_characters` - Enable/disable changing at least 4 characters for new password. Valid values: `enable`, `disable`.
* `min_number` - Min number of numeric characters in password.
* `min_non_alphanumeric` - Min number of non-alpha characters in password.
* `min_lower_case_letter` - Min number of lowercase characters in password.
* `min_upper_case_letter` - Min number of uppercase characters in password.
* `expire_status` - Enable/disable the password expiration. Valid values: `enable`, `disable`.
* `expire_day` - Set number of days after which admin users' password will expire.
* `apply_to` - Apply password policy to. Valid values: `admin-password`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System PasswordPolicy can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_passwordpolicy.labelname SystemPasswordPolicy

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_passwordpolicy.labelname SystemPasswordPolicy
$ unset "FORTISWITCH_IMPORT_TABLE"
```
