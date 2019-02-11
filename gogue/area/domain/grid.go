package domain

// Block a single tile on the gird
type Block struct {
	visible    bool
	passable   bool
	appearance rune
}

// Grid is a 2D array consisting of blocks
type Grid struct {
	tiles [][]*Block
}

// NewGrid creates a h by w size grid to use as a playable map
func NewGrid(w, h int) *Grid {
	a := make([][]*Block, h)
	for i := range a {
		a[i] = make([]*Block, w)
	}
	return &Grid{tiles: a}
}
