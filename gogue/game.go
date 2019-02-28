package gogue

import (
	"fmt"
	"os"

	"github.com/foxyblue/gogue/gogue/entity"
	"github.com/foxyblue/gogue/gogue/equipment"
	"github.com/foxyblue/gogue/gogue/feed"
	"github.com/gdamore/tcell"

	// Blank import registers the biomes
	_ "github.com/foxyblue/gogue/gogue/biome/blank"
)

// Game holds the instance of the game
type Game struct {
	Screen tcell.Screen

	// Level refers to the level at which the active area exists
	Level int

	// ActiveArea refers to the active area to which the player is in.
	ActiveArea *Area

	// Player refers to the user
	Player *entity.Creature

	// stdFeed is the in game feed
	Feed *feed.Feed

	ShowEquipmentScreen bool
	EquipmentScreen     *equipment.Screen
}

// NewGame creates a new game instance
func NewGame() *Game {
	level := 0
	playerX, playerY := 10, 10

	screen := newScreen()
	player := entity.NewPlayer(playerX, playerY)
	feed := feed.NewFeed(screen)
	area := NewArea(player, level, screen, feed)
	equipment := equipment.NewEquipmentScreen(screen)
	return &Game{
		Screen:              screen,
		ActiveArea:          area,
		Player:              player,
		Level:               0,
		Feed:                feed,
		ShowEquipmentScreen: false,
		EquipmentScreen:     equipment,
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
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	s.Clear()
	return s
}

// Draw will render the game on screen
func (game *Game) Draw() {
	game.Screen.Clear()
	game.ActiveArea.Draw()
	game.Feed.Draw()
	if game.ShowEquipmentScreen {
		game.EquipmentScreen.Draw()
	}
	game.Screen.Show()
}

func (game *Game) ToggleEquipmentScreen() {
	if game.ShowEquipmentScreen {
		game.ShowEquipmentScreen = false
	}
	game.ShowEquipmentScreen = true
}

// Start creates a game instance
func Start() {
	game := NewGame()
	game.Feed.Log("A new game has started!")
	player := game.Player
	playerTurnTaken := false
	game.Draw()

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
						game.ActiveArea.MoveCreature(player, 0, -1)
						playerTurnTaken = true
					case 'j':
						game.ActiveArea.MoveCreature(player, 0, 1)
						playerTurnTaken = true
					case 'h':
						game.ActiveArea.MoveCreature(player, -1, 0)
						playerTurnTaken = true
					case 'l':
						game.ActiveArea.MoveCreature(player, 1, 0)
						playerTurnTaken = true
					case 'g':
						game.ActiveArea.Pickup(player)
						playerTurnTaken = true
					case 'i':
						game.ToggleEquipmentScreen()
					}
				}
			}
			game.Draw()
			if playerTurnTaken {
				// Game steps exist here:
				game.ActiveArea.Turn()
				playerTurnTaken = false
			}
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
