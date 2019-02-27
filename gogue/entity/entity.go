package entity

import "github.com/gdamore/tcell"

// Type represents the type of entity that's on the map
type Type string

const (
	// Beast entities move around the map
	Beast Type = "Beast"
	// Item entities can be picked up
	Item Type = "Item"
)

type entity struct {
	Name       string
	X          int
	Y          int
	Style      tcell.Style
	Appearance rune
	Type       Type
}

// Entity represents the interface an object on the map should have
type Entity interface {
	Identify() Type
	Draw(x, y int, screen tcell.Screen)
}
