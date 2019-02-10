package area

import (
	"github.com/foxyblue/gogue/gogue/display"
	"github.com/gdamore/tcell"
)

// Area defines the area in which the user plays
type Area struct {
	Box *display.Box

	Screen tcell.Screen

	Start *PlayerStart
}

type PlayerStart struct {
	X int
	Y int
}

// NewArea creates a new playable area
func NewArea(level int, s tcell.Screen) *Area {
	maxW, maxH := s.Size()
	x, y := 1, 1
	w, h := maxW-2, int(float64(maxH)*(3./4.))

	b := display.NewBox(x, y, w, h, s)
	return &Area{
		Box:    b,
		Screen: s,
		Start:  &PlayerStart{X: 2, Y: 2},
	}
}

// Draw the contents of the area
func (a *Area) Draw() {
	a.Box.Draw()

	// This is where we should draw the contents of the game
}
