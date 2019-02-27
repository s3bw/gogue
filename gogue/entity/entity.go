package entity

import "github.com/gdamore/tcell"

// Type represents the type of entity that's on the map
type Type string

const (
	// TypeCreature entities move around the map
	TypeCreature Type = "Creature"
	// TypeItem entities can be picked up
	TypeItem Type = "Item"
)

// Entity represents the interface an object on the map should have
type Entity interface {
	Identify() Type
	Draw(x, y int, screen tcell.Screen)
}

// Base are the initial attributes one needs for an entity on the map
type Base struct {
	Name       string
	X          int
	Y          int
	Style      tcell.Style
	Appearance rune
	Type       Type
}

func (b *Base) ChangeStyle(style tcell.Style) {
	b.Style = style
}

func (b *Base) ChangeAppearence(r rune) {
	b.Appearance = r
}

func (b *Base) MakeItem() {
	b.Type = TypeItem
}
