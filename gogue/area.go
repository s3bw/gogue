package gogue

import (
	"github.com/foxyblue/gogue/gogue/biome"
	"github.com/foxyblue/gogue/gogue/biome/factory"
	"github.com/foxyblue/gogue/gogue/creature"
	"github.com/foxyblue/gogue/gogue/display"
	"github.com/gdamore/tcell"
)

// Area defines the area in which the user plays
type Area struct {
	Box *display.Box

	Screen tcell.Screen

	Grid biome.Grid

	Player *creature.Player
}

// NewArea creates a new playable area
func NewArea(player *creature.Player, level int, s tcell.Screen) *Area {
	maxW, maxH := s.Size()
	x, y := 1, 1
	w, h := maxW-2, int(float64(maxH)*(3./4.))
	pX, pY := player.Creature.X, player.Creature.Y

	b := display.NewBox(x, y, w, h, s)
	newBiome := NewBiome(x, y, pX, pY, w, h)
	newBiome.Generate()

	return &Area{
		Box:    b,
		Screen: s,
		Grid:   newBiome.GetGrid(),
		// Creatures: newBiome.GetCreatures(),
		Player: player,
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

func (a *Area) MoveCreature(obj *creature.Creature, dx, dy int) {
	// var target creature.Creature

	x := obj.X + dx
	y := obj.Y + dy
	// for creature := a.Creatures.Iterate() {
	// 	if creature.X == x && creature.Y == y {
	// 		target = creature
	// 	}
	// }
	//if !target {
	if a.Grid.Tiles[y][x].Passable {
		obj.Move(x, y)
	}
	// } else {
	// 	obj.Attack(target)
	// }
}

// Draw the contents of the area
func (a *Area) Draw() {
	c := tcell.Color(3 % a.Screen.Colors())
	st := tcell.StyleDefault

	a.Box.Draw()
	// a.Biome.Draw(a.Screen)
	// x, y := b.parameters.x, b.parameters.y
	// x, y := a.Biome.OffSet()
	x, y := 2, 2
	for ri, row := range a.Grid.Tiles {
		for ci, tile := range row {
			a.Screen.SetCell(x+ci, y+ri, st.Background(c), tile.Appearence)
		}
	}
	a.Player.Creature.Draw(x, y, a.Screen)

	// This is where we should draw the contents of the game
}
