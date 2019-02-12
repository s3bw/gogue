package creature

import (
	"github.com/gdamore/tcell"
)

// Creature defines the creature
type Creature struct {
	X, Y int

	background tcell.Style

	Appearance rune
}

type ListCreatures struct {
	list []*Creature
}

func NewCreature(x, y int) *Creature {
	style := tcell.StyleDefault
	return &Creature{
		X:          x,
		Y:          y,
		background: style.Background(tcell.Color(tcell.ColorBlack)),
		Appearance: '@',
	}
}

func (c *Creature) Move(x, y int) {
	c.X = x
	c.Y = y
}

func (c *Creature) Draw(x, y int, screen tcell.Screen) {
	screen.SetCell(c.X+x, c.Y+y, c.background, c.Appearance)
}
