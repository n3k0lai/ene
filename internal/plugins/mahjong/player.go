package Mahjong

import User "github.com/n3k0lai/ene/internal/users"

type Player struct {
	User             User.User
	GameHistory      []Game
	CurrentlyPlaying []Game
}
