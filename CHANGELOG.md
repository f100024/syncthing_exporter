## 0.3.17 / 2025-12-04
---
* Bump go version to 1.25.6
* Update dependecies
* Fixed GO-2026-4340,GO-2026-4341,GO-2026-4342
* Bump alpine to 3.23
## 0.3.16 / 2025-12-04
---
* Bump go version to 1.25.5
* Update dependecies
* Fixed CVE-2025-47914,CVE-2025-58181
## 0.3.15 / 2025-11-03
---
* Bump go version to 1.25.3
* Update dependecies
## 0.3.14 / 2025-08-11
---
* Fix CVE-2025-47907
* Bump go version to 1.24.6
* Update docker image
## 0.3.13 / 2025-06-11
---
* Added ability to choose custom endpoints for gathering metrics `--endpoints`
* Fix missed parameter `global_files`, Thanks @Emily3403 (#34)
* Update some dependencies
* Bump go version to 1.24.4
* Update docker image
## 0.3.12 / 2025-04-14
---
* Fix "Docker instances crashing, seemingly at random" (#30)
* Minor fixes
## 0.3.11 / 2025-04-11
---
* Added support /rest/config/devices (#29)
* Updated dependencies
* Minor fixes
## 0.3.10 / 2025-02-24
---
* Switched from promlog to promslog
* Switched to go 1.24
* Updated dependencies
* Minor fixes
## 0.3.9 / 2024-12-23
---
* Updated dependencies
* Minor fixes
## 0.3.8 / 2024-08-26
---
* Switched to go1.23
* Updated dependencies
## 0.3.7 / 2023-03-20
---
* Switched to go1.22
* Updated dependencies

## 0.3.6 / 2023-10-27
---
* Fixed bug (#25) node reporting down
* Updated Grafana Dashboard example
* Added linux/riscv64 build
* Switched to go1.21
* Updated dependencies
## 0.3.5 / 2023-06-08
---
* Switched to go1.20
* Updated dependencies
## 0.3.4 / 2022-12-02
---
* Fixed crash due to missing value `lastConnectionDurationS` (thanks to @benediktschlager) 
* Fixed arm* docker image (thanks to @Luxtech)
* Switched to go1.19
* Updated dependencies
## 0.3.3 / 2022-06-12
---
* Switch to go1.18
* Updated dependencies
* Updated promu to 0.13.0

## 0.3.2 / 2021-10-19
---
* Switch to go1.17
* Updated dependencies for compatibility with go1.17
* Updated promu to 0.12.0
* Added multiple platforms 

## 0.3.1 / 2021-04-15
---
* Created new metric `last_connection_timestamp`
* Removed the `lastSeen` label from `last_connection_duration`

## 0.3.0 / 2021-04-15
---
* Added support /rest/db/status
* Fixed logging
* Updated grafana dashboard
* Refactoring

## 0.2.4 / 2021-04-13
---
* Added support /rest/stats/device
* Minor refactoring

## 0.2.3 / 2021-03-31
---
* Added information about using exporter with docker

## 0.2.2 / 2021-03-31
---
* Swithched to go 1.16
* Update some dependencies

## 0.2.1 / 2020-12-30
---
* Added CI Github actions
* Environment variables renamed and clarified
* Added required flag to start parameters

## 0.1.0 / 2020-10-14
---
* Initial release
