/*Package styles defines a set of styles to use for the game.*/
package styles

import (
	"github.com/gdamore/tcell"
)

// DefaultStyle is the default game style
func DefaultStyle() tcell.Style {
	style := tcell.StyleDefault
	// lstyle = style.Background(tcell.Color(tcell.ColorBlack))
	return style
}

func DeadStyle() tcell.Style {
	style := tcell.StyleDefault
	style = style.Foreground(tcell.Color(tcell.ColorDarkRed)).
		Background(tcell.Color(tcell.ColorBlack))
	return style
}
