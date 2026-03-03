package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.NewWithID("com.unf0r9et.chat")

	myWindow := myApp.NewWindow("CHAT-TCP-UDP")

	go TCPlistener()

	myWindow.Resize(fyne.NewSize(200, 200))

	myWindow.ShowAndRun()
}
