package creature

import (
	"github.com/gdamore/tcell"
)

// Creature defines the creature
type Creature struct {
	X, Y int

	Color tcell.Color

	Appearance rune
}

func NewCreature(x, y int) *Creature {
	return &Creature{
		X:          x,
		Y:          y,
		Color:      tcell.Color(tcell.ColorBlack),
		Appearance: '@',
	}
}

func (c *Creature) Move(dx, dy int) {
	c.X += dx
	c.Y += dy
}
