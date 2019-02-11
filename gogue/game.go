package main

import (
	"fmt"
	"os"
	"time"

	"github.com/foxyblue/gogue/gogue/area"
	"github.com/foxyblue/gogue/gogue/creature"
	"github.com/foxyblue/gogue/gogue/feed"
	"github.com/gdamore/tcell"
)

// Game holds the instance of the game
type Game struct {
	Screen tcell.Screen

	// Level refers to the level at which the active area exists
	Level int

	// ActiveArea refers to the active area to which the player is in.
	ActiveArea *area.Area

	// Player refers to the user
	Player *creature.Player

	// stdFeed is the in game feed
	Feed *feed.Feed
}

// NewGame creates a new game instance
func NewGame() *Game {
	level := 0

	screen := newScreen()
	area := area.NewArea(level, screen)
	x := area.Start.X
	y := area.Start.Y
	feed := feed.NewFeed(screen)
	player := creature.NewPlayer(x, y)
	return &Game{
		Screen:     screen,
		ActiveArea: area,
		Player:     player,
		Level:      0,
		Feed:       feed,
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

// Draw will render the game on screen
func (game *Game) Draw() {
	game.Screen.Clear()
	game.ActiveArea.Draw()
	st := tcell.StyleDefault
	p := game.Player.Creature
	game.Screen.SetCell(p.X, p.Y, st.Background(p.Color), p.Appearance)
	game.Feed.Draw()
	game.Screen.Show()
}

func main() {
	game := NewGame()
	game.Feed.Log("A new game has started!")

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
