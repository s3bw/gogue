package biome

import (
	"math/rand"
	"time"

	"github.com/foxyblue/gogue/gogue/styles"
	"github.com/gdamore/tcell"
)

// Coord represents the location of a single point
type Coord struct {
	X, Y int
}

// RandomCoord generates a random co-ordinate
func RandomCoord(minX, maxX, minY, maxY int) *Coord {
	rand.Seed(time.Now().Unix())
	x := (rand.Intn(maxX-minX) + minX)
	y := (rand.Intn(maxY-minY) + minY)
	return &Coord{
		X: x,
		Y: y,
	}
}

// IsIn checks if x and y overlap a given list of co-ordinates
func IsIn(x, y int, list []*Coord) bool {
	for _, b := range list {
		if b.X == x && b.Y == y {
			return true
		}
	}
	return false
}

// Tile a single point on a grid
type Tile struct {
	Visible    bool
	Passable   bool
	Appearance rune
	Style      tcell.Style
	coord      Coord
}

// EmptyTile represents an empty space
func EmptyTile(x, y int) *Tile {
	style := styles.DefaultStyle()
	return &Tile{
		Visible:    false,
		Passable:   true,
		Style:      style,
		Appearance: '·',
	}
}

// WallTile represents a wall
func WallTile(x, y int) *Tile {
	style := styles.DefaultStyle()
	return &Tile{
		Visible:    false,
		Passable:   false,
		Style:      style,
		Appearance: '#',
	}
}

// Grid is a 2D array consisting of tiles
type Grid struct {
	OffsetX int
	OffsetY int
	Tiles   [][]*Tile
}

// NewGrid creates a 'h by w' size grid to use as a playable map
// width is generated first and then height to form the (x, y) indexing
func NewGrid(x, y, w, h int) *Grid {
	grid := make([][]*Tile, w)
	for row := range grid {
		grid[row] = make([]*Tile, h)
	}
	return &Grid{
		Tiles:   grid,
		OffsetX: x + 1,
		OffsetY: y + 1,
	}
}
