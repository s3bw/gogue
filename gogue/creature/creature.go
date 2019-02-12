package creature

import (
	"github.com/foxyblue/gogue/gogue/area/biome"
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

// Move is the creature's next action in the game
func (c *Creature) Move(dx, dy int, grid biome.Grid) {
	x := c.X + dx
	y := c.Y + dy
	if grid.Tiles[y][x].Passable {
		c.move(x, y)
	}
}

func (c *Creature) move(x, y int) {
	c.X = x
	c.Y = y
}

func (c *Creature) Draw(x, y int, screen tcell.Screen) {
	screen.SetCell(c.X+x, c.Y+y, c.background, c.Appearance)
}
