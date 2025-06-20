---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_debug"
description: |-
  Application and CLI debug values to set at startup and retain over reboot.
---

# fortiswitch_system_debug
Application and CLI debug values to set at startup and retain over reboot. Applies to FortiSwitch Version `>= 7.6.1`.

## Argument Reference

The following arguments are supported:

* `cli` - CMDB/CLI debug.
* `radvd` - router adv daemon
* `raguard` - raguard daemon
* `miglogd` - Log daemon.
* `dhcp6c` - DHCPv6 client.
* `eap_proxy` - EAP proxy daemon.
* `wpa_supp` - MKA/Fortilink macsec daemon.
* `macsec_srv` - MKA/Fortilink macsec cak server daemon.
* `dhcps` - DHCP server.
* `fnbamd` - Fortigate non-blocking auth daemon.
* `dhcprelay` - DHCP relay daemon.
* `snmpd` - Simple Network Managment Protocol (SNMP) daemon.
* `dnsproxy` - DNS proxy module.
* `sflowd` - sFlow collection and export daemon.
* `dhcpc` - DHCP client module.
* `router_launcher` - Routing system launcher daemon.
* `sshd` - Secure Sockets Shell(SSH) daemon.
* `ctrld` - Switch general control daemon.
* `stpd` - Spanning Tree (STP) daemon.
* `trunkd` - Link Aggregation Control Protocol (LACP) daemon.
* `lacpd` - Link Aggregation Control Protocol (LACP) debug.
* `lldpmedd` - Link Layer Discovery Protocol (LLDP) daemon.
* `ipconflictd` - IP Conflictd detection daemon.
* `httpsd` - HTTP/HTTPS daemon.
* `link_monitor` - Link monitor daemon.
* `libswitchd` - Switch library daemon.
* `switch_launcher` - Switching system launcher daemon.
* `alertd` - Monitor and Alert daemon.
* `l2d` - L2 daemon responsible to have core logic to assist L2 feature like MCLAG.
* `l2dbg` - Background daemon responsible to assist in any heavy HW related operation needed by L2d.
* `nwmcfgd` - Network monitor daemon responsible for handling configuration.
* `nwmonitord` - Network monitor daemon responsible for packet handling and parsing.
* `portspeedd` - Port speed daemon.
* `l3` - L3 debug.
* `mcast_snooping` - Multicast Snooping debug.
* `dmid` - DMI daemon debug.
* `scep` - SCEP
* `cu_swtpd` - Switch-controller CAPWAP daemon.
* `fortilinkd` - FortiLink daemon.
* `flcmdd` - FortiLink command daemon.
* `gvrpd` - GVRP daemon.
* `flan_mgr` - FortiLAN Cloud Manager daemon.
* `rsyslogd` - Remote SYSLOG daemon.
* `vrrpd` - Virtual Router Redundancy Protocol (VRRP) daemon.
* `fpmd` - FPMD HW routing daemon.
* `ospfd` - Open Shortest Path First (OSPF) routing daemon.
* `ospf6d` - Open Shortest Path First (OSPFv6) routing daemon.
* `pbrd` - Policy Based Routing (PBR) routing daemon.
* `isisd` - Isis daemon.
* `ripd` - Routing Information Protocol (RIP) daemon.
* `ripngd` - Routing Information Protocol NG (RIPNG) daemon.
* `bgpd` - Bgp daemon.
* `zebra` - Zebra routing daemon.
* `bfdd` - Bidirectional Forwarding Detection (BFD) daemon.
* `staticd` - Static route daemon.
* `pimd` - Protocol Independent Multicast (PIM) daemon.
* `ntpd` - Network Time Protocol (NTP) daemon.
* `wiredap` - Wired AP (802.1X port-based auth) daemon.
* `ip6addrd` - IPv6 address utility.
* `gratarp` - IP Conflict Gratuitious ARP utility.
* `radius_das` - Radius CoA daemon.
* `gui` - GUI service.
* `statsd` - Stats collection daemon.
* `flow_export` - Flow-export debug.
* `erspan_auto_mgr` - ERSPAN-auto mode configuration resolution daemon.
* `autod` - Automation Stitches.
* `email_server` - Email server.
* `auto_script` - Auto script.
* `apache` - Apache.
* `delayclid` - Delay CLI daemon.
* `rvi_daemon` - Multicast Snooping debug.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Debug can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_debug.labelname SystemDebug

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_debug.labelname SystemDebug
$ unset "FORTISWITCH_IMPORT_TABLE"
```
