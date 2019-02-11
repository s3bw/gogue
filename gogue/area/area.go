package area

import (
	"github.com/foxyblue/gogue/gogue/area/biome"
	"github.com/foxyblue/gogue/gogue/area/biome/factory"
	"github.com/foxyblue/gogue/gogue/display"
	"github.com/gdamore/tcell"
)

// Area defines the area in which the user plays
type Area struct {
	Box *display.Box

	Screen tcell.Screen

	Biome biome.Biome
}

// NewArea creates a new playable area
func NewArea(pX, pY, level int, s tcell.Screen) *Area {
	maxW, maxH := s.Size()
	x, y := 1, 1
	w, h := maxW-2, int(float64(maxH)*(3./4.))

	b := display.NewBox(x, y, w, h, s)

	// Specify parameters for the new domain.
	start := &biome.Coord{X: pX, Y: pY}
	params := make(map[string]interface{})
	params["start"] = start
	params["maxX"], params["maxY"] = w, h

	newBiome, err := factory.Create("blank", params)
	if err != nil {
		panic(err)
	}
	newBiome.Generate()
	return &Area{
		Box:    b,
		Screen: s,
		Biome:  newBiome,
	}
}

// Move the domain stuff here...
// func NewBiome() {}

func (a *Area) playerXY() (int, int) {
	coord := a.Biome.StartLocation()
	return coord.X, coord.Y
}

// Draw the contents of the area
func (a *Area) Draw() {
	a.Box.Draw()
	a.Biome.Draw()

	// This is where we should draw the contents of the game
}
