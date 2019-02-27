package entity

import "github.com/foxyblue/gogue/gogue/styles"

func NewRabbit(x, y int) *Creature {
	style := styles.DefaultStyle()
	ent := entity{
		Name:       "Rabbit",
		X:          x,
		Y:          y,
		Style:      style,
		Appearance: 'r',
		Type:       Beast,
	}
	return &Creature{
		entity: &ent,
		HP:     2,
	}
}
