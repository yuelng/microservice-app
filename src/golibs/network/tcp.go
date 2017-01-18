package main

import (
	"log"
	"net"
)

func Start() {
	addr, err := net.ResolveTCPAddr("tcp", "8080")
	ln, err := net.ListenTCP("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			log.Println("accept:", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn *net.TCPConn) {

}
