<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_serviceenginegroup"
sidebar_current: "docs-avi-datasource-serviceenginegroup"
description: |-
  Get information of Avi ServiceEngineGroup.
---

# avi_serviceenginegroup

This data source is used to to get avi_serviceenginegroup objects.

## Example Usage

```hcl
data "avi_serviceenginegroup" "foo_serviceenginegroup" {
    uuid = "serviceenginegroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search ServiceEngineGroup by name.
* `uuid` - (Optional) Search ServiceEngineGroup by uuid.
* `cloud_ref` - (Optional) Search ServiceEngineGroup by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `accelerated_networking` - Enable accelerated networking option for azure se. Accelerated networking enables single root i/o virtualization (sr-iov) to a se vm. This improves networking performance. Field introduced in 17.2.14,18.1.5,18.2.1.
* `active_standby` - Service engines in active/standby mode for ha failover.
* `aggressive_failure_detection` - Enable aggressive failover configuration for ha. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `algo` - In compact placement, virtual services are placed on existing ses until max_vs_per_se limit is reached. Enum options - PLACEMENT_ALGO_PACKED, PLACEMENT_ALGO_DISTRIBUTED.
* `allow_burst` - Allow ses to be created using burst license. Field introduced in 17.2.5.
* `app_cache_percent` - A percent value of total se memory reserved for applicationcaching. This is an se bootup property and requires se restart.requires se reboot. Allowed values are 0 - 100. Special values are 0- 'disable'. Field introduced in 18.2.3. Unit is percent. Allowed in basic(allowed values- 0) edition, essentials(allowed values- 0) edition, enterprise edition. Special default for basic edition is 0, essentials edition is 0, enterprise is 10.
* `app_cache_threshold` - The max memory that can be allocated for the app cache. This value will act as an upper bound on the cache size specified in app_cache_percent. Special values are 0- 'disable'. Field introduced in 20.1.1. Unit is gb.
* `app_learning_memory_percent` - A percent value of total se memory reserved for application learning. This is an se bootup property and requires se restart. Allowed values are 0 - 10. Field introduced in 18.2.3. Unit is percent.
* `archive_shm_limit` - Amount of se memory in gb until which shared memory is collected in core archive. Field introduced in 17.1.3. Unit is gb.
* `async_ssl` - Ssl handshakes will be handled by dedicated ssl threads.requires se reboot. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `async_ssl_threads` - Number of async ssl threads per se_dp.requires se reboot. Allowed values are 1-16.
* `auto_rebalance` - If set, virtual services will be automatically migrated when load on an se is less than minimum or more than maximum thresholds. Only alerts are generated when the auto_rebalance is not set. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `auto_rebalance_capacity_per_se` - Capacities of se for auto rebalance for each criteria. Field introduced in 17.2.4.
* `auto_rebalance_criteria` - Set of criteria for se auto rebalance. Enum options - SE_AUTO_REBALANCE_CPU, SE_AUTO_REBALANCE_PPS, SE_AUTO_REBALANCE_MBPS, SE_AUTO_REBALANCE_OPEN_CONNS, SE_AUTO_REBALANCE_CPS. Field introduced in 17.2.3.
* `auto_rebalance_interval` - Frequency of rebalance, if 'auto rebalance' is enabled. Unit is sec.
* `auto_redistribute_active_standby_load` - Redistribution of virtual services from the takeover se to the replacement se can cause momentary traffic loss. If the auto-redistribute load option is left in its default off state, any desired rebalancing requires calls to rest api. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `availability_zone_refs` - Availability zones for virtual service high availability. It is a reference to an object of type availabilityzone. Field introduced in 20.1.1.
* `bgp_peer_monitor_failover_enabled` - Enable bgp peer monitoring based failover. Field introduced in 21.1.3.
* `bgp_state_update_interval` - Bgp peer state update interval. Allowed values are 5-100. Field introduced in 17.2.14,18.1.5,18.2.1. Unit is sec.
* `buffer_se` - Excess service engine capacity provisioned for ha failover.
* `cloud_ref` - It is a reference to an object of type cloud.
* `compress_ip_rules_for_each_ns_subnet` - Compress ip rules into a single subnet based ip rule for each north-south ipam subnet configured in pcap mode in openshift/kubernetes node. Field introduced in 18.2.9, 20.1.1.
* `config_debugs_on_all_cores` - Enable config debugs on all cores of se. Field introduced in 17.2.13,18.1.5,18.2.1.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `connection_memory_percentage` - Percentage of memory for connection state. This will come at the expense of memory used for http in-memory cache. Allowed values are 10-90. Unit is percent.
* `core_shm_app_cache` - Include shared memory for app cache in core file.requires se reboot. Field introduced in 18.2.8, 20.1.1.
* `core_shm_app_learning` - Include shared memory for app learning in core file.requires se reboot. Field introduced in 18.2.8, 20.1.1.
* `cpu_reserve` - Boolean flag to set cpu_reserve.
* `cpu_socket_affinity` - Allocate all the cpu cores for the service engine virtual machines  on the same cpu socket. Applicable only for vcenter cloud.
* `custom_securitygroups_data` - Custom security groups to be associated with data vnics for se instances in openstack and aws clouds. Field introduced in 17.1.3.
* `custom_securitygroups_mgmt` - Custom security groups to be associated with management vnic for se instances in openstack and aws clouds. Field introduced in 17.1.3.
* `custom_tag` - Custom tag will be used to create the tags for se instance in aws. Note this is not the same as the prefix for se name.
* `data_network_id` - Subnet used to spin up the data nic for service engines, used only for azure cloud. Overrides the cloud level setting for service engine subnet. Field introduced in 18.2.3.
* `datascript_timeout` - Number of instructions before datascript times out. Allowed values are 0-100000000. Field introduced in 18.2.3.
* `deactivate_ipv6_discovery` - If activated, ipv6 address and route discovery are deactivated.requires se reboot. Field introduced in 21.1.1.
* `dedicated_dispatcher_core` - Dedicate the core that handles packet receive/transmit from the network to just the dispatching function. Don't use it for tcp/ip and ssl functions.
* `description` - User defined description for the object.
* `disable_avi_securitygroups` - By default, avi creates and manages security groups along with custom sg provided by user. Set this to true to disallow avi to create and manage new security groups. Avi will only make use of custom security groups provided by user. This option is supported for aws and openstack cloud types. Field introduced in 17.2.13,18.1.4,18.2.1.
* `disable_csum_offloads` - Stop using tcp/udp and ip checksum offload features of nics. Field introduced in 17.1.14, 17.2.5, 18.1.1.
* `disable_flow_probes` - Disable flow probes for scaled out vs'es. Field introduced in 20.1.3.
* `disable_gro` - Disable generic receive offload (gro) in dpdk poll-mode driver packet receive path. Gro is on by default on nics that do not support lro (large receive offload) or do not gain performance boost from lro. Field introduced in 17.2.5, 18.1.1.
* `disable_se_memory_check` - If set, disable the config memory check done in service engine. Field introduced in 18.1.2.
* `disable_tso` - Disable tcp segmentation offload (tso) in dpdk poll-mode driver packet transmit path. Tso is on by default on nics that support it. Field introduced in 17.2.5, 18.1.1.
* `disk_per_se` - Amount of disk space for each of the service engine virtual machines. Unit is gb.
* `distribute_load_active_standby` - Use both the active and standby service engines for virtual service placement in the legacy active standby ha mode. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `distribute_queues` - Distributes queue ownership among cores so multiple cores handle dispatcher duties. Requires se reboot. Deprecated from 18.2.8, instead use max_queues_per_vnic. Field introduced in 17.2.8. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `distribute_vnics` - Distributes vnic ownership among cores so multiple cores handle dispatcher duties.requires se reboot. Field introduced in 18.2.5. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `downstream_send_timeout` - Timeout for downstream to become writable. Field introduced in 21.1.1. Unit is milliseconds.
* `dp_aggressive_deq_interval_msec` - Dequeue interval for receive queue from se_dp in aggressive mode. Allowed values are 1-1000. Field introduced in 21.1.1. Unit is milliseconds. Allowed in basic edition, essentials edition, enterprise edition.
* `dp_aggressive_enq_interval_msec` - Enqueue interval for request queue to se_dp in aggressive mode. Allowed values are 1-1000. Field introduced in 21.1.1. Unit is milliseconds. Allowed in basic edition, essentials edition, enterprise edition.
* `dp_aggressive_hb_frequency` - Frequency of se - se hb messages when aggressive failure mode detection is enabled. Field introduced in 20.1.3. Unit is milliseconds.
* `dp_aggressive_hb_timeout_count` - Consecutive hb failures after which failure is reported to controller,when aggressive failure mode detection is enabled. Field introduced in 20.1.3.
* `dp_deq_interval_msec` - Dequeue interval for receive queue from se_dp. Allowed values are 1-1000. Field introduced in 21.1.1. Unit is milliseconds. Allowed in basic edition, essentials edition, enterprise edition.
* `dp_enq_interval_msec` - Enqueue interval for request queue to se_dp. Allowed values are 1-1000. Field introduced in 21.1.1. Unit is milliseconds. Allowed in basic edition, essentials edition, enterprise edition.
* `dp_hb_frequency` - Frequency of se - se hb messages when aggressive failure mode detection is not enabled. Field introduced in 20.1.3. Unit is milliseconds.
* `dp_hb_timeout_count` - Consecutive hb failures after which failure is reported to controller, when aggressive failure mode detection is not enabled. Field introduced in 20.1.3.
* `enable_gratarp_permanent` - Enable gratarp for vip_ip. Field introduced in 18.2.3.
* `enable_hsm_log` - Enable hsm luna engine logs. Field introduced in 21.1.1.
* `enable_hsm_priming` - (this is a beta feature). Enable hsm key priming. If enabled, key handles on the hsm will be synced to se before processing client connections. Field introduced in 17.2.7, 18.1.1.
* `enable_multi_lb` - Applicable only for azure cloud with basic sku lb. If set, additional azure lbs will be automatically created if resources in existing lb are exhausted. Field introduced in 17.2.10, 18.1.2.
* `enable_pcap_tx_ring` - Enable tx ring support in pcap mode of operation. Tso feature is not supported with tx ring enabled. Deprecated from 18.2.8, instead use pcap_tx_mode. Requires se reboot. Field introduced in 18.2.5.
* `ephemeral_portrange_end` - End local ephemeral port number for outbound connections. Field introduced in 17.2.13, 18.1.5, 18.2.1.
* `ephemeral_portrange_start` - Start local ephemeral port number for outbound connections. Field introduced in 17.2.13, 18.1.5, 18.2.1.
* `extra_config_multiplier` - Multiplier for extra config to support large vs/pool config.
* `extra_shared_config_memory` - Extra config memory to support large geo db configuration. Field introduced in 17.1.1. Unit is mb.
* `flow_table_new_syn_max_entries` - Maximum number of flow table entries that have not completed tcp three-way handshake yet. Field introduced in 17.2.5.
* `free_list_size` - Number of entries in the free list. Field introduced in 17.2.10, 18.1.2.
* `gcp_config` - Google cloud platform, service engine group configuration. Field introduced in 20.1.3.
* `gratarp_permanent_periodicity` - Gratarp periodicity for vip-ip. Allowed values are 5-30. Field introduced in 18.2.3. Unit is min.
* `ha_mode` - High availability mode for all the virtual services using this service engine group. Enum options - HA_MODE_SHARED_PAIR, HA_MODE_SHARED, HA_MODE_LEGACY_ACTIVE_STANDBY. Allowed in basic(allowed values- ha_mode_legacy_active_standby) edition, essentials(allowed values- ha_mode_legacy_active_standby) edition, enterprise edition. Special default for basic edition is ha_mode_legacy_active_standby, essentials edition is ha_mode_legacy_active_standby, enterprise is ha_mode_shared.
* `handle_per_pkt_attack` - Configuration to handle per packet attack handling.for example, dns reflection attack is a type of attack where a response packet is sent to the dns vs.this configuration tells if such packets should be dropped without further processing. Field introduced in 20.1.3.
* `hardwaresecuritymodulegroup_ref` - It is a reference to an object of type hardwaresecuritymodulegroup.
* `heap_minimum_config_memory` - Minimum required heap memory to apply any configuration. Allowed values are 0-100. Field introduced in 18.1.2. Unit is mb.
* `hm_on_standby` - Enable active health monitoring from the standby se for all placed virtual services. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition. Special default for basic edition is false, essentials edition is false, enterprise is true.
* `host_attribute_key` - Key of a (key, value) pair identifying a label for a set of nodes usually in container clouds. Needs to be specified together with host_attribute_value. Ses can be configured differently including ha modes across different se groups. May also be used for isolation between different classes of virtualservices. Virtualservices' se group may be specified via annotations/labels. A openshift/kubernetes namespace maybe annotated with a matching se group label as openshift.io/node-selector  apptype=prod. When multiple se groups are used in a cloud with host attributes specified,just a single se group can exist as a match-all se group without a host_attribute_key.
* `host_attribute_value` - Value of a (key, value) pair identifying a label for a set of nodes usually in container clouds. Needs to be specified together with host_attribute_key.
* `host_gateway_monitor` - Enable the host gateway monitor when service engine is deployed as docker container. Disabled by default. Field introduced in 17.2.4.
* `http_rum_console_log` - Enable javascript console logs on the client browser when collecting client insights. Field introduced in 21.1.1. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `http_rum_min_content_length` - Minimum response size content length to sample for client insights. Field introduced in 21.1.1. Allowed in basic(allowed values- 64) edition, essentials(allowed values- 64) edition, enterprise edition.
* `hypervisor` - Override default hypervisor. Enum options - DEFAULT, VMWARE_ESX, KVM, VMWARE_VSAN, XEN.
* `ignore_rtt_threshold` - Ignore rtt samples if it is above threshold. Field introduced in 17.1.6,17.2.2. Unit is milliseconds.
* `ingress_access_data` - Program se security group ingress rules to allow vip data access from remote cidr type. Enum options - SG_INGRESS_ACCESS_NONE, SG_INGRESS_ACCESS_ALL, SG_INGRESS_ACCESS_VPC. Field introduced in 17.1.5.
* `ingress_access_mgmt` - Program se security group ingress rules to allow ssh/icmp management access from remote cidr type. Enum options - SG_INGRESS_ACCESS_NONE, SG_INGRESS_ACCESS_ALL, SG_INGRESS_ACCESS_VPC. Field introduced in 17.1.5.
* `instance_flavor` - Instance/flavor name for se instance.
* `instance_flavor_info` - Additional information associated with instance_flavor. Field introduced in 20.1.1.
* `iptables` - Iptable rules. Maximum of 128 items allowed.
* `l7_conns_per_core` - Number of l7 connections that can be cached per core. Field introduced in 21.1.1.
* `l7_resvd_listen_conns_per_core` - Number of reserved l7 listener connections per core. Field introduced in 21.1.1.
* `labels` - Labels associated with this se group. Field introduced in 20.1.1. Maximum of 1 items allowed.
* `lbaction_num_requests_to_dispatch` - Number of requests to dispatch from the request. Queue at a regular interval. Field introduced in 21.1.1.
* `lbaction_rq_per_request_max_retries` - Maximum retries per request in the request queue. Field introduced in 21.1.1.
* `least_load_core_selection` - Select core with least load for new flow.
* `license_tier` - Specifies the license tier which would be used. This field by default inherits the value from cloud. Enum options - ENTERPRISE_16, ENTERPRISE, ENTERPRISE_18, BASIC, ESSENTIALS, SAAS. Field introduced in 17.2.5.
* `license_type` - If no license type is specified then default license enforcement for the cloud type is chosen. Enum options - LIC_BACKEND_SERVERS, LIC_SOCKETS, LIC_CORES, LIC_HOSTS, LIC_SE_BANDWIDTH, LIC_METERED_SE_BANDWIDTH. Field introduced in 17.2.5.
* `log_agent_compress_logs` - Flag to indicate if log files are compressed upon full on the service engine. Field introduced in 21.1.1.
* `log_agent_debug_enabled` - Enable debug logs by default on service engine. This includes all other debugging logs. Debug logs can also be explcitly enabled from the cli shell. Field introduced in 21.1.1.
* `log_agent_file_sz_appl` - Maximum application log file size before rollover. Field introduced in 21.1.1.
* `log_agent_file_sz_conn` - Maximum connection log file size before rollover. Field introduced in 21.1.1.
* `log_agent_file_sz_debug` - Maximum debug log file size before rollover. Field introduced in 21.1.1.
* `log_agent_file_sz_event` - Maximum event log file size before rollover. Field introduced in 21.1.1.
* `log_agent_log_storage_min_sz` - Minimum storage allocated for logs irrespective of memory and cores. Field introduced in 21.1.1. Unit is mb.
* `log_agent_max_concurrent_rsync` - Maximum concurrent rsync requests initiated from log-agent to the controller. Field introduced in 21.1.1.
* `log_agent_max_storage_excess_percent` - Excess percentage threshold of disk size to trigger cleanup of logs on the service engine. Field introduced in 21.1.1.
* `log_agent_max_storage_ignore_percent` - Maximum storage on the disk not allocated for logs on the service engine. Field introduced in 21.1.1.
* `log_agent_min_storage_per_vs` - Minimum storage allocated to any given virtualservice on the service engine. Field introduced in 21.1.1.
* `log_agent_sleep_interval` - Internal timer to stall log-agent and prevent it from hogging cpu cycles on the service engine. Field introduced in 21.1.1. Unit is milliseconds.
* `log_agent_trace_enabled` - Enable trace logs by default on service engine. Configuration operations are logged along with other important logs by service engine. Field introduced in 21.1.1.
* `log_agent_unknown_vs_timer` - Timeout to purge unknown virtual service logs from the service engine. Field introduced in 21.1.1. Unit is sec.
* `log_disksz` - Maximum disk capacity (in mb) to be allocated to an se. This is exclusively used for debug and log data. Unit is mb.
* `log_malloc_failure` - Se will log memory allocation related failure to the se_trace file, wherever available. Field introduced in 20.1.2. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition.
* `log_message_max_file_list_size` - Maximum number of file names in a log message. Field introduced in 21.1.1.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.7. Allowed in basic edition, essentials edition, enterprise edition.
* `max_concurrent_external_hm` - Maximum number of external health monitors that can run concurrently in a service engine. This helps control the cpu and memory use by external health monitors. Special values are 0- 'value will be internally calculated based on cpu and memory'. Field introduced in 18.2.7.
* `max_cpu_usage` - When cpu usage on an se exceeds this threshold, virtual services hosted on this se may be rebalanced to other ses to reduce load. A new se may be created as part of this process. Allowed values are 40-90. Unit is percent.
* `max_memory_per_mempool` - Max bytes that can be allocated in a single mempool. Field introduced in 18.1.5. Unit is mb.
* `max_num_se_dps` - Configures the maximum number of se_dp processes that handles traffic. If not configured, defaults to the number of cpus on the se. If decreased, it will only take effect after se reboot. Allowed values are 1-128. Field introduced in 20.1.1. Allowed in basic(allowed values- 0) edition, essentials(allowed values- 0) edition, enterprise edition.
* `max_public_ips_per_lb` - Applicable to azure platform only. Maximum number of public ips per azure lb. Field introduced in 17.2.12, 18.1.2.
* `max_queues_per_vnic` - Maximum number of queues per vnic setting to '0' utilises all queues that are distributed across dispatcher cores. Allowed values are 0,1,2,4,8,16. Field introduced in 18.2.7, 20.1.1. Allowed in basic(allowed values- 1) edition, essentials(allowed values- 1) edition, enterprise edition.
* `max_rules_per_lb` - Applicable to azure platform only. Maximum number of rules per azure lb. Field introduced in 17.2.12, 18.1.2.
* `max_scaleout_per_vs` - Maximum number of active service engines for the virtual service. Allowed values are 1-64.
* `max_se` - Maximum number of services engines in this group. Allowed values are 0-1000.
* `max_skb_frags` - Maximum of number of 4 kb pages allocated to the linux kernel gro subsystem for packet coalescing. This parameter is limited to supported kernels only. Requires se reboot. Allowed values are 1-17. Field introduced in 21.1.3.
* `max_vs_per_se` - Maximum number of virtual services that can be placed on a single service engine. Allowed values are 1-1000.
* `mem_reserve` - Boolean flag to set mem_reserve.
* `memory_for_config_update` - Indicates the percent of memory reserved for config updates. Allowed values are 0-100. Field introduced in 18.1.2. Unit is percent.
* `memory_per_se` - Amount of memory for each of the service engine virtual machines. Changes to this setting do not affect existing ses.
* `mgmt_network_ref` - Management network to use for avi service engines. It is a reference to an object of type network.
* `mgmt_subnet` - Management subnet to use for avi service engines.
* `min_cpu_usage` - When cpu usage on an se falls below the minimum threshold, virtual services hosted on the se may be consolidated onto other underutilized ses. After consolidation, unused service engines may then be eligible for deletion. Allowed values are 20-60. Unit is percent.
* `min_scaleout_per_vs` - Minimum number of active service engines for the virtual service. Allowed values are 1-64.
* `min_se` - Minimum number of services engines in this group (relevant for se autorebalance only). Allowed values are 0-1000. Field introduced in 17.2.13,18.1.3,18.2.1.
* `minimum_connection_memory` - Indicates the percent of memory reserved for connections. Allowed values are 0-100. Field introduced in 18.1.2. Unit is percent.
* `n_log_streaming_threads` - Number of threads to use for log streaming. Allowed values are 1-100. Field introduced in 17.2.12, 18.1.2.
* `name` - Name of the object.
* `netlink_poller_threads` - Number of threads to poll for netlink messages excluding the thread for default namespace. Requires se reboot. Allowed values are 1-32. Field introduced in 21.1.1.
* `netlink_sock_buf_size` - Socket buffer size for the netlink sockets. Requires se reboot. Allowed values are 1-128. Field introduced in 21.1.1. Unit is mega_bytes.
* `ngx_free_connection_stack` - Free the connection stack. Field introduced in 21.1.1.
* `non_significant_log_throttle` - This setting limits the number of non-significant logs generated per second per core on this se. Default is 100 logs per second. Set it to zero (0) to deactivate throttling. Field introduced in 17.1.3. Unit is per_second.
* `ns_helper_deq_interval_msec` - Dequeue interval for receive queue from ns helper. Allowed values are 1-1000. Field introduced in 21.1.1. Unit is milliseconds. Allowed in basic edition, essentials edition, enterprise edition.
* `num_dispatcher_cores` - Number of dispatcher cores (0,1,2,4,8 or 16). If set to 0, then number of dispatcher cores is deduced automatically.requires se reboot. Allowed values are 0,1,2,4,8,16. Field introduced in 17.2.12, 18.1.3, 18.2.1. Allowed in basic(allowed values- 0) edition, essentials(allowed values- 0) edition, enterprise edition.
* `num_flow_cores_sum_changes_to_ignore` - Number of changes in num flow cores sum to ignore.
* `objsync_config` - Configuration knobs for interse object distribution. Field introduced in 20.1.3.
* `objsync_port` - Tcp port on se management interface for interse object distribution. Supported only for externally managed security groups. Not supported on full access deployments. Requires se reboot. Field introduced in 20.1.3.
* `openstack_availability_zones` - Field introduced in 17.1.1. Maximum of 5 items allowed.
* `openstack_mgmt_network_name` - Avi management network name.
* `openstack_mgmt_network_uuid` - Management network uuid.
* `os_reserved_memory` - Amount of extra memory to be reserved for use by the operating system on a service engine. Unit is mb.
* `pcap_tx_mode` - Determines the pcap transmit mode of operation. Requires se reboot. Enum options - PCAP_TX_AUTO, PCAP_TX_SOCKET, PCAP_TX_RING. Field introduced in 18.2.8, 20.1.1.
* `pcap_tx_ring_rd_balancing_factor` - In pcap mode, reserve a configured portion of tx ring resources for itself and the remaining portion for the rx ring to achieve better balance in terms of queue depth. Requires se reboot. Allowed values are 10-100. Field introduced in 20.1.3. Unit is percent.
* `per_app` - Per-app se mode is designed for deploying dedicated load balancers per app (vs). In this mode, each se is limited to a max of 2 vss. Vcpus in per-app ses count towards licensing usage at 25% rate. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `per_vs_admission_control` - Enable/disable per vs level admission control.enabling this feature will cause the connection and packet throttling on a particular vs that has high packet buffer consumption. Field introduced in 20.1.3.
* `placement_mode` - If placement mode is 'auto', virtual services are automatically placed on service engines. Enum options - PLACEMENT_MODE_AUTO.
* `realtime_se_metrics` - Enable or deactivate real time se metrics.
* `reboot_on_panic` - Reboot the vm or host on kernel panic. Field introduced in 18.2.5.
* `resync_time_interval` - Time interval to re-sync se's time with wall clock time. Allowed values are 8-600000. Field introduced in 20.1.1. Unit is milliseconds.
* `sdb_flush_interval` - Sdb pipeline flush interval. Allowed values are 1-10000. Field introduced in 21.1.1. Unit is milliseconds. Allowed in basic edition, essentials edition, enterprise edition.
* `sdb_pipeline_size` - Sdb pipeline size. Allowed values are 1-10000. Field introduced in 21.1.1. Allowed in basic edition, essentials edition, enterprise edition.
* `sdb_scan_count` - Sdb scan count. Allowed values are 1-1000. Field introduced in 21.1.1. Allowed in basic edition, essentials edition, enterprise edition.
* `se_bandwidth_type` - Select the se bandwidth for the bandwidth license. Enum options - SE_BANDWIDTH_UNLIMITED, SE_BANDWIDTH_25M, SE_BANDWIDTH_200M, SE_BANDWIDTH_1000M, SE_BANDWIDTH_10000M. Field introduced in 17.2.5. Allowed in basic(allowed values- se_bandwidth_unlimited) edition, essentials(allowed values- se_bandwidth_unlimited) edition, enterprise edition.
* `se_delayed_flow_delete` - Delay the cleanup of flowtable entry. To be used under surveillance of avi support. Field introduced in 20.1.2. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition.
* `se_deprovision_delay` - Duration to preserve unused service engine virtual machines before deleting them. If traffic to a virtual service were to spike up abruptly, this se would still be available to be utilized again rather than creating a new se. If this value is set to 0, controller will never delete any ses and administrator has to manually cleanup unused ses. Allowed values are 0-525600. Unit is min.
* `se_dos_profile` - Dict settings for serviceenginegroup.
* `se_dp_hm_drops` - Internal only. Used to simulate se - se hb failure. Field introduced in 20.1.3.
* `se_dp_isolation` - Toggle support to run se datapath instances in isolation on exclusive cpus. This improves latency and performance. However, this could reduce the total number of se_dp instances created on that se instance. Supported for >= 8 cpus. Requires se reboot. Field introduced in 20.1.4.
* `se_dp_isolation_num_non_dp_cpus` - Number of cpus for non se-dp tasks in se datapath isolation mode. Translates total cpus minus 'num_non_dp_cpus' for datapath use. It is recommended to reserve an even number of cpus for hyper-threaded processors. Requires se reboot. Allowed values are 1-8. Special values are 0- 'auto'. Field introduced in 20.1.4.
* `se_dp_log_nf_enqueue_percent` - Internal buffer full indicator on the service engine beyond which the unfiltered logs are abandoned. Field introduced in 21.1.1.
* `se_dp_log_udf_enqueue_percent` - Internal buffer full indicator on the service engine beyond which the user filtered logs are abandoned. Field introduced in 21.1.1.
* `se_dp_max_hb_version` - The highest supported se-se heartbeat protocol version. This version is reported by secondary se to primary se in heartbeat response messages. Allowed values are 1-3. Field introduced in 20.1.1.
* `se_dp_vnic_queue_stall_event_sleep` - Time (in seconds) service engine waits for after generating a vnic transmit queue stall event before resetting thenic. Field introduced in 18.2.5.
* `se_dp_vnic_queue_stall_threshold` - Number of consecutive transmit failures to look for before generating a vnic transmit queue stall event. Field introduced in 18.2.5.
* `se_dp_vnic_queue_stall_timeout` - Time (in milliseconds) to wait for network/nic recovery on detecting a transmit queue stall after which service engine resets the nic. Field introduced in 18.2.5.
* `se_dp_vnic_restart_on_queue_stall_count` - Number of consecutive transmit queue stall events in se_dp_vnic_stall_se_restart_window to look for before restarting se. Field introduced in 18.2.5.
* `se_dp_vnic_stall_se_restart_window` - Window of time (in seconds) during which se_dp_vnic_restart_on_queue_stall_count number of consecutive stalls results in a se restart. Field introduced in 18.2.5.
* `se_dpdk_pmd` - Determines if dpdk pool mode driver should be used or not   0  automatically determine based on hypervisor/nic type 1  unconditionally use dpdk poll mode driver 2  don't use dpdk poll mode driver.requires se reboot. Allowed values are 0-2. Field introduced in 18.1.3.
* `se_flow_probe_retries` - Flow probe retry count if no replies are received.requires se reboot. Allowed values are 0-5. Field introduced in 18.1.4, 18.2.1.
* `se_flow_probe_retry_timer` - Timeout in milliseconds for flow probe retries.requires se reboot. Allowed values are 20-50. Field introduced in 18.2.5. Unit is milliseconds.
* `se_group_analytics_policy` - Analytics policy for serviceenginegroup. Field introduced in 20.1.3.
* `se_hyperthreaded_mode` - Controls the distribution of se data path processes on cpus which support hyper-threading. Requires hyper-threading to be enabled at host level. Requires se reboot. For more details please refer to se placement kb. Enum options - SE_CPU_HT_AUTO, SE_CPU_HT_SPARSE_DISPATCHER_PRIORITY, SE_CPU_HT_SPARSE_PROXY_PRIORITY, SE_CPU_HT_PACKED_CORES. Field introduced in 20.1.1.
* `se_ip_encap_ipc` - Determines if se-se ipc messages are encapsulated in an ip header       0        automatically determine based on hypervisor type    1        use ip encap unconditionally    ~[0,1]   don't use ip encaprequires se reboot. Field introduced in 20.1.3.
* `se_kni_burst_factor` - This knob controls the resource availability and burst size used between se datapath and kni. This helps in minimising packet drops when there is higher kni traffic (non-vip traffic from and to linux). The factor takes the following values      0-default. 1-doubles the burst size and kni resources. 2-quadruples the burst size and kni resources. Allowed values are 0-2. Field introduced in 18.2.6.
* `se_l3_encap_ipc` - Determines if se-se ipc messages use se interface ip instead of vip        0        automatically determine based on hypervisor type    1        use se interface ip unconditionally    ~[0,1]   don't use se interface iprequires se reboot. Field introduced in 20.1.3.
* `se_log_buffer_app_blocking_dequeue` - Internal flag that blocks dataplane until all application logs are flushed to log-agent process. Field introduced in 21.1.1.
* `se_log_buffer_conn_blocking_dequeue` - Internal flag that blocks dataplane until all connection logs are flushed to log-agent process. Field introduced in 21.1.1.
* `se_log_buffer_events_blocking_dequeue` - Internal flag that blocks dataplane until all outstanding events are flushed to log-agent process. Field introduced in 21.1.1.
* `se_lro` - Enable or disable large receive optimization for vnics. Requires se reboot. Field introduced in 18.2.5.
* `se_mp_ring_retry_count` - The retry count for the multi-producer enqueue before yielding the cpu. To be used under surveillance of avi support. Field introduced in 20.1.3. Allowed in basic(allowed values- 500) edition, essentials(allowed values- 500) edition, enterprise edition.
* `se_mtu` - Mtu for the vnics of ses in the se group. Allowed values are 512-9000. Field introduced in 18.2.8, 20.1.1.
* `se_name_prefix` - Prefix to use for virtual machine name of service engines.
* `se_pcap_lookahead` - Enables lookahead mode of packet receive in pcap mode. Introduced to overcome an issue with hv_netvsc driver. Lookahead mode attempts to ensure that application and kernel's view of the receive rings are consistent. Field introduced in 18.2.3.
* `se_pcap_pkt_count` - Max number of packets the pcap interface can hold and if the value is 0 the optimum value will be chosen. The optimum value will be chosen based on se-memory, cloud type and number of interfaces.requires se reboot. Field introduced in 18.2.5.
* `se_pcap_pkt_sz` - Max size of each packet in the pcap interface. Requires se reboot. Field introduced in 18.2.5.
* `se_pcap_qdisc_bypass` - Bypass the kernel's traffic control layer, to deliver packets directly to the driver. Enabling this feature results in egress packets not being captured in host tcpdump. Note   brief packet reordering or loss may occur upon toggle. Field introduced in 18.2.6.
* `se_pcap_reinit_frequency` - Frequency in seconds at which periodically a pcap reinit check is triggered. May be used in conjunction with the configuration pcap_reinit_threshold. (valid range   15 mins - 12 hours, 0 - disables). Allowed values are 900-43200. Special values are 0- 'disable'. Field introduced in 17.2.13, 18.1.3, 18.2.1. Unit is sec.
* `se_pcap_reinit_threshold` - Threshold for input packet receive errors in pcap mode exceeding which a pcap reinit is triggered. If not set, an unconditional reinit is performed. This value is checked every pcap_reinit_frequency interval. Field introduced in 17.2.13, 18.1.3, 18.2.1. Unit is metric_count.
* `se_probe_port` - Tcp port on se where echo service will be run. Field introduced in 17.2.2.
* `se_rl_prop` - Rate limiter properties. Field introduced in 20.1.1.
* `se_rum_sampling_nav_interval` - Minimum time to wait on server between taking sampleswhen sampling the navigation timing data from the end user client. Field introduced in 18.2.6. Unit is sec.
* `se_rum_sampling_nav_percent` - Percentage of navigation timing data from the end user client, used for sampling to get client insights. Field introduced in 18.2.6.
* `se_rum_sampling_res_interval` - Minimum time to wait on server between taking sampleswhen sampling the resource timing data from the end user client. Field introduced in 18.2.6. Unit is sec.
* `se_rum_sampling_res_percent` - Percentage of resource timing data from the end user client used for sampling to get client insight. Field introduced in 18.2.6.
* `se_sb_dedicated_core` - Sideband traffic will be handled by a dedicated core.requires se reboot. Field introduced in 16.5.2, 17.1.9, 17.2.3.
* `se_sb_threads` - Number of sideband threads per se.requires se reboot. Allowed values are 1-128. Field introduced in 16.5.2, 17.1.9, 17.2.3.
* `se_thread_multiplier` - Multiplier for se threads based on vcpu. Allowed values are 1-10. Allowed in basic(allowed values- 1) edition, essentials(allowed values- 1) edition, enterprise edition.
* `se_tracert_port_range` - Traceroute port range. Field introduced in 17.2.8.
* `se_tunnel_mode` - Determines if direct secondary return (dsr) from secondary se is active or not  0  automatically determine based on hypervisor type. 1  enable tunnel mode - dsr is unconditionally disabled. 2  disable tunnel mode - dsr is unconditionally enabled. Tunnel mode can be enabled or disabled at run-time. Allowed values are 0-2. Field introduced in 17.1.1. Allowed in basic(allowed values- 0) edition, essentials(allowed values- 0) edition, enterprise edition.
* `se_tunnel_udp_port` - Udp port for tunneled packets from secondary to primary se in docker bridge mode.requires se reboot. Field introduced in 17.1.3.
* `se_tx_batch_size` - Number of packets to batch for transmit to the nic. Requires se reboot. Field introduced in 18.2.5.
* `se_txq_threshold` - Once the tx queue of the dispatcher reaches this threshold, hardware queues are not polled for further packets. To be used under surveillance of avi support. Allowed values are 512-32768. Field introduced in 20.1.2. Allowed in basic(allowed values- 2048) edition, essentials(allowed values- 2048) edition, enterprise edition.
* `se_udp_encap_ipc` - Determines if se-se ipc messages are encapsulated in a udp header  0  automatically determine based on hypervisor type. 1  use udp encap unconditionally.requires se reboot. Allowed values are 0-1. Field introduced in 17.1.2.
* `se_use_dpdk` - Determines if dpdk library should be used or not   0  automatically determine based on hypervisor type 1  use dpdk if pcap is not enabled 2  don't use dpdk. Allowed values are 0-2. Field introduced in 18.1.3.
* `se_vnic_tx_sw_queue_flush_frequency` - Configure the frequency in milliseconds of software transmit spillover queue flush when enabled. This is necessary to flush any packets in the spillover queue in the absence of a packet transmit in the normal course of operation. Allowed values are 50-500. Special values are 0- 'disable'. Field introduced in 20.1.1. Unit is milliseconds.
* `se_vnic_tx_sw_queue_size` - Configure the size of software transmit spillover queue when enabled. Requires se reboot. Allowed values are 128-2048. Field introduced in 20.1.1.
* `se_vs_hb_max_pkts_in_batch` - Maximum number of aggregated vs heartbeat packets to send in a batch. Allowed values are 1-256. Field introduced in 17.1.1.
* `se_vs_hb_max_vs_in_pkt` - Maximum number of virtualservices for which heartbeat messages are aggregated in one packet. Allowed values are 1-1024. Field introduced in 17.1.1.
* `self_se_election` - Enable ses to elect a primary amongst themselves in the absence of a connectivity to controller. Field introduced in 18.1.2. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `send_se_ready_timeout` - Timeout for sending se_ready without ns helper registration completion. Allowed values are 10-600. Field introduced in 21.1.1. Unit is seconds. Allowed in basic edition, essentials edition, enterprise edition.
* `service_ip6_subnets` - Ipv6 subnets assigned to the se group. Required for vs group placement. Field introduced in 18.1.1. Maximum of 128 items allowed.
* `service_ip_subnets` - Subnets assigned to the se group. Required for vs group placement. Field introduced in 17.1.1. Maximum of 128 items allowed.
* `shm_minimum_config_memory` - Minimum required shared memory to apply any configuration. Allowed values are 0-100. Field introduced in 18.1.2. Unit is mb.
* `significant_log_throttle` - This setting limits the number of significant logs generated per second per core on this se. Default is 100 logs per second. Set it to zero (0) to deactivate throttling. Field introduced in 17.1.3. Unit is per_second.
* `ssl_preprocess_sni_hostname` - (beta) preprocess ssl client hello for sni hostname extension.if set to true, this will apply sni child's ssl protocol(s), if they are different from sni parent's allowed ssl protocol(s). Field introduced in 17.2.12, 18.1.3.
* `ssl_sess_cache_per_vs` - Number of ssl sessions that can be cached per vs. Field introduced in 21.1.1.
* `tenant_ref` - It is a reference to an object of type tenant.
* `transient_shared_memory_max` - The threshold for the transient shared config memory in the se. Allowed values are 0-100. Field introduced in 20.1.1. Unit is percent.
* `udf_log_throttle` - This setting limits the number of udf logs generated per second per core on this se. Udf logs are generated due to the configured client log filters or the rules with logging enabled. Default is 100 logs per second. Set it to zero (0) to deactivate throttling. Field introduced in 17.1.3. Unit is per_second.
* `upstream_connect_timeout` - Timeout for backend connection. Field introduced in 21.1.1. Unit is milliseconds.
* `upstream_connpool_enable` - Enable upstream connection pool,. Field introduced in 21.1.1.
* `upstream_read_timeout` - Timeout for data to be received from backend. Field introduced in 21.1.1. Unit is milliseconds.
* `upstream_send_timeout` - Timeout for upstream to become writable. Field introduced in 21.1.1. Unit is milliseconds. Allowed in basic(allowed values- 3600000) edition, essentials(allowed values- 3600000) edition, enterprise edition.
* `use_hyperthreaded_cores` - Enables the use of hyper-threaded cores on se. Requires se reboot. Field introduced in 20.1.1.
* `use_legacy_netlink` - Enable legacy model of netlink notifications. Field introduced in 21.1.1.
* `use_objsync` - Enable interse objsyc distribution framework. Field introduced in 20.1.3. Allowed in basic edition, essentials edition, enterprise edition.
* `use_standard_alb` - Use standard sku azure load balancer. By default cloud level flag is set. If not set, it inherits/uses the use_standard_alb flag from the cloud. Field introduced in 18.2.3.
* `user_agent_cache_config` - Configuration for user-agent cache used in bot management. Field introduced in 21.1.1.
* `user_defined_metric_age` - Defines in seconds how long before an unused user-defined-metric is garbage collected. Field introduced in 21.1.1. Unit is sec.
* `uuid` - Unique object identifier of the object.
* `vcenter_clusters` - Dict settings for serviceenginegroup.
* `vcenter_datastore_mode` - Enum options - vcenter_datastore_any, vcenter_datastore_local, vcenter_datastore_shared.
* `vcenter_datastores` - List of list.
* `vcenter_datastores_include` - Boolean flag to set vcenter_datastores_include.
* `vcenter_folder` - Folder to place all the service engine virtual machines in vcenter.
* `vcenter_hosts` - Dict settings for serviceenginegroup.
* `vcenters` - Vcenter information for scoping at host/cluster level. Field introduced in 20.1.1.
* `vcpus_per_se` - Number of vcpus for each of the service engine virtual machines. Changes to this setting do not affect existing ses.
* `vip_asg` - When vip_asg is set, vip configuration will be managed by avi.user will be able to configure vip_asg or vips individually at the time of create. Field introduced in 17.2.12, 18.1.2.
* `vnic_dhcp_ip_check_interval` - Dhcp ip check interval. Allowed values are 1-1000. Field introduced in 21.1.1. Unit is sec. Allowed in basic edition, essentials edition, enterprise edition.
* `vnic_dhcp_ip_max_retries` - Dhcp ip max retries. Field introduced in 21.1.1. Allowed in basic edition, essentials edition, enterprise edition.
* `vnic_ip_delete_interval` - Wait interval before deleting ip. Field introduced in 21.1.1. Unit is sec. Allowed in basic edition, essentials edition, enterprise edition.
* `vnic_probe_interval` - Probe vnic interval. Field introduced in 21.1.1. Unit is sec. Allowed in basic edition, essentials edition, enterprise edition.
* `vnic_rpc_retry_interval` - Time interval for retrying the failed vnic rpc requests. Field introduced in 21.1.1. Unit is sec. Allowed in basic edition, essentials edition, enterprise edition.
* `vnicdb_cmd_history_size` - Size of vnicdb command history. Allowed values are 0-65535. Field introduced in 21.1.1. Allowed in basic edition, essentials edition, enterprise edition.
* `vs_host_redundancy` - Ensure primary and secondary service engines are deployed on different physical hosts. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition. Special default for basic edition is true, essentials edition is true, enterprise is true.
* `vs_scalein_timeout` - Time to wait for the scaled in se to drain existing flows before marking the scalein done. Unit is sec.
* `vs_scalein_timeout_for_upgrade` - During se upgrade, time to wait for the scaled-in se to drain existing flows before marking the scalein done. Unit is sec.
* `vs_scaleout_timeout` - Time to wait for the scaled out se to become ready before marking the scaleout done. Unit is sec.
* `vs_se_scaleout_additional_wait_time` - Wait time for sending scaleout ready notification after virtual service is marked up. In certain deployments, there may be an additional delay to accept traffic. For example, for bgp, some time is needed for route advertisement. Allowed values are 0-20. Field introduced in 18.1.5,18.2.1. Unit is sec.
* `vs_se_scaleout_ready_timeout` - Timeout in seconds for service engine to sendscaleout ready notification of a virtual service. Allowed values are 0-90. Field introduced in 18.1.5,18.2.1. Unit is sec.
* `vs_switchover_timeout` - During se upgrade in a legacy active/standby segroup, time to wait for the new primary se to accept flows before marking the switchover done. Field introduced in 17.2.13,18.1.4,18.2.1. Unit is sec.
* `vss_placement` - Parameters to place virtual services on only a subset of the cores of an se. Field introduced in 17.2.5.
* `vss_placement_enabled` - If set, virtual services will be placed on only a subset of the cores of an se. Field introduced in 18.1.1.
* `waf_mempool` - Enable memory pool for waf.requires se reboot. Field introduced in 17.2.3.
* `waf_mempool_size` - Memory pool size used for waf.requires se reboot. Field introduced in 17.2.3. Unit is kb.

