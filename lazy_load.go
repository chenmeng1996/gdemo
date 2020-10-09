package main

import (
	"sync"
)

var names map[string]string
var loadOnce sync.Once

func initLoad() {
	names = map[string]string{
		"1": "1",
	}
}

func Get(key string) string {
	loadOnce.Do(initLoad)
	return names[key]
}
