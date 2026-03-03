package main

import (
	"fmt"
	"net"
)

func TCPlistener() {
	listener, err := net.Listen("tcp", "10.220.76.218:8080")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("ERROR: ", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		fmt.Printf("Получено: %s", string(buffer[:n]))
		conn.Write([]byte("Сообщение получено"))
	}
}
