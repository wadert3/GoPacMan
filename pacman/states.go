package pacman

func (game *Game) setStatePaused() {
	game.state = "paused"
}

func (game *Game) setStateActive() {
	game.state = "active"
}

func (game *Game) setStateLoading() {
	game.state = "loading"
}

func (game *Game) setStateClosed() {
	game.state = "closed"
}
