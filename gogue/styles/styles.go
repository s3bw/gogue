/*Package styles defines a set of styles to use for the game.*/
package styles

import "github.com/gdamore/tcell"

// DefaultStyle is the default game style
func DefaultStyle() tcell.Style {
	style := tcell.StyleDefault
	style.Foreground(tcell.Color(tcell.ColorBlack))
	return style
}
