package pacman

func (game *Game) collectCookies() {
	if game.gameMap[game.Pac.indexPlayerY][game.Pac.indexPlayerX] == 1 {
		game.gameMap[game.Pac.indexPlayerY][game.Pac.indexPlayerX] = 0
	}
}
