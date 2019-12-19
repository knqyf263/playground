package main

import (
	"fmt"

	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	NewMessage,
	NewGreeter,
	NewEvent,
)

type Message string

func NewMessage() Message {
	return "Hello World!"
}

type Greeter struct {
	Message Message
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	event := InitializeEvent()

	event.Start()
}
