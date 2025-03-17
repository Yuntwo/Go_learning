package main

import (
	"runtime"
	"sync"
	"time"

	"github.com/manucorporat/stats"
)

var (
	ips        = stats.New()
	messages   = stats.New()
	users      = stats.New()
	mutexStats sync.RWMutex
	savedStats map[string]uint64
)

// statsWorker is a goroutine that updates stats data in var savedStats every second.
func statsWorker() {
	c := time.Tick(1 * time.Second)
	var lastMallocs uint64
	var lastFrees uint64
	for range c {
		var stats runtime.MemStats
		runtime.ReadMemStats(&stats)

		// This is a RW Lock
		mutexStats.Lock()
		// This map is replaced everytime
		savedStats = map[string]uint64{
			"timestamp":  uint64(time.Now().Unix()),
			"HeapInuse":  stats.HeapInuse,
			"StackInuse": stats.StackInuse,
			"Mallocs":    stats.Mallocs - lastMallocs,
			"Frees":      stats.Frees - lastFrees,
			"Inbound":    uint64(messages.Get("inbound")),
			"Outbound":   uint64(messages.Get("outbound")),
			"Connected":  connectedUsers(),
		}
		lastMallocs = stats.Mallocs
		lastFrees = stats.Frees
		messages.Reset()
		mutexStats.Unlock()
	}
}

func connectedUsers() uint64 {
	connected := users.Get("connected") - users.Get("disconnected")
	if connected < 0 {
		return 0
	}
	return uint64(connected)
}

// Stats returns savedStats data.
func Stats() map[string]uint64 {
	mutexStats.RLock()
	defer mutexStats.RUnlock()

	return savedStats
}
