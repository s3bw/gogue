package biome

// Coord represents the location of a single point
type Coord struct {
	X, Y int
}

// Tile a single point on a grid
type Tile struct {
	visible    bool
	passable   bool
	appearance rune
	coord      Coord
}

// Grid is a 2D array consisting of tiles
type Grid struct {
	tiles [][]*Tile
}

// NewGrid creates a h by w size grid to use as a playable map
func NewGrid(w, h int) *Grid {
	a := make([][]*Tile, h)
	for i := range a {
		a[i] = make([]*Tile, w)
	}
	return &Grid{tiles: a}
}
