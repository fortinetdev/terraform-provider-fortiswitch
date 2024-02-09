---
layout: "fortiswitch"
page_title: "Provider: FortiSwitch"
sidebar_current: "docs-fortiswitch-index"
description: |-
  The FortiSwitch provider interacts with FortiSwitch.
---

# FortiSwitch Provider

The FortiSwitch provider is used to interact with the resources supported by FortiSwitch. We need to configure the provider with the proper credentials before it can be used.

## Configuration for FortiSwitch

### Example Usage

```hcl
# Configure the FortiSwitch Provider
provider "fortiswitch" {
  hostname     = "192.168.52.177"
  username = "admin"
  password = "fortinet"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"
}

# Create a system interface entry
resource "fortiswitch_system_interface" "interface" {
        description = "interface"
        ip = "192.168.22.23/24"
        mode = "static"
        name = "test"
}
```

If it is used for testing, you can set `insecure` to "true" and unset `cabundlefile` to quickly set the provider up, for example:

```hcl
provider "fortiswitch" {
  hostname = "192.168.52.177"
  username = "admin"
  password = "fortinet"
  insecure = "true"
}
```

Please refer to the Argument Reference below for more help on `insecure` and `cabundlefile`.

### Authentication

The FortiSwitch provider offers a way of providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables

#### Static credentials

Usage:

```hcl
provider "fortiswitch" {
  hostname     = "192.168.52.177"
  username = "admin"
  password = "fortinet"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"
}
```

#### Environment variables

You can provide your credentials via the `FORTISWITCH_INSECURE` and `FORTISWITCH_CA_CABUNDLE` environment variables. Note that setting your FortiSwitch credentials using static credentials variables will override the environment variables.

Usage:

```shell
$ export "FORTISWITCH_ACCESS_HOSTNAME"="192.168.52.177"
$ export "FORTISWITCH_INSECURE"="false"
$ export "FORTISWITCH_ACCESS_USERNAME"="admin"
$ export "FORTISWITCH_ACCESS_PASSWORD"="fortinet"
$ export "FORTISWITCH_CA_CABUNDLE"="/path/yourCA.crt"
```

### Argument Reference

The following arguments are supported:

* `hostname` - (Optional) The hostname or IP address of FortiSwitch unit. It must be provided, but it can also be sourced from the `FORTISWITCH_ACCESS_HOSTNAME` environment variable.

* `username` - (Optional) The username of FortiSwitch unit. It must be provided, but it can also be sourced from the `FORTISWITCH_ACCESS_USERNAME` environment variable.

* `password` - (Optional) The password of FortiSwitch unit. It must be provided, but it can also be sourced from the `FORTISWITCH_ACCESS_PASSWORD` environment variable.

* `insecure` - (Optional) Control whether the Provider to perform insecure SSL requests. If omitted, the `FORTISWITCH_INSECURE` environment variable is used. If neither is set, default value is `false`.

* `cabundlefile` - (Optional) The path of a custom CA bundle file. You can specify a path to the file, or you can specify it by the `FORTISWITCH_CA_CABUNDLE` environment variable.

* `cabundlecontent` - (Optional) The content of a custom CA bundle file. Note: `cabundlefile` and `cabundlecontent` cannot exist at the same time! Please only configure one of them.

* `http_proxy` - (Optional) HTTP proxy address. You can also specify it by the environment variable `HTTPS_PROXY` or `HTTP_PROXY`.

## Release
Check out the FortiSwitch provider release notes and additional information from: [the FortiSwitch provider releases](https://github.com/fortinetdev/terraform-provider-fortiswitch/releases).

## Versioning

The provider can cover FortiSwitch 7.0, 7.2 and 7.4 versions, the configuration of all parameters should be based on the relevant FortiSwitch version manual.
