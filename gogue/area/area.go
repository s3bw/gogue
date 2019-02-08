package area

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

// Area defines the area in which the player plays
type Area struct {
	X, Y, Width, Height int

	backgroundColor tcell.Color

	border bool

	borderColor tcell.Color

	Screen tcell.Screen
}

// NewArea create a new playable area
func NewArea(level int, s tcell.Screen) *Area {
	w, h := s.Size()

	seedtime := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seedtime)
	x := random.Int() % w
	y := random.Int() % h
	lw := (random.Int() % (w - x)) + 2
	lh := (random.Int() % (h - y)) + 2

	// Assign this glyph as the background texture
	// gl := '.'
	st := tcell.Color(random.Int() % s.Colors())

	return &Area{
		X:               x,
		Y:               y,
		backgroundColor: st,
		Width:           lw,
		Height:          lh,
		Screen:          s,
	}
}

// Draw the contents of the area
func (a *Area) Draw() {
	s := a.Screen
	st := tcell.StyleDefault
	for row := 0; row < a.Height; row++ {
		for col := 0; col < a.Width; col++ {
			s.SetCell(a.X+col, a.Y+row, st.Background(a.backgroundColor), '.')
		}
	}
}
