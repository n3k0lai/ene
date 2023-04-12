package Mahjong

import core "github.com/n3k0lai/ene/cmd"

type Player struct {
	User             core.User
	GameHistory      []Game
	CurrentlyPlaying []Game
}
