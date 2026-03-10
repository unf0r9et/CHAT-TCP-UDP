package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

func TCPConnect(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	handleTCPConnection(conn)
}

func handleTCPConnection(conn net.Conn) {
	addr := conn.RemoteAddr().String()

	handshake := TCPPacket{Type: "handshake", Name: USERNAME}
	data, _ := json.Marshal(handshake)
	conn.Write(append(data, '\n'))

	peer := &Peer{Conn: conn, Name: "Unknown"}

	mu.Lock()
	Connections[addr] = peer
	mu.Unlock()

	defer func() {
		conn.Close()
		mu.Lock()
		delete(Connections, addr)
		mu.Unlock()
		AppendLog(fmt.Sprintf("Узел отключился: %s (%s)", peer.Name, addr))
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		var packet TCPPacket
		
		if err := json.Unmarshal([]byte(line), &packet); err != nil {
			continue
		}

		switch packet.Type {
		case "handshake":
			peer.Name = packet.Name
			AppendLog(fmt.Sprintf("Обнаружен новый узел: %s (%s)", peer.Name, addr))
			
			
			go sendBroadcastMessage()

		case "message":
			AppendLog(fmt.Sprintf("%s (%s): %s", packet.Name, addr, packet.Message))
		}
	}
}

func BroadcastMessage(text string) {
	packet := TCPPacket{
		Type:    "message",
		Name:    USERNAME,
		Message: text,
	}
	data, _ := json.Marshal(packet)
	data = append(data, '\n')

	mu.Lock()
	defer mu.Unlock()

	for addr, peer := range Connections {
		_, err := peer.Conn.Write(data)
		if err != nil {
		 
			fmt.Printf("Не удалось отправить сообщение на %s\n", addr)
		}
	}
}