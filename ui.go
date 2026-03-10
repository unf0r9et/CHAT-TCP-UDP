package main

import (
	"fmt"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var chatHistory binding.StringList

func ShowLoginUI() {
	label := widget.NewLabel("Введите ваше имя:")
	nameEntry := widget.NewEntry()

	buttonAccept := widget.NewButton("Войти в чат", func() {
		if nameEntry.Text != "" {
			USERNAME = nameEntry.Text
			StartNetworking()
			ShowChatUI()
		}
	})

	content := container.NewVBox(
		label,
		nameEntry,
		buttonAccept,
	)

	Window.SetContent(container.NewCenter(content))
}

func ShowChatUI() {
	chatHistory = binding.NewStringList()
	list := widget.NewListWithData(
		chatHistory,
		func() fyne.CanvasObject {
			label := widget.NewLabel("template")
			label.Wrapping = fyne.TextWrapWord
			return label
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		},
	)

	messageEntry := widget.NewEntry()
	messageEntry.PlaceHolder = "Введите сообщение..."

	sendButton := widget.NewButton("Отправить", func() {
		text := messageEntry.Text
		if text != "" {
			BroadcastMessage(text)
			AppendLog(fmt.Sprintf("Вы: %s", text))
			messageEntry.SetText("")
		}
	})

	bottomPanel := container.NewBorder(nil, nil, nil, sendButton, messageEntry)
	content := container.NewBorder(nil, bottomPanel, nil, nil, list)

	Window.SetContent(content)
	AppendLog(fmt.Sprintf("Добро пожаловать в чат, %s!", USERNAME))
}

func AppendLog(msg string) {
	timestamp := time.Now().Format("15:04:05")
	formatted := fmt.Sprintf("[%s] %s", timestamp, msg)
	_ = chatHistory.Append(formatted)
}
