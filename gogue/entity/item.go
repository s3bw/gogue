package entity

import "github.com/gdamore/tcell"

type Item struct {
	*Base

	Weight int
}

func (i *Item) Identify() Type {
	return i.Base.Type
}

func (i *Item) Draw(x, y int, screen tcell.Screen) {
	entity := i.Base
	screen.SetCell(entity.X+x, entity.Y+y, entity.Style, entity.Appearance)
}
