package pacman

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"strings"
	"time"
)

type Game struct {
	screen           tcell.Screen
	state            string
	forbiddenValues  []int
	Pac              PacMan
	ticker           *time.Ticker
	fpsCounterTicker *time.Ticker
	frame            int
}

type PacMan struct {
	playerX       int
	playerY       int
	indexPlayerX  int
	indexPlayerY  int
	pacStyle      tcell.Style
	pacRune       rune
	isMoving      bool
	moveDirection string
	motionTicker  *time.Ticker
}

func buildBoard(layout [30][13]int) string {

	var board strings.Builder

	board.WriteString("\n")
	for _, row := range layout {
		board.WriteRune('|')
		for i := 0; i <= 12; i++ {
			board.WriteRune(writeCharacter(row[i]))
		}
		for i := 12; i >= 0; i-- {
			board.WriteRune(writeCharacter(row[i]))
		}
		board.WriteString("|\n")
	}

	return board.String()
}

func writeCharacter(value int) rune {
	if value == 9 {
		return '&'
	}
	if value == 8 {
		return ' '
	}
	if value == 1 {
		return '•'
	}

	return 'x'
}

func getPacManLocation(layout [30][13]int) (coords [2]int, err error) {
	for i := 0; i < 29; i++ {
		for j := 0; j <= 12; j++ {
			if layout[i][j] == 10 {
				coords = [2]int{j + 1, i + 1}
				err = nil
				return
			}
		}
	}
	return coords, fmt.Errorf("pacman location invalid")
}

func (game *Game) writeBoardToScreen(board string) {
	row := 0
	col := 0

	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	for _, r := range board {
		if r == '\n' {
			row++
			col = 0
			continue
		}
		game.screen.SetContent(col, row, r, nil, style)
		col++
	}
}

func (game *Game) draw() {
	board := buildBoard(gameMap)
	game.writeBoardToScreen(board)
	game.screen.SetContent(game.Pac.playerX, game.Pac.playerY, game.Pac.pacRune, nil, game.Pac.pacStyle)

}

func (game *Game) updateLoop() {
	game.screen.Clear()
	game.HandleMotion()
	game.draw()
	game.screen.Show()
}

func (game *Game) RunGame() {
	for game.state == "active" {
		select {
		case <-game.ticker.C:
			game.updateLoop()
			game.frame++
		case <-game.fpsCounterTicker.C:
			fmt.Printf("\nFPS: %d", game.frame)
			game.frame = 0
		}
	}
}
