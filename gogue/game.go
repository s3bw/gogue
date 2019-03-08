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
	// Screen holds the terminal screen we draw into
	Screen tcell.Screen
	// Level refers to the level at which the active area exists
	Level int
	// ActiveArea refers to the active area to which the player is in.
	ActiveArea *Area
	// Player refers to the user
	Player *entity.Creature
	// stdFeed is the in game feed
	Feed                *feed.Feed
	ShowEquipmentScreen bool

	// We might not have to create a pointer to this, and deploy it
	// only when asked for.. vv
	EquipmentScreen *equipment.Screen
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

// These keys could be consts in a key.go
func IsActionKey(k rune) bool {
	var actionKeys = []rune{
		'k', 'j', 'h', 'l', 'g',
	}
	for _, v := range actionKeys {
		if k == v {
			return true
		}
	}
	return false
}

// PlayerAction contains the actions a player can take
func (game *Game) PlayerAction(p *entity.Creature, key rune) {
	switch key {
	// Moves the player up.
	case 'k':
		game.ActiveArea.MoveCreature(p, 0, -1)
	// Moves the player down.
	case 'j':
		game.ActiveArea.MoveCreature(p, 0, 1)
	// Moves the player left.
	case 'h':
		game.ActiveArea.MoveCreature(p, -1, 0)
	// Moves the player right.
	case 'l':
		game.ActiveArea.MoveCreature(p, 1, 0)
	// Has the player pickup an item they stand on
	case 'g':
		game.ActiveArea.Pickup(p)
	}
}

// Start creates a game instance
func Start() {
	game := NewGame()
	game.Feed.Log("A new game has started!")
	player := game.Player
	turnTaken := false
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
					key := ev.Rune()
					switch {
					// If the key is an action key we have the player take
					// an action, which means they have taken a turn.
					// When implementing equipment screen we might have to
					// use the same keys.
					// (IsActionKey(key) && EquipScreenToggled)
					case IsActionKey(key):
						game.PlayerAction(player, key)
						turnTaken = true
					// Below are non-action keys, these allow the player
					// to do thing that will not cost a turn.
					case key == 'i':
						game.ToggleEquipmentScreen()
					}
				}
			}
			game.Draw()
			if turnTaken {
				// Game steps exist here:
				game.ActiveArea.Turn()
				turnTaken = false
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
