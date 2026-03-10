package main

import (
	_ "bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
)

func StartNetworking() {
	go TCPListener()
	go func() {
		for PORT == 0 {
		}
		go sendBroadcastMessage()
		go UDPListener()
	}()
}

func UDPListener() {
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:8989")
	if err != nil {
		AppendLog("Ошибка настройки UDP: " + err.Error())
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		AppendLog("Ошибка прослушивания UDP: " + err.Error())
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, sender, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}

		var msg UDPPacket
		err = json.Unmarshal(bytes.TrimRight(buf[:n], "\x00\n"), &msg)
		if err != nil || msg.Name == USERNAME || msg.Port == PORT {
			continue  
		}

		targetAddr := fmt.Sprintf("%s:%d", sender.IP.String(), msg.Port)
		
		mu.Lock()
		_, exists := Connections[targetAddr]
		mu.Unlock()

		if !exists {
			go TCPConnect(targetAddr)
		}
	}
}

func sendBroadcastMessage() {
	conn, err := net.Dial("udp", "255.255.255.255:8989")
	if err != nil {
		return
	}
	defer conn.Close()

	packet := UDPPacket{Name: USERNAME, Port: PORT}
	data, _ := json.Marshal(packet)
	conn.Write(append(data, '\n'))
}

func TCPListener() {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return
	}
	PORT = listener.Addr().(*net.TCPAddr).Port

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleTCPConnection(conn)
	}
}