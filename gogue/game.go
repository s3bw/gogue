package main

import (
	"fmt"
	"os"

	"github.com/foxyblue/gogue/gogue/area"
	"github.com/foxyblue/gogue/gogue/creature"
	"github.com/foxyblue/gogue/gogue/feed"
	"github.com/gdamore/tcell"

	// Blank import registers the biomes
	_ "github.com/foxyblue/gogue/gogue/area/biome/blank"
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
	playerX, playerY := 10, 10

	screen := newScreen()
	player := creature.NewPlayer(playerX, playerY)
	area := area.NewArea(player, level, screen)
	feed := feed.NewFeed(screen)
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
	game.Feed.Draw()
	game.Screen.Show()
}

func main() {
	game := NewGame()
	game.Feed.Log("A new game has started!")
	grid := game.ActiveArea.Biome.GetGrid()

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
				case tcell.KeyRune:
					switch ev.Rune() {
					case 'k':
						game.Player.Creature.Move(0, -1, grid)
					case 'j':
						game.Player.Creature.Move(0, 1, grid)
					case 'h':
						game.Player.Creature.Move(-1, 0, grid)
					case 'l':
						game.Player.Creature.Move(1, 0, grid)
					}
				}
			}
			// Game steps exist here:
			game.Draw()
			game.Feed.Log("Step")
		}
	}()

	// Main Gameloop
gameLoop:
	for {
		select {
		case <-quit:
			break gameLoop
		}
	}

	game.Screen.Fini()
	fmt.Println("Game has ended.")
}
