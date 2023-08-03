---
subcategory: "FortiSwitch Log"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_logsyslogd_overridesetting"
description: |-
  Settings for remote syslog server.
---

# fortiswitch_logsyslogd_overridesetting
Settings for remote syslog server.

## Argument Reference

The following arguments are supported:

* `status` - Whether to enable remote syslog log. Valid values: `enable`, `disable`.
* `enc_algorithm` - Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
* `certificate` - Certificate used to communicate with Syslog server.
* `facility` - Which facility for remote syslog. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
* `source_ip` - Source IP address of the syslog.
* `server` - Address of the remote syslog server.
* `mode` - Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
* `override` - Override syslog settings or use the global settings. Valid values: `enable`, `disable`.
* `csv` - Whether to enable CSV. Valid values: `enable`, `disable`.
* `port` - Port that the server listens at.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogSyslogd OverrideSetting can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_logsyslogd_overridesetting.labelname LogSyslogdOverrideSetting

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_logsyslogd_overridesetting.labelname LogSyslogdOverrideSetting
$ unset "FORTISWITCH_IMPORT_TABLE"
```
