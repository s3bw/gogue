package feed

import (
	"github.com/foxyblue/gogue/gogue/display"
	"github.com/gdamore/tcell"
)

const MaxSize = 5

type Feed struct {
	Box *display.Box

	Feed Queue
}

// NewArea creates a new playable area
func NewFeed(s tcell.Screen) *Feed {
	maxW, maxH := s.Size()
	w, y := maxW-2, int(float64(maxH)*(3./4.))+1
	x, h := 1, maxH-1

	b := display.NewBox(x, y, w, h, s)
	q := CreateQueue()
	return &Feed{
		Box:  b,
		Feed: q,
	}
}

func (f *Feed) Draw() {
	s := f.Box.Screen
	st := tcell.StyleDefault
	f.Box.Draw()
	// s.SetCell(f.Box.X+2, f.Box.Y+2, st.Background(f.Box.BackgroundColor), '#')
	for text := range f.Feed.Iterate() {
		for x := 0; x < len(text.(string)); x++ {
			s.SetCell(f.Box.X+x+1, f.Box.Y+1, st.Background(f.Box.BackgroundColor), rune(text.(string)[x]))
		}
	}
}

// func (f *Feed) Println(s string) {
// 	if len(f.feed) > MaxSize {
// 		f.feed = f.feed[1:]
// 	}
// 	f.feed = append(f.feed, s)
// 	f.Draw()
// }

// func (f *Feed) Draw() {
// 	f.Screen
// }
