package pacman

func (game *Game) moveRight() {
	x := getMirrorLocation(game.Pac.indexPlayerX+1, len(gameMap[game.Pac.indexPlayerY]))
	y := game.Pac.indexPlayerY

	game.Pac.pacRune = '<'
	if !isValidIndex([]int{x, y}) {
		game.Pac.isMoving = false
		return
	}

	containsForbidden := contains(game.forbiddenValues, gameMap[y][x])

	if !containsForbidden {
		game.Pac.playerX++
		game.Pac.indexPlayerX++
	} else {
		game.Pac.isMoving = false
	}

}

func (game *Game) moveLeft() {
	x := getMirrorLocation(game.Pac.indexPlayerX-1, len(gameMap[game.Pac.indexPlayerY]))
	y := game.Pac.indexPlayerY

	game.Pac.pacRune = '>'

	//TODO add warping function
	if !isValidIndex([]int{x, y}) {
		game.Pac.isMoving = false
		return
	}

	containsForbidden := contains(game.forbiddenValues, gameMap[y][x])

	if !containsForbidden {
		game.Pac.playerX--
		game.Pac.indexPlayerX--
	} else {

		game.Pac.isMoving = false
	}

}

func (game *Game) moveUp() {
	x := getMirrorLocation(game.Pac.indexPlayerX, len(gameMap[game.Pac.indexPlayerY]))
	y := game.Pac.indexPlayerY - 1

	game.Pac.pacRune = '∨'
	if !isValidIndex([]int{x, y}) {
		game.Pac.isMoving = false
		return
	}

	containsForbidden := contains(game.forbiddenValues, gameMap[y][x])

	if !containsForbidden {
		game.Pac.playerY--
		game.Pac.indexPlayerY--
	} else {
		game.Pac.isMoving = false
	}
}

func (game *Game) moveDown() {
	x := getMirrorLocation(game.Pac.indexPlayerX, len(gameMap[game.Pac.indexPlayerY]))
	y := game.Pac.indexPlayerY + 1

	game.Pac.pacRune = '∧'
	if !isValidIndex([]int{x, y}) {
		game.Pac.isMoving = false
		return
	}

	containsForbidden := contains(game.forbiddenValues, gameMap[y][x])

	if !containsForbidden {
		game.Pac.playerY++
		game.Pac.indexPlayerY++
	} else {
		game.Pac.isMoving = false
	}
}

func (game *Game) HandleMotion() {
	if game.Pac.isMoving {
		select {
		case <-game.Pac.motionTicker.C:
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
	}
}
