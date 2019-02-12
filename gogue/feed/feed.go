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

func NewFeed(s tcell.Screen) *Feed {
	maxW, maxH := s.Size()
	w, y := maxW-2, int(float64(maxH)*(3./4.))
	x, h := 0, maxH-1
	lines := (h - y) - 1

	b := display.NewBox(x, y, w, h, s)
	q := CreateQueue()
	return &Feed{
		Box:   b,
		Queue: q,
		lines: lines,
	}
}

// Log adds an item to the queue, the feed displays a maximum
// number of lines, thus we dequeue the older messages.
func (f *Feed) Log(text string) {
	if f.Queue.Size() >= f.lines {
		f.Queue.Dequeue()
	}
	f.Queue.Enqueue([]rune(text))
}

func (f *Feed) Draw() {
	s := f.Box.Screen
	f.Box.Draw()

	y := f.Box.Y + 1
	for text := range f.Queue.Iterate() {
		x := f.Box.X + 1
		for _, l := range text {
			// This has inherited the Box.Style, but we should
			// actually have each message getting it's own style
			s.SetCell(x, y, f.Box.Style, l)
			x++
		}
		y++
	}
}
