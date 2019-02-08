package main

import (
	"fmt"
	"os"
	"time"

	"github.com/foxyblue/gogue/gogue/area"
	"github.com/foxyblue/gogue/gogue/creature"
	"github.com/gdamore/tcell"
)

// Game holds the instance of the game
type Game struct {
	Screen tcell.Screen

	// Level refers to the level at which the active area exists
	Level int

	// CurrentArea refers to the active area to which the player is in.
	ActiveArea *area.Area

	// Player refers to the user
	Player *creature.Player
}

// NewGame creates a new game instance
func NewGame() *Game {
	return &Game{
		Screen: newScreen(),
		// Player: newPlayer(),
		Level: 0,
	}
}

// newScreen generates a screen instance for the game to be played
func newScreen() tcell.Screen {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err = s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorWhite))
	s.Clear()
	return s
}

func (g *Game) CreateArea() {
	area := area.NewArea(g.Level, g.Screen)
	g.ActiveArea = area
}

func (g *Game) CreatePlayer() {
	player := creature.NewPlayer(g.ActiveArea)
	g.Player = player
}

func (g *Game) Draw() {
	g.Screen.Clear()
	g.ActiveArea.Draw()
	st := tcell.StyleDefault
	p := g.Player.Creature
	g.Screen.SetCell(p.X, p.Y, st.Background(p.Color), p.Appearance)
	g.Screen.Show()
}

func main() {
	game := NewGame()
	game.CreateArea()
	game.CreatePlayer()

	// This is the Key Listener Channel
	quit := make(chan struct{})
	go func() {
		for {
			ev := game.Screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEnter:
					close(quit)
					return
				case tcell.KeyCtrlL:
					game.Screen.Sync()
				case tcell.KeyRune:
					switch ev.Rune() {
					case 'k':
						game.Player.Creature.Move(0, -1)
					case 'j':
						game.Player.Creature.Move(0, 1)
					case 'h':
						game.Player.Creature.Move(-1, 0)
					case 'l':
						game.Player.Creature.Move(1, 0)
					}
					game.Draw()
				}
			case *tcell.EventResize:
				game.Screen.Sync()
			}
		}
	}()

	// Main Gameloop
	cnt := 0
	dur := time.Duration(0)
gameLoop:
	for {
		select {
		case <-quit:
			break gameLoop
		case <-time.After(time.Millisecond * 50):
		}
		start := time.Now()
		game.Draw()
		cnt++
		dur += time.Now().Sub(start)
	}

	game.Screen.Fini()
	fmt.Println("Game has ended.")
}
