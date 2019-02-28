package equipment

import (
	"github.com/foxyblue/gogue/gogue/display"
	"github.com/gdamore/tcell"
)

type Screen struct {
	Box *display.Box
}

func NewEquipmentScreen(s tcell.Screen) *Screen {
	maxW, maxH := s.Size()
	w, y := maxW-2, int(float64(maxH)*(3./4.))
	x, h := 0, maxH-1

	b := display.NewBox(x, y, w, h, s)
	return &Screen{
		Box: b,
	}
}

func (s *Screen) Draw() {
	// screen := s.Box.Screen
	s.Box.Draw()
}
