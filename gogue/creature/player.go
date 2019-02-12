package creature

import "github.com/foxyblue/gogue/gogue/styles"

type Player struct {
	Creature *Creature
}

func playerCreature(x, y int) *Creature {
	style := styles.DefaultStyle()
	return &Creature{
		Name:       "Player",
		HP:         10,
		X:          x,
		Y:          y,
		Style:      style,
		Appearance: '@',
	}
}

// NewPlayer creates a player instance
func NewPlayer(x, y int) *Player {
	creature := playerCreature(x, y)
	return &Player{
		Creature: creature,
	}
}
