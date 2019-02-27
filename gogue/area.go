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

	// Entities is a list of the map entities
	Entities []entity.Entity

	Player *entity.Creature

	Feed *feed.Feed
}

// NewArea creates a new playable area
func NewArea(player *entity.Creature, level int, s tcell.Screen, feed *feed.Feed) *Area {
	maxW, maxH := s.Size()
	x, y := 0, 0
	w, h := maxW-2, int(float64(maxH)*(3./4.))
	pX, pY := player.X, player.Y

	b := display.NewBox(x, y, w, h, s)
	newBiome := NewBiome(x, y, pX, pY, w, h)
	newBiome.Generate()

	return &Area{
		Box:      b,
		Screen:   s,
		Grid:     newBiome.GetGrid(),
		Entities: newBiome.GetEntities(),
		Player:   player,
		Feed:     feed,
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

// Draw the contents of the area
func (a *Area) Draw() {
	a.Box.Draw()
	offsetX, offsetY := a.Grid.OffsetX, a.Grid.OffsetY
	for x, row := range a.Grid.Tiles {
		for y, tile := range row {
			a.Screen.SetCell(offsetX+x, offsetY+y, tile.Style, tile.Appearance)
		}
	}

	for _, e := range a.Entities {
		e.Draw(offsetX, offsetY, a.Screen)
	}
	a.Player.Draw(offsetX, offsetY, a.Screen)
}
