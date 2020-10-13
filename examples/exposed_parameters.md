
# Exposing metrics

[back to README.md](../README.md)

```
# HELP syncthing_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which syncthing_exporter was built.
# TYPE syncthing_exporter_build_info gauge
syncthing_exporter_build_info{branch="",goversion="go1.15",revision="",version=""} 1
# HELP syncthing_rest_svc_report_hash_performance Hash Performance value
# TYPE syncthing_rest_svc_report_hash_performance gauge
syncthing_rest_svc_report_hash_performance 184.61
# HELP syncthing_rest_svc_report_json_parse_failures Number of errors while parsing JSON.
# TYPE syncthing_rest_svc_report_json_parse_failures counter
syncthing_rest_svc_report_json_parse_failures 0
# HELP syncthing_rest_svc_report_memory_size Node memory size in megabytes
# TYPE syncthing_rest_svc_report_memory_size gauge
syncthing_rest_svc_report_memory_size 8191
# HELP syncthing_rest_svc_report_memory_usage_mb Memory usage by syncthc in MB
# TYPE syncthing_rest_svc_report_memory_usage_mb gauge
syncthing_rest_svc_report_memory_usage_mb 358
# HELP syncthing_rest_svc_report_node_info Node's string information
# TYPE syncthing_rest_svc_report_node_info gauge
syncthing_rest_svc_report_node_info{longVersion="syncthing v1.10.0 \"Fermium Flea\" (go1.15.2 windows-amd64) teamcity@build.syncthing.net 2020-09-15 17:38:23 UTC",platform="windows-amd64",uniqueID="d9uZew76",version="v1.10.0"} 1
# HELP syncthing_rest_svc_report_num_cpu Number of node CPU
# TYPE syncthing_rest_svc_report_num_cpu gauge
syncthing_rest_svc_report_num_cpu 4
# HELP syncthing_rest_svc_report_number_devices Number of devices in sync
# TYPE syncthing_rest_svc_report_number_devices gauge
syncthing_rest_svc_report_number_devices 3
# HELP syncthing_rest_svc_report_number_folders Number of folders in sync
# TYPE syncthing_rest_svc_report_number_folders gauge
syncthing_rest_svc_report_number_folders 2
# HELP syncthing_rest_svc_report_sha256_performance SHA256 Performance value
# TYPE syncthing_rest_svc_report_sha256_performance gauge
syncthing_rest_svc_report_sha256_performance 203.12
# HELP syncthing_rest_svc_report_total_data_in_MB Total data in megabytes
# TYPE syncthing_rest_svc_report_total_data_in_MB gauge
syncthing_rest_svc_report_total_data_in_MB 3.556006e+06
# HELP syncthing_rest_svc_report_total_files Total number of files
# TYPE syncthing_rest_svc_report_total_files gauge
syncthing_rest_svc_report_total_files 522111
# HELP syncthing_rest_svc_report_total_scrapes Current total Syncthings scrapes.
# TYPE syncthing_rest_svc_report_total_scrapes counter
syncthing_rest_svc_report_total_scrapes 4610
# HELP syncthing_rest_svc_report_up Was the last scrape of the Syncthing endpoint successful.
# TYPE syncthing_rest_svc_report_up gauge
syncthing_rest_svc_report_up 1
# HELP syncthing_rest_svc_report_uptime Syncthing uptime in seconds
# TYPE syncthing_rest_svc_report_uptime gauge
syncthing_rest_svc_report_uptime 582297
# HELP syncthing_rest_system_connections_json_parse_failures Number of errors while parsing JSON.
# TYPE syncthing_rest_system_connections_json_parse_failures counter
syncthing_rest_system_connections_json_parse_failures 0
# HELP syncthing_rest_system_connections_remote_device_connection_info Information of the connection with remote device. 1 - connected; 0 -disconnected.
# TYPE syncthing_rest_system_connections_remote_device_connection_info gauge
syncthing_rest_system_connections_remote_device_connection_info{address="",clientVersion="",crypto="",deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA",type=""} 0
syncthing_rest_system_connections_remote_device_connection_info{address="10.1.0.1:43364",clientVersion="v1.10.0",crypto="TLS1.3-TLS_AES_128_GCM_SHA256",deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA",type="tcp-server"} 1
syncthing_rest_system_connections_remote_device_connection_info{address="10.2.0.2:25241",clientVersion="v1.10.0",crypto="TLS1.3-TLS_AES_128_GCM_SHA256",deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA",type="tcp-client"} 1
# HELP syncthing_rest_system_connections_remote_device_in_bytes_total Total bytes received from remote device.
# TYPE syncthing_rest_system_connections_remote_device_in_bytes_total gauge
syncthing_rest_system_connections_remote_device_in_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 1.124147267e+09
syncthing_rest_system_connections_remote_device_in_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 1.137699873e+09
syncthing_rest_system_connections_remote_device_in_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 0
# HELP syncthing_rest_system_connections_remote_device_is_paused Is sync paused. 1 - paused, 0 - not paused.
# TYPE syncthing_rest_system_connections_remote_device_is_paused gauge
syncthing_rest_system_connections_remote_device_is_paused{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 0
syncthing_rest_system_connections_remote_device_is_paused{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 0
syncthing_rest_system_connections_remote_device_is_paused{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 0
# HELP syncthing_rest_system_connections_remote_device_out_bytes_total Total bytes transmitted to remote device.
# TYPE syncthing_rest_system_connections_remote_device_out_bytes_total gauge
syncthing_rest_system_connections_remote_device_out_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 7.478664224e+09
syncthing_rest_system_connections_remote_device_out_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 7.483754446e+09
syncthing_rest_system_connections_remote_device_out_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 0
# HELP syncthing_rest_system_connections_total_in_bytes_total Total bytes received to device.
# TYPE syncthing_rest_system_connections_total_in_bytes_total gauge
syncthing_rest_system_connections_total_in_bytes_total 2.42579682e+09
# HELP syncthing_rest_system_connections_total_out_bytes_total Total bytes transmitted from device.
# TYPE syncthing_rest_system_connections_total_out_bytes_total gauge
syncthing_rest_system_connections_total_out_bytes_total 1.5405146137e+10
# HELP syncthing_rest_system_connections_total_scrapes Current total Syncthings scrapes.
# TYPE syncthing_rest_system_connections_total_scrapes counter
syncthing_rest_system_connections_total_scrapes 4610
# HELP syncthing_rest_system_connections_up Was the last scrape of the Syncting system connections endpoint successful.
# TYPE syncthing_rest_system_connections_up gauge
syncthing_rest_system_connections_up 1

```