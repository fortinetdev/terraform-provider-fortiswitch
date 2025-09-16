## 1.1.7 (Unreleased)

## 1.1.6 (Sep 10, 2025)

IMPROVEMENTS:

* Support new version 7.6.3 and 7.6.4.
* Upgrade the Go version and all associated dependencies.

## 1.1.5 (June 20, 2025)

BUGFIXES:

* Fixed a crash that could occur when provider information was missing from the configuration.
* Now returns an error message when the resource does not exist.

IMPROVEMENTS:

* Support new version 7.6.1 and 7.6.2.

## 1.1.4 (November 21, 2024)

BUGFIXES:

* Fix the error in `fortiswitch_awitch_interface`.

## 1.1.3 (September 30, 2024)

BUGFIXES:

* Fix the provider crash issue in `fortiswitch_switch_ipmacbinding`, `fortiswitch_system_proxyarp` and `fortiswitch_systemsnmp_community`.
* Fix the issue in `data_source_router_policy` and delete the redundant data resource `data_source_router_policylist`.
* Fix the return data type change in v7.6.0 `fortiswitch_switchacl_ingress`, `fortiswitch_switchacl_egree` and `fortiswitch_switchacl_prelookup`.

IMPROVEMENTS:

* Support new version 7.6.0.
* Add a new documentation for concurrent operations on FortiSwitch.


## 1.1.2 (April 18, 2024)

BUGFIXES:

* Fix data type issues in `fortiswitch_switchacl_ingress` and `fortiswitch_switch_interface`;
* Fix the provider crash issue in `fortiswitch_systemdhcp_server`;
* Fix the issue that the sensitive parameter is changed to null while the parameter is not configured.

## 1.1.1 (February 9, 2024)

BUGFIXES:

* Fix data type issues in `fortiswitch_switchacl_egress` and `fortiswitch_switch_prelookup`;

IMPROVEMENTS:

* Support new version 7.4.2;

## 1.1.0 (October 24, 2023)

IMPROVEMENTS:

* Support new versions 7.2.4, 7.2.5, 7.4.0 and 7.4.1;
* Fix issue that mkey is missing in `fortiswitch_router_ospf` and `fortiswitch_router_ospf6`;
* Update the document for `fortiswitch_router_ospf` and `fortiswitch_router_ospf6`；
* Update the supported fortiSwitch versions in the document.

FEATURES:

* **New Resource:** `fortiswitch_routerospf6_interface`
* **New Resource:** `fortiswitch_routerospf_interface`
* **New Resource:** `fortiswitch_systemptp_interfacepolicy`
* **New Resource:** `fortiswitch_systemptp_profile`

## 1.0.0 (August 3, 2023)

FEATURES:

