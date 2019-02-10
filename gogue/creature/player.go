package creature

import (
	"math/rand"
	"time"

	"github.com/foxyblue/gogue/gogue/area"
)

type Player struct {
	Creature *Creature
}

func NewPlayer(a *area.Area) *Player {
	seedtime := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seedtime)
	x := (random.Int() % a.Width) + a.X
	y := (random.Int() % a.Height) + a.Y

	creature := NewCreature(x, y)
	return &Player{
		Creature: creature,
	}
}
