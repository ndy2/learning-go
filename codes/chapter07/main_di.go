package main

import "fmt"

type Sender interface {
	Send(msg string)
}

type EmailSender struct{}

func (e EmailSender) Send(msg string) {
	fmt.Println("EmailSender: " + msg)
}

type SmsSender struct{}

func (s SmsSender) Send(msg string) {
	fmt.Println("SmsSender: " + msg)
}

type SenderApp struct {
	sender Sender
}

func (sa SenderApp) Run() {
	fmt.Println("SenderApp is running")
	sa.sender.Send("Hello")
}

func main() {
	// s := EmailSender{}
	s := SmsSender{}

	SenderApp{sender: s}.Run()
}
