package area

import (
	"github.com/foxyblue/gogue/gogue/area/domain"
	"github.com/foxyblue/gogue/gogue/area/domain/factory"
	"github.com/foxyblue/gogue/gogue/display"
	"github.com/gdamore/tcell"
)

// Area defines the area in which the user plays
type Area struct {
	Box *display.Box

	Screen tcell.Screen

	Domain domain.Domain
}

// NewArea creates a new playable area
func NewArea(pX, pY, level int, s tcell.Screen) *Area {
	maxW, maxH := s.Size()
	x, y := 1, 1
	w, h := maxW-2, int(float64(maxH)*(3./4.))

	b := display.NewBox(x, y, w, h, s)

	start := &domain.Coord{X: pX, Y: pY}
	params := make(map[string]interface{})
	params["start"] = start

	domain, _ := factory.Create("blank", params)
	return &Area{
		Box:    b,
		Screen: s,
		Domain: domain,
	}
}

func (a *Area) playerXY() (int, int) {
	coord := a.Domain.StartLocation()
	return coord.X, coord.Y
}

// Draw the contents of the area
func (a *Area) Draw() {
	a.Box.Draw()

	// This is where we should draw the contents of the game
}
