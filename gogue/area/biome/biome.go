package biome

import "github.com/gdamore/tcell"

// Biome represents the level with which the player interacts
type Biome interface {
	// Draw should draw all the contents of the domain, e.g. the
	// creatures, items, walls, stairs and doors.
	Draw(s tcell.Screen)

	// Generate should create a new grid and fill the grid with
	// contents with which the player interacts
	Generate()

	// StartLocation defines the point at which the player enters
	// this is usually random.
	StartLocation() *Coord

	// EndLocation is the point on the grid at which the player
	// must exit to a lower domain
	EndLocation() *Coord
}
