package Tarot

//type Arcana {
//	"major" | "minor"
//}

type TarotCard struct {
	Name string
}

func NewTarotCard(name string) *TarotCard {
	return &TarotCard{
		Name: name,
	}
}

func (c *TarotCard) Render(adapterType string) {

}
