package core

import (
	"time"
)

type Adapter struct {
	Channel     string
	MsgRate     time.Duration
	Name        string
	Port        string
	PrivatePath string // oauth
	Server      string
}
