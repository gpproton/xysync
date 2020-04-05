package model

// instance.go

type Instance struct {
	Tag         *string      `json:"tag,omitempty"`
	Source      *SourceClass `json:"source,omitempty"`
	Destination *SourceClass `json:"destination,omitempty"`
	Syncs       []Sync       `json:"syncs"`
}


// item.go

type Item struct { }

