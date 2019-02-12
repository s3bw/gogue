package creature

import (
	"fmt"

	"github.com/foxyblue/gogue/gogue/feed"
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

func (c *Creature) Move(x, y int) {
	c.X = x
	c.Y = y
}

func (c *Creature) Kill() {
	c.Appearance = '%'
}

func (c *Creature) Attack(target *Creature, feed *feed.Feed) {
	target.HP--
	if target.HP <= 0 {
		target.Kill()
		feed.Log(fmt.Sprintf("The %s died!", target.Name))
	} else {
		feed.Log(fmt.Sprintf("%s hit, %d HP remaining!", target.Name, target.HP))
	}
}

func (c *Creature) Draw(x, y int, screen tcell.Screen) {
	screen.SetCell(c.X+x, c.Y+y, c.Style, c.Appearance)
}
