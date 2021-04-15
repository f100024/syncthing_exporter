
# Exposing metrics

[back to README.md](../README.md)

```
# HELP syncthing_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which syncthing_exporter was built.
# TYPE syncthing_exporter_build_info gauge
syncthing_exporter_build_info{branch="",goversion="go1.16",revision="",version=""} 1
# HELP syncthing_rest_db_status_errors Number of errors for current folder.
# TYPE syncthing_rest_db_status_errors gauge
syncthing_rest_db_status_errors{folderID="aaaaa-bb11b"} 0
# HELP syncthing_rest_db_status_global_bytes Number of bytes globally.
# TYPE syncthing_rest_db_status_global_bytes gauge
syncthing_rest_db_status_global_bytes{folderID="aaaaa-bb11b"} 4.047955455069e+12
# HELP syncthing_rest_db_status_global_deleted Number of bytes deleted.
# TYPE syncthing_rest_db_status_global_deleted gauge
syncthing_rest_db_status_global_deleted{folderID="aaaaa-bb11b"} 156112
# HELP syncthing_rest_db_status_global_directories Number of directories globally.
# TYPE syncthing_rest_db_status_global_directories gauge
syncthing_rest_db_status_global_directories{folderID="aaaaa-bb11b"} 346372
# HELP syncthing_rest_db_status_global_symlinks Number of symlinks globally.
# TYPE syncthing_rest_db_status_global_symlinks gauge
syncthing_rest_db_status_global_symlinks{folderID="aaaaa-bb11b"} 0
# HELP syncthing_rest_db_status_global_total_items Number of total items globally.
# TYPE syncthing_rest_db_status_global_total_items gauge
syncthing_rest_db_status_global_total_items{folderID="aaaaa-bb11b"} 1.197108e+06
# HELP syncthing_rest_db_status_ignore_patterns Is using ignore patterns.
# TYPE syncthing_rest_db_status_ignore_patterns gauge
syncthing_rest_db_status_ignore_patterns{folderID="aaaaa-bb11b"} 0
# HELP syncthing_rest_db_status_insync_bytes Number of bytes currently in sync.
# TYPE syncthing_rest_db_status_insync_bytes gauge
syncthing_rest_db_status_insync_bytes{folderID="aaaaa-bb11b"} 4.047955455069e+12
# HELP syncthing_rest_db_status_insync_files Number of files currently in sync.
# TYPE syncthing_rest_db_status_insync_files gauge
syncthing_rest_db_status_insync_files{folderID="aaaaa-bb11b"} 694624
# HELP syncthing_rest_db_status_json_parse_failures Number of errors while parsing JSON.
# TYPE syncthing_rest_db_status_json_parse_failures counter
syncthing_rest_db_status_json_parse_failures 0
# HELP syncthing_rest_db_status_local_bytes Number of bytes locally.
# TYPE syncthing_rest_db_status_local_bytes gauge
syncthing_rest_db_status_local_bytes{folderID="aaaaa-bb11b"} 4.047955455069e+12
# HELP syncthing_rest_db_status_local_deleted Number of bytes deleted locally.
# TYPE syncthing_rest_db_status_local_deleted gauge
syncthing_rest_db_status_local_deleted{folderID="aaaaa-bb11b"} 318
# HELP syncthing_rest_db_status_local_directories Number of local directories.
# TYPE syncthing_rest_db_status_local_directories gauge
syncthing_rest_db_status_local_directories{folderID="aaaaa-bb11b"} 346372
# HELP syncthing_rest_db_status_local_symlinks Number of local symlinks
# TYPE syncthing_rest_db_status_local_symlinks gauge
syncthing_rest_db_status_local_symlinks{folderID="aaaaa-bb11b"} 0
# HELP syncthing_rest_db_status_local_total_items Number of total items locally
# TYPE syncthing_rest_db_status_local_total_items gauge
syncthing_rest_db_status_local_total_items{folderID="aaaaa-bb11b"} 1.041314e+06
# HELP syncthing_rest_db_status_need_bytes Number of bytes need for sync.
# TYPE syncthing_rest_db_status_need_bytes gauge
syncthing_rest_db_status_need_bytes{folderID="aaaaa-bb11b"} 0
# HELP syncthing_rest_db_status_need_deletes Number of bytes need for deletes.
# TYPE syncthing_rest_db_status_need_deletes gauge
syncthing_rest_db_status_need_deletes{folderID="aaaaa-bb11b"} 0
# HELP syncthing_rest_db_status_need_directories Number of directories for sync.
# TYPE syncthing_rest_db_status_need_directories gauge
syncthing_rest_db_status_need_directories{folderID="aaaaa-bb11b"} 0
# HELP syncthing_rest_db_status_need_symlinks Number of symlinks need for sync.
# TYPE syncthing_rest_db_status_need_symlinks gauge
syncthing_rest_db_status_need_symlinks{folderID="aaaaa-bb11b"} 0
# HELP syncthing_rest_db_status_need_total_items Number of total items need to sync.
# TYPE syncthing_rest_db_status_need_total_items gauge
syncthing_rest_db_status_need_total_items{folderID="aaaaa-bb11b"} 0
# HELP syncthing_rest_db_status_pull_errors Number of pull errors.
# TYPE syncthing_rest_db_status_pull_errors gauge
syncthing_rest_db_status_pull_errors{folderID="aaaaa-bb11b"} 0
# HELP syncthing_rest_db_status_sequence Total bytes received from remote device.
# TYPE syncthing_rest_db_status_sequence gauge
syncthing_rest_db_status_sequence{folderID="aaaaa-bb11b",state="idle",stateChanged="2077-02-06T00:02:00-07:00"} 1.213411e+06
# HELP syncthing_rest_db_status_total_scrapes Current total Syncthings scrapes.
# TYPE syncthing_rest_db_status_total_scrapes counter
syncthing_rest_db_status_total_scrapes 1
# HELP syncthing_rest_stats_device_json_parse_failures Number of errors while parsing JSON.
# TYPE syncthing_rest_stats_device_json_parse_failures counter
syncthing_rest_stats_device_json_parse_failures 0
# HELP syncthing_rest_stats_device_last_connection_duration Duration of last connection with remote device in seconds.
# TYPE syncthing_rest_stats_device_last_connection_duration gauge
syncthing_rest_stats_device_last_connection_duration{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA",lastSeen="2077-12-31T00:00:50.3810375-08:00"} 0
syncthing_rest_stats_device_last_connection_duration{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-BBBBBBB",lastSeen="2077-04-14T22:14:49-07:00"} 42832.8784021
syncthing_rest_stats_device_last_connection_duration{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-CCCCCCC",lastSeen="1977-12-31T16:00:00-08:00"} 0
# HELP syncthing_rest_stats_device_total_scrapes Current total Syncthings scrapes.
# TYPE syncthing_rest_stats_device_total_scrapes counter
syncthing_rest_stats_device_total_scrapes 1
# HELP syncthing_rest_stats_device_up Was the last scrape of the Syncting system connections endpoint successful.
# TYPE syncthing_rest_stats_device_up gauge
syncthing_rest_stats_device_up 1
# HELP syncthing_rest_svc_report_hash_performance Hash Performance value
# TYPE syncthing_rest_svc_report_hash_performance gauge
syncthing_rest_svc_report_hash_performance 104.92
# HELP syncthing_rest_svc_report_json_parse_failures Number of errors while parsing JSON.
# TYPE syncthing_rest_svc_report_json_parse_failures counter
syncthing_rest_svc_report_json_parse_failures 0
# HELP syncthing_rest_svc_report_memory_size Node memory size in megabytes
# TYPE syncthing_rest_svc_report_memory_size gauge
syncthing_rest_svc_report_memory_size 8191
# HELP syncthing_rest_svc_report_memory_usage_mb Memory usage by syncthc in MB
# TYPE syncthing_rest_svc_report_memory_usage_mb gauge
syncthing_rest_svc_report_memory_usage_mb 264
# HELP syncthing_rest_svc_report_node_info Node's string information
# TYPE syncthing_rest_svc_report_node_info gauge
syncthing_rest_svc_report_node_info{longVersion="syncthing v1.15.1 \"Fermium Flea\" (go1.16.3 windows-amd64) teamcity@build.syncthing.net 2021-04-06 08:42:29 UTC",platform="windows-amd64",uniqueID="d9uZew76",version="v1.15.1"} 1
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
syncthing_rest_svc_report_sha256_performance 119.87
# HELP syncthing_rest_svc_report_total_data_in_MB Total data in megabytes
# TYPE syncthing_rest_svc_report_total_data_in_MB gauge
syncthing_rest_svc_report_total_data_in_MB 3.860566e+06
# HELP syncthing_rest_svc_report_total_files Total number of files
# TYPE syncthing_rest_svc_report_total_files gauge
syncthing_rest_svc_report_total_files 725584
# HELP syncthing_rest_svc_report_total_scrapes Current total Syncthings scrapes.
# TYPE syncthing_rest_svc_report_total_scrapes counter
syncthing_rest_svc_report_total_scrapes 1
# HELP syncthing_rest_svc_report_up Was the last scrape of the Syncthing endpoint successful.
# TYPE syncthing_rest_svc_report_up gauge
syncthing_rest_svc_report_up 1
# HELP syncthing_rest_svc_report_uptime Syncthing uptime in seconds
# TYPE syncthing_rest_svc_report_uptime gauge
syncthing_rest_svc_report_uptime 51271
# HELP syncthing_rest_system_connections_json_parse_failures Number of errors while parsing JSON.
# TYPE syncthing_rest_system_connections_json_parse_failures counter
syncthing_rest_system_connections_json_parse_failures 0
# HELP syncthing_rest_system_connections_remote_device_connection_info Information of the connection with remote device. 1 - connected; 0 -disconnected.
# TYPE syncthing_rest_system_connections_remote_device_connection_info gauge
syncthing_rest_system_connections_remote_device_connection_info{address="",clientVersion="",crypto="",deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA",type=""} 0
syncthing_rest_system_connections_remote_device_connection_info{address="",clientVersion="",crypto="",deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-BBBBBBB",type=""} 0
syncthing_rest_system_connections_remote_device_connection_info{address="127.0.0.1:1234",clientVersion="v1.15.1",crypto="TLS1.3-TLS_CHACHA20_POLY1305_SHA256",deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-CCCCCCC",type="tcp-client"} 1
# HELP syncthing_rest_system_connections_remote_device_in_bytes_total Total bytes received from remote device.
# TYPE syncthing_rest_system_connections_remote_device_in_bytes_total gauge
syncthing_rest_system_connections_remote_device_in_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 0
syncthing_rest_system_connections_remote_device_in_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-BBBBBBB"} 1.199599e+06
syncthing_rest_system_connections_remote_device_in_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-CCCCCCC"} 0
# HELP syncthing_rest_system_connections_remote_device_is_paused Is sync paused. 1 - paused, 0 - not paused.
# TYPE syncthing_rest_system_connections_remote_device_is_paused gauge
syncthing_rest_system_connections_remote_device_is_paused{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 1
syncthing_rest_system_connections_remote_device_is_paused{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-BBBBBBB"} 0
syncthing_rest_system_connections_remote_device_is_paused{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-CCCCCCC"} 0
# HELP syncthing_rest_system_connections_remote_device_out_bytes_total Total bytes transmitted to remote device.
# TYPE syncthing_rest_system_connections_remote_device_out_bytes_total gauge
syncthing_rest_system_connections_remote_device_out_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA"} 0
syncthing_rest_system_connections_remote_device_out_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-BBBBBBB"} 7.01416862e+08
syncthing_rest_system_connections_remote_device_out_bytes_total{deviceID="AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-AAAAAAA-CCCCCCC"} 0
# HELP syncthing_rest_system_connections_total_in_bytes_total Total bytes received to device.
# TYPE syncthing_rest_system_connections_total_in_bytes_total gauge
syncthing_rest_system_connections_total_in_bytes_total 1.199599e+06
# HELP syncthing_rest_system_connections_total_out_bytes_total Total bytes transmitted from device.
# TYPE syncthing_rest_system_connections_total_out_bytes_total gauge
syncthing_rest_system_connections_total_out_bytes_total 7.01416862e+08
# HELP syncthing_rest_system_connections_total_scrapes Current total Syncthings scrapes.
# TYPE syncthing_rest_system_connections_total_scrapes counter
syncthing_rest_system_connections_total_scrapes 1
# HELP syncthing_rest_system_connections_up Was the last scrape of the Syncting system connections endpoint successful.
# TYPE syncthing_rest_system_connections_up gauge
syncthing_rest_system_connections_up 1
```
