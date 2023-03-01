package main

import "github.com/tuannkhoi/Distributed-Cache/cache"

func main() {
	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}
	server := NewServer(opts, cache.New())
	err := server.Start()
	if err != nil {
		return
	}
}
