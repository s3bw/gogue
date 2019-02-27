package biome

import "github.com/foxyblue/gogue/gogue/entity"

// Biome represents the level with which the player interacts
type Biome interface {
	// Generate should create a new grid and fill the grid with
	// contents with which the player interacts
	Generate()

	// GetGrid returns the tiled map
	GetGrid() Grid

	GetCreatures() []*entity.Creature

	// StartLocation defines the point at which the player enters
	// this is usually random.
	StartLocation() *Coord

	// EndLocation is the point on the grid at which the player
	// must exit to a lower domain
	EndLocation() *Coord
}
