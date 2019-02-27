package entity

import "github.com/foxyblue/gogue/gogue/styles"

// type Player struct {
// 	Creature *Creature
// }

func playerCreature(x, y int) *Creature {
	style := styles.DefaultStyle()
	base := Base{
		Name:       "Player",
		X:          x,
		Y:          y,
		Style:      style,
		Appearance: '@',
		Type:       TypeCreature,
	}
	return &Creature{
		Base:     &base,
		HP:       10,
		Strength: 15,
	}
}

// func NewPlayer(x, y int) *Player {
// 	creature := playerCreature(x, y)
// 	return &Player{
// 		Creature: creature,
// 	}
// }

// NewPlayer creates a player instance
func NewPlayer(x, y int) *Creature {
	return playerCreature(x, y)
}
