package creature

import "github.com/foxyblue/gogue/gogue/styles"

func NewRabbit(x, y int) *Creature {
	style := styles.DefaultStyle()
	return &Creature{
		Name:       "Rabbit",
		HP:         2,
		X:          x,
		Y:          y,
		Style:      style,
		Appearance: 'r',
	}
}
