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
		HP:   5,
	}
}

func NewSword(x, y int) *entity.Item {
	style := styles.DefaultStyle()
	base := entity.Base{
		Name:       "Sword",
		X:          x,
		Y:          y,
		Style:      style,
		Appearance: '/',
		Type:       entity.TypeItem,
	}
	return &entity.Item{
		Base:   &base,
		Weight: 3,
	}
}
