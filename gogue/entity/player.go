package entity

import "github.com/foxyblue/gogue/gogue/styles"

type Player struct {
	Creature *Creature
}

func playerCreature(x, y int) *Creature {
	style := styles.DefaultStyle()
	ent := entity{
		Name:       "Player",
		X:          x,
		Y:          y,
		Style:      style,
		Appearance: '@',
		Type:       Beast,
	}
	return &Creature{
		entity: &ent,
		HP:     10,
	}
}

// NewPlayer creates a player instance
func NewPlayer(x, y int) *Player {
	creature := playerCreature(x, y)
	return &Player{
		Creature: creature,
	}
}
