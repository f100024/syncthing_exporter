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
