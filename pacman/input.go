package pacman

import "github.com/gdamore/tcell/v2"

func (game *Game) handleUserInput(event *tcell.EventKey) {
	switch event.Key() {
	case tcell.KeyEscape:
		game.setStateClosed()
	case tcell.KeyLeft:
		game.setMoveDirection("left")
	case tcell.KeyRight:
		game.setMoveDirection("right")
	case tcell.KeyUp:
		game.setMoveDirection("up")
	case tcell.KeyDown:
		game.setMoveDirection("down")
	}
}

func (game *Game) HandleEvents() {
	for game.state == "active" {
		switch event := game.screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.screen.Sync()
		case *tcell.EventKey:
			game.handleUserInput(event)
		}

	}
}