* **New Resource:** `fortiswitch_alertemail_setting`
* **New Resource:** `fortiswitch_antivirus_heuristic`
* **New Resource:** `fortiswitch_antivirus_profile`
* **New Resource:** `fortiswitch_antivirus_quarantine`
* **New Resource:** `fortiswitch_antivirus_settings`
* **New Resource:** `fortiswitch_application_custom`
* **New Resource:** `fortiswitch_application_group`
* **New Resource:** `fortiswitch_application_list`
* **New Resource:** `fortiswitch_application_name`
* **New Resource:** `fortiswitch_application_rulesettings`
* **New Resource:** `fortiswitch_authentication_rule`
* **New Resource:** `fortiswitch_authentication_scheme`
* **New Resource:** `fortiswitch_authentication_setting`
* **New Resource:** `fortiswitch_automation_setting`
* **New Resource:** `fortiswitch_certificate_ca`
* **New Resource:** `fortiswitch_certificate_crl`
* **New Resource:** `fortiswitch_certificate_local`
* **New Resource:** `fortiswitch_certificate_remote`
* **New Resource:** `fortiswitch_cifs_domaincontroller`
* **New Resource:** `fortiswitch_cifs_profile`
* **New Resource:** `fortiswitch_credentialstore_domaincontroller`
* **New Resource:** `fortiswitch_dlp_datatype`
* **New Resource:** `fortiswitch_dlp_dictionary`
* **New Resource:** `fortiswitch_dlp_filepattern`
* **New Resource:** `fortiswitch_dlp_fpdocsource`
* **New Resource:** `fortiswitch_dlp_fpsensitivity`
* **New Resource:** `fortiswitch_dlp_profile`
* **New Resource:** `fortiswitch_dlp_sensitivity`
* **New Resource:** `fortiswitch_dlp_sensor`
* **New Resource:** `fortiswitch_dlp_settings`
* **New Resource:** `fortiswitch_dnsfilter_domainfilter`
* **New Resource:** `fortiswitch_dnsfilter_profile`
* **New Resource:** `fortiswitch_dpdk_cpus`
* **New Resource:** `fortiswitch_dpdk_global`
* **New Resource:** `fortiswitch_emailfilter_blockallowlist`
* **New Resource:** `fortiswitch_emailfilter_bwl`
* **New Resource:** `fortiswitch_emailfilter_bword`
* **New Resource:** `fortiswitch_emailfilter_dnsbl`
* **New Resource:** `fortiswitch_emailfilter_fortishield`
* **New Resource:** `fortiswitch_emailfilter_iptrust`
* **New Resource:** `fortiswitch_emailfilter_mheader`
* **New Resource:** `fortiswitch_emailfilter_options`
* **New Resource:** `fortiswitch_emailfilter_profile`
* **New Resource:** `fortiswitch_endpointcontrol_client`
* **New Resource:** `fortiswitch_endpointcontrol_fctems`
* **New Resource:** `fortiswitch_endpointcontrol_forticlientems`
* **New Resource:** `fortiswitch_endpointcontrol_forticlientregistrationsync`
* **New Resource:** `fortiswitch_endpointcontrol_profile`
* **New Resource:** `fortiswitch_endpointcontrol_registeredforticlient`
* **New Resource:** `fortiswitch_endpointcontrol_settings`
* **New Resource:** `fortiswitch_extendercontroller_dataplan`
* **New Resource:** `fortiswitch_extendercontroller_extender`
* **New Resource:** `fortiswitch_extendercontroller_extender1`
* **New Resource:** `fortiswitch_extendercontroller_extenderprofile`
* **New Resource:** `fortiswitch_extensioncontroller_dataplan`
* **New Resource:** `fortiswitch_extensioncontroller_extender`
* **New Resource:** `fortiswitch_extensioncontroller_extenderprofile`
* **New Resource:** `fortiswitch_extensioncontroller_fortigate`
* **New Resource:** `fortiswitch_extensioncontroller_fortigateprofile`
* **New Resource:** `fortiswitch_filefilter_profile`
* **New Resource:** `fortiswitch_firewall_DoSpolicy`
* **New Resource:** `fortiswitch_firewall_DoSpolicy6`
* **New Resource:** `fortiswitch_firewall_accessproxy`
* **New Resource:** `fortiswitch_firewall_accessproxy6`
* **New Resource:** `fortiswitch_firewall_accessproxysshclientcert`
* **New Resource:** `fortiswitch_firewall_accessproxyvirtualhost`
* **New Resource:** `fortiswitch_firewall_address`
* **New Resource:** `fortiswitch_firewall_address6`
* **New Resource:** `fortiswitch_firewall_address6template`
* **New Resource:** `fortiswitch_firewall_addrgrp`
* **New Resource:** `fortiswitch_firewall_addrgrp6`
* **New Resource:** `fortiswitch_firewall_authportal`
* **New Resource:** `fortiswitch_firewall_centralsnatmap`
* **New Resource:** `fortiswitch_firewall_centralsnatmap_move`
* **New Resource:** `fortiswitch_firewall_centralsnatmap_sort`
* **New Resource:** `fortiswitch_firewall_city`
* **New Resource:** `fortiswitch_firewall_country`
* **New Resource:** `fortiswitch_firewall_decryptedtrafficmirror`
* **New Resource:** `fortiswitch_firewall_dnstranslation`
* **New Resource:** `fortiswitch_firewall_global`
* **New Resource:** `fortiswitch_firewall_identitybasedroute`
* **New Resource:** `fortiswitch_firewall_interfacepolicy`
* **New Resource:** `fortiswitch_firewall_interfacepolicy6`
* **New Resource:** `fortiswitch_firewall_internetservice`
* **New Resource:** `fortiswitch_firewall_internetserviceaddition`
* **New Resource:** `fortiswitch_firewall_internetserviceappend`
* **New Resource:** `fortiswitch_firewall_internetservicebotnet`
* **New Resource:** `fortiswitch_firewall_internetservicecustom`
* **New Resource:** `fortiswitch_firewall_internetservicecustomgroup`
* **New Resource:** `fortiswitch_firewall_internetservicedefinition`
* **New Resource:** `fortiswitch_firewall_internetserviceextension`
* **New Resource:** `fortiswitch_firewall_internetservicegroup`
* **New Resource:** `fortiswitch_firewall_internetserviceipblreason`
* **New Resource:** `fortiswitch_firewall_internetserviceipblvendor`
* **New Resource:** `fortiswitch_firewall_internetservicelist`
* **New Resource:** `fortiswitch_firewall_internetservicename`
* **New Resource:** `fortiswitch_firewall_internetserviceowner`
* **New Resource:** `fortiswitch_firewall_internetservicereputation`
* **New Resource:** `fortiswitch_firewall_ippool`
* **New Resource:** `fortiswitch_firewall_ippool6`
* **New Resource:** `fortiswitch_firewall_iptranslation`
* **New Resource:** `fortiswitch_firewall_ipv6ehfilter`
* **New Resource:** `fortiswitch_firewall_ldbmonitor`
* **New Resource:** `fortiswitch_firewall_localinpolicy`
* **New Resource:** `fortiswitch_firewall_localinpolicy6`
* **New Resource:** `fortiswitch_firewall_multicastaddress`
* **New Resource:** `fortiswitch_firewall_multicastaddress6`
* **New Resource:** `fortiswitch_firewall_multicastpolicy`
* **New Resource:** `fortiswitch_firewall_multicastpolicy6`
* **New Resource:** `fortiswitch_firewall_networkservicedynamic`
* **New Resource:** `fortiswitch_firewall_object_address`
* **New Resource:** `fortiswitch_firewall_object_addressgroup`
* **New Resource:** `fortiswitch_firewall_object_ippool`
* **New Resource:** `fortiswitch_firewall_object_service`
* **New Resource:** `fortiswitch_firewall_object_servicecategory`
* **New Resource:** `fortiswitch_firewall_object_servicegroup`
* **New Resource:** `fortiswitch_firewall_object_vip`
* **New Resource:** `fortiswitch_firewall_object_vipgroup`
* **New Resource:** `fortiswitch_firewall_policy`
* **New Resource:** `fortiswitch_firewall_policy46`
* **New Resource:** `fortiswitch_firewall_policy6`
* **New Resource:** `fortiswitch_firewall_policy64`
* **New Resource:** `fortiswitch_firewall_profilegroup`
* **New Resource:** `fortiswitch_firewall_profileprotocoloptions`
* **New Resource:** `fortiswitch_firewall_proxyaddress`
* **New Resource:** `fortiswitch_firewall_proxyaddrgrp`
* **New Resource:** `fortiswitch_firewall_proxypolicy`
* **New Resource:** `fortiswitch_firewall_proxypolicy_move`
* **New Resource:** `fortiswitch_firewall_proxypolicy_sort`
* **New Resource:** `fortiswitch_firewall_region`
* **New Resource:** `fortiswitch_firewall_security_policy`
* **New Resource:** `fortiswitch_firewall_security_policyseq`
* **New Resource:** `fortiswitch_firewall_security_policysort`
* **New Resource:** `fortiswitch_firewall_securitypolicy`
* **New Resource:** `fortiswitch_firewall_shapingpolicy`
* **New Resource:** `fortiswitch_firewall_shapingprofile`
* **New Resource:** `fortiswitch_firewall_sniffer`
* **New Resource:** `fortiswitch_firewall_sslserver`
* **New Resource:** `fortiswitch_firewall_sslsshprofile`
* **New Resource:** `fortiswitch_firewall_trafficclass`
* **New Resource:** `fortiswitch_firewall_ttlpolicy`
* **New Resource:** `fortiswitch_firewall_vendormac`
* **New Resource:** `fortiswitch_firewall_vip`
* **New Resource:** `fortiswitch_firewall_vip46`
* **New Resource:** `fortiswitch_firewall_vip6`
* **New Resource:** `fortiswitch_firewall_vip64`
* **New Resource:** `fortiswitch_firewall_vipgrp`
* **New Resource:** `fortiswitch_firewall_vipgrp46`
* **New Resource:** `fortiswitch_firewall_vipgrp6`
* **New Resource:** `fortiswitch_firewall_vipgrp64`
* **New Resource:** `fortiswitch_firewallconsolidated_policy`
* **New Resource:** `fortiswitch_firewallipmacbinding_setting`
* **New Resource:** `fortiswitch_firewallipmacbinding_table`
* **New Resource:** `fortiswitch_firewallschedule_group`
* **New Resource:** `fortiswitch_firewallschedule_onetime`
* **New Resource:** `fortiswitch_firewallschedule_recurring`
* **New Resource:** `fortiswitch_firewallservice_category`
* **New Resource:** `fortiswitch_firewallservice_custom`
* **New Resource:** `fortiswitch_firewallservice_group`
* **New Resource:** `fortiswitch_firewallshaper_peripshaper`
* **New Resource:** `fortiswitch_firewallshaper_trafficshaper`
* **New Resource:** `fortiswitch_firewallssh_hostkey`
* **New Resource:** `fortiswitch_firewallssh_localca`
* **New Resource:** `fortiswitch_firewallssh_localkey`
* **New Resource:** `fortiswitch_firewallssh_setting`
* **New Resource:** `fortiswitch_firewallssl_setting`
* **New Resource:** `fortiswitch_firewallwildcardfqdn_custom`
* **New Resource:** `fortiswitch_firewallwildcardfqdn_group`
* **New Resource:** `fortiswitch_fortimanager_devicemanager_device`
* **New Resource:** `fortiswitch_fortimanager_devicemanager_installdevice`
* **New Resource:** `fortiswitch_fortimanager_devicemanager_installpackage`
* **New Resource:** `fortiswitch_fortimanager_devicemanager_script`
* **New Resource:** `fortiswitch_fortimanager_devicemanager_script_execute`
* **New Resource:** `fortiswitch_fortimanager_firewall_object_address`
* **New Resource:** `fortiswitch_fortimanager_firewall_object_ippool`
* **New Resource:** `fortiswitch_fortimanager_firewall_object_service`
* **New Resource:** `fortiswitch_fortimanager_firewall_object_vip`
* **New Resource:** `fortiswitch_fortimanager_firewall_security_policy`
* **New Resource:** `fortiswitch_fortimanager_firewall_security_policypackage`
* **New Resource:** `fortiswitch_fortimanager_jsonrpc_request`
* **New Resource:** `fortiswitch_fortimanager_object_adom_revision`
* **New Resource:** `fortiswitch_fortimanager_system_admin`
* **New Resource:** `fortiswitch_fortimanager_system_admin_profiles`
* **New Resource:** `fortiswitch_fortimanager_system_admin_user`
* **New Resource:** `fortiswitch_fortimanager_system_adom`
* **New Resource:** `fortiswitch_fortimanager_system_dns`
* **New Resource:** `fortiswitch_fortimanager_system_global`
* **New Resource:** `fortiswitch_fortimanager_system_license_forticare`
* **New Resource:** `fortiswitch_fortimanager_system_license_vm`
* **New Resource:** `fortiswitch_fortimanager_system_network_interface`
* **New Resource:** `fortiswitch_fortimanager_system_network_route`
* **New Resource:** `fortiswitch_fortimanager_system_ntp`
* **New Resource:** `fortiswitch_fortimanager_system_syslogserver`
* **New Resource:** `fortiswitch_ftpproxy_explicit`
* **New Resource:** `fortiswitch_gui_console`
* **New Resource:** `fortiswitch_icap_profile`
* **New Resource:** `fortiswitch_icap_server`
* **New Resource:** `fortiswitch_icap_servergroup`
* **New Resource:** `fortiswitch_ips_custom`
* **New Resource:** `fortiswitch_ips_decoder`
* **New Resource:** `fortiswitch_ips_global`
* **New Resource:** `fortiswitch_ips_rule`
* **New Resource:** `fortiswitch_ips_rulesettings`
* **New Resource:** `fortiswitch_ips_sensor`
* **New Resource:** `fortiswitch_ips_settings`
* **New Resource:** `fortiswitch_ips_viewmap`
* **New Resource:** `fortiswitch_json_generic_api`
* **New Resource:** `fortiswitch_log_customfield`
* **New Resource:** `fortiswitch_log_eventfilter`
* **New Resource:** `fortiswitch_log_fortianalyzer_setting`
* **New Resource:** `fortiswitch_log_gui`
* **New Resource:** `fortiswitch_log_guidisplay`
* **New Resource:** `fortiswitch_log_setting`
* **New Resource:** `fortiswitch_log_syslog_setting`
* **New Resource:** `fortiswitch_log_threatweight`
* **New Resource:** `fortiswitch_logdisk_filter`
* **New Resource:** `fortiswitch_logdisk_setting`
* **New Resource:** `fortiswitch_logfortianalyzer2_filter`
* **New Resource:** `fortiswitch_logfortianalyzer2_overridefilter`
* **New Resource:** `fortiswitch_logfortianalyzer2_overridesetting`
* **New Resource:** `fortiswitch_logfortianalyzer2_setting`
* **New Resource:** `fortiswitch_logfortianalyzer3_filter`
* **New Resource:** `fortiswitch_logfortianalyzer3_overridefilter`
* **New Resource:** `fortiswitch_logfortianalyzer3_overridesetting`
* **New Resource:** `fortiswitch_logfortianalyzer3_setting`
* **New Resource:** `fortiswitch_logfortianalyzer_filter`
* **New Resource:** `fortiswitch_logfortianalyzer_overridefilter`
* **New Resource:** `fortiswitch_logfortianalyzer_overridesetting`
* **New Resource:** `fortiswitch_logfortianalyzer_setting`
* **New Resource:** `fortiswitch_logfortianalyzercloud_filter`
* **New Resource:** `fortiswitch_logfortianalyzercloud_overridefilter`
* **New Resource:** `fortiswitch_logfortianalyzercloud_overridesetting`
* **New Resource:** `fortiswitch_logfortianalyzercloud_setting`
* **New Resource:** `fortiswitch_logfortiguard_filter`
* **New Resource:** `fortiswitch_logfortiguard_overridefilter`
* **New Resource:** `fortiswitch_logfortiguard_overridesetting`
* **New Resource:** `fortiswitch_logfortiguard_setting`
* **New Resource:** `fortiswitch_logmemory_filter`
* **New Resource:** `fortiswitch_logmemory_globalsetting`
* **New Resource:** `fortiswitch_logmemory_setting`
* **New Resource:** `fortiswitch_lognulldevice_filter`
* **New Resource:** `fortiswitch_lognulldevice_setting`
* **New Resource:** `fortiswitch_logremote_setting`
* **New Resource:** `fortiswitch_logsyslogd2_filter`
* **New Resource:** `fortiswitch_logsyslogd2_overridefilter`
* **New Resource:** `fortiswitch_logsyslogd2_overridesetting`
* **New Resource:** `fortiswitch_logsyslogd2_setting`
* **New Resource:** `fortiswitch_logsyslogd3_filter`
* **New Resource:** `fortiswitch_logsyslogd3_overridefilter`
* **New Resource:** `fortiswitch_logsyslogd3_overridesetting`
* **New Resource:** `fortiswitch_logsyslogd3_setting`
* **New Resource:** `fortiswitch_logsyslogd4_filter`
* **New Resource:** `fortiswitch_logsyslogd4_overridefilter`
* **New Resource:** `fortiswitch_logsyslogd4_overridesetting`
* **New Resource:** `fortiswitch_logsyslogd4_setting`
* **New Resource:** `fortiswitch_logsyslogd_filter`
* **New Resource:** `fortiswitch_logsyslogd_overridefilter`
* **New Resource:** `fortiswitch_logsyslogd_overridesetting`
* **New Resource:** `fortiswitch_logsyslogd_setting`
* **New Resource:** `fortiswitch_logtacacsaccounting2_filter`
* **New Resource:** `fortiswitch_logtacacsaccounting2_setting`
* **New Resource:** `fortiswitch_logtacacsaccounting3_filter`
* **New Resource:** `fortiswitch_logtacacsaccounting3_setting`
* **New Resource:** `fortiswitch_logtacacsaccounting_filter`
* **New Resource:** `fortiswitch_logtacacsaccounting_setting`
* **New Resource:** `fortiswitch_logwebtrends_filter`
* **New Resource:** `fortiswitch_logwebtrends_setting`
* **New Resource:** `fortiswitch_networking_interface_port`
* **New Resource:** `fortiswitch_networking_route_static`
* **New Resource:** `fortiswitch_nsxt_servicechain`
* **New Resource:** `fortiswitch_nsxt_setting`
* **New Resource:** `fortiswitch_report_chart`
* **New Resource:** `fortiswitch_report_dataset`
* **New Resource:** `fortiswitch_report_layout`
* **New Resource:** `fortiswitch_report_setting`
* **New Resource:** `fortiswitch_report_style`
* **New Resource:** `fortiswitch_report_theme`
* **New Resource:** `fortiswitch_router_accesslist`
* **New Resource:** `fortiswitch_router_accesslist6`
* **New Resource:** `fortiswitch_router_aspathlist`
* **New Resource:** `fortiswitch_router_authpath`
* **New Resource:** `fortiswitch_router_bfd`
* **New Resource:** `fortiswitch_router_bfd6`
* **New Resource:** `fortiswitch_router_bgp`
* **New Resource:** `fortiswitch_router_communitylist`
* **New Resource:** `fortiswitch_router_gwdetect`
* **New Resource:** `fortiswitch_router_isis`
* **New Resource:** `fortiswitch_router_keychain`
* **New Resource:** `fortiswitch_router_multicast`
* **New Resource:** `fortiswitch_router_multicast6`
* **New Resource:** `fortiswitch_router_multicastflow`
* **New Resource:** `fortiswitch_router_ospf`
* **New Resource:** `fortiswitch_router_ospf6`
* **New Resource:** `fortiswitch_router_policy`
* **New Resource:** `fortiswitch_router_policy6`
* **New Resource:** `fortiswitch_router_prefixlist`
* **New Resource:** `fortiswitch_router_prefixlist6`
* **New Resource:** `fortiswitch_router_rip`
* **New Resource:** `fortiswitch_router_ripng`
* **New Resource:** `fortiswitch_router_routemap`
* **New Resource:** `fortiswitch_router_setting`
* **New Resource:** `fortiswitch_router_static`
* **New Resource:** `fortiswitch_router_static6`
* **New Resource:** `fortiswitch_router_vrf`
* **New Resource:** `fortiswitch_routerbgp_neighbor`
* **New Resource:** `fortiswitch_routerbgp_network`
* **New Resource:** `fortiswitch_routerbgp_network6`
* **New Resource:** `fortiswitch_routerospf6_ospf6interface`
* **New Resource:** `fortiswitch_routerospf_neighbor`
* **New Resource:** `fortiswitch_routerospf_network`
* **New Resource:** `fortiswitch_routerospf_ospfinterface`
* **New Resource:** `fortiswitch_sctpfilter_profile`
* **New Resource:** `fortiswitch_spamfilter_bwl`
* **New Resource:** `fortiswitch_spamfilter_bword`
* **New Resource:** `fortiswitch_spamfilter_dnsbl`
* **New Resource:** `fortiswitch_spamfilter_fortishield`
* **New Resource:** `fortiswitch_spamfilter_iptrust`
* **New Resource:** `fortiswitch_spamfilter_mheader`
* **New Resource:** `fortiswitch_spamfilter_options`
* **New Resource:** `fortiswitch_spamfilter_profile`
* **New Resource:** `fortiswitch_sshfilter_profile`
* **New Resource:** `fortiswitch_switch_autoislportgroup`
* **New Resource:** `fortiswitch_switch_autonetwork`
* **New Resource:** `fortiswitch_switch_domain`
* **New Resource:** `fortiswitch_switch_global`
* **New Resource:** `fortiswitch_switch_interface`
* **New Resource:** `fortiswitch_switch_ipmacbinding`
* **New Resource:** `fortiswitch_switch_mirror`
* **New Resource:** `fortiswitch_switch_phymode`
* **New Resource:** `fortiswitch_switch_physicalport`
* **New Resource:** `fortiswitch_switch_quarantine`
* **New Resource:** `fortiswitch_switch_raguardpolicy`
* **New Resource:** `fortiswitch_switch_securityfeature`
* **New Resource:** `fortiswitch_switch_staticmac`
* **New Resource:** `fortiswitch_switch_stormcontrol`
* **New Resource:** `fortiswitch_switch_trunk`
* **New Resource:** `fortiswitch_switch_virtualwire`
* **New Resource:** `fortiswitch_switch_vlan`
* **New Resource:** `fortiswitch_switch_vlantpid`
* **New Resource:** `fortiswitch_switchacl_8021X`
* **New Resource:** `fortiswitch_switchacl_egress`
* **New Resource:** `fortiswitch_switchacl_ingress`
* **New Resource:** `fortiswitch_switchacl_policer`
* **New Resource:** `fortiswitch_switchacl_prelookup`
* **New Resource:** `fortiswitch_switchacl_settings`
* **New Resource:** `fortiswitch_switchaclservice_custom`
* **New Resource:** `fortiswitch_switchcontroller_8021Xsettings`
* **New Resource:** `fortiswitch_switchcontroller_customcommand`
* **New Resource:** `fortiswitch_switchcontroller_dynamicportpolicy`
* **New Resource:** `fortiswitch_switchcontroller_flowtracking`
* **New Resource:** `fortiswitch_switchcontroller_fortilinksettings`
* **New Resource:** `fortiswitch_switchcontroller_global`
* **New Resource:** `fortiswitch_switchcontroller_igmpsnooping`
* **New Resource:** `fortiswitch_switchcontroller_lldpprofile`
* **New Resource:** `fortiswitch_switchcontroller_lldpsettings`
* **New Resource:** `fortiswitch_switchcontroller_location`
* **New Resource:** `fortiswitch_switchcontroller_macsyncsettings`
* **New Resource:** `fortiswitch_switchcontroller_managedswitch`
* **New Resource:** `fortiswitch_switchcontroller_nacdevice`
* **New Resource:** `fortiswitch_switchcontroller_nacsettings`
* **New Resource:** `fortiswitch_switchcontroller_networkmonitorsettings`
* **New Resource:** `fortiswitch_switchcontroller_portpolicy`
* **New Resource:** `fortiswitch_switchcontroller_quarantine`
* **New Resource:** `fortiswitch_switchcontroller_remotelog`
* **New Resource:** `fortiswitch_switchcontroller_sflow`
* **New Resource:** `fortiswitch_switchcontroller_snmpcommunity`
* **New Resource:** `fortiswitch_switchcontroller_snmpsysinfo`
* **New Resource:** `fortiswitch_switchcontroller_snmptrapthreshold`
* **New Resource:** `fortiswitch_switchcontroller_snmpuser`
* **New Resource:** `fortiswitch_switchcontroller_stormcontrol`
* **New Resource:** `fortiswitch_switchcontroller_stormcontrolpolicy`
* **New Resource:** `fortiswitch_switchcontroller_stpinstance`
* **New Resource:** `fortiswitch_switchcontroller_stpsettings`
* **New Resource:** `fortiswitch_switchcontroller_switchgroup`
* **New Resource:** `fortiswitch_switchcontroller_switchinterfacetag`
* **New Resource:** `fortiswitch_switchcontroller_switchlog`
* **New Resource:** `fortiswitch_switchcontroller_switchprofile`
* **New Resource:** `fortiswitch_switchcontroller_system`
* **New Resource:** `fortiswitch_switchcontroller_trafficpolicy`
* **New Resource:** `fortiswitch_switchcontroller_trafficsniffer`
* **New Resource:** `fortiswitch_switchcontroller_virtualportpool`
* **New Resource:** `fortiswitch_switchcontroller_vlan`
* **New Resource:** `fortiswitch_switchcontroller_vlanpolicy`
* **New Resource:** `fortiswitch_switchcontrollerautoconfig_custom`
* **New Resource:** `fortiswitch_switchcontrollerautoconfig_default`
* **New Resource:** `fortiswitch_switchcontrollerautoconfig_policy`
* **New Resource:** `fortiswitch_switchcontrollerinitialconfig_template`
* **New Resource:** `fortiswitch_switchcontrollerinitialconfig_vlans`
* **New Resource:** `fortiswitch_switchcontrollerptp_policy`
* **New Resource:** `fortiswitch_switchcontrollerptp_settings`
* **New Resource:** `fortiswitch_switchcontrollerqos_dot1pmap`
* **New Resource:** `fortiswitch_switchcontrollerqos_ipdscpmap`
* **New Resource:** `fortiswitch_switchcontrollerqos_qospolicy`
* **New Resource:** `fortiswitch_switchcontrollerqos_queuepolicy`
* **New Resource:** `fortiswitch_switchcontrollersecuritypolicy_8021X`
* **New Resource:** `fortiswitch_switchcontrollersecuritypolicy_captiveportal`
* **New Resource:** `fortiswitch_switchcontrollersecuritypolicy_localaccess`
* **New Resource:** `fortiswitch_switchigmpsnooping_globals`
* **New Resource:** `fortiswitch_switchlldp_profile`
* **New Resource:** `fortiswitch_switchlldp_settings`
* **New Resource:** `fortiswitch_switchmacsec_profile`
* **New Resource:** `fortiswitch_switchmldsnooping_globals`
* **New Resource:** `fortiswitch_switchnetworkmonitor_directed`
* **New Resource:** `fortiswitch_switchnetworkmonitor_settings`
* **New Resource:** `fortiswitch_switchptp_policy`
* **New Resource:** `fortiswitch_switchptp_settings`
* **New Resource:** `fortiswitch_switchqos_dot1pmap`
* **New Resource:** `fortiswitch_switchqos_ipdscpmap`
* **New Resource:** `fortiswitch_switchqos_qospolicy`
* **New Resource:** `fortiswitch_switchstp_instance`
* **New Resource:** `fortiswitch_switchstp_settings`
* **New Resource:** `fortiswitch_system3gmodem_custom`
* **New Resource:** `fortiswitch_system_accprofile`
* **New Resource:** `fortiswitch_system_acme`
* **New Resource:** `fortiswitch_system_admin`
* **New Resource:** `fortiswitch_system_admin_administrator`
* **New Resource:** `fortiswitch_system_admin_profiles`
* **New Resource:** `fortiswitch_system_affinityinterrupt`
* **New Resource:** `fortiswitch_system_affinitypacketredistribution`
* **New Resource:** `fortiswitch_system_alarm`
* **New Resource:** `fortiswitch_system_alertemail`
* **New Resource:** `fortiswitch_system_alias`
* **New Resource:** `fortiswitch_system_apiuser`
* **New Resource:** `fortiswitch_system_apiuser_setting`
* **New Resource:** `fortiswitch_system_arptable`
* **New Resource:** `fortiswitch_system_autoinstall`
* **New Resource:** `fortiswitch_system_automationaction`
* **New Resource:** `fortiswitch_system_automationdestination`
* **New Resource:** `fortiswitch_system_automationstitch`
* **New Resource:** `fortiswitch_system_automationtrigger`
* **New Resource:** `fortiswitch_system_autoscript`
* **New Resource:** `fortiswitch_system_bugreport`
* **New Resource:** `fortiswitch_system_centralmanagement`
* **New Resource:** `fortiswitch_system_clustersync`
* **New Resource:** `fortiswitch_system_console`
* **New Resource:** `fortiswitch_system_csf`
* **New Resource:** `fortiswitch_system_customlanguage`
* **New Resource:** `fortiswitch_system_ddns`
* **New Resource:** `fortiswitch_system_dedicatedmgmt`
* **New Resource:** `fortiswitch_system_dns`
* **New Resource:** `fortiswitch_system_dns64`
* **New Resource:** `fortiswitch_system_dnsdatabase`
* **New Resource:** `fortiswitch_system_dnsserver`
* **New Resource:** `fortiswitch_system_dscpbasedpriority`
* **New Resource:** `fortiswitch_system_emailserver`
* **New Resource:** `fortiswitch_system_external`
* **New Resource:** `fortiswitch_system_federatedupgrade`
* **New Resource:** `fortiswitch_system_fipscc`
* **New Resource:** `fortiswitch_system_flancloud`
* **New Resource:** `fortiswitch_system_flowexport`
* **New Resource:** `fortiswitch_system_fm`
* **New Resource:** `fortiswitch_system_fortiai`
* **New Resource:** `fortiswitch_system_fortianalyzer`
* **New Resource:** `fortiswitch_system_fortianalyzer2`
* **New Resource:** `fortiswitch_system_fortianalyzer3`
* **New Resource:** `fortiswitch_system_fortiguard`
* **New Resource:** `fortiswitch_system_fortimanager`
* **New Resource:** `fortiswitch_system_fortindr`
* **New Resource:** `fortiswitch_system_fortisandbox`
* **New Resource:** `fortiswitch_system_fssopolling`
* **New Resource:** `fortiswitch_system_fswcloud`
* **New Resource:** `fortiswitch_system_ftmpush`
* **New Resource:** `fortiswitch_system_geneve`
* **New Resource:** `fortiswitch_system_geoipcountry`
* **New Resource:** `fortiswitch_system_geoipoverride`
* **New Resource:** `fortiswitch_system_global`
* **New Resource:** `fortiswitch_system_gretunnel`
* **New Resource:** `fortiswitch_system_ha`
* **New Resource:** `fortiswitch_system_hamonitor`
* **New Resource:** `fortiswitch_system_ike`
* **New Resource:** `fortiswitch_system_interface`
* **New Resource:** `fortiswitch_system_ipam`
* **New Resource:** `fortiswitch_system_ipiptunnel`
* **New Resource:** `fortiswitch_system_ips`
* **New Resource:** `fortiswitch_system_ipsecaggregate`
* **New Resource:** `fortiswitch_system_ipsurlfilterdns`
* **New Resource:** `fortiswitch_system_ipsurlfilterdns6`
* **New Resource:** `fortiswitch_system_ipv6neighborcache`
* **New Resource:** `fortiswitch_system_ipv6tunnel`
* **New Resource:** `fortiswitch_system_license_forticare`
* **New Resource:** `fortiswitch_system_license_vdom`
* **New Resource:** `fortiswitch_system_license_vm`
* **New Resource:** `fortiswitch_system_linkmonitor`
* **New Resource:** `fortiswitch_system_location`
* **New Resource:** `fortiswitch_system_ltemodem`
* **New Resource:** `fortiswitch_system_macaddresstable`
* **New Resource:** `fortiswitch_system_managementtunnel`
* **New Resource:** `fortiswitch_system_mobiletunnel`
* **New Resource:** `fortiswitch_system_modem`
* **New Resource:** `fortiswitch_system_nat64`
* **New Resource:** `fortiswitch_system_ndproxy`
* **New Resource:** `fortiswitch_system_netflow`
* **New Resource:** `fortiswitch_system_networkvisibility`
* **New Resource:** `fortiswitch_system_npu`
* **New Resource:** `fortiswitch_system_ntp`
* **New Resource:** `fortiswitch_system_objecttag`
* **New Resource:** `fortiswitch_system_objecttagging`
* **New Resource:** `fortiswitch_system_passwordpolicy`
* **New Resource:** `fortiswitch_system_passwordpolicyguestadmin`
* **New Resource:** `fortiswitch_system_physicalswitch`
* **New Resource:** `fortiswitch_system_portpair`
* **New Resource:** `fortiswitch_system_pppoeinterface`
* **New Resource:** `fortiswitch_system_proberesponse`
* **New Resource:** `fortiswitch_system_proxyarp`
* **New Resource:** `fortiswitch_system_ptp`
* **New Resource:** `fortiswitch_system_replacemsggroup`
* **New Resource:** `fortiswitch_system_replacemsgimage`
* **New Resource:** `fortiswitch_system_limits`
* **New Resource:** `fortiswitch_system_saml`
* **New Resource:** `fortiswitch_system_sdnconnector`
* **New Resource:** `fortiswitch_system_sdwan`
* **New Resource:** `fortiswitch_system_sessionhelper`
* **New Resource:** `fortiswitch_system_sessionttl`
* **New Resource:** `fortiswitch_system_setting_dns`
* **New Resource:** `fortiswitch_system_setting_global`
* **New Resource:** `fortiswitch_system_setting_ntp`
* **New Resource:** `fortiswitch_system_settings`
* **New Resource:** `fortiswitch_system_sflow`
* **New Resource:** `fortiswitch_system_sittunnel`
* **New Resource:** `fortiswitch_system_smsserver`
* **New Resource:** `fortiswitch_system_snifferprofile`
* **New Resource:** `fortiswitch_system_speedtestschedule`
* **New Resource:** `fortiswitch_system_speedtestserver`
* **New Resource:** `fortiswitch_system_ssoadmin`
* **New Resource:** `fortiswitch_system_ssoforticloudadmin`
* **New Resource:** `fortiswitch_system_standalonecluster`
* **New Resource:** `fortiswitch_system_storage`
* **New Resource:** `fortiswitch_system_stp`
* **New Resource:** `fortiswitch_system_switchinterface`
* **New Resource:** `fortiswitch_system_tosbasedpriority`
* **New Resource:** `fortiswitch_system_vdom`
* **New Resource:** `fortiswitch_system_vdom_setting`
* **New Resource:** `fortiswitch_system_vdomdns`
* **New Resource:** `fortiswitch_system_vdomexception`
* **New Resource:** `fortiswitch_system_vdomlink`
* **New Resource:** `fortiswitch_system_vdomnetflow`
* **New Resource:** `fortiswitch_system_vdomproperty`
* **New Resource:** `fortiswitch_system_vdomradiusserver`
* **New Resource:** `fortiswitch_system_vdomsflow`
* **New Resource:** `fortiswitch_system_virtualswitch`
* **New Resource:** `fortiswitch_system_virtualwanlink`
* **New Resource:** `fortiswitch_system_virtualwirepair`
* **New Resource:** `fortiswitch_system_vnetunnel`
* **New Resource:** `fortiswitch_system_vxlan`
* **New Resource:** `fortiswitch_system_wccp`
* **New Resource:** `fortiswitch_system_web`
* **New Resource:** `fortiswitch_system_zone`
* **New Resource:** `fortiswitch_systemalias_command`
* **New Resource:** `fortiswitch_systemalias_group`
* **New Resource:** `fortiswitch_systemautoupdate_clientoverride`
* **New Resource:** `fortiswitch_systemautoupdate_override`
* **New Resource:** `fortiswitch_systemautoupdate_pushupdate`
* **New Resource:** `fortiswitch_systemautoupdate_schedule`
* **New Resource:** `fortiswitch_systemautoupdate_tunneling`
* **New Resource:** `fortiswitch_systemcertificate_ca`
* **New Resource:** `fortiswitch_systemcertificate_crl`
* **New Resource:** `fortiswitch_systemcertificate_local`
* **New Resource:** `fortiswitch_systemcertificate_ocsp`
* **New Resource:** `fortiswitch_systemcertificate_remote`
* **New Resource:** `fortiswitch_systemdhcp6_server`
* **New Resource:** `fortiswitch_systemdhcp_server`
* **New Resource:** `fortiswitch_systemlldp_networkpolicy`
* **New Resource:** `fortiswitch_systemreplacemsg_admin`
* **New Resource:** `fortiswitch_systemreplacemsg_alertmail`
* **New Resource:** `fortiswitch_systemreplacemsg_auth`
* **New Resource:** `fortiswitch_systemreplacemsg_automation`
* **New Resource:** `fortiswitch_systemreplacemsg_devicedetectionportal`
* **New Resource:** `fortiswitch_systemreplacemsg_ec`
* **New Resource:** `fortiswitch_systemreplacemsg_fortiguardwf`
* **New Resource:** `fortiswitch_systemreplacemsg_ftp`
* **New Resource:** `fortiswitch_systemreplacemsg_http`
* **New Resource:** `fortiswitch_systemreplacemsg_icap`
* **New Resource:** `fortiswitch_systemreplacemsg_mail`
* **New Resource:** `fortiswitch_systemreplacemsg_nacquar`
* **New Resource:** `fortiswitch_systemreplacemsg_nntp`
* **New Resource:** `fortiswitch_systemreplacemsg_spam`
* **New Resource:** `fortiswitch_systemreplacemsg_sslvpn`
* **New Resource:** `fortiswitch_systemreplacemsg_trafficquota`
* **New Resource:** `fortiswitch_systemreplacemsg_utm`
* **New Resource:** `fortiswitch_systemreplacemsg_webproxy`
* **New Resource:** `fortiswitch_systemschedule_group`
* **New Resource:** `fortiswitch_systemschedule_onetime`
* **New Resource:** `fortiswitch_systemschedule_recurring`
* **New Resource:** `fortiswitch_systemsnmp_community`
* **New Resource:** `fortiswitch_systemsnmp_mibview`
* **New Resource:** `fortiswitch_systemsnmp_sysinfo`
* **New Resource:** `fortiswitch_systemsnmp_user`
* **New Resource:** `fortiswitch_user_adgrp`
* **New Resource:** `fortiswitch_user_certificate`
* **New Resource:** `fortiswitch_user_device`
* **New Resource:** `fortiswitch_user_deviceaccesslist`
* **New Resource:** `fortiswitch_user_devicecategory`
* **New Resource:** `fortiswitch_user_devicegroup`
* **New Resource:** `fortiswitch_user_domaincontroller`
* **New Resource:** `fortiswitch_user_exchange`
* **New Resource:** `fortiswitch_user_fortitoken`
* **New Resource:** `fortiswitch_user_fsso`
* **New Resource:** `fortiswitch_user_fssopolling`
* **New Resource:** `fortiswitch_user_group`
* **New Resource:** `fortiswitch_user_krbkeytab`
* **New Resource:** `fortiswitch_user_ldap`
* **New Resource:** `fortiswitch_user_local`
* **New Resource:** `fortiswitch_user_nacpolicy`
* **New Resource:** `fortiswitch_user_passwordpolicy`
* **New Resource:** `fortiswitch_user_peer`
* **New Resource:** `fortiswitch_user_peergrp`
* **New Resource:** `fortiswitch_user_pop3`
* **New Resource:** `fortiswitch_user_quarantine`
* **New Resource:** `fortiswitch_user_radius`
* **New Resource:** `fortiswitch_user_saml`
* **New Resource:** `fortiswitch_user_securityexemptlist`
* **New Resource:** `fortiswitch_user_setting`
* **New Resource:** `fortiswitch_user_tacacs`
* **New Resource:** `fortiswitch_videofilter_profile`
* **New Resource:** `fortiswitch_videofilter_youtubechannelfilter`
* **New Resource:** `fortiswitch_videofilter_youtubekey`
* **New Resource:** `fortiswitch_voip_profile`
* **New Resource:** `fortiswitch_vpn_ipsec_phase1interface`
* **New Resource:** `fortiswitch_vpn_ipsec_phase2interface`
* **New Resource:** `fortiswitch_vpn_l2tp`
* **New Resource:** `fortiswitch_vpn_ocvpn`
* **New Resource:** `fortiswitch_vpn_pptp`
* **New Resource:** `fortiswitch_vpncertificate_ca`
* **New Resource:** `fortiswitch_vpncertificate_crl`
* **New Resource:** `fortiswitch_vpncertificate_local`
* **New Resource:** `fortiswitch_vpncertificate_ocspserver`
* **New Resource:** `fortiswitch_vpncertificate_remote`
* **New Resource:** `fortiswitch_vpncertificate_setting`
* **New Resource:** `fortiswitch_vpnipsec_concentrator`
* **New Resource:** `fortiswitch_vpnipsec_fec`
* **New Resource:** `fortiswitch_vpnipsec_forticlient`
* **New Resource:** `fortiswitch_vpnipsec_manualkey`
* **New Resource:** `fortiswitch_vpnipsec_manualkeyinterface`
* **New Resource:** `fortiswitch_vpnipsec_phase1`
* **New Resource:** `fortiswitch_vpnipsec_phase1interface`
* **New Resource:** `fortiswitch_vpnipsec_phase2`
* **New Resource:** `fortiswitch_vpnipsec_phase2interface`
* **New Resource:** `fortiswitch_vpnssl_client`
* **New Resource:** `fortiswitch_vpnssl_settings`
* **New Resource:** `fortiswitch_vpnsslweb_hostchecksoftware`
* **New Resource:** `fortiswitch_vpnsslweb_portal`
* **New Resource:** `fortiswitch_vpnsslweb_realm`
* **New Resource:** `fortiswitch_vpnsslweb_userbookmark`
* **New Resource:** `fortiswitch_vpnsslweb_usergroupbookmark`
* **New Resource:** `fortiswitch_waf_mainclass`
* **New Resource:** `fortiswitch_waf_profile`
* **New Resource:** `fortiswitch_waf_signature`
* **New Resource:** `fortiswitch_waf_subclass`
* **New Resource:** `fortiswitch_wanopt_authgroup`
* **New Resource:** `fortiswitch_wanopt_cacheservice`
* **New Resource:** `fortiswitch_wanopt_contentdeliverynetworkrule`
* **New Resource:** `fortiswitch_wanopt_peer`
* **New Resource:** `fortiswitch_wanopt_profile`
* **New Resource:** `fortiswitch_wanopt_remotestorage`
* **New Resource:** `fortiswitch_wanopt_settings`
* **New Resource:** `fortiswitch_wanopt_webcache`
* **New Resource:** `fortiswitch_webfilter_content`
* **New Resource:** `fortiswitch_webfilter_contentheader`
* **New Resource:** `fortiswitch_webfilter_fortiguard`
* **New Resource:** `fortiswitch_webfilter_ftgdlocalcat`
* **New Resource:** `fortiswitch_webfilter_ftgdlocalrating`
* **New Resource:** `fortiswitch_webfilter_ipsurlfiltercachesetting`
* **New Resource:** `fortiswitch_webfilter_ipsurlfiltersetting`
* **New Resource:** `fortiswitch_webfilter_ipsurlfiltersetting6`
* **New Resource:** `fortiswitch_webfilter_override`
* **New Resource:** `fortiswitch_webfilter_profile`
* **New Resource:** `fortiswitch_webfilter_searchengine`
* **New Resource:** `fortiswitch_webfilter_urlfilter`
* **New Resource:** `fortiswitch_webproxy_debugurl`
* **New Resource:** `fortiswitch_webproxy_explicit`
* **New Resource:** `fortiswitch_webproxy_forwardserver`
* **New Resource:** `fortiswitch_webproxy_forwardservergroup`
* **New Resource:** `fortiswitch_webproxy_global`
* **New Resource:** `fortiswitch_webproxy_profile`
* **New Resource:** `fortiswitch_webproxy_urlmatch`
* **New Resource:** `fortiswitch_webproxy_wisp`
* **New Resource:** `fortiswitch_wirelesscontroller_accesscontrollist`
* **New Resource:** `fortiswitch_wirelesscontroller_address`
* **New Resource:** `fortiswitch_wirelesscontroller_addrgrp`
* **New Resource:** `fortiswitch_wirelesscontroller_apcfgprofile`
* **New Resource:** `fortiswitch_wirelesscontroller_apstatus`
* **New Resource:** `fortiswitch_wirelesscontroller_arrpprofile`
* **New Resource:** `fortiswitch_wirelesscontroller_bleprofile`
* **New Resource:** `fortiswitch_wirelesscontroller_bonjourprofile`
* **New Resource:** `fortiswitch_wirelesscontroller_global`
* **New Resource:** `fortiswitch_wirelesscontroller_intercontroller`
* **New Resource:** `fortiswitch_wirelesscontroller_log`
* **New Resource:** `fortiswitch_wirelesscontroller_mpskprofile`
* **New Resource:** `fortiswitch_wirelesscontroller_nacprofile`
* **New Resource:** `fortiswitch_wirelesscontroller_qosprofile`
* **New Resource:** `fortiswitch_wirelesscontroller_region`
* **New Resource:** `fortiswitch_wirelesscontroller_setting`
* **New Resource:** `fortiswitch_wirelesscontroller_snmp`
* **New Resource:** `fortiswitch_wirelesscontroller_ssidpolicy`
* **New Resource:** `fortiswitch_wirelesscontroller_syslogprofile`
* **New Resource:** `fortiswitch_wirelesscontroller_timers`
* **New Resource:** `fortiswitch_wirelesscontroller_utmprofile`
* **New Resource:** `fortiswitch_wirelesscontroller_vap`
* **New Resource:** `fortiswitch_wirelesscontroller_vapgroup`
* **New Resource:** `fortiswitch_wirelesscontroller_wagprofile`
* **New Resource:** `fortiswitch_wirelesscontroller_widsprofile`
* **New Resource:** `fortiswitch_wirelesscontroller_wtp`
* **New Resource:** `fortiswitch_wirelesscontroller_wtpgroup`
* **New Resource:** `fortiswitch_wirelesscontroller_wtpprofile`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_anqp3gppcellular`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_anqpipaddresstype`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_anqpnairealm`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_anqpnetworkauthtype`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_anqproamingconsortium`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_anqpvenuename`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_anqpvenueurl`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_h2qpadviceofcharge`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_h2qpconncapability`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_h2qpoperatorname`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_h2qposuprovider`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_h2qposuprovidernai`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_h2qptermsandconditions`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_h2qpwanmetric`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_hsprofile`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_icon`
* **New Resource:** `fortiswitch_wirelesscontrollerhotspot20_qosmap`
* **New Data Source:** `fortiswitch_json_generic_api`
* **New Data Source:** `fortiswitch_router_accesslist`
* **New Data Source:** `fortiswitch_router_accesslist6`
* **New Data Source:** `fortiswitch_router_aspathlist`
* **New Data Source:** `fortiswitch_router_authpath`
* **New Data Source:** `fortiswitch_router_bgp`
* **New Data Source:** `fortiswitch_router_communitylist`
* **New Data Source:** `fortiswitch_router_isis`
* **New Data Source:** `fortiswitch_router_keychain`
* **New Data Source:** `fortiswitch_router_multicast`
* **New Data Source:** `fortiswitch_router_multicastflow`
* **New Data Source:** `fortiswitch_router_ospf`
* **New Data Source:** `fortiswitch_router_ospf6`
* **New Data Source:** `fortiswitch_router_policy`
* **New Data Source:** `fortiswitch_router_prefixlist`
* **New Data Source:** `fortiswitch_router_prefixlist6`
* **New Data Source:** `fortiswitch_router_rip`
* **New Data Source:** `fortiswitch_router_ripng`
* **New Data Source:** `fortiswitch_router_routemap`
* **New Data Source:** `fortiswitch_router_setting`
* **New Data Source:** `fortiswitch_router_static`
* **New Data Source:** `fortiswitch_router_static6`
* **New Data Source:** `fortiswitch_system_accprofile`
* **New Data Source:** `fortiswitch_system_admin`
* **New Data Source:** `fortiswitch_system_arptable`
* **New Data Source:** `fortiswitch_system_autoscript`
* **New Data Source:** `fortiswitch_system_automationaction`
* **New Data Source:** `fortiswitch_system_automationdestination`
* **New Data Source:** `fortiswitch_system_automationtrigger`
* **New Data Source:** `fortiswitch_system_centralmanagement`
* **New Data Source:** `fortiswitch_system_console`
* **New Data Source:** `fortiswitch_system_dns`
* **New Data Source:** `fortiswitch_system_dnsdatabase`
* **New Data Source:** `fortiswitch_system_dnsserver`
* **New Data Source:** `fortiswitch_system_emailserver`
* **New Data Source:** `fortiswitch_system_fm`
* **New Data Source:** `fortiswitch_system_fortiguard`
* **New Data Source:** `fortiswitch_system_fortimanager`
* **New Data Source:** `fortiswitch_system_global`
* **New Data Source:** `fortiswitch_system_interface`
* **New Data Source:** `fortiswitch_system_ipv6neighborcache`
* **New Data Source:** `fortiswitch_system_linkmonitor`
* **New Data Source:** `fortiswitch_system_managementtunnel`
* **New Data Source:** `fortiswitch_system_ntp`
* **New Data Source:** `fortiswitch_system_passwordpolicy`
* **New Data Source:** `fortiswitch_system_proxyarp`
* **New Data Source:** `fortiswitch_system_resourcelimits`
* **New Data Source:** `fortiswitch_system_sessionttl`
* **New Data Source:** `fortiswitch_system_sflow`
* **New Data Source:** `fortiswitch_system_tosbasedpriority`
* **New Data Source:** `fortiswitch_system_vxlan`
* **New Data Source:** `fortiswitch_system_zone`
* **New Data Source:** `fortiswitch_systemautoupdate_pushupdate`
* **New Data Source:** `fortiswitch_systemautoupdate_schedule`
* **New Data Source:** `fortiswitch_systemautoupdate_tunneling`
* **New Data Source:** `fortiswitch_systemdhcp_server`
* **New Data Source:** `fortiswitch_systemsnmp_community`
* **New Data Source:** `fortiswitch_systemsnmp_sysinfo`
* **New Data Source:** `fortiswitch_systemsnmp_user`
* **New Data Source:** `fortiswitch_routerbgp_neighbor`
* **New Data Source:** `fortiswitch_router_accesslistlist`
* **New Data Source:** `fortiswitch_router_accesslist6list`
* **New Data Source:** `fortiswitch_router_aspathlistlist`
* **New Data Source:** `fortiswitch_router_authpathlist`
* **New Data Source:** `fortiswitch_router_communitylistlist`
* **New Data Source:** `fortiswitch_router_keychainlist`
* **New Data Source:** `fortiswitch_router_multicastflowlist`
* **New Data Source:** `fortiswitch_router_policylist`
* **New Data Source:** `fortiswitch_router_prefixlistlist`
* **New Data Source:** `fortiswitch_router_prefixlist6list`
* **New Data Source:** `fortiswitch_router_routemaplist`
* **New Data Source:** `fortiswitch_router_staticlist`
* **New Data Source:** `fortiswitch_router_static6list`
* **New Data Source:** `fortiswitch_system_accprofilelist`
* **New Data Source:** `fortiswitch_system_adminlist`
* **New Data Source:** `fortiswitch_system_arptablelist`
* **New Data Source:** `fortiswitch_system_autoscriptlist`
* **New Data Source:** `fortiswitch_system_automationactionlist`
* **New Data Source:** `fortiswitch_system_automationdestinationlist`
* **New Data Source:** `fortiswitch_system_automationtriggerlist`
* **New Data Source:** `fortiswitch_system_dnsdatabaselist`
* **New Data Source:** `fortiswitch_system_dnsserverlist`
* **New Data Source:** `fortiswitch_system_interfacelist`
* **New Data Source:** `fortiswitch_system_ipv6neighborcachelist`
* **New Data Source:** `fortiswitch_system_linkmonitorlist`
* **New Data Source:** `fortiswitch_system_proxyarplist`
* **New Data Source:** `fortiswitch_system_tosbasedprioritylist`
* **New Data Source:** `fortiswitch_system_vxlanlist`
* **New Data Source:** `fortiswitch_system_zonelist`
* **New Data Source:** `fortiswitch_systemdhcp_serverlist`
* **New Data Source:** `fortiswitch_systemsnmp_communitylist`
* **New Data Source:** `fortiswitch_systemsnmp_userlist`
* **New Data Source:** `fortiswitch_routerbgp_neighborlist`