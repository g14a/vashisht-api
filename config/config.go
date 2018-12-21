package config

import "sync"

var (
	DATABASE   = "vashisht"
	COLLECTION = "users"
	mu         sync.Mutex
	size       int
	err        error
)
