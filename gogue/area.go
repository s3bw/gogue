package gogue

import (
	"github.com/foxyblue/gogue/gogue/biome"
	"github.com/foxyblue/gogue/gogue/biome/factory"
	"github.com/foxyblue/gogue/gogue/display"
	"github.com/foxyblue/gogue/gogue/entity"
	"github.com/foxyblue/gogue/gogue/feed"
	"github.com/gdamore/tcell"
)

// Area defines the area in which the user plays
type Area struct {
	Box *display.Box

	Screen tcell.Screen

	Grid biome.Grid

	Creatures []*entity.Creature

	Player *entity.Player

	Feed *feed.Feed
}

// NewArea creates a new playable area
func NewArea(player *entity.Player, level int, s tcell.Screen, feed *feed.Feed) *Area {
	maxW, maxH := s.Size()
	x, y := 0, 0
	w, h := maxW-2, int(float64(maxH)*(3./4.))
	pX, pY := player.Creature.X, player.Creature.Y

	b := display.NewBox(x, y, w, h, s)
	newBiome := NewBiome(x, y, pX, pY, w, h)
	newBiome.Generate()

	return &Area{
		Box:       b,
		Screen:    s,
		Grid:      newBiome.GetGrid(),
		Creatures: newBiome.GetCreatures(),
		Player:    player,
		Feed:      feed,
	}
}

// NewBiome generates a new biome with given parameters
func NewBiome(x, y, px, py, w, h int) biome.Biome {
	start := &biome.Coord{X: px, Y: py}
	params := make(map[string]interface{})
	params["start"] = start
	params["x"], params["y"] = x, y
	params["maxX"], params["maxY"] = w, h

	newBiome, err := factory.Create("blank", params)
	if err != nil {
		panic(err)
	}
	return newBiome
}

// targetCreatures is effectively an iterator
func (a *Area) targetCreatures() []*entity.Creature {
	l := []*entity.Creature{}
	list := append(l, a.Creatures...)
	return append(list, a.Player.Creature)
}

// MoveCreature will move a creature in the area, checking for collisions
func (a *Area) MoveCreature(obj *entity.Creature, dx, dy int) {
	var target *entity.Creature

	x := obj.X + dx
	y := obj.Y + dy
	target = nil
	// I'm going to have to use an iterator which filters
	// for types.
	for _, monster := range a.targetCreatures() {
		if monster.X == x && monster.Y == y {
			target = monster
			break
		}
	}
	if target == nil {
		if a.Grid.Tiles[x][y].Passable {
			obj.Move(x, y)
		}
	} else {
		// This can kill, when a creature dies, it transforms
		// into an object. I won't have to remove it from the
		// list. I might have to replace it.
		obj.Attack(target, a.Feed)
	}
}

// Draw the contents of the area
func (a *Area) Draw() {
	a.Box.Draw()
	offsetX, offsetY := a.Grid.OffsetX, a.Grid.OffsetY
	for x, row := range a.Grid.Tiles {
		for y, tile := range row {
			a.Screen.SetCell(offsetX+x, offsetY+y, tile.Style, tile.Appearance)
		}
	}

	for _, monster := range a.Creatures {
		a.Screen.SetCell(offsetX+monster.X, offsetY+monster.Y, monster.Style, monster.Appearance)
	}
	a.Player.Creature.Draw(offsetX, offsetY, a.Screen)
}
