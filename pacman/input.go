package pacman

import "github.com/gdamore/tcell/v2"

func (game *Game) handleUserInput(event *tcell.EventKey) {
	switch event.Key() {
	case tcell.KeyEscape:
		game.setStateClosed()
	case tcell.KeyLeft:
		game.Pac.isMoving = true
		game.Pac.moveDirection = "left"
	case tcell.KeyRight:
		game.Pac.isMoving = true
		game.Pac.moveDirection = "right"
	case tcell.KeyUp:
		game.Pac.isMoving = true
		game.Pac.moveDirection = "up"
	case tcell.KeyDown:
		game.Pac.isMoving = true
		game.Pac.moveDirection = "down"
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
