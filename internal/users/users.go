package Users

import (
	Twitch "github.com/gempir/go-twitch-irc/v4"
)

type User struct {
	Username    string
	Displayname string
	AvatarUrl   string
}

func NewUser(username string) *User {
	return &User{
		Username: username,
	}
}

func GetTwitchUser(twtUser Twitch.User) *User {
	return &User{
		Username: twtUser.Name,
	}
}
