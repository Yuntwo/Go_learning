package main

import "github.com/dustin/go-broadcast"

var roomChannels = make(map[string]broadcast.Broadcaster)

func openListener(roomId string) chan interface{} {
	listener := make(chan interface{})
	room(roomId).Register(listener)
	return listener
}

func closeListener(roomId string, listener chan interface{}) {
	room(roomId).Unregister(listener)
	close(listener)
}

func room(roomId string) broadcast.Broadcaster {
	b, ok := roomChannels[roomId]
	if !ok {
		b = broadcast.NewBroadcaster(10)
		roomChannels[roomId] = b
	}
	return b
}
