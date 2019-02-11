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

func NewCreature(x, y int) *Creature {
	style := tcell.StyleDefault
	return &Creature{
		X:          x,
		Y:          y,
		background: style.Background(tcell.Color(tcell.ColorBlack)),
		Appearance: '@',
	}
}

func (c *Creature) Move(dx, dy int) {
	c.X += dx
	c.Y += dy
}

func (c *Creature) Draw(screen tcell.Screen) {
	screen.SetCell(c.X, c.Y, c.background, c.Appearance)
}
