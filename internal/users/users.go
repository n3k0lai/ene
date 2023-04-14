package Users

import (
	Adapters "github.com/n3k0lai/ene/internal/adapters"
)

type User struct {
	Username    string
	Displayname string
	AvatarUrl   string
	AdapterType Adapters.AdapterType
}

func NewUser() *User {
	return &User{}
}
