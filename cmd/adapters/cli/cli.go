package CliAdapter

import (
	"fmt"
)

func Send(m Message) {
	fmt.Println(m.Text)
}

func Respond(c Conversation) {

}

func OnMessage(m Message) {

}
