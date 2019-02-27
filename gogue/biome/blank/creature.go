package blank

import (
	"github.com/foxyblue/gogue/gogue/entity"
	"github.com/foxyblue/gogue/gogue/styles"
)

func NewRabbit(x, y int) *entity.Creature {
	style := styles.DefaultStyle()
	base := entity.Base{
		Name:       "Rabbit",
		X:          x,
		Y:          y,
		Style:      style,
		Appearance: 'r',
		Type:       entity.TypeCreature,
	}
	return &entity.Creature{
		Base: &base,
		HP:   2,
	}
}
