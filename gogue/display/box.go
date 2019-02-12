package display

import (
	"github.com/foxyblue/gogue/gogue/styles"
	"github.com/gdamore/tcell"
)

// Box defines a window on the terminal screen
type Box struct {
	X, Y, Width, Height int

	Style tcell.Style

	Corners []*Corner

	Screen tcell.Screen
}

// Corner represents a box's corner
type Corner struct {
	x, y int
}

// NewBox creates a new display box
func NewBox(x, y, w, h int, s tcell.Screen) *Box {
	style := styles.DefaultStyle()
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
		Style:   style,
		Screen:  s,
	}
}

// Draw creates the border for the box
func (b *Box) Draw() {
	s := b.Screen
	// Draw corners
	for _, corner := range b.Corners {
		s.SetCell(corner.x, corner.y, b.Style, '+')
	}
	// Draw Top and Bottom
	for x := b.X + 1; x < b.Width; x++ {
		s.SetCell(x, b.Y, b.Style, '-')
		s.SetCell(x, b.Height, b.Style, '-')
	}
	// Draw sides
	for y := b.Y + 1; y < b.Height; y++ {
		s.SetCell(b.X, y, b.Style, '|')
		s.SetCell(b.Width, y, b.Style, '|')
	}

}
