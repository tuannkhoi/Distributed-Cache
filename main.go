package main

import (
	"log"
	"net"
	"time"

	"github.com/tuannkhoi/Distributed-Cache/cache"
)

func main() {
	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}

	go func() {
		time.Sleep(time.Second * 2)
		conn, err := net.Dial("tcp", ":3000")
		if err != nil {
			log.Fatal(err)
		}

		conn.Write([]byte("SET KHOI FAANG 25000"))
	}()

	server := NewServer(opts, cache.New())

	err := server.Start()
	if err != nil {
		return
	}
}
