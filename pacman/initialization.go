package pacman

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
	"time"
)

func InitializeScreen() tcell.Screen {

	screen, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error creating screen: %v\n", e)
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	return screen
}

func InitializeGame(screen tcell.Screen) (game Game) {

	game.ticker = time.NewTicker(time.Second / 60)
	game.fpsCounterTicker = time.NewTicker(time.Second)
	game.forbiddenValues = []int{9, 8}
	game.setStateActive()
	game.screen = screen
	game.gameMap = gameMap
	game.Pac = game.initializePacMan()

	log.Printf("%v", game)
	return
}

func (game *Game) initializePacMan() (pac PacMan) {
	var pacManLocation, pacManLocationError = getPacManLocation(game.gameMap)

	if pacManLocationError != nil {
		log.Fatal(pacManLocationError)
	}

	pac.motionTicker = time.NewTicker(time.Second / 10)
	pac.playerX = pacManLocation[0]
	pac.playerY = pacManLocation[1]
	pac.indexPlayerX = pacManLocation[0] - 1
	pac.indexPlayerY = pacManLocation[1] - 1
	pac.pacStyle = tcell.StyleDefault.Foreground(tcell.ColorGreen)
	pac.pacRune = '>'
	return
}
