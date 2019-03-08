package entity

import (
	"fmt"

	"github.com/foxyblue/gogue/gogue/feed"
	"github.com/foxyblue/gogue/gogue/styles"
	"github.com/gdamore/tcell"
)

// Creature defines the creature
type Creature struct {
	*Base

	// HP are the hit points this creature has
	HP int

	Backpack []Item

	Strength int
}

func (c *Creature) CanCarry(weight int) bool {
	return c.Strength > weight
}

// Identify returns the type associated with the creature
func (c *Creature) Identify() Type {
	return c.Base.Type
}

func (c *Creature) Draw(x, y int, screen tcell.Screen) {
	entity := c.Base
	screen.SetCell(entity.X+x, entity.Y+y, entity.Style, entity.Appearance)
}

func (c *Creature) Move(x, y int) {
	c.Base.X = x
	c.Base.Y = y
}

func (c *Creature) Kill() {
	c.Base.ChangeAppearence('%')
	c.Base.MakeOrnamentation()
	dead := styles.DeadStyle()
	c.Base.ChangeStyle(dead)
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
