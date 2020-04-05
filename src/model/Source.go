package model

// source.go

type Source struct {
	HTTP     *HTTP           `json:"http,omitempty"`
	Database *Database `json:"database,omitempty"`
}

// sync.go

type Sync struct {
	Destination *Source `json:"destination,omitempty"`
	Source      *Source  `json:"source,omitempty"`
}

// SourceClass.go

type SourceClass struct {
	Mode           *string         `json:"mode,omitempty"`
	Frequency      *string         `json:"frequency,omitempty"`
	HTTPConfig     *HTTPConfig     `json:"httpConfig,omitempty"`
	DatabaseConfig *DatabaseConfig `json:"databaseConfig,omitempty"`
}

