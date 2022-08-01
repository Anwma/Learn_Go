package main

import (
	"bytes"
	"sync"
)

func main() {

}

type Info struct {
	mu  sync.Mutex
	Str string
	// ... other fields, e.g.: Str string
}

func Update(info *Info) {
	info.mu.Lock()
	// critical section:
	info.Str = "3"
	// new value
	// end critical section
	info.mu.Unlock()
}

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}
