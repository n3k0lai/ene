package Mahjong

type IGame interface {
  Start()
	RenderTable()
	GetHand(p Player)
  Close()
}

type Game struct {
	Id      string
	Players []Player
}

func NewGame(players []Player) *Game {
  return &Game{
    Players: players,
  }
}

func Close(g *Game) {
  
}