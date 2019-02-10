package creature

type Player struct {
	Creature *Creature
}

func NewPlayer(x, y int) *Player {
	creature := NewCreature(x, y)
	return &Player{
		Creature: creature,
	}
}
