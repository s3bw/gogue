package creature

import (
	"math/rand"
	"time"

	"github.com/foxyblue/gogue/gogue/area"
	"github.com/gdamore/tcell"
)

// Creature defines the creature
type Creature struct {
	X, Y int

	Color tcell.Color

	Appearance rune
}

type Player struct {
	Creature *Creature
}

func NewCreature(x, y int) *Creature {
	return &Creature{
		X:          x,
		Y:          y,
		Color:      tcell.Color(tcell.ColorBlack),
		Appearance: '@',
	}
}

func NewPlayer(a *area.Area) *Player {
	seedtime := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seedtime)

	x := (random.Int() % a.Width) + a.X
	y := (random.Int() % a.Height) + a.Y

	creature := NewCreature(x, y)
	return &Player{
		Creature: creature,
	}
}

func (c *Creature) Move(dx, dy int) {
	c.X += dx
	c.Y += dy
}
