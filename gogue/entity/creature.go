package entity

import (
	"fmt"

	"github.com/foxyblue/gogue/gogue/feed"
	"github.com/gdamore/tcell"
)

// Creature defines the creature
type Creature struct {
	*entity

	// HP are the hit points this creature has
	HP int
}

func (c *Creature) Move(x, y int) {
	c.entity.X = x
	c.entity.Y = y
}

func (c *Creature) Kill() {
	c.entity.Appearance = '%'
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
	entity := c.entity
	screen.SetCell(entity.X+x, entity.Y+y, entity.Style, entity.Appearance)
}
