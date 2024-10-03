---
subcategory: ""
layout: "fortiswitch"
page_title: "Concurrent operations on FortiSwitch"
description: |-
  Concurrent operations on FortiSwitch.
---

# Concurrent operations on FortiSwitch

FortiSwitch has a limit on the number of concurrent administrative sessions it can handle at once. If too many sessions are opened simultaneously, it may hit this limit, causing problems.

In Terraform, the `-parallelism` option controls the number of concurrent operations Terraform can perform during apply, plan, or destroy commands. When you set -parallelism=1, it ensures that Terraform will only perform one operation at a time, which can be useful when dealing with resources that may conflict when created in parallel, or if there are dependency issues. The default parallelism value in Terraform is 10, while the maximum concurrent administrative sessions for FortiSwitch is 3, therefore, set `-parallelism=1`, `-parallelism=2` or `-parallelism=3` at a time to avoid hitting the FortiSwitch's concurrent session limit, leading to a smoother operation.