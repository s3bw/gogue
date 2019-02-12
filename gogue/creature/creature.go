package creature

import (
	"github.com/gdamore/tcell"
)

// Creature defines the creature
type Creature struct {

	// Name is the name of the creature
	Name string

	// HP are the hit points this creature has
	HP int

	// X, Y are the co ordinates of the creature
	X, Y int

	Style tcell.Style

	Appearance rune
}

type ListCreatures struct {
	list []*Creature
}

// func NewRabbit(x, y int) *Creature {
//
// }

func (c *Creature) Move(x, y int) {
	c.X = x
	c.Y = y
}

func (c *Creature) Draw(x, y int, screen tcell.Screen) {
	screen.SetCell(c.X+x, c.Y+y, c.Style, c.Appearance)
}
