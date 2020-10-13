package collector

type connectionParameters struct {
	IPAddress     string `json:"address"`
	AtDate        string `json:"at"`
	ClientVersion string `json:"clientVersion"`
	Connected     bool   `json:"connected"`
	CryptoConfig  string `json:"crypto"`
	InBytesTotal  int    `json:"inBytesTotal"`
	OutBytesTotal int    `json:"outBytesTotal"`
	Paused        bool   `json:"paused"`
	Type          string `json:"type"`
}

// SystemConnectionsResponse defines response struct
// for /rest/system/connections endpoint.
// Tested with Syncthing version 1.9.0.
type SystemConnectionsResponse struct {
	CS    map[string]connectionParameters `json:"connections"`
	Total connectionParameters            `json:"total"`
}
