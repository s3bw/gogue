package feed

import (
	"github.com/foxyblue/gogue/gogue/display"
	"github.com/gdamore/tcell"
)

const MaxSize = 5

type Feed struct {
	Box *display.Box

	lines int

	Queue Queue
}

// NewArea creates a new playable area
func NewFeed(s tcell.Screen) *Feed {
	maxW, maxH := s.Size()
	w, y := maxW-2, int(float64(maxH)*(3./4.))+1
	x, h := 1, maxH-1
	lines := (h - y) - 1

	b := display.NewBox(x, y, w, h, s)
	q := CreateQueue()
	return &Feed{
		Box:   b,
		Queue: q,
		lines: lines,
	}
}

func (f *Feed) Log(text string) {
	if f.Queue.Size() >= f.lines {
		f.Queue.Dequeue()
	}
	f.Queue.Enqueue([]rune(text))
}

func (f *Feed) Draw() {
	s := f.Box.Screen
	st := tcell.StyleDefault
	f.Box.Draw()

	y := f.Box.Y + 1
	for text := range f.Queue.Iterate() {
		x := f.Box.X + 1
		for _, l := range text {
			s.SetCell(x, y, st.Background(f.Box.BackgroundColor), l)
			x++
		}
		y++
	}
}
