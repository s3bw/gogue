package display

import (
	"github.com/gdamore/tcell"
)

// Box defines a window on the terminal screen
type Box struct {
	X, Y, Width, Height int

	BackgroundColor tcell.Color

	Corners []*Corner

	BorderColor tcell.Color

	Screen tcell.Screen
}

// Corner represents a box's corner
type Corner struct {
	x, y int
}

// NewBox creates a new display box
func NewBox(x, y, w, h int, s tcell.Screen) *Box {
	// 90 is arbitrary
	c := tcell.ColorWhite
	corners := []*Corner{
		&Corner{x: x, y: y},
		&Corner{x: x, y: h},
		&Corner{x: w, y: y},
		&Corner{x: w, y: h},
	}
	return &Box{
		X:       x,
		Y:       y,
		Width:   w,
		Height:  h,
		Corners: corners,
		// 170 is arbitrary
		BorderColor:     tcell.Color(170 % s.Colors()),
		BackgroundColor: c,
		Screen:          s,
	}
}

// Draw creates the border for the box
func (b *Box) Draw() {
	s := b.Screen
	st := tcell.StyleDefault
	// Draw corners
	for _, corner := range b.Corners {
		s.SetCell(corner.x, corner.y, st.Background(b.BackgroundColor), '+')
	}
	// Draw Top and Bottom
	for x := b.X + 1; x < b.Width; x++ {
		s.SetCell(x, b.Y, st.Background(b.BackgroundColor), '-')
		s.SetCell(x, b.Height, st.Background(b.BackgroundColor), '-')
	}
	// Draw sides
	for y := b.Y + 1; y < b.Height; y++ {
		s.SetCell(b.X, y, st.Background(b.BackgroundColor), '|')
		s.SetCell(b.Width, y, st.Background(b.BackgroundColor), '|')
	}

}
