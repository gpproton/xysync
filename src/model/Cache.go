package model

type Cache struct {
	Type           *string `json:"type,omitempty"`
	Host           *string `json:"host,omitempty"`
	Port           *int64  `json:"port,omitempty"`
	DefaultTimeout *string `json:"defaultTimeout,omitempty"`
}
