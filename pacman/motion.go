package pacman

func (game *Game) moveRight() {
	game.Pac.playerX++
	game.Pac.indexPlayerX++
	game.Pac.pacRune = '<'
}

func (game *Game) moveLeft() {
	game.Pac.playerX--
	game.Pac.indexPlayerX--
	game.Pac.pacRune = '>'
}

func (game *Game) moveUp() {
	game.Pac.playerY--
	game.Pac.indexPlayerY--
	game.Pac.pacRune = '∨'
}

func (game *Game) moveDown() {
	game.Pac.playerY++
	game.Pac.indexPlayerY++
	game.Pac.pacRune = '∧'
}

func (game *Game) HandleMotion() {
	if game.Pac.isMoving {
		select {
		case <-game.Pac.motionTicker.C:
			game.setMoveDirection(game.Pac.moveDirection)
		}
	}
}

func (game *Game) setMoveDirection(direction string) {
	x, y := game.getCoordinates(direction)

	//TODO add warping function to handle pac-man map teleportation
	if !isValidIndex([]int{x, y}) {
		game.Pac.isMoving = false
		return
	}

	containsForbidden := contains(game.forbiddenValues, game.gameMap[y][x])

	if !containsForbidden {
		game.Pac.isMoving = true
		game.Pac.moveDirection = direction
		game.move()
		return
	}

	game.Pac.isMoving = false
	return
}

func (game *Game) getCoordinates(direction string) (x int, y int) {
	switch direction {
	case "left":
		x = getMirrorLocation(game.Pac.indexPlayerX-1, len(game.gameMap[game.Pac.indexPlayerY]))
		y = game.Pac.indexPlayerY
	case "right":
		x = getMirrorLocation(game.Pac.indexPlayerX+1, len(game.gameMap[game.Pac.indexPlayerY]))
		y = game.Pac.indexPlayerY
	case "down":
		x = getMirrorLocation(game.Pac.indexPlayerX, len(game.gameMap[game.Pac.indexPlayerY]))
		y = game.Pac.indexPlayerY + 1
	case "up":
		x = getMirrorLocation(game.Pac.indexPlayerX, len(game.gameMap[game.Pac.indexPlayerY]))
		y = game.Pac.indexPlayerY - 1
	}
	return
}

func (game *Game) move() {
	switch game.Pac.moveDirection {
	case "left":
		game.moveLeft()
	case "right":
		game.moveRight()
	case "up":
		game.moveUp()
	case "down":
		game.moveDown()
	}
}
