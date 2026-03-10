package main

import (
	"net"
	"sync"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Peer struct {
	Conn net.Conn
	Name string
}

type TCPPacket struct {
	Type    string `json:"type"`  
	Name    string `json:"name"`
	Message string `json:"message"`
}

type UDPPacket struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

var (
	Connections = make(map[string]*Peer)
	mu          sync.Mutex

	PORT     int
	USERNAME string

	App    fyne.App
	Window fyne.Window
)

func main() {
	App = app.NewWithID("com.chat.p2p")
	Window = App.NewWindow("P2P Local Chat")
	
	ShowLoginUI()
	Window.Resize(fyne.NewSize(600, 400))
	Window.ShowAndRun()
}