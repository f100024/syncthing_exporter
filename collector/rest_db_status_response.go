package collector

// DBStatusResponse is a struct of the `syncthing` instance DB status
type DBStatusResponse struct {
	Errors float64 `json:"errors"`

	GlobalBytes       float64 `json:"globalBytes"`
	GlobalDeleted     float64 `json:"globalDeleted"`
	GlobalDirectories float64 `json:"globalDirectories"`
	GlobalFiles       float64 `json:"globalFiles"`
	GlobalSymlinks    float64 `json:"globalSymlinks"`
	GlobalTotalItems  float64 `json:"globalTotalItems"`

	IgnorePatters bool `json:"ignorePatterns"`

	InSyncBytes float64 `json:"inSyncBytes"`
	InSyncFiles float64 `json:"inSyncFiles"`

	Invalid string `json:"invalid,omitempty"` // Deprecated, retains external API for now

	LocalBytes       float64 `json:"localBytes"`
	LocalDeleted     float64 `json:"localDeleted"`
	LocalDirectories float64 `json:"localDirectories"`
	LocalFiles       float64 `json:"localFiles"`
	LocalSymlinks    float64 `json:"localSymlinks"`
	LocalTotalItems  float64 `json:"localTotalItems"`

	NeedBytes       float64 `json:"needBytes"`
	NeedDeletes     float64 `json:"needDeletes"`
	NeedDirectories float64 `json:"needDirectories"`
	NeedFiles       float64 `json:"needFiles"`
	NeedSymlinks    float64 `json:"needSymlinks"`
	NeedTotalItems  float64 `json:"needTotalItems"`

	PullErrors float64 `json:"pullErrors"`
	Sequence   float64 `json:"sequence"`

	State        string  `json:"state"`
	StateChanged string  `json:"stateChanged"`
	Version      float64 `json:"version,omitempty"` // Deprecated, retains external API for now

}
