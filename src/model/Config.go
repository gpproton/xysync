package model

type Welcome struct {
	ThreadsCount *int64     `json:"threadsCount,omitempty"`
	Cache        *Cache     `json:"cache,omitempty"`
	Instances    []Instance `json:"instances"`
}