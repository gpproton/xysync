package model


// DatabaseConfig.go

type DatabaseConfig struct {
	Type       *string `json:"type,omitempty"`
	HostOrPath *string `json:"hostOrPath,omitempty"`
	Port       *int64  `json:"port,omitempty"`
	Charset    *string `json:"charset,omitempty"`
	Dbname     *string `json:"dbname,omitempty"`
	Username   *string `json:"username,omitempty"`
	Password   *string `json:"password,omitempty"`
}

// database.go

type Database struct {
	Query    *Query     `json:"query,omitempty"`
	SubQuery []SubQuery `json:"subQuery"`
}

// query.go

type Query struct {
	Check   *string `json:"check,omitempty"`
	Perform *string `json:"perform,omitempty"`
	Confirm *string `json:"confirm,omitempty"`
}

// subquery.go

type SubQuery struct {
	Name  *string `json:"name,omitempty"`
	Query *Query  `json:"query,omitempty"`
}